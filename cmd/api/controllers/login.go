package controllers

import (
	"api/cmd/api/core/models"
	"api/cmd/api/repositories"
	"api/cmd/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSignup realiza a criação de um novo usuário e retorna um jwt válido
func (gc *GeneralController) UserSignup(c *gin.Context) {
	var user models.User

	if c.Request.Body == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "corpo inválido"})
		return
	}

	c.BindJSON(&user)

	hp, err := utils.SecurePassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user.Password = hp

	repo := repositories.NewUsersRepo(gc.Database)

	ID, err := repo.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := utils.CreateToken(ID, "admin")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id_inserido": ID,
		"token":       token,
	})
}

func (gc *GeneralController) UserLogin(c *gin.Context) {
	var user models.User

	if c.Request.Body == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "corpo inválido"})
		return
	}

	c.BindJSON(&user)

	repo := repositories.NewUsersRepo(gc.Database)
	userDB, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := utils.VerifyPassword(userDB.Password, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "e-mail ou senha incorretos"})
		return
	}

	token, err := utils.CreateToken(userDB.ID, userDB.AccessLevel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
