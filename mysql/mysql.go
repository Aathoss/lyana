package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gitlab.com/lyana/logger"
)

var (
	member   sMember
	msgcount sMessageCount
	count    int
)

// Member stock les informations de la base de données
type sMember struct {
	id             int
	uid_discord    string
	player_mc      string
	date_whitelist int
}

type sMessageCount struct {
	id          int
	uid_discord string
	count_msg   int
}

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
	db, e := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	checkError(e)
	return db
}

func checkError(err error) {
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}

/*-------------------------------------------*/
/*-------------------------------------------*/
/*-------------- Section Select -------------*/
/*-------------------------------------------*/
/*-------------------------------------------*/

func SelectCount(tab, colonne, uid string) int {
	db := dbConn()
	defer db.Close()

	e := db.QueryRow("SELECT COUNT(*) FROM " + tab + " WHERE " + colonne + " = " + uid).Scan(&count)
	checkError(e)

	return count
}

func AddWhitelist(uid_discord, playermc string) {
	db := dbConn()
	defer db.Close()

	SelectCount("membre", "tag_discord", uid_discord)

	if count == 0 {
		tNow := time.Now()
		tUnix := tNow.Unix()

		insert, e := db.Prepare("INSERT INTO membre(tag_discord, player_mc, date_whitelist) VALUES(?,?,?)")
		checkError(e)
		insert.Exec(uid_discord, playermc, tUnix)
	}
}

func VerifPlayerMC(player string) int {
	db := dbConn()
	defer db.Close()

	e := db.QueryRow("SELECT COUNT(*) FROM membre WHERE player_mc = '" + player + "'").Scan(&count)
	checkError(e)

	return count
}

func NewCountMessage(author string) {
	db := dbConn()
	defer db.Close()

	SelectCount("message_count", "uid", author)

	if count == 1 {

		insert, e := db.Prepare("UPDATE message_count SET count_msg=count_msg+1 WHERE uid=?")
		checkError(e)
		insert.Exec(author)
		return
	}
	if count == 0 {
		insert, e := db.Prepare("INSERT INTO message_count(uid, count_msg) VALUES(?, ?)")
		checkError(e)
		insert.Exec(author, 1)
	}
}
