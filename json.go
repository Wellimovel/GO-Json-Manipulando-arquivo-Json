package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	Nome   string `json: "Nome"`
	Cargo  string `json: "Cargo"`
	Idade  int    `json: "Idade"`
	Social Social `json: "Social"`
}
type Social struct {
	Facebook string `json: "facebook"`
	Linkedin string `json: "Linkedin"`
}

func main() {
	jsonFile, err := os.Open("json.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sucesso")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var usuarios Usuarios
	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuario Nome: " + usuarios.Usuarios[i].Nome)
		fmt.Println("Usuario Idade: " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Usuario Cargo: " + usuarios.Usuarios[i].Cargo)
		fmt.Println("Usuario Soacial: " + usuarios.Usuarios[i].Social.Facebook)
		fmt.Println("Usuario Soacial: " + usuarios.Usuarios[i].Social.Linkedin)
	}
}
