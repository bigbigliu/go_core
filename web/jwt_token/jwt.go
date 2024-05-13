package jwt_token

import (
	"time"

	goJwt "github.com/golang-jwt/jwt/v5"
)

// IappJwt ...
type IappJwt interface {
	NewToken(username string) (string, error)
	ParseToken(token string) (*Claims, error)
}

// Claims ...
type Claims struct {
	Username string `json:"username"`
	goJwt.RegisteredClaims
}

// appJWT ...
type appJWT struct {
	Secret  string `json:"secret"`  // jwt密钥
	Timeout int    `json:"timeout"` // jwt过期时间
}

// NewToken 生成新token
func (c *appJWT) NewToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(c.Timeout) * time.Second)

	claims := Claims{
		Username: username,
		RegisteredClaims: goJwt.RegisteredClaims{
			Issuer:    "app_os",                         // 发签人
			ExpiresAt: goJwt.NewNumericDate(expireTime), // 过期时间
			NotBefore: goJwt.NewNumericDate(nowTime),    // 生效时间
			IssuedAt:  goJwt.NewNumericDate(nowTime),    // 签发时间
		},
	}

	t := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(c.Secret))
	return token, err
}

// ParseToken 解析token
func (c *appJWT) ParseToken(token string) (*Claims, error) {
	jwtSecret := []byte(c.Secret)
	tokenClaims, err := goJwt.ParseWithClaims(token, &Claims{}, func(token *goJwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
