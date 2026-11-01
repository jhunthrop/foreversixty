@testable import ForeverSixty
import DesignSystem
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F11 screen and state on iPhone 17 (402 x 874), at the default size and at AX5, from the
/// contract fixtures. Recorded on the shared iOS 26.5 simulator like ReportsSnapshotTests; the
/// PNGs double as the review captures in design/reviews/2026-10-31-mobile-wave3-rankings.
@MainActor
@Suite("Rankings snapshots")
struct RankingsSnapshotTests {
    private let now = RankingsTestData.now
    private let clock = RankingsTestData.clock
    private let build = RankingsTestData.build

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: SimsSnapshotTests.safeArea, size: CGSize(width: SimsSnapshotTests.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    /// The screen as pushed, so the back control shows (root titled `back`).
    private func pushed(back: String, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(back).navigationDestination(for: Int.self) { _ in
                ScrollView { content() }
                    .background(FSColor.bg)
                    .reportsBar(title: back) { EmptyView() }
            }
        }
        .tint(FSColor.gold)
    }

    private func loaded(_ page: RankingsPage, offline: Date? = nil) -> BoardPhase {
        .loaded(BoardContent(first: page, rows: page.rows, page: 1, exhausted: false, offlineSince: offline))
    }

    private func board(
        _ phase: BoardPhase, filters: RankingsFilters = RankingsTestData.filters, account: RankingsAccount = RankingsTestData.account()
    ) -> some View {
        pushed(back: Copy.tabHome) {
            RankingsView(phase: phase, filters: filters, encounterName: "Onyxia", account: account, now: now)
        }
    }

    private var signedOutFilters: RankingsFilters {
        var filters = RankingsTestData.filters
        filters.classSlug = ""
        filters.spec = ""
        filters.ruleset = ""
        return filters
    }

    // MARK: F11.1 Rankings

    @Test func rankings() throws {
        snap(board(loaded(try RankingsTestData.page())), "rankings-ios")
    }

    @Test func rankings_large() throws {
        snap(board(loaded(try RankingsTestData.page())), "rankings-large-ios", height: 2400, large: true)
    }

    @Test func rankings_offPage() throws {
        snap(board(loaded(try RankingsTestData.offPage())), "rankings-offpage-ios")
    }

    @Test func rankings_unranked() throws {
        snap(board(loaded(try RankingsTestData.emptyPage())), "rankings-empty-ios")
    }

    @Test func rankings_signedOut() throws {
        snap(board(loaded(try RankingsTestData.mixed()), filters: signedOutFilters, account: .signedOut), "rankings-signedout-ios")
    }

    @Test func rankings_offline() throws {
        snap(board(loaded(try RankingsTestData.page(), offline: RaidTime.date("2026-12-17T18:00:00Z"))), "rankings-offline-ios")
    }

    @Test func rankings_error() {
        snap(board(.failed), "rankings-error-ios")
    }

    @Test func rankings_loading() {
        snap(board(.loading), "rankings-loading-ios")
    }

    // MARK: F11.2 Filters

    private func setup() throws -> RankingsSetup {
        let phases = try RankingsTestData.phases().phases
        return RankingsSetup(
            encounters: try RankingsTestData.encounters().encounters, phases: phases,
            defaults: RankingsDefaulting.defaults(RankingsTestData.lightning, specs: RankingsTestData.specs, phases: phases, now: now),
            encounter: "onyxia"
        )
    }

    @Test func filters() throws {
        let setup = try setup()
        let sheet = SheetFrame(height: 800, behind: board(loaded(try RankingsTestData.page()))) {
            FiltersView(filters: RankingsTestData.filters, setup: setup, apply: { _ in }, reset: {}, openEncounters: {})
        }
        snap(sheet, "rankings-filters-ios")
    }

