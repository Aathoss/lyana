package event

import (
	"strconv"
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

var (
	count60seconde = 60
)

func UpdateEvent(secondeboucle time.Duration) {
	logger.DebugLogger.Println("Starting UpdateEvent")

	session := framework.Session

	for {
		time.Sleep(time.Second * secondeboucle)

		//Update de l'embed lors de la création d'un nouvelle évent
		if framework.EventConstruction == true {
			tab, err := mysql.GetCreationEvent(framework.EventContructionIndex)
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}

			ConstructionEmbedEvent(0, session, tab)
		}

		//Update multi évent public
		if count60seconde >= 60 {
			tab, err := mysql.GetMultiEvent()
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}

			for i := 0; i <= len(tab)-1; i++ {
				if tab[i][1] == "terminer" {
					continue
				}

				if tab[i][1] != "prepterminer" {
					t1 := time.Now()
					t2, _ := strconv.Atoi(tab[i][7])
					num, _ := strconv.Atoi(tab[i][0])

					if t1.Unix() >= int64(t2) {
						tab[i][1] = "en cours"
						mysql.EditStatus(2, num)
					}
				}

				ConstructionEmbedEvent(0, session, tab[i])

				if tab[i][1] == "prepterminer" {
					num, _ := strconv.Atoi(tab[i][0])
					mysql.EditStatus(4, num)
				}
			}

			count60seconde = 0
		}

		count60seconde = count60seconde + 5
	}
}
