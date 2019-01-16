package types

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"errors"
	"time"
)

//表名 sc_partner

type ScPartner struct {
	Id           int    `json:"id"`
	CompanyName  string `json:"companyName"`
	CompanyLogo  string `json:"companyLogo"`
	CompanyIntro string `json:"companyIntro"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Weichat      string `json:"weichat"`
	Status       int    `json:"status"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type SqlScPartner struct {
	DBHandle *sql.DB
	Result   []ScPartner
	Row      int
	Data     *ScPartner
}

type JsonPartner struct {
	Message   string  `json:"message"`
	ErrCode	  int 	`json:"succFailCode"`
	Data      *[]ScPartner `json:"data"`
}

func (this *SqlScPartner) Query() error{
	this.Clear()
	rows, err := this.DBHandle.Query("select * from sc_partner")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScPartner
		err = rows.Scan(&result.Id,
			&result.CompanyName,
			&result.CompanyLogo,
			&result.CompanyIntro,
			&result.Name,
			&result.Phone,
			&result.Weichat,
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

func (this *SqlScPartner) QueryByCondition(condition string) error{
	this.Clear()
	sql := "select * from sc_partner "
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
		var result ScPartner
		err = rows.Scan(&result.Id,
			&result.CompanyName,
			&result.CompanyLogo,
			&result.CompanyIntro,
			&result.Name,
			&result.Phone,
			&result.Weichat,
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

func (this *SqlScPartner) Delete(condition string) error{
	if len(condition) == 0 {
		return errors.New("condition is null")
	}
	sql := "delete from sc_partner where "
	sql += condition
	fmt.Println(sql)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScPartner) UpdataState(id int) error {
	sql := fmt.Sprintf("update sc_partner set status=1 where id=%d", id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}


func (this *SqlScPartner) Update() error{
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("update sc_partner set company_name = '%s',company_logo = '%s',company_intro = '%s',"+
		"status = %d, create_time = '%s', update_time = '%s' "+
		"where id = %d",
		this.Data.CompanyName,
		this.Data.CompanyLogo,
		this.Data.CompanyIntro,
		this.Data.Status,
		t,
		t,
		this.Data.Id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}


func (this *SqlScPartner) Insert() error{
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("insert into sc_partner(company_name, company_logo, company_intro,"+
		"status, create_time, update_time) "+
		"values('%s','%s','%s', %d,'%s','%s')",
		this.Data.CompanyName,
		this.Data.CompanyLogo,
		this.Data.CompanyIntro,
		this.Data.Status,
		t,
		t)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScPartner) Clear() {
	if this.Row > 0 {
		this.Result = append(this.Result[this.Row:])
		this.Row = 0
	}
}

func (this *SqlScPartner) SerializeResult() (string,error){
	jsonPartner := &JsonPartner{
		ErrCode:2000,
		Message:"OK",
		Data:&this.Result,
	}
	retJson,err := json.Marshal(jsonPartner)
	if err != nil{
		fmt.Println(err.Error())
	}
	return string(retJson),err
}

func (this *SqlScPartner) UnSerializeJson(request string) error{
	err := json.Unmarshal([]byte(request),&this.Data)
	if err != nil{
		fmt.Println(err.Error())
	}
	return err
}