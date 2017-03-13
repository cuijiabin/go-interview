package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func DetailRecite(w http.ResponseWriter, req *http.Request) {
	id := req.PostFormValue("id")
	recite := GetReciteById(id)
	f := map[string]interface{}{
		"code":   0,
		"recite": recite,
	}
	json.NewEncoder(w).Encode(f)
}

func ReciteList(w http.ResponseWriter, req *http.Request) {
	page := PageRecite{Cp: 1, Ps: 10, tC: 0, tP: 0}
	page = PageMgoRecite(page)
	w.Header().Set("Content-Type", "application/json")
	f := map[string]interface{}{
		"Cp":   page.Cp,
		"Ps":   page.Ps,
		"tC":   page.tC,
		"tP":   page.tP,
		"list": page.list,
	}
	json.NewEncoder(w).Encode(f)
}

func AddRecite(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	content := req.PostFormValue("content")
	tip := req.PostFormValue("tip")
	r := Recite{Title: title, Content: content, Tip: tip}
	result := AddMgoRecite(r)
	f := map[string]interface{}{
		"id":   result,
		"code": 0,
	}
	json.NewEncoder(w).Encode(f)
}

func EditRecite(w http.ResponseWriter, req *http.Request) {
	id := req.PostFormValue("id")
	title := req.PostFormValue("title")
	content := req.PostFormValue("content")
	tip := req.PostFormValue("tip")
	r := Recite{Title: title, Content: content, Tip: tip}
	result := EditMgoRecite(id, r)
	f := map[string]interface{}{
		"id":   result,
		"code": 0,
	}
	json.NewEncoder(w).Encode(f)

}

func AddRepeat(w http.ResponseWriter, req *http.Request) {

	rId := req.PostFormValue("rId")
	recite := GetReciteById(rId)
	if recite == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": -1, "msg": "找不到背诵的内容"})
		return
	}
	rpContent := req.PostFormValue("rpContent")
	remark := req.PostFormValue("remark")
	isCorrect := strings.EqualFold(recite.Content, rpContent)

	re := Repeat{RpContent: rpContent, Remark: remark, IsCorrect: isCorrect}
	rpId := AddMgoRepeat(re, rId)

	StatMgoRecite(rId)
	resultMap := map[string]interface{}{
		"id":   rpId,
		"code": 0,
	}
	json.NewEncoder(w).Encode(resultMap)

}

func RepeatList(w http.ResponseWriter, req *http.Request) {
	rId := req.PostFormValue("rId")
	list := ListMgoRepeat(rId)
	json.NewEncoder(w).Encode(list)
}
