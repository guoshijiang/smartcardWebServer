package types

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

type ScTeamIntroduce struct {
	MemberId    int    `json:"memberId"`
	MemberName  string `json:"memberName"`
	MemberIntro string `json:"memberIntro"`
	MemberImg   string `json:"image"`
	Status      int    `json:"status"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type SqlScTeamIntroduce struct {
	DBHandle *sql.DB
	Result   []ScTeamIntroduce
	Row      int
	Data     *ScTeamIntroduce
}

type JsonTeamIntroduce struct {
	Message string             `json:"message"`
	ErrCode int                `json:"succFailCode"`
	Data    *[]ScTeamIntroduce `json:"data"`
}

func (this *SqlScTeamIntroduce) Query() error {
	this.Clear()
	rows, err := this.DBHandle.Query("select * from sc_team_introduce")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		var result ScTeamIntroduce
		err = rows.Scan(&result.MemberId,
			&result.MemberName,
			&result.MemberIntro,
			&result.MemberImg,
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

func (this *SqlScTeamIntroduce) QueryByCondition(condition string) error {
	this.Clear()
	sql := "select * from sc_team_introduce "
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
		var result ScTeamIntroduce
		err = rows.Scan(&result.MemberId,
			&result.MemberName,
			&result.MemberIntro,
			&result.MemberImg,
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

func (this *SqlScTeamIntroduce) Delete(condition string) error {
	if len(condition) == 0 {
		return errors.New("conditon is null")
	}
	sql := "delete from sc_team_introduce where "
	sql += condition
	fmt.Println(sql)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScTeamIntroduce) UpdataState(id int) error {
	sql := fmt.Sprintf("update sc_team_introduce set status=1 where member_id=%d", id)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScTeamIntroduce) Update() error{
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("update sc_team_introduce set member_name = '%s',member_intro = '%s'," +
		"image = '%s',status = %d,create_time = '%s',update_time = '%s' "+
		"where member_id = %d",
		this.Data.MemberName,
		this.Data.MemberIntro,
		this.Data.MemberImg,
		this.Data.Status,
		t,
		t,
			this.Data.MemberId)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScTeamIntroduce) Insert() error {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("insert into sc_team_introduce(member_name, member_intro,image,status,"+
		"create_time,update_time) "+
		"values('%s','%s','%s',%d,'%s','%s')",
		this.Data.MemberName,
		this.Data.MemberIntro,
		this.Data.MemberImg,
		this.Data.Status,
		t,
		t)
	_, err := this.DBHandle.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *SqlScTeamIntroduce) Clear() {
	if this.Row > 0 {
		this.Result = append(this.Result[this.Row:])
		this.Row = 0
	}
}

func (this *SqlScTeamIntroduce) SerializeResult() (string, error) {
	jsoneamIntroduce := &JsonTeamIntroduce{
		ErrCode: 2000,
		Message: "OK",
		Data:    &this.Result,
	}
	retJson, err := json.Marshal(jsoneamIntroduce)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(retJson), err
}

func (this *SqlScTeamIntroduce) UnSerializeJson(request string) error {
	err := json.Unmarshal([]byte(request), &this.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
