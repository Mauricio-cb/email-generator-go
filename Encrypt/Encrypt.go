package Encrypt

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func EncryptPassword() string{
	
	rand.Seed(time.Now().UnixNano())

	hash := sha1.New()
    hash.Write([]byte(string(rune(rand.Intn(3000)))))

	hashString := fmt.Sprintf("%x",  hash.Sum(nil))
	

	return hashString
}

func RandomEmail() string {

	email := nameList() + secondArgumentToEmail() + "@gmail.com"

	return email
}

func nameList() string {
	names := []string{
		"Miguel",
		"Davi",
		"Gabriel",
		"Arthur",
		"Lucas",
		"Matheus",
		"Pedro",
		"Guilherme",
		"Gustavo",
		"Rafael",
		"Felipe",
		"Bernardo",
		"Enzo",
		"Nicolas",
		"JoãoPedro",
		"PedroHenrique",
		"Cauã",
		"Vitor",
		"Eduardo",
		"Daniel",
		"Henrique",
		"Murilo",
		"Vinicius",
		"Samuel",
		"Pietro",
		"João Vitor",
		"Leonardo",
		"Caio",
		"Heitor",
	}
	rand.Seed(time.Now().UnixNano())

	name := names[rand.Intn(len(names))]

	return name
}

func secondArgumentToEmail() string {
	args := []string{
		"javeiro",
		"jogador",
		"odeiaPhp",
		"javascriptAoContrario",
		"jogadordelol",
		"vascaino",
		"flamenguista",
		"rapzeiro",
		"fantantasma",
		"gotico",
	}

	rand.Seed(time.Now().UnixNano())

	arg := args[rand.Intn(len(args))]

	return arg
}
