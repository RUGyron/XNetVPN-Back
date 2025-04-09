package models

import (
	"XNetVPN-Back/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var key = []byte(config.Config.JwtKey)

func makeToken(userId string, expDate int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    expDate,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *Tokens) resetTokens() {
	j.AccessToken = ""
	j.RefreshToken = ""
}

func (j *Tokens) GenerateTokens(userId string) error {
	accessToken, err := makeToken(userId, time.Now().Add(time.Second*time.Duration(config.Config.JwtAccessExpiration)).Unix())
	if err != nil {
		j.resetTokens()
		return err
	}

	refreshToken, err := makeToken(userId, time.Now().Add(time.Second*time.Duration(config.Config.JwtRefreshExpiration)).Unix())
	if err != nil {
		j.resetTokens()
		return err
	}

	j.AccessToken = accessToken
	j.RefreshToken = refreshToken

	return nil
}

func (j *Tokens) UpdateAccessToken(userId string) error {
	err := j.ValidateRefreshToken()
	if err != nil {
		j.resetTokens()
		return err
	}

	accessToken, err := makeToken(userId, time.Now().Add(time.Minute*time.Duration(config.Config.JwtAccessExpiration)).Unix())
	if err != nil {
		j.resetTokens()
		return err
	}

	j.AccessToken = accessToken

	return err
}

func (j *Tokens) ValidateAccessToken() error {
	token, err := jwt.Parse(j.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid access token")
	}

	return nil
}

func (j *Tokens) ValidateRefreshToken() error {
	token, err := jwt.Parse(j.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid refresh token")
	}

	return nil
}
