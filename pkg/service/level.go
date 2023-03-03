package service

import (
	"context"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

func (s *Service) LevelByID(c context.Context, id int) (*models.Level, error) {
	lvl := &models.Level{ID: id}
	err := s.DB.NewSelect().Model(lvl).WherePK().Scan(c)
	return lvl, err
}
func (s *Service) LvlAdd(c context.Context, lvl *models.Level) error {
	_, err := s.DB.NewInsert().Model(lvl).Exec(c)
	return err
}
func (s *Service) LvlUpdate(c context.Context, lvl *models.Level) error {
	_, err := s.DB.NewUpdate().Model(lvl).
		WherePK().Exec(c)
	return err
}
func (s *Service) LvlAll(c context.Context) *[]models.Level {
	lvls := new([]models.Level)
	s.DB.NewSelect().Model(lvls).Scan(c)
	return lvls
}
func (s *Service) LvlDelete(c context.Context, id int) error {
	d := (*models.Level)(nil)
	_, err := s.DB.NewDelete().Model(d).
		Where("id = ?", id).Exec(c)
	return err
}
