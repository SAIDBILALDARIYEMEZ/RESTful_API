package dataloaders

import (
	"RESTful_API/models"
	"RESTful_API/utils"
	"encoding/json"
)

func LoadUsers() []models.User {
	bytes, _ := utils.ReadFile("./json/users.json")
	var data []models.User
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterests() []models.Interest {
	bytes, _ := utils.ReadFile("./json/interests.json")
	var data []models.Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMappings() []models.InterestMapping {
	bytes, _ := utils.ReadFile("./json/userInterestMappings.json")
	var data []models.InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
