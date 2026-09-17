package main

import{
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
}

func main() {
	log.SetReportCaller(true)
	var r=chi.Mux=chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil{
		log.Error(err)
	}
}

var UnAuthorizedError=errors.New("Invalid username or token")

func Authorization(next http.Handler) http.handler{
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		var username string=r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username == "" || token == ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err=tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
			return 
		}

		var loginDetails *tools.LoginDetails
		loginDetails=(*database).GetUserLoginDetails(username)
		if(loginDetails==nil || (token!=(*&loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}