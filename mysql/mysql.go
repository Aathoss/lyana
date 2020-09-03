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
	sanction sSanction
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

type sSanction struct {
	id           int
	uid          string
	id_message   string
	id_msg_notif string
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

//Count le message sur le discord
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

func AddSanctionLimit(uid_discord, id_message, id_message_notif string) {
	db := dbConn()
	defer db.Close()

	insert, e := db.Prepare("INSERT INTO sanction(uid, id_message, id_msg_notif) VALUES(?, ?, ?)")
	checkError(e)
	insert.Exec(uid_discord, id_message, id_message_notif)
}

func RemoveSanctionLimit(uid_discord string) (sanctionID_msg, sanctionID_msg_notif string) {
	db := dbConn()
	defer db.Close()

	e := db.QueryRow("SELECT * FROM sanction WHERE uid = "+uid_discord).Scan(&sanction.id, &sanction.uid, &sanction.id_message, &sanction.id_msg_notif)
	checkError(e)

	_, e = db.Query("DELETE FROM sanction WHERE uid = " + uid_discord)
	checkError(e)

	return sanction.id_message, sanction.id_msg_notif
}

/* rows, e := db.Query("SELECT * FROM ticket WHERE id_member=" + strconv.Itoa(int(member.id)) + " ORDER BY num DESC LIMIT 10")
checkError(e)

for rows.Next() {
	var info []string

	e := rows.Scan(&ticket.id, &ticket.num, &ticket.timesec, &ticket.status, &ticket.idmember, &ticket.request)
	checkError(e)

	info = append(info, strconv.Itoa(int(i+1)), strconv.Itoa(int(ticket.num)), modules.Calcseconde(ticket.timesec), ticket.request, strconv.FormatBool(ticket.status))
	tab = append(tab, info)

	e = rows.Err()
	checkError(e)
	i++
} */
