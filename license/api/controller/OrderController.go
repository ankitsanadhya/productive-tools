package controller

import (
	"fmt"
	"license/api/apihelper"
	"license/api/constant"
	"license/api/vo"

	"github.com/gin-gonic/gin"
)

//OrderController struct
type OrderController struct {
}

func (oc *OrderController) PlaceOrder(c *gin.Context) {
	fmt.Println("Place order")

	/*//create sub product
	subProduct := vo.SubProduct{
		SubProductID:  "subProdID",
		Name:          "subProd",
		Description:   "Desc",
		IsDefault:     false,
		DefaultPrice:  20.2,
		Specification: "{a:2,b:3}",
	}

	var subProducts []vo.SubProduct
	subProducts = append(subProducts, subProduct)

	//create product
	product := vo.Product{
		ProductID:     "ProdId1",
		Name:          "prod",
		Description:   "desc",
		IsDefault:     true,
		DefaultPrice:  100.0,
		Specification: "fhg fyyg gy",
		SubProducts:   subProducts,
	}

	// create distribution product
	distributionProduct := vo.DistributionProduct{
		DistributionProductID: "distprodId1",
		Product:               product,
		Price:                 99,
	}

	//
	var distributionProducts []vo.DistributionProduct
	distributionProducts = append(distributionProducts, distributionProduct)

	// create distribution

	distribution := vo.Distribution{
		DistributionID:      "distId",
		Name:                "dist1",
		DistributionProduct: distributionProducts,
		Description:         "dsec gdfgh",
		LicenseType:         "paid",
		Validity:            "30",
		MaxActivation:       1,
		ActivationPerDevice: 1,
	}

	//create payment method

	payments := vo.PaymentMethod{
		PaymentID: "paymentId",
		Name:      "paymentNAme",
		Detail:    "details",
	}
	//cretae OrderProduct
	orderProduct := vo.OrderProduct{
		OrderProductID:        "ordrProdID",
		DistributionProductVo: distributionProduct,
	}
	var orderProducts []vo.OrderProduct
	orderProducts = append(orderProducts, orderProduct)

	//create order
	order := vo.Order{
		OrderID:        "ordreID1",
		Price:          20.2,
		TotalPrice:     20.2,
		CouponID:       "couponId1",
		TransactionID:  "tansactionId1",
		LicenceCount:   2,
		DistributionVo: distribution,
		PaymentVo:      payments,
		OrderProducts:  orderProducts,
	}*/

	result := make(map[string]interface{})
	result[constant.DATA()] = getOrder("")
	apihelper.CreateResponse(c, result)
}

func getOrder(orderId string) vo.Order {
	//create sub product
	subProduct := vo.SubProduct{
		SubProductID:  "subProdID",
		Name:          "subProd",
		Description:   "Desc",
		IsDefault:     false,
		DefaultPrice:  20.2,
		Specification: "{a:2,b:3}",
	}

	var subProducts []vo.SubProduct
	subProducts = append(subProducts, subProduct)

	//create product
	product := vo.Product{
		ProductID:     "ProdId1",
		Name:          "prod",
		Description:   "desc",
		IsDefault:     true,
		DefaultPrice:  100.0,
		Specification: "fhg fyyg gy",
		SubProducts:   subProducts,
	}

	// create distribution product
	distributionProduct := vo.DistributionProduct{
		DistributionProductID: "distprodId1",
		Product:               product,
		Price:                 99,
	}

	//
	var distributionProducts []vo.DistributionProduct
	distributionProducts = append(distributionProducts, distributionProduct)

	// create distribution

	distribution := vo.Distribution{
		DistributionID:      "distId",
		Name:                "dist1",
		DistributionProduct: distributionProducts,
		Description:         "dsec gdfgh",
		LicenseType:         "paid",
		Validity:            "30",
		MaxActivation:       1,
		ActivationPerDevice: 1,
	}

	//create payment method

	payments := vo.PaymentMethod{
		PaymentID: "paymentId",
		Name:      "paymentNAme",
		Detail:    "details",
	}
	//cretae OrderProduct
	orderProduct := vo.OrderProduct{
		OrderProductID:        "ordrProdID",
		DistributionProductVo: distributionProduct,
	}
	var orderProducts []vo.OrderProduct
	orderProducts = append(orderProducts, orderProduct)

	//create order
	order := vo.Order{
		OrderID:        "ordreID1",
		Price:          20.2,
		TotalPrice:     20.2,
		CouponID:       "couponId1",
		TransactionID:  "tansactionId1",
		LicenceCount:   2,
		DistributionVo: distribution,
		PaymentVo:      payments,
		OrderProducts:  orderProducts,
	}
	return order
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	fmt.Println("Get order")
	var orderId = c.Param("id")

	result := make(map[string]interface{})
	result[constant.DATA()] = getOrder(orderId)
	apihelper.CreateResponse(c, result)
}
