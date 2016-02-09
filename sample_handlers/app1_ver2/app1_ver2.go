package app1_ver2

//Our new app with language support

import(
	"net/http"
	"fmt"
)

func SayHelloV2(r http.ResponseWriter, w *http.Request){
	//do something
	fmt.Println("Hello World with language support!")
}

func SayHolaV1(r http.ResponseWriter, w *http.Request) {
	//do something
	fmt.Println("Hola Mundo!")
}

func SayNamaste(r http.ResponseWriter, w *http.Request) {
	//do something
	fmt.Println("Namaste Duniya!")
}
