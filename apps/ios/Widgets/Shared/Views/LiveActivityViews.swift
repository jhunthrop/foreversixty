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

/// The raid countdown on the lock screen (board F14.3): the zone and guild, the countdown with
/// `until invites at 7:45 pm`, the lead bar, and `You are In · DPS · {name}` with `38 of 40`.
struct RaidLiveLockView: View {
    let attributes: RaidActivityAttributes
    let state: RaidActivityAttributes.ContentState
    var fixedNow: Date?
    var zone: TimeZone = .current

    var body: some View {
        let figures = WidgetFigures(.live(attributes, state), now: fixedNow ?? .now, zone: zone)
        VStack(alignment: .leading, spacing: 10) {
            HStack(spacing: 8) {
                SealMark()
                Text(attributes.zoneName).font(ScaledType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                Spacer()
                Text(attributes.guildName).font(ScaledType.meta).foregroundStyle(FSColor.muted).lineLimit(1)
            }
            if state.status == "cancelled" {
                Text(Copy.widgetCancelled).font(ScaledType.stat).foregroundStyle(FSColor.wipe)
            } else {
                HStack(alignment: .firstTextBaseline, spacing: 10) {
                    LiveCountdown(target: state.targetDate, fixedNow: fixedNow, font: .system(size: 36, weight: .semibold, design: .monospaced))
                    Text(figures.liveUntil ?? "").font(ScaledType.body).foregroundStyle(FSColor.text).lineLimit(2)
                }
                LiveLeadBar(target: state.targetDate, fixedNow: fixedNow)
            }
            HStack {
                Text(Self.you(attributes, state)).font(ScaledType.body).foregroundStyle(FSColor.text).lineLimit(1)
                Spacer()
                if let signed = figures.signedUp { Text(signed).font(ScaledType.mono).foregroundStyle(FSColor.text) }
            }
        }
        .padding(16)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Self.spoken(attributes, state, now: fixedNow ?? .now))
    }

    static func you(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState) -> String {
        Copy.liveYou(
            state: WidgetFigures.stateWord(state.myState) ?? Copy.widgetNotSignedUp,
            role: WidgetFigures.roleWord(state.role), name: attributes.characterName
        )
    }

    /// `Invites for Onyxia's Lair in 38 minutes. You are In.`; the countdown is not read per tick.
    static func spoken(_ attributes: RaidActivityAttributes, _ state: RaidActivityAttributes.ContentState, now: Date) -> String {
        if state.status == "cancelled" { return "\(attributes.zoneName), \(Copy.widgetCancelled)" }
        let minutes = max(0, Countdown(from: now, to: state.targetDate).totalMinutes)
        return Copy.liveA11yRaid(
            zone: attributes.zoneName, span: Copy.widgetA11yMinutes(n: minutes),
            state: WidgetFigures.stateWord(state.myState) ?? Copy.widgetNotSignedUp
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
                    Text(attributes.zoneName).font(ScaledType.name).foregroundStyle(FSColor.strong).lineLimit(1)
                    Text([figures.liveUntil, figures.you].compactMap { $0 }.joined(separator: " · "))
                        .font(ScaledType.meta).foregroundStyle(FSColor.text).lineLimit(2)
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
