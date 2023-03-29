package dealer

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// returns modelId and tag and starts creating a new model
	router.HandlerFunc(http.MethodPost, "/v1/createModel", app.createModelHandler)
	// returns images for a model
	router.HandlerFunc(http.MethodPost, "/v1/getModels", app.getModelsHandler)
	// returns images for a model
	router.HandlerFunc(http.MethodPost, "/v1/getImages/", app.getImagesHandler)

	return router
}
