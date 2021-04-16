package mysql

import (
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
)

func VerifRule(uuid string) (int, error) {
	err := framework.DBLyana.QueryRow("SELECT COUNT(*) FROM rule WHERE uid=" + uuid).Scan(&count)
	return count, err
}

func VerifRuleTimestamp() ([]string, error) {
	t1 := time.Now()
	var temp []string

	rows, err := framework.DBLyana.Query("SELECT * FROM rule")
	if err != nil {
		return temp, err
	}
	defer rows.Close()

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
	insert, err := framework.DBLyana.Prepare("INSERT INTO rule(uid, timestamp) VALUES(?,?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	insert.Exec(uuid, timestamp+(60*60*24*3))
	insert.Close()
}

func RemoveRule(uuid string) error {
	delete, err := framework.DBLyana.Query("DELETE FROM rule WHERE uid = " + uuid)
	if err != nil {
		return err
	}
	delete.Close()
	return nil
}
