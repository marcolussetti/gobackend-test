//This is an experiment, a playground to discover whether Golang truly is

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcolussetti/gobackend-test/pkg/models"
)

// Articles represents the fake collection of blog entries
var Articles []Article

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
	fmt.Println("Endpoint GET: /")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint GET: /articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]
	fmt.Println("Endpoint GET: /article/" + articleID)

	for _, article := range Articles {
		if article.ID == articleID {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Endpoint POST: /article")

	var newArticle Article

	json.Unmarshal(reqBody, &newArticle)
	Articles = append(Articles, newArticle)

	json.NewEncoder(w).Encode(newArticle)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]
	fmt.Println("Endpoint DELETE: /article/" + articleID)

	for i, article := range Articles {
		if article.ID == articleID {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	articleID := vars["id"]
	fmt.Println("Endpoint POST: /article/" + articleID)

	var newArticle Article

	json.NewDecoder(r.Body).Decode(&newArticle)
	newArticle.ID = articleID

	for i, article := range Articles {
		if article.ID == articleID {
			var newArticles []Article
			newArticles = append(Articles[:i], newArticle)
			newArticles = append(newArticles, Articles[i+1:]...)
			Articles = newArticles

			json.NewEncoder(w).Encode(newArticle)
		}
	}

}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/article", returnAllArticles).Methods("GET")
	router.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
