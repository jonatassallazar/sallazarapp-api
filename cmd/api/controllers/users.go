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

// GetUser realiza a busca pelo repositório com o ID fornecidor via parâmetro de url
//
// O usuário pode buscar somente suas próprias informações
func (gc *GeneralController) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	idToken, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if idToken != id {
		c.AbortWithError(http.StatusUnauthorized, errors.New("o usuário pode buscar somente suas próprias informações"))
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

// UpdateUser realiza a atualização de nome e email pelo repositório com o ID fornecidor via parâmetro de url
//
// O usuário pode atualizar somente suas próprias informações
func (gc *GeneralController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	idToken, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if idToken != id {
		c.AbortWithError(http.StatusUnauthorized, errors.New("o usuário pode buscar somente suas próprias informações"))
		return
	}

	var user models.User

	if c.Request.Body == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New(""))
		return
	}

	c.BindJSON(&user)

	repo := repositories.NewUsersRepo(gc.Database)
	err = repo.UpdateUserByID(id, user)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}

// DeleteUser deleta um usuário com o ID fornecidor via parâmetro de url
//
// O usuário pode deletar somente a si mesmo
func (gc *GeneralController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	idToken, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if idToken != id {
		c.AbortWithError(http.StatusUnauthorized, errors.New("o usuário pode buscar somente suas próprias informações"))
		return
	}

	repo := repositories.NewUsersRepo(gc.Database)
	err = repo.DeleteUserByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}
