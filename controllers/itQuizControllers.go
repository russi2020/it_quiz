package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"it_quiz/models"
	u "it_quiz/utils"
	"net/http"
	"strconv"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	category := &models.Category{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	category.UserId = user
	resp := category.CreateCategoryRecord()
	u.Respond(w, resp)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllCategories()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetCategory(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func CreateTheme(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	theme := &models.Theme{}
	err := json.NewDecoder(r.Body).Decode(theme)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	theme.UserId = user
	resp := theme.CreateThemeRecord()
	u.Respond(w, resp)
}

func GetThemes(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllThemes()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetTheme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetTheme(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	question := &models.Question{}
	err := json.NewDecoder(r.Body).Decode(question)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	question.UserId = user
	resp := question.CreateQuestionRecord()
	u.Respond(w, resp)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllQuestions()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetQuestion(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	answer := &models.Answer{}
	err := json.NewDecoder(r.Body).Decode(answer)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	answer.UserId = user
	resp := answer.Create()
	u.Respond(w, resp)
}

func GetAnswers(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllAnswers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetAnswer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAnswer(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetAnswersByQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["question_id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAnswersByQuestionId(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetRightAnswerByQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionId, err := strconv.Atoi(params["question_id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetRightAnswerByQuestionId(uint(questionId))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetRightAnswersByTheme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	themeId, err := strconv.Atoi(params["theme_id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetRightAnswersByThemeId(uint(themeId))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
