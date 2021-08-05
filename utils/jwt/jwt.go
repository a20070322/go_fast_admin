package jwt

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserGroupType = string

const UserGroupCustomer UserGroupType = "customer"
const UserGroupAdmin UserGroupType = "admin"
const UserGroupAll UserGroupType = "all"

func GetTokenId(ctx *gin.Context) (string, error) {
	if cUid, uidBool := ctx.Get("uid"); uidBool == true {
		uid := cUid.(string)
		return uid, nil
	} else {
		return "", errors.New("未找到此用户")
	}
}

type Claims struct {
	jwt.StandardClaims
	UserID         string           `json:"user_id"`
	UserGroup      UserGroupType `json:"user_group"`
	IsRefreshToken bool          `json:"is_refresh_token"`
}

type AuthReturn struct {
	Token            string `json:"token"`
	ExpiresAt        int64  `json:"expires_at"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresAt int64  `json:"refresh_expires_at"`
}

// 将token缓存至redis
func SetCatChToken(claims *Claims, token string, timer time.Duration) error {
	str := ""
	if claims.IsRefreshToken == true {
		str = "refresh_"
	}
	_, err := global.Rdb.Set(context.Background(), str+"token_"+claims.UserGroup+"_"+claims.UserID, token, timer).Result()
	return err
}

// 检验token是否在缓存汇总
func CheckTokenCatch(claims *Claims, token string, IsRefreshToken bool) (bool, error) {
	str := ""
	if IsRefreshToken == true {
		str = "refresh_"
	}
	catchToken, err := global.Rdb.Get(context.Background(), str+"token_"+claims.UserGroup+"_"+claims.UserID).Result()
	if err != nil {
		return false, err
	}
	return catchToken == token, nil
}

// 获取Token
func GenToken(claims *Claims) (*AuthReturn, error) {
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second*time.Duration(global.AppSetting.Jwt.TokenExpireDuration)).UnixNano() / 1e6
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(global.AppSetting.Jwt.JwtSecret))
	if err != nil {
		return nil, err
	}
	refreshClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: claims.ExpiresAt + (86400 * 7),
		},
		UserID:         claims.UserID,
		UserGroup:      claims.UserGroup,
		IsRefreshToken: true,
	}
	refreshToken, err2 := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(global.AppSetting.Jwt.JwtSecret))
	if err2 != nil {
		return nil, err2
	}
	jwtAuthReturn := &AuthReturn{}
	jwtAuthReturn.Token = token
	jwtAuthReturn.ExpiresAt = claims.ExpiresAt
	jwtAuthReturn.RefreshToken = refreshToken
	jwtAuthReturn.RefreshExpiresAt = refreshClaims.ExpiresAt
	err3 := SetCatChToken(claims, token, time.Second*time.Duration(global.AppSetting.Jwt.TokenExpireDuration))
	if err3 != nil {
		global.Logger.Error(err3.Error())
		return nil, err3
	}
	err4 := SetCatChToken(refreshClaims, refreshToken, time.Second*time.Duration(global.AppSetting.Jwt.TokenExpireDuration+(86400*7)))
	if err4 != nil {
		global.Logger.Error(err4.Error())
		return nil, err4
	}
	return jwtAuthReturn, nil
}

//验证jwt token
func VerifyAction(strToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(strToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.AppSetting.Jwt.JwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("token expired")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("jwt analysis error")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("jwt invalid")
	}
	return claims, nil
}
