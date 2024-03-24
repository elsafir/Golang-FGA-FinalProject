package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertStringParamToInt(ctx *gin.Context, param string) (int, error) {
	idString := ctx.Param(param)

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}
