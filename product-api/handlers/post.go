package handlers

import (
	"net/http"

	"github.com/huavanthong/microservice-golang/product-api/data"
)

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
