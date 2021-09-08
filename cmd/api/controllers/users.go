package controllers

import (
	"api/cmd/api/core/models"
	"api/cmd/api/repositories"
	"api/cmd/api/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (gc *GeneralController) UserSignup(c *gin.Context) {
	var user models.User

	if c.Request.Body == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New(""))
		return
	}

	c.BindJSON(&user)

	hp, err := utils.SecurePassword(user.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.Password = hp

	repo := repositories.NewUsersRepo(gc.Database)

	ID, err := repo.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id_inserido": ID,
	})
}

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
