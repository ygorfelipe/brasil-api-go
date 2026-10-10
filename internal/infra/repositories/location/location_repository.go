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

type LocationRepository struct{}

func NewLocationRepository() *LocationRepository {
	// recuperando a referencia da memoria em * e o & atribui de onde iremos pegar
	return &LocationRepository{}
}

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
		fmt.Printf("Erro ao decodificar estados: %v\n", err)
		return nil, err
	}

	// declarando uma var da entidade estados
	var states []entities.StateEntity

	// for (_, seria for var i = 0;)
	// segundo valor, valordesejado,
	// range, i++;

	for _, s := range statesResponse {
		states = append(states, entities.StateEntity{
			Acronym: s.Sigla,
			Name:    s.Nome,
			Capital: s.Capital,
			Regiao: entities.RegionEntity{
				Id:    s.Regiao.Id,
				Sigla: s.Regiao.Nome,
				Nome:  s.Regiao.Sigla,
			},
		})
	}
	return states, nil
}

func (l *LocationRepository) GetAddressByCep(cep string) (*dto.ViacepApiResponse, error) {
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("Erro ao buscar endereço: ", err)
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Erro ao fechar o Body da buscar do endereco: %v\n", err)
		}
	}()

	var addressResponse dto.ViacepApiResponse

	err = json.NewDecoder(resp.Body).Decode(&addressResponse)

	if err != nil {
		return nil, err
	}

	return &addressResponse, nil

}
