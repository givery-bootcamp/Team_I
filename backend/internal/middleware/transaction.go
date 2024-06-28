package middleware

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/external/database"
)

func Transaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.DB.Begin()
		defer func() {
			if 400 <= ctx.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		ctx.Set("db", db)
		ctx.Next()
	}
}