    @Test func encounterList() throws {
        let setup = try setup()
        let sheet = SheetFrame(height: 800, behind: board(loaded(try RankingsTestData.page()))) {
            VStack(alignment: .leading, spacing: 0) {
                Text(Copy.rankFilterEncounter.uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                    .padding(.horizontal, 16).padding(.top, 16)
                EncounterList(encounters: setup.encounters, selected: "onyxia") { _ in }
            }
        }
        snap(sheet, "rankings-encounters-ios")
    }

    // MARK: F11.4 to F11.7 a character

    private func page(
        _ phase: CharacterPhase, name: String = "Voltaris", items: [Int: Item] = [:], account: RankingsAccount = RankingsTestData.account(),
        watchFailed: Bool = false
    ) throws -> some View {
        let ref = try #require(CharacterRef(key: "us/normal/\(name.lowercased())", name: name))
        return pushed(back: Copy.rankTitle) {
            CharacterView(
                ref: ref, phase: phase, items: items, account: account, specs: RankingsTestData.specs, build: build,
                watchFailed: watchFailed, now: now, clock: clock, actions: CharacterActions(openReport: { _ in })
            )
        }
    }

    private func profile(_ profile: CharacterProfile) -> CharacterPhase {
        .loaded(CharacterContent(profile: profile, offlineSince: nil))
    }

    @Test func character_other() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-other-ios")
    }

    @Test func character_yours() throws {
        snap(try page(profile(try RankingsTestData.lightningbolltz()), name: "Lightningbolltz", items: try RankingsTestData.items(),
                      account: RankingsTestData.account(history: true)), "character-yours-ios")
    }

    @Test func character_large() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-large-ios", height: 2600, large: true)
    }

    @Test func character_gear() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-gear-ios", height: 2300)
    }

    @Test func character_gear_large() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-gear-large-ios", height: 5200, large: true)
    }

    @Test func character_noGear() throws {
        let profile = try RankingsTestData.voltaris { $0["gear"] = [] as [Any]; $0["gear_gap_total"] = NSNull() }
        snap(try page(self.profile(profile)), "character-nogear-ios", height: 1700)
    }

    @Test func character_noLogs() throws {
        snap(try page(profile(try RankingsTestData.freshling()), name: "Freshling"), "character-nologs-ios")
    }

    @Test func character_notFound() throws {
        snap(try page(.notFound, name: "Voltxris"), "character-notfound-ios")
    }

    @Test func character_loading() throws {
        snap(try page(.loading), "character-loading-ios")
    }

    @Test func character_offline() throws {
        let phase = CharacterPhase.loaded(CharacterContent(profile: try RankingsTestData.voltaris(), offlineSince: RaidTime.date("2026-12-17T18:00:00Z")))
        snap(try page(phase, items: try RankingsTestData.items()), "character-offline-ios")
    }

    @Test func character_watchFailed() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items(), watchFailed: true), "character-watchfailed-ios")
    }

    /// The Watch star, off and on (and offline), drawn in a bar of its own: iOS snapshots do not
    /// draw the iOS 26 glass toolbar items.
    @Test func watchStar() throws {
        let profile = try RankingsTestData.voltaris()
        func bar(watching: Bool, offline: Bool) -> some View {
            HStack(spacing: 4) {
                Spacer()
                WatchStar(name: "Voltaris", watching: watching, offline: offline) {}
                Image(systemName: "square.and.arrow.up").foregroundStyle(FSColor.strong).frame(width: 44, height: 44)
            }
        }
        func card(watching: Bool, offline: Bool) -> some View {
            VStack(alignment: .leading, spacing: 8) {
                bar(watching: watching, offline: offline)
                CharacterHead(profile: profile, specs: RankingsTestData.specs, yours: false)
            }
            .padding(16)
            .quietCard()
        }
        let view = VStack(alignment: .leading, spacing: 20) {
            card(watching: false, offline: false)
            card(watching: true, offline: false)
            card(watching: true, offline: true)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        .background(FSColor.bg)
        snap(view, "character-watch-ios")
    }

    // MARK: Home

    @Test func home_rankingsRow() throws {
        var presenter = HomeFixtures.presenter()
        presenter.logsLine = ReportsCopy.homeLine(try ReportsTestData.list())
        presenter.rankingsLine = RankingsCopy.homeLine(page: try RankingsTestData.page(), filters: RankingsTestData.filters, encounterName: "Onyxia")
        let phase = HomePhase.loaded(try HomeFixtures.home(), offlineSince: nil)
        let view = HomeCaptureFrame {
            HomeAnswerView(phase: phase, presenter: presenter, build: build)
        }
        snap(view, "home-rankings-ios", height: 2200)
    }
}
