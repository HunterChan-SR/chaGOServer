package bean

// create table recontext
// (
//
//	id           int primary key auto_increment,
//	needhelpid   int references needhelp (id),
//	recontext    varchar(2047),
//	createtime   datetime default current_timestamp
//
// );
type Recontext struct {
	Id         int    `json:"id"`
	Needhelpid int    `json:"needhelpid"`
	Recontext  string `json:"recontext"`
	Createtime string `json:"createtime"`
}

func (Recontext) TableName() string {
	return "recontext"
}
