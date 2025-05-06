package aijobs

import "gorm.io/gorm"

type Type string

const (
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	DB           *gorm.DB
	DataBaseType Type
}

type JobConfig struct {
	DB *gorm.DB
}
