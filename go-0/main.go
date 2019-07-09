package main

import (
	"fmt"
	"sort"
)

type infoEstado struct {
	nome string
	area float64
}

func main() {
	os10maioresEstadosDoBrasil()
	return
}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var dezMais []string

	estados, err := getEstados()
	if err != nil {
		return nil, err
	}

	sliceEstados := convertEstadosMap(estados)

	sort.Slice(sliceEstados, func(a, b int) bool {
		return sliceEstados[a].area > sliceEstados[b].area
	})
	for index, estado := range sliceEstados[:10] {
		fmt.Printf("#%v: %v\n", index+1, estado.nome)
		dezMais = append(dezMais, estado.nome)
	}
	return dezMais, nil

}

func convertEstadosMap(entries map[string]float64) []infoEstado {
	convertedEstados := make([]infoEstado, 0, len(entries))
	for estado, area := range entries {
		convertedEstados = append(convertedEstados, infoEstado{estado, area})
	}
	return convertedEstados
}

/*
Criando função para estados pois os mesmos
poderiam estar sendo retornados de uma fonte externa (API,DB etc).
*/

func getEstados() (map[string]float64, error) {

	estados := map[string]float64{
		"Rondônia": 237765.233,

		"Acre": 164123.738,

		"Amazonas": 1559168.117,

		"Roraima": 224273.831,

		"Pará": 1245759.305,

		"Amapá": 142470.762,

		"Tocantins": 277720.404,

		"Maranhão": 329642.17,

		"Piauí": 251616.823,

		"Ceará": 148894.757,

		"Rio Grande do Norte": 52809.602,

		"Paraíba": 56467.239,

		"Pernambuco": 98068.021,

		"Alagoas": 27843.295,

		"Sergipe": 21926.908,

		"Bahia": 564722.611,

		"Minas Gerais": 586521.121,

		"Espirito Santo": 46074.444,

		"Rio de Janeiro": 43750.423,

		"São Paulo": 248219.481,

		"Paraná": 199305.236,

		"Santa Catarina": 95730.921,

		"Rio Grande do Sul": 281707.151,

		"Mato Grosso do Sul": 357145.535,

		"Mato Grosso": 903206.997,

		"Goiás": 340125.715,

		"Distrito Federal": 5760.783}

	return estados, nil
}
