package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

// Usecase ...
type Usecase struct {
	userManager    pkg_user.UserManagerUC
	secret         string
	signingKey     []byte
	expireDuration time.Duration
}

type AuthClaims struct {
	jwt.StandardClaims
	User *pkg_user.User `json:"user"`
}

// NewUsecase ...
func NewUsecase(
	userManager pkg_user.UserManagerUC,
	secret string,
	signingKey []byte,
	tokenTTL time.Duration) *Usecase {
	return &Usecase{
		userManager:    userManager,
		secret:         secret,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

func (u *Usecase) LDAPSignIn(email, password string) (*pkg_user.User, string, error) {
	var (
		user  *pkg_user.User
		token string
	)

	lUser, ok := u.ldapAuthenticate(email, password)
	if !ok {
		return &pkg_user.User{}, "", ErrLoginFailed
	}

	user, err := u.userManager.GetUserByEmail(email)
	if err != nil {
		user = &pkg_user.User{
			Email:      email,
			Name:       lUser.Name,
			Department: lUser.Department,
		}
		_, err := u.userManager.CreateUser(user)
		user, err = u.userManager.GetUserByEmail(email)
		if err != nil {
			return &pkg_user.User{}, "", err
		}
	} else {
		if user.Name != lUser.Name || user.Department != lUser.Department {
			user.Name = lUser.Name
			user.Department = lUser.Department
			u.userManager.CreateUser(user)
		}
	}
	token, err = u.GenerateToken(user)
	if err != nil {
		return &pkg_user.User{}, "", err
	}
	return user, token, nil
}

func (u *Usecase) GenerateToken(user *pkg_user.User) (string, error) {
	var (
		token    *jwt.Token
		strToken string
		err      error
	)

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(u.expireDuration)),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString(u.signingKey)
	if err != nil {
		logger.LogError(ErrCreateToken.Error(), "auth/usecase", user.Name, err)
		return "", ErrCreateToken
	}

	return strToken, nil
}

func (u *Usecase) ParseToken(ctx context.Context, accessToken string) (*pkg_user.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, nil
}

func (u *Usecase) ldapAuthenticate(email, password string) (ldapUser, bool) {
	ldap := NewLdapAuthenticator(
		viper.GetString("app.auth.ldap.server"),
		viper.GetString("app.auth.ldap.baseDN"),
		viper.GetString("app.auth.ldap.filterDN"),
	)
	user, err := ldap.Auth(email, password)
	if err != nil {
		return user, false
	}
	return user, true
}
