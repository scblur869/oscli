package configure

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"src/local/oscli/structs"
)

func createFolderFile(configFile []byte) {

	user, error := user.Current()
	if error != nil {
		panic(error)
	}
	_, err := os.Stat(user.HomeDir + "/.oscli")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(user.HomeDir+"/.oscli", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	e := ioutil.WriteFile(user.HomeDir+"/.oscli/config.json", configFile, 0644)
	if e != nil {
		panic(e)
	}
	fmt.Println("config file written..")

}

func Setup() {
	currentCredentials := GetCredentials()
	currentProfile := "default"
	currentHost := strings.Replace(currentCredentials.Host, "\n", "", -1)
	if currentCredentials.Profile != "" {
		currentProfile = currentCredentials.Profile
	}

	configRead := bufio.NewReader(os.Stdin)
	fmt.Print("AWS Cli Profile [" + currentProfile + "]: ")
	profileTxt, _ := configRead.ReadString('\n')
	if profileTxt == "\n" {
		profileTxt = currentProfile
	} else {
		profileTxt = strings.Replace(profileTxt, "\n", "", -1)
	}

	fmt.Print("opensearch Host [" + currentHost + "]: ")
	hostTxt, _ := configRead.ReadString('\n')
	if hostTxt == "\n" {
		hostTxt = currentHost
	} else {
		hostTxt = strings.Replace(hostTxt, "\n", "", -1)
	}

	configData := structs.ConfigFile{
		Profile: profileTxt,
		Host:    hostTxt,
	}

	fileBytes := new(bytes.Buffer)
	enc := json.NewEncoder(fileBytes)
	enc.SetIndent("", "  ")
	enc.Encode(configData)

	createFolderFile(fileBytes.Bytes())
}

func ReadConfigFile() []byte {
	user, error := user.Current()
	if error != nil {
		panic(error)
	}
	fd, e := ioutil.ReadFile(user.HomeDir + "/.oscli/config.json")
	if e != nil {
		panic(e)
	}
	return fd
}

// gets and returns the credentials from the aws cli config file
func GetCredentials() structs.ConfigFile {
	fileInterface := structs.ConfigFile{}
	config := ReadConfigFile()
	json.Unmarshal(config, &fileInterface)
	return fileInterface
}
