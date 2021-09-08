package controllers

import (
	"api/cmd/api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (gc *GeneralController) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	repo := repositories.NewUsersRepo(gc.Database)
	user, err := repo.GetUserByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
