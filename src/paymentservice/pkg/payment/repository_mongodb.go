package payment

import (
	"os"

	"github.com/juju/mgosession"
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool *mgosession.Pool
}

//NewMongoRepository create new repository
func NewMongoRepository(p *mgosession.Pool) *MongoRepository {
	return &MongoRepository{
		pool: p,
	}
}

//Find a payment
func (r *MongoRepository) Find(id entity.ID) (*entity.Payment, error) {
	result := entity.Payment{}
	session := r.pool.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("payment")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Store a payment
func (r *MongoRepository) Store(p *entity.Payment) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("payment")
	err := coll.Insert(p)
	if err != nil {
		return entity.ID(0), err
	}
	return p.ID, nil
}

//Update a payment
func (r *MongoRepository) Update(id entity.ID, p *entity.Payment) error {
	session := r.pool.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("payment")

	err := coll.Update(bson.M{"_id": id}, p)
	if err != nil {
		return err
	}
	return nil
}

//FindAll payments
func (r *MongoRepository) FindAll() ([]*entity.Payment, error) {
	// var d []*entity.Payment
	d := make([]*entity.Payment, 0)
	session := r.pool.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("payment")
	err := coll.Find(nil).Sort("name").All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Delete a payment
func (r *MongoRepository) Delete(id entity.ID) error {
	session := r.pool.Session(nil)
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("payment")
	return coll.Remove(bson.M{"_id": id})
}
