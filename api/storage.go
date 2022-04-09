package api

type Storage interface {
	CommunityService
	UserService
	ConversationService
	CommentService
	PaymentService
}

// Service represents a service for managing environment(endpoint) data.
type CommunityService interface {
	GetCommunities() ([]Community, error)
	GetCommunityByZid(string) (*Community, error)
	InsertCommunity(*Community) error

	AddUserToCommunity(communityZid, userDid string) error
	RemoveUserToCommunity(communityZid, userDid, leftReason string) error
}

// Service represents a service for managing environment(endpoint) data.
type UserService interface {
	GetUsers() ([]User, error)
	GetUserByDid(did string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	InsertUser(*User) error
}

type ConversationService interface {
	GetConversations() ([]Conversation, error)
	InsertConversation(*Conversation) error
}

type CommentService interface {
	GetComments() ([]Comment, error)
	InsertComment(*Comment) error
}

type PaymentService interface {
	GetPayments() ([]Payment, error)
	InsertPayment(*Payment) error
}
