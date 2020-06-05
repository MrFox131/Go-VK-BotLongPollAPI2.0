package vkapi

var typeSpecifier = make(map[string]interface{})

//MESSAGETYPINGSTATE_START

//MessageTypingState representes event of message_typing_state type
type MessageTypingState struct {
	State  string `mapstructure:"state"`
	FromID int64  `mapstructure:"from_id"`
	ToID   int64  `mapstructure:"to_id"`
}

//MESSAGETYPINGSTATE_END

//MESSAGENEW_START

//ClientInfoStruct struct of user's device available functions. Comes with NewMessage struct in ClientInfo field
type ClientInfoStruct struct {
	ButtonActions  []string `mapstructure:"button_actions"`
	Keyboard       bool     `mapstructure:"keyboard"`
	InlineKeyboard bool     `mapstructure:"inline_keyboard"`
	Carousel       bool     `mapstructure:"carousel"`
	LangID         int64    `mapstructure:"lang_id"`
}

// type AttachementsStruct struct {
// }

//GeoStruct struct with latitide and longtitude of place and information about it
type GeoStruct struct {
	Type        string `mapstructure:"type"`
	Coordinates struct {
		Latitude   float64 `mapstructure:"latitude"`
		Longtitude float64 `mapstructure:"longtitude"`
	} `mapstructure:"coordinates"`
	Place struct {
		ID         int64   `mapstructure:"id"`
		Title      string  `mapstructure:"title"`
		Latitude   float64 `mapstructure:"latitude"`
		Longtitude float64 `mapstructure:"longtitude"`
		Created    int64   `mapstructure:"created"`
		Icon       string  `mapstructure:"icon"`
		Country    string  `mapstructure:"country"`
		City       string  `mapstructure:"city"`
	} `mapstructure:"place"`
}

// type ButtonStruct struct {

// }

//KeyboardStruct is struct with information about keyboard on user's side
type KeyboardStruct struct {
	OneTime bool          `mapstructure:"one_time"`
	Buttons []interface{} `mapstructure:"buttons"`
	Inline  bool          `mapstructure:"inline"`
}

// type ActionStruct struct {
// }

//MessageStruct struct of message for message_new event
type MessageStruct struct {
	ID        int64  `mapstructure:"id"`
	Date      int64  `mapstructure:"date"`
	PeerID    int64  `mapstructure:"peer_id"`
	FromID    int64  `mapstructure:"from_id"`
	Text      string `mapstructure:"text"`
	RandomID  int64  `mapstructure:"random_id"`
	Ref       string `mapstructure:"ref"`
	RefSource string `mapstructure:"ref_source"`
	//Attachments  AttachementsStruct `mapstructure:"attachments"`
	Attachments  interface{}     `mapstructure:"attachments"`
	Important    bool            `mapstructure:"important"`
	Geo          GeoStruct       `mapstructure:"geo"`
	Payload      string          `mapstructure:"payload"`
	Keyboard     KeyboardStruct  `mapstructure:"keyboard"`
	FwdMessages  []MessageStruct `mapstructure:"fwd_messages"`
	ReplyMessage *MessageStruct  `mapstructure:"reply_message"`
	//Action       ActionStruct    `mapstructure:"action"`
	Action interface{} `mapstructure:"action"`
}

//MessageNew struct of new message event
type MessageNew struct {
	Message    MessageStruct    `mapstructure:"message"`
	ClientInfo ClientInfoStruct `mapstructure:"client_info"`
}

//MESSAGENEW_END

//MESSAGEREPLY_START

//MessageReply is struct for message_reply event
type MessageReply struct {
	Message MessageStruct `mapstructure:"message"`
}

//MESSAGEREPLY_END

//MESSAGEEDIT_START

//MessageEdit is struct for message_edit event
type MessageEdit struct {
	Message MessageStruct `mapstructure:"message"`
}

//MESSAGEEDIT_END

//MESSAGEALLOW_START

