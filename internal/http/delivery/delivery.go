package delivery

import "backend/internal/http/usecase"

type Handler struct {
	UseCase *usecase.UseCase
}

func NewHandler(uc *usecase.UseCase) *Handler {
	return &Handler{UseCase: uc}
}
