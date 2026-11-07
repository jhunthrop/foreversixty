import Darwin
import SnapshotTesting
import SwiftUI
import UIKit

/// The native parity gate's text side (design/parity.md "Native"): a state's words, strictly. The
/// accessibility tree's labels and values in reading order are a text golden beside the image
/// golden, so a copy change fails until it is re-recorded on purpose.
enum ParityText {
    /// The accessibility labels and values of `view` rendered at `size`, one line per element in
    /// reading order (`label: value` when an element has both).
    @MainActor
    static func lines(of view: some View, size: CGSize) -> [String] {
        AccessibilityAutomation.enable()
        return ParityHost.render(view, size: size) { host in elements(under: host.view).compactMap(line) }
    }

    /// Asserts the state's words against its text golden (`__Snapshots__/<suite>/<name>.txt`).
    @MainActor
    static func assertText(
        of view: some View,
        size: CGSize,
        named name: String,
        fileID: StaticString = #fileID,
        file: StaticString = #filePath,
        testName: String = #function,
        line: UInt = #line,
        column: UInt = #column
    ) {
        assertSnapshot(
            of: lines(of: view, size: size).joined(separator: "\n"),
            as: .lines,
            named: name,
            fileID: fileID,
            file: file,
            testName: testName,
            line: line,
            column: column
        )
    }

    /// Every accessibility element under `root`, depth first, in the order VoiceOver reads them.
    @MainActor
    static func elements(under root: NSObject) -> [NSObject] {
        if root.isAccessibilityElement { return [root] }
        return children(of: root).flatMap { elements(under: $0) }
    }

    /// An element's accessibility children: its listed elements, the indexed ones a SwiftUI
    /// hosting view vends, else its subviews.
    @MainActor
    private static func children(of root: NSObject) -> [NSObject] {
        if let listed = root.accessibilityElements as? [NSObject], !listed.isEmpty {
            return listed
        }
        let count = root.accessibilityElementCount()
        if count > 0, count != NSNotFound {
            return (0 ..< count).compactMap { root.accessibilityElement(at: $0) as? NSObject }
        }
        return (root as? UIView)?.subviews ?? []
    }

    @MainActor
    private static func line(_ element: NSObject) -> String? {
        let label = element.accessibilityLabel?.trimmingCharacters(in: .whitespacesAndNewlines) ?? ""
        let value = element.accessibilityValue?.trimmingCharacters(in: .whitespacesAndNewlines) ?? ""
        switch (label.isEmpty, value.isEmpty) {
        case (true, true): return nil
        case (false, true): return label
        case (true, false): return value
        case (false, false): return "\(label): \(value)"
        }
    }
}

/// SwiftUI builds its accessibility tree only while an assistive technology or UI automation is
/// on. The test host turns automation on the way XCUITest does (libAccessibility's
/// `_AXSSetAutomationEnabled`), once, so the text golden can read the tree. Test code only.
private enum AccessibilityAutomation {
    private typealias SetEnabled = @convention(c) (Int32) -> Void

    @MainActor private static var enabled = false

    @MainActor
    static func enable() {
        guard !enabled else { return }
        guard let library = dlopen("/usr/lib/libAccessibility.dylib", RTLD_NOW),
              let symbol = dlsym(library, "_AXSSetAutomationEnabled")
        else {
            fatalError("ParityText: libAccessibility's automation switch is not available")
        }
        unsafeBitCast(symbol, to: SetEnabled.self)(1)
        enabled = true
    }
}
