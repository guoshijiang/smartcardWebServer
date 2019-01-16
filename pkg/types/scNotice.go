package types

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"encoding/json"
	"errors"
	"time"
)

type ScNotic struct {
	NoticeId      int    `json:"noticeId"`
	NoticeTitle   string `json:"noticeTitle"`
	NoticeContent string `json:"noticeContent"`
	NoticeImg     string `json:"noticeImg"`
	Status        int    `json:"status"`
	NoticeRemark  string `json:"noticeRemark"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

type SqlScNotic struct {
	DBHandle *sql.DB
	Result   []ScNotic
	Row      int
	Data     ScNotic
}

type JsonNotic struct {
	Message string     `json:"message"`
	ErrCode int        `json:"succFailCode"`
	Data    *[]ScNotic `json:"data"`
}

func (this *SqlScNotic) Query() error {
	this.Clear()
	rows, err := this.DBHandle.Query("select * from sc_notice")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScNotic
		err = rows.Scan(&result.NoticeId,
			&result.NoticeTitle,
			&result.NoticeContent,
			&result.NoticeImg,
			&result.Status,
			&result.NoticeRemark,
			&result.CreateTime,
			&result.UpdateTime)
		if err != nil {
			fmt.Println(err.Error())
		}
		this.Result = append(this.Result, result)
		this.Row++
	}
	return err
}

func (this *SqlScNotic) QueryByCondition(condition string) error {
	this.Clear()
	sql := "select * from sc_notice "
	if len(condition) > 0 {
		sql += "where "
		sql += condition
	}
	rows, err := this.DBHandle.Query(sql)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScNotic
		err = rows.Scan(&result.NoticeId,
			&result.NoticeTitle,
			&result.NoticeContent,
			&result.NoticeImg,
			&result.Status,
			&result.NoticeRemark,
			&result.CreateTime,
			&result.UpdateTime)
		if err != nil {
			fmt.Println(err.Error())
		}
		this.Result = append(this.Result, result)
		this.Row++
	}
	return nil
}

func (this *SqlScNotic) Delete(condition string) error {
	if len(condition) == 0 {
		return errors.New("condtion is NULL")
	}
	sql := "delete from sc_notice where "
	sql += condition
	fmt.Println(sql)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNotic) UpdataState(id int) error {
	sql := fmt.Sprintf("update sc_notice set status=1 where notice_id=%d", id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNotic) Update() error{
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("update sc_notice set notice_title = '%s',notice_content = '%s'," +
		"notice_img = '%s',status = %d,notice_remark = '%s',create_time = '%s',update_time = '%s' "+
			"where notice_id = %d",
		this.Data.NoticeTitle,
		this.Data.NoticeContent,
		this.Data.NoticeImg,
		this.Data.Status,
		this.Data.NoticeRemark,
		t,
		t,
		this.Data.NoticeId)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNotic) Insert() error {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("insert into sc_notice(notice_title,notice_content,notice_img,status,"+
		"notice_remark,create_time,update_time) "+
		"values('%s','%s','%s',%d,'%s','%s','%s')",
		this.Data.NoticeTitle,
		this.Data.NoticeContent,
		this.Data.NoticeImg,
		this.Data.Status,
		this.Data.NoticeRemark,
		t,
		t)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNotic) Clear() {
	if this.Row > 0 {
		this.Result = append(this.Result[this.Row:])
		this.Row = 0
	}
}

func (this *SqlScNotic) SerializeResult() (string, error) {
	jsonNotice := &JsonNotic{
		ErrCode: 2000,
		Message: "OK",
		Data:    &this.Result,
	}
	retJson, err := json.Marshal(jsonNotice)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(retJson), err
}

func (this *SqlScNotic) UnSerializeJson(request string) error {
	err := json.Unmarshal([]byte(request), &this.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
