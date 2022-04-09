package api

const (
	ZION_JOIN_COMMUNITY      = "ZionJoinCommunity"
	ZION_LEAVE_COMMUNITY     = "ZionLeaveCommunity"
	ZION_COMMUNITY_KICK_USER = "ZionCommunityKickUser"
	ZION_COMMENT             = "ZionComment"

	COLLECTIONS_QUERY  = "CollectionsQuery"
	COLLECTIONS_WRITE  = "CollectionsWrite"
	COLLECTIONS_COMMIT = "CollectionsCommit"
	COLLECTIONS_DELETE = "CollectionsDelete"

	THREADS_QUERY  = "ThreadsQuery"
	THREADS_CREATE = "ThreadsCreate"
	THREADS_REPLY  = "ThreadsReply"
	THREADS_CLOSE  = "ThreadsClose"
	THREADS_DELETE = "ThreadsDelete"

	PERMISSIONS_REQUEST = "PermissionsRequest"
	PERMISSIONS_QUERY   = "PermissionsQuery"
	PERMISSIONS_GRANT   = "PermissionsGrant"
	PERMISSIONS_REVOKE  = "PermissionsRevoke"
)
