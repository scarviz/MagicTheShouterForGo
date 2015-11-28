package magictheshouter

const (
	Nothing   = -1
	Ready     = 1000
	FireBall  = 1001
	Dragon    = 1002
	Begiragon = 1003
	Megante   = 1004
	IceStorm  = 1005
)

const (
	StrFireBall  = "ファイヤーボール"
	StrDragon    = "ドラゴン"
	StrBegiragon = "ベギラゴン"
	StrMegante   = "メガンテ"
	StrIceStorm  = "アイスストーム"
)

/*
音声データ
*/
type VoiceData struct {
	Data string `json:data`
}

/*
魔法データ
*/
type MagicData struct {
	DateTime int64 `json:"datetime"`
	MagicNo  int   `json:"magicno"`
}
