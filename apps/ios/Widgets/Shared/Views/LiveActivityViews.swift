import AppIntents
import DesignSystem
import DesignTokens
import Models
import SwiftUI
import WidgetKit

extension WidgetSnapshot {
    /// The raid Live Activity as a snapshot, so its words come from the same WidgetFigures as
    /// every widget (one derivation, one clock).
    static func live(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState) -> WidgetSnapshot {
        let raid = WidgetNextRaid(
            raidId: Int64(attributes.raidID) ?? 0, guildId: Int64(attributes.guildID) ?? 0, guildName: attributes.guildName,
            zone: "", zoneName: attributes.zoneName, startsAt: state.startDate, inviteAt: state.inviteDate,
            pullAt: state.startDate, endsAt: state.startDate.addingTimeInterval(WidgetFigures.liveTail), size: state.size,
            signedUp: state.signedUp, status: state.status,
            my: WidgetMine(state: state.myState, spec: "live", role: state.role), canSignUp: state.status == "open"
        )
        let widget = Models.Widget(
            generatedAt: .now, character: WidgetCharacter(key: "", name: attributes.characterName),
            guild: WidgetGuild(id: raid.guildId, name: attributes.guildName), nextRaid: raid, live: WidgetLive(simRunsRunning: 0)
        )
        return WidgetSnapshot(widget: widget, fetchedAt: .now)
    }
}

/// The countdown: a system timer that ticks on its own (no update is needed per second), or a
/// fixed `38:12` for snapshots and review captures.
struct LiveCountdown: View {
    let target: Date
    let fixedNow: Date?
    let font: Font

    var body: some View {
        Group {
            if let fixedNow {
                Text(Self.clock(target.timeIntervalSince(fixedNow)))
            } else {
                Text(timerInterval: Date.now...max(target, .now), countsDown: true, showsHours: false)
            }
        }
        .font(font)
        .monospacedDigit()
        .foregroundStyle(FSColor.gold)
    }

    static func clock(_ seconds: TimeInterval) -> String {
        let total = max(0, Int(seconds))
        return String(format: "%d:%02d", total / 60, total % 60)
    }
}

/// The gold progress over the lead, from an hour before the target to the target.
struct LiveLeadBar: View {
    let target: Date
    let fixedNow: Date?

    var body: some View {
        let start = target.addingTimeInterval(-WidgetFigures.liveLead)
        Group {
            if let fixedNow {
                ProgressView(value: min(1, max(0, fixedNow.timeIntervalSince(start) / WidgetFigures.liveLead)))
            } else {
                ProgressView(timerInterval: start...max(target, start), countsDown: false) { EmptyView() } currentValueLabel: { EmptyView() }
            }
        }
        .progressViewStyle(.linear)
        .tint(FSColor.gold)
        .frame(height: 6)
        .accessibilityHidden(true)
    }
}

