package handler

import (
	"Pengenalan-Microservice/config"
	"Pengenalan-Microservice/utils"
	"net/http"
)

type AuthHandler struct {
	Config config.Auth
}

// Menjalankan validasi terlebih dahulu -> baru next handler
func (handler *AuthHandler) ValidateAdmin(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest(http.MethodPost, handler.Config.Host+"/validate-admin", nil)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header

		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		/*responseBody, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}*/

		/*var responseData map[string]interface{}
		err = json.Unmarshal(responseBody, &responseBody)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}*/

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, "invalid auth", authResponse.StatusCode)
			return
		}
	}
}
