package mysql

import (
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

/*-------------------------------------------*/
/*-------------------------------------------*/
/*-------------- Section Select -------------*/
/*-------------------------------------------*/
/*-------------------------------------------*/

func AddSanctionLimit(uid_discord, pseudomc, id_message, id_message_notif string) {
	insert, err := framework.DBLyana.Prepare("INSERT INTO sanction(uid, pseudomc, id_message, id_msg_notif) VALUES(?, ?, ?, ?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	insert.Exec(uid_discord, pseudomc, id_message, id_message_notif)
	insert.Close()
}

func RemoveSanctionLimit(uid_discord string) (pseudomc, sanctionID_msg, sanctionID_msg_notif string) {
	err := framework.DBLyana.QueryRow("SELECT * FROM sanction WHERE uid = "+uid_discord).Scan(&sanction.id, &sanction.uid, &sanction.pseudomc, &sanction.id_message, &sanction.id_msg_notif)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	delete, err := framework.DBLyana.Query("DELETE FROM sanction WHERE uid = " + uid_discord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	delete.Close()

	return sanction.pseudomc, sanction.id_message, sanction.id_msg_notif
}
