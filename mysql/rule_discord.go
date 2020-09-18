package mysql

import (
	"time"

	"gitlab.com/lyana/logger"
)

func VerifRule(uuid string) (int, error) {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT COUNT(*) FROM rule WHERE uid=" + uuid).Scan(&count)
	return count, err
}

func VerifRuleTimestamp() ([]string, error) {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	var temp []string

	rows, err := db.Query("SELECT * FROM rule")
	if err != nil {
		return temp, err
	}

	for rows.Next() {
		err := rows.Scan(&rule.id, &rule.uid, &rule.timestamp)
		if err != nil {
			return temp, err
		}

		err = rows.Err()
		if err != nil {
			return temp, err
		}

		if t1.Unix() >= rule.timestamp {
			temp = append(temp, rule.uid)
		}
	}

	return temp, nil
}

func AddRule(uuid string, timestamp int64) {
	db := dbConn()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO rule(uid, timestamp) VALUES(?,?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	insert.Exec(uuid, timestamp+(60*60*24*3))
}

func RemoveRule(uuid string) {
	db := dbConn()
	defer db.Close()

	_, err := db.Query("DELETE FROM rule WHERE uid = " + uuid)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
