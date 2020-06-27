package motion

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
// )

// // Slide renders a an assignment.
// func Slide() projector.CallableFunc {
// 	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
// 		var element struct {
// 			ID   int    `json:"id"`
// 			Mode string `json:"mode"`
// 		}
// 		if err := json.Unmarshal(e, &element); err != nil {
// 			return nil, fmt.Errorf("decoding element: %w", err)
// 		}

// 		m := ds.Get(fmt.Sprintf("%s:%d", "motions/motion", element.ID))
// 		if m == nil {
// 			return nil, fmt.Errorf("model motions/motion:%d does not exist", element.ID)
// 		}

// 		var mode string
// 		if err := ds.ConfigValue("motions_recommendation_text_mode", &mode); err != nil {
// 			return nil, fmt.Errorf("get config value for motions_recommendation_text_mode: %w", err)
// 		}
// 	}
// }
