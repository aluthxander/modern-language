package simple

type Database struct {
	Name string
}

type  DatabasePostgreSQL Database
type  DatabaseMongoDB Database


func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&DatabaseMongoDB{Name: "MongoDB"})
}

func NewDatabasePostgreSql() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&DatabasePostgreSQL{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSql *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSql: postgreSQL,
		DatabaseMongoDB:    mongoDB,
	}
}
