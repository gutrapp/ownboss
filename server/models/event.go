package models

type Event struct {
	ID        uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	CompanyID uint   `json:"companyId"`
	Date      string `json:"date"`
}

func CreateEvent(event *Event) error {
	tx := db.Create(event)
	return tx.Error
}

func UpdateEvent(event *Event) error {
	tx := db.Save(&event)
	return tx.Error
}

func DeleteEvent(id uint) error {
	tx := db.Unscoped().Delete(&Event{}, id)
	return tx.Error
}

func GetCompanyEvents(id uint) ([]Event, error) {
	var events []Event

	tx := db.Where("company_id = ?", id).Find(&events)
	if tx.Error != nil {
		return []Event{}, tx.Error
	}

	return events, nil
}

func GetEvent(id uint) (Event, error) {
	var event Event

	tx := db.Where("id = ?", id).First(&event)
	if tx.Error != nil {
		return Event{}, tx.Error
	}

	return event, nil
}

func GetEvents() ([]Event, error) {
	var events []Event

	tx := db.Find(&events)
	if tx.Error != nil {
		return []Event{}, tx.Error
	}

	return events, nil
}
