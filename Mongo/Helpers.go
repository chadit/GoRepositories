package Mongo

import "gopkg.in/mgo.v2/bson"

// GetNewBsonIDString get a new objectID as string
func GetNewBsonIDString() string {
	return bson.NewObjectId().Hex()
}

// IsBsonIDValid checks if bson id is valid
func IsBsonIDValid(ID string) bool {
	if ID == "" {
		return false
	}

	return bson.IsObjectIdHex(ID)
}
