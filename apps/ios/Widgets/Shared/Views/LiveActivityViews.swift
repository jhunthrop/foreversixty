import AppIntents
import DesignSystem
import DesignTokens
import Models
import SwiftUI
import WidgetKit

/// The raid Live Activity's words (flows F2, E7; mock ios-live-activity), one place for the lock
/// screen and the Island.
struct RaidLiveFigures: Equatable, Sendable {
    let zone: String
    /// `In, as Thandrel, Damage`
    let answer: String
    /// `Starts 8:00 pm CST · 18 in, 3 late · 2 tanks, 5 heals`
    let starts: String
    /// The Island's compact name: the zone's first three letters (`Ony`).
    let short: String
    let cancelled: Bool
    let done: Bool
    /// `Onyxia's Lair done`
    let doneTitle: String
    /// The summary's parts before the loot (`1 kill`), the loot line, and after it (`report ready`).
    let doneLead: [String]
    let loot: RaidLiveLoot?
    let doneTail: [String]

    init(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState, zone timeZone: TimeZone) {
        let clock = WidgetClock(zone: timeZone)
        zone = attributes.zoneName
        answer = RaidGlance.answerLine(state: state.myState, name: attributes.characterName, role: state.role)
        var parts = [Copy.liveStartsAt(time: clock.timeWithZone(state.startDate))]
        if let inCount = state.inCount {
            parts.append(Copy.liveCounts(n: inCount, late: state.late ?? 0))
        } else {
            parts.append(state.size.map { Copy.widgetSignedUp(n: state.signedUp, size: $0) } ?? Copy.raidCountNoSize(n: state.signedUp))
        }
        if let tanks = state.tanks, let healers = state.healers {
            parts.append("\(Copy.liveTanks(n: tanks)), \(Copy.liveHeals(n: healers))")
        }
        starts = parts.joined(separator: " · ")
        short = String((attributes.zoneName.split(separator: " ").first ?? Substring(attributes.zoneName)).prefix(3))
        cancelled = state.status == "cancelled"
        done = state.status == RaidActivityAttributes.ContentState.done
        doneTitle = Copy.liveDone(zone: attributes.zoneName)
        doneLead = state.kills.map { [Copy.liveKills(n: $0)] } ?? []
        loot = state.loot
        doneTail = state.reportReady == true ? [Copy.liveReportReady] : []
    }

    /// The summary as one line, the loot in its quality colour.
    var doneLine: AttributedString {
        let separator = AttributedString(" · ")
        var pieces = doneLead.map(AttributedString.init)
        if let loot {
            let whole = Copy.liveLoot(item: loot.name)
            var line = AttributedString(whole)
            if let range = line.range(of: loot.name) { line[range].foregroundColor = ItemQuality(value: loot.quality).text }
            pieces.append(line)
        }
        pieces += doneTail.map(AttributedString.init)
        return pieces.enumerated().reduce(into: AttributedString()) { out, piece in
            if piece.offset > 0 { out += separator }
            out += piece.element
        }
    }

    var spoken: String {
        if done { return "\(doneTitle). \(String(doneLine.characters))" }
        if cancelled { return "\(zone), \(Copy.widgetCancelled)" }
        return "\(zone). \(answer). \(starts)"
    }
}

/// The countdown: a system timer that ticks on its own (no update is needed per second), or a
/// fixed `42:00` for snapshots and the parity gallery.
struct LiveCountdown: View {
    let target: Date
    let fixedNow: Date?
    let role: TypeRole

    var body: some View {
        Group {
            if let fixedNow {
                Text(Self.clock(target.timeIntervalSince(fixedNow)))
            } else {
                Text(timerInterval: Date.now...max(target, .now), countsDown: true, showsHours: false)
            }
        }
        .type(role)
        .foregroundStyle(FSColor.gold)
        .multilineTextAlignment(.trailing)
    }

    static func clock(_ seconds: TimeInterval) -> String {
        let total = max(0, Int(seconds))
        return String(format: "%d:%02d", total / 60, total % 60)
    }
}

enum LiveType {
    static var name: TypeRole { TypeRole(face: FSFont.bodyBold, size: 16, lineHeight: 20, scalesWith: .headline) }
    static var clock: TypeRole { TypeRole(face: FSFont.bodyBold, size: 24, lineHeight: 28, scalesWith: .title2, tabular: true) }
    static var island: TypeRole { TypeRole(face: FSFont.bodyBold, size: 14, lineHeight: 17, scalesWith: .subheadline, tabular: true) }
    static var small: TypeRole { TypeRole(FSType.small) }
    static var percent: TypeRole { TypeRole(face: FSFont.bodyBold, size: 16, lineHeight: 28, scalesWith: .headline, tabular: true) }
}

enum LiveMetrics {
    static let padH: CGFloat = 16
    static let padV: CGFloat = 14
    static let gap: CGFloat = 10
    static let icon: CGFloat = 36
    static let radius: CGFloat = 20
    /// The name and the line under it, two 24 pt lines as the mock sets them.
    static let textBlock: CGFloat = 48
    /// The expanded Island: the row, then 10 and 12 more to the answers.
    static let answersGap: CGFloat = 12
}

