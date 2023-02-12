package admin

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
	"src/local/oscli/structs"
)

func NewRole(role structs.Role, name string) {

	resp := http_handler.PutRequest("_opendistro/_security/api/roles/"+name, role)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}
