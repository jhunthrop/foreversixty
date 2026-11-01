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
        GuildScreenScroll(title: Copy.historyTitle, offlineStrip: offlineStrip) {
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

/// A history row (72 min): the date block, the zone, two meta lines and the `You were there`
/// tag under them; rows top-aligned. Default size: each meta line on one line. Largest text: the
/// date block above the text, and a meta line wraps only between its segments.
struct HistoryRow: View {
    let night: HistoryNight
    let clock: RaidClock
    let open: () -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let large = typeSize.isAccessibilitySize
        Button(action: open) {
            HStack(alignment: .top, spacing: 12) {
                (large
                    ? AnyLayout(VStackLayout(alignment: .leading, spacing: 10))
                    : AnyLayout(HStackLayout(alignment: .top, spacing: 12))) {
                    DateBlock(date: night.start, clock: clock, width: large ? 64 : 48, height: 52)
                    text(large: large)
                }
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

    private func text(large: Bool) -> some View {
        VStack(alignment: .leading, spacing: 3) {
            Text(night.zone).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            meta(night.factsLine, large: large)
            meta(night.sizeLine, large: large)
            if night.youWereThere {
                if large {
                    WrappingTag(text: Copy.historyThere, color: FSColor.kill).padding(.top, 4)
                } else {
                    TagView(text: Copy.historyThere, color: FSColor.kill).padding(.top, 4)
                }
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    /// The body face with tabular figures, so `1 h 48 m` reads tight.
    @ViewBuilder
    private func meta(_ line: String, large: Bool) -> some View {
        let text = Text(line).font(ScaledType.meta.monospacedDigit()).foregroundStyle(FSColor.muted)
        if large {
            text.fixedSize(horizontal: false, vertical: true)
        } else {
            text.lineLimit(1).minimumScaleFactor(0.8)
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
        .guildPushedBar(Copy.historyTitle)
        .task { await store.load() }
    }
}
