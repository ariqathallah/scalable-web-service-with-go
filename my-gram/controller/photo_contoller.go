package controller

import (
	"my-gram/model/web"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PhotoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *PhotoController {
	return &PhotoController{photoService}
}

func (c *PhotoController) CreatePhoto(ctx *gin.Context) {
	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.PhotoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// call create photo service
	response, err := c.photoService.CreatePhoto(intUSerID, request)
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

func (c *PhotoController) GetAllPhotos(ctx *gin.Context) {
	// get all photos
	photos, err := c.photoService.GetAllPhotos()
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
		Data:    photos,
	})
}

func (c *PhotoController) UpdatePhoto(ctx *gin.Context) {
	// get params
	photoID := ctx.Param("id")
	intPhotoID, _ := strconv.Atoi(photoID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.PhotoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// update photo
	response, err := c.photoService.UpdatePhoto(intUSerID, intPhotoID, request)
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

func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
	// get params
	photoID := ctx.Param("id")
	intPhotoID, _ := strconv.Atoi(photoID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// delete photo
	err := c.photoService.DeletePhoto(intUSerID, intPhotoID)
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
		Message: "Your photo has been successfully deleted",
		Data:    nil,
	})
}
