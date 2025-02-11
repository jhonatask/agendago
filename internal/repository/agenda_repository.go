package repository

import (
    "agenda/internal/entity"
    "gorm.io/gorm"
)

type AgendaRepository struct {
    db *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) *AgendaRepository {
    return &AgendaRepository{db: db}
}

func (r *AgendaRepository) Create(agenda *entity.Agenda) error {
    return r.db.Create(agenda).Error
}

func (r *AgendaRepository) GetAll() ([]entity.Agenda, error) {
    var agendas []entity.Agenda
    err := r.db.Find(&agendas).Error
    return agendas, err
}

func (r *AgendaRepository) GetByID(id uint) (*entity.Agenda, error) {
    var agenda entity.Agenda
    err := r.db.First(&agenda, id).Error
    return &agenda, err
}

func (r *AgendaRepository) Update(agenda *entity.Agenda) error {
    return r.db.Save(agenda).Error
}

func (r *AgendaRepository) Delete(id uint) error {
    return r.db.Delete(&entity.Agenda{}, id).Error
}