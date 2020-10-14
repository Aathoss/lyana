package mysql

import (
	"time"

	"gitlab.com/lyana/logger"
)

//VerifPlayerMC Permet de voir si la personn est déjà whitelist
func VerifPlayerMC(player string) int {
	db := dbConn()
	defer db.Close()

	err := db.QueryRow("SELECT COUNT(*) FROM membre WHERE player_mc = '" + player + "'").Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return count
}

//AddWhitelist ajoute une personne à la whitelist
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

//GetWhitelist retourne les informations de la whitelist d'un utilisateur données
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

func DeleteUserWhitelist(uuid string) {
	db := dbConn()
	defer db.Close()

	insert, err := db.Prepare("DELETE FROM membre WHERE tag_discord=?")
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

/* func StatsMC(pseudoMC string) error {
	db := dbConnMC()
	defer db.Close()

	rows, err := db.Query(`SELECT VARIABLE,CONTENT FROM PLAYERDATA WHERE PLAYER=?`, pseudoMC)
	if err != nil {
		return err
	}

	var breakitems int
	var craftitems int
	var deaths int
	var jump int
	var mineblocks int
	var mineancientdebris int
	var minediamondore int
	var mobkills int
	var playertime int
	var useitems int

	for rows.Next() {
		var variable string
		var content interface{}

		err := rows.Scan(&variable, &content)
		if err != nil {
			return err
		}

		fmt.Println(variable)
		if variable == "breakitems" {
			rows.Scan(variable, breakitems)
			fmt.Print(breakitems)
			continue
		}
		if variable == "craftitems" {
			rows.Scan(variable, craftitems)
			fmt.Print(craftitems)
			continue
		}
		if variable == "deaths" {
			rows.Scan(variable, deaths)
			fmt.Print(deaths)
			continue
		}
		if variable == "jump" {
			rows.Scan(variable, jump)
			fmt.Print(jump)
			continue
		}
		if variable == "mineblocks" {
			rows.Scan(variable, mineblocks)
			fmt.Print(mineblocks)
			continue
		}
		if variable == "mine_ancient_debris" {
			rows.Scan(variable, mineancientdebris)
			fmt.Print(mineancientdebris)
			continue
		}
		if variable == "mine_diamond_ore" {
			rows.Scan(variable, minediamondore)
			fmt.Print(minediamondore)
			continue
		}
		if variable == "mobkills" {
			rows.Scan(variable, mobkills)
			fmt.Print(mobkills)
			continue
		}
		if variable == "playertime" {
			rows.Scan(variable, playertime)
			fmt.Print(playertime)
			continue
		}
		if variable == "useitems" {
			rows.Scan(variable, useitems)
			fmt.Print(useitems)
			continue
		}

		err = rows.Err()
		if err != nil {
			return err
		}
	}
	return nil
} */
