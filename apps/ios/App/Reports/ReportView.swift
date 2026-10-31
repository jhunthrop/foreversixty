import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// What the report hands back to its screen.
struct ReportActions {
    var openFight: (FightRow) -> Void = { _ in }
    var retry: () -> Void = {}
    var retryParses: () -> Void = {}
}

/// F10.2 and F10.5 as drawn from the store's states: stateless, so snapshots drive every state.
struct ReportView: View {
    let phase: ReportPhase
    let parses: ParsesState
    let coach: CoachState
    let now: Date
    let clock: ReportsClock
    var actions = ReportActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            switch phase {
            case .loading:
                ReportSkeleton()
            case .failed:
                ReportsErrorCard(text: Copy.reportFailed, retry: actions.retry).padding(16)
            case .notYours:
                ReportStateCard(title: Copy.reportNotYoursTitle, detail: Copy.reportNotYoursBody, color: FSColor.strong).padding(16)
            case let .loaded(content):
                loaded(content)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.bottom, SimsListView.tabBarClearance)
    }

    @ViewBuilder
    private func loaded(_ content: ReportContent) -> some View {
        let report = content.report
        if let since = content.offlineSince {
            OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now)))
        }
        SimsTitle(text: report.zone ?? report.title)
        header(report)
        VStack(alignment: .leading, spacing: 22) {
            switch report.status {
            case ReportStatus.processing:
                ReportStateCard(title: Copy.reportProcessingTitle, detail: Copy.reportProcessingBody, color: FSColor.strong)
            case ReportStatus.failed:
                ReportStateCard(title: Copy.reportUnreadableTitle, detail: Copy.reportUnreadableBody, color: FSColor.wipe)
            default:
                if report.status == ReportStatus.live {
                    ReportStateCard(title: Copy.reportLiveTitle(n: secondsSinceUpdate(report)), detail: Copy.reportLiveBody, color: FSColor.kill)
                }
                if parses != .hidden {
                    ParsesCard(state: parses, retry: actions.retryParses)
                }
                FightsCard(report: report, rows: content.fightRows, open: actions.openFight)
                if coach != .hidden {
                    VStack(alignment: .leading, spacing: 8) {
                        MicroLabel(Copy.reportCoach)
                        CoachCard(state: coach, clock: clock)
                    }
                }
                Text(Copy.reportFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .dynamicTypeSize(...ReportsLayout.rowTypeCap)
        .padding(.horizontal, 16)
        .padding(.top, 16)
    }

    private func header(_ report: Report) -> some View {
        VStack(alignment: .leading, spacing: 6) {
            HStack(spacing: 8) {
                MicroLabel(Copy.reportTag, color: FSColor.gold)
                TagView(text: ReportsCopy.visibility(ReportVisibility(report.visibility)), color: FSColor.nav)
            }
            Text(ReportsCopy.reportMeta(report, clock: clock)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
        .dynamicTypeSize(...ReportsLayout.rowTypeCap)
        .padding(.horizontal, 16)
    }

    private func secondsSinceUpdate(_ report: Report) -> Int {
        let last = report.heartbeatAt ?? report.lastUploadAt ?? report.createdAt
        return max(0, Int(now.timeIntervalSince(last)))
    }
}

/// A report state's two lines in a quiet card (spec F10.5).
struct ReportStateCard: View {
    let title: String
    let detail: String
    let color: Color

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(title).font(ScaledType.lead).foregroundStyle(color)
                .fixedSize(horizontal: false, vertical: true)
            Text(detail).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
        .accessibilityElement(children: .combine)
    }
}

/// Your parses, the one primary card: a row per character kill, else the no-kills line.
struct ParsesCard: View {
    let state: ParsesState
    let retry: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            MicroLabel(Copy.reportYourParses, color: FSColor.gold)
                .padding(.horizontal, 16).padding(.top, 14).padding(.bottom, 4)
            switch state {
            case .hidden:
                EmptyView()
            case .loading:
                HStack(spacing: 12) {
                    Circle().fill(FSColor.cardTop).frame(width: ReportsLayout.crest, height: ReportsLayout.crest)
                    VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 120, height: 15); SkeletonBar(width: 180, height: 12) }
                    Spacer()
                }
                .padding(.horizontal, 16)
                .frame(height: ReportsLayout.parseRow)
                .accessibilityHidden(true)
            case .failed:
                VStack(alignment: .leading, spacing: 12) {
                    Text(Copy.reportParsesFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                    SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry)
                }
                .padding(16)
            case let .loaded(me):
                let rows = ParseRows.rows(me)
                if rows.isEmpty {
                    Text(Copy.reportNoKills).font(ScaledType.body).foregroundStyle(FSColor.text)
                        .fixedSize(horizontal: false, vertical: true)
                        .padding(16)
                } else {
                    ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                        if index > 0 { Divider().overlay(FSColor.lineWarm) }
                        ParseRowView(row: row)
                    }
                }
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
    }
}

