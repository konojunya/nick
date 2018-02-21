package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/konojunya/nick/model"
	scan "github.com/mattn/go-scan"
	"github.com/urfave/cli"
)

// Save save .nick.json file from package.json
func Save(c *cli.Context) {
	raw, err := ioutil.ReadFile("./package.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	js := strings.NewReader(string(raw))

	var dependencies map[string]interface{}
	var devDependencies map[string]interface{}

	scan.ScanJSON(js, "/dependencies", &dependencies)
	scan.ScanJSON(js, "/devDependencies", &devDependencies)

	var dependenciesKeys []string
	var devDependenciesKeys []string
	for key := range dependencies {
		dependenciesKeys = append(dependenciesKeys, key)
	}
	for key := range devDependencies {
		devDependenciesKeys = append(devDependenciesKeys, key)
	}

	output, err := json.Marshal(&model.Json{
		Dependencies:    dependenciesKeys,
		DevDependencies: devDependenciesKeys,
	})
	if err != nil {
		log.Fatal(err)
	}

	context := []byte(output)
	ioutil.WriteFile(".nick.json", context, os.ModePerm)
}

// Load load
func Load(c *cli.Context) {
	raw, err := ioutil.ReadFile(".nick.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var js model.Json
	err = json.Unmarshal(raw, &js)
	if err != nil {
		log.Fatal(err)
	}

	loadModules(js)
}

func loadModules(js model.Json) {
	if len(js.Dependencies) != 0 {
		b, err := exec.Command("npm", "install", "--save", strings.Join(js.Dependencies, " ")).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}

	if len(js.DevDependencies) != 0 {
		b, err := exec.Command("npm", "install", "--save-dev", strings.Join(js.DevDependencies, " ")).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
