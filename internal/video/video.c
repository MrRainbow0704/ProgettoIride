#include <arv.h>
#include <assert.h>
#include <png.h>

int camera_init(ArvCamera* camera, GError** error) {
    camera = arv_camera_new(NULL, error);
    if (&error != NULL) {
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}

int camera_capture_buffer(ArvCamera* camera, ArvBuffer** outBuffer, GError** error) {
    if (ARV_IS_CAMERA(camera)) {
        ArvBuffer* buffer = arv_camera_acquisition(camera, 0, error);

        if (ARV_IS_BUFFER(buffer)) {
            *outBuffer = buffer;
            return EXIT_SUCCESS;
        }

        if (&error != NULL) {
            return EXIT_FAILURE;
        }
    } else {
        return EXIT_FAILURE;
    }
}

void camera_delete(ArvCamera* camera) {
    g_clear_object(&camera);
}

void camera_buffer_process(ArvBuffer* buffer, char* filename) {
    assert(arv_buffer_get_payload_type(buffer) == ARV_BUFFER_PAYLOAD_TYPE_IMAGE);

    size_t buffer_size;
    int width, height;
    char* buffer_data = (char*)arv_buffer_get_data(buffer, &buffer_size);                       // raw data
    arv_buffer_get_image_region(buffer, NULL, NULL, &width, &height);                           // get width/height
    int bit_depth = ARV_PIXEL_FORMAT_BIT_PER_PIXEL(arv_buffer_get_image_pixel_format(buffer));  // bit(s) per pixel
    int arv_row_stride = width * bit_depth / 8;                                                 // bytes per row, for constructing row pointers
    int color_type = PNG_COLOR_TYPE_GRAY;

    // boilerplate libpng stuff without error checking
    png_structp png_ptr = png_create_write_struct(PNG_LIBPNG_VER_STRING, NULL, NULL, NULL);
    png_infop info_ptr = png_create_info_struct(png_ptr);
    FILE* f = fopen(filename, "wb");
    png_init_io(png_ptr, f);
    png_set_IHDR(png_ptr, info_ptr, width, height, bit_depth, color_type,
                 PNG_INTERLACE_NONE, PNG_COMPRESSION_TYPE_BASE, PNG_FILTER_TYPE_BASE);
    png_write_info(png_ptr, info_ptr);

    // Need to create pointers to each row of pixels for libpng
    png_bytepp rows = (png_bytepp)(png_malloc(png_ptr, height * sizeof(png_bytep)));
    int i = 0;
    for (i = 0; i < height; ++i)
        rows[i] = (png_bytep)(buffer_data + (height - i) * arv_row_stride);
    // Actually write image
    png_write_image(png_ptr, rows);
    png_write_end(png_ptr, NULL);
    png_free(png_ptr, rows);
    png_destroy_write_struct(&png_ptr, &info_ptr);
    fclose(f);
}

void camera_buffer_clear(ArvBuffer* buffer) {
    g_clear_object(&buffer);
}