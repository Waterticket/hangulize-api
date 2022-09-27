package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hangulize/hangulize"
)

type PronouncePacket struct {
	Id        int64
	Original  string
	Pronounce string
}

type RequestData struct {
	Id   int64  `json:"id"`
	Kana string `json:"kana"`
}

type RequestPacket struct {
	Data []RequestData `json:"data"`
}

func main() {
	http.HandleFunc("/pronounciation/j2k/group", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var data RequestPacket
		json.Unmarshal(body, &data)

		var list = data.Data
		Packet := make([]PronouncePacket, len(list))

		for cnt, item := range list {
			han := hangulize.Hangulize("jpn-ck", item.Kana)
			Packet[cnt].Id = item.Id
			Packet[cnt].Original = item.Kana
			Packet[cnt].Pronounce = han
		}
		doc, _ := json.Marshal(Packet)

		w.Write([]byte(doc))
	})

	http.HandleFunc("/pronounciation/j2k/solo", func(w http.ResponseWriter, req *http.Request) {
		item := req.FormValue("q")
		han := hangulize.Hangulize("jpn-ck", item)
		w.Write([]byte(han))
	})

	http.ListenAndServe(":5000", nil)
}
