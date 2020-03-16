package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/TiagoAlmeida3000/cursogo2/3avançado/2POST/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Tiago"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar p JSON do objeto usuario. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.net/r/s8ckmzs8", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para requestbin. Erro: ",
			err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do request bin. Erro: ",
			err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo retornado pelo request bin. Erro: ",
				err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
