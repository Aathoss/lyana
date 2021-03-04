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
	logger.InfoLogger.Println("----- Démarrage de la boucle UpdateEvent")

	session := framework.Session

	for {
		logger.DebugLogger.Println("debug 1")
		time.Sleep(time.Second * secondeboucle)
		logger.DebugLogger.Println("debug 2")

		//Update de l'embed lors de la création d'un nouvelle évent
		if framework.EventConstruction == true {
			logger.DebugLogger.Println("debug 3")
			tab, err := mysql.GetCreationEvent(framework.EventContructionIndex)
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}

			framework.ConstructionEmbedEvent(0, session, tab)
		}

		//Update multi évent public
		if count60seconde >= 60 {
			logger.DebugLogger.Println("debug 4")
			tab, err := mysql.GetMultiEvent()
			if err != nil {
				logger.ErrorLogger.Println(err)
			}

			logger.DebugLogger.Println("debug 5")
			for i := 0; i <= len(tab)-1; i++ {
				if tab[i][1] == "terminer" {
					continue
				}

				logger.DebugLogger.Println("debug 6")
				if tab[i][1] != "prepterminer" {
					t1 := time.Now()
					t2, _ := strconv.Atoi(tab[i][7])
					num, _ := strconv.Atoi(tab[i][0])

					if t1.Unix() >= int64(t2) {
						tab[i][1] = "en cours"
						mysql.EditStatus(2, num)
					}
				}

				logger.DebugLogger.Println("debug 7")
				framework.ConstructionEmbedEvent(0, session, tab[i])
				logger.DebugLogger.Println("debug 8")

				if tab[i][1] == "prepterminer" {
					num, _ := strconv.Atoi(tab[i][0])
					mysql.EditStatus(4, num)
				}
				logger.DebugLogger.Println("debug 9")
			}
			logger.DebugLogger.Println("debug 10")
			count60seconde = 0
		}
		logger.DebugLogger.Println("debug 11")
		count60seconde = count60seconde + 5
	}

	logger.InfoLogger.Println("----- Arrêt de la boucle UpdateEvent")
}
