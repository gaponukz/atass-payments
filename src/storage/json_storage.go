package storage

import (
	"encoding/json"
	"os"
	"payments/src/entities"
)

type jsonStorage struct {
	filePath string
}

func NewJsonPaymentsStorage(filePath string) jsonStorage {
	return jsonStorage{filePath: filePath}
}

func (s jsonStorage) Create(payment entities.Payment) error {
	payments, err := s.readPaymentsFromFile()
	if err != nil {
		return err
	}

	payments = append(payments, payment)
	err = s.writePaymentsToFile(payments)

	if err != nil {
		return err
	}

	return nil
}

func (s jsonStorage) readPaymentsFromFile() ([]entities.Payment, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var payments []entities.Payment
	err = json.NewDecoder(file).Decode(&payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (s jsonStorage) writePaymentsToFile(payments []entities.Payment) error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(payments)
	if err != nil {
		return err
	}

	return nil
}
