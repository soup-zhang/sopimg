package sopimg

import (
	"fmt"
	"testing"
)

func TestMAIL(t *testing.T){
	var mc ImgInfo
	mc.Base64 =	""  //base64字符串
	mc.Type ="png"
	mc.Path = "./static/image"
	re,err :=mc.Base64ToFile()

	fmt.Println("re:",re,"err:",err)

}

