package handlers

import (
	jwtservice "auth/services/jwtService"
	"encoding/json"
	"fmt"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	/*rfmap := map[string]string{}
	  err := json.NewDecoder(r.Body).Decode(&rfmap)
	  if err != nil {
	  	fmt.Println(err)
	  }
	  rfToken := rfmap[""]*/

	t, username, err := jwtservice.TokenValid(r)
	if err != nil {
		fmt.Println(err)
	}
	if !t {
		tokens, err := jwtservice.CreateTokenPair(username)
		if err != nil {
			fmt.Println(err)
		}
		jstokens, err := json.Marshal(&tokens)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jstokens)
		w.Write([]byte("this is a token "))
	} else {
		accessToken, err := jwtservice.CreateAccessToken(username)
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(accessToken))
		w.Write([]byte("this is a access token"))
	}

}
