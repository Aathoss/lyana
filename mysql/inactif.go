package mysql

import (
	"strconv"
	"strings"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
)

func UpdateInactifPlayer() {
	t1 := time.Now()
	t2 := t1.Unix()

	if len(framework.ListPlayer) == 0 {
		return
	}

	for _, player := range framework.ListPlayer {
		player := strings.Replace(player, ",", " ", -1)

		insert, err := framework.DBLyana.Prepare("UPDATE membre SET inactif=?, notif=0 WHERE player_mc=?")
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		insert.Exec(t2, player)
		insert.Close()
	}
	return
}

func UpdateInactifDiscord(uuid string) {
	t1 := time.Now()
	t2 := t1.Unix()

	insert, err := framework.DBLyana.Prepare("UPDATE membre SET inactif=?, notif=0 WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	insert.Exec(t2, uuid)
	insert.Close()
}

func VerifInactif() ([][]string, error) {
	t1 := time.Now()
	t2 := t1.Unix()
	tab := [][]string{}

	rows, err := framework.DBLyana.Query("SELECT id, tag_discord, inactif, notif FROM membre")
	if err != nil {
		return tab, err
	}
	defer rows.Close()

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
	t1 := time.Now()
	t2 := t1.Unix()

	insert, err := framework.DBLyana.Prepare("UPDATE membre SET inactif=?, notif=notif+1 WHERE tag_discord=?")
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	insert.Exec(t2, uuid)
	insert.Close()
	return nil
}

func CompteInactif() (int, error) {
	err := framework.DBLyana.QueryRow("SELECT COUNT(notif) FROM membre WHERE notif>0").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
