import DesignSystem
import Models
import SwiftUI

/// What the board hands back to its screen.
struct RankingsActions {
    var open: (RankingRow) -> Void = { _ in }
    var retry: () -> Void = {}
    var loadMore: () -> Void = {}
    var openFilters: () -> Void = {}
    var chooseSpec: () -> Void = {}
    var toggleWatch: (RankingRow) -> Void = { _ in }
    /// Nil while the reports switch is off: the menu has no `Open the report`.
    var openReport: ((RankingRow) -> Void)?
}

/// F11.1 and F11.3 as drawn from the store: stateless, so snapshots drive every state.
struct RankingsView: View {
    let phase: BoardPhase
    let filters: RankingsFilters?
    let encounterName: String
    let account: RankingsAccount
    var watching: Set<String> = []
    var pageFailed = false
    let now: Date
    var actions = RankingsActions()

    var body: some View {
        ScrollViewReader { proxy in
            VStack(alignment: .leading, spacing: 0) {
                if case let .loaded(content) = phase, let since = content.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
                }
                if case let .loaded(content) = phase, let since = content.staleSince {
                    StaleStrip(time: OfflineStamp.format(since, now: now), retry: actions.retry)
                }
                SimsTitle(text: Copy.rankTitle)
                VStack(alignment: .leading, spacing: 12) {
                    switch phase {
                    case .loading:
                        if account.signedIn { AnswerSkeleton() }
                        if let filters { filterLine(filters) }
                        RankingsSkeleton()
                    case .failed:
                        if let filters { filterLine(filters) }
                        ReportsErrorCard(text: Copy.rankError, retry: actions.retry)
                    case let .loaded(content):
                        loaded(content, proxy: proxy)
                    }
                }
                .dynamicTypeSize(...RankingsLayout.rowTypeCap)
                .padding(.horizontal, 16)
                .padding(.top, 8)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.bottom, SimsListView.tabBarClearance)
        }
    }

    @ViewBuilder
    private func loaded(_ content: BoardContent, proxy: ScrollViewProxy) -> some View {
        if let filters {
            if let answer = RankingsCopy.answer(page: content.first, filters: filters, encounterName: encounterName, current: account.current) {
                AnswerCard(answer: answer)
            }
            filterLine(filters)
            // Your row exists only for a signed-in viewer's current character.
            let own = account.current == nil ? OwnRow.none : OwnRow.placement(rows: content.rows, characterRank: content.first.characterRank)
            if content.rows.isEmpty {
                Text(Copy.rankEmpty).font(ScaledType.body).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(16)
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .quietCard()
            } else {
                if case let .inPlace(rank) = own {
                    Button(Copy.rankJump(ordinal: ReportsCopy.ordinal(rank))) {
                        withAnimation { proxy.scrollTo(rank, anchor: .center) }
                    }
                    .font(ScaledType.bodyStrong)
                    .foregroundStyle(FSColor.gold)
                    .multilineTextAlignment(.leading)
                    .frame(minHeight: 44)
                }
                if case .pinned = own, let pinned = content.first.characterRank {
                    pinnedSection(pinned.row, content: content, filters: filters)
                }
                board(content, filters: filters, own: own)
            }
        }
    }

    private func filterLine(_ filters: RankingsFilters) -> some View {
        FilterLine(filters: filters, encounterName: encounterName, signedIn: account.signedIn, actions: actions)
    }

    private func pinnedSection(_ row: RankingRow, content: BoardContent, filters: RankingsFilters) -> some View {
        VStack(alignment: .leading, spacing: 6) {
            MicroLabel(Copy.rankPin)
            rowButton(row, content: content, filters: filters, mine: true, crests: RankingsBoard.showsCrests(content.rows))
                .clipShape(RoundedRectangle(cornerRadius: FSRadius.panel))
                .quietCard()
            FSColor.lineSoft.frame(height: 1).padding(.top, 8)
        }
        .padding(.top, 4)
    }

    private func board(_ content: BoardContent, filters: RankingsFilters, own: OwnRow) -> some View {
        let crests = RankingsBoard.showsCrests(content.rows)
        return VStack(spacing: 0) {
            ForEach(Array(content.rows.enumerated()), id: \.element.rank) { index, row in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                rowButton(row, content: content, filters: filters, mine: own == .inPlace(row.rank), crests: crests)
                    .id(row.rank)
                    .onAppear { if row.rank == content.rows.last?.rank { actions.loadMore() } }
            }
            pageFooter(content)
        }
        .clipShape(RoundedRectangle(cornerRadius: FSRadius.panel))
        .quietCard()
    }

    private func rowButton(_ row: RankingRow, content: BoardContent, filters: RankingsFilters, mine: Bool, crests: Bool) -> some View {
        let yours = RankingsBoard.isYours(row, keys: Set(account.characterKeys))
        return Button { actions.open(row) } label: {
            RankRowView(row: row, total: content.first.total, metric: filters.metric, yours: yours, mine: mine, crest: crests, now: now)
        }
        .buttonStyle(RowPressStyle())
        .contextMenu {
            RankRowMenu(row: row, signedIn: account.signedIn, watched: watching.contains(row.player.key), actions: actions)
        }
    }

    @ViewBuilder
    private func pageFooter(_ content: BoardContent) -> some View {
        if pageFailed {
            HStack {
                Text(Copy.rankError).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                Spacer()
                SecondaryButton(title: Copy.commonTryAgain, action: actions.loadMore)
            }
            .padding(16)
        } else if content.hasMorePages && content.offlineSince == nil {
            ProgressView().tint(FSColor.gold).frame(maxWidth: .infinity, minHeight: ReportsLayout.spinnerRow)
                .accessibilityLabel(Copy.rankLoadingMore)
        }
    }
}

