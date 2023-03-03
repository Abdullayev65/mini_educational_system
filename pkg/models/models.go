package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Username      string `bun:",unique,nullzero,notnull" json:"username"`
	Password      string `bun:",nullzero,notnull" json:"password"`
	FullName      string `bun:",nullzero" json:"fullName"`
}
type Attendance struct {
	bun.BaseModel `bun:"table:attendance"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Type          int       `bun:",notnull" json:"type"`
	Time          time.Time `bun:",nullzero,notnull" json:"time"`
	UserID        int       `bun:",nullzero,notnull" json:"userID"`
}
type Department struct {
	bun.BaseModel `bun:"table:department"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",nullzero,notnull" json:"name"`
}
type Level struct {
	bun.BaseModel `bun:"table:level"`
	ID            int `bun:",pk,autoincrement" json:"id"`
	Type          int `bun:",nullzero,notnull" json:"type"`
}
type Subject struct {
	bun.BaseModel `bun:"table:subject"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",nullzero,notnull" json:"name"`
	Type          int    `bun:",nullzero,notnull" json:"type"`
}
type Position struct {
	bun.BaseModel `bun:"table:position"`
	ID            int `bun:",pk,autoincrement" json:"id"`
	Type          int `bun:",nullzero,notnull" json:"type"`
	TypeDay       int `bun:",nullzero,notnull" json:"TypeDay"`
	UserID        int `bun:",nullzero,notnull"`
	DepartmentID  int `bun:",nullzero,notnull"`
	LevelID       int `bun:",nullzero,notnull"`
	SubjectID     int `bun:",nullzero,notnull"`
}
