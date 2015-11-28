package magictheshouter

const (
	Nothing      = -1
	Ready        = 1000
	FireBall     = 1001
	Dragon       = 1002
	Begiragon    = 1003
	Megante      = 1004
	IceStorm     = 1005
	Raidein      = 1006
	ThunderStorm = 1007
	ScrewWave    = 1008
	BigWave      = 1009
	Gigadoriru   = 1010
	Kurushio     = 1011
	TigerBazooka = 1012
	Paropunte    = 1013
	SJK          = 1014
	FinalDragoon = 1015
	RocknRoll    = 1016
	RainField    = 1017
)

const (
	StrFireBall     = "ファイヤーボール"
	StrDragon       = "ドラゴン"
	StrBegiragon    = "ベギラゴン"
	StrMegante      = "メガンテ"
	StrIceStorm     = "アイスストーム"
	StrRaidein      = "ライデイン"
	StrThunderStorm = "サンダーストーム"
	StrScrewWave    = "スクリューウェーブ"
	StrBigWave      = "ビッグウェーブ"
	StrGigadoriru   = "ギガドリル"
	StrKurushio     = "クルーシオ"
	StrTigerBazooka = "タイガーバズーカ"
	StrParopunte    = "パロプンテ"
	StrSJK          = "スリジャヤワルダナプラコッテ"
	StrFinalDragoon = "ファイナルドラグーン"
	StrRocknRoll    = "ロックンロール"
	StrRainField    = "レインフィールド"
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
