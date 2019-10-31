package decta

//
// Products
//

func (d *Decta) GetProducts() (products []Product, err error) {
	err = d.GET("/products/", &products)
	return
}

//
// Clients
//

func (d *Decta) GetClients() (clients []Client, err error) {
	err = d.GET("/clients/", &clients)
	return
}

func (d *Decta) GetClient(id string) (client Client, err error) {
	err = d.GET("/clients/"+id, &client)
	return
}

func (d *Decta) NewClient(c *Client) error {
	return d.POST("/clients/", c, c)
}

func (d *Decta) DeleteClient(id string) error {
	return d.DELETE("/clients/"+id, nil)
}

//
// Orders
//

func (d *Decta) GetOrders() (orders []Order, err error) {
	err = d.GET("/orders/", &orders)
	return
}

func (d *Decta) GetOrder(id string) (order Order, err error) {
	err = d.GET("/orders/"+id, &order)
	return
}

func (d *Decta) NewOrder(order *Order) error {
	return d.POST("/orders/", order, &order)
}

func (d *Decta) DeleteOrder(id string) error {
	return d.DELETE("/orders/"+id, nil)
}

func (d *Decta) CancelOrder(id string) error {
	return d.POST("/orders/"+id+"/cancel/", nil, nil)
}
