package main

import (
	_ "go_practice/docs"
	"go_practice/test" // test 패키지 import

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// test 패키지의 핸들러를 라우팅
	r.GET("/test/ping", test.PingHandler)

	// Swagger 엔드포인트 설정
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 서버 실행
	r.Run(":8080")
}
