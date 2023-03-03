package ioPut

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

type Sign struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserAdd struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	FullName     *string `bun:",nullzero" json:"fullName"`
	LevelID      *int    `json:"levelID"`
	DepartmentID *int    `json:"departmentID"`
	SubjectID    *int    `json:"subjectID"`
	Type         *int    `json:"type"`
	TypeDay      *int    `json:"typeDay"`
}
type DepInput struct {
	Name *string `json:"name"`
}
type LvlInput struct {
	Type *int `json:"type"`
}
type SubjectInput struct {
	Name *string `json:"name"`
	Type *int    `json:"type"`
}
type AttendanceInput struct {
	Type *int `json:"type"`
}

func (ua *UserAdd) MapToUser() (*models.User, *models.Position) {
	user := models.User{}
	if ua.Username != nil {
		user.Username = *ua.Username
	}
	if ua.Password != nil {
		user.Password = *ua.Password
	}
	if ua.FullName != nil {
		user.FullName = *ua.FullName
	}
	p := models.Position{}
	if ua.LevelID != nil {
		p.LevelID = *ua.LevelID
	}
	if ua.DepartmentID != nil {
		p.DepartmentID = *ua.DepartmentID
	}
	if ua.SubjectID != nil {
		p.SubjectID = *ua.SubjectID
	}
	if ua.Type != nil {
		p.Type = *ua.Type
	}
	if ua.TypeDay != nil {
		p.TypeDay = *ua.TypeDay
	}
	return &user, &p
}

func (ua *UserAdd) ValidForAdding() (valid bool, fieldName string) {
	if ua.Username == nil || len(*ua.Username) < 3 {
		fieldName = "username"
		return
	}
	if ua.Password == nil || len(*ua.Password) < 3 {
		fieldName = "password"
		return
	}
	if ua.Type == nil || (*ua.Type < 1 || *ua.Type > 3) {
		fieldName = "type"
		return
	}
	return true, ""
}
