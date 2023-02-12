package cat

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
)

func GetShardsInfo() {

	response := http_handler.GetRequest("_cat/shards?v")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
