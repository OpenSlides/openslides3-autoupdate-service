package agenda

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

type treeInfoDepth struct {
	treeInfo
	Depth int `json:"depth"`
}

type treeInfo struct {
	TitleInformation map[string]*projector.OptionalStr `json:"title_information"`
	Collection       string                            `json:"collection"`
}

// ItemListSlide renders a list of items.
func ItemListSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			OnlyMainItems bool `json:"only_main_items"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		rItems := ds.GetCollection("agenda/item")

		items := make(map[int]agendaItem)
		for _, rItem := range rItems {
			var item agendaItem
			if err := json.Unmarshal(rItem, &item); err != nil {
				return nil, fmt.Errorf("decoding item: %w", err)
			}
			if item.ID == 0 {
				return nil, fmt.Errorf("item with no id")
			}

			items[item.ID] = item
		}
		var out json.RawMessage
		if element.OnlyMainItems {
			var agendaItems []treeInfo
			for _, item := range sortedItems(items) {
				if item.ParentID != 0 || item.Type != 1 {
					continue
				}
				titleInformation := item.TitleInformation
				titleInformation["_agenda_item_number"] = projector.NewOptionalStr(item.ItemNumber)
				agendaItems = append(agendaItems, treeInfo{
					TitleInformation: titleInformation,
					Collection:       item.ContentObject.Collection,
				})
			}
			var err error
			out, err = json.Marshal(agendaItems)
			if err != nil {
				return nil, fmt.Errorf("decoding agenda items: %w", err)
			}
		} else {
			var err error
			agendaItems, err := getFlatTree(items, 0)
			if err != nil {
				return nil, fmt.Errorf("get flat tree: %w", err)
			}
			out, err = json.Marshal(agendaItems)
			if err != nil {
				return nil, fmt.Errorf("decoding agenda items: %w", err)
			}
		}

		return []byte(fmt.Sprintf(`{"items":%s}`, out)), nil
	}
}

func sortedItems(items map[int]agendaItem) []agendaItem {
	itemList := make([]agendaItem, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}

	sort.Slice(itemList, func(i, j int) bool {
		if itemList[i].Weight == itemList[j].Weight {
			return itemList[i].ID < itemList[j].ID
		}
		return itemList[i].Weight < itemList[j].Weight
	})

	return itemList
}

func getFlatTree(agendaItems map[int]agendaItem, parentID int) ([]treeInfoDepth, error) {
	children := make(map[int][]int)
	for _, item := range sortedItems(agendaItems) {
		if item.Type == 1 {
			children[item.ParentID] = append(children[item.ParentID], item.ID)
		}
	}

	var tree []treeInfoDepth
	var buildTree func(itemIDs []int, depth int)
	buildTree = func(itemIDs []int, depth int) {
		for _, itemID := range itemIDs {
			item := agendaItems[itemID]
			titleInformation := item.TitleInformation
			titleInformation["_agenda_item_number"] = projector.NewOptionalStr(item.ItemNumber)
			tree = append(tree, treeInfoDepth{
				treeInfo: treeInfo{
					TitleInformation: titleInformation,
					Collection:       item.ContentObject.Collection},
				Depth: depth,
			})
			buildTree(children[itemID], depth+1)
		}
	}
	buildTree(children[parentID], 0)

	return tree, nil
}

// ListOfSpeakersSlide renders a list of speakers.
func ListOfSpeakersSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var los listOfSpeakers
		projector.ModelFromElement(ds, e, "agenda/list-of-speakers", &los)

		var element struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			var element struct {
				ID interface{} `json:"id"`
			}

			// Fallback for better error messages
			if err := json.Unmarshal(e, &element); err == nil {
				return nil, projector.NewClientError("agenda/list-of-speakers with id %s does not exist", element.ID)
			}
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		if element.ID == 0 {
			return nil, projector.NewClientError("id is required for list of speakers slide")
		}

		if err := ds.Get("agenda/list-of-speakers", element.ID, &los); err != nil {
			return nil, projector.NewClientError("agenda/list-of-speakers with id %d does not exist", element.ID)
		}

		return listOfSpeakerSlideData(ds, los)
	}
}

type formattedSpeaker struct {
	User         json.RawMessage         `json:"user"`
	Marked       json.RawMessage         `json:"marked"`
	PointOfOrder bool                    `json:"point_of_order"`
	ProSpeech    *projector.OptionalBool `json:"pro_speech"`
	Weight       *projector.OptionalInt  `json:"weight"`
	EndTime      json.RawMessage         `json:"end_time"`
}

func titleInformation(ds projector.Datastore, los listOfSpeakers) (map[string]*projector.OptionalStr, error) {
	titleInformation := los.TitleInformation
	var contentObject struct {
		AgendaItemID int `json:"agenda_item_id"`
	}
	if err := ds.Get(los.ContentObject.Collection, los.ContentObject.ID, &contentObject); err != nil {
		var doesNotExist interface {
			DoesNotExist() string
		}
		if errors.As(err, &doesNotExist) {
			// Content Object does not exist.
			return titleInformation, nil
		}
		return nil, fmt.Errorf("getting content object: %w", err)
	}
	// contentObject does exist.
	// Set titleInformation["_agenda_item_number"].
	var item struct {
		Number string `json:"item_number"`
	}
	if err := ds.Get("agenda/item", contentObject.AgendaItemID, &item); err != nil {
		var doesNotExist interface {
			DoesNotExist() string
		}
		if errors.As(err, &doesNotExist) {
			// Agenda item does not exist.
			return titleInformation, nil
		}
		return nil, fmt.Errorf("getting content object: %w", err)
	}
	titleInformation["_agenda_item_number"] = projector.NewOptionalStr(item.Number)
	return titleInformation, nil
}

func listOfSpeakerSlideData(ds projector.Datastore, los listOfSpeakers) (json.RawMessage, error) {
	var showLastSpeakers int
	if err := ds.ConfigValue("agenda_show_last_speakers", &showLastSpeakers); err != nil {
		return nil, fmt.Errorf("loading agenda_show_last_speakers: %w", err)
	}

	var speakersWaiting []formattedSpeaker = make([]formattedSpeaker, 0)
	var speakersFinished []formattedSpeaker
	var currentSpeaker *formattedSpeaker
	for _, speaker := range los.Speakers {
		username, err := user.GetUserName(ds, speaker.UserID)
		if err != nil {
			return nil, fmt.Errorf("getting username: %w", err)
		}
		fs := formattedSpeaker{
			User:         username,
			Marked:       speaker.Marked,
			PointOfOrder: speaker.PointOfOrder,
			ProSpeech:    speaker.ProSpeech,
			Weight:       speaker.Weight,
			EndTime:      speaker.EndTime,
		}

		if bytes.Equal(speaker.BeginTime, []byte("null")) && bytes.Equal(speaker.EndTime, []byte("null")) {
			speakersWaiting = append(speakersWaiting, fs)
			continue
		}

		if bytes.Equal(speaker.EndTime, []byte("null")) {
			currentSpeaker = &fs
			continue
		}

		if showLastSpeakers > 0 {
			speakersFinished = append(speakersFinished, fs)
		}
	}

	sort.Slice(speakersWaiting, func(i, j int) bool {
		return speakersWaiting[i].Weight.Value() < speakersWaiting[j].Weight.Value()
	})

	if showLastSpeakers > 0 {
		sort.Slice(speakersFinished, func(i, j int) bool {
			return speakersFinished[i].Weight.Value() < speakersFinished[j].Weight.Value()
		})

		l := len(speakersFinished) - showLastSpeakers
		if l > 0 {
			speakersFinished = speakersFinished[l:]
		}
	}

	var showNextSpeakers int
	if err := ds.ConfigValue("agenda_show_next_speakers", &showNextSpeakers); err != nil {
		return nil, fmt.Errorf("loading agenda_show_next_speakers: %w", err)
	}

	if showNextSpeakers != -1 {
		showNextSpeakerCount := showNextSpeakers

		if len(speakersWaiting) < showNextSpeakerCount {
			showNextSpeakerCount = len(speakersWaiting)
		}

		speakersWaiting = speakersWaiting[0:showNextSpeakerCount]
	}

	if speakersFinished == nil {
		speakersFinished = make([]formattedSpeaker, 0)
	}

	ti, err := titleInformation(ds, los)
	if err != nil {
		return nil, fmt.Errorf("get title information: %w", err)
	}

	out := struct {
		Waiting                 []formattedSpeaker                `json:"waiting"`
		Current                 *formattedSpeaker                 `json:"current"`
		Finished                []formattedSpeaker                `json:"finished"`
		ContentObjectCollection string                            `json:"content_object_collection"`
		TitleInformation        map[string]*projector.OptionalStr `json:"title_information"`
		Closed                  json.RawMessage                   `json:"closed"`
	}{
		speakersWaiting,
		currentSpeaker,
		speakersFinished,
		los.ContentObject.Collection,
		ti,
		los.Closed,
	}
	b, err := json.Marshal(out)
	if err != nil {
		return nil, fmt.Errorf("encoding outgoing data: %w", err)
	}
	return b, nil
}

// CurrentListOfSpeakersSlide renders the current list of speakers.
func CurrentListOfSpeakersSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		rp, err := referenceProjector(ds, pid)
		if err != nil {
			return nil, fmt.Errorf("getting reference projector: %w", err)
		}

		los, err := currentListOfSpeakers(ds, rp)
		if err != nil {
			if err == errNoListOfSpeakers {
				return []byte(`{}`), nil
			}
			return nil, fmt.Errorf("getting current list of speakers: %w", err)
		}

		return listOfSpeakerSlideData(ds, los)
	}
}

func referenceProjector(ds projector.Datastore, id int) (json.RawMessage, error) {
	var thisProjector struct {
		ReferenceID int `json:"reference_projector_id"`
	}
	if err := ds.Get("core/projector", id, &thisProjector); err != nil {
		return nil, fmt.Errorf("getting projector: %w", err)
	}

	referenceID := id
	if thisProjector.ReferenceID != 0 {
		referenceID = thisProjector.ReferenceID
	}

	var reference json.RawMessage

	if err := ds.Get("core/projector", referenceID, &reference); err != nil {
		return nil, fmt.Errorf("getting projector: %w", err)
	}
	return reference, nil
}

func currentListOfSpeakers(ds projector.Datastore, p json.RawMessage) (listOfSpeakers, error) {
	var pr struct {
		Elements []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"elements"`
	}
	if err := json.Unmarshal(p, &pr); err != nil {
		var je *json.UnmarshalTypeError
		if errors.As(err, &je) {
			if je.Field == "elements.id" {
				return listOfSpeakers{}, projector.NewClientError("elements.id is not an int.")
			}
		}
		return listOfSpeakers{}, fmt.Errorf("decoding projector: %w", err)
	}

	var los listOfSpeakers
	for _, element := range pr.Elements {
		if element.ID == 0 {
			continue
		}

		var model struct {
			LosID int `json:"list_of_speakers_id"`
		}
		if err := ds.Get(element.Name, element.ID, &model); err != nil {
			var doesNotExist interface {
				DoesNotExist() string
			}
			if errors.As(err, &doesNotExist) {
				continue
			}
			return listOfSpeakers{}, fmt.Errorf("getting element: %w", err)
		}

		if err := ds.Get("agenda/list-of-speakers", model.LosID, &los); err != nil {
			var doesNotExist interface {
				DoesNotExist() string
			}
			if errors.As(err, &doesNotExist) {
				continue
			}
			return listOfSpeakers{}, fmt.Errorf("getting list of speakers: %w", err)
		}
		return los, nil
	}
	return listOfSpeakers{}, errNoListOfSpeakers
}

