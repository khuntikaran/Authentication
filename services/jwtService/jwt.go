package jwtservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(username string) (string, error) {
	var err error
	jwt_key := []byte("M@Her7505")

	jtCliams := jwt.MapClaims{}
	jtCliams["authorized"] = true
	jtCliams["username"] = username
	jtCliams["expiration_time"] = time.Now().Add(time.Minute * 5).Unix()
	jt := jwt.NewWithClaims(jwt.SigningMethodHS256, jtCliams)
	token, err := jt.SignedString(jwt_key)
	if err != nil {
		fmt.Println(err)
	}
	return token, nil
}

type Tokens struct {
	AccessToken  string
	Acc_Exp_Time int64
	RefreshToken string
	Ref_Exp_Time int64
}

func CreateTokenPair(username string) (*Tokens, error) {
	var err error

	td := &Tokens{}

	td.Acc_Exp_Time = time.Now().Add(time.Minute * 2).Unix()
	td.Ref_Exp_Time = time.Now().Add(time.Minute * 5).Unix()
	//Creating access token
	acc_key := []byte("M@Her7505")
	acClaims := jwt.MapClaims{}
	acClaims["authorized"] = true
	acClaims["username"] = username
	acClaims["acc_ExpirationTime"] = td.Acc_Exp_Time
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, acClaims)
	td.AccessToken, err = at.SignedString(acc_key)
	if err != nil {
		fmt.Println(err)
	}

	//Creating refresh token
	ref_key := []byte("R@jsh@kh@5050")
	refClaims := jwt.MapClaims{}
	refClaims["authorized"] = true
	refClaims["username"] = username
	refClaims["ref_ExpirationTime"] = td.Ref_Exp_Time
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaims)
	td.RefreshToken, err = rt.SignedString(ref_key)
	if err != nil {
		fmt.Println(err)
	}

	return td, nil

}

type TokenS struct {
	token string
}

func CheckRefExpTime(token_exp_time int64) bool {
	if time.Unix(token_exp_time, 0).Sub(time.Now()) > time.Second*60 {
		return true
	}
	fmt.Println("this is expire checker function")
	return false
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("authorized")
	rfmap := map[string]string{}
	tokeen := TokenS{}
	fmt.Println(r.Body)
	to := json.NewDecoder(r.Body).Decode(&rfmap)
	fmt.Println(to)
	fmt.Println(token, "this is the token we get from extract function")
	st := strings.Split(tokeen.token, ".")
	if len(st) == 2 {
		fmt.Println(st)
		return st[1]
	}
	fmt.Println("this is extract function", st)
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	var err error

	rfmap := map[string]string{}
	err = json.NewDecoder(r.Body).Decode(&rfmap)
	if err != nil {
		fmt.Println(err)
	}
	rfToken := rfmap["refresh_token"]
	fmt.Println(rfToken, "this is valid function")
	//ref_key := []byte("R@jsh@kh@5050")
	tokenn, err := jwt.Parse(rfToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header)

		}
		return []byte("R@jsh@kh@5050"), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := tokenn.Claims.(jwt.Claims); !ok && !tokenn.Valid {
		return nil, fmt.Errorf("token not valid")
	}
	return tokenn, nil
}

func TokenValid(r *http.Request) (bool, string, error) {
	var err error
	rfToken, err := VerifyToken(r)
	if err != nil {
		fmt.Println("this is error occuring", err)
	}
	claims, ok := rfToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(ok, "something went wrong ")
	}
	exp_time := claims["ref_ExpirationTime"].(float64)
	username := claims["username"].(string)
	exp := int64(exp_time)
	t := CheckRefExpTime(exp)
	fmt.Println(t)

	return t, username, nil

}
