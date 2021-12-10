package domain

import (
	"bookstore/errors/restError"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *restError.RestErr)
	Create(AccessToken) *restError.RestErr
	UpdateExpirationTime(AccessToken) *restError.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *restError.RestErr)
	Create(AccessToken) *restError.RestErr
	UpdateExpirationTime(AccessToken) *restError.RestErr
}

type tokenService struct {
	repository Repository
}

func NewTokenService(repo Repository) Service {
	return &tokenService{
		repository: repo,
	}
}

func (t *tokenService) GetById(accessTokenId string) (*AccessToken, *restError.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, restError.NewBadRequestError("invalid access token")
	}

	accessToken, err := t.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *tokenService) Create(at AccessToken) *restError.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *tokenService) UpdateExpirationTime(at AccessToken) *restError.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
