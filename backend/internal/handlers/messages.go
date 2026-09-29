package handlers

const (
	// Auth
	msgNotAuthenticated = "user not authenticated"
	msgInvalidLogin     = "invalid login"
	msgForbidden        = "forbidden"
	msgTokenFailed      = "could not create token"

	// Request-Validierung
	msgInvalidRequestBody = "invalid request body"
	msgInvalidDateFormat  = "invalid date format, expected ISO 8601"
	msgMissingQueryParam  = "missing query parameter 'q'"

	// Collection
	msgInvalidCollectionID     = "invalid collection id"
	msgCollectionIDRequired    = "collection id required"
	msgCollectionNotFound      = "collection not found"
	msgFetchCollections        = "error fetching collections"
	msgCreateCollection        = "could not create collection"
	msgCreateCollectionSuccess = "collection created successfully"
	msgUpdateCollection        = "could not update collection"

	// Item
	msgInvalidItemID = "invalid item id"
	msgItemNotFound  = "item not found"
	msgTitleRequired = "title is required"
	msgCreateItem    = "could not create item"
	msgDeleteItem    = "could not delete item"

	// User
	msgInvalidUserID      = "invalid user id"
	msgUserNotFound       = "user not found"
	msgEmailUsernameReq   = "email and username are required"
	msgEmailLength        = "email must be between 5 and 100 characters long"
	msgUsernameLength     = "username must be between 3 and 50 characters long"
	msgPasswordLength     = "password must be between 12 and 60 characters long"
	msgEmailTaken         = "email is already registered"
	msgInviteCodeRequired = "registration is disabled, an invite code is required"
	msgInvalidInviteCode  = "invalid invite code"
	msgCreateUser         = "could not create user"

	// Invite
	msgInvalidInviteID = "invalid invite id"
	msgCreateInvite    = "could not create invite code"
	msgInviteNotFound  = "invite code not found"
	msgFetchInvites    = "could not fetch invite codes"

	// Edition
	msgFetchEditions = "error fetching editions"

	// TVDB
	msgTVDBUnavailable = "tvdb request failed"

	// TMDB
	msgTMDBUnavailable = "tmdb request failed"

	// Discogs
	msgDiscogsUnavailable = "discogs request failed"

	// Settings
	msgInvalidSettingID = "invalid setting id"
	msgSettingNameReq   = "setting name is required"
	msgSettingNotFound  = "setting not found"
	msgSettingExists    = "a setting with this name already exists"
	msgFetchSettings    = "could not fetch settings"
	msgCreateSetting    = "could not create setting"
	msgUpdateSetting    = "could not update setting"
	msgDeleteSetting    = "could not delete setting"

	msgUserCreated       = "user created"
	msgLoginSuccessful   = "login successful"
	msgCollectionDeleted = "collection deleted"
	msgItemDeleted       = "item deleted"
	msgItemCreated       = "item created"
	msgInviteCreated     = "invite code created"
	msgInviteDeleted     = "invite code deleted"
	msgSettingCreated    = "setting created"
	msgSettingUpdated    = "setting updated"
	msgSettingDeleted    = "setting deleted"
)