//MessageAllow is struct for message_allow event
type MessageAllow struct {
	UserID int64  `mapstructure:"user_id"`
	Key    string `mapstructure:"key"`
}

//MESSAGEALLOW_END

//MESSAGEDENY_START

//MessageDeny is struct for message_deny event
type MessageDeny struct {
	UserID int64 `mapstructure:"user_id"`
}

//MESSAGE_END

//PHOTONEW_STRAT

//PhotoNew is struct for photo_new event
type PhotoNew struct {
	ID      int64  `mapstructure:"id"`
	AlbumID int64  `mapstructure:"album_id"`
	OwnerID int64  `mapstructure:"owner_id"`
	UserID  int64  `mapstructure:"user_id"`
	Text    string `mapstructure:"text"`
	Date    int64  `mapstructure:"date"`
	Sizes   struct {
		Type   string `mapstructure:"type"`
		URL    string `mapstructure:"url"`
		Width  int64  `mapstructure:"width"`
		Height int64  `mapstructure:"height"`
	} `mapstructure:"sizes"`
	Width  int64 `mapstructure:"width*"`
	Height int64 `mapstructure:"height*"`
}

//PHOTONEW_END

//PHOTOCOMMENTNEW/EDIT/RESTORE_START

//CommentStruct is struct for information about comment
type CommentStruct struct {
	ID             int64       `mapstructure:"id"`
	FromID         int64       `mapstructure:"from_id"`
	Date           int64       `mapstructure:"date"`
	Text           string      `mapstructure:"text"`
	ReplyToUser    int64       `mapstructure:"reply_to_user"`
	ReplyToComment int64       `mapstructure:"reply_to_comment"`
	Attachments    interface{} `mapstructure:"attachments"`
	ParantsStack   []int64     `mapstructure:"parants_stack"`
	Thread         struct {
		Count           int64           `mapstructure:"count"`
		Items           []CommentStruct `mapstructure:"items"`
		CanPost         bool            `mapstructure:"can_post"`
		ShowReplyButton bool            `mapstructure:"show_reply_button"`
		GroupsCanPost   bool            `mapstructure:"groups_can_post"`
	} `mapstructure:"thread"`
}

//PhotoCommentNew is struct for event photo_comment_new
type PhotoCommentNew struct {
	PhotoID        int64       `mapstructure:"photo_id"`
	PhotoOwnerID   int64       `mapstructure:"photo_owner_id"`
	ID             int64       `mapstructure:"id"`
	FromID         int64       `mapstructure:"from_id"`
	Date           int64       `mapstructure:"date"`
	Text           string      `mapstructure:"text"`
	ReplyToUser    int64       `mapstructure:"reply_to_user"`
	ReplyToComment int64       `mapstructure:"reply_to_comment"`
	Attachments    interface{} `mapstructure:"attachments"`
	ParantsStack   []int64     `mapstructure:"parants_stack"`
	Thread         struct {
		Count           int64           `mapstructure:"count"`
		Items           []CommentStruct `mapstructure:"items"`
		CanPost         bool            `mapstructure:"can_post"`
		ShowReplyButton bool            `mapstructure:"show_reply_button"`
		GroupsCanPost   bool            `mapstructure:"groups_can_post"`
	} `mapstructure:"thread"`
}

//PhotoCommentEdit is struct for event photo_comment_edit
type PhotoCommentEdit = PhotoCommentNew

//PhotoCommentRestore is struct for event photo_comment_restore
type PhotoCommentRestore = PhotoCommentNew

//PHOTOCOMMENTNEW/EDIT/RESTORE_END

//PHOTOCOMMENTDELETE_START

//PhotoCommentDelete is struct for event photo_comment_delete
type PhotoCommentDelete struct {
	OwnerID   int64 `mapstructure:"owner_id"`
	ID        int64 `mapstructure:"id"`
	UserID    int64 `mapstructure:"user_id"`
	DeleterID int64 `mapstructure:"deleter_id"`
	PhotoID   int64 `mapstructure:"photo_id"`
}

//PHOTOCOMMENT_END
