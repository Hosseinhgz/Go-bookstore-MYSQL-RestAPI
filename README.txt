Dependencies and Instalations:
go mod init github.com/github.com/Hosseinhgz/{your-repository-name}

// for connection and interact with Database
go get "github.com/jinzhu/gorm"

// MYSQL inside the gorm
go get "github.com/jinzhu/gorm/dialects/mysql"

// Gorilla
go get "github.com/gorilla/mux"




// After code is finished:
// go to cmd/main folder (it should contain main.go file)
go build
// after that you should fix your errors if you have some
