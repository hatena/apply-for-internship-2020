package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"io/ioutil"
	"time"
)

func askDoYouWantToApply() (bool, error) {
	confirmation := false
	prompt := &survey.Confirm{
		Message: "Do you want to apply for our internship?",
	}
	if err := survey.AskOne(prompt, &confirmation); err != nil {
		return confirmation, err
	}
	return confirmation, nil
}

func askName() (string, error) {
	name := ""
	prompt := &survey.Input{
		Message:  "What is your name?",
	}
	if err := survey.AskOne(prompt, &name); err != nil {
		return name, err
	}
	return name, nil
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	b, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("invalid private key data")
	}
	return x509.ParseECPrivateKey(block.Bytes)
}

func loadPublicKey() (*ecdsa.PublicKey, error) {
	b, err := ioutil.ReadFile("./public.pem")
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("invalid public key data")
	}
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	var publicKey *ecdsa.PublicKey
	var ok bool
	if publicKey, ok = parsedKey.(*ecdsa.PublicKey); !ok {
		return nil, fmt.Errorf("data doesn't contain valid ECDSA Public Key")
	}

	return publicKey, nil
}

func generateToken(iss string) ([]byte, error) {
	claims := jwt.New()
	if err := claims.Set(jwt.IssuerKey, iss); err != nil {
		return nil, err
	}
	if err := claims.Set(jwt.SubjectKey, "I apply."); err != nil {
		return nil, err
	}
	if err := claims.Set(jwt.AudienceKey, "Hatena Co., Ltd."); err != nil {
		return nil, err
	}
	if err := claims.Set(jwt.IssuedAtKey, time.Now()); err != nil {
		return nil, err
	}

	privateKey, err := loadPrivateKey()
	if err != nil {
		return nil, err
	}

 	token, err := jwt.Sign(claims, jwa.ES256, privateKey)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func verify(token []byte) (jwt.Token, error) {
	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, err
	}
	_, err = jws.Verify(token, jwa.ES256, publicKey)
	if err != nil {
		return nil, err
	}
	return jwt.ParseBytes(token)
}

func main() {
	wantToApply, err := askDoYouWantToApply()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !wantToApply {
		fmt.Println("See you again.")
		return
	}
	fmt.Println("Thank you!")

	name, err := askName()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Hello, %s.\n", name)

	token, err := generateToken(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Your token is here:\n%s", token)
}
