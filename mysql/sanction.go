package mysql

import "gitlab.com/lyana/logger"

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
