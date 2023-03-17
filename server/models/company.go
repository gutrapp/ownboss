package models

type Company struct {
	ID          uint       `gorm:"primaryKey:autoIncrement" json:"id"`
	Services    []Service  `gorm:"OnDelete:CASCADE,ForeignKey:CompanyID"`
	Employees   []Employee `gorm:"OnDelete:CASCADE,ForeignKey:CompanyID"`
	Events      []Event    `gorm:"OnDelete:CASCADE,ForeignKey:CompanyID"`
	Name        string     `gorm:"unique" json:"name"`
	Description string     `json:"description"`
	Balance     uint       `json:"balance"`
}

func CreateCompany(company Company) (Company, error) {
	tx := db.Create(&company)
	return company, tx.Error
}

func UpdateCompany(company *Company) error {
	tx := db.Save(&company)
	return tx.Error
}

func DeleteCompany(id uint) error {
	tx := db.Unscoped().Delete(&Company{}, id)
	return tx.Error
}

func GetCompanies() ([]Company, error) {
	var companies []Company

	tx := db.Find(&companies)
	if tx.Error != nil {
		return []Company{}, tx.Error
	}

	return companies, nil
}

func GetCompany(id uint) (Company, error) {
	var companie Company

	tx := db.Where("id = ?", id).First(&companie)
	if tx.Error != nil {
		return Company{}, tx.Error
	}

	return companie, nil
}
