package mysql

import (
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

func Countuser(uuiddiscord string) (usercount int, channelcount string) {
	framework.DBLyana.QueryRow("SELECT COUNT(uuid) FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&usercount)
	framework.DBLyana.QueryRow("SELECT channelid FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&channelcount)

	return usercount, channelcount
}

func ReturnConfigChannel(uuiddiscord string) (channelname string, channeluserlimit int) {
	framework.DBLyana.QueryRow("SELECT channelname, channeluserlimit FROM tempvoc WHERE uuid = ?", uuiddiscord).Scan(&channelname, &channeluserlimit)
	return channelname, channeluserlimit
}

func InsertCreation(uuid, channelname string, channeluserlimit int) error {
	insert, err := framework.DBLyana.Prepare("INSERT INTO tempvoc(uuid, channelname, channeluserlimit) VALUES(?, ?, ?)")
	_, err = insert.Exec(uuid, channelname, channeluserlimit)
	insert.Close()
	return err
}

func UpdateChannelID(uuiddiscord, channelid string) error {
	update, err := framework.DBLyana.Query("UPDATE tempvoc SET channelid = ? WHERE uuid = ?", channelid, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	update.Close()
	return err
}

func UpdateChannelName(uuiddiscord, channelname string) error {
	update, err := framework.DBLyana.Query("UPDATE tempvoc SET channelname = ? WHERE uuid = ?", channelname, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	update.Close()
	return err
}

func UpdateChannelUserLimit(uuiddiscord string, channeluserlimit int) error {
	update, err := framework.DBLyana.Query("UPDATE tempvoc SET channeluserlimit = ? WHERE uuid = ?", channeluserlimit, uuiddiscord)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	update.Close()
	return err
}

func RemoveChannelID(channelid string) error {
	update, err := framework.DBLyana.Query("UPDATE tempvoc SET channelid = ? WHERE channelid = ?", "", channelid)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	update.Close()
	return err
}

func ReturnChannelIDAll() []string {
	info := []string{}
	i := 0

	var channelid string

	rows, err := framework.DBLyana.Query("SELECT channelid FROM tempvoc")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer rows.Close()

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
