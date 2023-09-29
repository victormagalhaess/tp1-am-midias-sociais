package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/victormagalhaess/go-api-template/pkg/model"
)

func ensureFolderExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return os.MkdirAll(folderPath, 0755)
	}
	return nil
}

// Function to save a struct to a file
func saveStructToFile(filename string, data interface{}) error {
	// Encode the struct as JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	path := "data"

	if ensureFolderExists(path) != nil {
		fmt.Printf("Error creating the folder")
	}

	// Write the JSON data to the file
	err = os.WriteFile(fmt.Sprintf("%s/%s", path, filename), jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DoGetRequest(url string, user string) (*http.Response, error) {
	res, err := http.Get(fmt.Sprintf(url, user))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUsers(user string) (*model.User, error) {
	res, err := DoGetRequest("https://api-hexastats.vercel.app/summoners/br1/%s", user)
	if err != nil {
		return nil, err
	}

	bodyUsers, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	unmarshalledUser := &model.User{}
	err = json.Unmarshal(bodyUsers, unmarshalledUser)
	if err != nil {
		return nil, err
	}
	return unmarshalledUser, nil
}

func GetStats(user string) (*model.Stats, error) {
	res, err := DoGetRequest("https://api-hexastats.vercel.app/summoners/br1/%s/stats", user)
	if err != nil {
		return nil, err
	}

	bodyStats, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	unmarshalledStats := &model.Stats{}
	err = json.Unmarshal(bodyStats, unmarshalledStats)
	if err != nil {
		return nil, err
	}
	return unmarshalledStats, nil
}

// func Mine(request *model.Request) {
// user, err := GetUsers(request.StartUser)
// if err != nil {
// 	return
// }
// saveStructToFile(fmt.Sprintf("%s_user.json", request.StartUser), user)

// 	stats, err := GetStats(request.StartUser)
// 	if err != nil {
// 		return
// 	}
// 	saveStructToFile(fmt.Sprintf("%s_stats.json", request.StartUser), stats)
// }

func Mine(request *model.Request) {
	visitedUsers := make(map[string]bool)
	depth := 4
	snowballSampling(request.StartUser, depth, visitedUsers)
}

func snowballSampling(user string, depth int, visitedUsers map[string]bool) {
	if depth <= 0 || visitedUsers[user] {
		return
	}

	visitedUsers[user] = true

	stats, err := GetStats(user)
	if err != nil {
		panic(err)
	}
	saveStructToFile(fmt.Sprintf("%s_stats.json", user), stats)

	users, err := GetUsers(user)
	if err != nil {
		return
	}
	saveStructToFile(fmt.Sprintf("%s_user.json", user), users)

	for _, friend := range stats.Friends {
		snowballSampling(friend.Name, depth-1, visitedUsers)
	}
}