/// One kill: crest, boss and meta left; the figure and the parse number right (stacked at AX sizes).
struct ParseRowView: View {
    let row: ParseRow
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            ClassCrest(className: row.character.className, size: ReportsLayout.crest)
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: 6) { words; figure(alignment: .leading) }
                    .frame(maxWidth: .infinity, alignment: .leading)
            } else {
                words.frame(maxWidth: .infinity, alignment: .leading)
                figure(alignment: .trailing)
            }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 14)
        .frame(minHeight: ReportsLayout.parseRow)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(ReportsCopy.parseA11y(row))
    }

    private var words: some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(row.fight.encounter).font(ScaledType.name).foregroundStyle(FSColor.strong)
            Text(ReportsCopy.parseMeta(row)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
            if row.fight.newBest { TagView(text: Copy.reportNewBest, color: FSColor.gold) }
        }
    }

    private func figure(alignment: HorizontalAlignment) -> some View {
        VStack(alignment: alignment, spacing: 2) {
            HStack(alignment: .firstTextBaseline, spacing: 6) {
                Text(ReportsCopy.figure(row.fight.value)).font(ReportsLayout.figure).foregroundStyle(FSColor.strong)
                Text(ReportsCopy.unit(row.fight.unit)).font(ScaledType.micro).foregroundStyle(FSColor.muted)
            }
            if let percentile = row.fight.percentile {
                HStack(alignment: .firstTextBaseline, spacing: 4) {
                    Text("\(percentile)").font(ReportsLayout.parseNumber).foregroundStyle(ParseLadder.color(ParseTier(percentile)))
                    Text(Copy.reportPercentile).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
            }
        }
        .fixedSize()
    }
}

/// The parse ladder's text colours (spec §4 Parse number, the rankings page's floors).
enum ParseLadder {
    static func color(_ tier: ParseTier) -> Color {
        switch tier {
        case .grey: FSColor.parseGreyText
        case .green: FSColor.parseGreen
        case .blue: Color(tokenHex: 0x3D94F0)
        case .purple: Color(tokenHex: 0xB866F5)
        case .orange: FSColor.parseOrange
        case .pink: Color(tokenHex: 0xF08CC0)
        }
    }
}

/// Fights: every boss pull in order (a boss pulled more than once numbered), each zone's trash as
/// one entry after its boss pulls. At the accessibility sizes the label sits over the count.
struct FightsCard: View {
    let report: Report
    let rows: [FightRow]
    let open: (FightRow) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let pulls = FightRows.pullNumbers(report)
        VStack(alignment: .leading, spacing: 8) {
            header
            VStack(spacing: 0) {
                ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                    if index > 0 { Divider().overlay(FSColor.lineSoft) }
                    if case .inProgress = row {
                        FightRowView(row: row, pulls: pulls, chevron: false)
                    } else {
                        Button { open(row) } label: { FightRowView(row: row, pulls: pulls, chevron: true) }
                            .buttonStyle(RowPressStyle())
                            .accessibilityHint(Copy.reportA11yShowsMeters)
                    }
                }
            }
            .quietCard()
        }
    }

    @ViewBuilder
    private var header: some View {
        let count = MicroLabel(ReportsCopy.bossPulls(FightRows.bossPulls(report))).fixedSize()
        if typeSize.isAccessibilitySize {
            VStack(alignment: .leading, spacing: 4) { MicroLabel(Copy.reportFights); count }
        } else {
            HStack(alignment: .firstTextBaseline) { MicroLabel(Copy.reportFights); Spacer(); count }
        }
    }
}

/// One Fights row: the outcome dot and the chevron centre on the first text line.
struct FightRowView: View {
    let row: FightRow
    let pulls: [Int: PullNumber]
    let chevron: Bool
    @ScaledMetric(relativeTo: .subheadline) private var firstLine = ReportsLayout.firstLine

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Circle().fill(dot).frame(width: ReportsLayout.outcomeDot, height: ReportsLayout.outcomeDot)
                .frame(height: firstLine)
            VStack(alignment: .leading, spacing: 2) {
                Text(ReportsCopy.title(row, pulls: pulls)).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Text(ReportsCopy.meta(row)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if chevron { ReportsChevron().frame(height: firstLine) }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 10)
        .frame(minHeight: ReportsLayout.fightRow)
        .contentShape(Rectangle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(ReportsCopy.fightA11y(row, pulls: pulls))
        .accessibilityAddTraits(chevron ? .isButton : [])
    }

    private var dot: Color {
        switch row {
        case let .boss(fight): fight.kill ? FSColor.kill : FSColor.wipe
        case .inProgress: FSColor.gold
        case .trash, .loneTrash: FSColor.muted
        }
    }
}

/// The report while it first loads: the title, the parses card and the fights at their heights.
struct ReportSkeleton: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            SkeletonBar(width: 220, height: 30)
            SkeletonBar(width: 260, height: 12)
            RoundedRectangle(cornerRadius: FSRadius.panel).fill(FSColor.cardTop).frame(height: ReportsLayout.parseRow + 40)
            RoundedRectangle(cornerRadius: FSRadius.panel).fill(FSColor.cardTop).frame(height: ReportsLayout.fightRow * 4)
        }
        .padding(16)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.reportsBack)
    }
}
