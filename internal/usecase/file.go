package usecase

import "github.com/alsolovyev/dummy-api/internal/entity"

type FileUseCase struct {
	Repository entity.FileRepositoryer
}

func newFileUseCase(r entity.FileRepositoryer) *FileUseCase {
	return &FileUseCase{
		Repository: r,
	}
}

func (f *FileUseCase) GetFile(p string) (interface{}, *entity.Error) {
	d, err := f.Repository.GetFile(p)

	if err != nil {
		return nil, err
	}

	return d, nil
}
