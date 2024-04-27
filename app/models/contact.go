package models

type Contact struct {
	ID          string `json:"id" gorm:"primaryKey"`
	ContactName string `json:"contact_name" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`

	GroupMember []GroupMember `json:"group_member"`
	Message     []Message     `json:"message"`
}
