package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/util/log"
)

type LoginClaims struct {
	Id        string `json:"id,omitempty"`
	Role      int    `json:"role,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	LiftTime  int64  `json:"lift_time,omitempty"`
	LoginTime int64  `json:"login_time,omitempty"`
}

func (c LoginClaims) Valid() error {
	if c.LiftTime == 0 {
		return errors.New("lift time is zero")
	}
	return nil
}

func Encode(c LoginClaims, secret string) (string, error) {
	if c.LiftTime == 0 {
		c.LiftTime = 7200
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Decode(tokenStr string, secret string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		log.Infof("uid: %s, role: %v", claims.Id, claims.Role)
		return claims, nil
	}
	return nil, err
}
