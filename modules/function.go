package modules

import "github.com/spf13/viper"

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
