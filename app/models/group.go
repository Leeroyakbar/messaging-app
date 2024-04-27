package models

type Group struct {
	ID        string `json:"id" gorm:"primaryKey"`
	GroupName string `json:"group_name" gorm:"not null"`

	GroupMember []GroupMember `json:"group_member"`
	Message     []Message     `json:"message"`
}
