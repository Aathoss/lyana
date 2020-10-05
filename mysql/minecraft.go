package mysql

import (
	"time"

	"gitlab.com/lyana/logger"
)

func VerifPlayerMC(player string) int {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT COUNT(*) FROM membre WHERE player_mc = '" + player + "'").Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return count
}

func AddWhitelist(uid_discord, playermc string) error {
	db := dbConn()
	defer db.Close()

	SelectCount("membre", "tag_discord", uid_discord)

	if count == 0 {
		t1 := time.Now()
		t2 := t1.Unix()

		insert, err := db.Prepare("INSERT INTO membre(tag_discord, player_mc, date_whitelist, inactif) VALUES(?,?,?,?)")
		if err != nil {
			logger.ErrorLogger.Println(err)
			return err
		}
		_, err = insert.Exec(uid_discord, playermc, t2, t2)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return err
		}
	}
	return nil
}

func GetWhitelist(uuid string) (string, string, int64, error) {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM membre WHERE tag_discord = "+uuid).Scan(&member.id, &member.uid_discord, &member.player_mc, &member.date_whitelist, &member.inactif, &member.notif)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", "", 0, err
	}

	return member.uid_discord, member.player_mc, member.date_whitelist, nil
}
