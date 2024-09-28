package tests

import (
	"api/dao"
	"api/db"
	"api/repository"
	"api/router"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testProductIDs []string
var engine *gin.Engine

func TestMain(m *testing.M) {

	db.InitInMemoryDB()
	engine = router.InitRouter()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func setup() {
	// Add initial items to the database
	addInitialItems()
}

func teardown() {

	// Clean up database
	log.Printf("Flushing database...")
	db.GetDatabase.Exec("DELETE FROM products")
	testProductIDs = []string{}
}

func addInitialItems() {
	log.Print("Adding initial items to the database...")
	// Add initial items to the database
	products := []dao.Product{
		{Name: "Test Product 1", Description: "Initial Description 1", SKU: "SKU1", Image: "http://example.com/image1.jpg", Price: 10.00, Stock: 100},
		{Name: "Test Product 2", Description: "Initial Description 2", SKU: "SKU2", Image: "http://example.com/image2.jpg", Price: 20.00, Stock: 200},
		{Name: "Test Product 3", Description: "Initial Description 3", SKU: "SKU3", Image: "http://example.com/image3.jpg", Price: 30.00, Stock: 300},
	}

	for _, product := range products {

		prod, err := repository.CreateProduct(&product)
		if err != nil {
			log.Fatalf("Failed to add product: %v", err)
		}

		testProductIDs = append(testProductIDs, prod.ID.String())
		log.Printf("Added product with ID: %s", prod.ID)

	}
	log.Print("Adding initial items to the database Finished...")
}

// TestGetProduct tests the GetProduct function
func TestGetProduct(t *testing.T) {
	setup()
	defer teardown()
	log.Printf("Testing GetProduct")

	w := httptest.NewRecorder()

	// Set up the expected query for the specific product ID

	req, _ := http.NewRequest("GET", "/api/v1/products/"+testProductIDs[0], nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "product")
	assert.Equal(t, "Test Product 1", response["product"].(map[string]interface{})["name"])

}

// TestDeleteProduct tests the DeleteProduct function
func TestDeleteProduct(t *testing.T) {
	setup()
	defer teardown()
	log.Printf("Testing DeleteProduct")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/products/delete/"+testProductIDs[0], nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "product deleted successfully", response["message"])

}

// TestPostProduct tests the PostProduct function
func TestPostProduct(t *testing.T) {
	defer teardown()
	log.Printf("Testing PostProduct")

	var requestBody struct {
		Products []dao.Product `json:"products"`
	}
	requestBody.Products = append(requestBody.Products, dao.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		SKU:         "TESTSKU123",
		Image:       "http://example.com/image.jpg",
		Price:       99.99,
		Stock:       100,
	})

	productJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal dummy product: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products/create", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	engine.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "Product created Successfully", response["message"])

}

// TestPostProduct tests the PostProduct function when a mandatory product field is missing
func TestPostProductWithoutDependencies(t *testing.T) {
	defer teardown()
	log.Printf("Testing PostProduct")

	var requestBody dao.ProductsList
	requestBody.Products = append(requestBody.Products, dao.Product{
		Description: "This is a test product",
		SKU:         "TESTSKU123",
		Image:       "http://example.com/image.jpg",
		Price:       99.99,
		Stock:       100,
	})

	productJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal dummy product: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products/create", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	engine.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "error")

}

// TestDeleteProducts tests the DeleteProducts function
func TestDeleteProducts(t *testing.T) {
	defer teardown()
	log.Printf("Testing DeleteProducts")
	var ids []string
	// Add products to delete
	products := []dao.Product{
		{Name: "Test Product 1", Description: "To be deleted 1", SKU: "DEL1", Image: "http://example.com/del1.jpg", Price: 10.00, Stock: 100},
		{Name: "Test Product 2", Description: "To be deleted 2", SKU: "DEL2", Image: "http://example.com/del2.jpg", Price: 20.00, Stock: 200},
	}
	for _, product := range products {
		prod, err := repository.CreateProduct(&product)
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}
		ids = append(ids, prod.ID.String())

	}

	idsJSON, err := json.Marshal(ids)
	if err != nil {
		t.Fatalf("Failed to marshal product IDs: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/products/bulk/delete", bytes.NewBuffer(idsJSON))
	req.Header.Set("Content-Type", "application/json")
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "products deleted successfully", response["message"])

	// Verify all products are deleted
	var remainingProducts []dao.Product
	result := db.GetDatabase.Find(&remainingProducts)
	assert.NoError(t, result.Error)
	assert.Empty(t, remainingProducts)
	log.Printf("Remaining products: %v", remainingProducts)
}

// TestGetProducts tests the GetProducts function
func TestGetProducts(t *testing.T) {
	setup()
	defer teardown()
	log.Printf("Testing GetProducts")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products?limit=10&page=1", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "products")
	assert.IsType(t, []interface{}{}, response["products"])
}

// TestPutProduct tests the PutProduct function
func TestPutProduct(t *testing.T) {
	setup()
	defer teardown()
	log.Printf("Testing PutProduct")

	// Create a dummy product to update
	var requestBody struct {
		Product dao.UpdateProduct `json:"product"`
	}
	requestBody.Product = dao.UpdateProduct{
		Name:        "Updated Test Product",
		Description: "This is an updated test product",
		SKU:         "UPDATEDSKU123",
		Image:       "http://example.com/updated_image.jpg",
		Price:       199.99,
		Stock:       50,
	}

	productJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal dummy product: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/products/update/"+testProductIDs[0], bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "product with id "+testProductIDs[0]+" updated successfully", response["message"])

}
