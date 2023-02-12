package admin

import (
	"fmt"
	"io/ioutil"

	"src/local/oscli/http_handler"
	"src/local/oscli/structs"
)

func AddRoleMapping(rolemapping structs.RoleMapping, name string) {

	resp := http_handler.PutRequest("_opendistro/_security/api/rolesmapping/"+name, rolemapping)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}
