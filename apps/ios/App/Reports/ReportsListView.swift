import DesignSystem
import Models
import SwiftUI

/// What the list hands back to its screen.
struct ReportsActions {
    var open: (ReportRow) -> Void = { _ in }
    var retry: () -> Void = {}
    var loadMore: () -> Void = {}
    var upload: () -> Void = {}
    var sample: () -> Void = {}
    var signIn: () -> Void = {}
}

/// Your reports (mocks ios-reports, -empty, -failed; web /reports): the title, the lead and the
/// count, Upload a log, then the one Plate of reports, newest first. Stateless, so snapshots drive
/// every state.
struct ReportsListView: View {
    let phase: ReportsPhase
    var pageFailed = false
    var isSigningIn = false
    let now: Date
    let clock: ReportsClock
    var actions = ReportsActions()

    /// `.pc` over the plate: 20 under the head.
    static let plateGap: CGFloat = FSSpace.space5

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            if case let .loaded(content) = phase {
                if let since = content.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
                }
                if let since = content.staleSince {
                    StaleStrip(time: OfflineStamp.format(since, now: now), retry: actions.retry)
                }
            }
            VStack(alignment: .leading, spacing: Self.plateGap) {
                header
                Plate { plate }
            }
            .padding(.horizontal, FSSpace.gutterPhone)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, FSSpace.space6)
        .parityContent()
    }

    private var count: Int? {
        if case let .loaded(content) = phase, !content.rows.isEmpty { return content.total }
        return nil
    }

    private var header: some View {
        PageHeader(title: Copy.reportsTitle) {
            HStack(spacing: FSSpace.space4) {
                SmallText(Copy.reportsLead)
                if let count { SmallText(Copy.reportsCount(n: count)) }
            }
        } actions: {
            if phase != .signedOut {
                FSButton(Copy.reportsUpload, fullWidth: false, action: actions.upload)
            }
        }
    }

    @ViewBuilder
    private var plate: some View {
        switch phase {
        case .signedOut:
            EmptyState(Copy.reportsSignedOut) {
                FSButton(Copy.signinButton, kind: .primary, fullWidth: false, action: actions.signIn).disabled(isSigningIn)
            }
        case .loading:
            ReportsSkeleton()
        case .failed:
            LoadError(Copy.reportsFailed, retryTitle: Copy.commonTryAgain, retry: actions.retry)
        case let .loaded(content) where content.rows.isEmpty:
            EmptyState(Copy.reportsEmpty) {
                FSButton(Copy.reportsUpload, kind: .primary, fullWidth: false, action: actions.upload)
                ActionLink(Copy.reportsSample, action: actions.sample)
            }
        case let .loaded(content):
            rows(content)
        }
    }

    private func rows(_ content: ReportsContent) -> some View {
        VStack(spacing: 0) {
            ForEach(content.rows, id: \.id) { row in
                Button { actions.open(row) } label: {
                    ReportRowView(row: row, now: now, clock: clock)
                }
                .buttonStyle(RowPressStyle())
                .onAppear { if row.id == content.rows.last?.id { actions.loadMore() } }
            }
            pageFooter(content)
        }
        .dynamicTypeSize(...ReportsLayout.rowTypeCap)
    }

    @ViewBuilder
    private func pageFooter(_ content: ReportsContent) -> some View {
        if pageFailed {
            LoadError(Copy.reportsFailed, retryTitle: Copy.commonTryAgain, inline: true, retry: actions.loadMore)
        } else if content.hasMorePages && content.offlineSince == nil {
            ProgressView().tint(FSColor.gold).frame(maxWidth: .infinity, minHeight: ReportsLayout.spinnerRow)
        }
    }
}

/// One report row (`ui/Row`, no lead): the name, the date with bosses and kills, then the state
/// tag and the visibility glyph under them.
struct ReportRowView: View {
    let row: ReportRow
    let now: Date
    let clock: ReportsClock
    @ScaledMetric(relativeTo: .footnote) private var glyphSize = ReportsLayout.glyph

