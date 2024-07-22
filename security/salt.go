package security

import (
	"crypto/rand"
	"strings"
)

// 生成盐
type Salt interface {
	GenerateSalt(length int) (string, error) //生成盐
}

type RandomStrSalt struct {
}

func NewRandomStrSalt() *RandomStrSalt {
	return &RandomStrSalt{}
}

func (r *RandomStrSalt) GenerateSalt(length int) (string, error) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// 生成n个随机字节
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// 将每个字节映射到字符集中的一个字符
	var sb strings.Builder
	for _, b := range bytes {
		sb.WriteByte(letters[int(b)%len(letters)])
	}

	return sb.String(), nil
}
