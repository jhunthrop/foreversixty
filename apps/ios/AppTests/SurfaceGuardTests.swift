import Foundation
import Testing

/// Only the DesignSystem primitives draw a surface (mobile redesign spec 3.4, the web's
/// surface-guard.test.ts). Outside `Sources/DesignSystem`, a token colour as a background
/// (`bg`, the screen ground, aside), a stroked rounded box, and the retired styles (`primaryCard`,
/// `quietCard`, `surfaceBand`, `MicroLabel`, `TagView`, the Barlow Condensed `ScaledType` heads, the
/// gradient `BrandButtonStyle`, caps titles) fail: use Plate, FSSection, FSRow, FSTag and FSButton.
///
/// It ships as a ratchet: `allowance` holds the files still drawing their own, with their counts.
/// A file may only go down; when it does, lower or delete its line, so the list only shrinks.
@Suite("Surface guard")
struct SurfaceGuardTests {
    static let appRoot = HomeFixtures.repoRoot.appending(path: "apps/ios")
    static let scanned = ["App", "Widgets", "NotificationContent"]

    /// What a page may not draw for itself.
    static let sources = [
        #"\.background\(\s*(?:FSColor\.(?!bg\b)\w+|GameColors\.\w+)"#,
        #"\.overlay\(\s*(?:RoundedRectangle|Capsule)\([^)]*\)\s*\.stroke"#,
        #"\.(?:primaryCard|quietCard|surfaceBand)\("#,
        #"\b(?:MicroLabel|TagView)\("#,
        #"\bScaledType\.(?:title|h2|stat|sec|cardHead)\b"#,
        #"\bBrandButtonStyle\("#,
        #"\bText\([^\n]*\.uppercased\(\)"#,
        #"\.textCase\(\.uppercase\)"#,
    ]
    static let patterns = sources.compactMap { try? NSRegularExpression(pattern: $0) }

    /// The files still drawing their own surfaces, as of the restyle of the secondary screens
    /// (flows spec section 2). Only ever lower a count or delete a line.
    static let allowance: [String: Int] = [
        "App/Account/AccountArt.swift": 2,
        "App/Account/DeleteAccountView.swift": 5,
        "App/Account/SignedOutView.swift": 12,
        "App/Account/YouView.swift": 8,
        "App/Bis/BisPickerSheet.swift": 6,
        "App/Bis/BisView.swift": 7,
        "App/Common/NewDataStrip.swift": 2,
        "App/Common/SegmentBar.swift": 2,
        "App/Guides/ChangesView.swift": 7,
        "App/Guides/GuideBody.swift": 8,
        "App/Guides/GuideParts.swift": 7,
        "App/Guides/GuideReaderView.swift": 3,
        "App/Guides/GuidesHomeView.swift": 9,
        "App/Guides/TiersView.swift": 5,
        "App/Guild/Views/CalendarScreen.swift": 5,
        "App/Guild/Views/GuildCommonViews.swift": 5,
        "App/Guild/Views/HistoryScreen.swift": 2,
        "App/Guild/Views/RosterScreen.swift": 5,
        "App/Home/GearView.swift": 3,
        "App/Home/HomeAnswerView.swift": 16,
        "App/Home/HomeScreen.swift": 1,
        "App/Items/ItemTooltipSheet.swift": 7,
        "App/Planner/PlannerEditorView.swift": 9,
        "App/Planner/PlannerSheets.swift": 11,
        "App/Planner/SharedBuildView.swift": 4,
        "App/Planner/YourBuildsView.swift": 7,
        "App/Premium/PlanViews.swift": 9,
        "App/Raids/Views/CompositionView.swift": 1,
        "App/Raids/Views/RaidComponents.swift": 5,
        "App/Raids/Views/RaidNightView.swift": 4,
        "App/Raids/Views/RaidPageView.swift": 14,
        "App/Raids/Views/RaidsTabView.swift": 17,
        "App/Raids/Views/RosterViews.swift": 7,
        "App/Raids/Views/SignupViews.swift": 10,
        "App/Rankings/CharacterView.swift": 16,
        "App/Rankings/FiltersSheet.swift": 8,
        "App/Rankings/RankingsView.swift": 11,
        "App/Reports/CoachCardView.swift": 6,
        "App/Reports/MetersView.swift": 5,
        "App/Reports/ReportView.swift": 10,
        "App/Search/ItemPageView.swift": 8,
        "App/Search/SearchRows.swift": 5,
        "App/Search/SearchTalentSheet.swift": 1,
        "App/Search/SearchViews.swift": 4,
        "App/Shell/CharacterPortrait.swift": 1,
        "App/Shell/ProfileDrawer.swift": 1,
        "App/Shell/ProfileMenu.swift": 1,
        "App/Sims/SimResultView.swift": 3,
        "App/Sims/SimRowCards.swift": 1,
        "App/Sims/SimsListView.swift": 9,
        "App/System/IntentViews.swift": 2,
        "App/Talents/TalentRankSheet.swift": 5,
        "App/Talents/TalentTreeView.swift": 3,
        "App/Talents/TalentsView.swift": 5,
        "App/Widgets/LiveSurfaceSettings.swift": 2,
        "Widgets/Shared/Views/LiveActivityViews.swift": 2,
        "Widgets/Shared/Views/NextRaidViews.swift": 2,
        "Widgets/Shared/Views/UpgradeViews.swift": 1,
        "Widgets/Shared/Views/WidgetChrome.swift": 3,
    ]

