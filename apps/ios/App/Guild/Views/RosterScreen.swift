import DesignSystem
import Models
import SwiftUI

/// The roster's filter state: the chip, the sort and the search.
struct RosterQuery: Equatable {
    var role: RosterRoleFilter = .all
    var sort: RosterSort = .rank
    var search = ""
    var searching = false
    /// Only these characters (the Raid tab's not-synced raiders, row 5); nil lists everyone.
    var only: Set<String>?
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
        let filtered = query.only.map(list.count(only:))
        RosterCountsCard(counts: list.counts, shown: filtered)
        RosterChips(query: $query, onlyCount: filtered)
        let rows = list.visible(role: query.role, query: query.search, sort: query.sort, only: query.only)
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

/// The counts card: three figures at stat size over their captions; filtered, the roster's reads
/// `4 of 38` (round 5, row 2).
struct RosterCountsCard: View {
    let counts: GuildMemberCounts
    /// How many the not-synced filter shows; nil when the roster is not filtered.
    var shown: Int?
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 12))
            : AnyLayout(HStackLayout(alignment: .top, spacing: 12))
        layout {
            ForEach(Self.figures(counts, shown: shown), id: \.caption) { figure($0.value, $0.caption) }
        }
        .raidCardBox(16).quietCard()
    }

    /// The plate's figures: filtered, the one population the list shows (`4 of 38 on the roster`,
    /// round 6, row 2); else the three guild-wide counts.
    static func figures(_ counts: GuildMemberCounts, shown: Int?) -> [(value: String, caption: String)] {
        if let shown { return [(Copy.rosterCountsOf(shown: shown, total: counts.onRoster), Copy.rosterCountsOnRoster)] }
        return [
            ("\(counts.onRoster)", Copy.rosterCountsOnRoster),
            ("\(counts.inGameNoAccount)", Copy.rosterCountsNoAccount),
            ("\(counts.waiting)", Copy.rosterCountsWaiting),
        ]
    }

    private func figure(_ value: String, _ caption: String) -> some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(value).font(ScaledType.stat).foregroundStyle(FSColor.strong)
            Text(caption).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .frame(maxWidth: typeSize.isAccessibilitySize ? nil : .infinity, alignment: .leading)
        .accessibilityElement(children: .combine)
    }
}

/// All, Tanks, Healers, DPS, then `Sort: {by}`; from the Raid tab's not-synced line, the filter's own
/// chip leads, selected, with its clear control, and All is not selected (round 5, row 2).
struct RosterChips: View {
    @Binding var query: RosterQuery
    /// The not-synced filter's count; nil when the roster is not filtered.
    var onlyCount: Int?

    var body: some View {
        WrapLayout(spacing: 8) {
            if let onlyCount { RosterFilterChip(count: onlyCount) { query.only = nil } }
            ForEach(RosterRoleFilter.allCases) { role in
                let selected = query.role == role && (role != .all || onlyCount == nil)
                Button {
                    query.role = role
                    // All is everyone: it clears the not-synced filter too.
                    if role == .all { query.only = nil }
                } label: {
                    Text(role.label).font(ScaledType.meta.weight(.semibold))
                        .foregroundStyle(selected ? FSColor.gold : FSColor.text)
                        .padding(.horizontal, 14)
                        .frame(minHeight: 36)
                        .chipSurface(fill: selected ? GameColors.pressedWash : GameColors.controlFill, line: selected ? FSColor.gold : FSColor.line)
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
                .chipSurface(fill: GameColors.controlFill, line: FSColor.line)
                .padding(.vertical, 4)
                .contentShape(.rect)
            }
        }
    }
}

/// The not-synced filter's chip in the selected style: `Not synced this week · 4`, then its clear
/// control, a 44 pt target.
struct RosterFilterChip: View {
    let count: Int
    let clear: () -> Void

    var body: some View {
        HStack(spacing: 0) {
            Text(Copy.rosterChipNotSynced(n: count)).font(ScaledType.meta.weight(.semibold))
                .padding(.leading, FSSpace.space3)
            Button(action: clear) {
                Image(systemName: "xmark").font(ScaledType.meta.weight(.bold))
                    .frame(width: FSTarget.min, height: FSButtonStyle.smallHeight)
                    .contentShape(.rect.inset(by: (FSButtonStyle.smallHeight - FSTarget.min) / 2))
            }
            .buttonStyle(.plain)
            .accessibilityLabel(Copy.rosterChipClear)
        }
        .foregroundStyle(FSColor.gold)
        .frame(minHeight: FSButtonStyle.smallHeight)
        .guildPill(fill: GameColors.pressedWash, line: FSColor.gold)
        .padding(.vertical, (FSTarget.min - FSButtonStyle.smallHeight) / 2)
        .accessibilityElement(children: .contain)
        .accessibilityAddTraits(.isSelected)
    }
}

/// A roster row (64 min): the ringed class crest, name and meta, attendance, chevron.
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
        ClassCrest(className: entry.className, size: 32, flat: true).accessibilityHidden(true)
    }

    private var nameLine: some View {
        HStack(spacing: 4) {
            Text(entry.name).font(ScaledType.name).foregroundStyle(ClassColor.of(entry.className))
                .lineLimit(1).minimumScaleFactor(0.6)
            if entry.supporter {
                Image(systemName: "crown.fill").font(.system(size: 10, weight: .bold)).foregroundStyle(FSColor.gold)
                    .accessibilityLabel(Copy.sharedPremiumSupporterChip)
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
    @State private var query: RosterQuery
    @Environment(\.openURL) private var openURL

    init(store: RosterStore, clock: RaidClock, now: @escaping () -> Date, only: Set<String>? = nil) {
        self.store = store
        self.clock = clock
        self.now = now
        _query = State(initialValue: RosterQuery(only: only))
    }

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
