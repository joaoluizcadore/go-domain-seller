# Go Domain Seller

It is a golang web server to show a web page to tell the visitor the current domain is for sale.
there is a form to make contact.

## Features

1. [x] Send the form data to your email
2. [x] Docker infraestructure
3. [x] Easily configuration

## Configuration

You should create a file named .env inside the root folder of this project.

```bash
# Basic SMTP config.
SMTP_HOST=
SMTP_USERNAME=
SMTP_PASSWORD=
SMTP_PORT=587
SMTP_TLS=true

# The email address to receive the messages
SEND_NOTIFICATION_TO=

# If you choose a different port, you should update the docker files.
SERVER_PORT=8080
```

## Run

```bash
$ docker-compose up -d
```
