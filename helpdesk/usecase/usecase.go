package usecase

import (
	"context"

	"github.com/DarkSoul94/helpdesk2/helpdesk"
)

// Usecase ...
type Usecase struct {
	repo helpdesk.Repository
}

// NewUsecase ...
func NewUsecase(repo helpdesk.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// HelloWorld ...
func (u *Usecase) HelloWorld(c context.Context) {
	println("Hello")
}
