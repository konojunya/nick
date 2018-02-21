package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/konojunya/nick/model"
	scan "github.com/mattn/go-scan"
	"github.com/urfave/cli"
)

func Save(c *cli.Context) error {
	raw, err := ioutil.ReadFile("./package.json")
	if err != nil {
		return err
	}

	js := strings.NewReader(string(raw))

	var modules map[string]interface{}

	err = scan.ScanJSON(js, "/dependencies", &modules)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var keys []string

	for key := range modules {
		keys = append(keys, key)
	}
	output, err := json.Marshal(&model.Json{
		Dependencies: keys,
	})
	if err != nil {
		return err
	}

	context := []byte(output)
	ioutil.WriteFile(".nick.json", context, os.ModePerm)

	return nil
}

func Install(c *cli.Context) error {
	return nil
}
