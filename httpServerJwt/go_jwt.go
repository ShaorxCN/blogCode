package main
/**
	token 分三部分:part1.part.part3以这样的格式存在
	第一部分指定类型（这里就是jwt）以及算法（part1是base64转码后的）
	第二部分就是一些指定的数据依旧base64转码
	第三部分，那part1.part2 为材料，配合指定的算法以及密钥加密作为part3

	所以服务端校验可以先解码第一第二部分，然后那解码第三部分来验证
 */
import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	//"time"
	"errors"
	"time"
)

const (
	//密钥
	SingelKey = "jwtexgkond"
)

//自定义claims
type CustomeClaim struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type UserVerfiy struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	NeedToken bool   `json:"needToken"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

//type User struct{
//	ID int `json:"id"`
//	Name string `json:"name"`
//	UserName string `json:"username"`
//}

func writeJson(res interface{}, rw http.ResponseWriter) {
	json, err := json.Marshal(res)

	if err != nil {
		log.Panicln(err)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Write([]byte(json))
}

//获取token
func generateJWTHandler(w http.ResponseWriter, r *http.Request) {
	var user UserVerfiy
	var res Response
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Success = false
		res.Message = "invalid request"
		writeJson(res, w)
		return
	}
	if user.Password != "textjwt" {
		w.WriteHeader(http.StatusForbidden)
		res.Success = false
		res.Message = "invalid username or password"
	} else {
		w.WriteHeader(http.StatusOK)
		//这边自定义的claims生成token
		claims := CustomeClaim{1, jwt.StandardClaims{ExpiresAt: time.Now().Add(30 * time.Second).Unix(),
			IssuedAt: time.Now().Unix()}}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(SingelKey))
		//token := jwt.New(jwt.SigningMethodHS256)
		//claims := make(jwt.MapClaims)
		//claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		//claims["iat"] = time.Now().Unix()
		//token.Claims = claims
		//log.Println(token)
		//tokenString, err := token.SignedString([]byte(SingelKey))

		if err != nil {
			res.Success = false
			res.Message = "internel service error"
		}
		res.Success = true
		res.Message = "success"
		res.Token = tokenString

	}
	writeJson(res, w)
	return
}

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var res Response
	if r.URL.String() != "/login" {
		//token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		//	func(token *jwt.Token) (interface{}, error) {
		//		return []byte(SingelKey), nil
		//	})
		token, err := jwt.Parse(r.Header.Get("Authorization"), func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(SingelKey), nil
			} else {
				return nil, errors.New("interval error")
			}
		})
		if err == nil && token.Valid {
			if v, ok := token.Claims.(jwt.MapClaims); ok {
				log.Println(v["id"], v)
			}
			next(w, r)
		} else {
			res.Message = "授权失败"
			res.Success = false
			writeJson(res, w)
		}

	} else {
		next(w, r)
	}

}

func RecoveryHanlder(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	next(rw, req)
}

func main() {
	//自定义生成
	negroni.Classic()
	mux := http.NewServeMux()
	mux.HandleFunc("/goal", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})
	mux.HandleFunc("/login", generateJWTHandler)

	n := negroni.New(negroni.HandlerFunc(RecoveryHanlder), negroni.NewLogger(), negroni.HandlerFunc(ValidateToken))
	n.UseHandler(http.Handler(mux))
	n.Run(":8080")

}
