package bean

//create table submit
//(
//id        int primary key auto_increment,
//userid    int references user (id),
//problemid int references problem (id),
//state     varchar(255)
//);

type Submit struct {
	Id        int    `json:"id"`
	Userid    int    `json:"userid"`
	Problemid int    `json:"problemid"`
	State     string `json:"state"`
	Dfbyid    string `json:"dfbyid"`
}

func (Submit) TableName() string {
	return "submit"
}
