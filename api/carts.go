package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	username := fmt.Sprintf("%s", r.Context().Value("username")) // TODO: replace this

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		error := model.ErrorResponse{
			Error: "Request Product Not Found",
		}
		jsonError, err := json.Marshal(error)
		if err != nil {
			return
		}
		w.Write(jsonError)
		return
	}

	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	// TODO: answer here

	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	var totalPrice float64

	for _, j := range list {
		totalPrice += j.Total
	}

	if len(list) == 0 {
		w.WriteHeader(400)
		error := model.ErrorResponse{
			Error: "Request Product Not Found",
		}
		jsonError, err := json.Marshal(error)
		if err != nil {
			return
		}
		w.Write(jsonError)
		return
	}

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{
		Name:       username,
		Cart:       list,
		TotalPrice: totalPrice,
	} // TODO: replace this

	_, err = api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	w.WriteHeader(200)
	api.dashboardView(w, r)
}
