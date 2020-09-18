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

func AddWhitelist(uid_discord, playermc string) {
	db := dbConn()
	defer db.Close()

	SelectCount("membre", "tag_discord", uid_discord)

	if count == 0 {
		tNow := time.Now()

		insert, err := db.Prepare("INSERT INTO membre(tag_discord, player_mc, date_whitelist) VALUES(?,?,?)")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		insert.Exec(uid_discord, playermc, tNow.Unix)
	}
}

func GetWhitelist(uuid string) (string, string, int) {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM membre WHERE tag_discord = "+uuid).Scan(&member.id, &member.uid_discord, &member.player_mc, &member.date_whitelist)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return member.uid_discord, member.player_mc, member.date_whitelist
}
