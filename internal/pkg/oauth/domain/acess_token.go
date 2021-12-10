package domain

import (
	"bookstore/pkg/errors/restError"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (a *AccessToken) Validate() *restError.RestErr {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
	if a.AccessToken == "" {
		return restError.NewBadRequestError("invalid access token id")
	}

	if a.UserId <= 0 {
		return restError.NewBadRequestError("invalid user id")
	}

	if a.ClientId <= 0 {
		return restError.NewBadRequestError("invalid client id")
	}

	if a.Expires <= 0 {
		return restError.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now()
	expiration := time.Unix(at.Expires, 0)
	return now.After(expiration)
}
