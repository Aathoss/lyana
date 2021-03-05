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
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := viper.GetString("MySql.Lyana.dbuser")
	dbPass := viper.GetString("MySql.Lyana.dbmdp")
	dbName := viper.GetString("MySql.Lyana.dbname")
	dbIP := viper.GetString("MySql.Lyana.dbip")
	dbPort := viper.GetString("MySql.Lyana.dbport")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return db
}

func DbConnMC() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := viper.GetString("MySql.Minecraft.dbuser")
	dbPass := viper.GetString("MySql.Minecraft.dbmdp")
	dbName := viper.GetString("MySql.Minecraft.dbname")
	dbIP := viper.GetString("MySql.Minecraft.dbip")
	dbPort := viper.GetString("MySql.Minecraft.dbport")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return db
}
