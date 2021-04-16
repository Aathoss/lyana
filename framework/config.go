package framework

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	mcrcon "gitlab.com/lyana/library/package/mc_rcon"
	"gitlab.com/lyana/logger"
)

var (
	//Variable de version
	Version = "0.6.0"

	err         error
	DBLyana     *sql.DB
	DBMinecraft *sql.DB

	countGoRoutineMC int
)

//LoadConfiguration charge les paramètres / variables
func init() {

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	logger.InfoLogger.Println("----- [Lyana] Démarrage du bot")
	logger.InfoLogger.Println("----- [Config] en préparation")

	//Configuration de l'heure sûr le serveur
	logger.InfoLogger.Println("----- [Config] Initialisation l'heure")
	os.Setenv("TZ", "Europe/Paris")

	//Chargement de la configuration du serveur
	logger.InfoLogger.Println("----- [Config] Initialisation du fichier de config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		logger.ErrorLogger.Println(err)
		os.Exit(10)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.InfoLogger.Println("Config file changed:", e.Name)
	})

	//Démarrage des goroutine pour les connexion rcon
	OnlinePlayer = make([]int, len(viper.GetStringMapString("Minecraft")))
	ListPlayer = make([]string, len(viper.GetStringMapString("Minecraft")))
	OnlineServer = make([]string, len(viper.GetStringMapString("Minecraft")))
	ConnectMC = make([]*mcrcon.MCConn, len(viper.GetStringMapString("Minecraft")))
	countMAP := len(viper.GetStringMapString("Minecraft"))
	if viper.GetBool("Dev.test") != true {
		for i := 0; i < countMAP; i++ {

			logger.InfoLogger.Println("----- [Config] Initialisation de la connexion rcon [-" + viper.GetString("Minecraft."+strconv.Itoa(i)+".Name") + "-]")
			go StartRCON(i)

			//time.Sleep(time.Second * 1)
		}
	}

	if viper.GetBool("MySql.Lyana.online") == true {
		//Connexion à la base de données lyana
		logger.InfoLogger.Println("----- [Config] Initialisation de la base de données [-Lyana-]")
		dbUser := viper.GetString("MySql.Lyana.dbuser")
		dbPass := viper.GetString("MySql.Lyana.dbmdp")
		dbName := viper.GetString("MySql.Lyana.dbname")
		dbIP := viper.GetString("MySql.Lyana.dbip")
		dbPort := viper.GetString("MySql.Lyana.dbport")

		DBLyana, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		if err = DBLyana.Ping(); err != nil {
			logger.ErrorLogger.Println(err)
		}
	}

	if viper.GetBool("MySql.Minecraft.online") == true {
		//Connexion à la base de données minecraft
		logger.InfoLogger.Println("----- [Config] Initialisation de la base de données [-Minecraft-]")
		dbUserMC := viper.GetString("MySql.Minecraft.dbuser")
		dbPassMC := viper.GetString("MySql.Minecraft.dbmdp")
		dbNameMC := viper.GetString("MySql.Minecraft.dbname")
		dbIPMC := viper.GetString("MySql.Minecraft.dbip")
		dbPortMC := viper.GetString("MySql.Minecraft.dbport")
		DBMinecraft, err = sql.Open("mysql", dbUserMC+":"+dbPassMC+"@tcp("+dbIPMC+":"+dbPortMC+")/"+dbNameMC)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		if err = DBMinecraft.Ping(); err != nil {
			logger.ErrorLogger.Println(err)
		}
	}

	if err != nil {
		os.Exit(10)
	}
	logger.InfoLogger.Println("----- [Config] Configuration charger")
}
