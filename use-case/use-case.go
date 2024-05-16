package usecase

import "minerva-content-status/repository"

type UseCase struct{ repo *repository.Repository }

func Initialize(repo *repository.Repository) *UseCase {
	uc := UseCase{repo}
	return &uc
}
