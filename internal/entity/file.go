package entity

type FileRepositoryer interface {
	GetFile(p string) (interface{}, *Error)
}

type FileUseCaser interface {
	GetFile(p string) (interface{}, *Error)
}
