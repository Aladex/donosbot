## Description

There's FSB bot for sending messages to FSB department

## Install

```
go get github.com/go-telegram-bot-api/telegram-bot-api
go get github.com/spf13/viper
go build 
```

Change config for your bot and setup proxy pass from nginx to the port from config via http

Enjoy

## Nginx example

```
location /my-bot-token {
    proxy_pass http://127.0.0.1:3330/my-bot-token; 
    proxy_set_header Host $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect default;
}

```