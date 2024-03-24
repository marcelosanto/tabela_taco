package controllers

import (
	"go-alimentos/internal/app/models"
	"go-alimentos/internal/app/templates"
	"go-alimentos/internal/app/utils"
	"go-alimentos/internal/database"
	"log"
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

	component := templates.Index(templates.Alimentos(db), false)
	component.Render(c, c.Writer)

	//c.JSON(200, db)
}

func Buscar_Alimento(c *gin.Context) {
	comida := c.Request.FormValue("search")
	searchList := []models.Alimento{}

	log.Print("Comida:", utils.RemoveAccents(comida))

	db, err := database.StartDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

	}

	for _, alimento := range db {
		if strings.Contains(strings.ToLower(utils.RemoveAccents(alimento.Description)), strings.ToLower(comida)) {
			searchList = append(searchList, alimento)
		}
	}

	if len(searchList) == 0 {
		log.Println("Deu error")
		component := templates.Index(templates.Alimentos(db), true)
		component.Render(c, c.Writer)

	}

	component := templates.Index(templates.Alimentos(searchList), false)
	component.Render(c, c.Writer)
}
