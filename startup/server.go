package startup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igordonatti/REST-API/api/user"
	"github.com/igordonatti/REST-API/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Server() {
	// Conecta ao banco de dados
	database, err := gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	// Inicializa o Gin
	router := gin.Default()

	/// Inicializa o serviço e o controlador de usuário
	userService := user.NewUserService(database)
	userController := user.NewUserController(userService)

	userController.MountRoutes(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: router,
	}

	println("Servidor rodando!")

	if err := server.ListenAndServe(); err != nil {
		panic("Failed to start server.")
	}
}