/// Raid night on the lock screen (round 5, m-ios-lock; flows F2): the zone's icon, the raid, the
/// guild and your answer, the countdown in gold; `Starts 8:00 pm CST · 18 in, 3 late · 2 tanks, 5 heals`;
/// then In, Late, Bench and Out, the current answer teal, each answering without opening the app.
struct RaidLiveLockView: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState
    var fixedNow: Date?
    var zone: TimeZone = .current
    var icons: WidgetIcons = .shared

    static let iconSize: CGFloat = 44
    static let answers = ["in", "late", "bench", "out"]
    static let buttonHeight: CGFloat = 36

    var body: some View {
        VStack(alignment: .leading, spacing: 10) {
            HStack(spacing: 10) {
                zoneIcon
                VStack(alignment: .leading, spacing: 2) {
                    Text(attributes.zoneName).font(.custom(FSFont.bodyBold, size: 17, relativeTo: .headline))
                        .foregroundStyle(FSColor.strong).lineLimit(1)
                    Text(Self.answerLine(attributes, state)).font(.custom(FSFont.body, size: 13, relativeTo: .footnote))
                        .foregroundStyle(FSColor.text).lineLimit(1)
                }
                Spacer(minLength: 8)
                if state.status == "cancelled" {
                    Text(Copy.widgetCancelled).font(ScaledType.name).foregroundStyle(FSColor.wipe)
                } else {
                    LiveCountdown(target: state.targetDate, fixedNow: fixedNow, font: .custom(FSFont.bodyBold, size: 28, relativeTo: .title))
                }
            }
            Text(Self.factLine(state, zone: zone)).font(.custom(FSFont.body, size: 13, relativeTo: .footnote))
                .foregroundStyle(FSColor.text).lineLimit(2)
            if state.status != "cancelled" {
                HStack(spacing: 6) {
                    ForEach(Self.answers, id: \.self) { answer in
                        Button(intent: RaidLiveAnswerIntent(state: answer)) {
                            Text(WidgetFigures.stateWord(answer) ?? answer)
                                .font(.custom(FSFont.bodyBold, size: 13, relativeTo: .footnote))
                                .foregroundStyle(FSColor.strong)
                                .frame(maxWidth: .infinity, minHeight: Self.buttonHeight)
                                .background(answer == state.myState ? FSColor.tealEdge : Color.white.opacity(0.1), in: .capsule)
                        }
                        .buttonStyle(.plain)
                    }
                }
            }
        }
        .padding(14)
        .accessibilityElement(children: .contain)
        .accessibilityLabel(Self.spoken(attributes, state, now: fixedNow ?? .now))
    }

    @ViewBuilder
    private var zoneIcon: some View {
        if let icon = attributes.zoneIcon, let image = icons.image(icon) {
            Image(uiImage: image).resizable().scaledToFill()
                .frame(width: Self.iconSize, height: Self.iconSize)
                .clipShape(.rect(cornerRadius: 6))
                .overlay(RoundedRectangle(cornerRadius: 6).stroke(FSColor.line, lineWidth: 1))
                .accessibilityHidden(true)
        } else {
            SealMark().frame(width: Self.iconSize, height: Self.iconSize)
        }
    }

    /// `Olympus XXVII · You are In`.
    static func answerLine(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState) -> String {
        Copy.liveAnswer(guild: attributes.guildName, state: WidgetFigures.stateWord(state.myState) ?? Copy.widgetNotSignedUp)
    }

    /// `Starts 8:00 pm CST · 18 in, 3 late · 2 tanks, 5 heals` (invites first when the raid has them);
    /// the counts are left out by a server that does not send them.
    static func factLine(_ state: RaidActivityAttributes.ContentState, zone: TimeZone) -> String {
        let clock = WidgetClock(zone: zone)
        let when = state.inviteDate.map { Copy.liveInvitesAt(time: clock.zonedTime($0)) } ?? Copy.liveStartsAt(time: clock.zonedTime(state.startDate))
        guard let inCount = state.inCount, let late = state.late, let tanks = state.tanks, let healers = state.healers else { return when }
        return "\(when) · \(Copy.liveCounts(in: inCount, late: late, tanks: tanks, healers: healers))"
    }

    /// `Invites for Onyxia's Lair in 38 minutes. You are In.`; the countdown is not read per tick.
    static func spoken(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState, now: Date) -> String {
        if state.status == "cancelled" { return "\(attributes.zoneName), \(Copy.widgetCancelled)" }
        let span = Copy.widgetA11yMinutes(n: max(0, Countdown(from: now, to: state.targetDate).totalMinutes))
        let word = WidgetFigures.stateWord(state.myState) ?? Copy.widgetNotSignedUp
        return state.inviteAt == nil
            ? Copy.liveA11yRaidPull(zone: attributes.zoneName, span: span, state: word)
            : Copy.liveA11yRaid(zone: attributes.zoneName, span: span, state: word)
    }

    /// `You are In · DPS · Thandrel`: the Island's expanded line.
    static func you(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState) -> String {
        Copy.liveYou(
            state: WidgetFigures.stateWord(state.myState) ?? Copy.widgetNotSignedUp,
            role: WidgetFigures.roleWord(state.role), name: attributes.characterName
        )
    }
}

