package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/golang/glog"
	"smartcardWebServer/pkg/types"
	"smartcardWebServer/pkg/db"
	"net/http"
	"io"
	"encoding/json"
	"strings"
	"strconv"
	"io/ioutil"
)

type Id struct {
	Id  string `json:"noticeId"`
}

func QueryNotice(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNotic  := types.SqlScNotic{
		DBHandle:db.GetDB(),
	}

	err := sqlScNotic.Query()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScNotic.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func AddNotice(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNotic  := types.SqlScNotic{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	var data []byte

	_,err := r.Body.Read(data)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNotic.UnSerializeJson(string(data))
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNotic.Insert()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScNotic.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func DelNotice(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNotic  := types.SqlScNotic{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	var noticeId Id

	err = json.Unmarshal(data,&noticeId)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	arrId := strings.Split(noticeId.Id,",")
	for _,v :=range  arrId{
		id,err := strconv.Atoi(v)
		err = sqlScNotic.UpdataState(id)
		if err  != nil {
			glog.Errorf(err.Error())
			return
		}
	}

	jsonNotic,err := sqlScNotic.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonNotic)
	return
}


func UpdateNotice(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNotic  := types.SqlScNotic{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNotic.UnSerializeJson(string(data))
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNotic.Update()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	jsonNotic,err := sqlScNotic.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonNotic)
	return
}
