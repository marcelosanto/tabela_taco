package controllers

import (
	"go-alimentos/internal/app/models"
	"go-alimentos/internal/app/templates"
	"go-alimentos/internal/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Listar_Alimentos(c *gin.Context) {
	db, err := database.StartDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	component := templates.Index(templates.Alimentos(db))
	component.Render(c, c.Writer)

	//c.JSON(200, db)
}

func Buscar_Alimento(c *gin.Context) {
	comida := c.Param("comida")

	db, err := database.StartDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	all := []models.Alimento{}

	for _, alimento := range db {
		if strings.Contains(strings.ToLower(alimento.Description), strings.ToLower(comida)) {
			all = append(all, alimento)
		}
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Comida não encontrada",
		})
	}

	c.JSON(http.StatusOK, all)

}
