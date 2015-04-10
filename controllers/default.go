package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/daqiapp/models"
	"reflect"
	"strings"
	"time"
	//	"time"
)

type MainController struct {
	beego.Controller
}

type XmlGet struct {
	beego.Controller
}

type DataDelete struct {
	beego.Controller
}

type DataGet struct {
	beego.Controller
}

type DataSave struct {
	beego.Controller
}

type Xml struct {
	beego.Controller
}

type XmlSave struct {
	beego.Controller
}

type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,年龄："`
	Sex   string
	Intro string `form:",textarea"`
}

func (c *DataDelete) Post() {
	modname := c.GetString("FormModName")
	Id, _ := c.GetInt64("Id", 0)

	if Id < 1 {
		c.Ctx.WriteString(`{"success":false }`)
		return
	}

	o := orm.NewOrm()

	switch modname {
	case "DqContact":
		mod := m.DqContact{}
		mod.Id = Id
		o.Delete(&mod)
	}

	c.Ctx.WriteString(`{"success":true }`)
}

func (c *XmlSave) Post() {

	modname := c.GetString("FormModName")
	Deleteid, _ := c.GetInt("Deleteid", 0)
	o := orm.NewOrm()

	if Deleteid > 0 {
		mmd := m.Appmain{Id: Deleteid}
		mmd.Deleted = 1
		o.Update(&mmd)
		c.Ctx.WriteString(`{"success":true }`)
		return
	}

	iid, _ := c.GetInt("id", 0)

	t := c.GetString("title")
	mname := beego.ModXML[modname].Tablename

	mm := m.Appmain{Id: iid}
	mm.Title = t
	mm.Modname = modname
	mm.Deleted = 0

	objid := 0

	if iid > 0 {
		mm.Edittime = time.Now()
		o.Update(&mm)
		objid = iid
	} else {
		mm.Addtime = time.Now()
		nid, _ := o.Insert(&mm)
		objid = int(nid)
	}

	sql := ""
	var raw []string
	if objid > 0 {

		cc := len(beego.ModXML[modname].Fields.Filed)
		i := 0
		for i = 0; i < cc; i++ {
			if beego.ModXML[modname].Fields.Filed[i].Name == "title" {
				continue
			}
			f := c.GetString(beego.ModXML[modname].Fields.Filed[i].Name)
			raw = append(raw, " "+beego.ModXML[modname].Fields.Filed[i].Name+" ='"+f+"' ")
		}
		if iid > 0 {
			sql += " update " + mname + "   set " + strings.Join(raw, ",") + "  where aid =  '" + fmt.Sprintf("%d", objid) + "' limit 1 "
		} else {
			sql += " insert into " + mname + "   set " + strings.Join(raw, ",") + " ,aid =  '" + fmt.Sprintf("%d", objid) + "'"
		}
		fmt.Print(sql)
		o.Raw(sql).Exec()

	}

	c.Ctx.WriteString(`{"success":true }`)
}

func (c *XmlGet) Post() {
	rows, _ := c.GetInt("rows", 10)
	if rows < 1 {
		rows = 1
	}
	if rows > 300 {
		rows = 300
	}

	page, _ := c.GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * rows

	modname := c.GetString("FormModName")
	cc := len(beego.ModXML[modname].Fields.Filed)
	sqls := ""
	i := 0
	var flist []string
	for i = 0; i < cc; i++ {

		flist = append(flist, beego.ModXML[modname].Fields.Filed[i].Name)

	}
	o := orm.NewOrm()
	sqln := "select count(*) as c from appmain , " + beego.ModXML[modname].Tablename + " where id = aid  and deleted = 0 "
	var countc []orm.Params
	numss, errss := o.Raw(sqln).Values(&countc)
	fmt.Println(sqln)
	rowsall := ""
	if errss == nil && numss > 0 {
		rowsall = countc[0]["c"].(string)
		fmt.Println(countc[0]["c"].(string))

	}

	flist = append(flist, "id", "creatorid", "ownerid", "addtime", "edittime")
	sqls += strings.Join(flist, ",")
	sqls = " select " + sqls + "     from appmain , " + beego.ModXML[modname].Tablename + " where id = aid  and deleted = 0   limit  " + fmt.Sprintf("%d", start) + " , " + fmt.Sprintf("%d", rows)

	fmt.Println(sqls)

	var lists []orm.ParamsList

	num, err := o.Raw(sqls).ValuesList(&lists)
	var jstrs []string
	if err == nil && num > 0 {

		//fmt.Print(lists)
		var ii int64
		ii = 0
		for ii = 0; ii < num; ii++ {
			//strarr += lists[0][0].(string) + "  " + lists[0][7].(string) + "\n"
			ic := len(lists[ii])
			var iic int
			var jstr []string
			for iic = 0; iic < ic; iic++ {
				stmp := ""
				if lists[ii][iic] != nil {
					stmp = lists[ii][iic].(string)
				}

				//	jstr = append(jstr, "\""+flist[iic]+"\":\""+lists[ii][iic].(string)+"\"")
				jstr = append(jstr, "\""+flist[iic]+"\":\""+stmp+"\"")

			}
			jstrs = append(jstrs, "{"+strings.Join(jstr, ",")+"}")

		}
	}

	strarr := "{\"total\":\"" + rowsall + "\",\"rows\":[" + strings.Join(jstrs, ",") + "  ]}"
	c.Ctx.WriteString(strarr)
}

