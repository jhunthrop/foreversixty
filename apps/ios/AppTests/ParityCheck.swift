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

    /// EXPECTED_DRIFT, per screen. Lanes M1i to M5i (2026-11-10) took the restyled tabs to their
    /// round-5 mocks; what is left is text metrics: the runner's Chromium sets Open Sans a few points
    /// narrower than iOS, so `Main hand · Fel Ancient, Ashenvale` holds one line in the mock and wraps
    /// on the phone, moving New for you 13 down.
    static let expectedDrift: [String: [Drift]] = [
        // The web system's rows set the mock's Send me 12 shorter since batch H; Notifications keeps its
        // rows until the secondary screens' restyle.
        "notification-settings": [
            Drift(prefix: "section-when: y", since: "2026-11-10", lane: "secondary restyle"),
            Drift(prefix: "section-discord: y", since: "2026-11-10", lane: "secondary restyle"),
        ],
        // Raid night's answers are the tab roots' 44 pt answer control (one answer control everywhere);
        // the mock still draws the old 36 pt chips, so the sign-up section is 8 taller.
        "raid-night-live": nightAnswersDrift + [bandHead("section-your-sign-up")],
        "raid-night-signup": nightAnswersDrift + [bandHead("section-your-sign-up")],
        // The raid flows' heads sit on the band under the clear bar (N3), 16 above the content; the
        // mocks draw the head on the ground under an opaque bar, so every region below is ~20 lower.
        "raid-unanswered": ["controls", "section-not-answered", "section-healers-on-the-bench"].map(bandHead),
        // Home's raid-day card draws the one answer control (44 pt and its 1 pt ring, lane mobile-rest's
        // brand ruling: one answer control everywhere); the mock's older answers are shorter, so New for
        // you sits 8 to 9 lower and Next upgrades 13 to 14. Clears when the Home mock draws the segmented answers.
        "home-raid-day": [
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "mobile-rest"),
            Drift(prefix: "section-next-upgrades: y", since: "2026-11-11", lane: "batch K one answer control"),
        ],
        "home-raid-day-late": [
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "mobile-rest"),
            Drift(prefix: "section-next-upgrades: y", since: "2026-11-11", lane: "batch K one answer control"),
        ],
        "home-raid-day-officer": [
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "mobile-rest"),
            Drift(prefix: "section-next-upgrades: y", since: "2026-11-11", lane: "batch K one answer control"),
        ],
        "raid-applicant": applicantBandDrift,
        // Round 3 mocks set `.sub` lines at 18.29 and the Segmented inside a 1 px border; the app's row
        // line is RowType.sub's 18. Over the Raid tab's and Activity's ~30 lines that sums to 8 to 11 pt
        // by the last section, inside each section's own tolerance.
        "raid": [Drift(prefix: "section-olympus-xxvii: y", since: "2026-11-10", lane: "row line box 18.29")],
        "raid-signup": [Drift(prefix: "section-olympus-xxvii: y", since: "2026-11-10", lane: "row line box 18.29")],
        // The Guild tab (lane guild-tab): the same 18.29 row line box as the Raid tab's over ~40 rows, so the
        // guild's sections under the raid sit 9 to 17 pt higher than the mock's; each section's own height holds.
        "guild-tab": ["section-guild", "section-roster", "section-loot", "section-recent-logs", "section-progression", "section-pages"]
            .map(guildRowDrift),
        "guild-tab-officer": [
            "section-your-job", "section-who-is-coming", "section-before-you-pull", "section-guild", "section-roster", "section-loot",
            "section-recent-logs", "section-progression", "section-pages",
        ].map(guildRowDrift),
        // Under 60 the Close to 60 rows add four more 18.29 line boxes, so the sections under the roster sit 13 to
        // 15 pt higher than the mock's.
        "guild-tab-leveler": ["section-recent-logs", "section-progression", "section-pages"].map(guildRowDrift),
        // The recruiting rows' trail (Apply, or Applied Mon) sits 13 shorter than the mock's `.tr` margin.
        "guild-tab-none": [
            guildRowDrift("section-find-your-guild"),
            Drift(prefix: "section-guilds-recruiting-a-fury-warrior: height", since: "2026-11-11", lane: "guild-tab row trail"),
        ],
        "guild-tab-none-leveler": [guildRowDrift("section-find-your-guild")],
        "raid-officer": [
            Drift(prefix: "section-olympus-xxvii: y", since: "2026-11-10", lane: "row line box 18.29"),
        ],
        // Round-4 mock: Home's raid line title (`Onyxia's Lair · Sat 19 Dec, 7:00 pm CST`) holds one
        // line in the runner's Chromium and wraps on the phone, so the block is 16 pt taller.
        "home-60": [
            Drift(prefix: "countdowns: height", since: "2026-11-10", lane: "round-4 mocks"),
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "round-4 mocks"),
        ],
        "home-60-noavatar": [
            Drift(prefix: "countdowns: height", since: "2026-11-10", lane: "round-4 mocks"),
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "round-4 mocks"),
        ],
        "home-60-answered": [
            Drift(prefix: "countdowns: height", since: "2026-11-10", lane: "round-4 mocks"),
            Drift(prefix: "section-new-for-you: y", since: "2026-11-10", lane: "round-4 mocks"),
        ],
        "activity": [Drift(prefix: "section-this-week: y", since: "2026-11-10", lane: "row line box 18.29")],
    ]

    private static let nightAnswersDrift = [
        Drift(prefix: "section-your-sign-up: height", since: "2026-11-10", lane: "secondary restyle"),
        Drift(prefix: "section-who-is-coming: y", since: "2026-11-10", lane: "secondary restyle"),
        Drift(prefix: "section-before-you-pull: y", since: "2026-11-10", lane: "secondary restyle"),
    ]

    private static let applicantBandDrift = ["section-raid-ready", "section-their-note", "section-guild-history"].map(bandHead)

    /// A region moved down only because its screen's head now sits on the band (N3, iOS round 1).
    private nonisolated static func bandHead(_ region: String) -> Drift {
        Drift(prefix: "\(region): y", since: "2026-11-10", lane: "N3 band head")
    }

    /// A Guild tab region moved only by the row line box over the sections above it (lane guild-tab).
    private nonisolated static func guildRowDrift(_ region: String) -> Drift {
        Drift(prefix: "\(region): y", since: "2026-11-11", lane: "guild-tab row line box 18.29")
    }

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
