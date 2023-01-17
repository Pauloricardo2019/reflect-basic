package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    uint64 `tamanho:"1" required:"false"`
	Name  string `tamanho:"4" required:"true"`
	Email string `tamanho:"4" required:"true"`
	Age   uint8  `tamanho:"5" required:"false"`
}

func main() {
	//Gerar um objeto usuário qualquer
	//
	userTest := generateUser()

	//Faz o parse da struct para o tipo "reflect.Value"
	v := reflect.ValueOf(userTest)

	//Sempre que for um ponteiro usar o metodo "Elem()"

	//Verifica se o tipo parseado é realmente uma struct
	if v.Elem().Kind() != reflect.Struct {
		fmt.Println("V don't is a struct")
	}

	//Faz um looping dos campos da struct, campo a campo
	for i := 0; i < v.Elem().NumField(); i++ {

		//Captura o field type de cada campo
		fieldType := v.Elem().Type().Field(i)
		fmt.Println("Campo type:", fieldType)
		//Captura o valor bruto de cada campo, para converter teria que usar por exemplo: fielValue.String()
		fielValue := v.Elem().Field(i)

		//Função para capturar os valores das tags de cada campo
		t, r := getTag(fieldType)

		fmt.Println(fmt.Sprintf("Valor da tag tamanho:%s, valor da tag required %s", t, r))

		//Validação para verificar se Nome do campo é "Name"
		if fieldType.Name == "Name" {
			fmt.Println("O nome e: ", fielValue.String())
			//Alterando dinamicamente o valor do campo, que antes era "Paulo", para "Douglas"
			fielValue.SetString("Douglas")
		}
	}

	fmt.Println(fmt.Sprintf("Nome: %s", userTest.Name))

}

func getTag(f reflect.StructField) (string, string) {
	return f.Tag.Get("tamanho"), f.Tag.Get("required")
}

func generateUser() *User {
	return &User{
		ID:    uint64(1),
		Name:  "Paulo",
		Email: "paulo@gmail.com",
		Age:   22,
	}
}
