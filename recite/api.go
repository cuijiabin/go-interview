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

func DelRecite(w http.ResponseWriter, req *http.Request) {
	id := req.PostFormValue("id")
	list := ListMgoRepeat(id)
	if len(list) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"msg":  "还有背诵没有删除",
			"code": -1,
		})
		return
	}

	result := DelMgoRecite(id)
	f := map[string]interface{}{
		"msg":  result,
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

func DetailRepeat(w http.ResponseWriter, req *http.Request) {

	id := req.PostFormValue("id")
	repeat := GetRepeatById(id)
	f := map[string]interface{}{
		"code":   0,
		"repeat": repeat,
	}
	json.NewEncoder(w).Encode(f)

}

func DelRepeat(w http.ResponseWriter, req *http.Request) {

	id := req.PostFormValue("id")
	result := DelRepeatById(id)
	f := map[string]interface{}{
		"msg":  result,
		"code": 0,
	}
	json.NewEncoder(w).Encode(f)

}

func EditRepeat(w http.ResponseWriter, req *http.Request) {

	id := req.PostFormValue("id")
	remark := req.PostFormValue("remark")
	result := EditMgoRepeat(id, remark)
	f := map[string]interface{}{
		"id":   result,
		"code": 0,
	}
	json.NewEncoder(w).Encode(f)

}

func RepeatList(w http.ResponseWriter, req *http.Request) {
	rId := req.PostFormValue("rId")
	list := ListMgoRepeat(rId)
	json.NewEncoder(w).Encode(list)
}

//TODO 标签列表
func LabelList(w http.ResponseWriter, req *http.Request) {
	list := ListMgoLabel()
	json.NewEncoder(w).Encode(list)
}
//TODO 添加标签
func AddLabel(w http.ResponseWriter, req *http.Request) {
	content := req.PostFormValue("content")
	l := Label{Name: content}
	result := AddMgoLabel(l)
	f := map[string]interface{}{
		"id":   result,
		"code": 0,
	}
	json.NewEncoder(w).Encode(f)
}
//TODO 删除标签
//TODO 修改标签
//TODO 标签详情
