package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net/http"
	"protobuf/proto/user"
)

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		res := &LoginResponse{}

		if r.Method == "POST" {
			err := json.NewDecoder(r.Body).Decode(&res)
			if err != nil {
				w.Write([]uint8(err.Error()))
			}
			fmt.Println("res", res)
		}

		userJSONProto, err := json.Marshal(user.User{})
		if err != nil {
			fmt.Println(err)
		}

		userProto, err := proto.Marshal(&user.User{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(userJSONProto))
		fmt.Println(string(userProto))

		resp := &user.User{
			Name:   res.Username,
			Age:    10,
			Gender: 10,
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(resp)
		w.Write(jsonData)
	})

	http.ListenAndServe(":8080", nil)
}
