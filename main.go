package Aibolit

import (
	"encoding/json"
	"os"
	"sort"
)

type contract struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func Do(sourceAddress string, resultAddress string) error {

	patients, err := readPatients(sourceAddress)
	if err != nil {
		return err
	}

	sortByAge(*patients)

	err = writePatients(patients, resultAddress)
	if err != nil {
		return err
	}

	return nil

}

func readPatients(sourceAddress string) (*[]contract, error) {

	f, err := os.Open(sourceAddress)
	if err != nil {
		return &[]contract{}, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	jsonData := make([]contract, 0, 3)

	for decoder.More() {
		var c contract
		err = decoder.Decode(&c)
		if err != nil {
			return &[]contract{}, err
		}
		jsonData = append(jsonData, c)
	}

	return &jsonData, nil

}

func writePatients(patients *[]contract, resultAddress string) error {

	f, err := os.CreateTemp(resultAddress, "result_Json_")
	if err != nil {
		return err
	}
	err = json.NewEncoder(f).Encode(patients)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func sortByAge(patients []contract) {
	sort.Slice(patients, func(i, j int) bool {
		return patients[i].Age < patients[j].Age
	})
}
