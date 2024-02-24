package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isantoshgyawali/apiWebGo/handlers/api"
)

func RequestRouter() *gin.Engine {
	r := gin.Default()

   r.Group("/", api.UserRoutes().Handlers...)
   r.Group("/", api.OrgRoutes().Handlers...)

   return r
}


