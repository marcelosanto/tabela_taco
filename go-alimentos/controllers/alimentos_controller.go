package controllers

import (
	"go-alimentos/database"
	"go-alimentos/models"
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

	c.JSON(200, db)
}

func Buscar_Alimento(c *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// query := r.PathValue("nome")

	// resultados := alimentos
	// for _, alimento := range alimentos {
	// 	if strings.Contains(strings.ToLower(alimento.Description), strings.ToLower(query)) {
	// 		resultados = append(resultados, alimento)
	// 	}
	// }

	// jsonData, err = json.Marshal(resultados)
	// if err != nil {
	// 	http.Error(w, "Erro ao serializar dados para JSON", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)
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
