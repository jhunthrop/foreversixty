@testable import ForeverSixty
import DesignTokens
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The Talents screen's states (spec F6 §6, §7) at iPhone 17 size, at the default text size and at
/// AX5, with the repo's own icon files. Goldens are rendered on iOS 26.5 (the shared simulator and CI).
@MainActor
@Suite("Talents snapshots")
struct TalentsSnapshotTests {
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

    /// iPhone SE 3, the smallest supported phone (spec F6 §1).
    private static let iPhoneSE = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 20, left: 0, bottom: 0, right: 0),
        size: CGSize(width: 375, height: 667),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    private let shaman: TalentsFile

    init() throws {
        shaman = try TalentTestData.shaman()
    }

    private func presenter(
        yours: String? = TalentTestData.shamanYours,
        guide: TalentGuideBuild = .build(talents: TalentTestData.shamanGuide, band: 60)
    ) -> TalentsPresenter {
        let export = yours.map { TalentExport.ranks(TalentRules.ranks(ofBand: $0)) } ?? .none
        return TalentsPresenter(file: shaman, guide: guide, export: export, level: 60, specName: "Elemental")
    }

    private func screen(_ phase: TalentsPhase, shown: TalentBuildKind = .yours, tree: Int = 0, offline: String? = nil) -> some View {
        NavigationStack {
            TalentsView(phase: phase, build: BisTestData.build, offlineStamp: offline, shown: .constant(shown), tree: .constant(tree))
                .navigationTitle(Copy.talentsTitle)
                .navigationBarTitleDisplayMode(.inline)
                .toolbarBackground(FSColor.bg, for: .navigationBar)
                .toolbarBackground(.visible, for: .navigationBar)
        }
    }

    private func snap(_ view: some View, size: TextSize = .standard, device: ViewImageConfig = iPhone17, named name: String) {
        let traits = UITraitCollection { mutable in
            mutable.userInterfaceStyle = .dark
            mutable.preferredContentSizeCategory = size.category
        }
        assertSnapshot(
            of: view.environment(\.itemIconSource, RepoItemIconSource()).preferredColorScheme(.dark),
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: device), traits: traits),
            named: size.rawValue,
            testName: name
        )
    }

    private func sheet(_ name: String, build: TalentBuildKind = .yours) throws -> TalentSheetModel {
        let id = try #require(TalentRules.trees(shaman).flatMap(\.talents).first { $0.name == name }).id
        return try #require(presenter().sheet(build, talentID: id))
    }

    @Test("differs: dashes, chips, dots and the notes", arguments: TextSize.allCases)
    func differs(size: TextSize) {
        snap(screen(.loaded(presenter(), offlineSince: nil)), size: size, named: "talents-differs")
    }

    @Test("matches")
    func matches() {
        snap(screen(.loaded(presenter(yours: TalentTestData.shamanGuide), offlineSince: nil)), named: "talents-matches")
    }

    @Test("the guide segment")
    func guide() {
        snap(screen(.loaded(presenter(), offlineSince: nil), shown: .guide), named: "talents-guide")
    }

    @Test("no export: the guide alone")
    func noExport() {
        snap(screen(.loaded(presenter(yours: nil), offlineSince: nil), shown: .guide), named: "talents-noexport")
    }

    @Test("no build: your build alone")
    func noBuild() {
        snap(screen(.loaded(presenter(guide: .none), offlineSince: nil)), named: "talents-nobuild")
    }

    @Test("no spec")
    func noSpec() {
        snap(screen(.loaded(presenter(guide: .noSpec), offlineSince: nil)), named: "talents-nospec")
    }

    @Test("loading skeletons")
    func loading() {
        snap(screen(.loading), named: "talents-loading")
    }

    @Test("offline under the strip")
    func offline() {
        snap(screen(.loaded(presenter(), offlineSince: BisTestData.now), offline: "9:41 pm"), named: "talents-offline")
    }

    @Test("error with Try again")
    func failed() {
        snap(screen(.failed), named: "talents-error")
    }

    @Test("rank sheet: the bottom-row capstone, locked, with its requirements")
    func rankBottomRow() throws {
        snap(TalentRankContent(model: try sheet("Elemental Mastery"), build: BisTestData.build), named: "talents-rank-bottom-row")
    }

    @Test("rank sheet: the right column with the guide line", arguments: TextSize.allCases)
    func rankRightColumn(size: TextSize) throws {
        snap(TalentRankContent(model: try sheet("Call of Flame"), build: BisTestData.build), size: size, named: "talents-rank-right-column")
    }

    @Test("rank sheet: maxed, no Next rank")
    func rankMaxed() throws {
        snap(TalentRankContent(model: try sheet("Convection"), build: BisTestData.build), named: "talents-rank-maxed")
    }

    @Test("the Restoration tab: its dot, its one dashed cell")
    func restorationTab() {
        snap(screen(.loaded(presenter(), offlineSince: nil), tree: 2), named: "talents-restoration")
    }

    @Test("the first screen on an iPhone SE")
    func smallPhone() {
        snap(screen(.loaded(presenter(), offlineSince: nil)), device: Self.iPhoneSE, named: "talents-differs-se")
    }

    @Test("rank sheet at AX5, drawn whole to its last line")
    func rankLargeEnd() throws {
        let tall = ViewImageConfig(safeArea: .zero, size: CGSize(width: 402, height: 1700), traits: UITraitCollection(userInterfaceStyle: .dark))
        snap(TalentRankContent(model: try sheet("Call of Flame"), build: BisTestData.build), size: .ax5, device: tall, named: "talents-rank-large-end")
    }
}
