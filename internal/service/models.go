package service

type Server struct {
	DBase DB
}

type DB interface {
	CreateTable()
	Close()
}

const (
	MaxMoveDays = 400
)
