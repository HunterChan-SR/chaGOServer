package bean

// create table recontext
// (
//
//	id           int primary key auto_increment,
//	needhelpid   int references needhelp (id),
//	userid       int references user (id),
//	recontext    varchar(2047),
//	createtime   datetime default current_timestamp
//
// );
type Recontext struct {
	Id         int    `json:"id"`
	Needhelpid int    `json:"needhelpid"`
	Userid     int    `json:"userid"`
	Recontext  string `json:"recontext"`
	Createtime string `json:"createtime"`
}

func (Recontext) TableName() string {
	return "recontext"
}

// create view recontextview as
//
//	select
//	    recontext.id,
//	    recontext.needhelpid,
//	    user.nickname,
//	    recontext.recontext,
//	    recontext.createtime
//	from recontext,
//	     user
//	where recontext.userid = user.id;
type RecontextView struct {
	Id         int    `json:"id"`
	Needhelpid int    `json:"needhelpid"`
	Nickname   string `json:"nickname"`
	Recontext  string `json:"recontext"`
	Createtime string `json:"createtime"`
}

func (RecontextView) TableName() string {
	return "recontextview"
}
