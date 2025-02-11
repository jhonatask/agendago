package controller

import (
    "agenda/internal/entity"
    "agenda/internal/usecase"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type AgendaController struct {
    useCase *usecase.AgendaUseCase
}

func NewAgendaController(useCase *usecase.AgendaUseCase) *AgendaController {
    return &AgendaController{useCase: useCase}
}

func (ctrl *AgendaController) CreateAgenda(c *gin.Context) {
    var agenda entity.Agenda
    if err := c.ShouldBindJSON(&agenda); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.useCase.CreateAgenda(&agenda); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, agenda)
}

func (ctrl *AgendaController) GetAgendas(c *gin.Context) {
    agendas, err := ctrl.useCase.GetAgendas()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, agendas)
}

func (ctrl *AgendaController) GetAgendaByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    agenda, err := ctrl.useCase.GetAgendaByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Agenda not found"})
        return
    }
    c.JSON(http.StatusOK, agenda)
}

func (ctrl *AgendaController) UpdateAgenda(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var agenda entity.Agenda
    if err := c.ShouldBindJSON(&agenda); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    agenda.ID = uint(id)
    if err := ctrl.useCase.UpdateAgenda(&agenda); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, agenda)
}

func (ctrl *AgendaController) DeleteAgenda(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := ctrl.useCase.DeleteAgenda(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Agenda deleted"})
}