package Mongo

import (
	"time"

	"github.com/chadit/GoShare"

	"gopkg.in/mgo.v2/bson"
)

// Tenant model
type Tenant struct {
	// (ReadOnly) Id of the document, created by system utilizing Mongo Bson Id
	ID string `json:"Id"  bson:"_id" binding:"required"`
	// (ReadOnly) Date the document was created (UTC)
	DateCreated time.Time `json:"DateCreated" bson:"DateCreated" binding:"required"`
	// (ReadOnly) Date the document was modified (UTC)
	DateModified time.Time `json:"DateModified" bson:"DateModified"`
	// (ReadOnly) What user modified the document
	UserModified string `json:"UserModified" bson:"UserModified,omitempty"`
	// (ReadOnly) Has the record been deleted
	IsDeleted bool `json:"IsDeleted" bson:"IsDeleted,omitempty"`
	// (Migration/Sync only)ExternalId is used to sync with external systems (Not used for anything internally)
	ExternalID string `json:"ExternalId" bson:"ExternalId,omitempty"`
	// (Migration/Sync only)ExternalNumber is used to sync with external systems (Not used for anything internally)
	ExternalNumber string `json:"ExternalNumber" bson:"ExternalNumber,omitempty"`
	// TenantID pertains to the tenant the document belongs to.  this is not serialized out
	TenantID string `json:"-" bson:"TenantId"`

	APIKey           string      `json:"ApiKey" bson:"ApiKey" binding:"required"`
	Name             string      `json:"Name"  bson:"Name" binding:"required"`
	FullName         string      `json:"FullName"  bson:"FullName" binding:"required"`
	SubDomain        string      `json:"SubDomain"  bson:"SubDomain" binding:"required"`
	Comments         string      `json:"Comments"  bson:"Comments"`
	DatabaseNode     string      `json:"DatabaseNode"  bson:"DatabaseNode"`
	DatabaseName     string      `json:"DatabaseName"  bson:"DatabaseName"`
	IsEnabled        bool        `json:"IsEnabled"  bson:"IsEnabled"`
	PrimaryContactID string      `json:"PrimaryContactId"  bson:"PrimaryContactId"`
	Status           string      `json:"Status"  bson:"Status"`
	CustomerCode     string      `json:"CustomerCode"  bson:"CustomerCode"`
	Address          interface{} `json:"Address"  bson:"Address"`
	PhoneNumber      interface{} `json:"PhoneNumber"  bson:"PhoneNumber"`
	DispatchSystem   string      `json:"DispatchSystem"  bson:"DispatchSystem"`
}

// NewTenant gets a new object
func NewTenant(tenantID string) Tenant {
	newItem := Tenant{}
	newItem.Init(tenantID)
	return newItem
}

// Init sets the defaults
func (p *Tenant) Init(tenantID string) {
	eventTime := time.Now().UTC()
	p.ID = GetNewBsonIDString()
	if p.DateCreated.IsZero() {
		p.DateCreated = eventTime
	}
	p.DateModified = eventTime
	p.IsDeleted = false
	p.TenantID = tenantID
}

// SetupSave updates the object
func (p *Tenant) SetupSave(tenantID string) {
	eventTime := time.Now().UTC()
	if p.ID == "" {
		p.ID = bson.NewObjectId().Hex()
	}

	if p.DateCreated.IsZero() {
		p.DateCreated = eventTime
	}

	p.DateModified = eventTime
	p.TenantID = tenantID
}

// TenantConnectionInfo model
type TenantConnectionInfo struct {
	Tenant           Tenant `json:"-" bson:"-"`
	ConnectionString string `json:"-"  bson:"-"`
	DatabaseNode     string `json:"-"  bson:"-"`
	DatabaseName     string `json:"-"  bson:"-"`
	UserModified     string `json:"-" bson:"-"`
	CollectionName   string `json:"-" bson:"-"`
}

// InitBaseTenantConnectionInfo sets the defaults for a new userlogin
func (p *TenantConnectionInfo) InitBaseTenantConnectionInfo() {
	p.ConnectionString = Shares.MongoConnection
}

// Node model used to hold tenant connection information
type Node struct {
	// (ReadOnly) Id of the document, created by system utilizing Mongo Bson Id
	ID string `json:"Id"  bson:"_id" binding:"required"`
	// (ReadOnly) Date the document was created (UTC)
	DateCreated time.Time `json:"DateCreated" bson:"DateCreated" binding:"required"`
	// (ReadOnly) Date the document was modified (UTC)
	DateModified time.Time `json:"DateModified" bson:"DateModified"`
	// (ReadOnly) What user modified the document
	UserModified string `json:"UserModified" bson:"UserModified,omitempty"`
	// (ReadOnly) Has the record been deleted
	IsDeleted bool `json:"IsDeleted" bson:"IsDeleted,omitempty"`
	// (Migration/Sync only)ExternalId is used to sync with external systems (Not used for anything internally)
	ExternalID string `json:"ExternalId" bson:"ExternalId,omitempty"`
	// (Migration/Sync only)ExternalNumber is used to sync with external systems (Not used for anything internally)
	ExternalNumber string `json:"ExternalNumber" bson:"ExternalNumber,omitempty"`
	// TenantID pertains to the tenant the document belongs to.  this is not serialized out
	TenantID   string `json:"-" bson:"TenantId"`
	Name       string `json:"Name" bson:"Name" binding:"required"`
	Connection string `json:"Connection"  bson:"Connection" binding:"required"`
}

// InitNode sets the defaults for a new node
func (p *Node) InitNode() {
	eventTime := time.Now().UTC()
	p.ID = bson.NewObjectId().Hex()
	p.DateCreated = eventTime
	p.DateModified = eventTime
	p.IsDeleted = false
}

// GetBaseConnectionInformation will attempt to locate the tenant by provided informaiton
func GetBaseConnectionInformation(collectionName string) (TenantConnectionInfo, error) {
	item := NewTenant("global")
	item.ID = "global"
	item.Name = "global"
	baseConnectionInfo := TenantConnectionInfo{}
	baseConnectionInfo.InitBaseTenantConnectionInfo()
	baseConnectionInfo.Tenant = item
	baseConnectionInfo.CollectionName = collectionName

	return baseConnectionInfo, nil
}
