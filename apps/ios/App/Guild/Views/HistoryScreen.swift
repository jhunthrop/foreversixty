import DesignSystem
import Models
import SwiftUI

/// The history list's paging state, beside the store's phase.
struct HistoryPagingState: Equatable {
    var loadingMore = false
    var moreFailed = false
}

/// F16.3, pure over the store's phase so snapshot tests draw every state.
struct HistoryContentView: View {
    let phase: GuildLoadPhase<HistoryContent>
    let paging: HistoryPagingState
    let yourKeys: Set<String>
    let clock: RaidClock
    let now: () -> Date
    let open: (HistoryTarget) -> Void
    let loadMore: () -> Void
    let retry: () -> Void

    var body: some View {
        GuildScreenScroll(offlineStrip: offlineStrip) {
            switch phase {
            case .loading:
                SectionLabel(text: Copy.historyLabel)
                GuildSkeletonRows(count: 6, height: 72)
            case .failed:
                ErrorRegion(message: Copy.historyError, retry: retry)
            case let .loaded(content, _):
                loaded(content)
            }
        }
    }

    private var offlineStrip: String? {
        guard case let .loaded(_, since?) = phase else { return nil }
        return Copy.offlineStrip(time: clock.stripTime(since, now: now()))
    }

    @ViewBuilder
    private func loaded(_ content: HistoryContent) -> some View {
        if content.rows.isEmpty {
            GuildMessageCard(title: Copy.historyEmptyTitle, note: Copy.historyEmptyBody)
        } else {
            let nights = content.rows.map { HistoryNight(row: $0, yourKeys: yourKeys) }
            SectionLabel(text: Copy.historyLabel)
            LazyVStack(spacing: 0) {
                ForEach(nights) { night in
                    HistoryRow(night: night, clock: clock) { open(night.target) }
                    if night.id != nights.last?.id { Divider().overlay(FSColor.lineSoft) }
                }
            }
            .quietCard()
            if content.hasMore { moreRow }
            GuildFooter(text: Copy.historyFooter)
        }
    }

    @ViewBuilder
    private var moreRow: some View {
        VStack(alignment: .leading, spacing: 8) {
            Button(action: loadMore) {
                HStack(spacing: 8) {
                    Text(Copy.historyMore)
                    if paging.loadingMore { ProgressView().controlSize(.small).tint(FSColor.gold) }
                }
            }
            .buttonStyle(GoldLinkStyle())
            .disabled(paging.loadingMore)
            if paging.moreFailed {
                Text(Copy.historyMoreError).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                SecondaryButton(title: Copy.commonTryAgain, action: loadMore)
            }
        }
        .padding(.horizontal, 10)
    }
}

/// A history row (72 min): the date block, the zone, two meta lines that never wrap, and the
/// `You were there` tag under them; rows top-aligned.
struct HistoryRow: View {
    let night: HistoryNight
    let clock: RaidClock
    let open: () -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let large = typeSize.isAccessibilitySize
        Button(action: open) {
            HStack(alignment: .top, spacing: 12) {
                DateBlock(date: night.start, clock: clock, width: large ? 56 : 48, height: 52)
                VStack(alignment: .leading, spacing: 3) {
                    Text(night.zone).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                        .fixedSize(horizontal: false, vertical: true)
                    metaLines(large: large)
                    if night.youWereThere {
                        TagView(text: Copy.historyThere, color: FSColor.kill).padding(.top, 4)
                    }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                Image(systemName: "chevron.right").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .padding(.top, 4)
                    .accessibilityHidden(true)
            }
            .padding(14)
            .frame(minHeight: 72)
            .contentShape(.rect)
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(night.accessibilityLabel(clock: clock))
        .accessibilityAddTraits(.isButton)
    }

    /// Default: two single lines. Largest text: one clause a line, so a clause never breaks.
    @ViewBuilder
    private func metaLines(large: Bool) -> some View {
        if large {
            ForEach(Array((night.factsLine + " · " + night.sizeLine).components(separatedBy: " · ").enumerated()), id: \.offset) { item in
                Text(item.element).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted).lineLimit(1).minimumScaleFactor(0.6)
            }
        } else {
            Text(night.factsLine).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted).lineLimit(1).minimumScaleFactor(0.7)
            Text(night.sizeLine).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted).lineLimit(1).minimumScaleFactor(0.7)
        }
    }
}

/// F16.3 pushed from the Raids tab's Raid history row.
struct HistoryScreen: View {
    let store: HistoryStore
    let yourKeys: Set<String>
    let clock: RaidClock
    let now: () -> Date
    let openNight: (Int64) -> Void
    @Environment(\.openURL) private var openURL

    var body: some View {
        HistoryContentView(
            phase: store.phase, paging: HistoryPagingState(loadingMore: store.loadingMore, moreFailed: store.moreFailed),
            yourKeys: yourKeys, clock: clock, now: now,
            open: { target in
                switch target {
                case let .night(raid): openNight(raid)
                case let .report(url): openURL(url)
                }
            },
            loadMore: { Task { await store.loadMore() } },
            retry: { Task { await store.load() } }
        )
        .refreshable { await store.load() }
        .navigationTitle(Copy.historyTitle)
        .navigationBarTitleDisplayMode(.large)
        .task { await store.load() }
    }
}
