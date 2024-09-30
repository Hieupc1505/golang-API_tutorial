package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/hieupc09/simple_api/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type createTodoTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image"`
}

func (server *Server) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (server *Server) createTodoTask(ctx *gin.Context) {
	var req createTodoTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	des := pgtype.Text{String: req.Description, Valid: true}
	img := pgtype.Text{String: req.Image, Valid: true}
	arg := db.CreateTodoTaskParams{
		Title:       req.Title,
		Description: des,
		Image:       img,
	}
	task, err := server.store.CreateTodoTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

type getListTodoTaskRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getListTodoTask(ctx *gin.Context) {
	var req getListTodoTaskRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetTodoTaskListParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	tasks, err := server.store.GetTodoTaskList(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

type updateTodoTaskRequest struct {
	ID     int64         `json:"id" binding:"required"`
	Status db.TaskStatus `json:"status" binding:"required"`
}

func (server *Server) updateTodoTask(ctx *gin.Context) {
	var req updateTodoTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTodoTaskParams{
		ID:     req.ID,
		Status: req.Status,
	}
	_, err := server.store.UpdateTodoTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": 1})
}

func (server *Server) deleteTodoTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	num, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = server.store.DeleteTodoTask(ctx, int64(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": 1})
}

type registerUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
