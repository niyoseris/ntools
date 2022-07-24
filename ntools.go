package ntools

import (
	"io/ioutil"
	"net/http"
)


func WebToText(webPage string) (string, error){
	paginata, err := http.Get(webPage)

	if err != nil {
		return "Error : " , err
	}

	pagina, erro  := ioutil.ReadAll(paginata.Body)

	if erro != nil {
		return "Error : ", erro 
	}

	return string(pagina), nil 
}