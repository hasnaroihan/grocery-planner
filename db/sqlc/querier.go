// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateIngredient(ctx context.Context, arg CreateIngredientParams) (Ingredient, error)
	CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error)
	CreateRecipeIngredient(ctx context.Context, arg CreateRecipeIngredientParams) (RecipesIngredient, error)
	CreateSchedule(ctx context.Context, author uuid.NullUUID) (Schedule, error)
	CreateScheduleRecipe(ctx context.Context, arg CreateScheduleRecipeParams) (SchedulesRecipe, error)
	CreateUnit(ctx context.Context, name string) (Unit, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteIngredient(ctx context.Context, id int32) error
	DeleteRecipe(ctx context.Context, id int64) error
	DeleteRecipeIngredient(ctx context.Context, arg DeleteRecipeIngredientParams) error
	DeleteSchedule(ctx context.Context, id int64) error
	DeleteScheduleRecipe(ctx context.Context, arg DeleteScheduleRecipeParams) error
	DeleteUnit(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetIngredient(ctx context.Context, id int32) (Ingredient, error)
	GetLogin(ctx context.Context, username string) (User, error)
	GetRecipe(ctx context.Context, id int64) (Recipe, error)
	GetRecipeIngredients(ctx context.Context, recipeID int64) ([]GetRecipeIngredientsRow, error)
	GetSchedule(ctx context.Context, id int64) (Schedule, error)
	GetScheduleRecipe(ctx context.Context, scheduleID int64) ([]GetScheduleRecipeRow, error)
	GetUnit(ctx context.Context, id int32) (Unit, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	ListGroceries(ctx context.Context, scheduleID int64) ([]ListGroceriesRow, error)
	ListIngredients(ctx context.Context) ([]Ingredient, error)
	ListRecipes(ctx context.Context, arg ListRecipesParams) ([]Recipe, error)
	ListSchedules(ctx context.Context, arg ListSchedulesParams) ([]Schedule, error)
	ListUnits(ctx context.Context) ([]Unit, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	SearchIngredientName(ctx context.Context, name string) (Ingredient, error)
	SearchIngredients(ctx context.Context, name string) ([]SearchIngredientsRow, error)
	SearchRecipe(ctx context.Context, arg SearchRecipeParams) ([]SearchRecipeRow, error)
	UpdateIngredient(ctx context.Context, arg UpdateIngredientParams) (Ingredient, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (UpdatePasswordRow, error)
	UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error)
	UpdateRecipeIngredient(ctx context.Context, arg UpdateRecipeIngredientParams) (RecipesIngredient, error)
	UpdateUnit(ctx context.Context, arg UpdateUnitParams) (Unit, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerified(ctx context.Context, arg UpdateVerifiedParams) (User, error)
}

var _ Querier = (*Queries)(nil)
