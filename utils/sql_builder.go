package utils

import (
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

// 自定义SQL
type SQLBuilder struct {
	lock    sync.RWMutex
	sqlLock sync.RWMutex
	dbs     map[string]*gorm.DB // db
	hasDB   bool
	sqls    map[string]string // original sql
}

type DBAdapter SQLBuilder

// 占位符防止sql注入
var (
	SPACE = " "
	LIKE  = " LIKE ?"
	EQ    = " = ?"
	IN    = " IN (?)"
	PLACE = " ? "
	AND   = " AND "
)

// SQL 代理
var SQLAdapterObj *SQLBuilder

// 初始化方法
func init() {
	SQLAdapterObj = &SQLBuilder{
		dbs:   make(map[string]*gorm.DB),
		hasDB: false,
		sqls:  make(map[string]string),
	}
}

func New() *SQLBuilder {
	return &SQLBuilder{}
}

// adapter
func (builder *SQLBuilder) Adapter(k string, v *gorm.DB) *SQLBuilder {
	builder.lock.Lock()
	defer builder.lock.Unlock()
	builder.dbs[k] = v
	builder.hasDB = true
	return builder
}

func (builder *SQLBuilder) OriginSql(k string, sql string) *SQLBuilder {
	builder.sqlLock.Lock()
	defer builder.sqlLock.Unlock()
	builder.sqls[k] = sql
	return builder
}

// k key唯一key, n name属性名称, c command命令, v value 值
func (builder *SQLBuilder) And(k string, n string, c string, v interface{}) *SQLBuilder {
	builder.sqlLock.Lock()
	defer builder.sqlLock.Unlock()
	builder.sqls[k] = And(builder.dbs[k], builder.sqls[k], n, c, v)
	return builder
}

func And(db *gorm.DB, sql string, n string, c string, v interface{}) string {
	switch c {
	case "like", "-like-", "-like", "like-":
		sql = sql + AND + n + LIKE + PLACE
		db.Row()
	case "=":

	case "in":
		if str, ok := v.(string); ok {
			arr := strings.Split(str, ",")
			if len(arr) == 1 {
			} else {
			}
		}
		// [""]
		if val, ok := v.([]string); ok {
			if len(val) == 1 && val[0] != "" {
			} else if len(val) > 1 {
			}
		}
	default:
	}

	return sql
}

func (builder *SQLBuilder) Model(k string, v interface{}) *SQLBuilder {
	builder.lock.Lock()
	defer builder.lock.Unlock()
	builder.dbs[k] = builder.dbs[k].Model(v)
	return builder
}

// param, commad, value
func (builder *SQLBuilder) Where(k string, p string, c string, v interface{}) *SQLBuilder {
	builder.lock.Lock()
	defer builder.lock.Unlock()
	var db = builder.dbs[k]
	if !Pre(db, v) {
		return builder
	}

	builder.dbs[k] = Where(db, p, c, v)

	return builder
}

// 前置判断空值
func Pre(db *gorm.DB, v interface{}) bool {
	if db.Statement == nil {
		return false
	}

	if t, ok := v.(time.Time); ok {
		nilTime := time.Time{}
		if t == nilTime {
			return false
		}
	} else if v == nil {
		return false
	} else if s, ok := v.(string); ok && s == "" {
		return false
	}
	return true
}

// 条件查询
func Where(db *gorm.DB, p string, c string, v interface{}) *gorm.DB {
	switch c {
	case "like", "-like-":
		db.Where(p+LIKE, Like(v.(string), 0))
	case "-like":
		db.Where(p+LIKE, Like(v.(string), 1))
	case "like-":
		db.Where(p+LIKE, Like(v.(string), 2))
	case "=":
		db.Where(p+EQ, v)
	case "in":
		if str, ok := v.(string); ok {
			arr := strings.Split(str, ",")
			if len(arr) == 1 {
				db.Where(p+EQ, v)
			} else {
				db.Where(p+IN, arr)
			}
		}
		// [""]
		if val, ok := v.([]string); ok {
			if len(val) == 1 && val[0] != "" {
				db.Where(p+EQ, v)
			} else if len(val) > 1 {
				db.Where(p+IN, v)
			}
		}
	default:
		db.Where(p+SPACE+c+PLACE, v)
	}
	return db
}

func Order(db *gorm.DB, r string) *gorm.DB {
	if db.Statement != nil {
		var orders map[string]string
		if r != "" {
			rule := strings.Split(r, ",")
			orders = make(map[string]string, len(rule))
			for _, k := range rule {
				p := strings.Split(k, "-")
				orders[p[0]] = p[1]
			}
		}
		for k, v := range orders {
			order := k + " " + v
			db = db.Order(order)
		}

	}

	return db
}

func (builder *SQLBuilder) Order(k string, r string) *SQLBuilder {
	builder.lock.Lock()
	defer builder.lock.Unlock()
	builder.dbs[k] = Order(builder.dbs[k], r)
	return builder
}

func (builder *SQLBuilder) Page(k string, page int, pageSize int) *SQLBuilder {
	builder.lock.Lock()
	defer builder.lock.Unlock()
	builder.dbs[k] = Page(builder.dbs[k], page, pageSize)
	return builder
}

func Page(db *gorm.DB, page int, pageSize int) *gorm.DB {
	if db.Statement != nil {
		limit := pageSize
		offset := pageSize * (page - 1)
		db.Limit(limit).Offset(offset)
	}
	return db
}

func (builder *SQLBuilder) Go(k string) (tx *gorm.DB) {
	builder.lock.RLock()
	defer builder.lock.RUnlock()
	defer Close(k, builder)
	return builder.dbs[k]
}

func Close(k string, builder *SQLBuilder) {
	if builder.dbs[k] != nil {
		builder.dbs[k] = nil
	}
}

// 模糊查询
func Like(v string, t int) string {
	var r string
	for _, c := range v {
		s := string(c)
		// 通配符处理
		if s == "_" || s == "^" || s == "%" {
			r += "\\" + s
		} else {
			r += s
		}
	}
	if t == 0 {
		return "%" + v + "%"
	} else if t == 1 {
		return "%" + v
	} else {
		return v + "%"
	}
}

func (builder *SQLBuilder) Get(key string) (tx *gorm.DB) {
	builder.lock.RLock()
	defer builder.lock.RUnlock()
	return builder.dbs[key]
}
