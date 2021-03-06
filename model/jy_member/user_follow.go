package JYMemberDB

import (
	"errors"
	"fmt"
	"strings"

	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
	"changit.cn/contra/server/game_error"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

//This file is generate by scripts,don't edit it

//user_follow
//用户关注表

// +gen *
type UserFollow struct {
	Id         int   `db:"id" json:"id"`                   //
	Uid        int   `db:"uid" json:"uid"`                 // 关注者ID
	PassiveUid int   `db:"passive_uid" json:"passive_uid"` // 被关注者UID（关注话题时为0）
	Tid        int   `db:"tid" json:"tid"`                 // 话题分类ID（关注个人时为0）
	Type       int8  `db:"type" json:"type"`               // 订阅类型：1关注个人，2关注话题集合
	Status     int8  `db:"status" json:"status"`           // 状态（1关注中，9取消关注）
	CTime      int64 `db:"c_time" json:"c_time"`           // 添加时间
	UTime      int64 `db:"u_time" json:"u_time"`           // 更新时间
	RoomId     int8  `db:"room_id" json:"room_id"`         // 非0表示开通直播的
}

type userFollowOp struct{}

var UserFollowOp = &userFollowOp{}
var DefaultUserFollow = &UserFollow{}

// 按主键查询. 注:未找到记录的话将触发sql.ErrNoRows错误，返回nil, error
func (op *userFollowOp) Get(id int) (*UserFollow, error) {
	obj := &UserFollow{}
	sql := "select `id`,`uid`,`passive_uid`,`tid`,`type`,`status`,`c_time`,`u_time`,`room_id` from user_follow where id=? "
	err := db.JYMemberDB.Get(obj, sql,
		id,
	)

	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (op *userFollowOp) SelectAll() ([]*UserFollow, error) {
	objList := []*UserFollow{}
	sql := "select `id`,`uid`,`passive_uid`,`tid`,`type`,`status`,`c_time`,`u_time`,`room_id` from user_follow"
	err := db.JYMemberDB.Select(&objList, sql)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return objList, nil
}

func (op *userFollowOp) QueryByMap(m map[string]interface{}) ([]*UserFollow, error) {
	result := []*UserFollow{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`tid`,`type`,`status`,`c_time`,`u_time`,`room_id` from user_follow where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userFollowOp) QueryByMapComparison(m map[string]interface{}) ([]*UserFollow, error) {
	result := []*UserFollow{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`tid`,`type`,`status`,`c_time`,`u_time`,`room_id` from user_follow where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}
	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userFollowOp) QueryByClause(m map[string]interface{}, limit, offset int, orderby, clause []string) ([]*UserFollow, error) {
	result := []*UserFollow{}
	var params []interface{}

	sql := "select `id`,`uid`,`passive_uid`,`tid`,`type`,`status`,`c_time`,`u_time`,`room_id` from user_follow where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s? ", k)
		params = append(params, v)
	}

	if len(orderby) > 0 {
		for k, v := range orderby {
			if len(v) < 2 {
				continue
			}
			opr := v[:1]
			switch opr {
			case "-":
				orderby[k] = fmt.Sprintf("%s desc", v[1:len(v)])
			case "+":
				orderby[k] = fmt.Sprintf("%s", v[1:len(v)])
			default:
				orderby[k] = fmt.Sprintf("%s", v)
			}
		}
		sql += fmt.Sprintf(" order by %s", strings.Join(orderby, ", "))
	}

	if len(clause) > 0 {
		sql += fmt.Sprintf(" %s", strings.Join(clause, " "))
	}

	if limit > 0 {
		sql += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}

	zaplogger.Info("[SQL]:"+sql, zap.Reflect("| params:", params))

	err := db.JYMemberDB.Select(&result, sql, params...)
	if err != nil {
		zaplogger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (op *userFollowOp) GetByMap(m map[string]interface{}) (*UserFollow, error) {
	lst, err := op.QueryByMap(m)
	if err != nil {
		return nil, err
	}
	if len(lst) == 1 {
		return lst[0], nil
	} else if len(lst) == 0 {
		return nil, nil
	}

	return nil, errors.New("Get multi rows.")
}

// 插入数据，自增长字段将被忽略
func (op *userFollowOp) Insert(m *UserFollow) (int64, error) {
	return op.InsertTx(db.JYMemberDB, m)
}

// 插入数据，自增长字段将被忽略
func (op *userFollowOp) InsertTx(ext sqlx.Ext, m *UserFollow) (int64, error) {
	sql := "insert into user_follow(uid,passive_uid,tid,type,status,c_time,u_time,room_id) values(?,?,?,?,?,?,?,?)"
	result, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Tid,
		m.Type,
		m.Status,
		m.CTime,
		m.UTime,
		m.RoomId,
	)
	if err != nil {
		game_error.RaiseError(err)
		return -1, err
	}
	affected, _ := result.RowsAffected()
	return affected, nil
}

/*
func (i *UserFollow) Update() {
    _,err := db.JYMemberDB.Update(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFollowOp) Update(m *UserFollow) error {
	return op.UpdateTx(db.JYMemberDB, m)
}

// 用主键(属性)做条件，更新除主键外的所有字段
func (op *userFollowOp) UpdateTx(ext sqlx.Ext, m *UserFollow) error {
	sql := `update user_follow set uid=?,passive_uid=?,tid=?,type=?,status=?,c_time=?,u_time=?,room_id=? where id=?`
	_, err := ext.Exec(sql,
		m.Uid,
		m.PassiveUid,
		m.Tid,
		m.Type,
		m.Status,
		m.CTime,
		m.UTime,
		m.RoomId,
		m.Id,
	)

	if err != nil {
		game_error.RaiseError(err)
		return err
	}

	return nil
}

// 用主键做条件，更新map里包含的字段名
func (op *userFollowOp) UpdateWithMap(id int, m map[string]interface{}) error {
	return op.UpdateWithMapTx(db.JYMemberDB, id, m)
}

// 用主键做条件，更新map里包含的字段名
func (op *userFollowOp) UpdateWithMapTx(ext sqlx.Ext, id int, m map[string]interface{}) error {

	sql := `update user_follow set %s where 1=1 and id=? ;`

	var params []interface{}
	var set_sql string
	for k, v := range m {
		set_sql += fmt.Sprintf(" %s=? ", k)
		params = append(params, v)
	}
	params = append(params, id)
	_, err := ext.Exec(fmt.Sprintf(sql, set_sql), params...)
	return err
}

/*
func (i *UserFollow) Delete(){
    _,err := db.JYMemberDB.Delete(i)
    if err != nil{
        game_error.RaiseError(err)
    }
}
*/
// 根据主键删除相关记录
func (op *userFollowOp) Delete(id int) error {
	return op.DeleteTx(db.JYMemberDB, id)
}

// 根据主键删除相关记录,Tx
func (op *userFollowOp) DeleteTx(ext sqlx.Ext, id int) error {
	sql := `delete from user_follow where 1=1
        and id=?
        `
	_, err := ext.Exec(sql,
		id,
	)
	return err
}

// 返回符合查询条件的记录数
func (op *userFollowOp) CountByMap(m map[string]interface{}) int64 {

	var params []interface{}
	sql := `select count(*) from user_follow where 1=1 `
	for k, v := range m {
		sql += fmt.Sprintf(" and  %s=? ", k)
		params = append(params, v)
	}
	count := int64(-1)
	err := db.JYMemberDB.Get(&count, sql, params...)
	if err != nil {
		game_error.RaiseError(err)
	}
	return count
}

func (op *userFollowOp) DeleteByMap(m map[string]interface{}) (int64, error) {
	return op.DeleteByMapTx(db.JYMemberDB, m)
}

func (op *userFollowOp) DeleteByMapTx(ext sqlx.Ext, m map[string]interface{}) (int64, error) {
	var params []interface{}
	sql := "delete from user_follow where 1=1 "
	for k, v := range m {
		sql += fmt.Sprintf(" and %s=? ", k)
		params = append(params, v)
	}
	result, err := ext.Exec(sql, params...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
