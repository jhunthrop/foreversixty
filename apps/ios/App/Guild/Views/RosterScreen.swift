import DesignSystem
import Models
import SwiftUI

/// The roster's filter state: the chip, the sort and the search.
struct RosterQuery: Equatable {
    var role: RosterRoleFilter = .all
    var sort: RosterSort = .rank
    var search = ""
    var searching = false
}

/// F16.1 and F16.2, pure over the store's phase so snapshot tests draw every state.
struct RosterContentView: View {
    let phase: GuildLoadPhase<GuildHome>
    let clock: RaidClock
    let now: () -> Date
    @Binding var query: RosterQuery
    let open: (URL) -> Void
    let retry: () -> Void
    /// The trailing search button: shows the field, or hides and clears it.
    var toggleSearch: () -> Void = {}
    /// Lane F15: the stale strip's time.
    var staleSince: Date?
    @FocusState private var searchFocused: Bool
    @State private var width: CGFloat = 402

    var body: some View {
        GuildScreenScroll(title: Copy.rosterTitle, offlineStrip: offlineStrip, trailing: { searchButton }, content: {
            switch phase {
            case .loading:
                GuildSkeletonRows(count: 8, height: 64)
            case .failed:
                ErrorRegion(message: Copy.rosterError, retry: retry)
            case let .loaded(home, _):
                loaded(RosterList(home: home))
            }
        }, stale: staleSince.map { StaleStrip(time: clock.stripTime($0, now: now()), retry: retry) })
        .onGeometryChange(for: CGFloat.self) { $0.size.width } action: { width = $0 }
    }

    private var offlineStrip: String? {
        guard case let .loaded(_, since?) = phase else { return nil }
        return Copy.offlineStrip(time: clock.stripTime(since, now: now()))
    }

    @ViewBuilder
    private func loaded(_ list: RosterList) -> some View {
        if query.searching { searchField }
        RosterCountsCard(counts: list.counts)
        RosterChips(query: $query)
        let rows = list.visible(role: query.role, query: query.search, sort: query.sort)
        if rows.isEmpty {
            GuildMessageCard(title: Copy.rosterEmptyTitle, note: Copy.rosterEmptyBody)
        } else {
            LazyVStack(spacing: 0) {
                ForEach(rows) { entry in
                    GuildRosterRowView(entry: entry, showsBar: width >= GuildRosterRowView.barMinWidth, open: open)
                    if entry.id != rows.last?.id { Divider().overlay(FSColor.lineSoft) }
                }
            }
            .quietCard()
            GuildFooter(text: Copy.rosterFooter)
        }
    }

    /// The trailing search button (spec F16.1): a 44 pt circle on the title row.
    private var searchButton: some View {
        Button(action: toggleSearch) {
            Image(systemName: query.searching ? "xmark" : "magnifyingglass")
                .font(.system(size: 16, weight: .semibold))
                .foregroundStyle(FSColor.strong)
                .frame(width: 44, height: 44)
                .guildPill(fill: GameColors.controlFill, line: FSColor.line)
        }
        .buttonStyle(.plain)
        .accessibilityLabel(query.searching ? Copy.commonCancel : Copy.rosterSearch)
    }

    private var searchField: some View {
        HStack(spacing: 8) {
            Image(systemName: "magnifyingglass").foregroundStyle(FSColor.muted).accessibilityHidden(true)
            TextField(Copy.rosterSearch, text: $query.search)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.strong)
                .textInputAutocapitalization(.never)
                .autocorrectionDisabled()
                .focused($searchFocused)
                .submitLabel(.search)
        }
        .padding(.horizontal, 14)
        .frame(minHeight: 44)
        .guildPill(fill: GameColors.controlFill, line: searchFocused ? FSColor.gold : FSColor.line)
        .onAppear { searchFocused = query.search.isEmpty }
    }
}

/// The counts card: three figures at stat size over their captions.
struct RosterCountsCard: View {
    let counts: GuildMemberCounts
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 12))
            : AnyLayout(HStackLayout(alignment: .top, spacing: 12))
        layout {
            figure(counts.onRoster, Copy.rosterCountsOnRoster)
            figure(counts.inGameNoAccount, Copy.rosterCountsNoAccount)
            figure(counts.waiting, Copy.rosterCountsWaiting)
        }
        .raidCardBox(16).quietCard()
    }

    private func figure(_ value: Int, _ caption: String) -> some View {
        VStack(alignment: .leading, spacing: 2) {
            Text("\(value)").font(ScaledType.stat).foregroundStyle(FSColor.strong)
            Text(caption).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .frame(maxWidth: typeSize.isAccessibilitySize ? nil : .infinity, alignment: .leading)
        .accessibilityElement(children: .combine)
    }
}

/// All, Tanks, Healers, DPS, then `Sort: {by}`.
struct RosterChips: View {
    @Binding var query: RosterQuery

