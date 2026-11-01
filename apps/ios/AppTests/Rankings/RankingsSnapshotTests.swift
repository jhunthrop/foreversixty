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

    /// `keyWindow` renders through the window, which draws the iOS 26 glass bar items (the Watch
    /// star and share) that a layer render leaves out.
    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false, keyWindow: Bool = false) {
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
            as: keyWindow
                ? .image(drawHierarchyInKeyWindow: true, precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits)
                : .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
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

    @Test func filters_todayOnly() throws {
        let setup = try setup()
        var today = RankingsTestData.filters
        today.todayOnly = true
        let sheet = SheetFrame(height: 800, behind: board(loaded(try RankingsTestData.page()), filters: today)) {
            FiltersView(filters: today, setup: setup, apply: { _ in }, reset: {}, openEncounters: {})
        }
        snap(sheet, "rankings-filters-today-ios")
    }

    /// A row's long-press menu, its items drawn as the system menu lists them (an off-screen
    /// capture cannot open the real one).
    @Test func rowMenu() throws {
        let page = try RankingsTestData.page()
        let row = page.rows[0]
        let menu = VStack(alignment: .leading, spacing: 0) {
            RankRowView(row: row, total: page.total, metric: .dps, yours: false, mine: false, crest: false, now: now)
                .quietCard()
                .padding(.bottom, 12)
            VStack(alignment: .leading, spacing: 0) {
                RankRowMenu(row: row, signedIn: true, watched: false, actions: RankingsActions(openReport: { _ in }))
                    .labelStyle(MenuCaptureLabelStyle())
            }
            .background(GameColors.sheet, in: .rect(cornerRadius: 14))
            .frame(width: 250)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        .background(Color.black.opacity(0.6))
        .background(board(loaded(page)))
        snap(menu, "rankings-menu-ios")
    }

    @Test func encounterList() throws {
        let setup = try setup()
        let sheet = SheetFrame(height: 800, behind: board(loaded(try RankingsTestData.page()))) {
            VStack(alignment: .leading, spacing: 0) {
                Text(Copy.rankFilterEncounter.uppercased()).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                    .padding(.horizontal, 16).padding(.top, 16)
                EncounterList(encounters: setup.encounters, selected: "onyxia") { _ in }
            }
            .frame(maxHeight: .infinity, alignment: .top)
        }
        snap(sheet, "rankings-encounters-ios")
    }

    // MARK: F11.4 to F11.7 a character

    /// The character page as pushed over Rankings, with its bar: back, the Watch star and share.
    private func page(
        _ phase: CharacterPhase, name: String = "Voltaris", items: [Int: Item] = [:], account: RankingsAccount = RankingsTestData.account(),
        watching: Bool = false, offline: Bool = false, watchFailed: Bool = false
    ) throws -> some View {
        let ref = try #require(CharacterRef(key: "us/normal/\(name.lowercased())", name: name))
        let star: WatchStarState? = if case .loaded = phase, account.signedIn {
            WatchStarState(name: name, watching: watching, offline: offline)
        } else {
            nil
        }
        return NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(Copy.rankTitle).navigationDestination(for: Int.self) { _ in
                ScrollView {
                    CharacterView(
                        ref: ref, phase: phase, items: items, account: account, specs: RankingsTestData.specs, build: build,
                        watchFailed: watchFailed, now: now, clock: clock, actions: CharacterActions(openReport: { _ in })
                    )
                }
                .background(FSColor.bg)
                .characterBar(title: name, share: ref.webURL, star: star) {}
            }
        }
        .tint(FSColor.gold)
    }

    private func profile(_ profile: CharacterProfile) -> CharacterPhase {
        .loaded(CharacterContent(profile: profile, offlineSince: nil))
    }

    @Test func character_other() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-other-ios", keyWindow: true)
    }

    @Test func character_yours() throws {
        snap(try page(profile(try RankingsTestData.lightningbolltz()), name: "Lightningbolltz", items: try RankingsTestData.items(),
                      account: RankingsTestData.account(history: true)), "character-yours-ios", keyWindow: true)
    }

    @Test func character_large() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-large-ios", height: 2600, large: true, keyWindow: true)
    }

    @Test func character_gear() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-gear-ios", height: 2300, keyWindow: true)
    }

    @Test func character_gear_large() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items()), "character-gear-large-ios", height: 5200, large: true)
    }

    @Test func character_noGear() throws {
        let profile = try RankingsTestData.voltaris { $0["gear"] = [] as [Any]; $0["gear_gap_total"] = NSNull() }
        snap(try page(self.profile(profile)), "character-nogear-ios", height: 1700, keyWindow: true)
    }

    @Test func character_noLogs() throws {
        snap(try page(profile(try RankingsTestData.freshling()), name: "Freshling"), "character-nologs-ios", keyWindow: true)
    }

    @Test func character_notFound() throws {
        snap(try page(.notFound, name: "Voltxris"), "character-notfound-ios", keyWindow: true)
    }

    @Test func character_loading() throws {
        snap(try page(.loading), "character-loading-ios", keyWindow: true)
    }

    @Test func character_offline() throws {
        let phase = CharacterPhase.loaded(CharacterContent(profile: try RankingsTestData.voltaris(), offlineSince: RaidTime.date("2026-12-17T18:00:00Z")))
        snap(try page(phase, items: try RankingsTestData.items(), offline: true), "character-offline-ios", keyWindow: true)
    }

    @Test func character_error() throws {
        snap(try page(.failed), "character-error-ios", keyWindow: true)
    }

    /// The star on in the bar (Voltaris watched); `character-other-ios` is the star off.
    @Test func character_watching() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items(), watching: true), "character-watching-ios", keyWindow: true)
    }

    /// A failed watch: the star back off and the line under the head.
    @Test func character_watchFailed() throws {
        snap(try page(profile(try RankingsTestData.voltaris()), items: try RankingsTestData.items(), watchFailed: true), "character-watchfailed-ios", keyWindow: true)
    }

    // MARK: Home

    /// Lightningbolltz, the Elemental Shaman the board places: Home's row says his ladder.
    private static let lightningHome = HomeCharacter(
        key: "us/normal/lightningbolltz", region: "us", ruleset: "normal", name: "Lightningbolltz", race: "troll",
        className: "shaman", specName: "Elemental", level: 60, guild: "Hordebreakers", faction: "horde"
    )

    @Test func home_rankingsRow() throws {
        var presenter = HomeFixtures.presenter(character: Self.lightningHome)
        presenter.logsLine = ReportsCopy.homeLine(try ReportsTestData.list())
        presenter.rankingsLine = RankingsCopy.homeLine(page: try RankingsTestData.page(), filters: RankingsTestData.filters, encounterName: "Onyxia")
        let home = try HomeFixtures.home("character.home.yours")
        let view = HomeCaptureFrame(title: Self.lightningHome.name) {
            HomeAnswerView(phase: .loaded(home, offlineSince: nil), presenter: presenter, build: home.dataBuild)
        }
        snap(view, "home-rankings-ios", height: 2200)
    }

    /// Lightningbolltz's Your gear: his Home answer's upgrades over the shaman items file, the
    /// same dots his character page draws.
    @Test func home_gearYours() throws {
        let home = try HomeFixtures.home("character.home.yours")
        let rows = GearModel.rows(home: home, gear: [:], items: try RankingsTestData.items())
        let view = HomeCaptureFrame(title: Copy.gearTitle, large: true) {
            GearView(rows: rows, intro: Copy.gearIntro(band: "60 · raid"), build: home.dataBuild)
        }
        snap(view, "home-gear-yours-ios", height: 1300)
    }
}

/// A system menu row as captured: the title, then the glyph at the trailing edge, 44 pt.
private struct MenuCaptureLabelStyle: LabelStyle {
    func makeBody(configuration: Configuration) -> some View {
        HStack {
            configuration.title.font(ScaledType.body).foregroundStyle(FSColor.strong)
            Spacer()
            configuration.icon.foregroundStyle(FSColor.strong)
        }
        .padding(.horizontal, 16)
        .frame(minHeight: 44)
        .overlay(alignment: .bottom) { FSColor.lineSoft.frame(height: 1) }
    }
}
