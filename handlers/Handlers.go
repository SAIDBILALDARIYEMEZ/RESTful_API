package handlers

import (
	"RESTful_API/dataloaders"
	. "RESTful_API/models"
	"encoding/json"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Sayfa nesnesi oluşturma
	page := Page{
		ID:          7,
		Name:        "Kullanıcılar",
		Description: "Kullanıcı Listesi",
		URI:         "/users",
	}

	// Verileri yükleme
	users := dataloaders.LoadUsers()
	interests := dataloaders.LoadInterests()
	interestsMappings := dataloaders.LoadInterestMappings()

	// Yeni kullanıcı listesi oluşturma
	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestsMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	// JSON verisini oluşturma
	viewModel := UserViewModel{
		Page:  page,
		Users: newUsers,
	}
	data, _ := json.Marshal(viewModel)

	// Yanıtı gönderme
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
