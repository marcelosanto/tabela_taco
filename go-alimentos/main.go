package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Alimento struct {
	ID               int     `json:"id"`
	Description      string  `json:"description"`
	Category         string  `json:"category"`
	HumidityPercents float64 `json:"humidity_percents"`
	EnergyKcal       float64 `json:"energy_kcal"`
	EnergyKj         float64 `json:"energy_kj"`
	ProteinG         float64 `json:"protein_g"`
	LipidG           float64 `json:"lipid_g"`
	CholesterolMg    float64 `json:"cholesterol_mg"`
	CarbohydrateG    float64 `json:"carbohydrate_g"`
	FiberG           float64 `json:"fiber_g"`
	AshesG           float64 `json:"ashes_g"`
	CalciumMg        float64 `json:"calcium_mg"`
	MagnesiumMg      float64 `json:"magnesium_mg"`
	ManganeseMg      float64 `json:"manganese_mg"`
	PhosphorusMg     float64 `json:"phosphorus_mg"`
	IronMg           float64 `json:"iron_mg"`
	SodiumMg         float64 `json:"sodium_mg"`
	PotassiumMg      float64 `json:"potassium_mg"`
	CopperMg         float64 `json:"copper_mg"`
	ZincMg           float64 `json:"zinc_mg"`
	RetinolMcg       float64 `json:"retinol_mcg"`
	ReMcg            float64 `json:"re_mcg"`
	RaeMcg           float64 `json:"rae_mcg"`
	ThiamineMg       float64 `json:"thiamine_mg"`
	RiboflavinMg     float64 `json:"riboflavin_mg"`
	PyridoxineMg     float64 `json:"pyridoxine_mg"`
	NiacinMg         float64 `json:"niacin_mg"`
	VitaminCMg       float64 `json:"vitaminC_mg"`
	SaturatedG       float64 `json:"saturated_g"`
	MonounsaturatedG float64 `json:"monounsaturated_g"`
	PolyunsaturatedG float64 `json:"polyunsaturated_g"`
	One20G           float64 `json:"12:0_g"`
	One40G           float64 `json:"14:0_g"`
	One60G           float64 `json:"16:0_g"`
	One80G           float64 `json:"18:0_g"`
	Two00G           float64 `json:"20:0_g"`
	Two20G           float64 `json:"22:0_g"`
	Two40G           float64 `json:"24:0_g"`
	One41G           float64 `json:"14:1_g"`
	One61G           float64 `json:"16:1_g"`
	One81G           float64 `json:"18:1_g"`
	Two01G           float64 `json:"20:1_g"`
	One82N6G         float64 `json:"18:2 n-6_g"`
	One83N3G         float64 `json:"18:3 n-3_g"`
	Two04G           float64 `json:"20:4_g"`
	Two05G           float64 `json:"20:5_g"`
	Two25G           float64 `json:"22:5_g"`
	Two26G           float64 `json:"22:6_g"`
	One81TG          float64 `json:"18:1t_g"`
	One82TG          float64 `json:"18:2t_g"`
	TryptophanG      float64 `json:"tryptophan_g"`
	ThreonineG       float64 `json:"threonine_g"`
	IsoleucineG      float64 `json:"isoleucine_g"`
	LeucineG         float64 `json:"leucine_g"`
	LysineG          float64 `json:"lysine_g"`
	MethionineG      float64 `json:"methionine_g"`
	CystineG         float64 `json:"cystine_g"`
	PhenylalanineG   float64 `json:"phenylalanine_g"`
	TyrosineG        float64 `json:"tyrosine_g"`
	ValineG          float64 `json:"valine_g"`
	ArginineG        float64 `json:"arginine_g"`
	HistidineG       float64 `json:"histidine_g"`
	AlanineG         float64 `json:"alanine_g"`
	AsparticG        float64 `json:"aspartic_g"`
	GlutamicG        float64 `json:"glutamic_g"`
	GlycineG         float64 `json:"glycine_g"`
	ProlineG         float64 `json:"proline_g"`
	SerineG          float64 `json:"serine_g"`
}

func main() {
	jsonData, err := os.ReadFile("tabela_alimentos.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return
	}

	var alimentos []Alimento
	if err := json.Unmarshal(jsonData, &alimentos); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	http.HandleFunc("/tabela-alimentos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsonData, err := json.Marshal(alimentos)
		if err != nil {
			http.Error(w, "Erro ao serializar dados para JSON", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	})

	http.HandleFunc("/buscar-alimento", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		query := r.URL.Query().Get("nome")

		resultados := []Alimento{}
		for _, alimento := range alimentos {
			if strings.Contains(strings.ToLower(alimento.Description), strings.ToLower(query)) {
				resultados = append(resultados, alimento)
			}
		}

		jsonData, err := json.Marshal(resultados)
		if err != nil {
			http.Error(w, "Erro ao serializar dados para JSON", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	})

	port := 8080
	fmt.Printf("Servindo em http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
