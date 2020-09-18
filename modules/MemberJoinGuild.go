package modules

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"gitlab.com/lyana/logger"
)

const (
	x         = 800
	y         = 180
	imgprofil = 128
	rcircle   = imgprofil / 2
)

func MemberJoinGuild(username, url string) {

	backgroundGen()

	//dc := gg.NewContext(x, y)
	imgBackground, err := gg.LoadImage("library/img/bienvenue1.png")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	dc := gg.NewContext(x, y)
	dc.DrawRoundedRectangle(0, 0, x, y, 5)
	dc.Clip()
	dc.DrawImage(imgBackground, 0, 0)

	addProfilPicture(dc, url)
	addLogo(dc)
	addPseudo(dc, username)
	addBienvenueText(dc)

	err = dc.SavePNG("library/img/bienvenue2.png")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}

func addProfilPicture(dc *gg.Context, url string) {
	response, err := http.Get(url)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer response.Body.Close()
	m, _, err := image.Decode(response.Body)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	dc.DrawCircle(rcircle+38, y/2, rcircle)
	dc.Clip()
	dc.DrawImage(m, 38, y/2-(imgprofil/2))
	dc.Fill()
	dc.InvertMask()
}

func addLogo(dc *gg.Context) {
	logo, err := gg.LoadImage("library/img/logo.png")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	dc.DrawImage(logo, x/2-(logo.Bounds().Dx()/2)+75, 24)
	dc.Fill()
}

func addPseudo(dc *gg.Context, username string) {
	if len(username) > 32 {
		username = username[:32] + " ..."
	}

	fontPath, err := findfont.Find("arial.ttf")
	if err != nil {
		panic(err)
	}
	logger.DebugLogger.Println("Found 'arial.ttf' in '%s'\n", fontPath)

	// load the font with the freetype library
	/* fontData, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(fontData)
	if err != nil {
		panic(err)
	} */

	if err := dc.LoadFontFace("library/fonts/whitney-medium.ttf", 28); err != nil {
		panic(err)
	}
	dc.SetRGBA255(98, 98, 98, 255)
	dc.DrawStringAnchored(string("◇─ℒ𝓾𝓬𝓪𝓼𝓣𝓸𝓷𝓲_𝟓𝟑─◇#9171"), x/2+75, y/2, 0.5, 0.5)
}

func addBienvenueText(dc *gg.Context) {
	if err := dc.LoadFontFace("library/fonts/Arial.ttf", 26); err != nil {
		panic(err)
	}
	dc.SetRGBA255(98, 98, 98, 255)
	dc.DrawStringAnchored("Nous vous souhaitons la bienvenue !", x/2+75, y/2+45, 0.5, 0.5)
}

func backgroundGen() {
	dc := gg.NewContext(x, y)
	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	rand.Seed(time.Now().UnixNano())
	for ix := 5; ix <= x; ix += 14 {
		for iy := 10; iy <= y; iy += 14 {
			r := rand.Intn(1000)
			if r < 10 {
				dc.SetHexColor("5f27cd") // Violet
			}
			if r >= 10 && r < 240 {
				dc.SetHexColor("f1f1f1") // Gris 12%
			}
			if r >= 240 && r < 470 {
				dc.SetHexColor("ebebeb") // Gris 4%
			}
			if r >= 470 && r < 700 {
				dc.SetHexColor("e0e0e0") // Blanc 2%
			}
			if r >= 700 && r <= 1000 {
				dc.SetHexColor("ffffff") // Blanc 100%
			}
			dc.DrawRectangle(float64(ix), float64(iy), 5, 5)
			dc.Fill()
		}
	}

	srcImage := dc.Image()
	dstImage := image.NewRGBA(srcImage.Bounds())
	graphics.Blur(dstImage, srcImage, &graphics.BlurOptions{StdDev: 0.75})

	newImage, err := os.Create("library/img/bienvenue1.png")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer newImage.Close()
	png.Encode(newImage, dstImage)
}
