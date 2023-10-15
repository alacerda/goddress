package main

import (
	"fmt"
	"github.com/alacerda/goddress/helpers"
	"log"
	"os"
	"time"
)

type resposta struct {
	fonte    string
	endereco string
}

func main() {
	canal := make(chan resposta)

	// apicep.com
	go func() {
		resp, err := helpers.HttpRequest("https://cdn.apicep.com/file/apicep/" + os.Args[1] + ".json")
		if err == nil {
			cep := resposta{
				fonte:    "apicep.com",
				endereco: resp,
			}
			canal <- cep
		}
	}()

	// viacep.com.br
	go func() {
		resp, err := helpers.HttpRequest("http://viacep.com.br/ws/" + os.Args[1] + "/json/")
		if err == nil && len(resp) > 0 {
			cep := resposta{
				fonte:    "viacep.com.br",
				endereco: resp,
			}
			canal <- cep
		}
	}()

	select {
	case endereco := <-canal:
		fmt.Println("Endereco:", endereco.endereco)
		fmt.Println("Fonte:", endereco.fonte)

	case <-time.After(time.Second * 3):
		log.Fatalln("Timeout")
	}
}
