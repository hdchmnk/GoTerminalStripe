package cards

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, error, string) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Params:   stripe.Params{},
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeError, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeError.Code)
		}
		return nil, err, msg
	}
	return pi, nil, ""
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your cards was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your cards expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect ZIP code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is too large to charge to your cards"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is too small to charge to your cards"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Your postal code is invalid"
	default:
		msg = "Your cards was declined"
	}
	return msg
}
