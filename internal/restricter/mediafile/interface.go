package mediafile

import "github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"

type required interface {
	restricter.HasPermer
	IsSuperadmin(uid int) bool
	InGroups(uid int, groups []int) bool
}
