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

func QueryTeamIntroduce(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScTeamIntroduce  := types.SqlScTeamIntroduce{
		DBHandle:db.GetDB(),
	}

	err := sqlScTeamIntroduce.Query()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScTeamIntroduce.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func AddTeamIntroduce(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScTeamIntroduce  := types.SqlScTeamIntroduce{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	var data []byte

	_,err := r.Body.Read(data)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScTeamIntroduce.UnSerializeJson(string(data))
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	err = sqlScTeamIntroduce.Insert()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	retJson,err := sqlScTeamIntroduce.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,retJson)
	return
}

func DelTeamIntroduce(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScTeamIntroduce  := types.SqlScTeamIntroduce{
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
		err = sqlScTeamIntroduce.UpdataState(id)
		if err  != nil {
			glog.Errorf(err.Error())
			return
		}
	}

	jsonTeamIntroduce,err := sqlScTeamIntroduce.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonTeamIntroduce)
	return
}


func UpdateTeamIntroduce(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//PersonInfo, err := types.GetNoticeInfo()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sqlScTeamIntroduce  := types.SqlScTeamIntroduce{
		DBHandle:db.GetDB(),
	}

	defer r.Body.Close()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScTeamIntroduce.UnSerializeJson(string(data))
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	err = sqlScTeamIntroduce.Update()
	if err  != nil {
		glog.Errorf(err.Error())
		return
	}

	jsonTeamIntroduce,err := sqlScTeamIntroduce.SerializeResult()
	if err != nil {
		glog.Errorf(err.Error())
		return
	}

	io.WriteString(w,jsonTeamIntroduce)
	return
}
