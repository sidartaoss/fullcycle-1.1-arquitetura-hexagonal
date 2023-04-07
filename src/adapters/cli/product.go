package cli

import (
	"fmt"
	"github.com/codeedu/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s with the name %s has been created with price: %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
		break
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		enabledProduct, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", enabledProduct.GetName())
		break
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		disabledProduct, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", disabledProduct.GetName())
		break
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return result, nil
}
