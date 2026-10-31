import DesignSystem
import Models
import SwiftUI

/// What the list hands back to its screen.
struct ReportsActions {
    var open: (ReportRow) -> Void = { _ in }
    var retry: () -> Void = {}
    var loadMore: () -> Void = {}
    var getCompanion: () -> Void = {}
    var signIn: () -> Void = {}
}

/// F10.1 as drawn from a phase: stateless, so snapshots drive every state.
struct ReportsListView: View {
    let phase: ReportsPhase
    let account: ReportsAccount
    var pageFailed = false
    var isSigningIn = false
    let now: Date
    let clock: ReportsClock
    var actions = ReportsActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            switch phase {
            case .signedOut:
                SimsTitle(text: Copy.reportsTitle)
                signedOut
            case .loading:
                SimsTitle(text: Copy.reportsTitle)
                ReportsSkeleton()
            case .failed:
                SimsTitle(text: Copy.reportsTitle)
                ReportsErrorCard(text: Copy.reportsFailed, retry: actions.retry).padding(.horizontal, 16).padding(.top, 12)
            case let .loaded(content):
                if let since = content.offlineSince {
                    OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
                }
                SimsTitle(text: Copy.reportsTitle)
                loaded(content)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, SimsListView.tabBarClearance)
    }

    @ViewBuilder
    private func loaded(_ content: ReportsContent) -> some View {
        if content.rows.isEmpty {
            empty.padding(.horizontal, 16).padding(.top, 12)
        } else {
            let sections = ReportsList.sections(content.rows, now: now)
            VStack(alignment: .leading, spacing: 22) {
                if !sections.thisWeek.isEmpty { section(Copy.reportsThisWeek, sections.thisWeek, content: content) }
                if !sections.earlier.isEmpty { section(Copy.reportsEarlier, sections.earlier, content: content) }
                Text(Copy.reportsFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .padding(.horizontal, 16)
            .padding(.top, 12)
        }
    }

    private func section(_ title: String, _ rows: [ReportRow], content: ReportsContent) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(title)
            VStack(spacing: 0) {
                ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                    if index > 0 { Divider().overlay(FSColor.lineSoft) }
                    Button { actions.open(row) } label: {
                        ReportRowView(row: row, account: account, now: now, clock: clock)
                    }
                    .buttonStyle(RowPressStyle())
                    .onAppear { if row.id == content.rows.last?.id { actions.loadMore() } }
                }
                if rows.last?.id == content.rows.last?.id { pageFooter(content) }
            }
            .quietCard()
        }
    }

    @ViewBuilder
    private func pageFooter(_ content: ReportsContent) -> some View {
        if pageFailed {
            HStack {
                Text(Copy.reportsFailed).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                Spacer()
                SecondaryButton(title: Copy.commonTryAgain, action: actions.loadMore)
            }
            .padding(16)
        } else if content.hasMorePages && content.offlineSince == nil {
            ProgressView().tint(FSColor.gold).frame(maxWidth: .infinity, minHeight: ReportsLayout.spinnerRow)
        }
    }

    private var empty: some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(Copy.reportsEmpty).font(ScaledType.lead).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            Button(Copy.reportsGetCompanion, action: actions.getCompanion)
                .buttonStyle(BrandButtonStyle(.secondary))
                .fixedSize(horizontal: true, vertical: false)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }

    private var signedOut: some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(Copy.reportsSignedOut).font(ScaledType.lead).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            Button(action: actions.signIn) {
                if isSigningIn {
                    ProgressView().tint(FSColor.bg).frame(width: 20, height: 20)
                } else {
                    Text(Copy.signinButton)
                }
            }
            .buttonStyle(BrandButtonStyle(.primary))
            .disabled(isSigningIn)
            .accessibilityLabel(Copy.signinButton)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
        .padding(.horizontal, 16)
        .padding(.top, 12)
    }
}

/// One report row (spec F10.1): crest, name, meta, state tag, visibility glyph, chevron.
struct ReportRowView: View {
    let row: ReportRow
    let account: ReportsAccount
    let now: Date
    let clock: ReportsClock
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(alignment: typeSize.isAccessibilitySize ? .top : .center, spacing: 12) {
            ClassCrest(className: ReportsList.crestClass(row, characters: account.characters, currentClass: account.currentClass), size: ReportsLayout.crest)
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: 6) {
                    words
                    HStack(spacing: 10) { tag; glyph }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
            } else {
                words.frame(maxWidth: .infinity, alignment: .leading)
                tag
                glyph
            }
            SimChevron()
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: ReportsLayout.reportRow)
        .contentShape(Rectangle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(ReportsCopy.rowA11y(row, now: now, clock: clock))
        .accessibilityAddTraits(.isButton)
    }

    private var words: some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(ReportsList.name(row)).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            Text(ReportsCopy.rowMeta(row, now: now, clock: clock)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
    }

    @ViewBuilder
    private var tag: some View {
        switch ReportRowState(status: row.status) {
        case .live: LiveTag()
        case .processing: TagView(text: Copy.reportsStateProcessing, color: FSColor.gold)
        case .failed: TagView(text: Copy.reportsStateFailed, color: FSColor.wipe)
        case .complete: EmptyView()
        }
    }

    private var glyph: some View {
        let visibility = ReportVisibility(row.visibility)
        return Image(systemName: visibility.systemImage)
            .font(.system(size: ReportsLayout.glyph))
            .foregroundStyle(FSColor.muted)
            .accessibilityLabel(ReportsCopy.visibility(visibility))
    }
}

/// `Live` in kill green with its drawn 7 pt dot, on the tag's ground.
struct LiveTag: View {
    var body: some View {
        HStack(spacing: 5) {
            Circle().fill(FSColor.kill).frame(width: ReportsLayout.liveDot, height: ReportsLayout.liveDot)
                .accessibilityHidden(true)
            Text(Copy.reportsStateLive.uppercased()).font(ScaledType.micro).tracking(1)
        }
        .foregroundStyle(FSColor.kill)
        .padding(.horizontal, FSSpacing.space2)
        .frame(minHeight: 22)
        .background(FSColor.kill.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(FSColor.kill.opacity(0.35)))
        .fixedSize()
    }
}

/// An error in the failing region only: the sentence in wipe red and `Try again`.
struct ReportsErrorCard: View {
    let text: String
    let retry: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(text).font(ScaledType.lead).foregroundStyle(FSColor.wipe)
                .fixedSize(horizontal: false, vertical: true)
                .accessibilityAddTraits(.isStaticText)
            SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }
}

/// F10.7: four rows at their real height.
struct ReportsSkeleton: View {
    static let rows = 4

    var body: some View {
        VStack(spacing: 0) {
            ForEach(0..<Self.rows, id: \.self) { _ in
                HStack(spacing: 12) {
                    Circle().fill(FSColor.cardTop).frame(width: ReportsLayout.crest, height: ReportsLayout.crest)
                    VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 160, height: 15); SkeletonBar(width: 200, height: 12) }
                    Spacer()
                }
                .padding(.horizontal, 16)
                .frame(height: ReportsLayout.reportRow)
            }
        }
        .quietCard()
        .padding(.horizontal, 16)
        .padding(.top, 12)
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
    static let outcomeDot: CGFloat = 10
    static let spinnerRow: CGFloat = 56
    static let mineEdge: CGFloat = 3
    /// The meter bar's class colour strength (spec §4 Meter row).
    static let barOpacity = 0.22
    static let figure = Font.system(.title3, design: .monospaced, weight: .semibold)
    static let parseNumber = Font.system(.subheadline, design: .monospaced, weight: .bold)
}
