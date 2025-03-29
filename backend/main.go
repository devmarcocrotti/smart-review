package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jonathanhecl/gollama"
	"github.com/rs/cors"
)

type Response struct {
	Data string `json:"data"`
}

type Review struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	Text   string `json:"text"`
}

var Reviews = []Review{
	{"Cristina", 5, "Questo set di action figures è incredibile! La qualità dei dettagli è sorprendente, ogni personaggio sembra proprio uscito dal film. I colori sono vibranti e le articolazioni funzionano perfettamente. Un must per ogni vero collezionista!"},
	{"Ivan", 4, "La statuetta è davvero bella e ben fatta, ma il packaging potrebbe essere migliorato. Quando è arrivata, la scatola era leggermente danneggiata, anche se il prodotto interno era intatto. Nel complesso, un'ottima aggiunta alla mia collezione."},
	{"Marco", 2, "Deluso da questa edizione limitata. Sebbene l'idea fosse interessante, la qualità dei materiali non è all'altezza delle aspettative. Il pezzo si è graffiato facilmente e non sembra essere abbastanza robusto. Non consiglio a chi cerca un prodotto durevole."},
	{"Rosy", 5, "Una delle migliori acquisizioni della mia collezione! La replica del film è perfetta, il design è autentico e la rifinitura è impeccabile. Vale assolutamente il prezzo, e mi fa sentire come se avessi un pezzo di storia del cinema nella mia casa."},
	{"Alessio", 3, "Il prodotto è interessante, ma non sono completamente soddisfatto. La qualità non è male, ma alcune parti erano difficili da montare e una delle figure era leggermente sbilanciata. Non male, ma non all'altezza di altre edizioni da collezione che ho visto."},
}

var G *gollama.Gollama
var Ctx context.Context

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reviewsJSON, err := json.MarshalIndent(Reviews, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(reviewsJSON)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	prompt := "Fai riassunto, in una sola frase, di queste recensioni: "
	var reviewTexts []string

	for _, review := range Reviews {
		reviewTexts = append(reviewTexts, review.Text)
	}

	prompt += strings.Join(reviewTexts, ", ")

	type AIResponse struct {
		AIResponse string `required:"true"`
	}

	option := gollama.StructToStructuredFormat(AIResponse{})

	fmt.Printf("Option: %+v\n", option)

	output, err := G.Chat(Ctx, prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := Response{
		Data: output.Content,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func replyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(Reviews) {
		http.Error(w, "Indice non valido", http.StatusBadRequest)
		return
	}

	review := Reviews[index]

	prompt := "Rispondi con tono professionale a questa recensione: "
	prompt += review.Text

	type AIResponse struct {
		AIResponse string `required:"true"`
	}

	option := gollama.StructToStructuredFormat(AIResponse{})

	fmt.Printf("Option: %+v\n", option)

	output, err := G.Chat(Ctx, prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := Response{
		Data: output.Content,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	Ctx = context.Background()
	G = gollama.New("llama3.2")
	G.ServerAddr = "http://ollama:11434"
	G.Verbose = false
	if err := G.PullIfMissing(Ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}

	r.HandleFunc("/list", listHandler).Methods("GET")
	r.HandleFunc("/query", queryHandler).Methods("GET")
	r.HandleFunc("/reply/{id}", replyHandler).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	fmt.Println("Server avviato http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(r)))
}
