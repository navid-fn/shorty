package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
		utils.WriteError(w, http.StatusInternalServerError, "error detected")
		return
	}
	createdUrl, err := ul.urlStore.CreateUrl(&url)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		utils.WriteError(w, http.StatusInternalServerError, "error detected")
		return
	}
	utils.WriteJson(w, http.StatusOK, createdUrl)
}

func (ul *UrlHandler) HandleRedirectUrl(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")

	// get original url
	originalUrl, err := ul.urlStore.GetOrginalUrlByString(shortCode)
	if err != nil {
		fmt.Println("Error Occurred..", err)
		utils.WriteError(w, http.StatusInternalServerError, "error detected")
		return
	}
	if originalUrl == nil {
		utils.WriteError(w, http.StatusNotFound, "Not Found")
		return
	}
	fullUrl := "https://" + *originalUrl
	http.Redirect(w, r, fullUrl, http.StatusFound)

}
