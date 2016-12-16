package talk_training

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	product_id := GetQuery(w, r)

	response := GetTalks(product_id)

	// response := Talks{
	// 	ID:         1,
	// 	ProductID:  product_id,
	// 	Message:    "test",
	// 	CreateTime: time.Now(),
	// }
	json.NewEncoder(w).Encode(response)
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var m Messages
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	m.ProductID = GetQuery(w, r)
	PostTalk(m)

	response := GetTalks(m.ProductID)

	json.NewEncoder(w).Encode(response)
}

func GetQuery(w http.ResponseWriter, r *http.Request) int64 {
	query := r.URL.Query()

	id_query := query.Get("product_id")
	product_id, err := strconv.ParseInt(id_query, 10, 64)
	if err != nil {
		panic(err)
	}

	return product_id
}
