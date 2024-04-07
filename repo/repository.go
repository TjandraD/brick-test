package repo

import "gorm.io/gorm"

type RepositoryInterface interface {
	CreateTransaction(input Transaction) (Transaction, error)
	UpdateTransactionStatus(transactionId uint, status string) (bool, error)
}

type Repository struct {
	Db *gorm.DB
}

type Transaction struct {
	gorm.Model
	RecipientAccountNumber int64 `gorm:"index"`
	Amount                 int64
	Status                 string `gorm:"index"`
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Transaction{})

	return Repository{Db: db}
}

func (r Repository) CreateTransaction(input Transaction) (Transaction, error) {
	if err := r.Db.Create(&input).Error; err != nil {
		return Transaction{}, err
	}

	return input, nil
}

func (r Repository) UpdateTransactionStatus(transactionId uint, status string) (bool, error) {
	if err := r.Db.Model(&Transaction{}).Where("id = ?", transactionId).Update("status", status).Error; err != nil {
		return false, err
	}

	return true, nil
}
