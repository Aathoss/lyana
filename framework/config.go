package framework

import (
	"database/sql"
	"os"

	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gitlab.com/lyana/logger"
)

var (
	err     error
	DBLyana *sql.DB
)

//LoadConfiguration charge les paramètres / variables
func init() {
	logger.InfoLogger.Println("----- Démarrage du bot [Lyana]")
	logger.InfoLogger.Println("----- Configuration en préparation")

	//Configuration de l'heure sûr le serveur
	logger.InfoLogger.Println("----- Initialisation l'heure")
	os.Setenv("TZ", "Europe/Paris")

	//Chargement de la configuration du serveur
	logger.InfoLogger.Println("----- Initialisation du fichier de config")
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

	//Connexion à la base de données lyana
	logger.InfoLogger.Println("----- Initialisation de la base de données")
	dbDriver := "mysql"
	dbUser := viper.GetString("MySql.Lyana.dbuser")
	dbPass := viper.GetString("MySql.Lyana.dbmdp")
	dbName := viper.GetString("MySql.Lyana.dbname")
	dbIP := viper.GetString("MySql.Lyana.dbip")
	dbPort := viper.GetString("MySql.Lyana.dbport")

	DBLyana, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err = DBLyana.Ping(); err != nil {
		logger.ErrorLogger.Println(err)
		os.Exit(10)
	}

	DBLyana.SetConnMaxLifetime(150)
	DBLyana.SetMaxOpenConns(2)
	DBLyana.SetConnMaxIdleTime(300)
	//DBLyana.SetMaxIdleConns(0)
	//DBLyana.SetMaxOpenConns(5)

	logger.InfoLogger.Println("----- Configuration charger")
}
