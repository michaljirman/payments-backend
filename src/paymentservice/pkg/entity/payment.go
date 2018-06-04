package entity

// AccountInformation datat
type AccountInformation struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// SenderCharge data
type SenderCharge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// ChargesInformation data
type ChargesInformation struct {
	BearerCode              string         `json:"bearer_code"`
	SenderCharges           []SenderCharge `json:"sender_charges"`
	ReceiverChargesAmmount  string         `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string         `json:"receiver_charges_currency"`
}

// Fx data
type Fx struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// Attributes data
type Attributes struct {
	Amount               string             `json:"amount"`
	BeneficiaryParty     AccountInformation `json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `json:"charges_information"`
	Currency             string             `json:"currency"`
	DebtorParty          AccountInformation `json:"debtor_party"`
	EndToEndReference    string             `json:"end_to_end_reference"`
	Fx                   Fx                 `json:"fx"`
	NumericReference     string             `json:"numeric_reference"`
	PaymentID            string             `json:"payment_id"`
	PaymentPurpose       string             `json:"payment_purpose"`
	PaymentScheme        string             `json:"payment_scheme"`
	PaymentType          string             `json:"payment_type"`
	ProcessingDate       string             `json:"processing_date"`
	Reference            string             `json:"reference"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type"`
	SchemePaymentType    string             `json:"scheme_payment_type"`
	SponsorParty         AccountInformation `json:"sponsor_party"`
}

// Payment data
type Payment struct {
	ID             ID         `json:"id" bson:"_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}
