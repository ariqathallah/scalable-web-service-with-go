package controller

import (
	"my-gram/model/web"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{commentService}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// create photo
	response, err := c.commentService.CreateComment(intUSerID, request)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.WebResponse{
		Code:    http.StatusCreated,
		Message: "ok",
		Data:    response,
	})
}

func (c *CommentController) GetAllComments(ctx *gin.Context) {
	// get all comments
	comments, err := c.commentService.GetAllComments()
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    comments,
	})
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	// get params
	commentID := ctx.Param("id")
	intCommentID, _ := strconv.Atoi(commentID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// update comment
	response, err := c.commentService.UpdateComment(intUSerID, intCommentID, request)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    response,
	})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	// get params
	commentID := ctx.Param("id")
	intCommentID, _ := strconv.Atoi(commentID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// delete comment
	err := c.commentService.DeleteComment(intUSerID, intCommentID)
	if err != nil {
		ctx.JSON(err.Code, web.WebResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Your comment has been successfully deleted",
		Data:    nil,
	})
}
