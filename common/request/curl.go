package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Student struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	TotalEnergy  string `json:"totalEnergy"`
	RoundEnergy  string `json:"roundEnergy"`
	Rank         int    `json:"rank"`
	OnStageTimes int    `json:"onStageTimes"`
	AuthState    int    `json:"authState"`
}

type ClassStudent struct {
	UserList  []Student `json:"userList"`
	ClassName string    `json:"className"`
}

type Response struct {
	Stat int          `json:"stat"`
	Msg  string       `json:"msg"`
	Data ClassStudent `json:"data"`
}

func Post(ctx context.Context, url string, data interface{}) error {
	byteData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//fmt.Println(string(body))
	result := Response{}
	//result.Data.UserList = make([]Student, 0)
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	if len(result.Data.UserList) == 0 {
		return fmt.Errorf("%s", "No User")
	}
	fmt.Println("-", len(result.Data.UserList))
	return nil
}
