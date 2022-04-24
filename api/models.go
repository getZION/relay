package api

type Community struct {
	Created         int64     `json:"created,omitempty" gorm:"not null"`
	Deleted         bool      `json:"deleted,omitempty" gorm:"default:false"`
	Description     string    `json:"description" validate:"max=250" gorm:"size:250;not null"`
	EscrowAmount    int64     `json:"escrowAmount" validate:"gte=0,lt=100000" gorm:"not null"`
	Img             string    `json:"img,omitempty"`
	LastActive      int64     `json:"lastActive,omitempty"`
	Messages        []Message `json:"messages,omitempty"`
	Name            string    `json:"name" validate:"required,max=150" gorm:"size:150;unique;not null"`
	OwnerDid        string    `json:"ownerDid" validate:"required"`
	PricePerMessage int64     `json:"pricePerMessage" validate:"gte=0,lt=100000" gorm:"not null"`
	PriceToJoin     int64     `json:"priceToJoin" validate:"gte=0,lt=100000" gorm:"not null"`
	Updated         int64     `json:"updated,omitempty"`
	Users           []User    `json:"users,omitempty" gorm:"many2many:user_communities"`
	Zid             string    `json:"zid" gorm:"primary_key;unique;not null"`
}

type Conversation struct {
	CommunityZid string    `json:"communityZid" validate:"required" gorm:"not null"`
	Created      int64     `json:"created,omitempty" gorm:"not null"`
	Deleted      Bool      `json:"deleted,omitempty" gorm:"default:false"`
	Img          string    `json:"img,omitempty"`
	Link         string    `json:"link,omitempty" validate:"required_without=Text"`
	Messages     []Message `json:"messages,omitempty"`
	Text         string    `json:"text,omitempty" validate:"required_without=Link"`
	Title        string    `json:"title,omitempty"`
	Updated      int64     `json:"updated,omitempty"`
	UserDid      string    `json:"userDid" validate:"required"`
	Zid          string    `json:"zid" gorm:"primary_key;unique;not null"`
}

type Message struct {
	ConversationZid   string `json:"conversationZid,omitempty" gorm:"size:191"`
	CommunityZid      string `json:"communityZid,omitempty" gorm:"size:191"`
	Created           int64  `json:"created,omitempty" gorm:"not null"`
	Deleted           Bool   `json:"deleted,omitempty" gorm:"default:false"`
	Img               string `json:"img,omitempty"`
	Link              string `json:"link,omitempty" validate:"required_without=Text"`
	ReceivingUserDid  string `json:"receivingUserDid,omitempty"`
	ReplyToMessageZid string `json:"replyToMessageZid,omitempty"`
	Text              string `json:"text,omitempty" validate:"required_without=Link"`
	Updated           int64  `json:"updated,omitempty"`
	UserDid           string `json:"userDid,omitempty"`
	Zid               string `json:"zid" gorm:"primary_key;unique;not null"`
}

type Payment struct {
	Amount              int64  `json:"amount,omitempty"`
	Memo                string `json:"memo,omitempty"`
	MessageZid          string `json:"messageZid,omitempty"`
	RecipientDid        string `json:"recipientDid,omitempty"`
	RecipientNodePubkey string `json:"recipientNodePubkey,omitempty"`
	RecipientRelayUrl   string `json:"recipientRelayUrl,omitempty"`
	SenderDid           string `json:"senderDid,omitempty"`
	Status              string `json:"status,omitempty"`
	Type                int64  `json:"type,omitempty"`
	Zid                 string `json:"zid" gorm:"primary_key;unique;not null"`
}

type User struct {
	Bio         string      `json:"bio,omitempty"`
	Communities []Community `json:"communities,omitempty" gorm:"many2many:user_communities"`
	Created     int64       `json:"created,omitempty"`
	Did         string      `json:"did" validate:"required" gorm:"primary_key;unique;not null"`
	Img         string      `json:"img,omitempty"`
	ImgCover    string      `json:"imgCover,omitempty"`
	Link        string      `json:"link,omitempty"`
	Location    string      `json:"location,omitempty"`
	Name        string      `json:"name" validate:"required" gorm:"not null"`
	Updated     int64       `json:"updated,omitempty"`
	Username    string      `json:"username" validate:"required,username,min=6,max=16" gorm:"unique;not null"`
}

type JoinCommunity struct {
	CommunityZid string
	UserDid      string
}

type LeaveCommunity struct {
	CommunityZid string
	UserDid      string
}
