package mediafile

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders "topics/topic".
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		m := ds.Get("mediafiles/mediafile:" + strconv.Itoa(element.ID))
		if m == nil {
			return nil, fmt.Errorf("mediafiles/mediafile:%d does not exist", element.ID)
		}

		var mediafile struct {
			Path           string `json:"path"`
			Mimetype       string `json:"mimetype"`
			MediaURLPrefix string `json:"media_url_prefix"`
		}
		if err := json.Unmarshal(m, &mediafile); err != nil {
			return nil, fmt.Errorf("decoding mediafile: %w", err)
		}

		out, err := json.Marshal(mediafile)
		if err != nil {
			return nil, fmt.Errorf("encoding mediafile: %w", err)
		}
		return out, nil
	}
}
