import DesignSystem
import Models
import SwiftUI

/// What the list hands back to its screen.
struct SimsActions {
    var open: (SimRow) -> Void = { _ in }
    var runAgain: (String) -> Void = { _ in }
    var retry: () -> Void = {}
    var loadMore: () -> Void = {}
    var openPlans: () -> Void = {}
}

/// F5.1 as drawn from a phase: stateless, so snapshots drive every state.
struct SimsListView: View {
    let phase: SimsPhase
    let entitled: Bool
    let progress: [String: SimProgress]
    let rerun: [String: RerunStatus]
    var pageFailed = false
    var mayLinkToPlans = false
    let specs: SpecCatalog
    let build: String?
    let now: Date
    var actions = SimsActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            BandHead(title: Copy.simsTitle)
            switch phase {
            case .loading: SimsSkeleton()
            case .failed: errorCard(Copy.simsFailedLoad)
            case let .loaded(content): loaded(content)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, SimsListView.tabBarClearance)
    }

    /// Content ends above the floating tab bar (spec F5.1 Footer).
    static let tabBarClearance: CGFloat = 120

    @ViewBuilder
    private func loaded(_ content: SimsContent) -> some View {
        if let since = content.offlineSince {
            OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
        }
        if let since = content.staleSince {
            StaleStrip(time: OfflineStamp.format(since, now: now), retry: actions.retry)
        }
        VStack(alignment: .leading, spacing: 22) {
            if ServerRuns.showsSection(entitled: entitled, serverRows: content.serverRows) || content.serverFailed {
                serverSection(content)
            }
            mineSection(content)
        }
        .padding(.horizontal, 16)
        .padding(.top, BandReach.gap)
    }

    @ViewBuilder
    private func serverSection(_ content: SimsContent) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            SectionTitle(Copy.simsServer)
            if content.serverFailed {
                errorCard(Copy.simsServerFailed)
            } else if content.serverRows.isEmpty {
                Text(Copy.simsServerEmpty).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(16).frame(maxWidth: .infinity, alignment: .leading).quietCard()
            } else {
                let visible = ServerRuns.visible(content.serverRows)
                card(visible.rows) { row in
                    ServerRunRowView(
                        face: SimRowFace(row, specs: specs, now: now), now: now, build: build,
                        progress: row.simId.flatMap { progress[$0] }, rerun: row.simId.flatMap { rerun[$0] } ?? .idle,
                        offline: content.offlineSince != nil, mayLinkToPlans: mayLinkToPlans, actions: actions
                    )
                } footer: {
                    if visible.hasMore {
                        Text(Copy.simsOverflow).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            .padding(.horizontal, 16).padding(.vertical, 12)
                    }
                }
            }
        }
    }

    @ViewBuilder
    private func mineSection(_ content: SimsContent) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            SectionTitle(Copy.simsMine)
            if content.mineFailed {
                errorCard(Copy.simsFailedLoad)
            } else if content.mine.isEmpty {
                Text(Copy.simsEmpty).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(16).frame(maxWidth: .infinity, alignment: .leading).quietCard()
            } else {
                card(content.mine) { row in
                    MySimRowView(face: SimRowFace(row, specs: specs, now: now), now: now, build: build, open: actions.open)
                        .onAppear { if row.simId == content.mine.suffix(5).first?.simId { actions.loadMore() } }
                } footer: {
                    if pageFailed {
                        HStack {
                            Text(Copy.simsFailedLoad).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                            Spacer()
                            SecondaryButton(title: Copy.commonTryAgain, action: actions.loadMore)
                        }
                        .padding(16)
                    } else if content.hasMorePages && content.offlineSince == nil {
                        ProgressView().tint(FSColor.gold).frame(maxWidth: .infinity, minHeight: 56)
                    }
                }
            }
        }
    }

    private func card<Row: View, Footer: View>(
        _ rows: [SimRow], @ViewBuilder row: @escaping (SimRow) -> Row, @ViewBuilder footer: () -> Footer
    ) -> some View {
        VStack(spacing: 0) {
            ForEach(Array(rows.enumerated()), id: \.offset) { index, item in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                row(item)
            }
            footer()
        }
        .quietCard()
    }

    private func errorCard(_ text: String) -> some View {
        LoadError(text, retryLabel: Copy.commonTryAgain, retry: actions.retry)
    }
}

/// F5.6: three rows of each section at their real heights.
struct SimsSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 22) {
            section(height: SimsLayout.serverRowHeight, icon: SimsLayout.serverIcon)
            section(height: SimsLayout.mineRowHeight, icon: SimsLayout.mineIcon)
        }
        .padding(.horizontal, 16)
        .padding(.top, 12)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.simsTitle)
    }

    private func section(height: CGFloat, icon: CGFloat) -> some View {
        VStack(spacing: 0) {
            ForEach(0..<3, id: \.self) { _ in
                HStack(alignment: .top, spacing: 12) {
                    bar(icon, icon)
                    VStack(alignment: .leading, spacing: 8) { bar(180, 15); bar(120, 12) }
                    Spacer()
                    bar(64, 14)
                }
                .padding(.horizontal, 16).padding(.top, 14)
                .frame(height: height, alignment: .top)
            }
        }
        .quietCard()
    }

    private func bar(_ width: CGFloat, _ height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}
