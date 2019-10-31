package decta

import (
	"testing"
	"time"
)

var (
	productID = "60a6dffb-604c-4738-8733-1a2a6645177c"
	clientID  = "6bf85136-2f8f-4ec6-9d66-b438a05ccd93"
	orderID   = "e70e3f4a-86b8-4174-8421-2387806d3be8"
)

//
// Products
//

func TestProducts(t *testing.T) {
	products, err := decta.GetProducts()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, product := range products {
		productID = product.ID
	}

	t.Log(products)
}

//
// Clients
//

func TestNewClient(t *testing.T) {
	client := Client{
		Email:     "test@test.com",
		Phone:     "+123456789000",
		FirstName: "test-name",
		LastName:  "last-name",
	}

	if err := decta.NewClient(&client); err != nil {
		t.Error(err)
		t.FailNow()
	}

	clientID = client.ID

	t.Log(client.ID)
}

func TestGetClients(t *testing.T) {
	clients, err := decta.GetClients()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, client := range clients {
		clientID = client.ID
	}

	t.Log(clients)
}

func TestGetClient(t *testing.T) {
	client, err := decta.GetClient(clientID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if client.ID != clientID {
		t.Error("failed get client, ID not equal")
	}

	t.Log(client.ID)
}

//
// Orders
//

func TestNewOrder(t *testing.T) {
	order := Order{
		Language: "en",
		Due:      time.Now().Add(3 * time.Minute).Unix(),
		Client: &OrderClient{
			OriginalClientID: clientID,
			Email:            "test@test.com",
			SendToEmail:      true,
			SendToPhone:      false,
		},
		Products: []OrderProduct{
			OrderProduct{
				Title:    "test",
				ID:       productID,
				Price:    10,
				Quantity: 2,
			},
		},
	}

	err := decta.NewOrder(&order)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	orderID = order.ID

	t.Log(order.ID)
}

func TestGetOrder(t *testing.T) {
	order, err := decta.GetOrder(orderID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%+v", order)
}

func TestGetOrders(t *testing.T) {
	orders, err := decta.GetOrders()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, order := range orders {
		orderID = order.ID
	}

	t.Logf("%+v", orders)
}

func TestCancelOrder(t *testing.T) {
	if err := decta.CancelOrder(orderID); err != nil {
		t.Error(err)
	}
}

func TestDeleteOrder(t *testing.T) {
	if err := decta.DeleteOrder(orderID); err != nil {
		t.Error(err)
	}
}

func TestDeleteClient(t *testing.T) {
	if err := decta.DeleteClient(clientID); err != nil {
		t.Error(err)
	}
}
