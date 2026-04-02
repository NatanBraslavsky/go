package main

import (
	"log"

	"resty.dev/v3"
)

type Product struct {
	Name string         `json:"name"`
	Data map[string]any `json:"data"`
}

type Response struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Data      map[string]any `json:"data"`
	CreatedAt any            `json:"createdAt"`
	UpdatedAt any            `json:"updatedAt"`
}

func main() {
	client := resty.New()

	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// GET ALL
	res, err := client.R().
		Get("https://api.restful-api.dev/objects")
	if err != nil {
		log.Println("Erro GET ALL:", err)

		return
	}

	log.Println("GET ALL Status:", res.Status())

	// GET BY ID
	res, err = client.R().
		Get("https://api.restful-api.dev/objects/3")
	if err != nil {
		log.Println("Erro GET BY ID:", err)

		return
	}

	log.Println("GET BY ID Status:", res.Status())

	// POST
	body := Product{
		Name: "Apple MacBook",
		Data: map[string]any{
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
		log.Println("Erro POST:", err)

		return
	}

	log.Println("POST Status:", res.Status())

	created, ok := res.Result().(*Response)
	if !ok {
		log.Println("Erro ao converter response")

		return
	}

	id := created.ID
	log.Println("ID criado:", id)

	// PUT
	body = Product{
		Name: "Apple MacBook Pro 16",
		Data: map[string]any{
			"year":           2019,
			"price":          2049.99,
			"CPU model":      "Intel Core i9",
			"Hard disk size": "1 TB",
			"color":          "silver",
		},
	}

	url := "https://api.restful-api.dev/objects/" + id

	res, err = client.R().
		SetBody(body).
		SetResult(&Response{}).
		Put(url)
	if err != nil {
		log.Println("Erro PUT:", err)

		return
	}

	log.Println("PUT Status:", res.Status())

	updated, ok := res.Result().(*Response)
	if !ok {
		log.Println("Erro ao converter response (PUT)")

		return
	}

	log.Println("Atualizado (PUT):", updated)

	// PATCH
	patchBody := map[string]any{
		"name": "Apple MacBook Pro 16 (PATCHED)",
	}

	res, err = client.R().
		SetBody(patchBody).
		SetResult(&Response{}).
		Patch(url)
	if err != nil {
		log.Println("Erro PATCH:", err)

		return
	}

	log.Println("PATCH Status:", res.Status())

	patched, ok := res.Result().(*Response)
	if !ok {
		log.Println("Erro ao converter response (PATCH)")

		return
	}

	log.Println("Atualizado (PATCH):", patched)

	// DELETE
	res, err = client.R().
		Delete(url)
	if err != nil {
		log.Println("Erro DELETE:", err)

		return
	}

	log.Println("DELETE Status:", res.Status())
	log.Println("DELETE Body:", res.String())
}
