package mysql

import (
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

//VerifPlayerMC Permet de voir si la personn est déjà whitelist
func VerifPlayerMC(uiddiscord, player string) (countuuid, countplayer int) {

	framework.DBLyana.QueryRow("SELECT COUNT(*) FROM membre WHERE tag_discord = '" + uiddiscord + "'").Scan(&countuuid)
	framework.DBLyana.QueryRow("SELECT COUNT(*) FROM membre WHERE player_mc = '" + player + "'").Scan(&countplayer)

	return countuuid, countplayer
}

//AddWhitelist ajoute une personne à la whitelist
func AddWhitelist(uiddiscord, playermc string) error {

	if count == 0 {
		t1 := time.Now()
		t2 := t1.Unix()

		insert, err := framework.DBLyana.Prepare("INSERT INTO membre(tag_discord, player_mc, date_whitelist, inactif, notif) VALUES(?,?,?,?,?)")
		if err != nil {
			logger.ErrorLogger.Println(err)
			return err
		}
		_, err = insert.Exec(uiddiscord, playermc, t2, t2, 0)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return err
		}
	}
	return nil
}

//GetWhitelist retourne les informations de la whitelist d'un utilisateur données
func GetWhitelist(uuid string) (string, string, int64, error) {

	err := framework.DBLyana.QueryRow("SELECT * FROM membre WHERE tag_discord = "+uuid).Scan(&member.id, &member.uid_discord, &member.player_mc, &member.date_whitelist, &member.inactif, &member.notif)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", "", 0, err
	}

	return member.uid_discord, member.player_mc, member.date_whitelist, nil
}

func DeleteUserWhitelist(uuid string) {

	insert, err := framework.DBLyana.Prepare("DELETE FROM membre WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	_, err = insert.Exec(uuid)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
}
