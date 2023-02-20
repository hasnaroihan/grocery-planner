package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/hasnaroihan/grocery-planner/db/sqlc"
)

type NewRecipeRequest struct {
	Name            string                   `json:"name" binding:"required"`
	Author          uuid.UUID                `json:"author" binding:"required"`
	Portion         int32                    `json:"portion" binding:"required,number,min=1"`
	Steps           sql.NullString           `json:"steps" binding:"required,alpha"`
	ListIngredients []db.ListIngredientParam `json:"ingredients" binding:"required,structonly,min=1"`
}

func (server *Server) newRecipe(ctx *gin.Context) {
	var req NewRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.NewRecipeParams {
		Name: req.Name,
		Author: req.Author,
		Portion: req.Portion,
		Steps: req.Steps,
		ListIngredients: req.ListIngredients,
	}
	recipe, err := server.storage.NewRecipeTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

type deleteRecipeRequest struct {
	ID int64	`uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteRecipe(ctx *gin.Context) {
	var req deleteRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	err := server.storage.DeleteRecipe(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type deleteRecipeIngredientRequest struct {
	RecipeID     int64 `form:"recipeID" binding:"required,min=1"`
	IngredientID int32 `form:"ingredientID" binding:"required,min=1"`
}

func (server *Server) deleteRecipeIngredient(ctx *gin.Context) {
	var req deleteRecipeIngredientRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.DeleteRecipeIngredientParams {
		RecipeID: req.RecipeID,
		IngredientID: req.IngredientID,
	}

	err := server.storage.DeleteRecipeIngredient(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type getRecipeRequest struct {
	ID int64	`uri:"id" binding:"required,min=1"`
}

func (server *Server) getRecipe(ctx *gin.Context) {
	var req getRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	recipe, err := server.storage.GetRecipeTx(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

type listRecipesRequest struct {
	PageSize  int32 `form:"pageSize" binding:"required,number"`
	PageNum int32 `form:"pageNum" binding:"required,number"`
}

func (server *Server) listRecipes(ctx *gin.Context) {
	var req listRecipesRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRecipesParams {
		Limit: req.PageSize,
		Offset: req.PageNum,
	}
	recipes, err := server.storage.ListRecipes(ctx,arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, recipes)
}

type searchRecipeRequest struct {
	Name   string `form:"name" binding:"omitempty,lowercase"`
	PageSize  int32  `form:"pageSize" binding:"required,number"`
	PageNum int32 `form:"pageNum" binding:"required,number"`
}

func (server *Server) searchRecipe(ctx *gin.Context) {
	var req searchRecipeRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.SearchRecipeParams {
		Name: fmt.Sprintf("%%%s%%",req.Name),
		Limit: req.PageSize,
		Offset: req.PageNum,
	}
	recipes, err := server.storage.SearchRecipe(ctx,arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, recipes)
}