package main

import (
	//"errors"
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
	"net"
	"os"
	"strconv"
	//"strings"
	"time"
)

var counter int = 0
var conn_host string = ""
var conn_port string = ""
var serverPort string = ""

func errCheck(e error, s string) bool {
	if e != nil {
		fmt.Println(s)
		return true
	} else {
		return false
	}
}

func main() {

	conn_host = os.Args[1]
	conn_port = os.Args[2]
	serverPort = os.Args[3]
	fmt.Println("connection  ", conn_host, ":", conn_port)
	RunTCPServer(handleRequestCmd, serverPort)

}

// Handles incoming requests.
func handleRequestCmd(conn net.Conn, i int) {

	cmdbuf := make([]byte, 5)
	_, err := conn.Read(cmdbuf)

	if errCheck(err, "Problem getting command") {
		return
	}

	cmdStr := string(cmdbuf)

	fmt.Println("Received command ", cmdStr)
	fileName := RecvText(conn, "FILE0")
	fmt.Println("File to process is ", fileName)
	processImages(fileName)
}

func processImages(fileName string) {

	img := gocv.IMRead(fileName, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Unable to read Image file")
		return
	} else {
		go faceDetection(img)
	}
}

func faceDetection(img gocv.Mat) {

	defer img.Close()

	//xmlFile := "../cascade/haarcascade_frontalface_alt.xml"
	xmlFile := "./xmlfile.xml"

	// color for the rect when faces detected
	//blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	// detect faces
	rects := classifier.DetectMultiScale(img)
	fmt.Printf("found %d faces\n", len(rects))

	// draw a rectangle around each face on the original image,
	// along with text identifying as "Human"
	for _, r := range rects {
		gocv.Rectangle(&img, r, red, 2)
	}

	fileName := "/img/" + strconv.FormatInt(time.Now().Unix(), 10) + "_fd_image.jpg"
	b := gocv.IMWrite(fileName, img)
	if !b {
		fmt.Println("Writing Mat to file failed")
		return
	}
	fmt.Println("Just saved " + fileName)
	ConnectToSend(conn_host, conn_port, "RSP02", fileName)
}
