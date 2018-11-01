package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	"net/http"
	"crypto/tls"
	"io/ioutil"
)

var URL = "http://alknopfler.ddns.net:8080"

var a = &alexa.Alexa{ApplicationID: "amzn1.ask.skill.cbf2eb74-9b14-4752-b39c-2bec61845037", RequestHandler: &AlkAlarm{}, IgnoreTimestamp: true}

const cardTitle = "AlkAlarm"

// Alkalarm struct for request from the alkalarm skill.
type AlkAlarm struct{}

// Handle processes calls from Lambda
func Handle(ctx context.Context, requestEnv *alexa.RequestEnvelope) (interface{}, error) {
	return a.ProcessRequest(ctx, requestEnv)
}

// OnSessionStarted called when a new session is created.
func (h *AlkAlarm) OnSessionStarted(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *AlkAlarm) OnLaunch(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {
	speechText := "Welcome to AlkAlarm system. Starting the Alarm"

	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	response.SetSimpleCard(cardTitle, speechText)
	response.SetOutputSSML(speechText)
	response.SetRepromptSSML(speechText)

	response.ShouldSessionEnd = true

	return nil
}

// OnIntent called with a reqeust is received of type IntentRequest
func (h *AlkAlarm) OnIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {

	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)

	switch request.Intent.Name {
	case "activeAlarm":
		activateAlarm(request,response)
	case "deactivateAlarm":
		deactivateAlarm(request,response)
	case "statusAlarm":
		statusAlarm(request,response)
	default:
		return errors.New("Invalid Intent")
	}

	return nil
}

// OnSessionEnded called with a reqeust is received of type SessionEndedRequest
func (h *AlkAlarm) OnSessionEnded(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

func activateAlarm(request *alexa.Request, response *alexa.Response){
	log.Println("ActiveAlarm triggered")
	speechText := "AlkAlarm Activada puedes salir con seguridad de casa"

	if len(request.Intent.Slots) == 1 {
		log.Println(request.Intent.Slots["dentrode"].Resolutions)
		delay := request.Intent.Slots["dentrode"].Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
		log.Println("El delay será: "+delay)
	}

	reqNew, _ := http.NewRequest("POST", URL + "/activate/full", nil)
	reqNew.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respNew, _ := client.Do(reqNew)

	if respNew.StatusCode == http.StatusOK {
		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputSSML(speechText)
	}else{
		response.SetSimpleCard(cardTitle, "ERROR DOING THE ACTIVATION ALARM")
		response.SetOutputSSML("ERROR DOING THE ACTIVATION ALARM ")
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}

func deactivateAlarm(request *alexa.Request, response *alexa.Response){
	log.Println("DeactiveAlarm triggered")
	speechText := "AlkAlarm Desactivada puedes entrar con seguridad en casa"

	reqNew, _ := http.NewRequest("POST", URL + "/deactivate", nil)
	reqNew.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respNew, _ := client.Do(reqNew)

	if respNew.StatusCode == http.StatusOK {
		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputSSML(speechText)
	}else{
		response.SetSimpleCard(cardTitle, "ERROR DOING THE DEACTIVATION ALARM")
		response.SetOutputSSML("ERROR DOING THE DEACTIVATION ALARM ")
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}

func statusAlarm(request *alexa.Request, response *alexa.Response){
	log.Println("StatusAlarm triggered")
	activated := "La alarma esta activada"
	desactivated := "La alarma esta desactivada"



	reqNew, _ := http.NewRequest("POST", URL + "/status", nil)
	reqNew.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respNew, _ := client.Do(reqNew)
	data, _ := ioutil.ReadAll(respNew.Body)
	if string(data) == "full" {
		response.SetSimpleCard(cardTitle, activated)
		response.SetOutputSSML(activated)
	}else{
		response.SetSimpleCard(cardTitle, desactivated)
		response.SetOutputSSML(desactivated)
	}

	log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
}

func main() {
	lambda.Start(Handle)
}