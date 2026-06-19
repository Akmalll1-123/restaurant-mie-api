package xendit

import (
	"fmt"
	"time"
)

type InvoiceResponse struct {
	ID         string
	InvoiceURL string
}

func CreateInvoice(
	orderID uint,
	amount float64,
) (*InvoiceResponse, error) {

	invoiceID := fmt.Sprintf(
		"INV-%d-%d",
		orderID,
		time.Now().UnixNano(),
	)

	return &InvoiceResponse{
		ID: invoiceID,

		InvoiceURL: fmt.Sprintf(
			"https://checkout.xendit.co/%s",
			invoiceID,
		),
	}, nil
}
