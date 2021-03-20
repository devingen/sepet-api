package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// BucketStatus is the type for the Bucket's status
type BucketStatus string

const (
	// BucketStatusActive is status active
	BucketStatusActive BucketStatus = "active"
)

// Bucket defines the MongoDB and JSON structure of the bucket data
type Bucket struct {
	// DBRef fields
	Ref      string             `bson:"_ref,omitempty" json:"_ref,omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Database string             `bson:"_db,omitempty" json:"_db,omitempty"`

	// common model fields
	CreatedAt *time.Time `json:"_created,omitempty" bson:"_created,omitempty"`
	UpdatedAt *time.Time `json:"_updated,omitempty" bson:"_updated,omitempty"`
	Revision  int        `json:"_revision,omitempty" bson:"_revision,omitempty"`

	Domain              string       `json:"domain,omitempty" bson:"domain,omitempty"`
	Folder              string       `json:"folder,omitempty" bson:"folder,omitempty"`
	Version             string       `json:"version,omitempty" bson:"version,omitempty"`
	IndexPagePath       string       `json:"indexPagePath,omitempty" bson:"indexPagePath"`
	ErrorPagePath       string       `json:"errorPagePath,omitempty" bson:"errorPagePath"`
	IsCacheEnabled      bool         `json:"isCacheEnabled,omitempty" bson:"isCacheEnabled"`
	IsVersioningEnabled bool         `json:"isVersioningEnabled,omitempty" bson:"isVersioningEnabled"`
	Status              BucketStatus `json:"status,omitempty" bson:"status,omitempty"`
}

// AddCreationFields adds the necessary fields before inserting into database
func (b *Bucket) AddCreationFields() {
	b.ID = primitive.NewObjectID()
	now := time.Now()
	b.CreatedAt = &now
	b.UpdatedAt = &now
	b.Revision = 1
}

// PrepareUpdateFields sets the UpdatedAt and deletes the Revision. Giving 0 value to Revision results bson
// ignoring the revision field in $set function. It's incremented by the $inc command
func (b *Bucket) PrepareUpdateFields() {
	b.Revision = 0
	now := time.Now()
	b.UpdatedAt = &now
}
