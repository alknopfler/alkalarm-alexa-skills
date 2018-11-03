package function

import (
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	"log"
	"net/http"
	"crypto/tls"
	cfg "github.com/alknopfler/alkalarm-alexa-skills/config"

	"io/ioutil"
	"time"
)

func ActivateAlarmFull(request *alexa.Request, response *alexa.Response){
	log.Println("ActiveAlarm Full triggered")

	if len(request.Intent.Slots) == 1 {
		delay := request.Intent.Slots["dentrode"].Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
		log.Println("El delay será: "+delay)
		response.SetOutputText(cfg.SpeechDelay + delay + " segundos")
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechDelay, cfg.ImageSmall, cfg.ImageLong)
		d , _ := time.ParseDuration(delay+"s")
		time.Sleep(time.Second * d)
	}

	respNew := doRequest(http.MethodPost, cfg.URL + cfg.PathActivateFull)

	if respNew.StatusCode == http.StatusOK {
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnActivate, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnActivate)
	}else{
		response.SetSimpleCard(cfg.CardTitle, "ERROR DOING THE ACTIVATION ALARM")
		response.SetOutputText("ERROR DOING THE ACTIVATION ALARM ")
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}

func ActivateAlarmPartial(request *alexa.Request, response *alexa.Response){
	log.Println("ActiveAlarm Partial triggered")

	if len(request.Intent.Slots) == 1 {
		delay := request.Intent.Slots["dentrode"].Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
		log.Println("El delay será: "+delay)
		response.SetOutputText(cfg.SpeechDelay + delay + " segundos")
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechDelay, cfg.ImageSmall, cfg.ImageLong)
		d , _ := time.ParseDuration(delay+"s")
		time.Sleep(time.Second * d)
	}

	respNew := doRequest(http.MethodPost, cfg.URL + cfg.PathActivatePartial)

	if respNew.StatusCode == http.StatusOK {
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnActivate, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnActivate)
	}else{
		response.SetSimpleCard(cfg.CardTitle, "ERROR DOING THE ACTIVATION ALARM")
		response.SetOutputText("ERROR DOING THE ACTIVATION ALARM ")
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}


func DeactivateAlarm(request *alexa.Request, response *alexa.Response){
	log.Println("DeactiveAlarm triggered")

	respNew := doRequest(http.MethodPost, cfg.URL + cfg.PathDeactivate)

	if respNew.StatusCode == http.StatusOK {
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnDeactivate, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnDeactivate)
	}else{
		response.SetSimpleCard(cfg.CardTitle, "ERROR DOING THE DEACTIVATION ALARM")
		response.SetOutputText("ERROR DOING THE DEACTIVATION ALARM ")
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}

func StatusAlarm(request *alexa.Request, response *alexa.Response){
	log.Println("StatusAlarm triggered")

	respNew := doRequest(http.MethodGet, cfg.URL + cfg.PathStatus)
	body, _ := ioutil.ReadAll(respNew.Body)
	log.Println("el body es:" + string(body))
	switch string(body) {
	case "\"inactive\"":
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnStatusOFF, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnStatusOFF)
	case "\"full\"":
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnStatusONFull, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnStatusONFull)
	case "\"partial\"":
		response.SetStandardCard(cfg.CardTitle, cfg.SpeechOnStatusONPartial, cfg.ImageSmall, cfg.ImageLong)
		response.SetOutputText(cfg.SpeechOnStatusONPartial)
	default:
		response.SetSimpleCard(cfg.CardTitle, "ERROR DOING THE STATUS ALARM")
		response.SetOutputText("ERROR DOING THE STATUS ALARM ")

	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}


func doRequest(method, apiURL string) *http.Response{
	reqNew, _ := http.NewRequest(method, apiURL, nil)
	reqNew.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respNew, _ := client.Do(reqNew)
	return respNew
}