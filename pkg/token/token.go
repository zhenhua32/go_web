package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
)

// 记录登录信息的 JWT
type LoginToken struct {
	jwt.Payload
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

// 签名算法, 随机, 不保存密钥, 每次都是随机的
var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(
	jwt.ECDSAPublicKey(publicKey),
	jwt.ECDSAPrivateKey(privateKey),
)

// 签名
func Sign(id uint, username string) (string, error) {
	now := time.Now()
	pl := LoginToken{
		Payload: jwt.Payload{
			Issuer:         "coolcat",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		ID:       id,
		Username: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

// 验证
func Verify(token []byte) (*LoginToken, error) {
	pl := &LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}
