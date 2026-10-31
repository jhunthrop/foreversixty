import AppIntents
import DesignSystem
import DesignTokens
import Models
import Routing
import SwiftUI
import WidgetKit

/// The words a raid widget draws when there is no raid to show (board F14.2): never a spinner,
/// never an empty box.
struct RaidEmptyState: View {
    let figures: WidgetFigures
    let lastRaid: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            switch figures.state {
            case .signedOut:
                WidgetHeader(title: Copy.appName)
                Text(Copy.widgetSignedOut).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                Spacer(minLength: 0)
                Text(Copy.widgetOpenApp).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
            case .noCharacter:
                WidgetHeader(title: Copy.widgetNextRaid)
                Text(Copy.widgetNoCharacter).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                Spacer(minLength: 0)
                Text(Copy.widgetOpenApp).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
            case .noGuild:
                WidgetHeader(title: Copy.tabRaids)
                Text(Copy.widgetNoGuild).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                Text(Copy.widgetNoGuildSub).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            case .noRaid, .raid:
                WidgetHeader(title: Copy.widgetNextRaid)
                Text(Copy.widgetNoRaid).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                if let lastRaid { Text(lastRaid).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            }
            if let stale = figures.stale { Spacer(minLength: 0); StaleLine(text: stale) }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .accessibilityElement(children: .combine)
        .widgetTextCap()
    }
}

struct StaleLine: View {
    let text: String

    var body: some View {
        Text(text).font(ScaledType.meta).foregroundStyle(FSColor.muted).lineLimit(1)
    }
}

/// The zone, the time and the countdown: the left half of both raid widgets.
struct RaidHeadline: View {
    let figures: WidgetFigures

    var body: some View {
        VStack(alignment: .leading, spacing: 2) {
            WidgetHeader(title: Copy.widgetNextRaid)
                .padding(.bottom, 4)
            Text((figures.zone ?? "").uppercased())
                .font(ScaledType.sec)
                .foregroundStyle(FSColor.strong)
                .lineLimit(2)
                .minimumScaleFactor(0.7)
            if let time = figures.time {
                Text(time).font(ScaledType.meta).foregroundStyle(FSColor.muted).lineLimit(1)
            }
            if let countdown = figures.countdown {
                Text(countdown).font(ScaledType.mono).foregroundStyle(FSColor.gold).lineLimit(1).minimumScaleFactor(0.7)
            }
        }
    }
}

/// The answer in words at a raid widget's foot: the tag, the `You:` line when the tag is about
/// the raid rather than you (Locked, Cancelled), and `Updated {n} h ago` under them when stale.
struct RaidFoot: View {
    let figures: WidgetFigures
    /// The medium widget's checked button already says In, Late or Out.
    var buttonsSayIt = false

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            if let tag = figures.tag, let kind = figures.tagKind, !(buttonsSayIt && [.in, .late, .out].contains(kind)) {
                WidgetTag(text: tag, kind: kind)
            }
            if figures.tagKind == .locked || figures.tagKind == .cancelled, let you = figures.you {
                Text(you).font(ScaledType.meta).foregroundStyle(FSColor.text).lineLimit(1).minimumScaleFactor(0.8)
            }
            if let stale = figures.stale { StaleLine(text: stale) }
        }
    }
}

/// Next raid, small (spec §5.6): the headline, then your answer as the foot tag.
struct NextRaidSmallView: View {
    let figures: WidgetFigures
    let lastRaid: String?

    var body: some View {
        if figures.state == .raid {
            VStack(alignment: .leading, spacing: 0) {
                RaidHeadline(figures: figures)
                Spacer(minLength: 4)
                RaidFoot(figures: figures)
            }
            .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
            .accessibilityElement(children: .ignore)
            .accessibilityLabel(figures.accessibility)
            .widgetTextCap()
        } else {
            RaidEmptyState(figures: figures, lastRaid: lastRaid)
        }
    }
}

/// Next raid, medium: the headline and the count on top; under them, 8 pt clear, the answer in
/// words on the left and In, Late, Out on the right. A press answers in place through
/// RaidAnswerIntent; with no spec saved every button opens the spec ask.
struct NextRaidMediumView: View {
    let figures: WidgetFigures
    let lastRaid: String?
    let raid: WidgetNextRaid?
    let characterKey: String?

    static let answers = ["in", "late", "out"]
    static let buttonGap: CGFloat = 8

    var body: some View {
        if figures.state == .raid, let raid {
            VStack(alignment: .leading, spacing: 0) {
                HStack(alignment: .top, spacing: 12) {
                    RaidHeadline(figures: figures)
                        .frame(maxWidth: .infinity, alignment: .topLeading)
                        .accessibilityElement(children: .ignore)
                        .accessibilityLabel(figures.accessibility)
                    if let signed = figures.signedUp {
                        VStack(alignment: .trailing, spacing: 0) {
                            Text(signed).font(ScaledType.monoDay).foregroundStyle(FSColor.strong).lineLimit(1)
                            Text(Copy.widgetSignedUpLabel).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        }
                        .accessibilityElement(children: .combine)
                    }
                }
                Spacer(minLength: Self.buttonGap)
                HStack(alignment: .bottom, spacing: 8) {
                    RaidFoot(figures: figures, buttonsSayIt: figures.buttons)
                    Spacer(minLength: 0)
                    if figures.buttons { buttons(raid) }
                }
            }
            .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
            .widgetTextCap()
        } else {
            RaidEmptyState(figures: figures, lastRaid: lastRaid)
        }
    }

