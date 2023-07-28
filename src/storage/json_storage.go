package storage

import (
	"encoding/json"
	"os"
	"payments/src/entities"
	"payments/src/errors"
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
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		if err := s.createEmptyFile(); err != nil {
			return nil, err
		}
	}

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

func (s jsonStorage) createEmptyFile() error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	emptyArray := []byte("[]")
	_, err = file.Write(emptyArray)
	if err != nil {
		return err
	}

	return nil
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

func (s jsonStorage) Pop() (entities.Payment, error) {
	payments, err := s.readPaymentsFromFile()
	if err != nil {
		return entities.Payment{}, err
	}

	numPayments := len(payments)
	if numPayments == 0 {
		return entities.Payment{}, errors.ErrStorageEmpty
	}

	payment := payments[numPayments-1]
	payments = payments[:numPayments-1]

	err = s.writePaymentsToFile(payments)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}

func (s jsonStorage) Rollback(payment entities.Payment) error {
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
