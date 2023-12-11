package repository

import (
	"context"
	"time"

	"github.com/jeeo/pack-management/internal/model"
	"github.com/jmoiron/sqlx"
)

type PackRepository struct {
	db *sqlx.DB
}

type packDataModel struct {
	ID        string    `db:"id"`
	Amount    int       `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewPackRepository(db *sqlx.DB) PackRepository {
	return PackRepository{
		db: db,
	}
}

// translate from domain model to data model
func toPackDataModel(pack model.Pack) packDataModel {
	return packDataModel{
		ID:     pack.ID,
		Amount: pack.Amount,
	}
}

// translate from domain model to data model
func toPackDomainModel(pack packDataModel) model.Pack {
	return model.Pack{
		ID:     pack.ID,
		Amount: pack.Amount,
	}
}

func (p PackRepository) Create(ctx context.Context, pack model.Pack) error {
	packData := toPackDataModel(pack)
	insertQuery := "INSERT INTO pack (amount) VALUES ($1)"

	_, err := p.db.ExecContext(ctx, insertQuery, packData.Amount)
	return err
}

func (p PackRepository) FindByID(ctx context.Context, packID string) (model.Pack, error) {
	var pack packDataModel
	selectQuery := "SELECT id, amount FROM pack WHERE id = $1"

	err := p.db.GetContext(ctx, &pack, selectQuery, packID)
	if err != nil {
		return model.Pack{}, err
	}

	return toPackDomainModel(pack), nil
}

// returns packs ordered by amount
func (p PackRepository) FindAll(ctx context.Context) ([]model.Pack, error) {
	var packs []packDataModel
	var result []model.Pack
	selectQuery := "SELECT id, amount FROM pack ORDER BY amount ASC"

	err := p.db.SelectContext(ctx, &packs, selectQuery)
	if err != nil {
		return nil, err
	}

	for _, p := range packs {
		result = append(result, toPackDomainModel(p))
	}

	return result, nil
}

func (p PackRepository) Update(ctx context.Context, pack model.Pack) error {
	packData := toPackDataModel(pack)
	updateQuery := "UPDATE pack SET amount = $1, updated_at = now() WHERE id = $2"

	_, err := p.db.ExecContext(ctx, updateQuery, packData.Amount, packData.ID)
	return err
}

func (p PackRepository) Delete(ctx context.Context, packID string) error {
	deleteQuery := "DELETE FROM pack WHERE id = $1"

	_, err := p.db.ExecContext(ctx, deleteQuery, packID)
	return err
}
