package magictheshouter

import (
	"appengine"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/cancel", cancel)

	http.HandleFunc("/regist", regist)
	http.HandleFunc("/magic", magic)

	http.HandleFunc("/testdata", testdata)
}

/*
魔法準備
*/
func ready(w http.ResponseWriter, r *http.Request) {
	changeReady(w, r, Ready)
}

/*
魔法取り消し
*/
func cancel(w http.ResponseWriter, r *http.Request) {
	changeReady(w, r, Nothing)
}

/*
魔法準備状態の変更
*/
func changeReady(w http.ResponseWriter, r *http.Request, magicNo int) {
	c := appengine.NewContext(r)
	r.ParseForm()
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	magicData := MagicData{
		DateTime: time.Now().Unix(),
		MagicNo:  magicNo,
	}

	// 登録
	regerr := RegistMagic(c, magicData)
	if regerr != nil {
		c.Debugf("[changeReady] regist err: %v", regerr)
		http.Error(w, regerr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, strconv.Itoa(http.StatusOK))
}

/*
Masic情報を登録する
*/
func regist(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	r.ParseForm()
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	// リクエスト内容を取得
	data, readerr := ioutil.ReadAll(r.Body)
	if readerr != nil {
		c.Debugf("[regist] read body err: %v", readerr)
		http.Error(w, readerr.Error(), http.StatusInternalServerError)
		return
	}

	// リクエストJSON
	var voiceData VoiceData
	jsonerr := json.Unmarshal(data, &voiceData)
	if jsonerr != nil {
		c.Debugf("[regist] json err: %v", jsonerr)
		http.Error(w, jsonerr.Error(), http.StatusInternalServerError)
		return
	}

	magicData := MagicData{
		DateTime: time.Now().Unix(),
		MagicNo:  GetMagicNo(voiceData.Data),
	}

	// 登録
	regerr := RegistMagic(c, magicData)
	if regerr != nil {
		c.Debugf("[regist] regist err: %v", regerr)
		http.Error(w, regerr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, strconv.Itoa(http.StatusOK))
}

/*
Magic情報を取得する
*/
func magic(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	magicData, err := GetMagicData(c)

	if err != nil {
		c.Debugf("[magic] GetMagicData err: %v", err)
		w.Header().Add("Content-type", "text/html; charset=utf-8")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resjson, jsonerr := json.Marshal(magicData)

	if jsonerr != nil {
		c.Debugf("[magic] json err: %v", jsonerr)
		w.Header().Add("Content-type", "text/html; charset=utf-8")
		http.Error(w, jsonerr.Error(), http.StatusInternalServerError)
		return
	}

	// 結果を返す
	w.Header().Add("Content-type", "application/json")
	fmt.Fprintf(w, string(resjson))
}

/*
Magic番号を取得する
*/
func GetMagicNo(data string) (magicNo int) {
	reFB, _ := regexp.Compile(StrFireBall)
	reDR, _ := regexp.Compile(StrDragon)
	reBG, _ := regexp.Compile(StrBegiragon)
	reMG, _ := regexp.Compile(StrMegante)
	reIS, _ := regexp.Compile(StrIceStorm)

	fb := reFB.FindAllStringSubmatch(data, -1)
	dr := reDR.FindAllStringSubmatch(data, -1)
	bg := reBG.FindAllStringSubmatch(data, -1)
	mg := reMG.FindAllStringSubmatch(data, -1)
	is := reIS.FindAllStringSubmatch(data, -1)

	if 0 < len(fb) {
		magicNo = FireBall
	} else if 0 < len(dr) {
		magicNo = Dragon
	} else if 0 < len(bg) {
		magicNo = Begiragon
	} else if 0 < len(mg) {
		magicNo = Megante
	} else if 0 < len(is) {
		magicNo = IceStorm
	} else {
		magicNo = Nothing
	}

	return
}

func testdata(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	r.ParseForm()
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	magicData1 := MagicData{
		DateTime: time.Now().Unix(),
		MagicNo:  GetMagicNo("ファイヤーボール"),
	}

	magicData2 := MagicData{
		DateTime: time.Now().Unix() + int64(10),
		MagicNo:  GetMagicNo("ドラゴン！！！"),
	}

	// 登録
	regerr := RegistMagic(c, magicData1)
	regerr = RegistMagic(c, magicData2)
	if regerr != nil {
		c.Debugf("[regist] regist err: %v", regerr)
		http.Error(w, regerr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, strconv.Itoa(http.StatusOK))
}
