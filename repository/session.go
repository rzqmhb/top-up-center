package repository

import (
	"time"

	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/models"
)

type SessionRepository interface {
	Store(session *models.Session) error
	GetAll() (*[]models.Session, error)
	GetByToken(token string) (*models.Session, error)
	GetByUsername(username string) (*models.Session, error)
	Update(username string, session *models.Session) error
	Delete(token string) error
	IsExpired(session *models.Session) bool
}

type sessionRepository struct {
	postgresDB *database.PostgresDB
}

func NewSessionRepository(postgresDB *database.PostgresDB) SessionRepository {
	return &sessionRepository{postgresDB: postgresDB}
}

func (s *sessionRepository) Store(session *models.Session) error  {
	return s.postgresDB.StoreSession(session)
}

func (s *sessionRepository) GetAll() (*[]models.Session, error)  {
	return s.postgresDB.FetchSessions()
}

func (s *sessionRepository) GetByToken(token string) (*models.Session, error)  {
	return s.postgresDB.FetchSessionByToken(token)
}

func (s *sessionRepository) GetByUsername(username string) (*models.Session, error)  {
	return s.postgresDB.FetchSessionByUsername(username)
}

func (s *sessionRepository) Update(username string, session *models.Session) error  {
	return s.postgresDB.UpdateSession(username, session)
}

func (s *sessionRepository) Delete(token string) error  {
	return s.postgresDB.DeleteSession(token)
}

func (s *sessionRepository) IsExpired(session *models.Session) bool {
	return session.Expiry.Before(time.Now())
}