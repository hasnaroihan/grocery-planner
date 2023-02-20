// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: recipe.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipes (
    name,
    author,
    portion,
    steps
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, author, portion, steps, created_at, modified_at
`

type CreateRecipeParams struct {
	Name    string         `json:"name"`
	Author  uuid.UUID      `json:"author"`
	Portion int32          `json:"portion"`
	Steps   sql.NullString `json:"steps"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.Name,
		arg.Author,
		arg.Portion,
		arg.Steps,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Portion,
		&i.Steps,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const createRecipeIngredient = `-- name: CreateRecipeIngredient :one
INSERT INTO recipes_ingredients (
    ingredient_id,
    recipe_id,
    amount,
    unit_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING ingredient_id, recipe_id, amount, unit_id
`

type CreateRecipeIngredientParams struct {
	IngredientID int32   `json:"ingredientID"`
	RecipeID     int64   `json:"recipeID"`
	Amount       float32 `json:"amount"`
	UnitID       int32   `json:"unitID"`
}

func (q *Queries) CreateRecipeIngredient(ctx context.Context, arg CreateRecipeIngredientParams) (RecipesIngredient, error) {
	row := q.db.QueryRowContext(ctx, createRecipeIngredient,
		arg.IngredientID,
		arg.RecipeID,
		arg.Amount,
		arg.UnitID,
	)
	var i RecipesIngredient
	err := row.Scan(
		&i.IngredientID,
		&i.RecipeID,
		&i.Amount,
		&i.UnitID,
	)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1
`

func (q *Queries) DeleteRecipe(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, id)
	return err
}

const deleteRecipeIngredient = `-- name: DeleteRecipeIngredient :exec
DELETE FROM recipes_ingredients
WHERE recipe_id = $1 AND ingredient_id = $2
`

type DeleteRecipeIngredientParams struct {
	RecipeID     int64 `json:"recipeID"`
	IngredientID int32 `json:"ingredientID"`
}

func (q *Queries) DeleteRecipeIngredient(ctx context.Context, arg DeleteRecipeIngredientParams) error {
	_, err := q.db.ExecContext(ctx, deleteRecipeIngredient, arg.RecipeID, arg.IngredientID)
	return err
}

const getRecipe = `-- name: GetRecipe :one
SELECT id, name, author, portion, steps, created_at, modified_at from recipes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRecipe(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Portion,
		&i.Steps,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getRecipeIngredients = `-- name: GetRecipeIngredients :many
SELECT ri.recipe_id, ri.ingredient_id, i.name, ri. amount, ri.unit_id
from recipes_ingredients as ri
INNER JOIN ingredients as i
ON ri.ingredient_id = i.id
WHERE recipe_id = $1
FOR SHARE
`

type GetRecipeIngredientsRow struct {
	RecipeID     int64   `json:"recipeID"`
	IngredientID int32   `json:"ingredientID"`
	Name         string  `json:"name"`
	Amount       float32 `json:"amount"`
	UnitID       int32   `json:"unitID"`
}

func (q *Queries) GetRecipeIngredients(ctx context.Context, recipeID int64) ([]GetRecipeIngredientsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecipeIngredients, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecipeIngredientsRow
	for rows.Next() {
		var i GetRecipeIngredientsRow
		if err := rows.Scan(
			&i.RecipeID,
			&i.IngredientID,
			&i.Name,
			&i.Amount,
			&i.UnitID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRecipes = `-- name: ListRecipes :many
SELECT id, name, author, portion, steps, created_at, modified_at from recipes
ORDER BY modified_at
LIMIT $1
OFFSET $2
`

type ListRecipesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRecipes(ctx context.Context, arg ListRecipesParams) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Author,
			&i.Portion,
			&i.Steps,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchRecipe = `-- name: SearchRecipe :many
SELECT id, name, author, modified_at from recipes
WHERE name LIKE $1
LIMIT $2
OFFSET $3
`

type SearchRecipeParams struct {
	Name   string `json:"name"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type SearchRecipeRow struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Author     uuid.UUID `json:"author"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func (q *Queries) SearchRecipe(ctx context.Context, arg SearchRecipeParams) ([]SearchRecipeRow, error) {
	rows, err := q.db.QueryContext(ctx, searchRecipe, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchRecipeRow
	for rows.Next() {
		var i SearchRecipeRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Author,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRecipe = `-- name: UpdateRecipe :one
UPDATE recipes
    set name = $2,
    portion = $3,
    steps = $4,
    modified_at = $5
WHERE id = $1
RETURNING id, name, author, portion, steps, created_at, modified_at
`

type UpdateRecipeParams struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	Portion    int32          `json:"portion"`
	Steps      sql.NullString `json:"steps"`
	ModifiedAt time.Time      `json:"modifiedAt"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipe,
		arg.ID,
		arg.Name,
		arg.Portion,
		arg.Steps,
		arg.ModifiedAt,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Portion,
		&i.Steps,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateRecipeIngredient = `-- name: UpdateRecipeIngredient :one
UPDATE recipes_ingredients
    set amount = $3,
    unit_id = $4
WHERE recipe_id = $1 AND ingredient_id = $2
RETURNING ingredient_id, recipe_id, amount, unit_id
`

type UpdateRecipeIngredientParams struct {
	RecipeID     int64   `json:"recipeID"`
	IngredientID int32   `json:"ingredientID"`
	Amount       float32 `json:"amount"`
	UnitID       int32   `json:"unitID"`
}

func (q *Queries) UpdateRecipeIngredient(ctx context.Context, arg UpdateRecipeIngredientParams) (RecipesIngredient, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeIngredient,
		arg.RecipeID,
		arg.IngredientID,
		arg.Amount,
		arg.UnitID,
	)
	var i RecipesIngredient
	err := row.Scan(
		&i.IngredientID,
		&i.RecipeID,
		&i.Amount,
		&i.UnitID,
	)
	return i, err
}
