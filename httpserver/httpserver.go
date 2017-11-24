package httpserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	DB "syncorder/database"
	"log"
	"net/http"
	"syncorder/config"
	"io/ioutil"
	"time"
	"encoding/json"
)

type OrderInfo struct {
	ID                    int64         `json:"id"`
	Email                 string        `json:"email"`
	ClosedAt              interface{}   `json:"closed_at"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	Number                int           `json:"number"`
	Note                  interface{}   `json:"note"`
	Token                 string        `json:"token"`
	Gateway               string        `json:"gateway"`
	Test                  bool          `json:"test"`
	TotalPrice            string        `json:"total_price"`
	SubtotalPrice         string        `json:"subtotal_price"`
	TotalWeight           int           `json:"total_weight"`
	TotalTax              string        `json:"total_tax"`
	TaxesIncluded         bool          `json:"taxes_included"`
	Currency              string        `json:"currency"`
	FinancialStatus       string        `json:"financial_status"`
	Confirmed             bool          `json:"confirmed"`
	TotalDiscounts        string        `json:"total_discounts"`
	TotalLineItemsPrice   string        `json:"total_line_items_price"`
	CartToken             string        `json:"cart_token"`
	BuyerAcceptsMarketing bool          `json:"buyer_accepts_marketing"`
	Name                  string        `json:"name"`
	ReferringSite         string        `json:"referring_site"`
	LandingSite           string        `json:"landing_site"`
	CancelledAt           interface{}   `json:"cancelled_at"`
	CancelReason          interface{}   `json:"cancel_reason"`
	TotalPriceUsd         string        `json:"total_price_usd"`
	CheckoutToken         string        `json:"checkout_token"`
	Reference             interface{}   `json:"reference"`
	UserID                interface{}   `json:"user_id"`
	LocationID            interface{}   `json:"location_id"`
	SourceIdentifier      interface{}   `json:"source_identifier"`
	SourceURL             interface{}   `json:"source_url"`
	ProcessedAt           time.Time     `json:"processed_at"`
	DeviceID              interface{}   `json:"device_id"`
	Phone                 interface{}   `json:"phone"`
	CustomerLocale        string        `json:"customer_locale"`
	AppID                 int           `json:"app_id"`
	BrowserIP             interface{}   `json:"browser_ip"`
	LandingSiteRef        interface{}   `json:"landing_site_ref"`
	OrderNumber           int           `json:"order_number"`
	DiscountCodes         []interface{} `json:"discount_codes"`
	NoteAttributes        []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"note_attributes"`
	PaymentGatewayNames []string      `json:"payment_gateway_names"`
	ProcessingMethod    string        `json:"processing_method"`
	CheckoutID          int64         `json:"checkout_id"`
	SourceName          string        `json:"source_name"`
	FulfillmentStatus   interface{}   `json:"fulfillment_status"`
	TaxLines            []interface{} `json:"tax_lines"`
	Tags                string        `json:"tags"`
	ContactEmail        string        `json:"contact_email"`
	OrderStatusURL      string        `json:"order_status_url"`
	LineItems           []struct {
		ID                         int64         `json:"id"`
		VariantID                  int64         `json:"variant_id"`
		Title                      string        `json:"title"`
		Quantity                   int           `json:"quantity"`
		Price                      string        `json:"price"`
		Grams                      int           `json:"grams"`
		Sku                        string        `json:"sku"`
		VariantTitle               string        `json:"variant_title"`
		Vendor                     string        `json:"vendor"`
		FulfillmentService         string        `json:"fulfillment_service"`
		ProductID                  int64         `json:"product_id"`
		RequiresShipping           bool          `json:"requires_shipping"`
		Taxable                    bool          `json:"taxable"`
		GiftCard                   bool          `json:"gift_card"`
		PreTaxPrice                string        `json:"pre_tax_price"`
		Name                       string        `json:"name"`
		VariantInventoryManagement string        `json:"variant_inventory_management"`
		Properties                 []interface{} `json:"properties"`
		ProductExists              bool          `json:"product_exists"`
		FulfillableQuantity        int           `json:"fulfillable_quantity"`
		TotalDiscount              string        `json:"total_discount"`
		FulfillmentStatus          interface{}   `json:"fulfillment_status"`
		TaxLines                   []interface{} `json:"tax_lines"`
		OriginLocation             struct {
			ID           int64  `json:"id"`
			CountryCode  string `json:"country_code"`
			ProvinceCode string `json:"province_code"`
			Name         string `json:"name"`
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			City         string `json:"city"`
			Zip          string `json:"zip"`
		} `json:"origin_location"`
		DestinationLocation struct {
			ID           int64  `json:"id"`
			CountryCode  string `json:"country_code"`
			ProvinceCode string `json:"province_code"`
			Name         string `json:"name"`
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			City         string `json:"city"`
			Zip          string `json:"zip"`
		} `json:"destination_location"`
	} `json:"line_items"`
	ShippingLines []struct {
		ID                            int64         `json:"id"`
		Title                         string        `json:"title"`
		Price                         string        `json:"price"`
		Code                          string        `json:"code"`
		Source                        string        `json:"source"`
		Phone                         interface{}   `json:"phone"`
		RequestedFulfillmentServiceID interface{}   `json:"requested_fulfillment_service_id"`
		DeliveryCategory              interface{}   `json:"delivery_category"`
		CarrierIdentifier             interface{}   `json:"carrier_identifier"`
		DiscountedPrice               string        `json:"discounted_price"`
		TaxLines                      []interface{} `json:"tax_lines"`
	} `json:"shipping_lines"`
	BillingAddress struct {
		FirstName    string      `json:"first_name"`
		Address1     string      `json:"address1"`
		Phone        string      `json:"phone"`
		City         string      `json:"city"`
		Zip          string      `json:"zip"`
		Province     interface{} `json:"province"`
		Country      string      `json:"country"`
		LastName     string      `json:"last_name"`
		Address2     interface{} `json:"address2"`
		Company      string      `json:"company"`
		Latitude     interface{} `json:"latitude"`
		Longitude    interface{} `json:"longitude"`
		Name         string      `json:"name"`
		CountryCode  string      `json:"country_code"`
		ProvinceCode interface{} `json:"province_code"`
	} `json:"billing_address"`
	ShippingAddress struct {
		FirstName    string      `json:"first_name"`
		Address1     string      `json:"address1"`
		Phone        string      `json:"phone"`
		City         string      `json:"city"`
		Zip          string      `json:"zip"`
		Province     interface{} `json:"province"`
		Country      string      `json:"country"`
		LastName     string      `json:"last_name"`
		Address2     interface{} `json:"address2"`
		Company      string      `json:"company"`
		Latitude     interface{} `json:"latitude"`
		Longitude    interface{} `json:"longitude"`
		Name         string      `json:"name"`
		CountryCode  string      `json:"country_code"`
		ProvinceCode interface{} `json:"province_code"`
	} `json:"shipping_address"`
	Fulfillments  []interface{} `json:"fulfillments"`
	ClientDetails struct {
		BrowserIP      string      `json:"browser_ip"`
		AcceptLanguage string      `json:"accept_language"`
		UserAgent      string      `json:"user_agent"`
		SessionHash    interface{} `json:"session_hash"`
		BrowserWidth   int         `json:"browser_width"`
		BrowserHeight  int         `json:"browser_height"`
	} `json:"client_details"`
	Refunds  []interface{} `json:"refunds"`
	Customer struct {
		ID                  int64       `json:"id"`
		Email               string      `json:"email"`
		AcceptsMarketing    bool        `json:"accepts_marketing"`
		CreatedAt           time.Time   `json:"created_at"`
		UpdatedAt           time.Time   `json:"updated_at"`
		FirstName           string      `json:"first_name"`
		LastName            string      `json:"last_name"`
		OrdersCount         int         `json:"orders_count"`
		State               string      `json:"state"`
		TotalSpent          string      `json:"total_spent"`
		LastOrderID         int64       `json:"last_order_id"`
		Note                interface{} `json:"note"`
		VerifiedEmail       bool        `json:"verified_email"`
		MultipassIdentifier interface{} `json:"multipass_identifier"`
		TaxExempt           bool        `json:"tax_exempt"`
		Phone               interface{} `json:"phone"`
		Tags                string      `json:"tags"`
		LastOrderName       string      `json:"last_order_name"`
		DefaultAddress      struct {
			ID           int64       `json:"id"`
			CustomerID   int64       `json:"customer_id"`
			FirstName    string      `json:"first_name"`
			LastName     string      `json:"last_name"`
			Company      string      `json:"company"`
			Address1     string      `json:"address1"`
			Address2     interface{} `json:"address2"`
			City         string      `json:"city"`
			Province     interface{} `json:"province"`
			Country      string      `json:"country"`
			Zip          string      `json:"zip"`
			Phone        string      `json:"phone"`
			Name         string      `json:"name"`
			ProvinceCode interface{} `json:"province_code"`
			CountryCode  string      `json:"country_code"`
			CountryName  string      `json:"country_name"`
			Default      bool        `json:"default"`
		} `json:"default_address"`
	} `json:"customer"`
}

