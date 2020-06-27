package agenda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

type treeInfo struct {
	TitleInformation map[string]string `json:"title_information"`
	Collection       string            `json:"collection"`
	Depth            int               `json:"depth"`
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

		var agendaItems []treeInfo
		if element.OnlyMainItems {
			for _, item := range sortedItems(items) {
				if item.ParentID != 0 || item.Type != 1 {
					continue
				}
				titleInformation := item.TitleInformation
				titleInformation["_agenda_item_number"] = item.ItemNumber
				agendaItems = append(agendaItems, treeInfo{
					TitleInformation: titleInformation,
					Collection:       item.ContentObject.Collection,
				})
			}
		} else {
			var err error
			agendaItems, err = getFlatTree(items, 0)
			if err != nil {
				return nil, fmt.Errorf("get flat tree: %w", err)
			}
		}

		out, err := json.Marshal(agendaItems)
		if err != nil {
			return nil, fmt.Errorf("decoding agenda items: %w", err)
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

func getFlatTree(agendaItems map[int]agendaItem, parentID int) ([]treeInfo, error) {
	children := make(map[int][]int)
	for _, item := range sortedItems(agendaItems) {
		if item.Type == 1 {
			children[item.ParentID] = append(children[item.ParentID], item.ID)
		}
	}

	var tree []treeInfo
	var buildTree func(itemIDs []int, depth int)
	buildTree = func(itemIDs []int, depth int) {
		for _, itemID := range itemIDs {
			item := agendaItems[itemID]
			titleInformation := item.TitleInformation
			titleInformation["_agenda_item_number"] = item.ItemNumber
			tree = append(tree, treeInfo{
				TitleInformation: titleInformation,
				Collection:       item.ContentObject.Collection,
				Depth:            depth,
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
		var element struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		if element.ID == 0 {
			return nil, fmt.Errorf("id is required for list of speakers slide")
		}

		l := ds.Get(fmt.Sprintf("%s:%d", "agenda/list-of-speakers", element.ID))
		if l == nil {
			return nil, fmt.Errorf("model %s:%d does not exist", "agenda/list-of-speakers", element.ID)
		}

		return listOfSpeakerSlideData(ds, l)
	}
}

type formattedSpeaker struct {
	User    string          `json:"user"`
	Marked  json.RawMessage `json:"marked"`
	Weight  int             `json:"weight"`
	EndTime json.RawMessage `json:"end_time"`
}

func listOfSpeakerSlideData(ds projector.Datastore, l json.RawMessage) (json.RawMessage, error) {
	var los listOfSpeakers
	if err := json.Unmarshal(l, &los); err != nil {
		return nil, fmt.Errorf("decoding list of speakers: %w", err)
	}

	titleInformation := los.TitleInformation
	contentObject := ds.Get(fmt.Sprintf("%s:%d", los.ContentObject.Collection, los.ContentObject.ID))

	// Set titleInformation["_agenda_item_number"].
	if contentObject != nil {
		// contentObject does exist.
		var object struct {
			AgendaItemID int `json:"agenda_item_id"`
		}
		if err := json.Unmarshal(contentObject, &object); err != nil {
			return nil, fmt.Errorf("decoding content object: %w", err)
		}

		i := ds.Get(fmt.Sprintf("%s:%d", "agenda/item", object.AgendaItemID))
		if i != nil {
			var item struct {
				Number string `json:"item_number"`
			}
			if err := json.Unmarshal(i, &item); err != nil {
				return nil, fmt.Errorf("decoding item: %w", err)
			}
			titleInformation["_agenda_item_number"] = item.Number
		}
	}

	var showLastSpeakers int
	if err := ds.ConfigValue("agenda_show_last_speakers", &showLastSpeakers); err != nil {
		return nil, fmt.Errorf("loading agenda_show_last_speakers: %w", err)
	}

	var speakersWaiting []formattedSpeaker
	var speakersFinished []formattedSpeaker
	var currentSpeaker *formattedSpeaker
	for _, speaker := range los.Speakers {
		username, err := user.GetUserName(ds, speaker.UserID)
		if err != nil {
			return nil, fmt.Errorf("getting username: %w", err)
		}
		fs := formattedSpeaker{
			User:    username,
			Marked:  speaker.Marked,
			Weight:  speaker.Weight,
			EndTime: speaker.EndTime,
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
		return speakersWaiting[i].Weight < speakersWaiting[j].Weight
	})

	if showLastSpeakers > 0 {
		sort.Slice(speakersFinished, func(i, j int) bool {
			return speakersFinished[i].Weight < speakersFinished[j].Weight
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
		speakersWaiting = speakersWaiting[0:showNextSpeakers]
	}

	if speakersFinished == nil {
		speakersFinished = make([]formattedSpeaker, 0)
	}

	out := struct {
		Waiting                 []formattedSpeaker `json:"waiting"`
		Current                 *formattedSpeaker  `json:"current"`
		Finished                []formattedSpeaker `json:"finished"`
		ContentObjectCollection string             `json:"content_object_collection"`
		TitleInformation        map[string]string  `json:"title_information"`
		Closed                  json.RawMessage    `json:"closed"`
	}{
		speakersWaiting,
		currentSpeaker,
		speakersFinished,
		los.ContentObject.Collection,
		titleInformation,
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
			return nil, fmt.Errorf("getting current list of speakers: %w", err)
		}

		if los == nil {
			return []byte(`{}`), nil
		}

		return listOfSpeakerSlideData(ds, los)
	}
}

func referenceProjector(ds projector.Datastore, id int) (json.RawMessage, error) {
	t := ds.Get(fmt.Sprintf("%s:%d", "core/projector", id))
	var thisProjector struct {
		ReferenceID int `json:"reference_projector_id"`
	}
	if err := json.Unmarshal(t, &thisProjector); err != nil {
		return nil, fmt.Errorf("decoding projector: %w", err)
	}
	referenceID := id
	if thisProjector.ReferenceID != 0 {
		referenceID = thisProjector.ReferenceID
	}

	return ds.Get(fmt.Sprintf("%s:%d", "core/projector", referenceID)), nil
}

func currentListOfSpeakers(ds projector.Datastore, p json.RawMessage) (json.RawMessage, error) {
	var projector struct {
		Elements []struct {
			Name string `json:"name"`
			ID   int    `josn:"id"`
		}
	}
	if err := json.Unmarshal(p, &projector); err != nil {
		return nil, fmt.Errorf("decoding projector: %w", err)
	}

	var los json.RawMessage
	for _, element := range projector.Elements {
		if element.ID == 0 {
			continue
		}

		m := ds.Get(fmt.Sprintf("%s:%d", element.Name, element.ID))
		if m == nil {
			continue
		}

		var model struct {
			LosID int `json:"list_of_speakers_id"`
		}
		if err := json.Unmarshal(m, &model); err != nil {
			// Ignore invalid models.
			continue
		}

		los = ds.Get(fmt.Sprintf("%s:%d", "agenda/list-of-speakers", model.LosID))
		if los != nil {
			break
		}
	}
	return los, nil
}

// CurrentSpeakerChyronSlide renders the current list of speaker chyron.
func CurrentSpeakerChyronSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		p := ds.Get(fmt.Sprintf("%s:%d", "core/projector", pid))
		var projector struct {
			BColor json.RawMessage `json:"chyron_background_color"`
			FColor json.RawMessage `json:"chyron_font_color"`
		}
		if err := json.Unmarshal(p, &projector); err != nil {
			return nil, fmt.Errorf("decoding projector: %w", err)
		}

		currentSpeaker := make(map[string]json.RawMessage)
		currentSpeaker["background_color"] = projector.BColor
		currentSpeaker["font_color"] = projector.FColor

		rp, err := referenceProjector(ds, pid)
		if err != nil {
			return nil, fmt.Errorf("getting reference projector: %w", err)
		}

		l, err := currentListOfSpeakers(ds, rp)
		if err != nil {
			return nil, fmt.Errorf("getting current list of speakers: %w", err)
		}

		if l == nil {
			b, err := json.Marshal(currentSpeaker)
			if err != nil {
				return nil, fmt.Errorf("encoding projector: %w", err)
			}
			return b, nil
		}

		var los listOfSpeakers
		if err := json.Unmarshal(l, &los); err != nil {
			return nil, fmt.Errorf("decoding list of speakers: %w", err)
		}

		var current string
		for _, speaker := range los.Speakers {
			if !bytes.Equal(speaker.BeginTime, []byte("null")) && bytes.Equal(speaker.EndTime, []byte("null")) {
				current, err = user.GetUserName(ds, speaker.UserID)
				if err != nil {
					return nil, fmt.Errorf("getting username: %w", err)
				}
				break
			}
		}

		if current != "" {
			currentSpeaker["current_speaker"] = []byte(`"` + current + `"`)
		}

		b, err := json.Marshal(currentSpeaker)
		if err != nil {
			return nil, fmt.Errorf("encoding projector: %w", err)
		}
		return b, nil
	}
}
