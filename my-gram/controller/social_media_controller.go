package controller

import (
	"my-gram/model/web"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SocialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{socialMediaService}
}

func (c *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.SocialMediaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// call create social media service
	response, err := c.socialMediaService.Create(intUSerID, request)
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

func (c *SocialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	// get all social medias
	socialMedias, err := c.socialMediaService.GetAllSocialMedias()
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
		Data:    socialMedias,
	})
}

func (c *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	// get params
	socialMediaID := ctx.Param("id")
	intSocialMediaID, _ := strconv.Atoi(socialMediaID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// bind request
	var request web.SocialMediaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// call update social media service
	response, err := c.socialMediaService.Update(intUSerID, intSocialMediaID, request)
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

func (c *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	// get params
	socialMediaID := ctx.Param("id")
	intSocialMediaID, _ := strconv.Atoi(socialMediaID)

	// get userID from claims
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID := claims["sub"].(float64)
	intUSerID := int(userID)

	// call delete social media service
	err := c.socialMediaService.Delete(intUSerID, intSocialMediaID)
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
		Message: "Your social medias has been successfully deleted",
		Data:    nil,
	})
}
