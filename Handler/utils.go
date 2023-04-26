package Handler

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h Handler) GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
	}
}
