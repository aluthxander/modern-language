package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRolback := tx.Rollback()
		PanicError(errorRolback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}