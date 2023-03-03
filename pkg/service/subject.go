package service

import (
	"context"
	"errors"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) SubjectByID(c context.Context, id int) (*models.Subject, error) {
	subject := &models.Subject{ID: id}
	err := s.DB.NewSelect().Model(subject).WherePK().Scan(c)
	return subject, err
}
func (s *Service) SubjectAll(c *gin.Context) *[]models.Subject {
	subjects := new([]models.Subject)
	s.DB.NewSelect().Model(subjects).Scan(c)
	return subjects
}
func (s *Service) SubjectAdd(c *gin.Context, subject *models.Subject) error {
	if subject.Name == "" {
		return errors.New("name of subject can not be null or blank")
	}
	if subject.Type != 1 && subject.Type != 2 {
		return errors.New("type of subject can be only 1 or 2")
	}
	_, err := s.DB.NewInsert().Model(subject).Exec(c)
	return err
}
func (s *Service) SubjectDelete(c *gin.Context, id int) error {
	subject := (*models.Subject)(nil)
	_, err := s.DB.NewDelete().Model(subject).
		Where("id = ?", id).Exec(c)
	return err
}
func (s *Service) SubjectUpdate(c *gin.Context, subject *models.Subject) error {
	_, err := s.DB.NewUpdate().Model(subject).WherePK().Exec(c)
	return err
}
