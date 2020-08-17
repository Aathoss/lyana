package modules

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/spf13/viper"
	"gitlab.com/inovaperf/bot/modules"
	"gitlab.com/unispace/framework"
)

var (
	urlVersion = "https://papermc.io/api/v1/paper"
	urlBuild   = "https://papermc.io/api/v1/paper/1.16.1"

	NotifVersion      bool
	NotifBuild        bool
	NotifBuildVersion string
)

func VerifServerMCVersion() {
	response, err := http.Get(urlVersion)
	modules.CheckError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	matched, err := regexp.Match(viper.GetString("Minecraft.CheckVersion"), body)
	if matched != false && NotifVersion != true && framework.VersionMC != viper.GetString("Minecraft.CheckVersion") {
		LogDiscord("[:pushpin:] La version " + viper.GetString("Minecraft.CheckVersion") + " vient de sortir. <@&735283360322027600>")
		NotifVersion = true
	}
}

func VerifServerMCBuild() {
	response, err := http.Get(urlBuild)
	modules.CheckError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	modules.CheckError(err)

	var retuInfo map[string]interface{}
	json.Unmarshal(body, &retuInfo)
	buildMap := retuInfo["builds"].(map[string]interface{})["latest"].(string)

	if buildMap != NotifBuildVersion {
		NotifBuild = false
	}

	matched, err := regexp.Match(buildMap, body)
	if matched != false && NotifBuild != true {
		LogDiscord("[:pushpin:] Un nouveau build est disponible " + buildMap + ". <@&735283360322027600>")
		NotifBuild = true
		NotifBuildVersion = buildMap
	}
}
