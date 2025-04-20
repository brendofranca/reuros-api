package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
