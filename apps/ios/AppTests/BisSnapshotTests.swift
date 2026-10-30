@testable import ForeverSixty
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The BiS tab's states (spec F3) at iPhone 17 size, at the default text size and at AX5.
/// Recorded on the iPhone 17 simulator; re-record with `SNAPSHOT_TESTING_RECORD=all` and review
/// the PNG diff. Icons do not load here (no icon source), so each tile shows its quality ring.
@MainActor
@Suite("BiS snapshots")
struct BisSnapshotTests {
    enum TextSize: String, CaseIterable {
        case standard, ax5

        var category: UIContentSizeCategory {
            self == .standard ? .large : .accessibilityExtraExtraExtraLarge
        }
    }

    private static let iPhone17 = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0),
        size: CGSize(width: 402, height: 874),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    private let list: BisList
    private let context: BisContext
    private let catalog: SpecCatalog

    init() throws {
        let file = try BisTestData.bis()
        let band = try #require(file.band(for: BisTestData.raid.range, faction: "horde"))
        list = BisListBuilder(items: try BisTestData.items().items, worn: [12640], timeZone: TimeZone.gmt)
            .list(file: file, band: band, specLabel: "Fury Warrior", rangeLabel: Copy.bisRangeP1)
        context = BisContext(
            selection: BisTestData.raid,
            characterSelection: BisTestData.raid,
            characterName: "Obnoxious Yell",
            classes: try BisTestData.classes()
        )
        catalog = try BisTestData.catalog()
    }

    private func view(_ phase: BisPhase, context: BisContext? = nil) -> some View {
        BisView(
            phase: phase,
            context: context ?? self.context,
            catalog: catalog,
            searchText: .constant(""),
            now: BisTestData.now
        )
    }

    private func snap(_ view: some View, size: TextSize, named name: String) {
        let traits = UITraitCollection { mutable in
            mutable.userInterfaceStyle = .dark
            mutable.preferredContentSizeCategory = size.category
        }
        assertSnapshot(
            of: view.preferredColorScheme(.dark),
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: Self.iPhone17), traits: traits),
            named: size.rawValue,
            testName: name
        )
    }

    @Test("the list, Yours, with the worn pick", arguments: TextSize.allCases)
    func listState(size: TextSize) {
        snap(view(.loaded(list, offlineSince: nil)), size: size, named: "list")
    }

    @Test("offline: the saved list under the strip", arguments: TextSize.allCases)
    func offline(size: TextSize) {
        let saved = BisTestData.now.addingTimeInterval(-86_400)
        snap(view(.loaded(list, offlineSince: saved)), size: size, named: "offline")
    }

    @Test("loading: 17 skeleton rows", arguments: TextSize.allCases)
    func loading(size: TextSize) {
        snap(view(.loading), size: size, named: "loading")
    }

    @Test("error with Try again", arguments: TextSize.allCases)
    func failed(size: TextSize) {
        snap(view(.failed), size: size, named: "error")
    }

    @Test("offline and never fetched", arguments: TextSize.allCases)
    func notCached(size: TextSize) {
        snap(view(.notCached), size: size, named: "not-cached")
    }

    @Test("the picker sheet, away from the character's list", arguments: TextSize.allCases)
    func picker(size: TextSize) {
        var away = context
        away.selection = BisTestData.raid.withFaction("alliance")
        snap(BisPickerSheet(context: away, catalog: catalog, onSelect: { _ in }), size: size, named: "picker")
    }

    @Test("the item tooltip of a pick", arguments: TextSize.allCases)
    func tooltip(size: TextSize) throws {
        let row = try #require(list.rows.first)
        let pick = try #require(row.pick)
        snap(ItemTooltipSheet(content: BisTooltip.content(row: row, pick: pick, list: list)), size: size, named: "tooltip")
    }
}
