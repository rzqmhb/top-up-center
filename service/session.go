package service

import (
	"github.com/rzqmhb/top-up-center/models"
	repo "github.com/rzqmhb/top-up-center/repository"
)

type SessionService interface {
	GetByUsername(username string) (*models.Session, error)
}

type sessionService struct {
	sessionRepository repo.SessionRepository
}

func NewSessionService(sessionRepository repo.SessionRepository) SessionService {
	return &sessionService{sessionRepository: sessionRepository}
}

func (s *sessionService) GetByUsername(username string) (*models.Session, error) {
	return s.sessionRepository.GetByUsername(username)
}