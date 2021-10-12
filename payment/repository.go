package payment

import "gorm.io/gorm"

type Repository interface {
	Save(input Payment, tx *gorm.DB) (Payment, error)
	Update(input Payment, tx *gorm.DB) (Payment, error)
	Delete(paymentId, tx *gorm.DB) (bool, error)
	FindAllByUserId(userId int, tx *gorm.DB) ([]Payment, error)
	FindById(paymentId int, tx *gorm.DB) (Payment, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (s *repository) Save(input Payment, tx *gorm.DB) (Payment, error) {

	err := tx.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil

}

func (s *repository) Update(input Payment, tx *gorm.DB) (Payment, error) {

	err := tx.Save(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil

}

func (s *repository) Delete(paymentId, tx *gorm.DB) (bool, error) {

	err := tx.Where("id = ?", paymentId).Delete(&Payment{}).Error
	if err != nil {
		return false, err
	}

	return true, nil

}

func (s *repository) FindAllByUserId(userId int, tx *gorm.DB) ([]Payment, error) {

	var payments []Payment

	err := tx.Where("user_id = ?", userId).Find(&payments).Error
	if err != nil {
		return payments, err
	}

	return payments, nil

}

func (s *repository) FindById(paymentId int, tx *gorm.DB) (Payment, error) {

	var payment Payment

	err := tx.Where("id = ?", paymentId).Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil

}
