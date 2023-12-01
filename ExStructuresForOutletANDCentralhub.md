outlets := []Outlet{
    {
        outletID: "1",
        location: "Paris",
        inventory: map[string]*product.Product{
            "A": &product.Product{ /* ... Olive Oil details ... */ },
            "B": &product.Product{ /* ... Baguette details ... */ },
            "C": &product.Product{ /* ... Manchego Cheese details ... */ },
        },
    },
    {
        outletID: "2",
        location: "Lyon",
        inventory: map[string]*product.Product{
            "A": &product.Product{ /* ... Olive Oil details ... */ },
            "B": &product.Product{ /* ... Baguette details ... */ },
            "D": &product.Product{ /* ... Black Tea details ... */ },
        },
    },
    {
        outletID: "3",
        location: "Marseille",
        inventory: map[string]*product.Product{
            "A": &product.Product{ /* ... Olive Oil details ... */ },
            "C": &product.Product{ /* ... Manchego Cheese details ... */ },
            "D": &product.Product{ /* ... Black Tea details ... */ },
        },
    },
    {
        outletID: "4",
        location: "Nice",
        inventory: map[string]*product.Product{
            "B": &product.Product{ /* ... Baguette details ... */ },
            "C": &product.Product{ /* ... Manchego Cheese details ... */ },
            "D": &product.Product{ /* ... Black Tea details ... */ },
        },
    },
}

//省略了具体的产品信息


{
    hubID : "CentralHub",
    location : "Paris",
    resources : map[string]*product.Product{
            "A": &product.Product{ /* ... Olive Oil details ... */ },
            "B": &product.Product{ /* ... Baguette details ... */ },
            "C": &product.Product{ /* ... Manchego Cheese details ... */ },
            "D": &product.Product{ /* ... Black Tea details ... */ },
    }
}