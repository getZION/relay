package api

type Storage interface {
	CommunityService
	ConversationService
	MessageService
	PaymentService
	UserService
}

// Service represents a service for managing environment(endpoint) data.
type CommunityService interface {
	GetCommunities() ([]Community, error)
	GetCommunityByZid(string) (*Community, error)
	InsertCommunity(*Community) error

	AddUserToCommunity(communityZid, userDid string) error
	RemoveUserToCommunity(communityZid, userDid, leftReason string) error
}

type ConversationService interface {
	GetConversations() ([]Conversation, error)
	InsertConversation(*Conversation) error
}

type MessageService interface {
	GetMessages() ([]Message, error)
	InsertMessage(*Message) error
}

type PaymentService interface {
	GetPayments() ([]Payment, error)
	InsertPayment(*Payment) error
}

type UserService interface {
	GetUsers() ([]User, error)
	GetUserByDid(did string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	InsertUser(*User) error
}
