package domain

import "errors"

var ErrProductStockNegative = errors.New("the amount is bigger than the current quantity")
var ErrProductNotFound = errors.New("product not found")
