#ifndef _JFSDK_H_
#define _JFSDK_H_

typedef struct IMAGE_INFO_STRUCT
{
    int DataFilePTR;
}ImageInfoStruct;
typedef void  *LPVOID;

int InitImageFileFunc( ImageInfoStruct* sImageInfo, const char* Path );
int UnInitImageFileFunc( ImageInfoStruct* sImageInfo );
char * GetImageStreamFunc( ImageInfoStruct* sImageInfo, float fScale, int nImagePosX, int nImagePosY, int* nDataLength,  char** ImageStream );
int DeleteImageDataFunc( LPVOID pImageData );

#endif