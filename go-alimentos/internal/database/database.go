package database

import (
	"encoding/json"
	"fmt"
	"go-alimentos/internal/app/models"

	"os"
)

func StartDB() ([]models.Alimento, error) {
	jsonData, err := os.ReadFile("tabela_alimentos.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return []models.Alimento{}, err
	}

	var alimentos []models.Alimento
	if err := json.Unmarshal(jsonData, &alimentos); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return []models.Alimento{}, err
	}

	return alimentos, nil
}