    private func buttons(_ raid: WidgetNextRaid) -> some View {
        HStack(spacing: 6) {
            ForEach(Self.answers, id: \.self) { state in
                answerButton(state, raid: raid)
            }
        }
        .fixedSize()
    }

    @ViewBuilder
    private func answerButton(_ state: String, raid: WidgetNextRaid) -> some View {
        let label = AnswerButtonLabel(
            title: figures.notSaved == state ? Copy.widgetNotSaved : (WidgetFigures.stateWord(state) ?? state),
            checked: figures.checked == state, failed: figures.notSaved == state
        )
        if figures.tagKind == .chooseSpec {
            Link(destination: DeepLink.raid(guild: raid.guildId, raid: raid.raidId, spec: "ask", state: state).url) { label }
        } else {
            Button(intent: RaidAnswerIntent(state: state, character: characterKey)) { label }
                .buttonStyle(.plain)
        }
    }
}

/// One of In, Late, Out: tonal, the answer checked with the kill ground and a 2 pt kill underline.
struct AnswerButtonLabel: View {
    let title: String
    let checked: Bool
    let failed: Bool

    var body: some View {
        Text(title)
            .font(ScaledType.button)
            .foregroundStyle(failed ? FSColor.wipe : (checked ? FSColor.strong : FSColor.nav))
            // `Not saved` wraps to two lines rather than cut (spec §9: an answer never collapses).
            .lineLimit(failed ? 2 : 1)
            .multilineTextAlignment(.center)
            .minimumScaleFactor(0.7)
            .frame(minWidth: WidgetMetrics.buttonHeight, minHeight: WidgetMetrics.buttonHeight)
            .padding(.horizontal, 4)
            .background(checked ? FSColor.kill.opacity(0.16) : Color.white.opacity(0.06), in: .rect(cornerRadius: 12))
            .overlay(alignment: .bottom) {
                if checked { Rectangle().fill(FSColor.kill).frame(height: 2).padding(.horizontal, 8) }
            }
            .accessibilityAddTraits(checked ? [.isButton, .isSelected] : .isButton)
    }
}

// MARK: Lock screen (monochrome: the system tints them)

/// `Onyxia's Lair · Thu 8 pm · In`
struct RaidInlineView: View {
    let figures: WidgetFigures

    var body: some View {
        Text(figures.inline).accessibilityLabel(figures.accessibility)
    }
}

/// The zone, `Thu 8:00 pm · in 2 days 2 h`, and `You: In · DPS`.
struct RaidRectangularView: View {
    let figures: WidgetFigures

    var body: some View {
        VStack(alignment: .leading, spacing: 1) {
            if figures.state == .raid {
                Text(figures.zone ?? "").font(.headline).lineLimit(1).widgetAccentable()
                Text([figures.time, figures.countdown ?? figures.tag].compactMap { $0 }.joined(separator: " · "))
                    .font(.caption).lineLimit(1).minimumScaleFactor(0.7)
                Text(figures.tagKind == .locked || figures.tagKind == .cancelled ? (figures.tag ?? "") : (figures.you ?? ""))
                    .font(.caption).lineLimit(1)
            } else {
                Text(Copy.widgetNextRaid).font(.headline).lineLimit(1).widgetAccentable()
                Text(figures.inline).font(.caption).lineLimit(2)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(figures.accessibility)
    }
}

/// A ring of signed up over the raid's size, the count in the middle.
struct RaidSignedUpCircularView: View {
    let figures: WidgetFigures

    var body: some View {
        if let count = figures.signedUpCount, let size = figures.size, size > 0 {
            Gauge(value: Double(min(count, size)), in: 0...Double(size)) {
                Text(Copy.widgetSignedUpLabel)
            } currentValueLabel: {
                // The unit rides under the figure: `24` over `of 40` (widget.ofSize, split at the count).
                VStack(spacing: -2) {
                    Text("\(count)").font(.system(.title3, design: .rounded, weight: .semibold))
                    Text(Self.ofSize(count: count, size: size)).font(.system(size: 10, weight: .medium, design: .rounded))
                }
                .minimumScaleFactor(0.7)
            }
            .gaugeStyle(.accessoryCircularCapacity)
            .widgetAccentable()
            .accessibilityLabel(figures.signedUp.map { "\($0) \(Copy.widgetSignedUpLabel)" } ?? "")
        } else {
            CircularWords(text: figures.state == .raid ? (figures.signedUp ?? "") : "—")
        }
    }
}

extension RaidSignedUpCircularView {
    /// `of 40`: widget.ofSize with the count taken off its front, so the words stay the catalog's.
    static func ofSize(count: Int, size: Int) -> String {
        let whole = Copy.widgetOfSize(n: count, size: size)
        let lead = "\(count)"
        return whole.hasPrefix(lead) ? String(whole.dropFirst(lead.count)).trimmingCharacters(in: .whitespaces) : whole
    }
}

/// Days (or hours, or minutes) to the raid: `2d`.
struct RaidCountdownCircularView: View {
    let figures: WidgetFigures

    var body: some View {
        CircularWords(text: figures.short ?? (figures.tagKind == .cancelled ? "×" : "—"))
            .accessibilityLabel(figures.countdown ?? figures.inline)
    }
}

struct CircularWords: View {
    let text: String

    var body: some View {
        ZStack {
            AccessoryWidgetBackground()
            Text(text).font(.system(.title3, design: .rounded, weight: .semibold)).minimumScaleFactor(0.6)
                .widgetAccentable()
        }
    }
}
