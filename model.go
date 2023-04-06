package audit_log

type OperateLogs struct {
	AccountID    uint     `json:"accountId"`
	GroupID      uint     `json:"groupId"`
	Module       string   `json:"module"`
	IP           string   `json:"ip"`
	Content      string   `json:"content"`
	Detail       string   `json:"detail"`
	Fields       []string `json:"fields"`
	FieldsBefore []string `json:"beforeFields"`
	FieldsAfter  []string `json:"afterFields"`
}
