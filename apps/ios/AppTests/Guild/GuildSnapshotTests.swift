@testable import ForeverSixty
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F16 screen and state on iPhone 17 (402 x 874), at the default size and at AX5. The
/// PNGs, named by board cell, double as the review captures in
/// design/reviews/2026-10-31-mobile-wave3-guild/build/ios.
@MainActor
@Suite("Guild extras snapshots")
struct GuildSnapshotTests {
    static let safeArea = UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0)

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: Self.safeArea, size: CGSize(width: 402, height: height), traits: traits)
        let sized = view
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    /// The screen as pushed, so the back control shows (root titled `back`), like SimsSnapshotTests.
    private func pushed(_ title: String, back: String, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(back).navigationDestination(for: Int.self) { _ in
                content()
                    .navigationTitle(title)
                    .navigationBarTitleDisplayMode(.large)
            }
        }
    }

    // MARK: F16.1, F16.2 Roster

    private func roster(_ phase: GuildLoadPhase<GuildHome>, query: RosterQuery = RosterQuery()) -> some View {
        pushed(Copy.rosterTitle, back: Copy.tabRaids) {
            RosterContentView(phase: phase, clock: RaidTime.clock, now: { RaidTime.now }, query: .constant(query), open: { _ in }, retry: {})
                .toolbar { ToolbarItem(placement: .topBarTrailing) { Image(systemName: "magnifyingglass") } }
        }
    }

    private func loadedRoster(offline: Date? = nil) throws -> GuildLoadPhase<GuildHome> {
        .loaded(try GuildFixture.home(), offlineSince: offline)
    }

    @Test func rosterAll() throws { snap(roster(try loadedRoster()), "f16.1-roster-all", height: 1400) }
    @Test func rosterAll_xl() throws { snap(roster(try loadedRoster()), "f16.1-roster-all-xl", height: 2600, large: true) }

    @Test func rosterTanks() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(role: .tanks)), "f16.1-roster-tanks")
    }

    @Test func rosterHealers() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(role: .healers)), "f16.1-roster-healers", height: 1100)
    }

    @Test func rosterSortAttendance() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(sort: .attendance)), "f16.1-roster-sort-attendance", height: 3000)
    }

    @Test func rosterSortAttendance_xl() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(sort: .attendance)), "f16.1-roster-sort-attendance-xl", height: 2600, large: true)
    }

    @Test func rosterNoMatch() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(role: .tanks, search: "mend", searching: true)), "f16.2-no-match")
    }

    @Test func rosterNoMatch_xl() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(role: .tanks, search: "mend", searching: true)), "f16.2-no-match-xl", height: 1600, large: true)
    }

    @Test func rosterSearch() throws {
        snap(roster(try loadedRoster(), query: RosterQuery(search: "ra", searching: true)), "f16.1-roster-search")
    }

    @Test func rosterLoading() throws { snap(roster(.loading), "f16.1-roster-loading") }
    @Test func rosterError() throws { snap(roster(.failed), "f16.1-roster-error") }

    @Test func rosterOffline() throws {
        snap(roster(try loadedRoster(offline: RaidTime.now.addingTimeInterval(-3600))), "f16.1-roster-offline")
    }

    // MARK: F16.3 Raid history

    private func history(
        _ phase: GuildLoadPhase<HistoryContent>, paging: HistoryPagingState = HistoryPagingState()
    ) -> some View {
        pushed(Copy.historyTitle, back: Copy.tabRaids) {
            HistoryContentView(
                phase: phase, paging: paging, yourKeys: [GuildFixture.viewerKey], clock: RaidTime.clock, now: { RaidTime.now },
                open: { _ in }, loadMore: {}, retry: {}
            )
        }
    }

    private func page1(offline: Date? = nil) throws -> GuildLoadPhase<HistoryContent> {
        let page = try GuildFixture.historyPage1()
        return .loaded(HistoryContent(rows: page.rows, nextCursor: offline == nil ? page.nextCursor : nil, offlineSince: offline), offlineSince: offline)
    }

    private func bothPages() throws -> GuildLoadPhase<HistoryContent> {
        let rows = try GuildFixture.historyPage1().rows + GuildFixture.historyPage2().rows
        return .loaded(HistoryContent(rows: rows, nextCursor: nil, offlineSince: nil), offlineSince: nil)
    }

    @Test func historyPage1() throws { snap(history(try page1()), "f16.3-history-page1", height: 2200) }
    @Test func historyPage1_xl() throws { snap(history(try page1()), "f16.3-history-page1-xl", height: 2600, large: true) }
    @Test func historyPage2() throws { snap(history(try bothPages()), "f16.3-history-page2", height: 2400) }
    @Test func historyLoadingMore() throws {
        snap(history(try page1(), paging: HistoryPagingState(loadingMore: true)), "f16.3-history-loading-more", height: 2200)
    }

    @Test func historyMoreError() throws {
        snap(history(try page1(), paging: HistoryPagingState(moreFailed: true)), "f16.3-history-more-error", height: 2300)
    }

    @Test func historyEmpty() throws {
        snap(history(.loaded(HistoryContent(rows: [], nextCursor: nil, offlineSince: nil), offlineSince: nil)), "f16.3-history-empty")
    }

    @Test func historyLoading() throws { snap(history(.loading), "f16.3-history-loading") }
    @Test func historyError() throws { snap(history(.failed), "f16.3-history-error") }

    @Test func historyOffline() throws {
        snap(history(try page1(offline: RaidTime.now.addingTimeInterval(-3600))), "f16.3-history-offline")
    }

    // MARK: F16.4, F16.5 Calendar

    private static let link = "https://api.foreversixty.gg/v1/calendar/Rm9yZXZlciBTaXh0eSBmaXh0dXJlLCBub3QgcmVhbCE.ics"

    private func calendar(_ state: CalendarScreenState) -> some View {
        pushed(Copy.calendarTitle, back: Copy.tabYou) {
            CalendarContentView(state: state, actions: CalendarActions())
        }
    }

    private func on(_ which: CalendarWhich = .all) -> CalendarScreenState {
        CalendarScreenState(sync: .on, settings: CalendarSettings(enabled: true, which: which), link: .ready(Self.link))
    }

    @Test func calendarOn() { snap(calendar(on()), "f16.4-calendar-on", height: 1000) }
    @Test func calendarOn_xl() { snap(calendar(on()), "f16.4-calendar-on-xl", height: 3200, large: true) }
    @Test func calendarOff() { snap(calendar(CalendarScreenState(link: .ready(Self.link))), "f16.4-calendar-off", height: 1000) }
    @Test func calendarWhichMine() { snap(calendar(on(.mine)), "f16.4-which-mine", height: 1000) }

    @Test func calendarAccessOff() {
        snap(calendar(CalendarScreenState(sync: .accessOff, link: .ready(Self.link))), "f16.5-access-off", height: 1150)
    }

    @Test func calendarAccessOff_xl() {
        snap(calendar(CalendarScreenState(sync: .accessOff, link: .ready(Self.link))), "f16.5-access-off-xl", height: 3800, large: true)
    }

    @Test func calendarRemoved() { snap(calendar(CalendarScreenState(sync: .removed, link: .ready(Self.link))), "f16.4-calendar-removed", height: 1000) }

    @Test func calendarLinkStates() {
        snap(calendar(CalendarScreenState(link: .elsewhere)), "f16.4-link-elsewhere", height: 1000)
    }

    @Test func calendarLinkCopied() {
        var state = on()
        state.copied = true
        snap(calendar(state), "f16.4-link-copied", height: 1000)
    }

    @Test func calendarLinkOffline() {
        var state = on()
        state.linkOffline = true
        snap(calendar(state), "f16.4-link-offline", height: 1000)
    }

    @Test func calendarLinkFailed() {
        snap(calendar(CalendarScreenState(link: .failed, linkMessage: Copy.calendarLinkError)), "f16.4-link-error", height: 1000)
    }

    // MARK: F16.7 the Raids tab rows, and You › Calendar

    private func raidsTab(large: Bool = false) throws -> some View {
        let me = try RaidFixture.load(Me.self, "me")
        let raids = try RaidFixture.load(GuildRaids.self, "guild.raids")
        let character = try #require(me.mainCharacter)
        let content = RaidsTabContent(
            head: GuildHead(character: try #require(character.guild), guild: me.guilds.first),
            character: character, planned: raids.planned ?? [], nights: raids.rows,
            waiting: raids.waitingForApproval, offlineSince: nil
        )
        let controller = try makeController(content.planned.first)
        return NavigationStack {
            RaidsTabContentView(
                state: .loaded(content), clock: RaidTime.clock, now: { RaidTime.now },
                nextController: controller,
                openRaid: { _, _ in }, openNight: { _ in }, openWeb: { _ in }, retry: {},
                guildRows: GuildTabRows(onRoster: 38, openRoster: {}, openHistory: {})
            )
            .navigationTitle(Copy.tabRaids)
            .toolbarTitleDisplayMode(.inline)
        }
    }

    @Test func raidsTabRows() throws { snap(try raidsTab(), "f16.7-raids-tab", height: 1400) }
    @Test func raidsTabRows_xl() throws { snap(try raidsTab(), "f16.7-raids-tab-xl", height: 3000, large: true) }

    @Test func youCalendarRow() throws {
        let summary = try AccountSnapshotTests.summary("me.json")
        snap(NavigationStack { YouView(summary: summary, showsCalendar: true) {} }, "you-calendar-row")
    }

    @Test func youCalendarRow_xl() throws {
        let summary = try AccountSnapshotTests.summary("me.json")
        snap(NavigationStack { YouView(summary: summary, showsCalendar: true) {} }, "you-calendar-row-xl", height: 2400, large: true)
    }
}