/// A row's long-press menu (spec F11.1): `Watch` or `Stop watching` signed in, `Open the report`
/// while the reports switch is on, and `Share` (the web character page).
struct RankRowMenu: View {
    let row: RankingRow
    let signedIn: Bool
    let watched: Bool
    let actions: RankingsActions

    var body: some View {
        if signedIn {
            Button { actions.toggleWatch(row) } label: {
                Label(watched ? Copy.rankMenuUnwatch : Copy.rankMenuWatch, systemImage: watched ? "star.slash" : "star")
            }
        }
        if let openReport = actions.openReport {
            Button { openReport(row) } label: { Label(Copy.rankMenuReport, systemImage: "doc.text") }
        }
        if let ref = CharacterRef(key: row.player.key, name: row.player.name) {
            ShareButton(item: ShareText.character(key: ref.key, name: row.player.name, spec: nil, className: nil, level: nil)) {
                Label(Copy.rankMenuShare, systemImage: "square.and.arrow.up")
            }
        }
    }
}

/// The filter line (spec F11.1): the words wrap, never cut, beside the 44 pt `Filters` button.
/// Signed out it starts with `Choose a class and spec`, which opens the F3.2 picker.
struct FilterLine: View {
    let filters: RankingsFilters
    let encounterName: String
    let signedIn: Bool
    let actions: RankingsActions
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: 10))
            : AnyLayout(HStackLayout(alignment: .center, spacing: 10))
        layout {
            VStack(alignment: .leading, spacing: 2) {
                if !signedIn {
                    Button(choose, action: actions.chooseSpec)
                        .font(ScaledType.bodyStrong)
                        .foregroundStyle(FSColor.gold)
                        .frame(minHeight: 44)
                }
                Text(RankingsCopy.filterLine(filters, encounterName: encounterName))
                    .font(ScaledType.meta).foregroundStyle(FSColor.nav)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            Button(action: actions.openFilters) {
                HStack(spacing: 6) {
                    Image(systemName: "slider.horizontal.3").accessibilityHidden(true)
                    Text(Copy.rankFilters)
                }
            }
            .buttonStyle(BrandButtonStyle(.secondary, height: 44, fullWidth: false))
        }
    }

    /// `Choose a class and spec`, or the picked cohort once chosen.
    private var choose: String {
        filters.classSlug.isEmpty ? Copy.rankChooseSpec : "\(filters.spec) \(RankingsCopy.className(filters.classSlug))"
    }
}

/// Your rank (spec F11.1): the ordinal at stat size in its parse colour, the unit, the meta.
struct AnswerCard: View {
    let answer: RankAnswer

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            MicroLabel(Copy.rankLabel, color: FSColor.gold)
            switch answer {
            case let .ranked(ordinal, tier, unit, meta, name):
                ViewThatFits(in: .horizontal) {
                    HStack(alignment: .firstTextBaseline, spacing: 8) { figure(ordinal, tier); unitText(unit) }
                    VStack(alignment: .leading, spacing: 2) { figure(ordinal, tier); unitText(unit) }
                }
                metaText(meta, name: name)
            case let .unranked(title, meta):
                Text(title).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Text(meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 14)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
        .accessibilityElement(children: .combine)
    }

    private func figure(_ ordinal: String, _ tier: ParseTier) -> some View {
        Text(ordinal).font(ScaledType.stat).foregroundStyle(ParseLadder.color(tier))
    }

    private func unitText(_ unit: String) -> some View {
        Text(unit).font(ScaledType.body).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
    }

    /// The meta, its leading character name in gold (F11.1, F11.3).
    private func metaText(_ meta: String, name: String) -> some View {
        var text = AttributedString(meta)
        if let range = text.range(of: name) {
            text[range].foregroundColor = FSColor.gold
            text[range].font = ScaledType.bodyStrong
        }
        return Text(text).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            .fixedSize(horizontal: false, vertical: true)
    }
}

