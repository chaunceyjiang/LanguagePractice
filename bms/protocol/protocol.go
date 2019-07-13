package protocol

import (
	"bms/utils"
	"log"
	"path"
	"strconv"
	"strings"
)

func (p ProtocolConfig) Int2ip(i int64) string {
	return strconv.FormatInt(i>>24&0xff, 10) + "." + strconv.FormatInt(i>>16&0xff, 10) + "." +
		strconv.FormatInt(i>>4&0xff, 10) + "." + strconv.FormatInt(i&0xff, 10)
}

func (p ProtocolConfig) NormalizesHandler(record map[string]interface{}) {
	for _, n := range p.Normalizes.Normalize {
		n.NormalizeHandler(record, p.ProtocolFileName)
	}
}

func (n Normalize) NormalizeHandler(record map[string]interface{}, opts ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(n, record, err)
		}
	}()
	var name = n.Name
	if n.If != nil {
		for _, i := range n.If {
			if r, ok := i.FieldHandler(record, opts...).(bool); ok && r { //多个if标签处理
				for _, f := range i.Fields {
					if f.FieldHandler(record, opts...) != nil {
						record[name] = f.FieldHandler(record, opts...)
						return
					}
				}
			}
		}
	}
	if n.Else.Fields != nil {
		for _, f := range n.Else.Fields {
			if f.FieldHandler(record, opts...) != nil {
				record[name] = f.FieldHandler(record, opts...)
				return
			}
		}

	} else if n.Fields != nil {
		for _, f := range n.Fields {
			if f.FieldHandler(record, opts...) != nil {
				record[name] = f.FieldHandler(record, opts...)
				return //处理
			}
		}
	} else {
		record[name] = nil //没有任何处理 ,则该字段为空值
	}

}
func (p ProtocolConfig) PrepareHandler(record map[string]interface{}) {
	var f Field

	// prepare 字段采用xml.Text
	for _, f = range p.Prepare.Fields {
		record[f.Text] = f.FieldHandler(record, p.ProtocolFileName)
	}
}
func (p ProtocolConfig)Process(record map[string]interface{})  {
	p.PrepareHandler(record)
	p.NormalizesHandler(record)
}
func (f Attr) FieldHandler(record map[string]interface{}, opts ...interface{}) (r interface{}) {
	defer func(r *interface{}) {
		if err := recover(); err != nil {
			r = nil
		}
	}(&r)
	switch f.Op {
	case "as", "notNull", "isNone", "str_trim": //一元操作符
		if f.Value != "" {
			var v interface{}
			switch f.Type {
			case "":
				return MapFunc[f.Op](f.Value)
			case "int":
				v, _ = strconv.Atoi(f.Value)
				return MapFunc[f.Op](v)
			case "float":
				v, _ = strconv.ParseFloat(f.Value, 32)

			}
			return MapFunc[f.Op](v)
		}
		t := MapFunc[f.Op](record[f.Item1])
		switch oj := t.(type) {
		case bool:
			if oj && f.Result != "" {
				return record[f.Result]
			}
		default:
			return t
		}

	case "str_sub": //三元操作符
		return MapFunc[f.Op](record[f.Item1], f.Start, f.End)
	case "lookup":
		return MapFunc[f.Op](path.Join(opts[0].(string), f.File), record[f.Item1])
	default: //二元操作符 相应处理
		var y interface{}
		if f.Value != "" {
			y = f.Value
		} else {
			y = record[f.Item2]
		}
		return MapFunc[f.Op](record[f.Item1], y) //这里基本上是对固定值进行处理的
	}
	return MapFunc[f.Op]()
}

var MapFunc = map[string]func(s ...interface{}) interface{}{
	"str_add": func(s ...interface{}) interface{} {
		return s[0].(string) + s[1].(string)
	},
	"str_sub": func(s ...interface{}) interface{} {
		start, _ := strconv.Atoi(s[1].(string))
		end, _ := strconv.Atoi(s[2].(string))
		return s[0].(string)[start:end]
	},
	"str_trim": func(s ...interface{}) interface{} {
		return strings.Trim(s[0].(string), " ")
	},
	"str_regex": nil, // TODO
	"eq": func(b ...interface{}) interface{} {
		return b[0] == b[1]
	},
	"ne": func(b ...interface{}) interface{} {
		return b[0] != b[1]
	},
	"and": func(b ...interface{}) interface{} {
		return b[0].(bool) && b[1].(bool)
	},
	"or": func(b ...interface{}) interface{} {
		return b[0].(bool) || b[1].(bool)
	},
	"add": func(i ...interface{}) interface{} {
		return i[0].(int) + i[1].(int)
	},
	"sub": func(i ...interface{}) interface{} {
		return i[0].(int) - i[1].(int)
	},
	"multi": func(i ...interface{}) interface{} {
		return i[0].(int) * i[1].(int)
	},
	"divide": func(i ...interface{}) interface{} {
		return i[0].(int) / i[1].(int)
	},
	"notNull": func(i ...interface{}) interface{} {
		if i[0] != nil {
			return true
		}
		return false
	},
	"isNone": func(i ...interface{}) interface{} {
		if i[0] == nil {
			return true
		}
		return false
	},
	"in": func(s ...interface{}) interface{} {
		return strings.Contains(s[1].(string), s[0].(string))
	},
	"notIn": func(s ...interface{}) interface{} {
		return !strings.Contains(s[1].(string), s[0].(string))
	},
	"as": func(s ...interface{}) interface{} {
		return s[0]
	},
	"lookup": lookup,
}

func lookup(s ...interface{}) interface{} {
	if s[1] == nil {
		return nil
	}
	var data map[string]string
	data = utils.LoadCsv(s[0].(string))
	return data[s[1].(string)]
}
