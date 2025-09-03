#ifndef _IRIDE_INTERNAL_VIDEO_H_
#define _IRIDE_INTERNAL_VIDEO_H_
#include <arv.h>

int camera_init(ArvCamera* camera, GError* error);
void camera_delete(ArvCamera* camera);
int camera_capture_buffer(ArvCamera* camera, void* buffer, GError* error);
void camera_buffer_process(ArvBuffer* buffer, char* filename);
void camera_buffer_clear(ArvBuffer* buffer);

#endif  // _IRIDE_INTERNAL_VIDEO_H_