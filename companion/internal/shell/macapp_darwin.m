//go:build darwin && !nogui

// companion/internal/shell/macapp_darwin.m
// The Objective-C half of closehook_darwin.go. The classes it extends are
// the webview library's window delegate and the systray library's app
// delegate; methods are added at run time because neither library offers
// a hook. Built without ARC, like both libraries.
#import <Cocoa/Cocoa.h>
#include <objc/runtime.h>
#include <stdbool.h>
#include "_cgo_export.h"

// A close request hides the window (when Go says so) and never closes it:
// a closed window would end the library's run loop, and with it the
// menu-bar icon.
static BOOL fsWindowShouldClose(id self, SEL _cmd, id sender) {
	if (fsGoWindowShouldClose()) {
		[(NSWindow *)sender orderOut:nil];
	}
	return NO;
}

bool fsInstallCloseHook(void *window) {
	NSWindow *w = (NSWindow *)window;
	id delegate = [w delegate];
	if (delegate == nil) {
		return false;
	}
	Class cls = object_getClass(delegate);
	SEL sel = @selector(windowShouldClose:);
	class_addMethod(cls, sel, (IMP)fsWindowShouldClose, "c@:@");
	if (class_getMethodImplementation(cls, sel) != (IMP)fsWindowShouldClose) {
		return false; // the class answers it already, in its own way
	}
	// AppKit may have asked what the delegate answers when it was set.
	[w setDelegate:nil];
	[w setDelegate:delegate];
	return true;
}

void fsShowWindow(void *window) {
	NSWindow *w = (NSWindow *)window;
	[NSApp activateIgnoringOtherApps:YES];
	if ([w isMiniaturized]) {
		[w deminiaturize:nil];
	}
	[w makeKeyAndOrderFront:nil];
}

void fsHideWindow(void *window) {
	[(NSWindow *)window orderOut:nil];
}

// A Quit from the menu, the Dock or the tray library asks Go to stop and is
// cancelled here: the companion stops the way the tray's Quit stops it,
// flushing what it holds, and the process ends when Go is done. A logout,
// restart or shutdown cannot wait for that, or the system says the
// companion interrupted it, so Go is told to stop and the app ends now;
// the upload queue is on disk and resumes at the next launch.
static NSApplicationTerminateReply fsShouldTerminate(id self, SEL _cmd, id app) {
	fsGoQuit();
	NSAppleEventDescriptor *event = [[NSAppleEventManager sharedAppleEventManager] currentAppleEvent];
	if (event != nil && ([event attributeDescriptorForKeyword:kAEQuitReason] != nil ||
	                     [event paramDescriptorForKeyword:kAEQuitReason] != nil)) {
		return NSTerminateNow; // the system's quit carries its reason; a player's does not
	}
	return NSTerminateCancel;
}

// A click on the Dock icon brings the window back.
static BOOL fsShouldHandleReopen(id self, SEL _cmd, id app, BOOL visible) {
	fsGoReopen();
	return NO;
}

static NSMenuItem *fsItem(NSString *title, SEL action, NSString *key, NSEventModifierFlags mods) {
	NSMenuItem *item = [[NSMenuItem alloc] initWithTitle:title action:action keyEquivalent:key];
	[item setKeyEquivalentModifierMask:mods];
	return [item autorelease];
}

static NSMenuItem *fsSubmenu(NSString *title, NSArray<NSMenuItem *> *items) {
	NSMenu *menu = [[[NSMenu alloc] initWithTitle:title] autorelease];
	for (NSMenuItem *item in items) {
		[menu addItem:item];
	}
	NSMenuItem *holder = [[[NSMenuItem alloc] initWithTitle:title action:nil keyEquivalent:@""] autorelease];
	[holder setSubmenu:menu];
	return holder;
}

// The menu bar: the app menu, Edit (Command-C, Command-V and the rest only
// work in a web view through these items) and Window, whose Close hides.
static NSMenu *fsMainMenu(void) {
	NSEventModifierFlags cmd = NSEventModifierFlagCommand;
	NSMenu *bar = [[[NSMenu alloc] initWithTitle:@""] autorelease];
	[bar addItem:fsSubmenu(@"Forever Sixty", @[
		fsItem(@"Hide Forever Sixty", @selector(hide:), @"h", cmd),
		fsItem(@"Hide Others", @selector(hideOtherApplications:), @"h", cmd | NSEventModifierFlagOption),
		fsItem(@"Show All", @selector(unhideAllApplications:), @"", 0),
		[NSMenuItem separatorItem],
		fsItem(@"Quit Forever Sixty", @selector(terminate:), @"q", cmd),
	])];
	[bar addItem:fsSubmenu(@"Edit", @[
		fsItem(@"Undo", @selector(undo:), @"z", cmd),
		fsItem(@"Redo", @selector(redo:), @"z", cmd | NSEventModifierFlagShift),
		[NSMenuItem separatorItem],
		fsItem(@"Cut", @selector(cut:), @"x", cmd),
		fsItem(@"Copy", @selector(copy:), @"c", cmd),
		fsItem(@"Paste", @selector(paste:), @"v", cmd),
		fsItem(@"Select All", @selector(selectAll:), @"a", cmd),
	])];
	[bar addItem:fsSubmenu(@"Window", @[
		fsItem(@"Minimize", @selector(performMiniaturize:), @"m", cmd),
		fsItem(@"Close", @selector(performClose:), @"w", cmd),
	])];
	return bar;
}

void fsInstallAppHooks(void) {
	NSApplication *app = [NSApplication sharedApplication];
	// A bundled app is a regular app already; a bare binary run from a
	// terminal is not, and would get no Dock icon and no menu bar.
	[app setActivationPolicy:NSApplicationActivationPolicyRegular];
	id delegate = [app delegate];
	if (delegate != nil) {
		Class cls = object_getClass(delegate);
		class_addMethod(cls, @selector(applicationShouldTerminate:), (IMP)fsShouldTerminate, "Q@:@");
		class_addMethod(cls, @selector(applicationShouldHandleReopen:hasVisibleWindows:), (IMP)fsShouldHandleReopen, "c@:@c");
		// NSApplication reads what its delegate answers when it is set.
		[app setDelegate:nil];
		[app setDelegate:delegate];
	}
	[app setMainMenu:fsMainMenu()];
}
