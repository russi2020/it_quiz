package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"it_quiz/app"
	"it_quiz/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Printf("Starting on port %s\n", port)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")

	//Category routes
	router.HandleFunc("/api/it_quiz/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/it_quiz/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/it_quiz/categories/{id}", controllers.GetCategory).Methods("GET")

	//Theme routes
	router.HandleFunc("/api/it_quiz/themes", controllers.CreateTheme).Methods("POST")
	router.HandleFunc("/api/it_quiz/themes", controllers.GetThemes).Methods("GET")
	router.HandleFunc("/api/it_quiz/themes/{id}", controllers.GetTheme).Methods("GET")
	router.HandleFunc("/api/it_quiz/themes/answers/{theme_id}", controllers.GetRightAnswersByTheme).Methods("GET")

	//Question routes
	router.HandleFunc("/api/it_quiz/questions", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/it_quiz/questions", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/api/it_quiz/questions/{id}", controllers.GetQuestion).Methods("GET")
	router.HandleFunc("/api/it_quiz/questions/answers/{question_id}", controllers.GetAnswersByQuestion).Methods("GET")
	router.HandleFunc("/api/it_quiz/questions/answers/right/{question_id}", controllers.GetRightAnswerByQuestion).Methods("GET")

	//Answer routes
	router.HandleFunc("/api/it_quiz/answers", controllers.CreateAnswer).Methods("POST")
	router.HandleFunc("/api/it_quiz/answers", controllers.GetAnswers).Methods("GET")
	router.HandleFunc("/api/it_quiz/answers/{id}", controllers.GetAnswer).Methods("GET")

	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		log.Println(err)
	}
}
