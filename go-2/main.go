package main

import (
	"os"
	"bufio"
	"encoding/csv"
	"io"
	"sort"
	"strconv"
)

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
}
type jogador struct {
	name string
	wage float64
	age int
}
//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1() (int, error) {
	paises:= make(map[string] struct{}) 
	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return 0, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
			
		nationality:= record[14]
		if _, have := paises[nationality]; have || nationality == "" {
			continue
		}
		paises[nationality] = struct{}{}
	}

	return len(paises), nil
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q2() (int, error) {

	clubs:= make(map[string] struct{}) 
	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return 0, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
			
		club:= record[3]
		if _, have := clubs[club]; have || club == "" {
			continue
		}
		clubs[club] = struct{}{}
	}

	return len(clubs), nil
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3() ([]string, error) {
	
	var names []string

	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return nil, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for i:=0; i< 20; i++ {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
		if record[2] == "" {
			continue
		}
		names = append(names, record[2])
	}

	return names, nil
}

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {

	var top10 []string
	var jogadores []jogador
	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return []string{}, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
		wage, _ := strconv.ParseFloat(record[17], 64)
		jogadores = append(jogadores, jogador{
			name: record[2],
			wage: wage,
		})
	}
	sort.SliceStable(jogadores, func(a, b int) bool {
		return jogadores[a].wage > jogadores[b].wage
	})

	for _, jogador := range jogadores[0:10] {
		top10 = append(top10, jogador.name)
	}

	return top10, nil
}

//Quem são os 10 jogadores mais velhos (use como critério de desempate o campo `eur_wage`)?
func q5() ([]string, error) {


	var top10 []string
	var jogadores []jogador
	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return top10, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
		wage, _ := strconv.ParseFloat(record[17],64)
		age, _ := strconv.Atoi(record[6])
		jogadores = append(jogadores, jogador{
			name: record[2],
			wage: wage,
			age: age,
		})
	}
	sort.Slice(jogadores, func(a, b int) bool {
		if jogadores[a].age != jogadores[b].age {
			return jogadores[a].age > jogadores[b].age
		} 
		return jogadores[a].wage > jogadores[b].wage
	})

	for _, jogador := range jogadores[:10] {
		top10 = append(top10, jogador.name)
	}

	return top10, nil
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	idades := make(map[int]int)
	file, err:= os.Open("data.csv")
	defer file.Close()

	if err != nil{
		return map[int]int{}, err
	}

	reader:= csv.NewReader(bufio.NewReader(file))
	
	reader.Read() //skip on header

	for {
		record, err := reader.Read()
		if err==io.EOF {
			break
		}
			
		idade,_:= strconv.Atoi(record[6])
		idades[idade] = idades[idade] + 1
	}

	return idades, nil
}
