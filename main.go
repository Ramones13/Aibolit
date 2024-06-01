package Aibolit

import (
	"encoding/json"
	"os"
)

type contract struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func Do(sourceAddress string, resultAddress string) error {
	f, err := os.Open(sourceAddress)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	jsonData := make([]contract, 0, 3)

	for decoder.More() {
		var c contract
		err = decoder.Decode(&c)
		if err != nil {
			return err
		}
		jsonData = append(jsonData, c)
	}

	//
	f1, err := os.CreateTemp(resultAddress, "result_Json_")
	if err != nil {
		return err
	}
	err = json.NewEncoder(f1).Encode(jsonData)
	if err != nil {
		return err
	}
	err = f1.Close()
	if err != nil {
		return err
	}

	return nil

}
