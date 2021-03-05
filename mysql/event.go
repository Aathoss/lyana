package mysql

import (
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/lyana/logger"
)

func CountIndexEvent() (count int, err error) {
	db := DbConn()
	defer db.Close()

	err = db.QueryRow("SELECT `AUTO_INCREMENT` FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 's13_lyana' AND TABLE_NAME = 'event'").Scan(&count)
	return count, err
}

func CreateEvent(status, titre, auteur, messageid, channelid, emplacement, eventdate, description, recompense, participant string) (err error) {
	db := DbConn()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO event(status, messageid, channelid, titre, auteur, localisation, description, date, recompense, participant) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	_, err = insert.Exec(status, messageid, channelid, titre, auteur, emplacement, description, eventdate, recompense, participant)
	return err
}

func GetCreationEvent(idtab int) (info []string, err error) {
	db := DbConn()
	defer db.Close()

	var (
		id           int
		status       string
		messageid    string
		channelid    string
		titre        string
		auteur       string
		localisation string
		description  string
		date         string
		recompense   string
		participant  string
	)

	info = []string{}

	err = db.QueryRow("SELECT * FROM event WHERE id = ? LIMIT 1", idtab).Scan(&id, &status, &messageid, &channelid, &titre, &auteur, &localisation, &description, &date, &recompense, &participant)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return info, err
	}

	info = append(info, strconv.Itoa(id), status, messageid, channelid, titre, localisation, description, date, recompense, participant, auteur)

	return info, nil
}

func GetMultiEvent() (content [][]string, err error) {
	db := DbConn()
	defer db.Close()

	var (
		id           int
		status       string
		messageid    string
		channelid    string
		titre        string
		auteur       string
		localisation string
		description  string
		date         string
		recompense   string
		participant  string
		tab          [][]string
	)

	rows, err := db.Query("SELECT * FROM event WHERE status != ? AND status != ?", "dev", "terminer")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	for rows.Next() {
		var info []string

		err := rows.Scan(&id, &status, &messageid, &channelid, &titre, &auteur, &localisation, &description, &date, &recompense, &participant)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return tab, err
		}

		info = append(info, strconv.Itoa(id), status, messageid, channelid, titre, localisation, description, date, recompense, participant, auteur)
		tab = append(tab, info)

		err = rows.Err()
		if err != nil {
			logger.ErrorLogger.Println(err)
			return tab, err
		}
	}

	return tab, nil
}

//0 = dev
//1 = à venir
//2 = en cours
//3 = prepterminer
//4 = terminer
func EditStatus(status, id int) (err error) {
	db := DbConn()
	defer db.Close()
	situation := "dev"

	if status == 1 {
		situation = "à venir"
	} else if status == 2 {
		situation = "en cours"
	} else if status == 3 {
		situation = "prepterminer"
	} else if status == 4 {
		situation = "terminer"
	}

	_, err = db.Query("UPDATE event SET status = ? WHERE id = ?", situation, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditTitre(titre string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET titre = ? WHERE id = ?", titre, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditAuteur(auteur string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET auteur = ? WHERE id = ?", auteur, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditEmplacement(emplacement string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET localisation = ? WHERE id = ?", emplacement, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditDescription(description string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET description = ? WHERE id = ?", description, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditDate(date string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET date = ? WHERE id = ?", date, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditRecompense(recompense string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET recompense = ? WHERE id = ?", recompense, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditChannelID(ChannelID string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET channelid = ? WHERE id = ?", ChannelID, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func EditMessageID(MessageID string, id int) (err error) {
	db := DbConn()
	defer db.Close()

	_, err = db.Query("UPDATE event SET messageid = ? WHERE id = ?", MessageID, id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

//0 = Add participant
//1 = Remove participant
func ReactionParticipants(situation int, messageid, uuid string) {
	db := DbConn()
	defer db.Close()

	listparticipants := ""
	supp := false

	err := db.QueryRow("SELECT participant FROM event WHERE messageid = ?", messageid).Scan(&listparticipants)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	// Add participant à l'évent
	if situation == 0 {
		if len(listparticipants) == 0 {
			listparticipants = uuid
		} else {
			listparticipants = listparticipants + "," + uuid
		}
	}

	// Remove participant de l'évent
	if situation == 1 {
		re := regexp.MustCompile("," + uuid)
		if re.MatchString(listparticipants) == true {
			listparticipants = strings.Replace(listparticipants, ","+uuid, "", -1)
			supp = true
		}

		re = regexp.MustCompile(uuid)
		if re.MatchString(listparticipants) == true && supp == false {
			listparticipants = strings.Replace(listparticipants, uuid, "", -1)
		}
	}

	_, err = db.Query("UPDATE event SET participant = ? WHERE messageid = ?", listparticipants, messageid)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return

}
