package parameter

import (
	"encoding/json"
	"net/http"

	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/log"
)

// GetJSON ... get json data
func GetJSON(r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)
	if err != nil {
		ctx := r.Context()
		log.Warningm(ctx, "dec.Decode", err)
		return err
	}
	return nil
}
