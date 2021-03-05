package mysql

import (
	"gitlab.com/lyana/logger"
)

func Countuser(uuiddiscord string) (usercount int, channelcount string) {
	db := DbConn()
	defer db.Close()

	db.QueryRow("SELECT COUNT(uuid) FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&usercount)
	db.QueryRow("SELECT channelid FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&channelcount)

	return usercount, channelcount
}

func ReturnConfigChannel(uuiddiscord string) (channelname string, channeluserlimit int) {
	db := DbConn()
	defer db.Close()

	db.QueryRow("SELECT channelname, channeluserlimit FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&channelname, &channeluserlimit)
	return channelname, channeluserlimit
}

func InsertCreation(uuid, channelname string, channeluserlimit int) error {
	db := DbConn()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO tempvoc(uuid, channelname, channeluserlimit) VALUES(?, ?, ?)")
	_, err = insert.Exec(uuid, channelname, channeluserlimit)
	return err
}

func UpdateChannelID(uuiddiscord, channelid string) error {
	db := DbConn()
	defer db.Close()

	_, err := db.Query("UPDATE tempvoc SET channelid = ? WHERE uuid = ?", channelid, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}

func UpdateChannelName(uuiddiscord, channelname string) error {
	db := DbConn()
	defer db.Close()

	_, err := db.Query("UPDATE tempvoc SET channelname = ? WHERE uuid = ?", channelname, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}

func UpdateChannelUserLimit(uuiddiscord string, channeluserlimit int) error {
	db := DbConn()
	defer db.Close()

	_, err := db.Query("UPDATE tempvoc SET channeluserlimit = ? WHERE uuid = ?", channeluserlimit, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}

func RemoveChannelID(channelid string) error {
	db := DbConn()
	defer db.Close()

	_, err := db.Query("UPDATE tempvoc SET channelid = ? WHERE channelid = ?", "", channelid)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func ReturnChannelIDAll() []string {
	db := DbConn()
	defer db.Close()
	info := []string{}
	i := 0

	var channelid string

	rows, err := db.Query("SELECT channelid FROM tempvoc")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	for rows.Next() {

		err := rows.Scan(&channelid)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		info = append(info, channelid)

		err = rows.Err()
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		i++
	}
	return info
}
