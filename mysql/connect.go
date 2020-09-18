package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gitlab.com/lyana/logger"
)

/*-------------------------------------------*/
/*-------------------------------------------*/
/*---------- Connexion bdd / erreur ---------*/
/*-------------------------------------------*/
/*-------------------------------------------*/

// Gestion de la connexion à la base de données
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := viper.GetString("MySql.dbuser")
	dbPass := viper.GetString("MySql.dbmdp")
	dbName := viper.GetString("MySql.dbname")
	dbIP := viper.GetString("MySql.dbip")
	dbPort := viper.GetString("MySql.dbport")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return db
}
