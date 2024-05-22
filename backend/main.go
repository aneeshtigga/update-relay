package main

import (
    "encoding/base64"
    "encoding/json"
    "net/http"
    "os"

	"github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/go-resty/resty/v2"
)

type Update struct {
    Message string `json:"message"`
}

func relayUpdate(w http.ResponseWriter, r *http.Request) {
    var update Update
    json.NewDecoder(r.Body).Decode(&update)

    slackToken := os.Getenv("SLACK_USER_TOKEN")
    slackChannel := os.Getenv("SLACK_CHANNEL_ID")
    // sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
    jiraApiToken := os.Getenv("JIRA_EMAIL_ID")+":"+os.Getenv("JIRA_API_TOKEN")
    jiraAuthToken := base64.StdEncoding.EncodeToString([]byte(jiraApiToken))
    jiraIssueID := os.Getenv("JIRA_ISSUE_ID")

    client := resty.New()

    // Send to Slack
    client.R().
        SetHeader("Authorization", "Bearer "+slackToken).
        SetBody(map[string]string{
            "channel": slackChannel,
            "text":    update.Message,
        }).
        Post("https://slack.com/api/chat.postMessage")

    // // Send Email (using SendGrid)
    // client.R().
    //     SetHeader("Authorization", "Bearer "+sendgridAPIKey).
    //     SetBody(map[string]interface{}{
    //         "personalizations": []map[string]interface{}{
    //             {"to": []map[string]string{{"email": "recipient@example.com"}}},
    //         },
    //         "from":    map[string]string{"email": "sender@example.com"},
    //         "subject": "New Update",
    //         "content": []map[string]string{{"type": "text/plain", "value": update.Message}},
    //     }).
    //     Post("https://api.sendgrid.com/v3/mail/send")

    // Post to Jira
    client.R().
        SetHeader("Authorization", "Basic "+jiraAuthToken).
        SetBody(map[string]string{
            "body": update.Message,
        }).
        Post("https://yourjiraspace.atlassian.net/rest/api/2/issue/" + jiraIssueID + "/comment")

    w.WriteHeader(http.StatusOK)
}

func main() {
    godotenv.Load()

    r := mux.NewRouter()
    r.HandleFunc("/api/updates", relayUpdate).Methods("POST")

	corsHandler := cors.Default().Handler(r)

    http.ListenAndServe(":5000", corsHandler)
}
