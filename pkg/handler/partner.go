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

func QueryPartner(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScPartner  := types.SqlScPartner{
		DBHandle:db.GetDB(),
	}

	err := sqlScPartner.Query()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScPartner.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func AddPartner(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScPartner  := types.SqlScPartner{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	var data []byte

	_,err := r.Body.Read(data)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScPartner.UnSerializeJson(string(data))
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	err = sqlScPartner.Insert()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScPartner.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func DelPartner(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScPartner  := types.SqlScPartner{
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
		err = sqlScPartner.UpdataState(id)
		if err  != nil {
			glog.Errorf(err.Error())
			return
		}
	}

	jsonPartner,err := sqlScPartner.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonPartner)
	return
}


func UpdatePartner(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScPartner  := types.SqlScPartner{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScPartner.UnSerializeJson(string(data))
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScPartner.Update()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	jsonPartner,err := sqlScPartner.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonPartner)
	return
}
