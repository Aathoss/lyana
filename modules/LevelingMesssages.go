package modules

import (
	"image"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

type leveling struct {
	uuid      string
	niveau    int
	xp        int
	timestamp int
}

func LevelingMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		lvl leveling
	)

	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Logs chaque messages pour compter le nombre de message envoyer
	t1 := time.Now()
	insert, err := framework.DBLyana.Prepare("INSERT INTO logs(uuid, categorie, timestamp, content) VALUES(?, ?, ?, ?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer insert.Close()
	_, err = insert.Exec(m.Author.ID, "msgcount", t1.Unix(), "")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	// Système de leveling
	lvl.uuid = m.Author.ID
	lvl.timestamp = int(t1.Unix())
	framework.DBLyana.QueryRow("SELECT niveau,xp FROM level WHERE uuid=?", lvl.uuid).Scan(&lvl.niveau, &lvl.xp)

	lvl.xp = lvl.xp + (rand.Intn(10) + 10)
	calculNiveau := (5 * lvl.niveau) + (100 * lvl.niveau) + 100
	calculNiveauUP := (5 * (lvl.niveau + 1)) + (100 * (lvl.niveau + 1)) + 100

	if lvl.xp >= calculNiveau {
		lvl.niveau++
		lvl.xp = lvl.xp - calculNiveau

		const (
			x         = 800
			y         = 180
			imgprofil = 128
			rectangle = imgprofil / 2
		)

		dc := gg.NewContext(x, y)
		dc.SetHexColor("#E5DBCF")
		dc.Clear()

		response, err := http.Get(m.Author.AvatarURL("128"))
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		defer response.Body.Close()
		mdecode, _, err := image.Decode(response.Body)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		// Ajoute d'un coutour autour du profil en couleurs
		dc.DrawCircle(rectangle+38, y/2, rectangle+5)
		dc.SetHexColor("#373B4D")
		dc.Fill()

		// Ajoute de la photo de profil
		dc.DrawCircle(rectangle+38, y/2, rectangle)
		dc.Clip()
		dc.DrawImage(mdecode, 38, y/2-(imgprofil/2))
		dc.Fill()
		dc.InvertMask()

		// Ajoute du texte de félicitation
		if err := dc.LoadFontFace("/library/fonts/MrDafoe.ttf", 52); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		dc.SetHexColor("#373B4D")
		dc.DrawStringAnchored("Félicitation pour le level up.", x/2+75, 50, 0.5, 0.5)

		// Ajoute du texte
		if err := dc.LoadFontFace("/library/fonts/Raleway-Medium.ttf", 32); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		dc.SetHexColor("#AD9F91")
		dc.DrawStringAnchored("Niveau : "+strconv.Itoa(lvl.niveau)+" | Expérience : "+strconv.Itoa(lvl.xp)+"/"+strconv.Itoa(calculNiveauUP), x/2+75, 120, 0.5, 0.5)

		err = dc.SavePNG("library/leveling/" + lvl.uuid + ".png")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		// Ouverture de l'image pour l'envoie
		png, err := os.Open("library/leveling/" + lvl.uuid + ".png")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		defer png.Close()

		_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			File: &discordgo.File{
				Name:   lvl.uuid + ".png",
				Reader: png,
			},
		})
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
	}

	insert, err = framework.DBLyana.Prepare(`
	INSERT INTO level(uuid, niveau, xp, timestamp)
	VALUES(?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		niveau = ?,
		xp = ?,
		timestamp = ?
		`)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer insert.Close()
	_, err = insert.Exec(lvl.uuid, lvl.niveau, lvl.xp, lvl.timestamp, lvl.niveau, lvl.xp, lvl.timestamp)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
