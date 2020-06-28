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

// Slide renders a an assignment.
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			ID   int    `json:"id"`
			Mode string `json:"mode"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		var m motion
		if err := ds.Get("motions/motion", element.ID, &m); err != nil {
			return nil, fmt.Errorf("getting motion: %w", err)
		}

		out := map[string]json.RawMessage{
			"identifier":           m.Identifier,
			"title":                m.Title,
			"amendment_paragraphs": m.AmendmentParagraphs,
		}

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
		out["submitters"] = encSubmitters

		isChild := []byte("true")
		if isNull(m.ParentID) {
			isChild = []byte("false")
		}
		out["is_child"] = isChild

		baseStatute := []byte("null")
		if !isNull(m.StatuteParagraphID) {
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
			baseStatute = b
		}
		out["base_statute"] = baseStatute

		baseMotion := []byte("null")
		if isNull(m.StatuteParagraphID) && !isNull(m.ParentID) && !isNull(m.AmendmentParagraphs) {
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
			baseMotion = b
		}
		out["base_motion"] = baseMotion

		var changeRecommendations json.RawMessage
		var amendments json.RawMessage
		if isNull(m.StatuteParagraphID) && !(!isNull(m.ParentID) && !isNull(m.AmendmentParagraphs)) {
			lChangeRecommendations := make([]json.RawMessage, 0)
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

				lChangeRecommendations = append(lChangeRecommendations, cr)
			}

			changeRecommendations, err = json.Marshal(lChangeRecommendations)
			if err != nil {
				return nil, fmt.Errorf("encoding change recommendations: %w", err)
			}

			amendmentData := make([]json.RawMessage, 0)
			for _, id := range m.AmendmentIDs {
				var a amendment
				if err := ds.Get("motions/motion", id, &a); err != nil {
					return nil, fmt.Errorf("getting ammendment: %w", err)
				}

				intoDiff, err := mergedIntoDiff(ds, a)
				if err != nil {
					return nil, fmt.Errorf("calc merged into diff: %w", err)
				}

				intoFinal, err := mergedIntoFinal(ds, a)
				if err != nil {
					return nil, fmt.Errorf("calc merged into final: %w", err)
				}

				out := struct {
					ID         int             `json:"id"`
					Identifier json.RawMessage `json:"identifer"`
					Title      json.RawMessage `json:"title"`
					Paragraphs json.RawMessage `json:"amendment_paragraphs"`
					IntoDiff   int             `json:"merge_amendment_into_diff"`
					IntoFinal  int             `json:"merge_amendment_into_final"`
				}{
					id,
					a.Identifier,
					a.Title,
					a.Paragraphs,
					intoDiff,
					intoFinal,
				}
				b, err := json.Marshal(out)
				if err != nil {
					return nil, fmt.Errorf("ecoding amentment: %w", err)
				}
				amendmentData = append(amendmentData, b)
			}
			amendments, err = json.Marshal(amendmentData)
			if err != nil {
				return nil, fmt.Errorf("encoding amentment list: %w", err)
			}

		}
		out["change_recommendations"] = changeRecommendations
		out["amendments"] = amendments

		mode := element.Mode
		if mode == "" {
			if err := ds.ConfigValue("motions_recommendation_text_mode", &mode); err != nil {
				return nil, fmt.Errorf("get config value for motions_recommendation_text_mode: %w", err)
			}
		}
		if mode == "final" {
			out["modified_final_version"] = m.ModifiedFinalVersion
		}

		var buf json.RawMessage
		if err := ds.ConfigValue("motions_disable_sidebox_on_projector", &buf); err != nil {
			return nil, fmt.Errorf("getting motions_disable_sidebox_on_projector: %w", err)
		}
		out["show_meta_box"] = buf

		buf = buf[0:0:0]
		if err := ds.ConfigValue("motions_line_length", &buf); err != nil {
			return nil, fmt.Errorf("getting motions_line_length: %w", err)
		}
		out["line_length"] = buf

		buf = buf[0:0:0]
		if err := ds.ConfigValue("motions_preamble", &buf); err != nil {
			return nil, fmt.Errorf("getting motions_preamble: %w", err)
		}
		out["preamble"] = buf

		buf = buf[0:0:0]
		if err := ds.ConfigValue("motions_default_line_numbering", &buf); err != nil {
			return nil, fmt.Errorf("getting motions_default_line_numbering: %w", err)
		}
		out["line_numbering_mode"] = buf

		var b bool
		if err := ds.ConfigValue("motions_disable_text_on_projector", &b); err != nil {
			return nil, fmt.Errorf("getting motions_disable_text_on_projector: %w", err)
		}
		if !b {
			out["text"] = m.Text
		}

		if err := ds.ConfigValue("motions_hide_referring_motions", &b); err != nil {
			return nil, fmt.Errorf("getting motions_hide_referring_motions: %w", err)
		}
		showReferringMotions := !b
		out["show_referring_motions"] = []byte(`false`)
		if showReferringMotions {
			out["show_referring_motions"] = []byte(`true`)
			b, err := recommendationReferencingMotions(ds, m.ID)
			if err != nil {
				return nil, fmt.Errorf("getting recommendation reflecting motions: %w", err)
			}
			out["recommendation_referencing_motions"] = b
		}

		if err := ds.ConfigValue("motions_disable_reason_on_projector", &b); err != nil {
			return nil, fmt.Errorf("getting motions_disable_reason_on_projector: %w", err)
		}
		if !b {
			out["reason"] = m.Reason
		}

		if err := ds.ConfigValue("motions_disable_recommendation_on_projector", &b); err != nil {
			return nil, fmt.Errorf("getting motions_disable_recommendation_on_projector: %w", err)
		}

		if !b && !m.RecommendationID.Null() {
			var state struct {
				Label         json.RawMessage `json:"recommendation_label"`
				ShowExtension bool            `json:"show_recommendation_extension_field"`
			}
			if err := ds.Get("motions/state", m.RecommendationID.Value(), &state); err != nil {
				return nil, fmt.Errorf("getting state: %w", err)
			}
			out["recommendation"] = state.Label

			if state.ShowExtension {
				extension, err := extendReferenceMotions(ds, m.RecommendationExtension)
				if err != nil {
					return nil, fmt.Errorf("getting extension: %w", err)
				}
				out["recommendation_extension"] = m.RecommendationExtension
				out["referenced_motions"] = extension
			}

			recommenderConfig := "motions_statute_recommendations_by"
			if isNull(m.StatuteParagraphID) {
				recommenderConfig = "motions_recommendations_by"
			}
			if err := ds.ConfigValue(recommenderConfig, &buf); err != nil {
				return nil, fmt.Errorf("getting %s: %w", recommenderConfig, err)
			}
			out["recommender"] = buf
		}

		for k, v := range out {
			fmt.Println(k, string(v))
		}

		encoded, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encode motion data: %w", err)
		}
		return encoded, nil
	}
}

func recommendationReferencingMotions(ds projector.Datastore, motionID int) (json.RawMessage, error) {
	var v []json.RawMessage
	for _, rm := range ds.GetCollection("motions/motion") {
		var m motion
		if err := json.Unmarshal(rm, &m); err != nil {
			return nil, fmt.Errorf("decoding motion: %w", err)
		}

		if m.RecommendationID.Null() || isNull(m.RecommendationExtension) {
			continue
		}

		var state struct {
			ShowExtension bool `json:"show_recommendation_extension_field"`
		}
		if err := ds.Get("motions/state", m.RecommendationID.Value(), &state); err != nil {
			// TODO: if state does not exist, better error message.
			return nil, fmt.Errorf("getting state: %w", err)
		}
		if !state.ShowExtension {
			continue
		}

		r := regexp.MustCompile(`\[motion:(\d+)\]`)
		ids := make(map[int]bool)
		for _, match := range r.FindAllSubmatch(m.RecommendationExtension, -1) {
			id, err := strconv.Atoi(string(match[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid id: %w", err)
			}
			ids[id] = true
		}

		if ids[motionID] {
			out := struct {
				Title      json.RawMessage `json:"title"`
				Identifier json.RawMessage `json:"identifier"`
			}{
				m.Title,
				m.Identifier,
			}

			b, err := json.Marshal(out)
			if err != nil {
				return nil, fmt.Errorf("encode item: %w", err)
			}
			v = append(v, b)
		}
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("encoding list: %w", err)
	}
	return b, nil
}

func extendReferenceMotions(ds projector.Datastore, recommendation json.RawMessage) (json.RawMessage, error) {
	if isNull(recommendation) {
		return []byte(`{}`), nil
	}

	r := regexp.MustCompile(`\[motion:(\d+)\]`)
	recommendated := make(map[int]json.RawMessage)
	for _, match := range r.FindAllSubmatch(recommendation, -1) {
		id, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid id: %w", err)
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
			return nil, fmt.Errorf("getting recommendated motion: %w", err)
		}

		b, err := json.Marshal(motion)
		if err != nil {
			return nil, fmt.Errorf("encoding motion: %w", err)
		}
		recommendated[id] = b
	}
	b, err := json.Marshal(recommendated)
	if err != nil {
		return nil, fmt.Errorf("encoding recommendation list: %w", err)
	}
	return b, nil
}

func mergedIntoDiff(ds projector.Datastore, a amendment) (int, error) {
	if a.StateID.Null() {
		return 0, nil
	}

	var state struct {
		IntoFinal int `json:"merge_amendment_into_final"`
	}
	if err := ds.Get("motions/state", a.StateID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}

	if state.IntoFinal == -1 {
		return 0, nil
	}

	if state.IntoFinal == 1 {
		return 1, nil
	}

	if a.RecommendationID.Null() {
		return 0, nil
	}

	if err := ds.Get("motion/state", a.RecommendationID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}
	if state.IntoFinal == 1 {
		return 1, nil
	}
	return 0, nil
}

func mergedIntoFinal(ds projector.Datastore, a amendment) (int, error) {
	if a.StateID.Null() {
		return 0, nil
	}

	var state struct {
		IntoFinal int `json:"merge_amendment_into_final"`
	}
	if err := ds.Get("motions/state", a.StateID.Value(), &state); err != nil {
		// TODO: if state does not exist, better error message.
		return 0, fmt.Errorf("getting state: %w", err)
	}

	if state.IntoFinal == 1 {
		return 1, nil
	}
	return 0, nil
}

func isNull(b []byte) bool {
	return bytes.Equal(b, []byte("null"))
}
