package magictheshouter

import (
	"appengine"
	"appengine/datastore"
)

/*
Magic情報を登録する
*/
func RegistMagic(context appengine.Context, data MagicData) error {
	k := datastore.NewIncompleteKey(context, "MagicData", nil)
	_, err := datastore.Put(context, k, &data)
	return err
}

/*
Magic情報を取得する
*/
func GetMagicData(context appengine.Context) (MagicData, error) {
	// DateTimeの降順に並べる
	q := datastore.NewQuery("MagicData").Order("-DateTime")

	var datas []MagicData
	keys, err := q.GetAll(context, &datas)

	var magicData MagicData
	if datas != nil && 0 < len(datas) {
		context.Debugf("[GetMagicData] MagicData: %v", len(datas))
		// 最新情報を優先する
		magicData = datas[0]

		if magicData.MagicNo != Ready {
			// 準備状態でなければすべて削除する
			err = datastore.DeleteMulti(context, keys)
		}
	}

	return magicData, err
}
