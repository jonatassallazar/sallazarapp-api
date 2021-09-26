package controllers

import (
	"api/cmd/api/core/models"
	"api/cmd/api/repositories"
	"api/cmd/api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddClient realiza a criação de um novo cliente
func (gc *GeneralController) AddClient(c *gin.Context) {
	var client models.Client

	if c.Request.Body == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "corpo inválido"})
		return
	}

	c.BindJSON(&client)

	userID, accessLevel, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if accessLevel != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewClientsRepo(gc.Database)

	ID, err := repo.CreateClient(client, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id_inserido": ID,
	})
}

// GetAllClients realiza a busca de todos os clientes com o ID de usuário
//
// O usuário somente poderá atualizar os clientes que são pertencentes à ele
func (gc *GeneralController) GetAllClients(c *gin.Context) {
	userID, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewClientsRepo(gc.Database)
	clients, err := repo.GetClientsByUserID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	clientsJSON, err := convertClientsSliceToJSON(clients)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": clientsJSON,
	})
}

func convertClientsSliceToJSON(clients []models.Client) ([]models.ClientJSON, error) {
	var clientsJSON []models.ClientJSON

	for _, client := range clients {
		clientJSON, err := client.ClientToJSON()
		if err != nil {
			return []models.ClientJSON{}, err
		}

		clientsJSON = append(clientsJSON, clientJSON)
	}

	return clientsJSON, nil
}

// GetOneClient realiza a busca de um cliente com o ID fornecidor via parâmetro de url
//
// O usuário somente poderá atualizar os clientes que são pertencentes à ele
func (gc *GeneralController) GetOneClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewClientsRepo(gc.Database)
	client, err := repo.GetClientByID(userID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	clientJSON, err := client.ClientToJSON()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": clientJSON,
	})
}

// UpdateClient realiza a atualização dos dados de cliente com o ID fornecidor via parâmetro de url
//
// O usuário somente poderá atualizar os clientes que são pertencentes à ele
func (gc *GeneralController) UpdateClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.ClientJSON

	if c.Request.Body == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "corpo inválido"})
		return
	}

	c.BindJSON(&client)

	repo := repositories.NewClientsRepo(gc.Database)
	err = repo.UpdateClientByID(userID, id, client)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteClient deleta um cliente com o ID fornecidor via parâmetro de url
//
// O usuário somente poderá deletar os clientes que são pertencentes à ele
func (gc *GeneralController) DeleteClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _, err := utils.ExtractUserIdAndAccessLevel(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewClientsRepo(gc.Database)
	err = repo.DeleteClientByID(id, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
