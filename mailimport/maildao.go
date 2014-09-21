package mailimport

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

const (
    COLLECTION_MAIL = "mail"
)

// ---------------------- 
// Declarations

// DAO for Mail collection.
// @author Rémy MATHIEU
type MailDAO struct {
    mongo       *Mongo
    collection  *mgo.Collection
}

// ---------------------- 
// Methods

func NewMailDAO(c Config, m *Mongo) *MailDAO {
    return &MailDAO{m, m.GetCollection(c, COLLECTION_MAIL)}
}

func (d *MailDAO) Save(mail *Mail) error {
    if (len(mail.Id) > 0) {
        return d.collection.Update(bson.M{"_id": mail.Id}, mail)
    } else {
        return d.collection.Insert(mail)
    }
}
