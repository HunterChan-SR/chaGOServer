package bean

// create table savedfile
// (
//
//	id         int primary key auto_increment,
//	userid     int references user (id),
//	filename   varchar(255),
//	createtime datetime default current_timestamp
//
// );
type SavedFile struct {
	Id         int    `json:"id"`
	Userid     int    `json:"userid"`
	Filename   string `json:"filename"`
	Createtime string `json:"createtime"`
}

func (SavedFile) TableName() string {
	return "savedfile"
}
