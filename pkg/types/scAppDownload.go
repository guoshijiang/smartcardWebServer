package types

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"errors"
	"time"
)

//表名 sc_app_download

type ScAppDownload struct {
	DownloadId   int    `json:"downloadId"`
	AndroidTitle string `json:"androidTitle"`
	AndroidImg   string `json:"androidImg"`
	AndroidUrl   string `json:"androidUrl"`
	IosTitle     string `json:"iosTitle"`
	IosImg       string `json:"iosImg"`
	IosUrl       string `json:"iosUrl"`
	Status       int    `json:"status"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type SqlScAppDownload struct {
	DBHandle *sql.DB
	Result   []ScAppDownload
	Row      int
	Data     *ScAppDownload
}

type JsonAppDownload struct {
	Message string           `json:"message"`
	ErrCode int              `json:"succFailCode"`
	Data    *[]ScAppDownload `json:"data"`
}

func (this *SqlScAppDownload) Query() error {
	this.Clear()
	rows, err := this.DBHandle.Query("select * from sc_app_download")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScAppDownload
		err = rows.Scan(&result.DownloadId,
			&result.AndroidTitle,
			&result.AndroidImg,
			&result.AndroidUrl,
			&result.IosTitle,
			&result.IosImg,
			&result.IosUrl,
			&result.Status,
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

func (this *SqlScAppDownload) QueryByCondition(condition string) error {
	this.Clear()
	sql := "select * from sc_app_download "
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
		var result ScAppDownload
		err = rows.Scan(&result.DownloadId,
			&result.AndroidTitle,
			&result.AndroidImg,
			&result.AndroidUrl,
			&result.IosTitle,
			&result.IosImg,
			&result.IosUrl,
			&result.Status,
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

func (this *SqlScAppDownload) Delete(condition string) error {
	if len(condition) == 0 {
		return errors.New("conditon is null")
	}
	sql := "delete from sc_app_download where "
	sql += condition
	fmt.Println(sql)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScAppDownload) UpdataState(id int) error {
	sql := fmt.Sprintf("update sc_app_download set status=1 where download_id = %d", id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScAppDownload) Update() error {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("update sc_app_download set android_title = '%s'"+
		",android_img = '%s',android_url = '%s',ios_title = '%s',ios_img = '%s',ios_url = '%s',"+
		"status = %d,create_time = '%s',update_time = '%s' "+
		"where download_id = %d",
		this.Data.AndroidTitle,
		this.Data.AndroidImg,
		this.Data.AndroidUrl,
		this.Data.IosTitle,
		this.Data.IosImg,
		this.Data.IosUrl,
		this.Data.Status,
		t,
		t,
		this.Data.DownloadId)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScAppDownload) Insert() error {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("insert into sc_app_download(android_title,android_img,android_url,"+
		"ios_title,ios_img,ios_url,status,create_time,update_time) "+
		"values('%s','%s','%s','%s','%s','%s',%d,'%s','%s')",
		this.Data.AndroidTitle,
		this.Data.AndroidImg,
		this.Data.AndroidUrl,
		this.Data.IosTitle,
		this.Data.IosImg,
		this.Data.IosUrl,
		this.Data.Status,
		t,
		t)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err

}

func (this *SqlScAppDownload) Clear() {
	if this.Row > 0 {
		this.Result = append(this.Result[this.Row:])
		this.Row = 0
	}
}

func (this *SqlScAppDownload) SerializeResult() (string, error) {
	jsonAppDownload := &JsonAppDownload{
		ErrCode: 2000,
		Message: "OK",
		Data:    &this.Result,
	}
	retJson, err := json.Marshal(jsonAppDownload)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(retJson), err
}

func (this *SqlScAppDownload) UnSerializeJson(request string) error {
	err := json.Unmarshal([]byte(request), &this.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
