package mysql

import (
	"fmt"
	"strings"
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

func UpdateInactifPlayer() {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	t2 := t1.Unix()

	if len(framework.ListPlayer) == 0 {
		return
	}

	for _, player := range framework.ListPlayer {
		player := strings.Replace(player, ",", " ", -1)

		insert, err := db.Prepare("UPDATE membre SET inactif=? WHERE player_mc=?")
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		insert.Exec(t2, player)
	}
}

func UpdateInactifDiscord(uuid string) {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	t2 := t1.Unix()

	insert, err := db.Prepare("UPDATE membre SET inactif=? WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	insert.Exec(t2, uuid)
}

func VerifInactif() error {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	t2 := t1.Unix()

	fmt.Println(t2)
	fmt.Println(t2 - 20000)
	//fmt.Println(t2 - 1209600)

	rows, err := db.Query("SELECT id, player_mc, inactif, notif FROM membre")
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(&member.id, &member.player_mc, &member.inactif, &member.notif)
		if err != nil {
			return err
		}

		if member.inactif <= t2-20000 {
			fmt.Print(member.id)
			fmt.Print(" ")
			fmt.Print(member.player_mc)
			fmt.Print(" ")
			fmt.Print(member.inactif)
			fmt.Print(" ")
			fmt.Println(member.notif)
		}

		err = rows.Err()
		if err != nil {
			return err
		}
	}
	return nil
}
