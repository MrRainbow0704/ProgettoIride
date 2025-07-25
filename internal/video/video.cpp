extern "C" {
char* captureFrame() {
    char f[] = "this/is/a/file.test";
    int size = sizeof(f) * sizeof(char);
    // char* outfile = (char*)malloc(size);
    char* outfile = new char[size];
    return f;
}
}