package main

import (
	"fmt"

	"resty.dev/v3"
)

type Product struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Response struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt interface{}            `json:"createdAt"`
	UpdatedAt interface{}            `json:"updatedAt"`
}

func main() {
	client := resty.New()
	defer client.Close()

	// GET ALL
	res, err := client.R().
		Get("https://api.restful-api.dev/objects")

	if err != nil {
		fmt.Println("Erro GET ALL:", err)
		return
	}

	fmt.Println("GET ALL Status:", res.Status())

	// GET BY ID
	res, err = client.R().
		Get("https://api.restful-api.dev/objects/3")

	if err != nil {
		fmt.Println("Erro GET BY ID:", err)
		return
	}

	fmt.Println("GET BY ID Status:", res.Status())

	// POST
	body := Product{
		Name: "Apple MacBook",
		Data: map[string]interface{}{
			"year":           2019,
			"price":          1849.99,
			"CPU model":      "Intel Core i9",
			"Hard disk size": "1 TB",
		},
	}

	res, err = client.R().
		SetBody(body).
		SetResult(&Response{}).
		Post("https://api.restful-api.dev/objects")

	if err != nil {
		fmt.Println("Erro POST:", err)
		return
	}

	fmt.Println("POST Status:", res.Status())

	created := res.Result().(*Response)
	fmt.Println("Criado:", created)

	id := created.ID
	fmt.Println("ID criado:", id)

	// PUT
	body = Product{
		Name: "Apple MacBook Pro 16",
		Data: map[string]interface{}{
			"year":           2019,
			"price":          2049.99,
			"CPU model":      "Intel Core i9",
			"Hard disk size": "1 TB",
			"color":          "silver",
		},
	}

	url := fmt.Sprintf("https://api.restful-api.dev/objects/%s", id)

	res, err = client.R().
		SetBody(body).
		SetResult(&Response{}).
		Put(url)

	if err != nil {
		fmt.Println("Erro PUT:", err)
		return
	}

	fmt.Println("PUT Status:", res.Status())

	updated := res.Result().(*Response)
	fmt.Println("Atualizado (PUT):", updated)

	// PATCH
	patchBody := map[string]interface{}{
		"name": "Apple MacBook Pro 16 (PATCHED)",
	}

	res, err = client.R().
		SetBody(patchBody).
		SetResult(&Response{}).
		Patch(url)

	if err != nil {
		fmt.Println("Erro PATCH:", err)
		return
	}

	fmt.Println("PATCH Status:", res.Status())

	patched := res.Result().(*Response)
	fmt.Println("Atualizado (PATCH):", patched)

	// DELETE 
	res, err = client.R().
		Delete(url)

	if err != nil {
		fmt.Println("Erro DELETE:", err)
		return
	}

	fmt.Println("DELETE Status:", res.Status())
	fmt.Println("DELETE Body:", res.String())
}
