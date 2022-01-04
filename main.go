package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ecoprohcm/DMS_BackendServer/docs"
	"github.com/ecoprohcm/DMS_BackendServer/handlers"
	"github.com/ecoprohcm/DMS_BackendServer/models"
	"github.com/ecoprohcm/DMS_BackendServer/mqttSvc"
	"github.com/gin-gonic/gin" // swagger embed files
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func initSwagger(r *gin.Engine) {
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	// programmatically set swagger info
	// Note: Use config later
	docs.SwaggerInfo.Title = "DMS Backend API"
	// docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Description = "This is DMS backend server"
	docs.SwaggerInfo.Host = "http://iot.hcmue.space:8002"
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)

	// DI process
	gwSvc := models.NewGatewaySvc(db)
	areaSvc := models.NewAreaSvc(db)
	dlSvc := models.NewDoorlockSvc(db)
	glSvc := models.NewLogSvc(db)
	pwSvc := models.NewPasswordSvc(db)
	sSvc := models.NewStudentSvc(db)
	eSvc := models.NewEmployeeSvc(db)

	mqttHost := os.Getenv("SERVER_HOST")
	mqttPort := os.Getenv("MQTT_PORT")

	mqttClient := mqttSvc.MqttClient(mqttHost, mqttPort, glSvc, dlSvc, gwSvc)

	gwHdlr := handlers.NewGatewayHandler(gwSvc, mqttClient)
	areaHdlr := handlers.NewAreaHandler(areaSvc)
	dlHdlr := handlers.NewDoorlockHandler(dlSvc, mqttClient)
	glHdlr := handlers.NewGatewayLogHandler(glSvc)
	pwHdlr := handlers.NewPasswordHandler(pwSvc, mqttClient)
	sHdlr := handlers.NewStudentHandler(sSvc, mqttClient)
	eHdlr := handlers.NewEmployeeHandler(eSvc, mqttClient)

	// HTTP Serve
	r := setupRouter(gwHdlr, areaHdlr, dlHdlr, glHdlr, pwHdlr, sHdlr, eHdlr)
	initSwagger(r)
	r.Run(":8080")
	mqttClient.Disconnect(250)
}
