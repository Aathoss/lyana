package modules

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

type sShop struct {
	shop_solde float64 `json:"shop_solde"`
	shop1      struct {
		define  bool    `json:"shop1_define"`
		surface float64 `json:"shop1_surface"`
		coef    float64 `json:"shop1_coef"`
	}
	shop2 struct {
		define  bool    `json:"shop2_define"`
		surface float64 `json:"shop2_surface"`
		coef    float64 `json:"shop2_coef"`
	}
}

var (
	dateDernierUpdate string
)

func MinecraftCheckShop(secondeboucle time.Duration) {
	logger.InfoLogger.Println("----- [GoRoutine] Démarrage de la boucle MinecraftCheckShop")

	err := framework.DBLyana.QueryRow("SELECT content FROM logs WHERE categorie='varMinecraftActualise'").Scan(&dateDernierUpdate)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	db := framework.DBMinecraft

	for {

		var (
			player  string
			players []string
			result  float64
		)
		t1 := time.Now()
		num, _ := strconv.Atoi(dateDernierUpdate)

		if t1.Day() == num {
			time.Sleep(time.Second * secondeboucle)
			continue
		}

		framework.LogsChannel(":pushpin: **------------------------------[ Lancement du scan des shop ]------------------------------**")

		rows, err := db.Query("SELECT DISTINCT PLAYER FROM PLAYERDATA")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		for rows.Next() {
			err = rows.Scan(&player)
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}
			players = append(players, player)

			err = rows.Err()
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}
		}

		for _, player := range players {
			var (
				shop      sShop
				calcshop1 float64
				calcshop2 float64
			)
			err = db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "'AND VARIABLE='shop_solde'").Scan(&shop.shop_solde)
			if err != nil {
				//logger.ErrorLogger.Println("Impossible de lire les valeurs de l'utilisateur "+player)
				continue
			}

			db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop1_define'").Scan(&shop.shop1.define)
			db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop2_define'").Scan(&shop.shop2.define)

			if shop.shop1.define == true {
				db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop1_surface'").Scan(&shop.shop1.surface)
				db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop1_coef'").Scan(&shop.shop1.coef)
				calcshop1 = calcCoef(shop.shop1.surface)
			}

			if shop.shop2.define == true {
				db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop2_surface'").Scan(&shop.shop2.surface)
				db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='shop2_coef'").Scan(&shop.shop2.coef)
				calcshop2 = calcCoef(shop.shop2.surface)
			}

			if (calcshop1 + calcshop2) == 0 {
				continue
			}

			shop.shop_solde = shop.shop_solde - (calcshop1 + calcshop2)
			result += calcshop1 + calcshop2
			fmt.Println("Calcule du joueurs " + player + " en cours | soldes : " + fmt.Sprintf("%.2f", shop.shop_solde) + " | Mairie : " + fmt.Sprintf("%.2f", result))
			msgLogs := "**Joueurs :** `" + player + "`             **Solde :  " + fmt.Sprintf("%.2f", shop.shop_solde) + " / " + fmt.Sprintf("%.2f", ((calcshop1*7)+(calcshop2*7))*2) + "** Diamants             :moneybag: -" + fmt.Sprintf("%.2f", (calcshop1+calcshop2))

			if shop.shop_solde < -((calcshop1 * 7) + (calcshop2 * 7)) {
				framework.LogsChannel(":heart: " + msgLogs)
				framework.LogsRolePolicier(":heart: " + msgLogs)
			} else if shop.shop_solde < -((calcshop1*7)+(calcshop2*7))/2 {
				framework.LogsChannel(":orange_heart: " + msgLogs)
				framework.LogsRolePolicier(":orange_heart: " + msgLogs)

			} else if shop.shop_solde < 0 {
				framework.LogsChannel(":yellow_heart: " + msgLogs)
				framework.LogsRolePolicier(":yellow_heart: " + msgLogs)
			} else {
				framework.LogsChannel(":green_heart: " + msgLogs)
			}

			update, err := db.Query("UPDATE PLAYERDATA SET content=? WHERE VARIABLE='shop_solde' AND PLAYER=?", fmt.Sprintf("%.2f", shop.shop_solde), player)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
			update.Close()
		}

		for {
			err = framework.ConnectMC[1].Authenticate()
			if err != nil {
				framework.OnlineServer[1] = "offline"
				framework.Connect(1)
			}

			_, err := framework.ConnectMC[1].SendCommand("mycmd-variables add mairie_solde " + fmt.Sprintf("%.2f", result))
			if err != nil {
				logger.ErrorLogger.Println("MC-Host : "+viper.GetString("Minecraft."+strconv.Itoa(1)+".Name")+" | Command failed : ", err)
				time.Sleep(time.Second * 5)
				continue
			}
			framework.LogsChannel(":pushpin: **----------------------------------[ Mairie +" + fmt.Sprintf("%.2f", result) + " Diamants ]----------------------------------**")
			break
		}

		update, err := framework.DBLyana.Query("UPDATE logs SET content = ? WHERE categorie = ?", t1.Day(), "varMinecraftActualise")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		update.Close()
		dateDernierUpdate = strconv.Itoa(t1.Day())
	}
	logger.InfoLogger.Println("----- [GoRoutine] Arrêt de la boucle MinecraftCheckShop")
}

func calcCoef(surface float64) float64 {
	surfaceBase := viper.GetFloat64("Minecraft_shop_surface_basee")
	shopTarifsBase := viper.GetFloat64("Minecraft_shop_tarifs")
	multiplicateurMairieBase := viper.GetFloat64("Minecraft_shop_mairie_variable")

	shopCoef := math.Pow(((surface / surfaceBase) * multiplicateurMairieBase), 1.50)
	prixSemaine := shopTarifsBase * shopCoef
	prixJour := prixSemaine / 7

	return prixJour
}
