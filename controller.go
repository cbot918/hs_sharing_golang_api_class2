package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	D *Dao
}

func NewController(d *Dao) Controller {
	return Controller{
		D: d,
	}
}

func (c *Controller) CreatePostController(ctx *gin.Context) {

	var post *Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := c.D.CreatePostDao(post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (c *Controller) ReadPostsController(ctx *gin.Context) {

	idString := ctx.Param("id")
	idInt, _ := strconv.Atoi(idString)

	post, err := c.D.FindByIdPostDao(idInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
