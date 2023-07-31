package storage

import (
	"fmt"
	"payments/src/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Payment{}, &entities.OutboxData{})
	if err != nil {
		return nil, err
	}

	return &sqlUserStorage{
		db: db,
	}, nil
}

func (repo sqlUserStorage) Create(payment entities.Payment) error {
	tx := repo.db.Begin()

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	outbox := entities.OutboxData{
		PaymentID: payment.ID,
		RouteID:   payment.RouteID,
		Passenger: payment.Passenger,
	}

	if err := tx.Create(&outbox).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo sqlUserStorage) PopPayment() (entities.OutboxData, error) {
	var outboxDTO entities.OutboxData

	if err := repo.db.First(&outboxDTO).Error; err != nil {
		return entities.OutboxData{}, err
	}

	if err := repo.db.Delete(&outboxDTO).Error; err != nil {
		return entities.OutboxData{}, err
	}

	return outboxDTO, nil
}

func (repo sqlUserStorage) PushBack(payment entities.OutboxData) error {
	if err := repo.db.Create(&payment).Error; err != nil {
		return err
	}

	return nil
}

func (repo sqlUserStorage) DropTables() error {
	return repo.db.Migrator().DropTable(&entities.Payment{}, &entities.OutboxData{})
}
