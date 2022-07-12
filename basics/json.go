package basics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// UserList struct which contains an array of users
type UserList struct {
	Users []User `json:"users"`
}

// User struct which contains a name a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func getJsonFile() (*os.File, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Open our jsonFile
	jsonFile, err := os.Open(getwd + "/basics/users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully Opened users.json")
	return jsonFile, nil
}
func JSONLearn() {
	jsonFile, err := getJsonFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var list UserList
	// we unmarshal our byteArray which contains our jsonFile's content into 'list' which we defined above
	err = json.Unmarshal(byteValue, &list)
	if err != nil {
		return
	}

	fmt.Println(string(byteValue), " list : ", list) // look that, `type` field is missing.
	for i := 0; i < len(list.Users); i++ {
		fmt.Println("Age: " + strconv.Itoa(list.Users[i].Age))
		fmt.Println("Name: " + list.Users[i].Name)
	}

}
