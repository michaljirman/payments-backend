package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/payment"
)

// respondWithError generates error response from status code and message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	errorPayload := map[string]interface{}{
		"status": code,
		"detail": msg,
	}
	errorsPayload := []map[string]interface{}{errorPayload}

	response, _ := json.Marshal(map[string]interface{}{"errors": errorsPayload})
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithJson generates response from status code and payload
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if payload == nil {
		payload = make(map[string]interface{})
	}
	response, _ := json.Marshal(map[string]interface{}{"data": payload})
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(code)
	w.Write(response)

}

// paymentFindAll GET request for list of payments
func paymentFindAll(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payments, err := service.FindAll()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, payments)
	})
}

// paymentFind handles GET request to find a payment by its paymentID
func paymentFind(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paymentID, ok := vars["paymentID"]
		if !ok {
			respondWithError(w, http.StatusNotFound, "Missing route parameter 'paymentID'")
			return
		}
		if entity.IsValidID(paymentID) {
			payment, err := service.Find(entity.StringToID(paymentID))
			if err != nil {
				respondWithError(w, http.StatusNotFound, "Payment ID does not exist")
				return
			}
			respondWithJSON(w, http.StatusOK, payment)
		} else {
			respondWithError(w, http.StatusBadRequest, "Invalid Payment ID")
			return
		}
	})
}

// paymentCreate handles POST request to create a new payment
func paymentCreate(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var p *entity.Payment
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		p.ID, err = service.Store(p)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, p)
	})
}

// paymentUpdate handles PUT request to update an existing payment
func paymentUpdate(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paymentID, ok := vars["paymentID"]
		if !ok {
			respondWithError(w, http.StatusNotFound, "Missing route parameter 'paymentID'")
			return
		}
		defer r.Body.Close()
		var payment *entity.Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if err := service.Update(entity.StringToID(paymentID), payment); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, payment)
	})
}

// paymentDelete handles DELETE request of existing payment
func paymentDelete(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paymentID, ok := vars["paymentID"]
		if !ok {
			respondWithError(w, http.StatusNotFound, "Missing route parameter 'paymentID'")
			return
		}
		if entity.IsValidID(paymentID) {
			err := service.Delete(entity.StringToID(paymentID))
			if err != nil {
				respondWithError(w, http.StatusNotFound, "Payment ID does not exist")
				return
			}
			respondWithJSON(w, http.StatusNoContent, nil)
		} else {
			respondWithError(w, http.StatusBadRequest, "Invalid Payment ID")
			return
		}
	})
}

//MakePaymentHandlers creates a subrouter with handles for payments api
func MakePaymentHandlers(r *mux.Router, n negroni.Negroni, service payment.UseCase) {
	paymentsRouter := r.PathPrefix("/v1/payments").Subrouter()
	paymentsRouter.Handle("", n.With(
		negroni.Wrap(paymentFindAll(service)),
	)).Methods("GET", "OPTIONS")

	paymentsRouter.Handle("/{paymentID}", n.With(
		negroni.Wrap(paymentFind(service)),
	)).Methods("GET", "OPTIONS")

	paymentsRouter.Handle("", n.With(
		negroni.Wrap(paymentCreate(service)),
	)).Methods("POST", "OPTIONS")

	paymentsRouter.Handle("/{paymentID}", n.With(
		negroni.Wrap(paymentUpdate(service)),
	)).Methods("PUT", "OPTIONS")

	paymentsRouter.Handle("/{paymentID}", n.With(
		negroni.Wrap(paymentDelete(service)),
	)).Methods("DELETE", "OPTIONS")
}
