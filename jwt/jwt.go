package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	expiration_time_minute  = 1200
	expiration_renewal_time = 3600
	renewal_time_minute     = 180
)

type JwtUtil struct {
	ExpiresAt            int64   //过期时间，秒
	ExpiresRenewalTimeAt int64   //续签过期时间
	RenewalTime          float64 //续签时间
}

type User struct {
	Id   int64
	Data any
}

// 加密的秘钥
var singKey []byte

type UserClaims struct {
	User
	jwt.RegisteredClaims
}

func NewJwtUtil() *JwtUtil {
	return &JwtUtil{
		ExpiresAt:            expiration_time_minute,
		ExpiresRenewalTimeAt: renewal_time_minute,
		RenewalTime:          expiration_renewal_time,
	}
}

// 生成一个sh256加密的jwt
func (j *JwtUtil) NewHs256Token(user User) (string, error) {
	return j.newHs256Token(user, 0)
}

func (j *JwtUtil) newHs256Token(user User, pattern int) (string, error) {
	var userClaims UserClaims
	switch pattern {
	case 0:
		userClaims = UserClaims{user, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresAt) * time.Second)),
		}}
	case 1:
		userClaims = UserClaims{user, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresRenewalTimeAt) * time.Second)),
		}}
	default:
		return "", errors.New("Pattern does not exist")
	}

	//1、生成token
	//使用指定的签名方法创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	//2、把token加密
	ss, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (j *JwtUtil) Hs256Parse(token_ string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(token_, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return singKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JwtUtil) Hs256Verify(token_ string) bool {
	token, err := j.Hs256Parse(token_)
	if err != nil {
		return false
	}
	if token.Valid {
		return true
	}

	return false
}

// 续签token
func (j *JwtUtil) Hs256RefreshToken(token_ string) (string, error) {
	token, err := j.Hs256Parse(token_)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok {
		now := time.Now()
		expiresAt := claims.ExpiresAt.Time //拿到过期时间
		//判断是否要续签
		if expiresAt.After(now) &&
			expiresAt.Sub(now).Seconds() <= j.RenewalTime {
			newtoken, err := j.newHs256Token(claims.User, 1)
			if err != nil {
				return "", err
			}
			return newtoken, nil
		}

	}
	return token_, nil
}

// 修改签名秘钥
func InitSingKey(singKey1 string) {
	singKey = []byte(singKey1)
}

func (j *JwtUtil) ParseUserFromToken(token_ string) (*User, error) {
	token, err := j.Hs256Parse(token_)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		user := &User{
			Id:   claims.Id,
			Data: claims.Data,
		}
		return user, nil
	}
	return nil, errors.New("invalid token or token claims")
}

func init() {
	singKey = []byte("chjksdvbudfvsdvvf")
}
