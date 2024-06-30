package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

const (
	expiration_time_minute = 20
	expiration_time_hour   = 60
)

var (
	jwtUtil *JwtUtil
	once    sync.Once
)

type JwtUtil struct {
}

type User struct {
	Id       int64
	UserType int32
}

// 加密的秘钥
var singKey []byte

type UserClaims struct {
	User
	jwt.RegisteredClaims
}

func NewJwtUtil() *JwtUtil {
	once.Do(func() {
		jwtUtil = &JwtUtil{}
	})
	return jwtUtil
}

// 生成一个sh256加密的jwt
func (j *JwtUtil) NewHs256Token(user User) (string, error) {
	userClaims := UserClaims{user, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiration_time_minute) * time.Hour)),
	}}
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
		exat := claims.ExpiresAt //拿到过期时间
		//判断是否要续签
		if exat.Time.Unix() >= time.Now().Unix() &&
			exat.Unix()-time.Now().Unix() <= 60*3 {
			newtoken, err := j.NewHs256Token(claims.User)
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
			Id:       claims.Id,
			UserType: claims.UserType,
		}
		return user, nil
	}
	return nil, errors.New("invalid token or token claims")
}

func init() {
	singKey = []byte("chjksdvbudfvsdvvf")
}
