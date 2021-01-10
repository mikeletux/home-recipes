package cors

import (
	"net/http"

	"github.com/rs/cors"
)

type CorsEnabler interface {
	Handler(http.Handler) http.Handler
}

func NewRsCorsEnabler(options *cors.Options) CorsEnabler {
	if options != nil {
		return cors.New(*options)
	}
	return cors.Default()
}
