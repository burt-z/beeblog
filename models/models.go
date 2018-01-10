package models

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//分类信息
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Commnet struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

//文章信息
type Topic struct {
	Id              int64
	Uid             int64
	Title           string //默认的字节长度是255
	Content         string `orm:"size(5000)"` //内容，比较长这里默认的是写的比较大
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       int64 `orm:"index"`
	ReplyCount      string
	ReplyLastUserId string
	Category        string
}

func RegisterDB() {

	_, err := os.Stat("//Users/zhuweijie/go/src/sample/beeblog/data/beeblog.db")

	if os.IsNotExist(err) {
		fmt.Println("数据库不存在")
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Commnet))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) //10 是连接数
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	fmt.Println("name名称", name)
	var cate *Category = new(Category)
	cate.Title = name
	cate.TopicTime = time.Now()
	cate.Created = time.Now()
	//	var cate *Category = &Category{Title: name} //创建一个对象
	qs := o.QueryTable("category") //查询哪一个表
	//	Filter("title",name) 的意思是查询表里ititle字段与name 一致 ，One（）方法是因为只有一个而不是一个slice，所以用one
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		fmt.Println("没查到,")
		return err //说明没有查询到
	}
	//	没有查询到执行插入的操作
	fmt.Println("Category", cate.Title)
	_, err = o.Insert(cate)
	if err != nil {
		//		执行插入报错
		fmt.Println("插入报错")
		return err
	}
	return nil
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	fmt.Println("数据cates", cates)
	return cates, err
}

//删除的操作
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64) //10进制，int64
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	var cate *Category = new(Category) //删除的时候必须指定ID
	cate.Id = cid
	_, err = o.Delete(cate)
	return err
}
func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:    title,
		Content:  content,
		Created:  time.Now(),
		Updated:  time.Now(),
		Category: category,
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}
	//	更新分类
	cate := new(Category)
	qs := o.QueryTable("category")
	fmt.Println("category", category)
	err = qs.Filter("title", category).One(cate)

	fmt.Println("查询到重复的err", err)

	return err

}

func GetAllTopics(cate string, IsDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if IsDesc {
		fmt.Println("倒序排序")
		if len(cate) > 0 {
			qs = qs.Filter("category", cate) //这个判断的意思在cate有值时说明是点击的分类，进行过滤按照分类，这里需要进行赋值，不赋值的话会出现过滤无效，过滤之后需要保存
			fmt.Println("qsqs", qs)
		}
		_, err = qs.OrderBy("-created").All(&topics) //按照传入的标签进行排序  需要加一个杠
	} else {
		_, err = qs.All(&topics)
	}
	fmt.Println("topic顺序", topics)
	return topics, err
}
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++ //浏览次数增加
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		topic.Category = category
		o.Update(topic)
	}
	return nil
}
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	comment := &Commnet{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	_, err = o.Insert(comment)
	return err
}
func GetAllReplies(tid string) ([]*Commnet, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	replies := make([]*Commnet, 0)
	qs := o.QueryTable("commnet")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}
func DeleteReply(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Commnet{Id: tidNum}
	_, err = o.Delete(reply)
	return err
}
