package jwtTokenProvider

import (
	"RestAPI/component/tokenProvider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenProvider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenProvider.TokenPayload, expiry int) (*tokenProvider.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})
	myToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &tokenProvider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}

//func (j *jwtProvider) Validate(myToken string) *tokenProvider.TokenPayload {
//	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(j.secret), nil
//	})
//	if err != nil {
//		return nil, tokenProvider.ErrorInvalidToken
//	}
//	if !res.Valid {
//		return nil, tokenProvider.ErrorInvalidToken
//	}
//	return nil, nil
//
//}
