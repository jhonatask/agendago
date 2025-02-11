package usecase

import "agenda/internal/entity"

type AgendaRepository interface {
    Create(agenda *entity.Agenda) error
    GetAll() ([]entity.Agenda, error)
    GetByID(id uint) (*entity.Agenda, error)
    Update(agenda *entity.Agenda) error
    Delete(id uint) error
}

type AgendaUseCase struct {
    repo AgendaRepository
}

func NewAgendaUseCase(repo AgendaRepository) *AgendaUseCase {
    return &AgendaUseCase{repo: repo}
}

func (uc *AgendaUseCase) CreateAgenda(agenda *entity.Agenda) error {
    return uc.repo.Create(agenda)
}

func (uc *AgendaUseCase) GetAgendas() ([]entity.Agenda, error) {
    return uc.repo.GetAll()
}

func (uc *AgendaUseCase) GetAgendaByID(id uint) (*entity.Agenda, error) {
    return uc.repo.GetByID(id)
}

func (uc *AgendaUseCase) UpdateAgenda(agenda *entity.Agenda) error {
    return uc.repo.Update(agenda)
}

func (uc *AgendaUseCase) DeleteAgenda(id uint) error {
    return uc.repo.Delete(id)
}