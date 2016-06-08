package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	router                              *mux.Router
	slackToken, slackChannel, slackTeam string
)

func init() {
	router = mux.NewRouter()

	slackToken = os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Fatal("YOU HAVE TO SET SLACK_TOKEN")
	}

	slackChannel = os.Getenv("SLACK_CHANNEL")
	if slackChannel == "" {
		log.Fatal("YOU HAVE TO SET SLACK_CHANNEL")
	}

	slackTeam = os.Getenv("SLACK_TEAM")
	if slackTeam == "" {
		log.Fatal("YOU HAVE TO SET SLACK_TEAM")
	}

}
func main() {

	// GET HEAD POST PUT OPTIONS DELETE TRACE CONNECT
	router.Methods("HEAD", "OPTIONS", "TRACE", "CONNECT", "PUT", "DELETE").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Nope!", http.StatusMethodNotAllowed)
	})

	router.Methods("POST").Path("/deploy").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/*
		   token=AzwMDTFpcgdD3YtPMYHpsRQu
		   team_id=T0001
		   team_domain=example
		   channel_id=C2147483705
		   channel_name=test
		   user_id=U2147483697
		   user_name=Steve
		   command=/weather
		   text=94070
		   response_url=https://hooks.slack.com/commands/1234/5678
		*/
		incomingToken := r.FormValue("token")
		incomingChannel := r.FormValue("channel_name")
		incomingTeam := r.FormValue("team_domain")

		log.Print("incoming request....")

		if incomingToken != slackToken {
			http.Error(w, "BAD TOKEN", http.StatusUnauthorized)
			return
		}

		if slackChannel != incomingChannel {
			http.Error(w, "BAD CHANNEL", http.StatusForbidden)
			return
		}

		if slackTeam != incomingTeam {
			http.Error(w, "BAD TEAM", http.StatusForbidden)
			return
		}

		/*
		   {
		       "response_type": "in_channel",
		       "text": "It's 80 degrees right now.",
		       "attachments": [
		           {
		               "text":"Partly cloudy today and tomorrow"
		           }
		       ]
		   }
		*/

		responseTemplate := `{
    "response_type": "in_channel",
    "text": "{{.Text}}",
    "attachments": [
        {
            "text":"{{.Text}}"
        }
    ]
}`

		data := struct {
			Text string
		}{"Test"}
		tmpl, err := template.New("json").Parse(responseTemplate)
		if err != nil {
			http.Error(w, "Template Problem", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
		w.Write([]byte("thanks\n"))

		return

	})

	http.Handle("/", router)
	httpServer := http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(httpServer.ListenAndServe())
}
