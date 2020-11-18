package motion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders a a motion.
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			ID   int    `json:"id"`
			Mode string `json:"mode"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			var element struct {
				ID interface{} `json:"id"`
			}

			// Fallback for better error messages
			if err := json.Unmarshal(e, &element); err == nil {
				return nil, projector.NewClientError("motions/motion with id %s does not exist", element.ID)
			}
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		if element.ID == 0 {
			return nil, projector.NewClientError("id is required for motions/motion slide")
		}

		var m motion
		if err := ds.Get("motions/motion", element.ID, &m); err != nil {
			return nil, projector.NewClientError("motions/motion with id %d does not exist", element.ID)
		}

		submitters, err := slideSubmitters(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get motion submitters: %w", err)
		}

		isChild := jsonBool(!isNull(m.ParentID))

		baseStatute, err := slideBaseStatute(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get motion base statute: %w", err)
		}

		baseMotion, err := slideBaseMotion(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get base motion: %w", err)
		}

		changeRecommendations, err := slideChangeRecommendations(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get change recommendations: %w", err)
		}

		amendments, err := slideAmendments(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get amendments: %w", err)
		}

		var showMetaBox json.RawMessage
		if err := ds.ConfigValue("motions_disable_sidebox_on_projector", &showMetaBox); err != nil {
			return nil, fmt.Errorf("getting motions_disable_sidebox_on_projector: %w", err)
		}
		showMetaBox = jsonNot(showMetaBox)

		var lineLength json.RawMessage
		if err := ds.ConfigValue("motions_line_length", &lineLength); err != nil {
			return nil, fmt.Errorf("getting motions_line_length: %w", err)
		}

		var preamble json.RawMessage
		if err := ds.ConfigValue("motions_preamble", &preamble); err != nil {
			return nil, fmt.Errorf("getting motions_preamble: %w", err)
		}

		var lineNumberingMode json.RawMessage
		if err := ds.ConfigValue("motions_default_line_numbering", &lineNumberingMode); err != nil {
			return nil, fmt.Errorf("getting motions_default_line_numbering: %w", err)
		}

		showReferingMotions, recommendationReferencingMotions, err := slideReferringMotions(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get refering motions: %w", err)
		}

		out := map[string]json.RawMessage{
			"identifier":             m.Identifier,
			"title":                  m.Title,
			"amendment_paragraphs":   m.AmendmentParagraphs,
			"submitters":             submitters,
			"is_child":               isChild,
			"base_statute":           baseStatute,
			"base_motion":            baseMotion,
			"change_recommendations": changeRecommendations,
			"amendments":             amendments,
			"show_meta_box":          showMetaBox,
			"line_length":            lineLength,
			"preamble":               preamble,
			"line_numbering_mode":    lineNumberingMode,
			"show_referring_motions": jsonBool(showReferingMotions),
		}

		if showReferingMotions {
			out["recommendation_referencing_motions"] = recommendationReferencingMotions
		}

		mode := element.Mode
		if mode == "" {
			if err := ds.ConfigValue("motions_recommendation_text_mode", &mode); err != nil {
				return nil, fmt.Errorf("get config value for motions_recommendation_text_mode: %w", err)
			}
		}
		if mode == "final" {
			out["modified_final_version"] = m.ModifiedFinalVersion
		}

		showText, err := slideShowText(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get show text: %w", err)
		}
		if showText {
			out["text"] = m.Text
		}

		showReason, err := slideShowReason(ds, &m)
		if err != nil {
			return nil, fmt.Errorf("get show reason: %w", err)
		}
		if showReason {
			out["reason"] = m.Reason
		}

		if err := slideRecommendation(out, ds, &m); err != nil {
			return nil, fmt.Errorf("get recommendation: %w", err)
		}

		encoded, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encode motion data: %w", err)
		}
		return encoded, nil
	}
}

// SlideMotionBlock renders a an motion block.
func SlideMotionBlock() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var mb struct {
			Motions []int           `json:"motions_id"`
			Title   json.RawMessage `json:"title"`
		}
		if err := projector.ModelFromElement(ds, e, "motions/motion-block", &mb); err != nil {
			return nil, fmt.Errorf("getting motion block: %w", err)
		}

		var motions []json.RawMessage
		referenced := make(map[int]json.RawMessage)
		for _, mID := range mb.Motions {
			var m motion
			if err := ds.Get("motions/motion", mID, &m); err != nil {
				return nil, fmt.Errorf("getting motion: %w", err)
			}

			out := map[string]json.RawMessage{
				"title":      m.Title,
				"identifier": m.Identifier,
			}

			if !m.RecommendationID.Null() {
				var recommendation struct {
					Label              json.RawMessage `json:"recommendation_label"`
					CSS                json.RawMessage `json:"css_class"`
					ShowExtensionField bool            `json:"show_recommendation_extension_field"`
				}
				if err := ds.Get("motions/state", m.RecommendationID.Value(), &recommendation); err != nil {
					return nil, fmt.Errorf("getting recommendation state: %w", err)
				}

				var outRecommendation = struct {
					Name json.RawMessage `json:"name"`
					CSS  json.RawMessage `json:"css_class"`
				}{
					recommendation.Label,
					recommendation.CSS,
				}

				bs, err := json.Marshal(outRecommendation)
				if err != nil {
					return nil, fmt.Errorf("encoding recommendation: %w", err)
				}
				out["recommendation"] = bs

				if recommendation.ShowExtensionField {
					out["recommendation_extension"] = m.RecommendationExtension
					if err := extendReferenceMotions(ds, m.RecommendationExtension, referenced); err != nil {
						return nil, fmt.Errorf("getting extension: %w", err)
					}
				}
			}
			bs, err := json.Marshal(out)
			if err != nil {
				return nil, fmt.Errorf("encoding motion: %w", err)
			}
			motions = append(motions, bs)
		}

		var out = struct {
			Title      json.RawMessage         `json:"title"`
			Motions    []json.RawMessage       `json:"motions"`
			Referenced map[int]json.RawMessage `json:"referenced_motions"`
		}{
			mb.Title,
			motions,
			referenced,
		}
		bs, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding motion block: %w", err)
		}
		return bs, nil
	}
}

// SlideMotionPoll renders slides for motions/motion-poll.
func SlideMotionPoll() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var mp map[string]json.RawMessage
		if err := projector.ModelFromElement(ds, e, "motions/motion-poll", &mp); err != nil {
			return nil, fmt.Errorf("getting motion poll: %w", err)
		}

		var mid int
		if err := json.Unmarshal(mp["motion_id"], &mid); err != nil {
			return nil, fmt.Errorf("decoding motion id: %w", err)
		}

		pollData := map[string]json.RawMessage{
			"title":                   mp["title"],
			"type":                    mp["type"],
			"pollmethod":              mp["pollmethod"],
			"state":                   mp["state"],
			"onehundred_percent_base": mp["onehundred_percent_base"],
			"majority_method":         mp["majority_method"],
		}

		if bytes.Equal(mp["state"], []byte("4")) {
			var oids []int
			if err := json.Unmarshal(mp["options_id"], &oids); err != nil {
				return nil, fmt.Errorf("decoding options_id: %w", err)
			}

			var rawOption struct {
				Yes     string `json:"yes"`
				No      string `json:"no"`
				Abstain string `json:"abstain"`
			}
			if err := ds.Get("motions/motion-option", oids[0], &rawOption); err != nil {
				return nil, fmt.Errorf("getting motion-option: %w", err)
			}

			yes, err := strconv.ParseFloat(rawOption.Yes, 32)
			if err != nil {
				return nil, fmt.Errorf("decoding option.yes: %w", err)
			}
			no, err := strconv.ParseFloat(rawOption.No, 32)
			if err != nil {
				return nil, fmt.Errorf("decoding option.no: %w", err)
			}
			abstain, err := strconv.ParseFloat(rawOption.Abstain, 32)
			if err != nil {
				return nil, fmt.Errorf("decoding option.abstain: %w", err)
			}

			outOption := struct {
				Yes     float64 `json:"yes"`
				No      float64 `json:"no"`
				Abstain float64 `json:"abstain"`
			}{
				yes,
				no,
				abstain,
			}
			t := []interface{}{outOption}

			b, err := json.Marshal(t)
			if err != nil {
				return nil, fmt.Errorf("encoding motion option: %w", err)
			}
			pollData["options"] = b
			pollData["votesvalid"] = mp["votesvalid"]
			pollData["votesinvalid"] = mp["votesinvalid"]
			pollData["votescast"] = mp["votescast"]
		}

		var m motion
		if err := ds.Get("motions/motion", mid, &m); err != nil {
			return nil, fmt.Errorf("getting motion: %w", err)
		}
		out := struct {
			Motion struct {
				Title      json.RawMessage `json:"title"`
				Identifier json.RawMessage `json:"identifier"`
			} `json:"motion"`
			Poll map[string]json.RawMessage `json:"poll"`
		}{
			struct {
				Title      json.RawMessage `json:"title"`
				Identifier json.RawMessage `json:"identifier"`
			}{
				m.Title,
				m.Identifier,
			},
			pollData,
		}
		b, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding data: %w", err)
		}
		return b, nil
	}
}

func extendReferenceMotions(ds projector.Datastore, recommendation json.RawMessage, recommendated map[int]json.RawMessage) error {
	if isNull(recommendation) {
		return nil
	}

	r := regexp.MustCompile(`\[motion:(\d+)\]`)
	for _, match := range r.FindAllSubmatch(recommendation, -1) {
		id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}
		var motion struct {
			Title      json.RawMessage `json:"title"`
			Identifier json.RawMessage `json:"identifier"`
		}
		if err := ds.Get("motions/motion", id, &motion); err != nil {
			var doesNotExist interface {
				DoesNotExist()
			}
			if errors.As(err, &doesNotExist) {
				continue
			}
			return fmt.Errorf("getting recommendated motion: %w", err)
		}

		b, err := json.Marshal(motion)
		if err != nil {
			return fmt.Errorf("encoding motion: %w", err)
		}
		recommendated[id] = b
	}

	return nil
}

func mergedIntoDiff(ds projector.Datastore, m *motion) (int, error) {
	if m.StateID.Null() {
		return 0, nil
	}

	var state struct {
		IntoFinal int `json:"merge_amendment_into_final"`
	}
	if err := ds.Get("motions/state", m.StateID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}

	if state.IntoFinal == -1 {
		return 0, nil
	}

	if state.IntoFinal == 1 {
		return 1, nil
	}

	if m.RecommendationID.Null() {
		return 0, nil
	}

	if err := ds.Get("motion/state", m.RecommendationID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}
	if state.IntoFinal == 1 {
		return 1, nil
	}
	return 0, nil
}

func mergedIntoFinal(ds projector.Datastore, m *motion) (int, error) {
	if m.StateID.Null() {
		return 0, nil
	}

	var state struct {
		IntoFinal int `json:"merge_amendment_into_final"`
	}
	if err := ds.Get("motions/state", m.StateID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}

	if state.IntoFinal == 1 {
		return 1, nil
	}
	return 0, nil
}

func slideSubmitters(ds projector.Datastore, m *motion) (json.RawMessage, error) {
	sort.Slice(m.Submitters, func(i, j int) bool {
		return m.Submitters[i].Weight < m.Submitters[j].Weight
	})
	submitters := make([]string, len(m.Submitters))
	for i, s := range m.Submitters {
		username, err := user.GetUserName(ds, s.UserID)
		if err != nil {
			return nil, fmt.Errorf("getting username: %w", err)
		}
		submitters[i] = username
	}
	encSubmitters, err := json.Marshal(submitters)
	if err != nil {
		return nil, fmt.Errorf("encode submitters: %w", err)
	}
	return encSubmitters, nil
}

func slideBaseStatute(ds projector.Datastore, m *motion) (json.RawMessage, error) {
	if isNull(m.StatuteParagraphID) {
		return []byte("null"), nil
	}

	var statute struct {
		Title json.RawMessage `json:"title"`
		Text  json.RawMessage `json:"text"`
	}
	if err := ds.Get("motions/statute-paragraph", m.ID, &statute); err != nil {
		return nil, fmt.Errorf("getting statute paragraph: %w", err)
	}
	b, err := json.Marshal(statute)
	if err != nil {
		return nil, fmt.Errorf("decoding statute: %w", err)
	}
	return b, nil
}

func slideBaseMotion(ds projector.Datastore, m *motion) (json.RawMessage, error) {
	if !isNull(m.StatuteParagraphID) || isNull(m.ParentID) || isNull(m.AmendmentParagraphs) {
		return []byte("null"), nil
	}

	var pid int
	if err := json.Unmarshal(m.ParentID, &pid); err != nil {
		return nil, fmt.Errorf("decoding parent id: %w", err)
	}

	var parent struct {
		Identifier json.RawMessage `json:"identifier"`
		Title      json.RawMessage `json:"title"`
		Text       json.RawMessage `json:"text"`
	}
	if err := ds.Get("motions/motion", pid, &parent); err != nil {
		return nil, fmt.Errorf("getting parent motion: %w", err)
	}
	b, err := json.Marshal(parent)
	if err != nil {
		return nil, fmt.Errorf("decoding parent: %w", err)
	}
	return b, nil

}

func slideChangeRecommendations(ds projector.Datastore, m *motion) (json.RawMessage, error) {
	var crs []json.RawMessage
	for _, id := range m.ChangeRecommendationIDs {
		var cr json.RawMessage
		if err := ds.Get("motions/motion-change-recommendation", id, &cr); err != nil {
			var doesNotExist interface {
				DoesNotExist()
			}
			if errors.As(err, &doesNotExist) {
				continue
			}
			return nil, fmt.Errorf("decoding change recommendation: %w", err)
		}

		var crInternal struct {
			Internal bool `json:"internal"`
		}
		if err := json.Unmarshal(cr, &crInternal); err != nil {
			return nil, fmt.Errorf("decoding internal value of change recommendation: %w", err)
		}

		if crInternal.Internal {
			continue
		}

		crs = append(crs, cr)
	}

	if len(crs) == 0 {
		return []byte("[]"), nil
	}

	c, err := json.Marshal(crs)
	if err != nil {
		return nil, fmt.Errorf("encoding change recommendations: %w", err)
	}

	return c, nil
}

func slideAmendments(ds projector.Datastore, m *motion) (json.RawMessage, error) {
	if !isNull(m.StatuteParagraphID) || (!isNull(m.ParentID) && !isNull(m.AmendmentParagraphs)) {
		return []byte("[]"), nil
	}

	var amendmentData []json.RawMessage
	for _, id := range m.AmendmentIDs {
		var amendment motion
		if err := ds.Get("motions/motion", id, &amendment); err != nil {
			return nil, fmt.Errorf("getting ammendment: %w", err)
		}

		// Get change recommendations for the amendment.
		crs, err := slideChangeRecommendations(ds, &amendment)
		if err != nil {
			return nil, fmt.Errorf("get change recommendations for amendment: %w", err)
		}

		intoDiff, err := mergedIntoDiff(ds, &amendment)
		if err != nil {
			return nil, fmt.Errorf("calc merged into diff: %w", err)
		}

		intoFinal, err := mergedIntoFinal(ds, &amendment)
		if err != nil {
			return nil, fmt.Errorf("calc merged into final: %w", err)
		}

		out := struct {
			ID                    int             `json:"id"`
			Identifier            json.RawMessage `json:"identifier"`
			Title                 json.RawMessage `json:"title"`
			Paragraphs            json.RawMessage `json:"amendment_paragraphs"`
			ChangeRecommendations json.RawMessage `json:"change_recommendations"`
			IntoDiff              int             `json:"merge_amendment_into_diff"`
			IntoFinal             int             `json:"merge_amendment_into_final"`
		}{
			id,
			amendment.Identifier,
			amendment.Title,
			amendment.AmendmentParagraphs,
			crs,
			intoDiff,
			intoFinal,
		}
		b, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("ecoding amentment: %w", err)
		}
		amendmentData = append(amendmentData, b)
	}

	if len(amendmentData) == 0 {
		return []byte("[]"), nil
	}

	amendments, err := json.Marshal(amendmentData)
	if err != nil {
		return nil, fmt.Errorf("encoding amentment list: %w", err)
	}

	return amendments, nil
}

func slideShowText(ds projector.Datastore, m *motion) (bool, error) {
	var b bool
	if err := ds.ConfigValue("motions_disable_text_on_projector", &b); err != nil {
		return false, fmt.Errorf("getting motions_disable_text_on_projector: %w", err)
	}
	return !b, nil
}

func slideShowReason(ds projector.Datastore, m *motion) (bool, error) {
	var b bool
	if err := ds.ConfigValue("motions_disable_reason_on_projector", &b); err != nil {
		return false, fmt.Errorf("getting motions_disable_reason_on_projector: %w", err)
	}
	return !b, nil
}

func slideReferringMotions(ds projector.Datastore, m *motion) (bool, json.RawMessage, error) {
	var hideRMotion bool
	if err := ds.ConfigValue("motions_hide_referring_motions", &hideRMotion); err != nil {
		return false, nil, fmt.Errorf("getting motions_hide_referring_motions: %w", err)
	}

	if hideRMotion {
		return false, nil, nil
	}

	var v []json.RawMessage
	for _, rm := range ds.GetCollection("motions/motion") {
		var im motion
		if err := json.Unmarshal(rm, &im); err != nil {
			return false, nil, fmt.Errorf("decoding motion: %w", err)
		}

		if im.RecommendationID.Null() || isNull(im.RecommendationExtension) {
			continue
		}

		var state struct {
			ShowExtension bool `json:"show_recommendation_extension_field"`
		}
		if err := ds.Get("motions/state", im.RecommendationID.Value(), &state); err != nil {
			// TODO: if state does not exist, better error message.
			return false, nil, fmt.Errorf("getting state: %w", err)
		}
		if !state.ShowExtension {
			continue
		}

		r := regexp.MustCompile(`\[motion:(\d+)\]`)
		ids := make(map[int]bool)
		for _, match := range r.FindAllSubmatch(im.RecommendationExtension, -1) {
			id, err := strconv.Atoi(string(match[1]))
			if err != nil {
				return false, nil, fmt.Errorf("invalid id: %w", err)
			}
			ids[id] = true
		}

		if ids[m.ID] {
			out := struct {
				Title      json.RawMessage `json:"title"`
				Identifier json.RawMessage `json:"identifier"`
			}{
				im.Title,
				im.Identifier,
			}

			b, err := json.Marshal(out)
			if err != nil {
				return false, nil, fmt.Errorf("encode item: %w", err)
			}
			v = append(v, b)
		}
	}
	r, err := json.Marshal(v)
	if err != nil {
		return false, nil, fmt.Errorf("encoding list: %w", err)
	}
	return true, r, nil
}

func slideRecommendation(out map[string]json.RawMessage, ds projector.Datastore, m *motion) error {
	var disableRecommendation bool
	if err := ds.ConfigValue("motions_disable_recommendation_on_projector", &disableRecommendation); err != nil {
		return fmt.Errorf("getting motions_disable_recommendation_on_projector: %w", err)
	}

	if disableRecommendation || m.RecommendationID.Null() {
		return nil
	}

	var state struct {
		Label         json.RawMessage `json:"recommendation_label"`
		ShowExtension bool            `json:"show_recommendation_extension_field"`
	}
	if err := ds.Get("motions/state", m.RecommendationID.Value(), &state); err != nil {
		return fmt.Errorf("getting state: %w", err)
	}
	out["recommendation"] = state.Label

	if state.ShowExtension {
		extension := make(map[int]json.RawMessage)
		if err := extendReferenceMotions(ds, m.RecommendationExtension, extension); err != nil {
			return fmt.Errorf("getting extension: %w", err)
		}
		bs, err := json.Marshal(extension)
		if err != nil {
			return fmt.Errorf("encoding extension: %w", err)
		}
		out["recommendation_extension"] = m.RecommendationExtension
		out["referenced_motions"] = bs
	}

	recommenderConfig := "motions_statute_recommendations_by"
	if isNull(m.StatuteParagraphID) {
		recommenderConfig = "motions_recommendations_by"
	}
	var recommender json.RawMessage
	if err := ds.ConfigValue(recommenderConfig, &recommender); err != nil {
		return fmt.Errorf("getting %s: %w", recommenderConfig, err)
	}
	out["recommender"] = recommender
	return nil
}

func isNull(b []byte) bool {
	return bytes.Equal(b, []byte("null"))
}

func jsonNot(v json.RawMessage) json.RawMessage {
	t := []byte("true")
	f := []byte("false")
	if bytes.Equal(v, t) {
		return f
	}
	if bytes.Equal(v, f) {
		return t
	}
	return v
}

func jsonBool(b bool) json.RawMessage {
	if b {
		return []byte("true")
	}
	return []byte("false")
}
