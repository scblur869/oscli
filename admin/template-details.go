package admin

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
)

func GetTemplateDetails(name string) {

	response := http_handler.GetRequest("_index_template/" + name)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
