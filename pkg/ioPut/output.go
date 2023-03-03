package ioPut

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

type AttendanceInfo struct {
	FullName    string
	Username    string
	Department  string
	Position    string
	Attendances []models.Attendance
}

type UserInfo struct {
	ID          *int    `json:"id"`
	Username    *string `json:"username"`
	FullName    *string `json:"fullName"`
	Type        *int    `json:"type"`
	TypeDay     *int    `json:"typeDay"`
	Department  *string `json:"department"`
	Level       *int    `json:"level"`
	SubjectName *string `json:"SubjectName"`
	SubjectType *int    `json:"subjectType"`
}
type UserDep struct {
	ID          *int    `json:"id"`
	Username    *string `json:"username"`
	FullName    *string `json:"fullName"`
	Type        *int    `json:"type"`
	TypeDay     *int    `json:"typeDay"`
	Level       *int    `json:"level"`
	SubjectName *string `json:"SubjectName"`
	SubjectType *int    `json:"subjectType"`
}

type DepInfo struct {
	ID       *int       `json:"id"`
	Name     *string    `json:"username"`
	Admins   *[]UserDep `json:"admins"`
	Students *[]UserDep `json:"students"`
	Teachers *[]UserDep `json:"teachers"`
}

func MapUserInfo(user *models.User, position *models.Position, dep *models.Department,
	lvl *models.Level, subject *models.Subject) *UserInfo {
	ui := new(UserInfo)
	ui.ID = &user.ID
	ui.Username = &user.Username
	ui.FullName = &user.FullName
	if position != nil {
		ui.Type = &position.Type
		ui.TypeDay = &position.TypeDay
	}
	if dep != nil {
		ui.Department = &dep.Name
	}
	if lvl != nil {
		ui.Level = &lvl.Type
	}
	if subject != nil {
		ui.SubjectType = &subject.Type
		ui.SubjectName = &subject.Name
	}

	return ui
}
func NewUserDep(user *models.User, p *models.Position, lvl *models.Level, subject *models.Subject) *UserDep {
	return &UserDep{
		ID:          &user.ID,
		Username:    &user.Username,
		FullName:    &user.FullName,
		Type:        &p.Type,
		TypeDay:     &p.TypeDay,
		Level:       &lvl.Type,
		SubjectName: &subject.Name,
		SubjectType: &subject.Type,
	}
}
