package handler

import (
	"context"
	"fmt"
	"go-microservice/pkg/proto"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var grpcClient proto.MessageServiceClient

func InitRoutes(cl proto.MessageServiceClient) {
	grpcClient = cl
	router := gin.Default()

	router.GET("/sendgroup", SendGroupChat)
	router.GET("/sendchannel", SendChannel)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8081")

}

// @Summary Send message
// @Description Send message to group chat
// @ID send-message-group-chat
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Param priority query string true "priority message"
// @Router /sendgroup [get]
func SendGroupChat(c *gin.Context) {
	grpcClient.SendGroupChat(context.Background(), &proto.Mes{Text: c.Query("text"), Priority: c.Query("priority")})
}

// @Summary Send message
// @Description Send message to channel
// @ID send-message-channel
// @Accept  json
// @Produce  json
// @Param text query string true "message text"
// @Param priority query string true "priority message"
// @Router /sendchannel [get]
func SendChannel(c *gin.Context) {
	fmt.Println("aknflnsdgkns")
	grpcClient.SendChannel(context.Background(), &proto.Mes{Text: c.Query("text"), Priority: c.Query("priority")})
}