// CurrentSpeakerChyronSlide renders the current list of speaker chyron.
func CurrentSpeakerChyronSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var projector struct {
			BColor json.RawMessage `json:"chyron_background_color"`
			FColor json.RawMessage `json:"chyron_font_color"`
		}

		if err := ds.Get("core/projector", pid, &projector); err != nil {
			return nil, fmt.Errorf("getting projector: %w", err)
		}

		currentSpeaker := make(map[string]json.RawMessage)
		currentSpeaker["background_color"] = projector.BColor
		currentSpeaker["font_color"] = projector.FColor

		rp, err := referenceProjector(ds, pid)
		if err != nil {
			return nil, fmt.Errorf("getting reference projector: %w", err)
		}

		los, err := currentListOfSpeakers(ds, rp)
		if err != nil {
			if err == errNoListOfSpeakers {
				b, err := json.Marshal(currentSpeaker)
				if err != nil {
					return nil, fmt.Errorf("encoding projector: %w", err)
				}
				return b, nil
			}
			return nil, fmt.Errorf("getting current list of speakers: %w", err)
		}

		var current []byte
		for _, speaker := range los.Speakers {
			if !bytes.Equal(speaker.BeginTime, []byte("null")) && bytes.Equal(speaker.EndTime, []byte("null")) {
				current, err = user.GetUserName(ds, speaker.UserID)
				if err != nil {
					return nil, fmt.Errorf("getting username: %w", err)
				}
				break
			}
		}

		if current != nil {
			currentSpeaker["current_speaker"] = current
		}

		b, err := json.Marshal(currentSpeaker)
		if err != nil {
			return nil, fmt.Errorf("encoding projector: %w", err)
		}
		return b, nil
	}
}