/// One board row (spec F11.1, 68 min): the rank numeral in its parse colour, the name in the
/// class colour with `You`, the meta, the value with its unit word, the chevron. At the
/// accessibility sizes the value moves under the name. Your row: the row wash and a gold edge.
struct RankRowView: View {
    let row: RankingRow
    let total: Int
    let metric: RankMetric
    let yours: Bool
    /// Your row, marked in place or pinned.
    let mine: Bool
    let crest: Bool
    let now: Date
    @Environment(\.dynamicTypeSize) private var typeSize
    @ScaledMetric(relativeTo: .subheadline) private var numeral = RankingsLayout.numeralWidth

    var body: some View {
        HStack(alignment: typeSize.isAccessibilitySize ? .top : .center, spacing: 12) {
            if !typeSize.isAccessibilitySize {
                rankNumeral.frame(minWidth: min(numeral, RankingsLayout.numeralCap), alignment: .trailing)
                if crest { ClassCrest(className: row.player.className, size: RankingsLayout.crest) }
            }
            VStack(alignment: .leading, spacing: 2) {
                // At the accessibility sizes the numeral and crest lead the name's line above it.
                if typeSize.isAccessibilitySize {
                    HStack(spacing: 8) {
                        rankNumeral
                        if crest { ClassCrest(className: row.player.className, size: RankingsLayout.crest) }
                    }
                }
                HStack(spacing: 6) {
                    Text(row.player.name).font(ScaledType.bodyStrong.weight(.bold)).foregroundStyle(ClassColor.of(row.player.className))
                        .fixedSize(horizontal: false, vertical: true)
                    if yours { TagView(text: Copy.rankYou, color: FSColor.gold) }
                }
                Text(RankingsCopy.rowMeta(row, now: now)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                if typeSize.isAccessibilitySize { value(alignment: .leading).padding(.top, 4) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if !typeSize.isAccessibilitySize { value(alignment: .trailing) }
            ReportsChevron()
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: RankingsLayout.row)
        .background(mine ? GameColors.rowWash : .clear)
        .overlay(alignment: .leading) {
            if mine { FSColor.gold.frame(width: ReportsLayout.mineEdge) }
        }
        .contentShape(Rectangle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(RankingsCopy.a11y(row, total: total, metric: metric, yours: yours))
        .accessibilityAddTraits(.isButton)
    }

    private var rankNumeral: some View {
        Text("\(row.rank)").font(ScaledType.mono.weight(.bold))
            .foregroundStyle(ParseLadder.color(RankColor.tier(rank: row.rank, total: total)))
            .fixedSize()
    }

    private func value(alignment: HorizontalAlignment) -> some View {
        VStack(alignment: alignment, spacing: 0) {
            Text(RankingsCopy.value(row, metric: metric)).font(ScaledType.mono).foregroundStyle(FSColor.strong)
            Text(RankingsCopy.unit(metric)).font(ScaledType.micro).foregroundStyle(FSColor.muted)
        }
        .fixedSize()
    }
}

/// The answer card while it loads.
struct AnswerSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            SkeletonBar(width: 70, height: 11)
            SkeletonBar(width: 220, height: 28)
            SkeletonBar(width: 240, height: 12)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
        .accessibilityHidden(true)
    }
}

/// The board while it loads: five rows at their real height.
struct RankingsSkeleton: View {
    static let rows = 5

    var body: some View {
        VStack(spacing: 0) {
            ForEach(0..<Self.rows, id: \.self) { _ in
                HStack(spacing: 12) {
                    SkeletonBar(width: 22, height: 15)
                    VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 130, height: 15); SkeletonBar(width: 210, height: 12) }
                    Spacer()
                    SkeletonBar(width: 36, height: 15)
                }
                .padding(.horizontal, 16)
                .frame(height: RankingsLayout.row)
            }
        }
        .quietCard()
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.rankTitle)
    }
}

enum RankingsLayout {
    static let row: CGFloat = 68
    static let crest: CGFloat = 28
    static let numeralWidth: CGFloat = 30
    static let numeralCap: CGFloat = 64
    static let headCrest: CGFloat = 64
    static let bar: CGFloat = 8
    static let star: CGFloat = 44
    static let sparkWidth: CGFloat = 96
    static let sparkHeight: CGFloat = 34
    static let bestRow: CGFloat = 64
    static let fightRow: CGFloat = 56
    static let barLabel: CGFloat = 86
    static let disabledOpacity = 0.5
    /// A text link's hit area beyond its line, to a 44 pt target.
    static let linkSlop: CGFloat = 12
    /// Row text stops growing here, as F10's rows do (spec §9, §11).
    static let rowTypeCap = DynamicTypeSize.accessibility3
}
