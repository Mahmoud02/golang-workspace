package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

// Product

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

type Category struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
type BlogPost struct {
	Title   string
	Content string
}
type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	categories := []Category{
		{
			CategoryID:   "1NF",
			CategoryName: "Laptops",
		},
		{
			CategoryID:   "2NF",
			CategoryName: "PC",
		},
	}
	categoriesJSON, _ := json.Marshal(categories)

	w.Write(categoriesJSON)

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		j, err := json.Marshal(productList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		productList = append(productList, product)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		println("before category")
		handler.ServeHTTP(writer, request)
		println("after category")
	})
}
func blogPostHandler(w http.ResponseWriter, r *http.Request) {
	blogPost := BlogPost{Title: "First Article", Content: "bla bla bla bla"}
	template, _ := template.New("blog.tmpl").Parse(`<h1>{{.Title}}</h1><p>{{.Content}}</p>`)
	template.Execute(w, blogPost)
}

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vdS",
			"upc": "939581000000",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
		  },
		  {
			"productId": 2,
			"manufacturer": "Hessel, Schimmel and Feeney",
			"sku": "i7v300kmx",
			"upc": "740979000000",
			"pricePerUnit": "282.29",
			"quantityOnHand": 9217,
			"productName": "leg warmers"
		  },
		  {
			"productId": 3,
			"manufacturer": "Swaniawski, Bartoletti and Bruen",
			"sku": "q0L657ys7",
			"upc": "111730000000",
			"pricePerUnit": "436.26",
			"quantityOnHand": 5905,
			"productName": "lamp shade"
		  }
	]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	categoryHandler := http.HandlerFunc(CategoryHandler)
	http.Handle("/foo", &fooHandler{Message: "Hello Belal"})
	http.Handle("/category", middlewareHandler(categoryHandler))
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/blogs", blogPostHandler)
	http.ListenAndServe(":5000", nil)
}
