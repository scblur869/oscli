package cat

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"

	color "github.com/logrusorgru/aurora"
)

func GetHealthInfo() {

	response := http_handler.GetRequest("_cat/health?v")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(color.Bold(string(body)))
}
