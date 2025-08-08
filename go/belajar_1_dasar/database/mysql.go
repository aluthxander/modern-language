package database

var connection string

func init(){
	connection = "Mysql Connection"
}

func GetDatabase() string {
	return connection
}