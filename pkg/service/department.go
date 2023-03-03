package service

import (
	"context"
	"errors"
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
)

func (s *Service) DepAdd(c context.Context, dep *models.Department) error {
	if dep.Name == "" {
		return errors.New("name of department can not be null or blank")
	}
	_, err := s.DB.NewInsert().Model(dep).Exec(c)
	return err
}
func (s *Service) DepUpdate(c context.Context, dep *models.Department) error {
	if dep.Name == "" {
		return errors.New("name of department can not be null or blank")
	}
	_, err := s.DB.NewUpdate().Model(dep).
		WherePK().Exec(c)
	return err
}
func (s *Service) DepAll(c context.Context) *[]models.Department {
	deps := new([]models.Department)
	s.DB.NewSelect().Model(deps).Scan(c)
	return deps
}
func (s *Service) DepByID(c context.Context, id int) (*models.Department, error) {
	dep := &models.Department{ID: id}
	err := s.DB.NewSelect().Model(dep).WherePK().Scan(c)
	return dep, err
}
func (s *Service) DepDelete(c context.Context, id int) error {
	d := (*models.Department)(nil)
	_, err := s.DB.NewDelete().Model(d).
		Where("id = ?", id).Exec(c)
	return err
}
func (s *Service) UsersByDepID(c context.Context, depID int) (*ioPut.DepInfo, error) {
	dep := &models.Department{ID: depID}
	err := s.DB.NewSelect().Model(dep).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	depInfo := &ioPut.DepInfo{ID: &dep.ID, Name: &dep.Name}
	positions := make([]models.Position, 0)
	err = s.DB.NewSelect().Model(&positions).
		Where("department_id = ?", depID).Scan(c)
	if err != nil {
		return nil, err
	}
	for _, p := range positions {
		user, err := s.UserByID(c, p.UserID)
		if err != nil {
			return nil, err
		}
		lvl, err := s.LevelByID(c, p.LevelID)
		if err != nil {
			return nil, err
		}
		subject, err := s.SubjectByID(c, p.SubjectID)
		if err != nil {
			return nil, err
		}
		{
			depInfo.Admins = new([]ioPut.UserDep)
			depInfo.Students = new([]ioPut.UserDep)
			depInfo.Teachers = new([]ioPut.UserDep)
		}
		var temp *[]ioPut.UserDep
		switch p.Type {
		case 1:
			temp = depInfo.Admins
		case 2:
			temp = depInfo.Students
		case 3:
			temp = depInfo.Teachers
		}
		*temp = append(*temp, *ioPut.NewUserDep(user, &p, lvl, subject))
	}

	return depInfo, nil
}
