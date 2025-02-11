package router

import (
    "agenda/internal/controller"
    "agenda/internal/repository"
    "agenda/internal/usecase"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    agendaRepo := repository.NewAgendaRepository(db)
    agendaUseCase := usecase.NewAgendaUseCase(agendaRepo)
    agendaController := controller.NewAgendaController(agendaUseCase)

    r.GET("/agendas", agendaController.GetAgendas)
    r.POST("/agendas", agendaController.CreateAgenda)
    r.GET("/agendas/:id", agendaController.GetAgendaByID)
    r.PUT("/agendas/:id", agendaController.UpdateAgenda)
    r.DELETE("/agendas/:id", agendaController.DeleteAgenda)

    return r
}