/// The raid on the lock screen (flows F2): the instance icon, the zone and your answer as your
/// character, the countdown in gold, then `Starts 8:00 pm CST · 18 in, 3 late · 2 tanks, 5 heals`;
/// after the raid, `Onyxia's Lair done` and the night's summary.
struct RaidLiveLockView: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState
    let icons: WidgetIcons
    var fixedNow: Date?
    var zone: TimeZone = .current

    var body: some View {
        let figures = RaidLiveFigures(attributes, state, zone: zone)
        VStack(alignment: .leading, spacing: LiveMetrics.gap) {
            HStack(spacing: 12) {
                zoneIcon
                VStack(alignment: .leading, spacing: 0) {
                    Text(figures.done ? figures.doneTitle : figures.zone).type(LiveType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                    if figures.done {
                        Text(figures.doneLine).type(LiveType.small).foregroundStyle(FSColor.muted).lineLimit(2)
                    } else {
                        Text(figures.answer).type(LiveType.small).foregroundStyle(FSColor.muted).lineLimit(1)
                    }
                }
                .frame(maxWidth: .infinity, minHeight: LiveMetrics.textBlock, alignment: .leading)
                if figures.cancelled {
                    Text(Copy.widgetCancelled).type(LiveType.name).foregroundStyle(FSColor.wipe)
                } else if !figures.done {
                    LiveCountdown(target: state.startDate, fixedNow: fixedNow, role: LiveType.clock).fixedSize()
                }
            }
            if !figures.done {
                Text(figures.starts).type(LiveType.small).foregroundStyle(FSColor.muted).lineLimit(2)
            }
        }
        .padding(.horizontal, LiveMetrics.padH)
        .padding(.vertical, LiveMetrics.padV)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(figures.spoken)
    }

    private var zoneIcon: some View {
        WidgetItemIcon(icon: attributes.icon ?? "", size: LiveMetrics.icon, icons: icons)
    }
}

/// The Island's compact raid: `Ony` and the countdown.
struct RaidIslandCompactLeading: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState

    var body: some View {
        Text(RaidLiveFigures(attributes, state, zone: .current).short).type(LiveType.island).foregroundStyle(FSColor.strong)
    }
}

/// The Island's expanded raid: the icon, the zone and the countdown, then In, Late, Bench and Out
/// (flows F2), the current answer selected, each answered in place through RaidLiveAnswerIntent.
struct RaidIslandExpandedView: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState
    let icons: WidgetIcons
    var fixedNow: Date?

    var body: some View {
        VStack(spacing: LiveMetrics.gap) {
            HStack(spacing: LiveMetrics.gap) {
                WidgetItemIcon(icon: attributes.icon ?? "", size: LiveMetrics.icon, icons: icons)
                Text(attributes.zoneName).type(LiveType.island).foregroundStyle(FSColor.strong).lineLimit(1)
                Spacer(minLength: 0)
                if state.status == "cancelled" {
                    Text(Copy.widgetCancelled).type(LiveType.island).foregroundStyle(FSColor.wipe)
                } else {
                    LiveCountdown(target: state.startDate, fixedNow: fixedNow, role: LiveType.island).fixedSize()
                }
            }
            if state.status == "open" {
                HStack(spacing: 8) {
                    ForEach(RaidGlance.answers, id: \.self) { answer in
                        Button(intent: RaidLiveAnswerIntent(state: answer)) {
                            Text(WidgetFigures.stateWord(answer) ?? answer).lineLimit(1).minimumScaleFactor(0.7)
                        }
                        .buttonStyle(FSButtonStyle(state.myState == answer ? .primary : .secondary))
                        .accessibilityAddTraits(state.myState == answer ? .isSelected : [])
                    }
                }
                .padding(.top, LiveMetrics.answersGap)
            }
        }
    }
}

/// A server run on the lock screen (mock ios-live-activity `Sim running`): `Sim: {title}` with the
/// percent in gold over the progress; then `Finished` with the DPS, or `Failed` and `Open the run`.
struct SimLiveLockView: View {
    let attributes: SimActivityAttributes
    let state: SimActivityAttributes.ContentState

    var body: some View {
        let figures = SimLiveFigures(title: attributes.title, state: state)
        VStack(alignment: .leading, spacing: LiveMetrics.gap) {
            HStack(spacing: 12) {
                Text(Copy.liveSimTitle(title: figures.title)).type(LiveType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                Spacer(minLength: 0)
                if figures.finished || figures.failed {
                    Text(figures.status).type(LiveType.small)
                        .foregroundStyle(figures.finished ? FSColor.kill : FSColor.wipe)
                } else if let fraction = figures.fraction {
                    Text(fraction, format: .percent.precision(.fractionLength(0))).type(LiveType.percent).foregroundStyle(FSColor.gold)
                }
            }
            if let detail = figures.detail, figures.finished || figures.failed {
                Text(detail).type(LiveType.small).foregroundStyle(figures.failed ? FSColor.gold : FSColor.text)
            } else {
                LiveMeter(fraction: figures.fraction ?? 0)
            }
        }
        .padding(.horizontal, LiveMetrics.padH)
        .padding(.vertical, LiveMetrics.padV)
        .accessibilityElement(children: .combine)
        // The time left is spoken, not drawn: the line holds the title, the percent and the bar.
        .accessibilityValue(figures.left ?? "")
    }
}

/// The mock's Meter: 8 pt, the gold share over `lineSoft`.
struct LiveMeter: View {
    let fraction: Double

    var body: some View {
        GeometryReader { geo in
            ZStack(alignment: .leading) {
                RoundedRectangle(cornerRadius: 3).fill(FSColor.lineSoft)
                RoundedRectangle(cornerRadius: 3).fill(FSColor.gold).frame(width: geo.size.width * min(1, max(0, fraction)))
            }
        }
        .frame(height: 8)
        .accessibilityHidden(true)
    }
}
