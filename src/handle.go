package src

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"io"
	"log"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type ret_upload struct {
	Err_no 		int 		`json:"err_no"`
	Err_msg		string		`json:"err_msg"`
	Image_url	string		`json:"image_url"`
}

func (data *ret_upload) general_ret(errNo int , errMsg string ){
	data.Err_no = errNo
	data.Err_msg = errMsg
	log.Printf("image server err , errNo: %d errMsg: %s \n" , errNo ,  errMsg)
}


func UploadHandler(w http.ResponseWriter, r *http.Request) {

	ret := ret_upload{
		Err_no:0 ,
		Err_msg:"success" ,
	}


	//随机生成一个fileid
	//var imgid string
	//imgid=MakeImageID()

	//上传参数为uploadfile
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		ret.general_ret(-1 , err.Error())
		bret , _ := Json_marshal(ret)
		w.Write(bret)
		return
	}
	defer file.Close()
	//检测文件类型
	body , err := ioutil.ReadAll(file)

    if err != nil {
		ret.general_ret(-1 , err.Error())
		bret , _ := Json_marshal(ret)
		w.Write(bret)
		return
    }
    filetype := http.DetectContentType(body)
	imagetype := strings.Split(filetype , "/")

	if imagetype[0]!="image"{
		ret.general_ret(-1 , "file type not image")
		bret , _ := Json_marshal(ret)
		w.Write(bret)
		return
	}
	//回绕文件指针
	log.Println(filetype)
	if  _, err = file.Seek(0, 0); err!=nil{
		log.Println(err)
	}
	//提前创建整棵存储树
	if err != nil{
		log.Println(err)
	}
	//log.Println(ImageID2Path(imgid))

	// 获取文件md5
	md5Ctx := md5.New()
	md5Ctx.Write(body)
	md5 := hex.EncodeToString(md5Ctx.Sum(nil))

	dirpath , dir , err := BuildTree(md5)

	path := dirpath + "/" + md5 + "." + imagetype[1]

	if FileExist(path) {
		ret.Image_url = Conf.Domain + "/" + dir + "/" + md5 + "." + imagetype[1]
		bret , _ := Json_marshal(ret)
		w.Write(bret)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0775)
	defer f.Close()
	if err != nil {
		ret.general_ret(-1 , err.Error())
		bret , _ := Json_marshal(ret)
		w.Write(bret)
		return
	}
	io.Copy(f, file)

	ret.Image_url = Conf.Domain + "/" + dir + "/" + md5 + "." + imagetype[1]
	bret , _ := Json_marshal(ret)
	w.Write(bret)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageid := vars["imgid"]
	if len([]rune(imageid)) != 16 {
		w.Write([]byte("Error:ImageID incorrect."))
		return
	}
	imgpath := ImageID2Path(imageid)
	if !FileExist(imgpath) {
		w.Write([]byte("Error:Image Not Found."))
		return
	}

	http.ServeFile(w, r, imgpath)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><center><h1>It Works!</h1></center><hr><center>Quick Image Server</center></body></html>"))
}
