package mysql

import "gitlab.com/lyana/logger"

func SelectCount(tab, colonne, uid string) int {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT COUNT(*) FROM " + tab + " WHERE " + colonne + " = " + uid).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return count
}

//Count le message sur le discord
func NewCountMessage(author string) {
	db := dbConn()
	defer db.Close()

	SelectCount("message_count", "uid", author)

	if count == 1 {

		insert, err := db.Prepare("UPDATE message_count SET count_msg=count_msg+1 WHERE uid=?")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		insert.Exec(author)
		return
	}
	if count == 0 {
		insert, err := db.Prepare("INSERT INTO message_count(uid, count_msg) VALUES(?, ?)")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		insert.Exec(author, 1)
	}
}
