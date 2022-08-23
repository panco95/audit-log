package fields_log

import (
	"testing"
)

func TestUpdateLog(t *testing.T) {
	beforeFields := map[string]interface{}{
		"username": "panco",
		"remark":   "牛逼",
		"mobile":   "10086",
		"password": "123456",
		"status":   false,
	}
	afterFields := map[string]interface{}{
		"username": "panco",
		"remark":   "super牛逼",
		"mobile":   "10010",
		"password": "abcdefg",
		"status":   true,
	}
	result, err := UpdateFieldsLog(beforeFields, afterFields, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
