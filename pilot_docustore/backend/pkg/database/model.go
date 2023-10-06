package database

// Document is the base JSON document stored in the database.
type Document struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Data string `json:"data" bson:"data"`
}
