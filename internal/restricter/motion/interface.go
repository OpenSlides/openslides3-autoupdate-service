package motion

import "github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"

type required interface {
	restricter.HasPermer
	InGroups(uid int, groups []int) bool
}
