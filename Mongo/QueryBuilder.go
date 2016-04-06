package Mongo

import "gopkg.in/mgo.v2/bson"

// QueryOptions are a set of options that can be passed to the mongo database
type QueryOptions struct {
	Projection bson.M
	Sort       string
	Limit      int
	Skip       int
}

// SetQueryOptionDefaults - initialize the defaults
func (queryOptions *QueryOptions) SetQueryOptionDefaults() {
	queryOptions.Sort = "$natural"
}
