package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	logger "github.com/alxarno/swap/logger"

	db "github.com/alxarno/swap/db2"
	"github.com/alxarno/swap/settings"
	"github.com/robbert229/jwt"
)

const (
	successResult = "Success"
	errorResult   = "Error"
)

func decodeFail(ref string, err error, r *http.Request, w *http.ResponseWriter) {
	var p []byte
	r.Body.Read(p)
	sendAnswerError(ref, err, string(p), failedDecodeData, 0, w)
}

func getSecret() (string, error) {
	secret, err := settings.GetSettings()
	if err != nil {
		return "", err
	}
	return secret.Backend.SecretKeyForToken, nil
}

func sendAnswerError(reference string, err error, data string, eType int, errCode int, w *http.ResponseWriter) {
	e := fmt.Sprintf("%s %d", reference, errCode)
	if err != nil {
		e = fmt.Sprintf("%s %s", e, err.Error())
	}
	if data != "" {
		e = fmt.Sprintf("%s %s", e, data)
	}
	logger.Logger.Print(e)

	var answer = make(map[string]interface{})
	answer["result"] = errorResult
	answer["code"] = errCode
	answer["type"] = eType
	finish, _ := json.Marshal(answer)
	fmt.Fprintf((*w), string(finish))
}

func sendAnswerSuccess(w *http.ResponseWriter) {
	var x = make(map[string]string)
	x["result"] = successResult
	finish, _ := json.Marshal(x)
	fmt.Fprintf((*w), string(finish))
}

func generateToken(id int64) (string, error) {
	secret, err := getSecret()
	if err != nil {
		return "", err
	}
	algorithm := jwt.HmacSha256(secret)
	claims := jwt.NewClaim()
	claims.Set("id", id)
	claims.Set("time", time.Now().AddDate(0, 0, 30).Unix())
	token, err := algorithm.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func getJSON(target interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func TestUserToken(token string) (*db.User, error) {
	secret, err := getSecret()
	if err != nil {
		return nil, err
	}
	algorithm := jwt.HmacSha256(secret)
	claims, err := algorithm.Decode(token)
	if err != nil {
		return nil, errors.New("token is wrong 1")
	}
	id, err := claims.Get("id")
	if err != nil {
		return nil, errors.New("token is wrong 2")
	}
	tokenTime, err := claims.Get("time")
	if err != nil {
		return nil, errors.New("token is wrong 3")
	}

	if int64(tokenTime.(float64)) > time.Now().Unix() {
		u, err := db.GetUserByID(int64(id.(float64)))
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	return nil, errors.New("token time is done")
}

func getUserByToken(r *http.Request) (*db.User, error) {
	var token string
	if token = r.Header.Get("X-Auth-Token"); len(token) == 0 {
		return nil, errors.New("Token is undefined in X-Auth-Token header")
	}
	u, err := TestUserToken(token)
	if err != nil {
		return nil, err
	}
	return u, nil
}
