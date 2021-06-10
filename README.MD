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
add existing config.yml with cp from example :
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
firstly you need this package to run the migration https://github.com/golang-migrate/migrate
to install the package you need to run this
#### MacOS

```bash
$ brew install golang-migrate
```

#### Windows

Using [scoop](https://scoop.sh/)

```bash
$ scoop install migrate
```

#### Linux (*.deb package)

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```
#### usage
```bash
$ migrate -help
```
and use this command to do the migration
```bash
$ migrate -database postgresql://postgres:postgres@localhost:5432/products -path product/migration/ up
```
#### With Go toolchain

##### Versioned

```bash
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate
$ cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ git checkout $TAG  # e.g. v4.1.0
$ # Go 1.15 and below
$ go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ # Go 1.16+
$ go install -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG
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
1. Open API (swagger), you can access it on (http://localhost:3000/swagger/index.html)
2. DB Migration using golang-migrate (https://github.com/golang-migrate/migrate)
3. validation by struct
4. using config with viper (https://github.com/spf13/viper)
5. using echo HTTP framework (https://echo.labstack.com/) 
6. using grom (https://gorm.io/) as ORM 

## License
This package is open-sourced software licensed under the MIT license.