    var body: some View {
        FSRow(
            title: ReportsList.name(row), sub: ReportsCopy.rowMeta(row, now: now, clock: clock), trailPlacement: .below,
            lead: { EmptyView() },
            trail: {
                HStack(spacing: FSSpace.space2) {
                    tag
                    glyph
                }
            }
        )
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(ReportsCopy.rowA11y(row, now: now, clock: clock))
        .accessibilityAddTraits(.isButton)
    }

    @ViewBuilder
    private var tag: some View {
        switch ReportRowState(status: row.status) {
        case .live: LiveTag()
        case .processing: FSTag(Copy.reportsStateProcessing, tone: .gold)
        case .failed: FSTag(Copy.reportsStateFailed, tone: .wipe)
        case .complete: EmptyView()
        }
    }

    private var glyph: some View {
        let visibility = ReportVisibility(row.visibility)
        return Image(systemName: visibility.systemImage)
            .font(.system(size: min(glyphSize, ReportsLayout.glyphCap)))
            .foregroundStyle(FSColor.muted)
            .accessibilityLabel(ReportsCopy.visibility(visibility))
    }
}

/// The row chevron, sized with the text so it stays readable at every size.
struct ReportsChevron: View {
    var body: some View {
        Image(systemName: "chevron.right").font(ScaledType.meta.weight(.semibold))
            .foregroundStyle(FSColor.muted).accessibilityHidden(true)
    }
}

/// `Live`: its drawn 7 pt kill-green dot beside the kill Tag.
struct LiveTag: View {
    var body: some View {
        HStack(spacing: ReportsLayout.liveDotGap) {
            Circle().fill(FSColor.kill).frame(width: ReportsLayout.liveDot, height: ReportsLayout.liveDot)
                .accessibilityHidden(true)
            FSTag(Copy.reportsStateLive, tone: .kill)
        }
    }
}

/// An error in the failing region only: `ui/LoadError` in the Plate.
struct ReportsErrorCard: View {
    let text: String
    let retry: () -> Void

    var body: some View {
        Plate { LoadError(text, retryTitle: Copy.commonTryAgain, retry: retry) }
    }
}

/// F10.7: four rows at their real height, inside the Plate.
struct ReportsSkeleton: View {
    static let rows = 4

    var body: some View {
        VStack(spacing: 0) {
            ForEach(0..<Self.rows, id: \.self) { _ in
                VStack(alignment: .leading, spacing: FSSpace.space2) { SkeletonBar(width: 160, height: 15); SkeletonBar(width: 200, height: 12) }
                    .padding(.horizontal, FSSpace.rowPadX)
                    .frame(maxWidth: .infinity, minHeight: ReportsLayout.reportRow, alignment: .leading)
                    .overlay(alignment: .bottom) { Hairline() }
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.reportsTitle)
    }
}

struct SkeletonBar: View {
    let width: CGFloat?
    let height: CGFloat

    var body: some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}

enum ReportsLayout {
    static let reportRow: CGFloat = 68
    static let parseRow: CGFloat = 72
    static let fightRow: CGFloat = 60
    static let meterRow: CGFloat = 44
    static let crest: CGFloat = 36
    static let meterIcon: CGFloat = 28
    static let glyph: CGFloat = 16
    static let liveDot: CGFloat = 7
    static let liveDotGap: CGFloat = 6
    /// Crest and glyph grow with the text up to these (spec §9 largest text).
    static let crestCap: CGFloat = 44
    static let glyphCap: CGFloat = 24
    static let meterIconCap: CGFloat = 40
    /// One `body` line's height at the default size: chevrons and dots centre on it.
    static let firstLine: CGFloat = 20
    /// Row text stops growing here, so the large title stays the largest text and nothing fixed
    /// (a tag) outgrows the row (spec §9, §11).
    static let rowTypeCap = DynamicTypeSize.accessibility3
    static let outcomeDot: CGFloat = 10
    static let spinnerRow: CGFloat = 56
    static let mineEdge: CGFloat = 3
    /// The meter bar's class colour strength (spec §4 Meter row).
    static let barOpacity = 0.22
    static let figure = Font.system(.title3, design: .monospaced, weight: .semibold)
    static let parseNumber = Font.system(.subheadline, design: .monospaced, weight: .bold)
}
