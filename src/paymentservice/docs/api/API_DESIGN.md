# Payments REST API - Design
----
  _This document provides a basic design for the Payments REST API._
---

# Create a payment

Create a payment.

**URL** : `/v1/payments/`

**Method** : `POST`

**Headers** : Content-Type: application/json

**Auth required** : NO

**Permissions required** : None

**Data example** All fields must be sent.

```json
{    
    "type": "Payment",    
    "version": 0,
    "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "attributes": {
        "amount": "100.21",
        "beneficiary_party": {
            "account_name": "W Owens",
            "account_number": "31926819",
            "account_number_code": "BBAN",
            "account_type": 0,
            "address": "1 The Beneficiary Localtown SE2",
            "bank_id": "403000",
            "bank_id_code": "GBDSC",
            "name": "Wilfred Jeremiah Owens"
        },
        "charges_information": {
            "bearer_code": "SHAR",
            "sender_charges": [
                {
                    "amount": "5.00",
                    "currency": "GBP"
                },
                {
                    "amount": "10.00",
                    "currency": "USD"
                }
            ],
            "receiver_charges_amount": "1.00",
            "receiver_charges_currency": "USD"
        },
        "currency": "GBP",
        "debtor_party": {
            "account_name": "EJ Brown Black",
            "account_number": "GB29XABC10161234567801",
            "account_number_code": "IBAN",
            "address": "10 Debtor Crescent Sourcetown NE1",
            "bank_id": "203301",
            "bank_id_code": "GBDSC",
            "name": "Emelia Jane Brown"
        },
        "end_to_end_reference": "Wil piano Jan",
        "fx": {
            "contract_reference": "FX123",
            "exchange_rate": "2.00000",
            "original_amount": "200.42",
            "original_currency": "USD"
        },
        "numeric_reference": "1002001",
        "payment_id": "123456789012345678",
        "payment_purpose": "Paying for goods/services",
        "payment_scheme": "FPS",
        "payment_type": "Credit",
        "processing_date": "2017-01-18",
        "reference": "Payment for Em's piano lessons",
        "scheme_payment_sub_type": "InternetBanking",
        "scheme_payment_type": "ImmediatePayment",
        "sponsor_party": {
            "account_number": "56781234",
            "bank_id": "123123",
            "bank_id_code": "GBDSC"
        }
    }
}
```

## Success Response

**Condition** : If everything is OK and a payment didn't exist.

**Code** : `201 CREATED`

**Content example**

```json
{
  "data": {
    "id": "5b14fc8d530a6940e3ffc7b2",
    "type": "Payment",
    "version": 0,
    "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "attributes": {
      "amount": "100.21",
      "beneficiary_party": {
        "account_name": "W Owens",
        "account_number": "31926819",
        "account_number_code": "BBAN",
        "account_type": 0,
        "address": "1 The Beneficiary Localtown SE2",
        "bank_id": "403000",
        "bank_id_code": "GBDSC",
        "name": "Wilfred Jeremiah Owens"
      },
      "charges_information": {
        "bearer_code": "SHAR",
        "sender_charges": [
          {
            "amount": "5.00",
            "currency": "GBP"
          },
          {
            "amount": "10.00",
            "currency": "USD"
          }
        ],
        "receiver_charges_amount": "1.00",
        "receiver_charges_currency": "USD"
      },
      "currency": "GBP",
      "debtor_party": {
        "account_name": "EJ Brown Black",
        "account_number": "GB29XABC10161234567801",
        "account_number_code": "IBAN",
        "account_type": 0,
        "address": "10 Debtor Crescent Sourcetown NE1",
        "bank_id": "203301",
        "bank_id_code": "GBDSC",
        "name": "Emelia Jane Brown"
      },
      "end_to_end_reference": "Wil piano Jan",
      "fx": {
        "contract_reference": "FX123",
        "exchange_rate": "2.00000",
        "original_amount": "200.42",
        "original_currency": "USD"
      },
      "numeric_reference": "1002001",
      "payment_id": "123456789012345678",
      "payment_purpose": "Paying for goods/services",
      "payment_scheme": "FPS",
      "payment_type": "Credit",
      "processing_date": "2017-01-18",
      "reference": "Payment for Em's piano lessons",
      "scheme_payment_sub_type": "InternetBanking",
      "scheme_payment_type": "ImmediatePayment",
      "sponsor_party": {
        "account_name": "",
        "account_number": "56781234",
        "account_number_code": "",
        "account_type": 0,
        "address": "",
        "bank_id": "123123",
        "bank_id_code": "GBDSC",
        "name": ""
      }
    }
  }
}
```

## Error Responses
**Condition** : If an invalid payment's payload was sent.

