package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/juju/mgosession"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/payment"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	env := os.Getenv("PAYMENT_ENV")
	if env == "" {
		env = "dev"
	}
	err := godotenv.Load("config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session, err := mgo.Dial(os.Getenv("MONGODB_HOST"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	cPool, err := strconv.Atoi(os.Getenv("MONGODB_CONNECTION_POOL"))
	if err != nil {
		log.Println(err.Error())
		cPool = 10
	}
	mPool := mgosession.NewPool(nil, session, cPool)
	defer mPool.Close()

	paymentRepo := payment.NewMongoRepository(mPool)
	paymentService := payment.NewService(paymentRepo)
	all, err := paymentService.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(all) == 0 {
		log.Fatal(entity.ErrNotFound.Error())
	}
	for _, j := range all {
		fmt.Printf("%s %s \n", j.ID.String(), j.Type)
	}
}
