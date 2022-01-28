package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJson(fileName string) (result []int, err error) {
	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()

	if err != nil {
		return result, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	parseErr := json.Unmarshal(byteValue, &result)

	return result, parseErr
}
