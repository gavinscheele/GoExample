package repository

type Database interface {
	NewDatabase(username, password, dbName string)
}
