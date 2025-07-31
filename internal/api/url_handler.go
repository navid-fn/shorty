package api

import (

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/navid-fn/shorty/internal/store"
	"github.com/navid-fn/shorty/internal/utils"
)

type UrlHandler struct {
	urlStore store.UrlStore
}

func NewUrlHandler(urlStore store.UrlStore) *UrlHandler {
	return &UrlHandler{
		urlStore: urlStore,
	}
}

func (ul *UrlHandler) HandleCreateUrl(w http.ResponseWriter, r *http.Request) {
	var url store.Url
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		http.Error(w, "Error has been Occurred...", http.StatusInternalServerError)
		return
	}
	createdUrl, err := ul.urlStore.CreateUrl(&url)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		http.Error(w, "Error has been Occurred...", http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, createdUrl)
	return
}
