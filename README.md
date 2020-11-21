## Snippetbox in GO

Snippetbox is a server side rendered application built with Go which lets people paste and share snippets of text, a bit like Pastebin or GitHub’s Gists. 

This repo contains code that's given in the book Let's Go

### Features
 - users are able to save and view snippets via the app.
 - Snippets will be saved to a MySql database.
 - User authentication and authorization is included.
 - Supports HTTPS.
 - Uses a express like middleware based approach

### Usage
 - Use air -c .air.toml for live reloading
 - Use go build -o ./tmp/main.exe ./src/cmd/web/* && ./tmp/main.exe for development
 - Use go build -o ./tmp/main.exe ./src/cmd/web/* && ./tmp/main.exe --addr=":5000" to change address