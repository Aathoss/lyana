package framework

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// VerifStaff retourne un int
func VerifStaff(grade []string) (r int) {

	r = 0
	if len(grade) != 0 {
		for _, s := range grade {
			if s == viper.GetString("Staff.Moderateur") || s == viper.GetString("Staff.Moderatrice") && r <= 0 {
				r = 1
			}
			if s == viper.GetString("Staff.Responsable") && r <= 2 {
				r = 2
			}
			if s == viper.GetString("Staff.Administrateur") && r <= 3 {
				r = 3
			}
		}
	}
	return
}

func RemoveIndex(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func FormatTime(timestamp int64) {
	/* date1 := time.Date(0, 0, 0, 0, 0, 36, 0, time.UTC)
	date2 := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	if date1.After(date2) {
		date1, date2 = date2, date1
	}

	stringdate := diff(date1, date2) */

	Calculetime(1597530893, 0)
}

//timestamp = 0 heure actuel - heure données | timestamp = 1 calcule uniquement l'heure données
func Calculetime(timesec int64, timestamp int) string {
	seconde := timesec

	/* 	var ans int     // 60 * 60 * 24 * 30 * 12 = 31 104 000
	 */var mois int // 60 * 60 * 24 * 30 = 2 592 000
	var jours int   // 60 * 60 * 24 = 86 400
	var heures int  // 60 * 60 = 3 600
	var minutes int // 60 = 60

	if timestamp != 1 {
		t1 := time.Now()
		seconde = t1.Unix() - timesec
	}

	fmt.Println(seconde)
	for seconde >= 0 {
		/* if seconde >= 31104000 {
			ans++
			seconde = seconde - 31104000
			continue
		} */
		if seconde >= 2592000 {
			mois++
			seconde = seconde - 2592000
			continue
		}
		if seconde >= 86400 {
			jours++
			seconde = seconde - 86400
			continue
		}
		if seconde >= 3600 {
			heures++
			seconde = seconde - 3600
			continue
		}
		if seconde >= 60 {
			minutes++
			seconde = seconde - 60
			continue
		}
		break
	}

	/* fmt.Print(ans)
	fmt.Print(" Années ") */

	/* fmt.Print(mois)
	fmt.Print(" mois ")
	fmt.Print(jours)
	fmt.Print(" jours ")
	fmt.Print(heures)
	fmt.Print(" heures ")
	fmt.Print(minutes)
	fmt.Print(" minutes ")
	fmt.Print(seconde)
	fmt.Println(" secondes ") */
	return strconv.Itoa(mois) + " mois " + strconv.Itoa(jours) + " jours " + strconv.Itoa(heures) + " heures " + strconv.Itoa(minutes) + " minutes " + strconv.Itoa(int(seconde)) + " secondes "

}
