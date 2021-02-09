package base_common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

const  (
	privKeyPath = "keys/app.rsa"
	pubKeyPath = "keys/app.rsa.pub"
 )

var (
	verifyKey, signKey []byte
)

func initKeys()  {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("[initKeys]: %s\n", err)
		panic(err)
	}
}

func GenerateToken(name, role string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["iss"]  = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	t.Claims["exp"] = time.Now().Add(time.Hour *24).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
