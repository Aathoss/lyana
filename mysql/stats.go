package mysql

import (
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
