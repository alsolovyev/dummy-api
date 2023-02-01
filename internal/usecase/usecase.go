package usecase

import "github.com/alsolovyev/dummy-api/internal/entity"

type UseCase struct {
	File entity.FileUseCaser
}

func New(f entity.FileRepositoryer) *UseCase {
	return &UseCase{
		File: newFileUseCase(f),
	}
}
