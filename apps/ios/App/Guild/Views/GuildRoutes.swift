import SwiftUI

/// The pushed Roster: the tab's store when it is this guild's (the count row already loaded it),
/// else its own, kept for the screen's life.
struct RosterRoute: View {
    let extras: GuildExtrasDependencies
    @State private var store: RosterStore

    init(extras: GuildExtrasDependencies, guild: Int64, shared: RosterStore?) {
        self.extras = extras
        _store = State(initialValue: shared.flatMap { $0.guild == guild ? $0 : nil } ?? extras.rosterStore(guild: guild))
    }

    var body: some View {
        RosterScreen(store: store, clock: extras.clock, now: extras.now)
    }
}

/// The pushed Raid history, its store kept for the screen's life.
struct HistoryRoute: View {
    let extras: GuildExtrasDependencies
    let yourKeys: Set<String>
    let openNight: (Int64) -> Void
    @State private var store: HistoryStore

    init(extras: GuildExtrasDependencies, guild: Int64, yourKeys: Set<String>, openNight: @escaping (Int64) -> Void) {
        self.extras = extras
        self.yourKeys = yourKeys
        self.openNight = openNight
        _store = State(initialValue: extras.historyStore(guild: guild))
    }

    var body: some View {
        HistoryScreen(store: store, yourKeys: yourKeys, clock: extras.clock, now: extras.now, openNight: openNight)
    }
}

/// You › Calendar: the link model for the signed-in account, kept for the screen's life.
struct CalendarRoute: View {
    let sync: CalendarSyncController
    @State private var link: CalendarLinkModel

    init(extras: GuildExtrasDependencies, sync: CalendarSyncController, account: String) {
        self.sync = sync
        _link = State(initialValue: extras.linkModel(account: account))
    }

    var body: some View {
        CalendarScreen(sync: sync, link: link)
    }
}
