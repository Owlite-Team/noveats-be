package postgres

import (
	"database/sql"
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
)

type menuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) repository.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *entity.MenuItem) error {
	query := `
		INSERT INTO MenuItems (id, name, description, price, category, image_url,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(query,
		menu.ID,
		menu.Name,
		menu.Description,
		menu.Price,
		menu.Category,
		menu.ImageURL,
		menu.CreatedAt,
		menu.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return err
}

func (r *menuRepository) FindById(id string) (*entity.MenuItem, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM MenuItems
		WHERE id = $1
	`
	items := &entity.MenuItem{}
	err := r.db.QueryRow(query, id).Scan(
		&items.ID,
		&items.Name,
		&items.Description,
		&items.Price,
		&items.Category,
		&items.ImageURL,
		&items.CreatedAt,
		&items.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("menu not found")
	}
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *menuRepository) FindByName(name string) (*entity.MenuItem, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM MenuItems
		WHERE name = $1
	`
	items := &entity.MenuItem{}
	err := r.db.QueryRow(query, name).Scan(
		&items.ID,
		&items.Name,
		&items.Description,
		&items.Price,
		&items.Category,
		&items.ImageURL,
		&items.CreatedAt,
		&items.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("menu not found")
	}
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *menuRepository) FindAll() ([]*entity.MenuItem, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM MenuItems
		ORDER BY name
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // TODO(Fix)

	var items []*entity.MenuItem
	for rows.Next() {
		item := &entity.MenuItem{}
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Category,
			&item.ImageURL,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *menuRepository) Update(menu *entity.MenuItem) error {
	query := `
		UPDATE MenuItems
		SET name = $1, description = $2, price = $3, category = $4, image_url = $5
		WHERE id = $6
	`
	result, err := r.db.Exec(query,
		menu.Name,
		menu.Description,
		menu.Price,
		menu.Category,
		menu.ImageURL,
		menu.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("menu not found")
	}
	return nil
}

func (r *menuRepository) Delete(id string) error {
	query := `DELETE FROM MenuItems WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("menu not found")
	}
	return nil
}
