package db

import (
	"bookstore/internal/pkg/oauth/clients/cassandra"
	"bookstore/internal/pkg/oauth/domain"
	"bookstore/pkg/errors/restError"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? where access_token=?;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetById(string) (*domain.AccessToken, *restError.RestErr)
	Create(token domain.AccessToken) *restError.RestErr
	UpdateExpirationTime(domain.AccessToken) *restError.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*domain.AccessToken, *restError.RestErr) {
	var result domain.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, restError.NewNotFoundError("no access token found with given id")
		}
		return nil, restError.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at domain.AccessToken) *restError.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at domain.AccessToken) *restError.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken).Exec(); err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	return nil
}
