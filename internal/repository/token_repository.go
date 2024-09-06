package repository

import (
	"dalabio/internal/entity"
)

type TokenRepository interface {
	FindByToken(token string) (*entity.Token, error)
	Create(token *entity.Token) error
}