    var body: some View {
        WrapLayout(spacing: 8) {
            ForEach(RosterRoleFilter.allCases) { role in
                let selected = query.role == role
                Button { query.role = role } label: {
                    Text(role.label).font(ScaledType.meta.weight(.semibold))
                        .foregroundStyle(selected ? FSColor.gold : FSColor.text)
                        .padding(.horizontal, 14)
                        .frame(minHeight: 36)
                        .guildPill(fill: selected ? GameColors.pressedWash : GameColors.controlFill, line: selected ? FSColor.gold : FSColor.line)
                        .padding(.vertical, 4)
                        .contentShape(.rect)
                }
                .buttonStyle(.plain)
                .accessibilityAddTraits(selected ? .isSelected : [])
            }
            Menu {
                Picker(Copy.rosterSort(by: query.sort.label), selection: $query.sort) {
                    ForEach(RosterSort.allCases) { Text($0.label).tag($0) }
                }
            } label: {
                HStack(spacing: 4) {
                    Text(Copy.rosterSort(by: query.sort.label)).font(ScaledType.meta.weight(.semibold))
                    Image(systemName: "chevron.down").font(.system(size: 10, weight: .bold)).accessibilityHidden(true)
                }
                .foregroundStyle(FSColor.muted)
                .padding(.horizontal, 14)
                .frame(minHeight: 36)
                .guildPill(fill: GameColors.controlFill, line: FSColor.line)
                .padding(.vertical, 4)
                .contentShape(.rect)
            }
        }
    }
}

/// A roster row (64 min): spec icon, name and meta, attendance, chevron.
struct GuildRosterRowView: View {
    let entry: RosterEntry
    /// False on a screen narrower than `barMinWidth`.
    var showsBar = true
    let open: (URL) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize
    /// Below this width the bar hides and the caption stays (spec 9).
    static let barMinWidth: CGFloat = 340

    var body: some View {
        if let url = entry.pageURL {
            Button { open(url) } label: { content(chevron: true) }
                .buttonStyle(RowPressStyle())
        } else {
            content(chevron: false)
        }
    }

    private func content(chevron: Bool) -> some View {
        let large = typeSize.isAccessibilitySize
        return HStack(alignment: large ? .top : .center, spacing: 12) {
            // Largest text: the icon sits on the name line, centred on it.
            if !large { icon }
            VStack(alignment: .leading, spacing: 2) {
                if large {
                    HStack(alignment: .center, spacing: 10) { icon; nameLine }
                } else {
                    nameLine
                }
                Text(entry.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                if large { standing(showsBar: false).padding(.top, 4) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if !large { standing(showsBar: showsBar).fixedSize() }
            if chevron {
                Image(systemName: "chevron.right").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
        }
        .padding(.horizontal, 14)
        .padding(.vertical, 10)
        .frame(minHeight: 64)
        .contentShape(.rect)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(entry.accessibilityLabel)
        .accessibilityAddTraits(chevron ? .isButton : [])
    }

    private var icon: some View {
        SpecOrCrest(icon: entry.specIcon, className: entry.className, size: 32)
    }

    private var nameLine: some View {
        HStack(spacing: 4) {
            Text(entry.name).font(ScaledType.name).foregroundStyle(ClassColor.of(entry.className))
                .lineLimit(1).minimumScaleFactor(0.6)
            if entry.supporter {
                Image(systemName: "crown.fill").font(.system(size: 10, weight: .bold)).foregroundStyle(FSColor.gold)
                    .accessibilityLabel(Copy.rosterSupporter)
            }
        }
    }

    @ViewBuilder
    private func standing(showsBar: Bool) -> some View {
        switch entry.standing {
        case .attendance:
            VStack(alignment: .trailing, spacing: 4) {
                if showsBar, let bar = entry.bar { AttendanceBarView(bar: bar) }
                if let caption = entry.attendanceCaption {
                    Text(caption).font(ScaledType.meta.monospacedDigit()).foregroundStyle(FSColor.muted)
                }
            }
        case .noNights:
            Text(Copy.rosterNoNights).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        case .waiting:
            TagView(text: Copy.rosterWaitingTag, color: FSColor.gold)
        case .noAccount:
            TagView(text: Copy.rosterNoAccountTag, color: FSColor.muted)
        case .unknown:
            EmptyView()
        }
    }
}

/// One 7 x 12 pt segment per night: kill when present, the line colour when absent. Hidden
/// from VoiceOver (the row carries the words).
struct AttendanceBarView: View {
    let bar: AttendanceBar

    var body: some View {
        HStack(spacing: 2) {
            ForEach(0..<bar.segments, id: \.self) { index in
                RoundedRectangle(cornerRadius: 1.5)
                    .fill(index < bar.filled ? FSColor.kill : FSColor.line)
                    .frame(width: 7, height: 12)
            }
        }
        .accessibilityHidden(true)
    }
}

/// F16.1 pushed from the Raids tab's Roster row.
struct RosterScreen: View {
    let store: RosterStore
    let clock: RaidClock
    let now: () -> Date
    @State private var query = RosterQuery()
    @Environment(\.openURL) private var openURL

    var body: some View {
        RosterContentView(
            phase: store.phase, clock: clock, now: now, query: $query,
            open: { openURL($0) }, retry: { Task { await store.load() } },
            toggleSearch: {
                query.searching.toggle()
                if !query.searching { query.search = "" }
            },
            staleSince: store.staleSince
        )
        .refreshable { await store.load() }
        .guildPushedBar(Copy.rosterTitle)
        .task { await store.load() }
    }
}
