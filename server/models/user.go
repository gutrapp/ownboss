package models

type User struct {
	ID       uint     `gorm:"primaryKey:autoIncrement" json:"id"`
	Email    string   `gorm:"unique" json:"email"`
	Password string   `json:"password"`
	Employee Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func RegisterUser(user *User) error {
	tx := db.Create(&user)
	return tx.Error
}

func CheckCredentials(email string) (User, bool) {
	var data User
	tx := db.Where("email = ?", email).First(&data)

	if tx.Error != nil {
		return User{}, false
	}

	return data, true
}

func GetUser(id string) (User, error) {
	var user User

	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func GetUsers() ([]User, error) {
	var users []User

	tx := db.Find(&users)
	if tx.Error != nil {
		return []User{}, tx.Error
	}

	return users, nil
}

func GetJobs(id uint) ([]Employee, error) {
	var jobs []Employee

	tx := db.Where("user_id = ?", id).Find(&jobs)
	if tx.Error != nil {
		return []Employee{}, tx.Error
	}

	return jobs, nil
}
