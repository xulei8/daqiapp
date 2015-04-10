package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表

type Appmain struct {
	Id        int
	Title     string    `orm:"size(200);null"`
	Creatorid int       `orm:"null"`
	Ownerid   int       `orm:"null"`
	Deleted   int       `orm:"null"`
	Modname   string    `orm:"size(20);null"`
	Addtime   time.Time `orm:"null"`
	Edittime  time.Time `orm:"null"`
}
type DqContact struct {
	Id      int64  ` edithide:"true" `
	Sex     int    `orm:" null"  juiclass:"easyui-combobox"   allowList:"ok" label:"性别"  op:"textField: 'label', valueField: 'value',data: [{label: '男',value: '2'},{label: '女',value: '1'}]"  `
	Uname   string `orm:"size(16);null" juiclass:"easyui-numberbox"  allowList:"true"  `
	Piname  string `orm:"size(16);null"  juiclass:"easyui-datetimespinner"  `
	Tel     string `orm:"size(22);null;index"     allowList:"true"    juiclass:"easyui-datebox"    `
	Tel2    string `orm:"size(22);null"  value:"33"   allowList:"true"   `
	Comname string `orm:"size(33);null"   label:"公司"   `
	Mname   string `orm:"size(16);null"   label:"名称"   `
	Address string `orm:"size(32);null"    label:"地址"  `
	Note    string `orm:"size(152);null"   label:"备注"  `
	Email   string `orm:"size(32);null"   label:"邮件" `
	Qq      string `orm:"size(32);null"   label:"QQ" `

	Called        int       `orm:" null"  `
	Hits          int       `orm:" null"  label:"次数"  value:"30" juiclass:"easyui-numberspinner"  op:"min:10,max:100,editable:false"    `
	CallHits      int       `orm:" null"  label:"拨打" value:"40"   juiclass:"easyui-numberspinner"  op:"min:30,max:600,editable:true"   `
	Statu         int       `orm:" null"  label:"状态"  value:"33"      juiclass:"easyui-slider"    op:"showTip:true,rule:[0,'|',25,'|',50,'|',75,'|',100]"  `
	Lockit        int       `orm:" null"  `
	Addtime       time.Time `orm:"auto_now;type(datetime); null"`
	FirstCallTime time.Time `orm:"auto_now;type(datetime); null"`
	LastCalltime  time.Time `orm:"auto_now;type(datetime); null"`
	EditTime      time.Time `orm:"auto_now;type(datetime); null"`
}

type DqTest struct {
	Id      int64
	Sex     int    `orm:" null"  label:"性别"  `
	Uname   string `orm:"size(16);null" juiclass:"easyui-numberbox"  `
	Piname  string `orm:"size(16);null"  juiclass:"easyui-datetimespinner"  `
	Tel     string `orm:"size(22);null;index"    juiclass:"easyui-datebox"    `
	Tel2    string `orm:"size(22);null"  `
	Comname string `orm:"size(33);null"  `
	Mname   string `orm:"size(16);null"  `
	Address string `orm:"size(32);null"  `
	Note    string `orm:"size(152);null"  `
	Email   string `orm:"size(32);null"  `
	Qq      string `orm:"size(32);null"  `

	Called        int       `orm:" null"  `
	Hits          int       `orm:" null"  `
	CallHits      int       `orm:" null"  `
	Statu         int       `orm:" null"  `
	Lockit        int       `orm:" null"  `
	Addtime       time.Time `orm:"auto_now;type(datetime); null"`
	FirstCallTime time.Time `orm:"auto_now;type(datetime); null"`
	LastCalltime  time.Time `orm:"auto_now;type(datetime); null"`
	EditTime      time.Time `orm:"auto_now;type(datetime); null"`
}

func init() {

	orm.RegisterModel(new(Appmain))
	orm.RegisterModel(new(DqContact))
	orm.RegisterModel(new(DqTest))
}
