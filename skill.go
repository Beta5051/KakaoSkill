package kakaoskill

import (
	"encoding/json"
	"net/http"
)

const VERSION = "2.0"

func AddSkill(mux *http.ServeMux, path string, handler func (payload Payload) *Response) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload := Payload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := handler(payload)
		if resp == nil {
			http.Error(w, "response is null", http.StatusBadRequest)
			return
		}
		if resp.Version == "" {
			resp.Version = VERSION
		}
		bs, err := resp.Encode()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(bs)
	})
}