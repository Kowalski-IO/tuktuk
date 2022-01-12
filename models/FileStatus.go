package models

type FileStatus string

const (
	Active    FileStatus = "ACTIVE"
	Destroyed            = "DESTROYED"
)
