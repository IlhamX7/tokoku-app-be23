# Tokoku App Project BE 23

## Requirements
For running this project please install go. Go:
```
go version
go1.22.3 darwin/arm64
```

## The software you need
Recommended software to install. Example:
```
Visual Studio Code is editor code
Supabase is an open source Firebase alternative.
```

## Getting Started
Init folder name:
```
go mod init tokoku-app-be23
```

## Setup supabase
```
Login supabase
Insert project
Select menu table editor
Create new schema
Rename "tokoku"
```

## Setup database:
```
Add .env
setup env value according to your database settings, for example:
poshost= is your host name
posuser= is your username
pospw= is your password
posport= is port your database
dbname= is your db name
```

## Install library:
install several libraries needed for the project, for example:
```
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

## Run the app:
```
open terminal
go run main.go
```