func (c *DataGet) Post() {
	o := orm.NewOrm()
	var users []*m.DqContact
	num, err := o.QueryTable("dq_contact").All(&users)
	fmt.Println(num, err, "--------------")
	var i int64
	var raw []string
	for i = 0; i < num; i++ {
		uu := users[i]
		si := rowtoJson(reflect.TypeOf(*uu), reflect.ValueOf(uu).Elem())
		raw = append(raw, si)
		fmt.Println(num, users[i])
	}
	sn := fmt.Sprintf("%d", num)
	strarr := "{\"total\":\"" + sn + "\",\"rows\":[" + strings.Join(raw, ",") + "  ]}"

	c.Ctx.WriteString(strarr)
}

func rowtoJson(vtypes reflect.Type, refmod reflect.Value) string {

	count1 := vtypes.NumField()
	var raw []string
	for i := 0; i < count1; i++ {
		fname := vtypes.Field(i).Name
		typename := vtypes.Field(i).Type.Name()
		/*
		   {"total":"7","rows":[{"id":"188275","firstname":"g","lastname":"g","phone":"sf","email":""},{"id":"188284","firstname":"asdf","lastname":"asdf","phone":"son of bitch deleted all data","email":""},{"id":"188285","firstname":"sdsd","lastname":"sdfsdf","phone":"4545","email":"sdsd@dddd.dd"},{"id":"188287","firstname":"aaa","lastname":"aaaa","phone":"","email":""},{"id":"188288","firstname":"aaaa","lastname":"aaaaa","phone":"","email":""},{"id":"188289","firstname":"aaaaa","lastname":"aaaa","phone":"","email":""},{"id":"188290","firstname":"Tizio","lastname":"Sempronio","phone":"3409585632","email":""}]}
		*/
		datai := ""
		if typename == "string" {
			//refmod.Field(i).SetString(c.GetString(fname))
			datai = refmod.Field(i).String()
		}
		if typename == "int" || typename == "int64" {
			//refmod.Field(i).SetString(c.GetString(fname))
			tmpdata := refmod.Field(i).Int()
			datai = fmt.Sprintf("%d", tmpdata)
		}

		/*
			if typename == "int" || typename == "int64" {
				intt, _ := c.GetInt64(fname, 0)
				refmod.Field(i).SetInt(intt)
			}
		*/

		raw = append(raw, fmt.Sprintf("\"%s\":\"%s\"", fname, datai))
		fmt.Print(i, "		  to json 	", fname, "				 ", typename, "\n")
	}
	s := "{" + strings.Join(raw, ",") + "}"
	fmt.Println(s)
	return s
}
func (c *MainController) Get() {
	ss := m.DqContact{}
	c.Data["FormModName"] = "DqContact"
	c.Data["Form"] = &ss
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *Xml) Get() {
	ss := m.DqContact{}
	mname := "cusman"
	c.Data["Xmod"] = beego.ModXML[mname]
	c.Data["FormModName"] = "DqContact"
	c.Data["Form"] = &ss
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "xml.tpl"
}
func saveObj(vtypes reflect.Type, refmod reflect.Value, c *DataSave) {
	count1 := vtypes.NumField()

	for i := 0; i < count1; i++ {
		fname := vtypes.Field(i).Name
		typename := vtypes.Field(i).Type.Name()

		if typename == "string" {
			refmod.Field(i).SetString(c.GetString(fname))
		}

		if typename == "int" || typename == "int64" {
			intt, _ := c.GetInt64(fname, 0)
			refmod.Field(i).SetInt(intt)
		}

		fmt.Print(i, "			", fname, "				 ", typename, "\n")
	}

}

func DataWrite(md interface{}, c *DataSave) (int64, error) {
	o := orm.NewOrm()
	id, _ := c.GetInt("Id", 0)
	if id > 0 {
		return o.Update(md)
	} else {
		return o.Insert(md)
	}
}
func (c *DataSave) Post() {

	modname := c.GetString("FormModName")

	switch modname {
	case "DqContact":
		mod := m.DqContact{}
		saveObj(reflect.TypeOf(mod), reflect.ValueOf(&mod).Elem(), c)
		DataWrite(&mod, c)
	case "DqTest":
		mod := m.DqTest{}
		saveObj(reflect.TypeOf(mod), reflect.ValueOf(&mod).Elem(), c)
		DataWrite(&mod, c)
	}

	c.Ctx.WriteString(`{"success":true }`)
}
