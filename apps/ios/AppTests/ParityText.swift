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
        let lines = ParityHost.render(view, size: size) { host in settledLines(under: host.view) }
        // A host made while the switch was still taking can stay treeless; a fresh one is not.
        return lines.isEmpty ? ParityHost.render(view, size: size) { host in settledLines(under: host.view) } : lines
    }

    /// How long the tree may take to appear: SwiftUI builds it asynchronously once automation is
    /// on, and the first screen after the switch (on CI, Activity's) can read before it exists.
    static let settleTimeout: TimeInterval = 3
    static let settleStep: TimeInterval = 0.05

    /// The lines once the tree has appeared and two reads in a row agree, so a state never
    /// records a half-built (or empty) tree.
    @MainActor
    private static func settledLines(under root: UIView) -> [String] {
        let deadline = Date().addingTimeInterval(settleTimeout)
        var last = elements(under: root).compactMap(line)
        while Date() < deadline {
            RunLoop.main.run(until: Date().addingTimeInterval(settleStep))
            let next = elements(under: root).compactMap(line)
            if !next.isEmpty, next == last { return next }
            last = next
        }
        if last.isEmpty { print("ParityText: no accessibility lines; the tree:\n" + dump(root, depth: 0).joined(separator: "\n")) }
        return last
    }

    /// The accessibility tree as the walk sees it, for a state that reads as empty.
    @MainActor
    private static func dump(_ node: NSObject, depth: Int) -> [String] {
        guard depth < 12 else { return [] }
        let pad = String(repeating: "  ", count: depth)
        let head = "\(pad)\(type(of: node)) element=\(node.isAccessibilityElement) count=\(node.accessibilityElementCount()) "
            + "listed=\((node.accessibilityElements ?? []).count) label=\(node.accessibilityLabel ?? "-")"
        guard !node.isAccessibilityElement else { return [head] }
        return [head] + children(of: node).prefix(8).flatMap { dump($0, depth: depth + 1) }
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
    private static let warmUp: TimeInterval = 10

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
        // On a fresh CI simulator the first screen rendered after the switch never builds its tree
        // (Activity's, the first parity suite, read as empty). Render a probe until its tree
        // appears, so every real screen comes after the switch has taken.
        let deadline = Date().addingTimeInterval(warmUp)
        while Date() < deadline {
            let ready = ParityHost.render(Text(verbatim: "probe").accessibilityLabel("probe"), size: CGSize(width: 100, height: 100)) { host in
                RunLoop.main.run(until: Date().addingTimeInterval(0.1))
                return !ParityText.elements(under: host.view).isEmpty
            }
            if ready { return }
        }
    }
}
