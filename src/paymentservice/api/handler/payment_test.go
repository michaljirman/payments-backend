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

func fixture(path string) []byte {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return b
}

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
