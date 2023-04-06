## audit-log

Generate audio log.

Jow support gin now, but you can customize.

# How to use

1. Define the custom mapping
```
    import {
        audit_log "github.com/panco95/audit-log"
        ...
    }

	audit_log.SetLang("cn")
	fieldMapping := map[string]audit_log.MappingField{
		"name": {
			"cn": "姓名",
		},
		"age": {
			"cn": "年龄",
		},
	}
	audit_log.SetFieldMapping(fieldMapping)
```


2. Define gin middleware
```
    import {
        audit_log "github.com/panco95/audit-log"
        ...
    }

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.HTTPGzipEncoding)

	api := router.Group("/api/v1")
	api.Use(audit_log.NewAuditLogGinMiddleware(func(op *audit_log.OperateLogs) error {
		// save func, like database write to table, here we just print
		log.Print(map[string]interface{}{
			"account_id":   op.AccountID,    // write log from account id field if need
			"group_id":     op.GroupID,      // write lof from group id field if need
			"module":       op.Module,       // write log from module field if need
			"ip":           op.IP,           // write log from ip field if need
			"content":      op.Content,      // write log content field
			"detail":       op.Detail,       // write log detail field
			"fields":       op.Fields,       // write log fields list
			"beforeFields": op.FieldsBefore, // write log change before fields value list
			"afterFields":  op.FieldsAfter,  // write log change after fields value list
		})
		return nil
	}))
```

3. Write the fields and values of the operation in the code
```
    import {
        audit_log "github.com/panco95/audit-log"
        ...
    }

	type User struct {
		Name string
		Age  int
	}

	func Update(c *gin.Context) {
		c.Set(audit_log.KEY_LOG, &audit_log.OperateLogs{Module: "用户中心", Content: "更新资料"})
		// c.Set(audit_log.KEY_NEED_FIELDS, []string{"name", "age"})
		c.Set(audit_log.KEY_BEFORE, User{Name: "panco", Age: 18})
		c.Set(audit_log.KEY_AFTER, User{Name: "man", Age: 28})
	}
```

4. Result
```
	map[account_id:0 afterFields:[man 28] beforeFields:[panco 18] content:更新资料 detail: fields:[姓名 年龄] group_id:0 ip:127.0.0.1 module:用户中心]
```

Arrange the fields before and after the modification in the order of the fields...