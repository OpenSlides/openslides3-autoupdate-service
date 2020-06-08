package user

// HasPermer tells if a user has a specivic perm.
type HasPermer interface {
	HasPerm(uid int, perm string) bool
}
