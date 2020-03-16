package main

import (
	"fmt"
	"net/http"

	"github.com/TiagoAlmeida3000/cursogo2/3avançado/6ConexaoComSQL/manipulador"
	"github.com/TiagoAlmeida3000/cursogo2/3avançado/6ConexaoComSQL/repo"
)

func init() {
	fmt.Println("Subindo o servidor")
}

func main() {

	err := repo.AbreConexaoComBancoDeDados()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o baco de dados", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola mundo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Executamdo...")
	http.ListenAndServe(":8181", nil)
}
