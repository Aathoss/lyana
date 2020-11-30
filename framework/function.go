package framework

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// VerifStaff retourne un int
func VerifStaff(grade []string) (r int) {

	r = 0
	if len(grade) != 0 {
		for _, s := range grade {
			if s == viper.GetString("Staff.Staff") || s == viper.GetString("Staff.Moderateur") && r <= 0 {
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

func VerifGrade(grade []string) (count int) {
	if len(grade) != 0 {
		for _, s := range grade {
			if s == "757730769023008958" {
				return 1
			}
		}
	}
	return 0
}

func RemoveIndex(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

//timestamp = 0 heure actuel - heure données
// | timestamp = 1 heure données - heure actuel
// | timestamp = 2 calcule uniquement l'heure données
func Calculetime(timesec int64, timestamp int) string {
	seconde := timesec

	var ans int // 60 * 60 * 24 * 30 * 12 = 31 104 000
	var calcans bool
	var mois int // 60 * 60 * 24 * 30 = 2 592 000
	var calcmois bool
	var jours int // 60 * 60 * 24 = 86 400
	var calcjour bool
	var heures int // 60 * 60 = 3 600
	var calcheures bool
	var minutes int // 60 = 60
	var calcminutes bool
	var parse string

	if timestamp == 0 {
		t1 := time.Now()
		seconde = t1.Unix() - timesec
	}

	if timestamp == 1 {
		t1 := time.Now()
		seconde = timesec - t1.Unix()
	}

	for seconde >= 0 {
		if seconde >= 31104000 {
			calcans = true
			ans++
			seconde = seconde - 31104000
			continue
		}
		if seconde >= 2592000 {
			calcmois = true
			mois++
			seconde = seconde - 2592000
			continue
		}
		if seconde >= 86400 {
			calcjour = true
			jours++
			seconde = seconde - 86400
			continue
		}
		if seconde >= 3600 {
			calcheures = true
			heures++
			seconde = seconde - 3600
			continue
		}
		if seconde >= 60 {
			calcminutes = true
			minutes++
			seconde = seconde - 60
			continue
		}
		break
	}

	if calcans == true {
		parse = parse + plural(int(ans), "année")
	}

	if calcmois == true {
		parse = parse + strconv.Itoa(mois) + " mois "
	}

	if calcjour == true {
		parse = parse + plural(int(jours), "jour")
	}

	if calcheures == true {
		parse = parse + plural(int(heures), "heure")
	}

	if calcminutes == true && calcans == false {
		parse = parse + plural(int(minutes), "minute")
	}

	if calcmois == false {
		parse = parse + plural(int(seconde), "seconde")
	}

	return parse
}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + " " + singular + " "
	} else {
		result = strconv.Itoa(count) + " " + singular + "s "
	}
	return
}

func RequestAPI(method string, url string) ([]byte, error) {
	var body []byte

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return body, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return body, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	//fmt.Println(string(body))
	return body, nil
}