**Code** : `400 BAD REQUEST`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Invalid request payload",
      "status": 400
    }
  ]
}
```
<br>
<br>
---
<div style="page-break-after: always;"></div>

# Show all payments

Show all available payments.

**URL** : `/v1/payments/`

**Method** : `GET`

**Headers** : Content-Type: application/json

**Auth required** : NO

**Permissions required** : None

**Data constraints** : `{}`

## Success Responses

**Condition** : There are no payments available.

**Code** : `200 OK`

**Content** : `{"data": []}`

### OR

**Condition** : There is one or more payments available.

**Code** : `200 OK`

**Content** : In this example, thre are two payments:
```json
{
"data": [
    {
        "type": "Payment",
        "id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
        "version": 0,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "amount": "100.21",
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "account_type": 0,
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "amount": "5.00",
                        "currency": "GBP"
                    },
                    {
                        "amount": "10.00",
                        "currency": "USD"
                    }
                ],
                "receiver_charges_amount": "1.00",
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_amount": "200.42",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    },
    {
        "type": "Payment",
        "id": "216d4da9-e59a-4cc6-8df3-3da6e7580b77",
        "version": 0,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "amount": "100.21",
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "account_type": 0,
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "amount": "5.00",
                        "currency": "GBP"
                    },
                    {
                        "amount": "10.00",
                        "currency": "USD"
                    }
                ],
                "receiver_charges_amount": "1.00",
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_amount": "200.42",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }
  ]
}
```
---
<br>
<br>
<div style="page-break-after: always;"></div>

# Show a single payment

Find and show a single payment by ID.

**URL** : `/v1/payments/:pk`

**URL Parameters** : `pk=[ObjectID]` where `pk` is the ID of the payment on the server.

**Method** : `GET`

**Headers** : Content-Type: application/json

**Auth required** : NO

**Data**: `{}`

## Success Response

**Condition** : If payment exists.

**Code** : `200 OK`

**Content**

```json
{
  "data": {
    "id": "5b14ed45530a693ba4246e96",
    "type": "Payment",
    "version": 0,
    "organisation_id": "88888-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "attributes": {
      "amount": "100.21",
      "beneficiary_party": {
        "account_name": "W Owens",
        "account_number": "31926819",
        "account_number_code": "BBAN",
        "account_type": 0,
        "address": "1 The Beneficiary Localtown SE2",
        "bank_id": "403000",
        "bank_id_code": "GBDSC",
        "name": "Wilfred Jeremiah Owens"
      },
      "charges_information": {
        "bearer_code": "SHAR",
        "sender_charges": [
          {
            "amount": "5.00",
            "currency": "GBP"
          },
          {
            "amount": "10.00",
            "currency": "USD"
          }
        ],
        "receiver_charges_amount": "1.00",
        "receiver_charges_currency": "USD"
      },
      "currency": "GBP",
      "debtor_party": {
        "account_name": "EJ Brown Black",
        "account_number": "GB29XABC10161234567801",
        "account_number_code": "IBAN",
        "account_type": 0,
        "address": "10 Debtor Crescent Sourcetown NE1",
        "bank_id": "203301",
        "bank_id_code": "GBDSC",
        "name": "Emelia Jane Brown"
      },
      "end_to_end_reference": "Wil piano Jan",
      "fx": {
        "contract_reference": "FX123",
        "exchange_rate": "2.00000",
        "original_amount": "200.42",
        "original_currency": "USD"
      },
      "numeric_reference": "1002001",
      "payment_id": "123456789012345678",
      "payment_purpose": "Paying for goods/services",
      "payment_scheme": "FPS",
      "payment_type": "Credit",
      "processing_date": "2017-01-18",
      "reference": "Payment for Em's piano lessons",
      "scheme_payment_sub_type": "InternetBanking",
      "scheme_payment_type": "ImmediatePayment",
      "sponsor_party": {
        "account_name": "",
        "account_number": "56781234",
        "account_number_code": "",
        "account_type": 0,
        "address": "",
        "bank_id": "123123",
        "bank_id_code": "GBDSC",
        "name": ""
      }
    }
  }
}
```

## Error Responses

**Condition** : If payment does not exist with `id` of provided `pk` parameter.

**Code** : `404 NOT FOUND`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Payment ID does not exist",
      "status": 404
    }
  ]
}
```

### Or

**Condition** : If payment ID is in invalid format.

**Code** : `400 BAD REQUEST`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Invalid Payment ID",
      "status": 400
    }
  ]
}
```

## Notes

* 
* 
<br>
<br>
---
<div style="page-break-after: always;"></div>


# Delete a payment

Delete a payment from database.

**URL** : `/v1/payments/:pk`

**URL Parameters** : `pk=[ObjectID]` where `pk` is the ID of the payment on the server.

**Method** : `DELETE`

**Headers** : Content-Type: application/json

**Auth required** : NO

**Data**: `{}`

## Success Response

**Condition** : If the payment exists.

**Code** : `204 NO CONTENT`

**Content**: 
```json
{
  "data": {}
}
```

## Error Responses

**Condition** : If payment does not exist with `id` of provided `pk` parameter.

**Code** : `404 NOT FOUND`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Payment ID does not exist",
      "status": 404
    }
  ]
}
```

### Or

**Condition** : If payment ID is in invalid format.

