package main

/*
#cgo CFLAGS: -I ./libs
#cgo LDFLAGS: -L ./libs -lImageOperationLib
typedef struct IMAGE_INFO_STRUCT
{
    int DataFilePTR;
}ImageInfoStruct;
typedef void  *LPVOID;

int InitImageFileFunc( ImageInfoStruct* sImageInfo, const char* Path );
int UnInitImageFileFunc( ImageInfoStruct* sImageInfo );
char * GetImageStreamFunc( ImageInfoStruct* sImageInfo, float fScale, int nImagePosX, int nImagePosY, int* nDataLength,  char** ImageStream );
int DeleteImageDataFunc( LPVOID pImageData );
*/
import "C"

//import "log"

func main() {
	var h C.ImageInfoStruct
	//log.Println(" ") //but I put here[fmt | log].Println(" ") or [fmt | log].Print(" ") The problem is solved, I don't know why.

	r := C.InitImageFileFunc(&h, C.CString("demo.kfb"))
	if int(r) != 1 {
		panic("open file failed!")
	}
	defer C.UnInitImageFileFunc(&h)

	var dataLength C.int
	var imageStream *C.char

	for i := 0; i < 10; i++ {
		C.GetImageStreamFunc(&h, 10, 1*256, 1*256, &dataLength, &imageStream)
		C.DeleteImageDataFunc(C.LPVOID(imageStream))
		println("leave msg: index =", i, ",len =", dataLength)
	}
}
