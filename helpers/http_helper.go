package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpRequest(url string) (endereco string, erro error) {
	resp, erro := http.Get(url)
	if erro != nil || resp.StatusCode != 200 {
		erro = errors.New("Unsuccessful")
	} else {
		defer resp.Body.Close()
		responseBody, erro := ioutil.ReadAll(resp.Body)
		if erro == nil && !strings.Contains(string(responseBody), "erro\": true") {
			endereco = string(responseBody)
			fmt.Println(resp.StatusCode)
		}
	}

	return
}
