# :artificial_satellite: update-relay

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

App which relays your message to multiple channels

## :muscle: Supported Channels:

Message relays to these channels are currently supported:

- Slack
- Jira Cloud


## :book: How to use:

1. Clone the repo
2. Run the React App and http://localhost:3000 should open up in your browser
```bash
cd ./update-relay
npm start
```
3. Add the Keys and Token values in the ./backend/.env file
4. Start the backend service
```go
cd ./backend
go run main.go
```
5. Post updates using the form in http://localhost:3000