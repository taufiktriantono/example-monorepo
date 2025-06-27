package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

var (
	Admin    UserType = "admin"
	Customer UserType = "customer"
)

type AuthType string

var (
	Password     AuthType = "password"
	Phone        AuthType = "phone"
	Google       AuthType = "google"
	MagicLink    AuthType = "magic_link"
	ExternalOnly AuthType = "external_only"
)

type UserStatus string

var (
	Active    UserStatus = "active"
	Suspended UserStatus = "suspended"
	Disabled  UserStatus = "disabled"
)

type User struct {
	ID             string         `gorm:"column:id"`
	UserType       UserType       `gorm:"column:user_type"`
	AuthType       AuthType       `gorm:"auth_type"`
	OrganizationID string         `gorm:"column:organization_id"`
	ExternalID     string         `gorm:"column:external_id"`
	Username       string         `gorm:"column:username"`
	Password       string         `gorm:"column:password"`
	Status         UserStatus     `gorm:"column:status"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

type Profile struct {
	ID            string     `gorm:"column:id"`
	UserID        string     `gorm:"column:user_id;foreignKey:UserID;references:ID"`
	AvatarURL     string     `gorm:"column:avatar_url"`
	FirstName     string     `gorm:"column:first_name"`
	LastName      string     `gorm:"column:last_name"`
	Email         string     `gorm:"column:email"`
	Birthdate     string     `gorm:"column:birth_date;type:date"`
	EmailVerified *time.Time `gorm:"column:email_verified"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
}
