package service

import (
	"context"
	"errors"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

func (s *Service) AttendanceAdd(c context.Context, attendance *models.Attendance) error {
	if attendance.Type != 1 && attendance.Type != 2 {
		return errors.New("attendance type can be only 1 or 2")
	}
	err := s.DB.NewInsert().Model(attendance).Scan(c)
	return err
}
func (s *Service) AttendanceAll(c context.Context) *[]models.Attendance {
	all := new([]models.Attendance)
	s.DB.NewSelect().Model(all).Scan(c)
	return all
}
func (s *Service) AttendanceAllByUserID(c context.Context, userID int) (*[]models.Attendance, error) {
	all := new([]models.Attendance)
	err := s.DB.NewSelect().Model(all).
		Where("user_id = ?", userID).Scan(c)
	if err != nil {
		return nil, err
	}
	return all, nil
}
