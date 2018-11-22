package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"net/http"
	"text/template"
	"os"
	"time"
)

func status(w http.ResponseWriter, r *http.Request) { //展示界面


	fileName :=time.Now().Format("20060102150405")+".png"

	tableString:=""

	r.ParseForm()
	for i, v := range r.Form {
		if i == "pngValue" {
			if v[0] != "null" {
				data := []byte(v[0])
				data = data[22:]
				s := string(data)
				BASE64TOPNG(fileName,s)
				yHat, max :=Similarity(fileName)
				fmt.Printf("%d",max)
				tableString = "  <table> <tr> <th>Number</th>	<th>Similarity</th>	</tr>"
				for i:=0;i<10;i++{

					if i==max{
						tableString +=fmt.Sprintf(" <tr > <td ><strong>%d</strong></td>	<td><strong'>%f%%</strong></td></tr>", i,yHat.Cell[i][0]*100)
					}else{
						tableString +=fmt.Sprintf("<tr> <td>%d</td>	<td>%f%%</td></tr>", i,yHat.Cell[i][0]*100)

					}
				}
				tableString +=fmt.Sprintf(" </table><p><strong>The best match is %d!</strong></p>",max)

			}
		}
	}


	t, err := template.ParseFiles("./index.html") //使用text/template（html/template会导致输出部分被引号包围无法使用）
	if err != nil {
		panic(err)
		return
	}




	TableData := map[string]interface{}{"Data": tableString} //替换Status.html中的{{.ClientData}}
	t.Execute(w, TableData) //使用http.ResponseWrite输出网页

}

func main() {


	fmt.Printf("starting...\n")
	http.HandleFunc("/", status) //展示页面
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err:=http.ListenAndServe(":0718", nil) //创建服务器
	if err!=nil{
		fmt.Printf("%v",err)
	}
}

func Similarity(fileName string)(LKmath.Matrix, int){

	file, err := os.Open("./imageStore/"+fileName)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	img = resize.Resize(28, 28, img, resize.Lanczos3)

	X := LKmath.NewEmptyMatrix(28*28, 1)

	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			R, G, B, _ := img.At(j, i).RGBA()
			X.Cell[i*28+j][0] = 1 - ((0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)) / 256 / 256)
		}
	}

	hw := LKmath.ReadFromJson("hwT800.json")

	yHat, _ := hw.ForwardPropagation(X)
	yHat.Hprint("yhat")

	MaxNo :=0
	MaxValue :=yHat.Cell[0][0]

	for i:=1;i<10;i++{
		if MaxValue<yHat.Cell[i][0]{
			MaxValue=yHat.Cell[i][0]
			MaxNo =i
		}
	}

	return yHat, MaxNo

}

func BASE64TOPNG(imageName string,ImgBase64 string) {
	unbased, err := base64.StdEncoding.DecodeString(ImgBase64)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		fmt.Printf("Bad png")
	}

	f, err := os.OpenFile("./imageStore/"+imageName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("Cannot open file")
	}
	defer f.Close()

	png.Encode(f, im)
}
