package models

type Media struct {
	ID       string `json:"id" gorm:"primaryKey"`
	FileName string `json:"file_name" gorm:"not null"`
	FilePath string `json:"file_path" gorm:" not null"`
	FileSize int
	FileType string `json:"file_type" gorm:"not null"`

	Message []Message `json:"message"`
}
