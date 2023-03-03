package service

import (
	"context"
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/Abdullayev65/mini_educational_system/pkg/utill"
)

func (s *Service) UserByUsername(c context.Context, username string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.NewSelect().Model(user).
		Where("username = ?", username).Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UserInfoByID(c context.Context, userID int) (*ioPut.UserInfo, error) {
	user, err := s.UserByID(c, userID)
	if err != nil {
		return nil, err
	}
	p, err := s.PositionByUserID(c, userID)
	if err != nil {
		return nil, err
	}
	dep, err := s.DepByID(c, p.DepartmentID)
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
	info := ioPut.MapUserInfo(user, p, dep, lvl, subject)
	return info, nil
}

func (s *Service) UserByID(c context.Context, userID int) (*models.User, error) {
	user := &models.User{ID: userID}
	err := s.DB.NewSelect().Model(user).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) AddUserWithPosition(c context.Context, user *models.User, p *models.Position) error {
	_, err := s.DB.NewInsert().Model(user).Exec(c)
	p.UserID = user.ID
	_, err = s.DB.NewInsert().Model(p).Exec(c)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UserUpdate(c context.Context, userID int, ua *ioPut.UserAdd) error {
	tx, err := s.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}
	if utill.AnyNotNil(ua.FullName, ua.Username) {
		user, err := s.UserByID(c, userID)
		if err != nil {
			return err
		}
		if ua.FullName != nil {
			user.FullName = *ua.FullName
		}
		if ua.Username != nil {
			user.Username = *ua.Username
		}
		_, err = tx.NewUpdate().Model(user).WherePK().Exec(c)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if utill.AnyNotNil(ua.LevelID, ua.DepartmentID, ua.SubjectID, ua.Type) {
		p, err := s.PositionByUserID(c, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
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
		_, err = tx.NewUpdate().Model(p).WherePK().Exec(c)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
