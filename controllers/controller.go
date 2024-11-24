package controllers

type Controller struct {
	Brand    *BrandController
	Category *CategoryController
	Product  *ProductController
	Auth     *AuthController
	Order    *OrderController
}