    /// The source without its comments, so a doc comment naming a retired style never counts.
    static func code(_ text: String) -> String {
        let blocks = try? NSRegularExpression(pattern: #"/\*[\s\S]*?\*/"#)
        let unblocked = blocks?.stringByReplacingMatches(
            in: text, range: NSRange(text.startIndex..., in: text), withTemplate: ""
        ) ?? text
        return unblocked.split(separator: "\n", omittingEmptySubsequences: false)
            .filter { !$0.trimmingCharacters(in: .whitespaces).hasPrefix("//") }
            .joined(separator: "\n")
    }

    static func uses(in text: String) -> Int {
        let code = code(text)
        let range = NSRange(code.startIndex..., in: code)
        return patterns.reduce(0) { $0 + $1.numberOfMatches(in: code, range: range) }
    }

    /// Every scanned Swift file's count, keyed by its path under apps/ios.
    static func scan() throws -> [String: Int] {
        var counts: [String: Int] = [:]
        let manager = FileManager.default
        for folder in scanned {
            let root = appRoot.appending(path: folder)
            guard let walker = manager.enumerator(at: root, includingPropertiesForKeys: nil) else { continue }
            for case let url as URL in walker where url.pathExtension == "swift" {
                let count = uses(in: try String(contentsOf: url, encoding: .utf8))
                if count > 0 {
                    counts[String(url.path.dropFirst(appRoot.path.count + 1))] = count
                }
            }
        }
        return counts
    }

    @Test("no file outside the design system draws more surfaces than its allowance")
    func noNewSurfaces() throws {
        let counts = try Self.scan()
        #expect(!counts.isEmpty || Self.allowance.isEmpty, "the scan found no Swift files under \(Self.appRoot.path)")
        for (file, count) in counts.sorted(by: { $0.key < $1.key }) {
            let allowed = Self.allowance[file] ?? 0
            #expect(
                count <= allowed,
                "\(file) draws \(count) retired surface(s) or style(s), allowed \(allowed): use Plate, FSSection, FSRow, FSTag, FSButton and the type roles"
            )
        }
    }

    @Test("the allowance only shrinks: a file that dropped a surface lowers its line")
    func allowanceShrinks() throws {
        let counts = try Self.scan()
        for (file, allowed) in Self.allowance.sorted(by: { $0.key < $1.key }) {
            let count = counts[file] ?? 0
            #expect(count >= allowed, "\(file) is down to \(count) from \(allowed): lower its allowance (delete the line at 0)")
        }
    }

    @Test("the patterns catch each retired style and pass the primitives' own calls")
    func patternsCatch() {
        #expect(Self.patterns.count == Self.sources.count, "a guard pattern does not compile")
        #expect(Self.uses(in: ".background(FSColor.raised, in: .rect(cornerRadius: 6))") == 1)
        #expect(Self.uses(in: ".background(FSColor.bg)") == 0)
        #expect(Self.uses(in: ".overlay(RoundedRectangle(cornerRadius: 4).strokeBorder(FSColor.line))") == 1)
        #expect(Self.uses(in: ".quietCard()\n.primaryCard()\n.surfaceBand()") == 3)
        #expect(Self.uses(in: "MicroLabel(title)\nTagView(text: t, color: c)") == 2)
        #expect(Self.uses(in: ".font(ScaledType.title)\n.font(ScaledType.meta)") == 1)
        #expect(Self.uses(in: ".buttonStyle(BrandButtonStyle(.primary))") == 1)
        #expect(Self.uses(in: "Text(title.uppercased())\nlet region = code.uppercased()") == 1)
        #expect(Self.uses(in: "// .quietCard() in a comment\n/* MicroLabel( */") == 0)
        #expect(Self.uses(in: "Plate { rows }\nFSRow(title: t) { lead }") == 0)
    }
}
