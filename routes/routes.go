package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/holanda-lucas/gotasks-back/controllers"
	"github.com/holanda-lucas/gotasks-back/middlewares"
)

func LoadRoutes() {
	r := gin.Default()

	// Rotas de usuário
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", middlewares.AuthMiddleware, controllers.EditUser)
	r.DELETE("/users/:id", middlewares.AuthMiddleware, controllers.DeleteUser)
	r.GET("/users/:id/tasks", middlewares.AuthMiddleware, middlewares.AuthUserOnlyMiddleware, controllers.GetUserTasks)

	// Rotas de tarefas
	r.POST("/tasks", middlewares.AuthMiddleware, controllers.CreateTask)
	r.GET("/tasks/:id", middlewares.AuthMiddleware, controllers.GetTask)
	r.PUT("/tasks/:id", middlewares.AuthMiddleware, controllers.EditTask)
	r.DELETE("/tasks/:id", middlewares.AuthMiddleware, controllers.DeleteTask)

	// Rotas de autenticação
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.Run(":8000")
}