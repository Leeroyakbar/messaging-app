package models

import "database/sql/driver"

type RoleMember string

const (
	ADMIN  RoleMember = "admin"
	MEMBER RoleMember = "member"
)

func (rm *RoleMember) Scan(value interface{}) error {
	*rm = RoleMember(value.([]byte))
	return nil
}

func (rm RoleMember) Value() (driver.Value, error) {
	return string(rm), nil
}

type GroupMember struct {
	ID        string     `gorm:"primaryKey"`
	Role      RoleMember `json:"role" gorm:"not null"`
	GroupID   string     `json:"group_id" gorm:"not null"`
	ContactID string     `json:"contact_id" gorm:"not null"`

	Group   Group   `json:"group" gorm:"foreignKey:GroupID"`
	Contact Contact `json:"contact" gorm:"foreignKey:ContactID"`
}
