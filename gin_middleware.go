package audit_log

import (
	"log"

	"github.com/gin-gonic/gin"
)

func NewAuditLogGinMiddleware(saveFn func(op *OperateLogs) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		go func() {
			op, exists := c.Get(KEY_LOG)
			if !exists {
				return
			}
			operateLog, ok := op.(*OperateLogs)
			if !ok {
				log.Print("op.(*OperateLogs) err")
				return
			}
			operateLog.IP = convIP(c.Request.Header)
			operateLog.AccountID = c.GetUint("id")
			operateLog.GroupID = c.GetUint("groupId")

			needFields := []string{}
			if val, exists := c.Get(KEY_NEED_FIELDS); exists {
				needFields, ok = val.([]string)
				if !ok {
					log.Print("needFields get err")
					return
				}
			}
			expectFields := []string{}
			if val, exists := c.Get(KEY_EXPECT_FIELDS); exists {
				expectFields, ok = val.([]string)
				if !ok {
					log.Print("existsFields get err")
					return
				}
			}

			beforeFields, existsBefore := c.Get(KEY_BEFORE)
			afterFields, existsAfter := c.Get(KEY_AFTER)

			if existsBefore || existsAfter {
				fieldsSlice, beforeSlice, afterSlice, err := GetFieldsLogSlice(beforeFields, afterFields, needFields, expectFields)
				if err != nil {
					log.Printf("GetFieldsLogSlice err %v", err)
					return
				}
				operateLog.Fields = fieldsSlice
				operateLog.FieldsBefore = beforeSlice
				operateLog.FieldsAfter = afterSlice
			}

			err := saveFn(operateLog)
			if err != nil {
				log.Printf("create operateLog %v", err)
				return
			}
		}()
	}
}
