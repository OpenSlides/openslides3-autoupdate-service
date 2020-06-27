package mediafile

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders "topics/topic".
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var mediafile struct {
			Path           string `json:"path"`
			Mimetype       string `json:"mimetype"`
			MediaURLPrefix string `json:"media_url_prefix"`
		}
		if err := projector.ModelFromElement(ds, e, "mediafiles/mediafile", &mediafile); err != nil {
			return nil, fmt.Errorf("getting mediafiles/mediafile: %w", err)
		}

		out, err := json.Marshal(mediafile)
		if err != nil {
			return nil, fmt.Errorf("encoding mediafile: %w", err)
		}
		return out, nil
	}
}
