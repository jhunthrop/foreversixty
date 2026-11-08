import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// One native parity state (design/parity.md "Native"), the four checks every screen lands with:
/// the regions and inventory against `design/mocks/app/regions/ios-<screen>.json`, the text golden,
/// and the image goldens at the default and the largest (AX5) text size.
@MainActor
enum ParityCheck {
    /// A known drift (parity.md "When it fails"): the diff it excuses, from when, and which lane
    /// owns it. The test fails once the drift is gone, so the list only shrinks.
    struct Drift {
        let prefix: String
        let since: String
        let lane: String
    }

    /// EXPECTED_DRIFT, per screen.
    static let expectedDrift: [String: [Drift]] = [:]

    /// Room under the content for the safe area and the screen's own bottom padding.
    static let textSlack: CGFloat = 200

    /// The device width, and a height that holds every reported region and item.
    static func textSize(_ size: CGSize, capture: ParityRegions.Capture) -> CGSize {
        let bottoms = capture.regions.values.map(\.maxY) + capture.items.map(\.frame.maxY)
        return CGSize(width: size.width, height: max(size.height, (bottoms.max() ?? 0) + textSlack))
    }

    static func assert(
        _ view: some View,
        screen: String,
        size: CGSize,
        fileID: StaticString = #fileID,
        file: StaticString = #filePath,
        testName: String = #function,
        line: UInt = #line,
        column: UInt = #column
    ) async throws {
        await ParityText.prepare()
        let screenView = view.preferredColorScheme(.dark).environment(\.locale, Locale(identifier: "en_US"))
        let capture = ParityRegions.capture(of: screenView, size: size)
        let diffs = ParityRegions.compare(capture, with: try ParityRegions.reference(screen: screen))
        let known = expectedDrift[screen] ?? []
        let unexpected = diffs.filter { diff in !known.contains { diff.hasPrefix($0.prefix) } }
        #expect(unexpected.isEmpty, "ios-\(screen) parity:\n\(unexpected.joined(separator: "\n"))")
        for drift in known where !diffs.contains(where: { $0.hasPrefix(drift.prefix) }) {
            Issue.record("ios-\(screen): the known drift \"\(drift.prefix)\" (\(drift.since), \(drift.lane)) is gone; remove it from EXPECTED_DRIFT")
        }
        for (suffix, category) in [("", UIContentSizeCategory.large), ("-ax5", .accessibilityExtraExtraExtraLarge)] {
            assertSnapshot(
                of: screenView,
                as: .image(
                    precision: SnapshotTolerance.precision,
                    perceptualPrecision: SnapshotTolerance.perceptualPrecision,
                    layout: .fixed(width: size.width, height: size.height),
                    traits: UITraitCollection { traits in
                        traits.userInterfaceStyle = .dark
                        traits.preferredContentSizeCategory = category
                    }
                ),
                named: "ios-\(screen)\(suffix)",
                fileID: fileID, file: file, testName: testName, line: line, column: column
            )
        }
        // The words of the whole content, as the regions count it: the text pass renders a window
        // tall enough that nothing scrolls.
        ParityText.assertText(
            of: screenView, size: textSize(size, capture: capture), named: "ios-\(screen)",
            fileID: fileID, file: file, testName: testName, line: line, column: column
        )
    }
}
