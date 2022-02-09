package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// HTTP-method routing along `pattern`
// Connect(pattern string, h http.HandlerFunc)
// Delete(pattern string, h http.HandlerFunc)
// Get(pattern string, h http.HandlerFunc)
// Head(pattern string, h http.HandlerFunc)
// Options(pattern string, h http.HandlerFunc)
// Patch(pattern string, h http.HandlerFunc)
// Post(pattern string, h http.HandlerFunc)
// Put(pattern string, h http.HandlerFunc)
// Trace(pattern string, h http.HandlerFunc)

//function to get the Article
// func getArticle(w  http.ResponseWriter  ,  r *http.Request){
  
// 	dataParam := chi.URLParam(r,"date")

// 	slugParam := chi.URLParam(r , "slug")

// 	article ,  err := database.getArticle(date, slug)

// 	if  err != nil {
// 		w.WriteHeader(422)
// 		w.WriteHeader([]byte(fmt.Sprintf("error fetching article %s=%s: &v" ,dataParam , slugParam,err)))
// 		return 
// 	}

// 	if article == nil {
// 		w.WriteHeader(404)
// 		w.Write([]byte("article not found"))
// 		return 

// 	}
// 	w.Write([]byte(article.Text))
// }



func main(){

  r := chi.NewRouter();//create new route
  
  r.Use(middleware.RequestID)
  r.Use(middleware.Logger) //logging  middleware
  r.Use(middleware.Recoverer)

  r.Get("/" , func(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("Testing the first route logging "))
  })

   http.ListenAndServe(":3000" , r)
}