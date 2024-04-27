package models

type Message struct {
	ID         string `json:"id" gorm:"primaryKey"`
	FromNumber string `json:"from_number" gorm:"not null"`
	ToNumber   string `json:"to_number" gorm:"not null"`
	Status     string `json:"status" gorm:"type:varchar(30)"`
	MediaID    string `json:"media_id" gorm:"not null"`
	ContactID  string `json:"contact_id" gorm:"not null"`
	GroupID    string `json:"group_id" gorm:"not null"`

	Media   Media   `json:"media" gorm:"foreignKey:MediaID"`
	Contact Contact `json:"contact" gorm:"foreignKey:ContactID"`
	Group   Group   `json:"group" gorm:"foreignKey:GroupID"`
}
