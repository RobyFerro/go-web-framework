package services

// GitServices exposes methods to clone a repository
type GitServices interface {
	Clone(repo, destination string) error
	Remove(destination string) error
	Update(repo, destination string) error
}
