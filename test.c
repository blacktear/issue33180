#include "stdio.h"
#include "libs/jfsdk.h"

//gcc test.c -L./libs -lImageOperationLib && ./a.out

int main() {
	ImageInfoStruct h;
	int r;

	r = InitImageFileFunc(&h, "demo.kfb");
	if (r != 1) {
		printf("open file failed!\n");
		return 0;
	}

	int len;
	char* imageStream;

	int i;
	for (i = 0; i < 10; i++) {
		GetImageStreamFunc(&h, 10, 1*256, 1*256, &len, &imageStream);
		DeleteImageDataFunc((LPVOID)imageStream);
		printf("leave msg: index = %d len = %d\n", i, len);
	}

	UnInitImageFileFunc(&h);
}
