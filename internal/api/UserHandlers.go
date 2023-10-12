package api

import (
	"encoding/json"
	"net/http"
)

// Структура для хранения информации о пользователе
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Проверка, что логин уникален

	// Хеширование пароля перед сохранением в базу данных
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// dbInit - объект Postgre
	// err = dbInit.InsertUser(User{
	// 	Login:    user.Login,
	// 	Password: string(hashedPassword),
	// })

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}
