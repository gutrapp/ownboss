package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	CompanyID   uint   `json:"companyId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantatiy"`
}

func CreateService(service *Service) error {
	tx := db.Create(service)
	return tx.Error
}

func UpdateService(service *Service) error {
	tx := db.Save(&service)
	return tx.Error
}

func DeleteService(id uint) error {
	tx := db.Unscoped().Delete(&Service{}, id)
	return tx.Error
}

func GetCompanyServices(id uint) ([]Service, error) {
	var services []Service

	tx := db.Where("company_id = ?", id).Find(&services)
	if tx.Error != nil {
		return []Service{}, tx.Error
	}

	return services, nil
}

func GetService(id uint) (Service, error) {
	var service Service

	tx := db.Where("id = ?", id).First(&service)
	if tx.Error != nil {
		return Service{}, tx.Error
	}

	return service, nil
}

func GetServices() ([]Service, error) {
	var services []Service

	tx := db.Find(&services)
	if tx.Error != nil {
		return []Service{}, tx.Error
	}

	return services, nil
}
