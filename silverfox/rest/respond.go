package rest

import (
	"encoding/json"
	"net/http"
)

// HALを適用したかったけど今回は見送り。簡易的なjosnレスポンス。
type Response struct {
	Entry   interface{} `json:"entity"`
	Message string      `json:"message"`
}

// 今回はすべてjsonで返却
func Respond(w http.ResponseWriter, status int, data Response) {
	// サンプル用なのですべて許可。公開APIでもない限り通常はドメインを絞る。
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Location")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
