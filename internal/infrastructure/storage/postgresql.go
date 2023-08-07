package storage

import (
	"fmt"
	"payments/internal/domain/entities"
	"payments/internal/domain/errors"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresCredentials struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
}

type sqlUserStorage struct {
	db *gorm.DB
}

func NewPostgresUserStorage(c PostgresCredentials) (*sqlUserStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.Dbname, c.Port, c.Sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&paymentModel{}, &outboxDataModel{})
	if err != nil {
		return nil, err
	}

	return &sqlUserStorage{
		db: db,
	}, nil
}

func (repo sqlUserStorage) Create(payment entities.Payment) error {
	payModel := paymentToModel(payment)

	tx := repo.db.Begin()
	if err := tx.Create(&payModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	outbox := outboxDataModel{
		CreatedAt:   time.Now(),
		PaymentID:   payment.ID,
		RouteID:     payment.RouteID,
		PassengerID: payment.Passenger.ID,
		Passenger:   passengerToModel(payment.Passenger),
	}

	if err := tx.Create(&outbox).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo sqlUserStorage) PopPayment() (entities.OutboxData, error) {
	var outboxDTO outboxDataModel

	if err := repo.db.Order("created_at").Preload("Passenger").First(&outboxDTO).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			return entities.OutboxData{}, errors.ErrStorageEmpty
		}

		return entities.OutboxData{}, err
	}

	if err := repo.db.Delete(&outboxDTO).Error; err != nil {
		return entities.OutboxData{}, err
	}

	return outboxDataFromModel(outboxDTO), nil
}

func (repo sqlUserStorage) PushBack(payment entities.OutboxData) error {
	model := outboxDataToModel(payment)
	model.CreatedAt = time.Now()
	return repo.db.Create(&model).Error
}

func (repo sqlUserStorage) DropTables() error {
	return repo.db.Migrator().DropTable(&paymentModel{}, &outboxDataModel{})
}
