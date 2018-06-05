package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/middleware"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/payment"
	"github.com/stretchr/testify/assert"
)

// TestPaymentFind tests API for find a payment call
func TestPaymentFind(t *testing.T) {

	repo := payment.NewInmemRepository()
	service := payment.NewService(repo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//payment
	MakePaymentHandlers(r, *n, service)

	ts := httptest.NewServer(r)
	defer ts.Close()

	payments := generateTestingPayments()
	for _, payment := range payments {
		_, _ = service.Store(&payment)
		res, err := http.Get(ts.URL + "/v1/payments/" + payment.ID.String())
		// t.Log(res)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

// TestPaymentFindAll tests API for find all payments call
func TestPaymentFindAll(t *testing.T) {

	repo := payment.NewInmemRepository()
	service := payment.NewService(repo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//payment
	MakePaymentHandlers(r, *n, service)

	ts := httptest.NewServer(r)
	defer ts.Close()

	payments := generateTestingPayments()
	for _, payment := range payments {
		_, _ = service.Store(&payment)
	}

	res, err := http.Get(ts.URL + "/v1/payments")
	// t.Log(res)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

// fixture prepares payload from json file
func fixture(path string) []byte {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return b
}

// TestPaymentCreate tests API for create a payment call
func TestPaymentCreate(t *testing.T) {

	repo := payment.NewInmemRepository()
	service := payment.NewService(repo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//payment
	MakePaymentHandlers(r, *n, service)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Post(ts.URL+"/v1/payments", "application/json; charset=utf-8", bytes.NewBuffer(fixture("payment_create.json")))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	payments, err := service.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(payments))
}

// TestPaymentDelete tests API for delete a payment call
func TestPaymentDelete(t *testing.T) {

	repo := payment.NewInmemRepository()
	service := payment.NewService(repo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//payment
	MakePaymentHandlers(r, *n, service)

	ts := httptest.NewServer(r)
	defer ts.Close()

	payments := generateTestingPayments()
	client := &http.Client{}
	for _, payment := range payments {
		_, _ = service.Store(&payment)
		req, err := http.NewRequest("DELETE", ts.URL+"/v1/payments/"+payment.ID.String(), nil)
		res, err := client.Do(req)
		// t.Log(res)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, res.StatusCode)
	}
}

// generateTestingPayments generates slice of testing entity.Payment objects for testing purposes
func generateTestingPayments() []entity.Payment {
	return []entity.Payment{
		entity.Payment{
			Type:           "Payment",
			Version:        0,
			OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
			Attributes: entity.Attributes{
				Amount: "100.21",
				BeneficiaryParty: entity.AccountInformation{
					AccountName:       "W Owens",
					AccountNumber:     "31926819",
					AccountNumberCode: "BBAN",
					AccountType:       0,
					Address:           "1 The Beneficiary Localtown SE2",
					BankID:            "403000",
					BankIDCode:        "GBDSC",
					Name:              "Wilfred Jeremiah Owens",
				},
				ChargesInformation: entity.ChargesInformation{
					BearerCode: "SHAR",
					SenderCharges: []entity.SenderCharge{
						entity.SenderCharge{
							Amount:   "5.00",
							Currency: "GBP",
						},
						entity.SenderCharge{
							Amount:   "10.00",
							Currency: "USD",
						},
					},
					ReceiverChargesAmmount:  "1.00",
					ReceiverChargesCurrency: "USD",
				},
				Currency: "GBP",
				DebtorParty: entity.AccountInformation{
					AccountName:       "EJ Brown Black",
					AccountNumber:     "GB29XABC10161234567801",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					BankID:            "203301",
					BankIDCode:        "GBDSC",
					Name:              "Emelia Jane Brown",
				},
				EndToEndReference: "Wil piano Jan",
				Fx: entity.Fx{
					ContractReference: "FX123",
					ExchangeRate:      "2.00000",
					OriginalAmount:    "200.42",
					OriginalCurrency:  "USD",
				},
				NumericReference:     "1002001",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				ProcessingDate:       "2017-01-18",
				Reference:            "Payment for Em's piano lessons",
				SchemePaymentSubType: "InternetBanking",
				SchemePaymentType:    "ImmediatePayment",
				SponsorParty: entity.AccountInformation{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
			},
		},
		entity.Payment{
			Type:           "Payment",
			Version:        0,
			OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
			Attributes: entity.Attributes{
				Amount: "100.21",
				BeneficiaryParty: entity.AccountInformation{
					AccountName:       "W Owens",
					AccountNumber:     "31926819",
					AccountNumberCode: "BBAN",
					AccountType:       0,
					Address:           "1 The Beneficiary Localtown SE2",
					BankID:            "403000",
					BankIDCode:        "GBDSC",
					Name:              "Wilfred Jeremiah Owens",
				},
				ChargesInformation: entity.ChargesInformation{
					BearerCode: "SHAR",
					SenderCharges: []entity.SenderCharge{
						entity.SenderCharge{
							Amount:   "5.00",
							Currency: "GBP",
						},
						entity.SenderCharge{
							Amount:   "10.00",
							Currency: "USD",
						},
					},
					ReceiverChargesAmmount:  "1.00",
					ReceiverChargesCurrency: "USD",
				},
				Currency: "GBP",
				DebtorParty: entity.AccountInformation{
					AccountName:       "EJ Brown Black",
					AccountNumber:     "GB29XABC10161234567801",
					AccountNumberCode: "IBAN",
					Address:           "10 Debtor Crescent Sourcetown NE1",
					BankID:            "203301",
					BankIDCode:        "GBDSC",
					Name:              "Emelia Jane Brown",
				},
				EndToEndReference: "Wil piano Jan",
				Fx: entity.Fx{
					ContractReference: "FX123",
					ExchangeRate:      "2.00000",
					OriginalAmount:    "200.42",
					OriginalCurrency:  "USD",
				},
				NumericReference:     "1002001",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				ProcessingDate:       "2017-01-18",
				Reference:            "Payment for Em's piano lessons",
				SchemePaymentSubType: "InternetBanking",
				SchemePaymentType:    "ImmediatePayment",
				SponsorParty: entity.AccountInformation{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
			},
		},
	}
}
