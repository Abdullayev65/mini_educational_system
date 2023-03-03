package service

import (
	"context"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

func (s *Service) PositionByID(c context.Context, id int) (*models.Position, error) {
	p := &models.Position{ID: id}
	err := s.DB.NewSelect().Model(p).WherePK().Scan(c)
	return p, err
}
func (s *Service) PositionByUserID(c context.Context, userID int) (*models.Position, error) {
	p := new(models.Position)
	err := s.DB.NewSelect().Model(p).Where("user_id = ?", userID).
		Limit(1).Scan(c)
	return p, err
}
func (s *Service) PositionAll(c context.Context) *[]models.Position {
	positions := new([]models.Position)
	s.DB.NewSelect().Model(positions).Scan(c)
	return positions
}
