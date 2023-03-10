package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hasnaroihan/grocery-planner/db/sqlc"
	"github.com/lib/pq"
)

type createIngredientRequest struct {
	Name        string        `json:"name" binding:"required,lowercase"`
	DefaultUnit sql.NullInt32 `json:"defaultUnit"`
}

func (server *Server) createIngredient(ctx *gin.Context) {
	var req createIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.CreateIngredientParams{
		Name: req.Name,
		DefaultUnit: sql.NullInt32{
			Int32: req.DefaultUnit.Int32,
			Valid: req.DefaultUnit.Valid,
		},
	}

	ingredient, err := server.storage.CreateIngredient(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				ctx.JSON(http.StatusConflict, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredient)
}

type deleteIngredientRequest struct {
	ID int32	`uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteIngredient(ctx *gin.Context) {
	var req deleteIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	err := server.storage.DeleteIngredient(ctx, req.ID)
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

type getIngredientRequest struct {
	ID int32	`uri:"id" binding:"required,min=1"`
}

func (server *Server) getIngredient(ctx *gin.Context) {
	var req getIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ingredient, err := server.storage.GetIngredient(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredient)
}

func (server *Server) listIngredients(ctx *gin.Context) {
	ingredients, err := server.storage.ListIngredients(ctx)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, ingredients)
}

type searchIngredientRequest struct {
	Name string		`form:"name" binding:"omitempty,lowercase"`
}

func (server *Server) searchIngredients(ctx *gin.Context) {
	var q searchIngredientRequest

	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ingredients, err := server.storage.SearchIngredients(ctx, fmt.Sprintf("%%%s%%",q.Name))
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, ingredients)
}

type updateIngredientUri struct {
	ID int32	`uri:"id" binding:"required,min=1"`
}

type updateIngredientJSON struct {
	ID 			int32		  `json:"id" binding:"required,min=1"`
	Name        string        `json:"name" binding:"required,lowercase"`
	DefaultUnit sql.NullInt32 `json:"defaultUnit"`
}

func (server *Server) updateIngredient(ctx *gin.Context) {
	var reqUri updateIngredientUri
	var reqJSON updateIngredientJSON

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if reqUri.ID != reqJSON.ID {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("mismatched uri and body ingredient id")))
		return
	}

	arg := db.UpdateIngredientParams{
		ID: reqUri.ID,
		Name: reqJSON.Name,
		DefaultUnit: sql.NullInt32{
			Int32: reqJSON.DefaultUnit.Int32,
			Valid: reqJSON.DefaultUnit.Valid,
		},
	}

	ingredient, err := server.storage.UpdateIngredient(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredient)
}
