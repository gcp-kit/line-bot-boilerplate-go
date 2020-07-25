package core

// environment variable key
const (
	EnvKeyBotName            = "BOT_NAME"
	EnvKeyChannelSecret      = "CHANNEL_SECRET"
	EnvKeyChannelAccessToken = "CHANNEL_ACCESS_TOKEN"
	EnvKeyMid                = "MID"
)

// TracerName - Events Name
type TracerName = string

// event Name
const (
	TracerTextMessage     TracerName = "TextMessage"
	TracerImageMessage    TracerName = "ImageMessage"
	TracerVideoMessage    TracerName = "VideoMessage"
	TracerAudioMessage    TracerName = "AudioMessage"
	TracerFileMessage     TracerName = "FileMessage"
	TracerLocationMessage TracerName = "LocationMessage"
	TracerStickerMessage  TracerName = "StickerMessage"

	TracerFollowEvent       TracerName = "follow"
	TracerUnfollowEvent     TracerName = "unfollow"
	TracerJoinEvent         TracerName = "join"
	TracerLeaveEvent        TracerName = "leave"
	TracerMemberJoinedEvent TracerName = "memberJoined"
	TracerMemberLeftEvent   TracerName = "memberLeft"
	TracerPostBackEvent     TracerName = "postback"
	TracerBeaconEvent       TracerName = "beacon"
	TracerAccountLinkEvent  TracerName = "accountLink"
	TracerThingsEvent       TracerName = "things"
)
