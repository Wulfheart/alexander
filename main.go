package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wulfheart/godip-influence/defaultInfluences"
	"github.com/zond/godip/variants"
	"net/http"
	"wulfheartalexander/common"
)

func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func preflight(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
}
func resolve(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	variantName := mux.Vars(r)["variant"]
	v, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	p := common.RequestDTO{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s := p.State(v)
	if err := s.Next(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := json.NewEncoder(w).Encode(common.NewResponseDTOfromState(s, p.Influence, v)); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	var variantName = mux.Vars(r)["variant"]
	v, found := variants.Variants[variantName]
	if !found {
		http.Error(w, fmt.Sprintf("Variant %q not found", variantName), 404)
		return
	}
	s, err := v.Start()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = json.NewEncoder(w).Encode(common.NewResponseDTOfromState(s, defaultInfluences.ConvertToInfluence(defaultInfluences.Classical), v)); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func main() {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(preflight)
	variants := r.Path("/{variant}").Subrouter()
	variants.Methods("POST").HandlerFunc(resolve)
	variants.Methods("GET").HandlerFunc(start)
	// r.Path("/").HandlerFunc(listVariants)
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8000", r)
}
