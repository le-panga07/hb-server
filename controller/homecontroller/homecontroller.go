package homecontroller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"hb-server/models"
	"hb-server/services"
	"html/template"
	"io/ioutil"
	"net/http"

	_ "hb-server/github.com/go-sql-driver/mysql"
	"hb-server/github.com/gorilla/mux"
)

//DataRes Result fromDb
type DataRes struct {
	Name           string `json:"name"`
	IsActive       bool   `json:"isactive"`
	ID             string `json:"publisherid"`
	Adslotid       string `json:"adslotid"`
	Extpublisherid string `json:"extpublisherid"`
	ProviderName   string `json:"providername"`
}

//Index function
func Index(db *sql.DB) http.HandlerFunc {

	fn := func(res http.ResponseWriter, req *http.Request) {
		tpl := template.Must(template.ParseFiles("../hb-server/views/test.html"))
		tpl.ExecuteTemplate(res, "test", nil)
	}
	return http.HandlerFunc(fn)
}

//GetConfigMap get config maps
func GetConfigMap(db *sql.DB) http.HandlerFunc {

	fn := func(res http.ResponseWriter, req *http.Request) {

		publisherID := mux.Vars(req)["id"]
		fmt.Println(publisherID)

		config := services.GetProviderConfigs(db, publisherID)

		//auctionResult := SendDataToNodeServer(config)
		GetScriptFileFromNodeServer(res, config)

		//res.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(res).Encode(auctionResult)

	}
	return http.HandlerFunc(fn)
}

//GetScriptFileFromNodeServer func
func GetScriptFileFromNodeServer(res http.ResponseWriter, configs *models.Config) {
	url := "http://localhost:3000/configs"
	fmt.Println("URL:> s", url)

	configsBuffer := new(bytes.Buffer)
	json.NewEncoder(configsBuffer).Encode(configs)

	requ, _ := http.NewRequest("POST", url, configsBuffer)

	requ.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, e := client.Do(requ)
	if e != nil {
		panic(e)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "text/javascript")
	res.Write([]byte(body))

}

//SendDataToNodeServer func
func SendDataToNodeServer(configs *models.Config) models.AuctionResult {
	url := "http://localhost:3000/configs"
	fmt.Println("URL:> s", url)

	configsBuffer := new(bytes.Buffer)
	json.NewEncoder(configsBuffer).Encode(configs)

	requ, _ := http.NewRequest("POST", url, configsBuffer)

	requ.Header.Set("X-Custom-Header", "goserver")
	requ.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, e := client.Do(requ)
	if e != nil {
		panic(e)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	auctionResult := make(models.AuctionResult)
	err = json.Unmarshal(body, &auctionResult)

	return auctionResult

}
