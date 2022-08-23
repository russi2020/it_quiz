package models

import (
	"github.com/jinzhu/gorm"
	u "it_quiz/utils"
	"log"
)

// Category содержит информацию о категориях квиза
type Category struct {
	gorm.Model
	Topic  string `json:"topic" ;sql:"topic"`
	UserId uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

// Theme содержит данные о доступных темах в соответсвующей категории
type Theme struct {
	gorm.Model
	Title      string `json:"title" ;sql:"title"`
	CategoryId uint   `json:"category_id" ;sql:"category_id" ;gorm:"foreignKey:CategoryRefer"`
	UserId     uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

// Question содержит вопрос относящийся к определенной теме
type Question struct {
	gorm.Model
	Title   string `json:"title" ;sql:"title"`
	Text    string `json:"text" ;sql:"text"`
	ThemeId uint   `json:"theme_id" ;sql:"theme_id" ;gorm:"foreignKey:ThemeRefer"`
	UserId  uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

// UserAttempt содержит данные о попытке прохождения тестирования
type UserAttempt struct {
	gorm.Model
	ThemeId             uint `json:"theme_id" ;sql:"theme_id" ;gorm:"foreignKey:ThemeRefer"`
	RightAnswersCounter uint `json:"right_answers_counter" ;sql:"right_answers_counter"` // Здесь количество правильных ответов
	WrongAnswersCounter uint `json:"wrong_answers_counter" ;sql:"wrong_answers_counter"` // Здесь количествл неправильных ответов
	UserId              uint `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

// Answer В базе данных более одного ответа для вопроса, правильный будет иметь поле Correct равным true.
type Answer struct {
	gorm.Model
	QuestionId uint   `json:"question_id" ;sql:"question_id" ;gorm:"foreignKey:QuestionRefer"`
	Text       string `json:"text" ;sql:"text"`
	UserId     uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
	Correct    bool   `json:"correct" ;sql:"correct"`
}

// UserAnswer это ответ пользователя на вопрос
type UserAnswer struct {
	gorm.Model
	UserAttemptId uint   `json:"user_attempt_id" ;sql:"user_attempt_id"`
	AnswerId      uint   `json:"answer_id" ;sql:"answer_id" ;gorm:"foreignKey:AnswerRefer"`
	QuestionId    uint   `json:"question_id" ;sql:"question_id" ;gorm:"foreignKey:QuestionRefer"`
	Text          string `json:"text" ;sql:"text"`
	UserId        uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
	Right         bool   `json:"right" ;sql:"right"`
}

// WikiQuestionUrl дает ссылку на url wiki в случае, если ответ на вопрос неправильный
type WikiQuestionUrl struct {
	gorm.Model
	QuestionId uint   `json:"question_id" ;sql:"question_id" ;gorm:"foreignKey:QuestionRefer"`
	WikiUrl    string `json:"wiki_url" ;sql:"wiki_url"`
}

func (c *Category) Validate() (map[string]interface{}, bool) {

	if c.Topic == "" {
		return u.Message(false, "Category topic should be on the payload"), false
	}

	if c.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (c *Category) CreateCategoryRecord() map[string]interface{} {

	if resp, ok := c.Validate(); !ok {
		return resp
	}

	GetDB().Create(c)

	if c.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	response := u.Message(true, "Category has been created")
	response["categories"] = c
	return response
}

func GetCategoryByName(topic string) map[string]interface{} {

	category := Category{}
	err := GetDB().Table("categories").Where("topic = ?", topic).First(category).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Category not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	resp := u.Message(true, "Get category")
	resp["categories"] = category

	return resp
}

func GetAllCategories() []*Category {
	categories := make([]*Category, 0)
	err := GetDB().Table("categories").Find(&categories).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return categories
}

func GetCategory(id uint) *Category {
	category := Category{}
	err := GetDB().Table("categories").Where("id = ?", id).Find(&category).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &category
}

func (t *Theme) Validate() (map[string]interface{}, bool) {

	if t.Title == "" {
		return u.Message(false, "Theme title should be on the payload"), false
	}

	if t.CategoryId <= 0 {
		return u.Message(false, "Category is not recognized"), false
	}

	if t.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (t *Theme) CreateThemeRecord() map[string]interface{} {

	if resp, ok := t.Validate(); !ok {
		return resp
	}

	GetDB().Create(t)

	if t.ID <= 0 {
		return u.Message(false, "Failed to create theme entry, connection error.")
	}

	response := u.Message(true, "Theme has been created")
	response["themes"] = t
	return response
}

func GetAllThemes() []*Theme {
	themes := make([]*Theme, 0)
	err := GetDB().Table("themes").Find(&themes).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return themes
}

func GetTheme(id uint) *Theme {
	theme := Theme{}
	err := GetDB().Table("themes").Where("id = ?", id).First(&theme).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &theme
}

func (question *Question) Validate() (map[string]interface{}, bool) {

	if question.Title == "" {
		return u.Message(false, "Theme title should be on the payload"), false
	}

	if question.Text == "" {
		return u.Message(false, "Text is not recognized"), false
	}

	if question.ThemeId <= 0 {
		return u.Message(false, "Theme is not recognized"), false
	}

	if question.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (question *Question) CreateQuestionRecord() map[string]interface{} {

	if resp, ok := question.Validate(); !ok {
		return resp
	}

	GetDB().Create(question)

	if question.ID <= 0 {
		return u.Message(false, "Failed to create question entry, connection error.")
	}

	response := u.Message(true, "Question has been created")
	response["questions"] = question
	return response
}

func GetAllQuestions() []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Find(&questions).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return questions
}

func GetQuestion(id uint) *Question {
	question := &Question{}
	err := GetDB().Table("questions").Where("id = ?", id).First(&question).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return question
}

func GetAllQuestionsByThemeId(themeId uint) []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Find(&questions, "theme_id = ?", themeId).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return questions
}

func (userAtt *UserAttempt) Validate() (map[string]interface{}, bool) {

	if userAtt.ThemeId <= 0 {
		return u.Message(false, "Theme is not recognized"), false
	}

	if userAtt.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (userAtt *UserAttempt) CreateAttempt() map[string]interface{} {
	if resp, ok := userAtt.Validate(); !ok {
		return resp
	}

	GetDB().Create(userAtt)

	if userAtt.ID <= 0 {
		return u.Message(false, "Failed to create user attempt entry, connection error.")
	}

	response := u.Message(true, "User attempt has been created")
	response["user_attempt"] = userAtt
	return response
}

func GetAllAttempts() []*UserAttempt {
	userAttempts := make([]*UserAttempt, 0)
	err := GetDB().Table("user_attempt").Find(&userAttempts).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return userAttempts
}

func GetUserAttempt(id uint) *UserAttempt {
	userAttempt := &UserAttempt{}
	err := GetDB().Table("user_attempt").Where("id = ?", id).First(&userAttempt).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return userAttempt
}

func (a *Answer) Validate() (map[string]interface{}, bool) {
	if a.QuestionId <= 0 {
		return u.Message(false, "Question is not recognized"), false
	}
	if a.Text == "" {
		return u.Message(false, "Text is not recognized"), false
	}
	if a.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (a *Answer) Create() map[string]interface{} {
	if resp, ok := a.Validate(); !ok {
		return resp
	}

	GetDB().Create(a)

	if a.ID <= 0 {
		return u.Message(false, "Failed to create answer entry, connection error.")
	}

	response := u.Message(true, "Answer has been created")
	response["answer"] = a
	return response
}

func GetAllAnswers() []*Answer {
	allAnswers := make([]*Answer, 0)
	err := GetDB().Table("answer").Find(&allAnswers).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return allAnswers
}

func GetAnswer(id uint) *Answer {
	answer := &Answer{}
	err := GetDB().Table("answer").Where("id = ?", id).First(&answer).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return answer
}

func GetAnswersByQuestionId(questionId uint) []*Answer {
	answersByQuestionId := make([]*Answer, 0)
	err := GetDB().Table("answers").Find(&answersByQuestionId, "question_id = ?", questionId).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return answersByQuestionId
}

func GetRightAnswerByQuestionId(questionId uint) *Answer {
	rightAnswer := &Answer{}
	err := GetDB().Table("answers").Where("question_id = ? and correct = true", questionId).Find(&rightAnswer).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return rightAnswer
}

func GetRightAnswersByThemeId(themeId uint) []*Answer {
	rightAnswers := make([]*Answer, 0)
	subQuery := GetDB().Select("id").Where("theme_id = ?", themeId).Table("questions")
	err := GetDB().Table("answers").Find(&rightAnswers, "question_id in (?) AND correct = true", subQuery.QueryExpr()).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return rightAnswers
}
