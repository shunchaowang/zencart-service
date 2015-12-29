package model

type (
    Product struct {
        Id string `json:"id"`
        Model string `json:"model"`
    }

    ProductType struct {
        Id string `json:"id"`
        Name string `json:"name"`
    }
)
