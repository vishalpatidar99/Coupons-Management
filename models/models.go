package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`         // Product name
	Description string  `json:"description"`  // Product description
	Price       float64 `json:"price"`        // Product price
	IsAvailable bool    `json:"is_available"` // Stock available or not
}

type Cart struct {
	gorm.Model
	UserID     uint      `json:"user_id"`     // User id
	Products   []Product `json:"products"`    // List of products in cart
	TotalPrice float64   `json:"total_price"` // Total Price after calculation
}

type Coupon struct {
	gorm.Model
	Label             string    `json:"label"`               //Descriptive name of the coupon
	Type              string    `json:"type"`                //Type of coupon: cart-wise, product-wise, bxgy etc.
	Discount          string    `json:"discount"`            //Actual Discount going to be apply on things
	TermsAndCondition string    `json:"terms_and_condition"` //Terms and condition of coupon to be apply
	Threshold         float64   `json:"threshold"`           //Threshold of price in case of cart-wise coupon
	Applied           bool      // Detects coupon is applied or not initial valued will be false
	ExpirationDate    time.Time `json:"expiration_date"` // Expriration date of coupon
}

type CouponApplication struct {
	gorm.Model
	CouponID   uint    `json:"coupon_id"`   // References the coupon
	CartID     uint    `json:"cart_id"`     // References the cart
	Discount   float64 `json:"discount"`    // Discount applied by the coupon
	AfterPrice float64 `json:"after_price"` //Price after discount deduction
}

type ApplicableCouponsResponse struct {
	Cart              Cart     // Cart object
	ApplicableCoupons []Coupon // List of all coupons can be apply on specific cart
	TotalDiscount     float64  // Total discount after calculation
}
