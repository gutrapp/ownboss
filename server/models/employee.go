package models

type Employee struct {
	ID        uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	CompanyID uint   `json:"companyId"`
	UserID    uint   `json:"userId"`
	Position  string `json:"title"`
	Name      string `json:"name"`
	Salary    uint   `json:"salary"`
	JobTitle  string `json:"jobTitle"`
}

func CreateEmployee(employee *Employee) error {
	tx := db.Create(&employee)
	return tx.Error
}

func UpdateEmployee(employee *Employee) error {
	tx := db.Save(&employee)
	return tx.Error
}

func DeleteEmployee(id uint) error {
	tx := db.Unscoped().Delete(&Employee{}, id)
	return tx.Error
}

func GetEmployee(id uint) (Employee, error) {
	var employee Employee

	tx := db.Where("id = ?", id).First(&employee)
	if tx.Error != nil {
		return Employee{}, tx.Error
	}

	return employee, nil
}

func GetEmployees() ([]Employee, error) {
	var employees []Employee

	tx := db.Find(&employees)
	if tx.Error != nil {
		return []Employee{}, tx.Error
	}

	return employees, nil
}

func GetEmployeesByCompanyId(id uint) ([]Employee, error) {
	var employees []Employee

	tx := db.Where("company_id = ?", id).Find(&employees)
	if tx.Error != nil {
		return []Employee{}, tx.Error
	}

	return employees, nil
}
