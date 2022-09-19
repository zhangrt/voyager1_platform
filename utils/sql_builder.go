package utils

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// 自定义SQL
type SQLBuilder struct {
	hasDB bool
}

type DBAdapter SQLBuilder

// 占位符防止sql注入
var (
	DB    *gorm.DB
	DBS   map[string]*gorm.DB
	SPACE = " "
	LIKE  = " like ?"
	EQ    = " = ?"
	IN    = " in (?)"
	PLACE = " ? "
)

// adapter
func (builder *SQLBuilder) Adapter(tx *gorm.DB) *SQLBuilder {
	DB = tx
	return builder
}

func (builder *SQLBuilder) AdapterSafety(k string, v *gorm.DB) *SQLBuilder {
	DBS = make(map[string]*gorm.DB)
	DBS[k] = v
	builder.hasDB = true
	return builder
}

func (builder *SQLBuilder) Model(v interface{}) *SQLBuilder {
	DB = DB.Model(v)
	return builder
}

func (builder *SQLBuilder) ModelSafety(k string, v interface{}) *SQLBuilder {
	DBS[k] = DBS[k].Model(v)
	return builder
}

// param, commad, value
func (builder *SQLBuilder) WhereSafety(k string, p string, c string, v interface{}) *SQLBuilder {
	var db = DBS[k]
	if !Pre(db, v) {
		return builder
	}

	db = Where(db, p, c, v)
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

// param, commad, value
func (builder *SQLBuilder) Where(p string, c string, v interface{}) *SQLBuilder {
	if !Pre(DB, v) {
		return builder
	}

	DB = Where(DB, p, c, v)

	return builder
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

func (builder *SQLBuilder) Order(r string) *SQLBuilder {
	DB = Order(DB, r)
	return builder
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
		if orders != nil {
			for k, v := range orders {
				order := k + " " + v
				db = db.Order(order)
			}
		}
	}

	return db
}

func (builder *SQLBuilder) OrderSafety(k string, r string) *SQLBuilder {
	DBS[k] = Order(DBS[k], r)
	return builder
}

func (builder *SQLBuilder) Page(page int, pageSize int) *SQLBuilder {
	DB = Page(DB, page, pageSize)
	return builder
}

func (builder *SQLBuilder) PageSafety(k string, page int, pageSize int) *SQLBuilder {
	DBS[k] = Page(DBS[k], page, pageSize)
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

func (builder *SQLBuilder) Go() (tx *gorm.DB) {
	defer Close()
	return DB
}

func (builder *SQLBuilder) GoSafety(k string) (tx *gorm.DB) {
	defer CloseSafety(k)
	return DBS[k]
}

func Close() {
	if DB != nil {
		DB = nil
	}
}

func CloseSafety(k string) {
	if DBS[k] != nil {
		DBS[k] = nil
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

func (builder *SQLBuilder) Get() (tx *gorm.DB) {
	return DB
}
