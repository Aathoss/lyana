package mysql

import (
	"strconv"
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

		insert, err := db.Prepare("UPDATE membre SET inactif=?, notif=0 WHERE player_mc=?")
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

	insert, err := db.Prepare("UPDATE membre SET inactif=?, notif=0 WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	insert.Exec(t2, uuid)
}

func VerifInactif() ([][]string, error) {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	t2 := t1.Unix()
	tab := [][]string{}

	rows, err := db.Query("SELECT id, tag_discord, inactif, notif FROM membre")
	if err != nil {
		return tab, err
	}

	for rows.Next() {
		var info []string

		err := rows.Scan(&member.id, &member.uid_discord, &member.inactif, &member.notif)
		if err != nil {
			return tab, err
		}

		if member.inactif <= t2-604800 {
			info = append(info, strconv.Itoa(member.id), member.uid_discord, strconv.Itoa(int(member.inactif)), strconv.Itoa(member.notif))
			tab = append(tab, info)
		}

		err = rows.Err()
		if err != nil {
			return tab, err
		}
	}
	return tab, nil
}

func UpdateMembresInactif(uuid string) error {
	db := dbConn()
	defer db.Close()

	t1 := time.Now()
	t2 := t1.Unix()

	insert, err := db.Prepare("UPDATE membre SET inactif=?, notif=notif+1 WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	insert.Exec(t2, uuid)
	return nil
}