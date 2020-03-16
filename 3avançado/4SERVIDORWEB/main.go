package main

import (
	"fmt"
	"net/http"

	"github.com/TiagoAlmeida3000/cursogo2/3avançado/4SERVIDORWEB/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Executamdo...")
	http.ListenAndServe(":8181", nil)
}
