package agenda

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

type agendaItem struct {
	ID               int               `json:"id"`
	ParentID         int               `json:"parent_id"`
	Weight           int               `json:"weight"`
	Type             int               `json:"type"`
	TitleInformation map[string]string `json:"title_information"`
	ItemNumber       string            `json:"item_number"`
	ContentObject    struct {
		Collection string `json:"collection"`
	} `json:"content_object"`
}

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
