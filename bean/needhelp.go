package bean

// create table needhelp
// (
//
//	id           int primary key auto_increment,
//	days         int,
//	userid       int references user (id),
//	problemtitle varchar(255),
//	subcode      varchar(255),
//	context      varchar(255),
//	recontext    varchar(255),
//	createtime   datetime default current_timestamp
//
// );
type NeedHelp struct {
	Id           int    `json:"id"`
	Days         int    `json:"days"`
	Userid       int    `json:"userid"`
	Problemtitle string `json:"problemtitle"`
	Subcode      string `json:"subcode"`
	Context      string `json:"context"`
	Createtime   string `json:"createtime"`
}

func (NeedHelp) TableName() string {
	return "needhelp"
}
