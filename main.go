package Aibolit

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
)

type patient struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Email string `xml:"Email,omitempty"`
}

type patients struct {
	List []patient `xml:"Patient"`
}

func Do(sourceAddress string, resultAddress string) error {

	patientsData, err := readPatients(sourceAddress)
	if err != nil {
		return err
	}

	sortByAge(*patientsData)

	err = writePatients(patientsData, resultAddress)
	if err != nil {
		return err
	}

	return nil

}

func readPatients(sourceAddress string) (*[]patient, error) {

	f, err := os.Open(sourceAddress)
	if err != nil {
		return &[]patient{}, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	jsonData := make([]patient, 0, 3)

	for decoder.More() {
		var c patient
		err = decoder.Decode(&c)
		if err != nil {
			return &[]patient{}, err
		}
		jsonData = append(jsonData, c)
	}

	return &jsonData, nil

}

func writePatients(patientsData *[]patient, resultAddress string) error {

	p := patients{
		List: *patientsData,
	}

	f, err := os.Create(resultAddress)
	if err != nil {
		return err
	}
	f.WriteString(xml.Header)
	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(p)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func sortByAge(patients []patient) {
	sort.Slice(patients, func(i, j int) bool {
		return patients[i].Age < patients[j].Age
	})
}
