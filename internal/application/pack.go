package application

import (
	"context"
	"database/sql"
	"errors"

	domainerrors "github.com/jeeo/pack-management/internal/errors"
	"github.com/jeeo/pack-management/internal/model"
)

type packRepository interface {
	Create(context.Context, model.Pack) error
	FindByID(context.Context, string) (model.Pack, error)
	FindAll(context.Context) ([]model.Pack, error)
	Update(context.Context, model.Pack) error
	Delete(context.Context, string) error
}

type PackApplication struct {
	packRepo packRepository
}

func NewPackApplication(packRepo packRepository) PackApplication {
	return PackApplication{
		packRepo: packRepo,
	}
}

func (a PackApplication) FindAll(ctx context.Context) ([]model.Pack, error) {
	return a.packRepo.FindAll(ctx)
}

func (a PackApplication) Create(ctx context.Context, amount int) error {
	pack := model.Pack{
		Amount: amount,
	}

	return a.packRepo.Create(ctx, pack)
}

func (a PackApplication) Update(ctx context.Context, pack model.Pack) error {
	if _, err := a.packRepo.FindByID(ctx, pack.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainerrors.NewErrPackNotFound(err, pack.ID)
		}
	}

	return a.packRepo.Update(ctx, pack)
}

func (a PackApplication) Delete(ctx context.Context, id string) error {
	if _, err := a.packRepo.FindByID(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainerrors.NewErrPackNotFound(err, id)
		}
	}

	return a.packRepo.Delete(ctx, id)
}
