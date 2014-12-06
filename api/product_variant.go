package shopify

import (
  
    "encoding/json"
  
    "fmt"
  
    "time"
  
)

type ProductVariant struct {
  
    Barcode string `json:"barcode"`
  
    CompareAtPrice string `json:"compare_at_price"`
  
    CreatedAt time.Time `json:"created_at"`
  
    FulfillmentService string `json:"fulfillment_service"`
  
    Grams int64 `json:"grams"`
  
    Id int64 `json:"id"`
  
    InventoryManagement string `json:"inventory_management"`
  
    InventoryPolicy string `json:"inventory_policy"`
  
    InventoryQuantity time.Time `json:"inventory_quantity"`
  
    OldInventoryQuantity string `json:"old_inventory_quantity"`
  
    InventoryQuantityAdjustment string `json:"inventory_quantity_adjustment"`
  
    Metafield string `json:"metafield"`
  
    Option string `json:"option"`
  
    Position string `json:"position"`
  
    Price float64 `json:"price"`
  
    ProductId int64 `json:"product_id"`
  
    RequiresShipping string `json:"requires_shipping"`
  
    Sku string `json:"sku"`
  
    Taxable string `json:"taxable"`
  
    Title string `json:"title"`
  
    UpdatedAt time.Time `json:"updated_at"`
  
    ImageId int64 `json:"image_id"`
  
}


func (api *API) ProductVariants() (*[]ProductVariant, error) {
  res, status, err := api.request("/admin/product_variants.json", "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := &map[string][]ProductVariant{}
  err = json.NewDecoder(res).Decode(r)

  fmt.Printf("things are: %v\n\n", *r)

  result := (*r)["products"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}




func (api *API) ProductVariant(id int64) (*ProductVariant, error) {
  endpoint := fmt.Sprintf("/admin/product_variants/%d.json", id)

  res, status, err := api.request(endpoint, "GET", nil)

  if err != nil {
    return nil, err
  }

  if status != 200 {
    return nil, fmt.Errorf("Status returned: %d", status)
  }

  r := map[string]ProductVariant{}
  err = json.NewDecoder(res).Decode(&r)

  fmt.Printf("things are: %v\n\n", r)

  result := r["product_variant"]

	if err != nil {
		return nil, err
  }

  return &result, nil
}









