// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID          int32         `json:"id"`
	Name        string        `json:"name"`
	CreatedAt   time.Time     `json:"createdAt"`
	DefaultUnit sql.NullInt32 `json:"defaultUnit"`
}

type Recipe struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	Author     uuid.UUID      `json:"author"`
	Portion    int32          `json:"portion"`
	Steps      sql.NullString `json:"steps"`
	CreatedAt  time.Time      `json:"createdAt"`
	ModifiedAt time.Time      `json:"modifiedAt"`
}

type RecipesIngredient struct {
	IngredientID int32   `json:"ingredientID"`
	RecipeID     int64   `json:"recipeID"`
	Amount       float32 `json:"amount"`
	UnitID       int32   `json:"unitID"`
}

type Schedule struct {
	ID        int64         `json:"id"`
	Author    uuid.NullUUID `json:"author"`
	CreatedAt time.Time     `json:"createdAt"`
}

type SchedulesRecipe struct {
	ScheduleID int64 `json:"scheduleID"`
	RecipeID   int64 `json:"recipeID"`
	Portion    int32 `json:"portion"`
}

type Unit struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         uuid.UUID    `json:"id"`
	Username   string       `json:"username"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	CreatedAt  time.Time    `json:"createdAt"`
	Role       string       `json:"role"`
	VerifiedAt sql.NullTime `json:"verifiedAt"`
}
