package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/ygorfelipe/brasil-api-go/internal/app/handlers/locations/dto"
	repositories "github.com/ygorfelipe/brasil-api-go/internal/infra/repositories/location"
)

type LocationHandler struct {
	locationRepository repositories.LocationRepository
}

func NewLocationRepository(locationRepository *repositories.LocationRepository) *LocationHandler {
	return &LocationHandler{locationRepository: *locationRepository}
}

func (l *LocationHandler) GetAllStates(c *gin.Context) {
	states, err := l.locationRepository.GetStates()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var stateResponse []dto.StatesResponse

	for _, s := range states {
		stateResponse = append(stateResponse, dto.StatesResponse{
			Acronym: s.Acronym,
			Name:    s.Name,
		})
	}

	c.JSON(200, stateResponse)
}
