#import <Cocoa/Cocoa.h>

void* Window_Create(int x, int y, int width, int height, const char* title);
void Window_MakeKeyAndOrderFront(void *wndPtr);