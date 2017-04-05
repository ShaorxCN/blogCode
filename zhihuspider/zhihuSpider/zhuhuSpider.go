package zhihuSpider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ShaorxCN/blogCode/zhihuspider/spider"
	simplejson "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/it512/sqlt"
	//"github.com/jmoiron/sqlx"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var Db *sql.DB

func init() {

	for i := 0; i < 50; i++ {
		spider.AddToPool(new(zhihuSpider))
	}

	Db, _ = sql.Open("mysql", "spider:123456@tcp(127.0.0.1:3306)/zhihuspider?charset=utf8")
	//loader := sqlt.NewDefaultSqlLoader("template/*.tpl")
	log.Println(Db.Ping())

}

type zhihuSpider struct {
	url string
	//html     []byte
	document *goquery.Document
	user     *User
	jsons    *simplejson.Json
}

//教育
type education struct {
	school string
	major  string
}

//工作
type job struct {
	company  string
	position string
}

//用户
type User struct {
	Name     string
	UserName string
	HeadLine string
	Vocation string
	//0:male 1:female 2:unknow
	Sex        int
	Educations []education
	Jobs       []job
	Locations  []string
}

//SetUrl implement Method SetUrl
func (s *zhihuSpider) SetUrl(url string) {
	s.url = url
}

//implement Method GetHtml
func (s *zhihuSpider) GetHtml() {
	resp, err := http.Get(s.url)
	if err != nil {
		log.Panicf("get Html from %s fail,跳过该url :%v", s.url, err)
	}

	defer resp.Body.Close()
	s.document, _ = goquery.NewDocument(s.url)

	log.Printf("get html from url:%s end ", s.url)
}

//implement Method GetUrls,暂定的是获取99页
func (s *zhihuSpider) GetUrls() []string {
	var todoUrls []string
	log.Printf("get urls from url:%s  ", formatUrl(s.user.UserName))

	urls := getFollowingPages(formatUrl(s.user.UserName), s.document)

	for _, url := range urls {

		resp, err := http.Get(url)

		if err != nil {
			log.Panicln(err)
		}

		defer resp.Body.Close()

		document, err := goquery.NewDocumentFromResponse(resp)

		if err != nil {
			log.Panicln(err)
		}

		se, _ := document.Find("div#data").Attr("data-state")

		js, _ := simplejson.NewJson([]byte(se))

		jc, _ := js.Get("entities").Get("users").Map()

		for k, _ := range jc {
			if k == s.user.UserName || k == "null" {
				continue
			}
			todoUrls = append(todoUrls, formatUserUrl(k))
		}

	}

	return todoUrls

}

//implement Method Analy
func (s *zhihuSpider) Analy() {

	user := new(User)

	user.Name = s.document.Find("span.ProfileHeader-name").Text()
	user.UserName = s.url[strings.LastIndex(s.url, "/")+1:]
	if s.document.Find("svg.Icon.Icon--male").Nodes != nil {
		user.Sex = 0
	} else if s.document.Find("svg.Icon.Icon--female").Nodes != nil {
		user.Sex = 1
	} else {
		user.Sex = 2
	}

	jsdata, _ := s.document.Find("div#data").Attr("data-state")
	js, err := simplejson.NewJson([]byte(jsdata))
	if err != nil {
		log.Panicf("get json from %s fail,跳过该url ", s.url)
	}

	s.jsons = js

	educations, _ := js.Get("entities").Get("users").Get(user.UserName).Get("educations").Array()

	var es []education
	for _, v := range educations {

		e := education{school: v.(map[string]interface{})["school"].(map[string]interface{})["name"].(string), major: v.(map[string]interface{})["major"].(map[string]interface{})["name"].(string)}
		es = append(es, e)
	}

	user.Educations = es
	user.Vocation, _ = js.Get("entities").Get("users").Get(user.UserName).Get("business").Get("name").String()

	jobs, _ := js.Get("entities").Get("users").Get(user.UserName).Get("employments").Array()
	user.HeadLine, _ = js.Get("entities").Get("users").Get(user.UserName).Get("headline").String()
	var bs []job
	for _, v := range jobs {

		j := job{company: v.(map[string]interface{})["company"].(map[string]interface{})["name"].(string), position: v.(map[string]interface{})["job"].(map[string]interface{})["name"].(string)}
		bs = append(bs, j)
	}

	locations, _ := js.Get("entities").Get("users").Get(user.UserName).Get("locations").Array()
	var ltmp []string
	for _, v2 := range locations {
		ltmp = append(ltmp, v2.(map[string]interface{})["name"].(string))
	}

	user.Jobs = bs
	user.Locations = ltmp

	s.user = user

}

//implement Method SaveData
func (s *zhihuSpider) SaveData() {
	log.Println(s.user)
	st, err := Db.Prepare("insert spider_user set username=?,name=?,vacation=?,headLine=?,sex=?")
	defer st.Close()

	if err != nil {
		log.Panicf("save data from url : %s failed %v", s.url, err)
	}

	_, err = st.Exec(s.user.UserName, s.user.Name, s.user.Vocation, s.user.HeadLine, s.user.Sex)
	if err != nil {
		log.Panicf("save data from url : %s failed %v", s.url, err)
	}
	log.Printf("save data from url:%s end ", s.url)
}

//拼接url
func formatUrl(urlToken string) string {
	return fmt.Sprintf("https://www.zhihu.com/people/%s/following", urlToken)
}

func formatUserUrl(urlToken string) string {
	return fmt.Sprintf("https://www.zhihu.com/people/%s", urlToken)
}

//获取following的链接
func getFollowingPages(url string, doc *goquery.Document) []string {

	var urls []string
	urls = append(urls, url)
	//判断关注的页数
	var n int = 1
	max := doc.Find("div.Pagination").Find("button")

	if max.Length() < 3 {
		return urls
	}

	maxInt, _ := strconv.Atoi(max.Eq(max.Length() - 2).Text())

	if maxInt > 99 {
		n = 99
	} else {
		n = maxInt
	}
	for i := 2; i < n; i++ {
		urls = append(urls, fmt.Sprintf("%s?page=%d", url, i))
	}

	return urls
}