/// The Island's expanded raid: the zone, `until invites at …`, the countdown, then Late and Out.
struct RaidIslandExpandedView: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState
    var fixedNow: Date?
    var zone: TimeZone = .current

    var body: some View {
        let figures = WidgetFigures(.live(attributes, state), now: fixedNow ?? .now, zone: zone)
        VStack(spacing: 12) {
            HStack(alignment: .center) {
                VStack(alignment: .leading, spacing: 2) {
                    HStack(spacing: 6) {
                        SealMark()
                        Text(attributes.zoneName).font(ScaledType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                    }
                    Text(figures.liveUntil ?? "").font(ScaledType.meta).foregroundStyle(FSColor.text).lineLimit(1)
                    if state.status != "cancelled" {
                        Text(RaidLiveLockView.you(attributes, state)).font(ScaledType.meta).foregroundStyle(FSColor.text).lineLimit(1)
                    }
                    if let signed = figures.signedUp {
                        Text(signed).font(ScaledType.tabularMeta).foregroundStyle(FSColor.muted)
                    }
                }
                Spacer()
                if state.status == "cancelled" {
                    Text(Copy.widgetCancelled).font(ScaledType.name).foregroundStyle(FSColor.wipe)
                } else {
                    LiveCountdown(target: state.targetDate, fixedNow: fixedNow, font: .system(size: 32, weight: .semibold, design: .monospaced))
                }
            }
            if state.status != "cancelled" {
                HStack(spacing: 10) {
                    ForEach(["late", "out"], id: \.self) { answer in
                        Button(intent: RaidLiveAnswerIntent(state: answer)) {
                            Text(WidgetFigures.stateWord(answer) ?? answer)
                                .font(ScaledType.button)
                                .foregroundStyle(FSColor.strong)
                                .frame(maxWidth: .infinity, minHeight: WidgetMetrics.buttonHeight)
                                .background(Color.white.opacity(0.10), in: .capsule)
                        }
                        .buttonStyle(.plain)
                    }
                }
            }
        }
    }
}

/// A server run on the lock screen (board F14.4): `{Spec} · {kind}`, the progress, then
/// `Finished` with the DPS or `Failed` with `Open the run`.
struct SimLiveLockView: View {
    let attributes: SimActivityAttributes
    let state: SimActivityAttributes.ContentState

    var body: some View {
        let figures = SimLiveFigures(title: attributes.title, state: state)
        VStack(alignment: .leading, spacing: 10) {
            HStack {
                Image(systemName: "flask").foregroundStyle(FSColor.gold).accessibilityHidden(true)
                Text(figures.title).font(ScaledType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                Spacer()
                Text(figures.status).font(ScaledType.meta)
                    .foregroundStyle(figures.finished ? FSColor.kill : (figures.failed ? FSColor.wipe : FSColor.text))
            }
            if figures.finished, let value = figures.value {
                HStack(alignment: .firstTextBaseline, spacing: 8) {
                    Text(value).font(ScaledType.stat).foregroundStyle(FSColor.strong)
                    Text(figures.valueRest ?? "").font(ScaledType.body).foregroundStyle(FSColor.text)
                }
            } else {
                HStack {
                    if let detail = figures.detail {
                        Text(detail).font(ScaledType.body).foregroundStyle(figures.failed ? FSColor.gold : FSColor.text)
                    }
                    Spacer()
                    if let left = figures.left { Text(left).font(ScaledType.body).foregroundStyle(FSColor.text) }
                }
                if let fraction = figures.fraction {
                    ProgressView(value: fraction).progressViewStyle(.linear).tint(FSColor.gold).frame(height: 6)
                }
            }
        }
        .padding(16)
        .accessibilityElement(children: .combine)
    }
}
