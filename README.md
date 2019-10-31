# Decta API client

API Client for [Decta](https://decta.com/)

### Install

```sh
go get github.com/sg3des/decta
```

### Usage


```go
d, err := decta.NewDecta(secret)
if err != nil {...}


products, err := d.GetProducts()
// handle error

// create new client
client := Client{
	Email:     "test@test.com",
	Phone:     "+123456789000",
	FirstName: "test-name",
	LastName:  "last-name",
}

err = d.NewClient(&client);
// handle error


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

err = decta.NewOrder(&order)
// handle error
```
