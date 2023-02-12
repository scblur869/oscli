package cat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"src/local/oscli/http_handler"
)

func GetClusterDefaults() {

	response := http_handler.GetRequest("_cluster/settings?include_defaults=true")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Println(string(prettyJSON.Bytes()))
}
