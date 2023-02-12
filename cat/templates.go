package cat

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
)

func GetIndexTemplates() {

	response := http_handler.GetRequest("_cat/templates")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