**Code** : `400 BAD REQUEST`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Invalid Payment ID",
      "status": 400
    }
  ]
}
```

## Notes

* 
* 
<br>
<br>
---
<div style="page-break-after: always;"></div>


# Update a payment

Fully update/override a payment on database. Use PATCH for partial update.

**URL** : `/v1/payments/:pk`

**URL Parameters** : `pk=[ObjectID]` where `pk` is the ID of the payment on the server.

**Method** : `PUT`

**Headers** : Content-Type: application/json

**Auth required** : NO

**Data**: 
```json
{
    "id": "5b14222b530a69fd4390de25",
    "type": "Payment",    
    "version": 0,
    "organisation_id": "88888-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "attributes": {
        "amount": "100.21",
        "beneficiary_party": {
            "account_name": "W Owens",
            "account_number": "31926819",
            "account_number_code": "BBAN",
            "account_type": 0,
            "address": "1 The Beneficiary Localtown SE2",
            "bank_id": "403000",
            "bank_id_code": "GBDSC",
            "name": "Wilfred Jeremiah Owens"
        },
        "charges_information": {
            "bearer_code": "SHAR",
            "sender_charges": [
                {
                    "amount": "5.00",
                    "currency": "GBP"
                },
                {
                    "amount": "10.00",
                    "currency": "USD"
                }
            ],
            "receiver_charges_amount": "1.00",
            "receiver_charges_currency": "USD"
        },
        "currency": "GBP",
        "debtor_party": {
            "account_name": "EJ Brown Black",
            "account_number": "GB29XABC10161234567801",
            "account_number_code": "IBAN",
            "address": "10 Debtor Crescent Sourcetown NE1",
            "bank_id": "203301",
            "bank_id_code": "GBDSC",
            "name": "Emelia Jane Brown"
        },
        "end_to_end_reference": "Wil piano Jan",
        "fx": {
            "contract_reference": "FX123",
            "exchange_rate": "2.00000",
            "original_amount": "200.42",
            "original_currency": "USD"
        },
        "numeric_reference": "1002001",
        "payment_id": "123456789012345678",
        "payment_purpose": "Paying for goods/services",
        "payment_scheme": "FPS",
        "payment_type": "Credit",
        "processing_date": "2017-01-18",
        "reference": "Payment for Em's piano lessons",
        "scheme_payment_sub_type": "InternetBanking",
        "scheme_payment_type": "ImmediatePayment",
        "sponsor_party": {
            "account_number": "56781234",
            "bank_id": "123123",
            "bank_id_code": "GBDSC"
        }
    }
}
```

## Success Response

**Condition** : If the payment exists.

**Code** : `200 OK`

**Content**: 
```json
{
  "data": {
    "id": "5b14ed45530a693ba4246e96",
    "type": "Payment",
    "version": 0,
    "organisation_id": "88888-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "attributes": {
      "amount": "100.21",
      "beneficiary_party": {
        "account_name": "W Owens",
        "account_number": "31926819",
        "account_number_code": "BBAN",
        "account_type": 0,
        "address": "1 The Beneficiary Localtown SE2",
        "bank_id": "403000",
        "bank_id_code": "GBDSC",
        "name": "Wilfred Jeremiah Owens"
      },
      "charges_information": {
        "bearer_code": "SHAR",
        "sender_charges": [
          {
            "amount": "5.00",
            "currency": "GBP"
          },
          {
            "amount": "10.00",
            "currency": "USD"
          }
        ],
        "receiver_charges_amount": "1.00",
        "receiver_charges_currency": "USD"
      },
      "currency": "GBP",
      "debtor_party": {
        "account_name": "EJ Brown Black",
        "account_number": "GB29XABC10161234567801",
        "account_number_code": "IBAN",
        "account_type": 0,
        "address": "10 Debtor Crescent Sourcetown NE1",
        "bank_id": "203301",
        "bank_id_code": "GBDSC",
        "name": "Emelia Jane Brown"
      },
      "end_to_end_reference": "Wil piano Jan",
      "fx": {
        "contract_reference": "FX123",
        "exchange_rate": "2.00000",
        "original_amount": "200.42",
        "original_currency": "USD"
      },
      "numeric_reference": "1002001",
      "payment_id": "123456789012345678",
      "payment_purpose": "Paying for goods/services",
      "payment_scheme": "FPS",
      "payment_type": "Credit",
      "processing_date": "2017-01-18",
      "reference": "Payment for Em's piano lessons",
      "scheme_payment_sub_type": "InternetBanking",
      "scheme_payment_type": "ImmediatePayment",
      "sponsor_party": {
        "account_name": "",
        "account_number": "56781234",
        "account_number_code": "",
        "account_type": 0,
        "address": "",
        "bank_id": "123123",
        "bank_id_code": "GBDSC",
        "name": ""
      }
    }
  }
}
```

## Error Responses

**Condition** : If payment does not exist with `id` of provided `pk` parameter.

**Code** : `404 NOT FOUND`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Payment ID does not exist",
      "status": 404
    }
  ]
}
```

### Or

**Condition** : If payment ID is in invalid format.

**Code** : `400 BAD REQUEST`

**Content** : 
```json
{
  "errors": [
    {
      "detail": "Invalid Payment ID",
      "status": 400
    }
  ]
}
```

## Notes

* Additional notes 
* 
<br>
<br>
---