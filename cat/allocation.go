package cat

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
)

func GetDiskAllocation() {

	response := http_handler.GetRequest("_cat/allocation?v")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
