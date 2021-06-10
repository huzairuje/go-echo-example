# Go Echo Example

## Installation
You can install and run the project with these steps.
#### 1. clone this project
Go into your go work directory
```$xslt
    cd GO-WORK-DIRECTORY/
```

#### 2. install dependencies
install the dependencies from go mod with
```$xslt
    go mod install
```

#### 3. configuration
add existing config.yml with cp from example
```$xslt
    cp config.yml.example config.yml
```
and then fill the necessary data, example
```$xslt
ServerPort: ":3000"
DatabaseHost: "localhost"
DatabasePort: "5432"
DatabaseUser: "postgres"
DatabasePassword: "postgres"
DatabaseName: "products"
```

#### 4. migrate database
migrate the database, with this command
```$xslt
    go mod install
```
#### 5. run the server
go to this project root folder
```$xslt
    cd ROOT-FOLDER-PROJECT/
```
and then
```$xslt
    go run main.go
```
## Features

## License
This package is open-sourced software licensed under the MIT license.