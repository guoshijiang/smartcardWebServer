package types

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"errors"
	"encoding/json"
	"time"
)

type ScNews struct {
	NewsId      int    `json:"newsId"`
	NewsTitle   string `json:"newsTitle"`
	NewsContent string `json:"newsContent"`
	NewsImg     string `json:"newsImg"`
	Status      int    `json:"status"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type SqlScNews struct {
	DBHandle *sql.DB
	Result   []ScNews
	Row      int
	Data     *ScNews
}

type JsonNews struct {
	Message string    `json:"message"`
	ErrCode int       `json:"succFailCode"`
	Data    *[]ScNews `json:"data"`
}

func (this *SqlScNews) Query() error {
	this.Clear()
	rows, err := this.DBHandle.Query("select * from sc_news")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScNews
		err = rows.Scan(&result.NewsId,
			&result.NewsTitle,
			&result.NewsContent,
			&result.NewsImg,
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

func (this *SqlScNews) QueryByCondition(condition string) error {
	this.Clear()
	sql := "select * from sc_news where status = 0"
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
		var result ScNews
		err = rows.Scan(&result.NewsId,
			&result.NewsTitle,
			&result.NewsContent,
			&result.NewsImg,
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

func (this *SqlScNews) Delete(condition string) error {
	if len(condition) == 0 {
		return errors.New("condition is null")
	}
	sql := "delete from sc_news where "
	sql += condition
	fmt.Println(sql)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNews) UpdataState(id int) error {
	fmt.Println("id = ", id)
	sql := fmt.Sprintf("update sc_news set status = 1 where news_id = %d",  id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNews) Update() error{
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("update  sc_news set news_title = '%s',news_content = '%s'," +
		"news_img = '%s',status = %d, create_time = '%s', update_time = '%s' "+
		"where news_id = %d",
		this.Data.NewsTitle,
		this.Data.NewsContent,
		this.Data.NewsImg,
		this.Data.Status,
		t,
		t,
		this.Data.NewsId)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNews) Insert() error {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("insert into sc_news(news_title,news_content,news_img,status,"+
		"create_time,update_time) "+
		"values('%s','%s','%s',%d,'%s','%s')",
		this.Data.NewsTitle,
		this.Data.NewsContent,
		this.Data.NewsImg,
		this.Data.Status,
		t,
		t)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScNews) Clear() {
	if this.Row > 0 {
		this.Result = append(this.Result[this.Row:])
		this.Row = 0
	}
}

func (this *SqlScNews) SerializeResult() (string, error) {
	jsonNews := &JsonNews{
		ErrCode: 2000,
		Message: "OK",
		Data:    &this.Result,
	}
	retJson, err := json.Marshal(jsonNews)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(retJson), err
}

func (this *SqlScNews) UnSerializeJson(request string) error {
	err := json.Unmarshal([]byte(request), &this.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
