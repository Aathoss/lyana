package mysql

import (
	"gitlab.com/lyana/logger"
)

/*-------------------------------------------*/
/*-------------------------------------------*/
/*-------------- Section Select -------------*/
/*-------------------------------------------*/
/*-------------------------------------------*/

func AddSanctionLimit(uid_discord, id_message, id_message_notif string) {
	db := dbConn()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO sanction(uid, id_message, id_msg_notif) VALUES(?, ?, ?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	insert.Exec(uid_discord, id_message, id_message_notif)
}

func RemoveSanctionLimit(uid_discord string) (sanctionID_msg, sanctionID_msg_notif string) {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM sanction WHERE uid = "+uid_discord).Scan(&sanction.id, &sanction.uid, &sanction.id_message, &sanction.id_msg_notif)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	_, err = db.Query("DELETE FROM sanction WHERE uid = " + uid_discord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

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
