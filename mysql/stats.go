package mysql

import (
	"fmt"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

func SelectCount(tab, colonne, uid string) int {
	err := framework.DBLyana.QueryRow("SELECT COUNT(*) FROM " + tab + " WHERE " + colonne + " = " + uid).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return count
}

//Count le message sur le discord
func NewCountMessage(author string) {
	SelectCount("message_count", "uid", author)

	if count == 1 {

		insert, err := framework.DBLyana.Prepare("UPDATE message_count SET count_msg=count_msg+1 WHERE uid=?")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		fmt.Println("[Mysql] [Débug] [Ligne 35 stats.go] uuid : " + author)

		insert.Exec(author)
		return
	}
	if count == 0 {
		insert, err := framework.DBLyana.Prepare("INSERT INTO message_count(uid, count_msg) VALUES(?, ?)")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		insert.Exec(author, 1)
	}
}

func ReturnNumMessages(uuid string) (int, error) {
	err := framework.DBLyana.QueryRow("SELECT count_msg FROM message_count WHERE uid=?", uuid).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return 0, err
	}

	return count, nil
}
