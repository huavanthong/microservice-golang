package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/daos"
	"github.com/huavanthong/microservice-golang/user-api/models"

	googleCred "github.com/huavanthong/microservice-golang/user-api/security/google"
	"github.com/huavanthong/microservice-golang/user-api/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Define user manages
type GoogleUser struct {
	utils    utils.Utils
	guserDAO daos.GoogleUser
}

var cred googleCred.Credentials
var conf *oauth2.Config

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {
	file, err := ioutil.ReadFile("./config/google-credentials.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &cred); err != nil {
		log.Println("unable to marshal data")
		return
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// IndexHandler handles the location /.
func (gu *GoogleUser) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func (gu *GoogleUser) AuthHandler(ctx *gin.Context) {
	// Handle the exchange code to initiate a transport.
	// get session where stored info users
	session := sessions.Default(ctx)

	retrievedState := session.Get("state")
	queryState := ctx.Request.URL.Query().Get("state")

	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		ctx.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
		return
	}
	code := ctx.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := models.GoogleUser{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}
	session.Set("user-id", u.Email)
	err = session.Save()
	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	seen := false
	if _, mongoErr := gu.guserDAO.LoadUser(u.Email); mongoErr == nil {
		seen = true
	} else {
		err = gu.guserDAO.SaveUser(&u)
		if err != nil {
			log.Println(err)
			ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving user. Please try again."})
			return
		}
	}

	ctx.HTML(http.StatusOK, "battle.tmpl", gin.H{"email": u.Email, "seen": seen})
}

// LoginHandler handles the login procedure.
func (gu *GoogleUser) LoginHandler(ctx *gin.Context) {
	state, err := googleCred.RandToken(32)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while generating random data."})
		return
	}
	session := sessions.Default(ctx)
	session.Set("state", state)

	err = session.Save()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(state)
	ctx.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}

// FieldHandler is a rudementary handler for logged in users.
func (gu *GoogleUser) FieldHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user-id")
	ctx.HTML(http.StatusOK, "field.tmpl", gin.H{"user": userID})
}