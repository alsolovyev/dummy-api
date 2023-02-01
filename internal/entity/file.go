package entity

type FileRepositoryer interface {
	GetFile(p string) (interface{}, error)
}

type FileUseCaser interface {
	GetFile(p string) (interface{}, error)
}
