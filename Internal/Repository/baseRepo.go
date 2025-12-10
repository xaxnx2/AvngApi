package Repository

import (
	database "HOMEWORKAVENGINVAPI/Internal/Database"
	"database/sql"

	_ "github.com/lib/pq"
)

type BaseRepository struct {
	db *sql.DB
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{
		db: database.ConnectDB(),
	}
}
