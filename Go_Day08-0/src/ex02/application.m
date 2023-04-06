#import "application.h"
#include "_cgo_export.h"

void InitApplication() {
    [NSApplication sharedApplication];
}

void RunApplication() {
    @autoreleasepool {
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [NSApp activateIgnoringOtherApps:YES];
        [NSApp run];
    }
}
