package api

type Community struct {
	Created         int64  `json:"created,omitempty" gorm:"not null"`
	Deleted         bool   `json:"deleted,omitempty" gorm:"default:false"`
	Description     string `json:"description" validate:"max=250" gorm:"size:250;not null"`
	EscrowAmount    int64  `json:"escrowAmount" validate:"gte=0,lt=100000" gorm:"not null"`
	Img             string `json:"img,omitempty"`
	LastActive      int64  `json:"lastActive,omitempty"`
	Name            string `json:"name" validate:"required,max=150" gorm:"size:150;unique;not null"`
	OwnerDid        string `json:"ownerDid" validate:"required"`
	PricePerMessage int64  `json:"pricePerMessage" validate:"gte=0,lt=100000" gorm:"not null"`
	PriceToJoin     int64  `json:"priceToJoin" validate:"gte=0,lt=100000" gorm:"not null"`
	Updated         int64  `json:"updated,omitempty"`
	Zid             string `json:"zid" gorm:"primary_key;unique;not null"`
}

type User struct {
	Bio      string `json:"bio,omitempty"`
	Created  int64  `json:"created,omitempty"`
	Did      string `json:"did" validate:"required" gorm:"primary_key;unique;not null"`
	Img      string `json:"img,omitempty"`
	Link     string `json:"link,omitempty"`
	Location string `json:"location,omitempty"`
	Name     string `json:"name" validate:"required" gorm:"not null"`
	Updated  int64  `json:"updated,omitempty"`
	Username string `json:"username" validate:"required,username,min=6,max=16" gorm:"unique;not null"`
}

type Comment struct{}
type Conversation struct{}
type Payment struct{}