func orderCreate(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	if r.Method == "POST" {
		hmac246 := r.Header.Get("X-Shopify-Hmac-Sha256")
		result, _ := ioutil.ReadAll(r.Body)
		caculateHmac256 := ComputeHmac256(string(result), config.AppConfig.Shopify.SharedSecret)

		if hmac246 == caculateHmac256 {

			orderStruct := &OrderInfo{}

			err := json.Unmarshal(result, orderStruct)

			if err != nil {
				log.Fatal(err.Error())
			}

			db := DB.GetInstance()
			db.Write("unfulfilled", orderStruct.Name, orderStruct)

		} else {
			log.Fatal("hmac256 verified failed")
		}

		rw.Write([]byte(config.AppConfig.Fetchx.Authorization))

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("method not allowed"))
	}

}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Start() {
	log.Println("Start listening on port " + config.AppConfig.HttpPort)
	http.HandleFunc("/orderCreate", orderCreate)
	err := http.ListenAndServe(":" + config.AppConfig.HttpPort, nil)

	if err != nil {
		log.Println(err.Error())
	}

}

func writeDB() {

	//
	//// Read a fish from the database (passing fish by reference)
	//onefish := Fish{}
	//if err := db.Read("unfulfilled", "onefish", &onefish); err != nil {
	//	fmt.Println("Error", err)
	//}
	//

}

