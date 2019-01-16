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

func QueryNews(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNews  := types.SqlScNews{
		DBHandle:db.GetDB(),
	}

	err := sqlScNews.Query()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScNews.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func AddNews(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNews  := types.SqlScNews{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	var data []byte

	_,err := r.Body.Read(data)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNews.UnSerializeJson(string(data))
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNews.Insert()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScNews.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func DelNews(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNews  := types.SqlScNews{
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
		err = sqlScNews.UpdataState(id)
		if err  != nil {
			glog.Errorf(err.Error())
			return
		}
	}

	jsonNews,err := sqlScNews.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonNews)
	return
}


func UpdateNews(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScNews  := types.SqlScNews{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNews.UnSerializeJson(string(data))
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScNews.Update()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	jsonNews,err := sqlScNews.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonNews)
	return
}
