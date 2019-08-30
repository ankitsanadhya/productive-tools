package controller

import (
	"encoding/json"
	"license/api/apihelper"
	"license/api/constant"
	"license/api/service"
	"license/api/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PartnerController struct {
}

func (pc *PartnerController) ListPartner(c *gin.Context) {
	// fmt.Println("Test Partner Controller")

	// listPartner := []vo.Partner{}
	// partner := vo.Partner{}

	// partner.PartnerID = "PatnerID1"
	// partner.Name = "Patner 1"
	// listPartnerUser := []vo.PartnerUser{}
	// partnerUser := vo.PartnerUser{}

	// partnerUser.UpID = "UserpartnerID 1"
	// partnerUser.PartnerID = "PartnerID 1"
	// partnerUser.AuthUserID = "AuthUserID 1"
	// partnerUser.UserRole = "admin"
	// partnerUser.Status = "Enable"

	// listPartnerUser = append(listPartnerUser, partnerUser)
	// partner.Users = listPartnerUser

	// partner.Description = "Patner Description"

	// listTenant := []vo.Tenant{}
	// tenant := vo.Tenant{}

	// tenant.TenantID = "TenantID1"
	// tenant.TenantName = "TenantName 1"
	// tenant.Description = "Description 1"

	// listTenantProduct := []vo.Product{}
	// tenantproduct := vo.Product{}
	// tenantproduct.ProductID = "ProductID 1"
	// tenantproduct.Name = "productName 1"
	// tenantproduct.Description = "productDescription1"
	// tenantproduct.IsDefault = true
	// tenantproduct.DefaultPrice = 1248
	// tenantproduct.Specification = ""

	// tenantsubProducts := vo.SubProduct{}
	// tenantsubProducts.SubProductID = "SubProductID1"
	// tenantsubProducts.Name = "productName1"
	// tenantsubProducts.Description = "subProducts Description "
	// tenantsubProducts.IsDefault = false
	// tenantsubProducts.DefaultPrice = 1248
	// tenantsubProducts.Specification = "Specification fggf"

	// listtenantSubProducts := []vo.SubProduct{}
	// listtenantSubProducts = append(listtenantSubProducts, tenantsubProducts)
	// tenantproduct.SubProducts = listtenantSubProducts

	// listTenantProduct = append(listTenantProduct, tenantproduct)
	// tenant.Products = listTenantProduct

	// listDistribution := []vo.Distribution{}

	// subProduct1 := vo.SubProduct{
	// 	SubProductID:  "subProdID",
	// 	Name:          "subProd",
	// 	Description:   "Desc",
	// 	IsDefault:     false,
	// 	DefaultPrice:  20.2,
	// 	Specification: "{a:2,b:3}",
	// }

	// subProducts1 := []vo.SubProduct{}
	// subProducts1 = append(subProducts1, subProduct1)

	// product1 := vo.Product{
	// 	ProductID:     "ProdId1",
	// 	Name:          "prod",
	// 	Description:   "desc",
	// 	IsDefault:     true,
	// 	DefaultPrice:  100.0,
	// 	Specification: "fhg fyyg gy",
	// 	SubProducts:   subProducts1,
	// }
	// distributionProduct := vo.DistributionProduct{
	// 	DistributionProductID: "distprodId1",
	// 	Product:               product1,
	// 	Price:                 99,
	// }
	// distributionProducts := []vo.DistributionProduct{}
	// distributionProducts = append(distributionProducts, distributionProduct)

	// tenant.Distributions = listDistribution

	// listtenantPartnerUser := []vo.PartnerUser{}
	// tenant.Users = listtenantPartnerUser

	// listTenantPaymentMethod := []vo.PaymentMethod{}

	// tenantpaymentMethod := vo.PaymentMethod{}
	// tenantpaymentMethod.PaymentID = "PaymentID"
	// tenantpaymentMethod.Name = "paymentMethodName"
	// tenantpaymentMethod.Detail = "paymentMethod Detail"

	// listTenantPaymentMethod = append(listTenantPaymentMethod, tenantpaymentMethod)
	// tenant.PaymentMethods = listTenantPaymentMethod

	// listTenant = append(listTenant, tenant)

	// partner.Tenants = listTenant

	// listProduct := []vo.Product{}

	// product := vo.Product{}
	// product.ProductID = "ProductID 1"
	// product.Name = "productName 1"
	// product.Description = "productDescription1"
	// product.IsDefault = true
	// product.DefaultPrice = 1248
	// product.Specification = ""

	// subProducts := vo.SubProduct{}
	// subProducts.SubProductID = "SubProductID1"
	// subProducts.Name = "productName1"
	// subProducts.Description = "subProducts Description "
	// subProducts.IsDefault = false
	// subProducts.DefaultPrice = 1248
	// subProducts.Specification = "{a:2,b:3}"

	// listSubProducts := []vo.SubProduct{}
	// listSubProducts = append(listSubProducts, subProducts)
	// product.SubProducts = listSubProducts

	// listProduct = append(listProduct, product)
	// partner.Product = listProduct

	// partner.Currency = "USD"

	// listPaymentMethod := []vo.PaymentMethod{}
	// paymentMethod := vo.PaymentMethod{}
	// paymentMethod.PaymentID = "PaymentID"
	// paymentMethod.Name = "paymentMethodName"
	// paymentMethod.Detail = "paymentMethod Detail"

	// listPaymentMethod = append(listPaymentMethod, paymentMethod)
	// partner.PaymentMethod = listPaymentMethod

	// listPartner = append(listPartner, partner)

	ps := new(service.PartnerService)
	var listPartner *[]vo.Partner
	err := ps.ListAllPartner(listPartner)
	if err != nil {
		apihelper.CreateResponseWithErrExplanation(c, err.Error(), err.Error(), http.StatusUnauthorized)
		return
	}
	result := make(map[string]interface{})
	result[constant.DATA()] = listPartner
	apihelper.CreateResponse(c, result)
}

func (pc *PartnerController) AddPartner(c *gin.Context) {

	partner := new(vo.Partner)
	err := json.NewDecoder(c.Request.Body).Decode(&partner)
	if err != nil {
		apihelper.CreateResponseWithErrExplanation(c, err.Error(), err.Error(), http.StatusBadRequest)
		return
	}
	ps := new(service.PartnerService)
	err = ps.AddPartner(partner)
	if err != nil {
		apihelper.CreateResponseWithErrExplanation(c, err.Error(), err.Error(), http.StatusUnauthorized)
		return
	}

	result := make(map[string]interface{})
	result[constant.DATA()] = partner
	apihelper.CreateResponse(c, result)

}

func (pc *PartnerController) AddProduct(c *gin.Context) {

	product := new(vo.Product)
	err := json.NewDecoder(c.Request.Body).Decode(&product)
	if err != nil {
		apihelper.CreateResponseWithErrExplanation(c, err.Error(), err.Error(), http.StatusBadRequest)
		return
	}
	PartnerID := c.Params.ByName("id")

	ps := new(service.PartnerService)
	err = ps.AddProduct(product, PartnerID)
	if err != nil {
		apihelper.CreateResponseWithErrExplanation(c, err.Error(), err.Error(), http.StatusUnauthorized)
		return
	}

	result := make(map[string]interface{})
	result[constant.DATA()] = product
	apihelper.CreateResponse(c, result)

}
