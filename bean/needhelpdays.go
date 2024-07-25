package bean

type NeedHelpDays struct {
	Id int `json:"id"`
}

func (NeedHelpDays) TableName() string {
	return "needhelpdays"
}
