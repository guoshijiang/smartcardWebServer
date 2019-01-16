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

func QueryAppDownload(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScAppDownload  := types.SqlScAppDownload{
		DBHandle:db.GetDB(),
	}

	err := sqlScAppDownload.Query()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScAppDownload.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func AddAppDownload(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScAppDownload  := types.SqlScAppDownload{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	var data []byte

	_,err := r.Body.Read(data)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScAppDownload.UnSerializeJson(string(data))
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	err = sqlScAppDownload.Insert()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScAppDownload.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func DelAppDownload(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScAppDownload  := types.SqlScAppDownload{
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
		err = sqlScAppDownload.UpdataState(id)
		if err  != nil {
			glog.Errorf(err.Error())
			return
		}
	}

	jsonAppDownload,err := sqlScAppDownload.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonAppDownload)
	return
}


func UpdateAppDownload(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScAppDownload  := types.SqlScAppDownload{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScAppDownload.UnSerializeJson(string(data))
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScAppDownload.Update()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	jsonAppDownload,err := sqlScAppDownload.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonAppDownload)
	return
}
