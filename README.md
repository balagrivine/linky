# Linky
Linky is a lightweight, efficient, and easy-to-use URL shortener service that allows users to transform long, cumbersome URLs into short, manageable links. Whether you're sharing links on social media, in emails, or anywhere else online, Linky provides a seamless way to keep your URLs concise and user-friendly.

## With Linky you can:
**Shorten Long URLs: Convert long, complex URLs into short, memorable links.
**Create Custom Aliases: Create custom short URLs that are meaningful and easy to remember

## Clone the repository
Clone the linky repository using the command below to get started.
```shell
git clone https://github.com/balagrivine/linky.git
```

# Database Setup:
The project uses PostgreSQL. While Windows is technically supported, using a Unix-like system (macOS or Linux) is recommended for better long-term compatibility.
For database interaction sqlc is used to generate type-safe code. If you don't have sqlc installed visit ['sqlc official documantation page']("https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html") to get started.

```shell
sqlc generate
```
For database migrations, the project uses goose. Follow the steps below to get started
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest

export PATH=$PATH:$HOME/go/bin
```
Run the command below to get up-to date with the migrations.
```shell
GOOSE_DRIVER=postgres GOOSE_DBSTRING="user={db_user} dbname={db_name} sslmode=disable" goose -dir ./migrations/schema up
```

Run the command below to download project dependancies.
```shell
go mod download go.mod go.sum
```

Goodluck with you coding!!!!
