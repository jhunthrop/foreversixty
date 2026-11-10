//go:build darwin

// companion/internal/install/system_darwin.m
// The Objective-C half of system_darwin.go: App Translocation's original
// path, the Trash, and the login item. Built without ARC, like the shell's.
#import <Foundation/Foundation.h>
#import <ServiceManagement/ServiceManagement.h>
#include <dlfcn.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

static char *fsCopyString(NSString *s) {
	return strdup(s == nil ? "" : [s UTF8String]);
}

// SecTranslocate is in Security.framework but declared only in a private
// header, so it is looked up at run time, as LetsMove does.
typedef Boolean (*fsIsTranslocatedFn)(CFURLRef, bool *, CFErrorRef *);
typedef CFURLRef (*fsOriginalFn)(CFURLRef, CFErrorRef *);

char *fsOriginalPath(const char *path) {
	@autoreleasepool {
		void *sec = dlopen("/System/Library/Frameworks/Security.framework/Security", RTLD_LAZY);
		if (sec == NULL) {
			return NULL;
		}
		fsIsTranslocatedFn isTranslocated = (fsIsTranslocatedFn)dlsym(sec, "SecTranslocateIsTranslocatedURL");
		fsOriginalFn original = (fsOriginalFn)dlsym(sec, "SecTranslocateCreateOriginalPathForURL");
		if (isTranslocated == NULL || original == NULL) {
			return NULL;
		}
		NSURL *url = [NSURL fileURLWithPath:[NSString stringWithUTF8String:path]];
		bool translocated = false;
		if (!isTranslocated((CFURLRef)url, &translocated, NULL) || !translocated) {
			return NULL;
		}
		CFURLRef out = original((CFURLRef)url, NULL);
		if (out == NULL) {
			return NULL;
		}
		char *result = fsCopyString([(NSURL *)out path]);
		CFRelease(out);
		return result;
	}
}

char *fsTrash(const char *path) {
	@autoreleasepool {
		NSURL *url = [NSURL fileURLWithPath:[NSString stringWithUTF8String:path]];
		NSError *err = nil;
		if ([[NSFileManager defaultManager] trashItemAtURL:url resultingItemURL:nil error:&err]) {
			return NULL;
		}
		return fsCopyString([err localizedDescription]);
	}
}

bool fsLoginItems(void) {
	if (@available(macOS 13.0, *)) {
		return true;
	}
	return false;
}

char *fsSetLoginItem(bool on) {
	@autoreleasepool {
		if (@available(macOS 13.0, *)) {
			SMAppService *service = [SMAppService mainAppService];
			NSError *err = nil;
			BOOL ok = on ? [service registerAndReturnError:&err] : [service unregisterAndReturnError:&err];
			// Unregistering what was never registered is not a failure.
			if (ok || (!on && [service status] == SMAppServiceStatusNotRegistered)) {
				return NULL;
			}
			return fsCopyString([err localizedDescription]);
		}
		return fsCopyString(@"opening at login needs macOS 13 or later");
	}
}
