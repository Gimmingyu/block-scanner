package presenter

import "github.com/gin-gonic/gin"

func BindQuery[T any](ctx *gin.Context) (*T, error) {
	var query = new(T)
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, err
	}
	return query, nil
}

func BindJSON[T any](ctx *gin.Context) (*T, error) {
	var json = new(T)
	if err := ctx.ShouldBindJSON(&json); err != nil {
		return nil, err
	}
	return json, nil

}
