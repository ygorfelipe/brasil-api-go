package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ygorfelipe/brasil-api-go/internal/domain/entities"
	"github.com/ygorfelipe/brasil-api-go/internal/infra/repositories/location/dto"
)

// EM GO nao existe classes, mas sim é type
// GO não tem interface, mas iremos criar uma estrutura

type LocationRepository struct{}

func NewLocationRepository() *LocationRepository {
	// recuperando a referencia da memoria em * e o & atribui de onde iremos pegar
	return &LocationRepository{}
}

// funcao "metodos"

// retornando nossas entidades
// seria um ARRAY mas é um slices, procurar sobre, sempre irá retornar EIther/Nil, ou seja, entidade ou erro/valor ou erro
// pois não tem tryCatch
// tudo que tiver após o erro, basta acrescentar "," e sair colocando os multiplos retornos
func (l *LocationRepository) GetStates() ([]entities.StateEntity, error) {
	// buscando os dados brasil-api
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}

	// é possível trabalhar com wildCard = _, __, ___
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/ibge/uf/v1", nil)

	// sempre tratar o erro antes = premature return
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer serve para cancelar, caso o context esteja ativo após a execução, ele irá cancelar
	defer cancel()

	req = req.WithContext(ctx)

	// executando o response

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao buscar estados: ", err)
		return nil, err
	}

	// fechando a requisição do array de bits, fica dentro da funcao resp
	// funcao anonima
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error ao fechar o Body de buscar de estados %v", err)
		}

	}()

	var statesResponse []dto.BrasilApiStateResponse
	// convertendo
	// decoded

	// não é necessário declarar uma nova variavel para erro, pois ja existe lá em cima
	err = json.NewDecoder(resp.Body).Decode(&statesResponse)
	if err != nil {
		fmt.Println("Erro ao decodificar estados: ", err)
		return nil, err
	}

	// declarando uma var da entidade estados
	var states []entities.StateEntity

	// for (_, seria for var i = 0;)
	// segundo valor, valordesejado,
	// range, i++;

	for _, s := range statesResponse {
		// pegando a "classe" entidade, adicionando os valores da classe vindo do for de S
		// declaro a "lista" para adicionar, retorno a lista e o nil (vazio/erro)
		states = append(states, entities.StateEntity{
			Acronym: s.Sigla,
			Name:    s.Nome,
		})
	}
	return states, nil
}
