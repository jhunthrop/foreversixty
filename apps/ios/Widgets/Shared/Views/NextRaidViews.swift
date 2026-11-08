import AppIntents
import DesignSystem
import DesignTokens
import Models
import Routing
import SwiftUI
import WidgetKit

/// The raid side of every home-screen widget (mocks ios-widgets-*): the label, the countdown, the
/// zone, then the time with its zone and the answer naming its character (flows E7).
struct RaidGlanceColumn: View {
    let glance: RaidGlance

    var body: some View {
        VStack(alignment: .leading, spacing: WidgetMetrics.gap) {
            Text(glance.label).type(WidgetType.label).foregroundStyle(FSColor.muted).lineLimit(1)
            if !glance.big.isEmpty {
                Text(glance.big).type(WidgetType.big).foregroundStyle(FSColor.strong).lineLimit(1).minimumScaleFactor(0.6)
            }
            Text(glance.name).type(WidgetType.name).foregroundStyle(FSColor.strong).lineLimit(3)
            ForEach(Array(glance.lines.enumerated()), id: \.offset) { _, line in
                Text(line).type(WidgetType.sub)
                    .foregroundStyle(glance.kind == .empty ? FSColor.gold : FSColor.muted)
                    .lineLimit(2)
            }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(glance.spoken)
    }
}

/// Next raid, small: the raid column alone.
struct NextRaidSmallView: View {
    let glance: RaidGlance

    var body: some View {
        RaidGlanceColumn(glance: glance).widgetTextCap()
    }
}

/// Next raid, medium (flows F1 from the widget): the raid column, and In, Late, Bench and Out as
/// App Intents, the current answer selected, so a change of plan is one press. With no spec saved
/// every answer opens the spec ask; a locked or cancelled raid shows the column alone.
struct RaidAnswerMediumView: View {
    let glance: RaidGlance
    let raid: WidgetNextRaid?
    let characterKey: String?

    private static let columns = [GridItem(.flexible(), spacing: 8), GridItem(.flexible(), spacing: 8)]

    var body: some View {
        HStack(alignment: .center, spacing: 16) {
            RaidGlanceColumn(glance: glance)
            if glance.canAnswer || glance.asksSpec, let raid {
                WidgetDivider()
                LazyVGrid(columns: Self.columns, spacing: 8) {
                    ForEach(RaidGlance.answers, id: \.self) { state in answer(state, raid: raid) }
                }
                .frame(maxWidth: .infinity)
            }
        }
        .widgetTextCap()
    }

    @ViewBuilder
    private func answer(_ state: String, raid: WidgetNextRaid) -> some View {
        let failed = glance.notSaved == state
        let title = failed ? Copy.widgetNotSaved : (WidgetFigures.stateWord(state) ?? state)
        let style = FSButtonStyle(glance.selected == state ? .primary : (failed ? .danger : .secondary))
        if glance.asksSpec {
            Link(destination: DeepLink.raid(guild: raid.guildId, raid: raid.raidId, spec: "ask", state: state).url) {
                Text(title).lineLimit(1).minimumScaleFactor(0.7)
            }
            .buttonStyle(style)
        } else {
            Button(intent: RaidAnswerIntent(state: state, character: characterKey)) {
                Text(title).lineLimit(1).minimumScaleFactor(0.7)
            }
            .buttonStyle(style)
            .accessibilityAddTraits(glance.selected == state ? .isSelected : [])
        }
    }
}

// MARK: Lock screen (monochrome: the system tints them)

/// `33d 8h · Onyxia's Lair`, as the mock's lock-screen raid line.
struct RaidInlineView: View {
    let glance: RaidGlance

    var body: some View {
        Text(glance.big.isEmpty ? glance.name : "\(glance.big) · \(glance.name)")
            .accessibilityLabel(glance.spoken)
    }
}

/// The zone, then `Wed 9 Dec, 8:00 pm CST` and `In, as Thandrel, Damage`.
struct RaidRectangularView: View {
    let glance: RaidGlance

    var body: some View {
        VStack(alignment: .leading, spacing: 1) {
            Text(glance.big.isEmpty ? glance.name : "\(glance.big) · \(glance.name)")
                .font(.headline).lineLimit(1).minimumScaleFactor(0.7).widgetAccentable()
            ForEach(Array(glance.lines.prefix(2).enumerated()), id: \.offset) { _, line in
                Text(line).font(.caption).lineLimit(1).minimumScaleFactor(0.7)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(glance.spoken)
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
            CircularWords(text: figures.state == .raid ? (figures.signedUp ?? "") : "-")
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
        CircularWords(text: figures.short ?? (figures.tagKind == .cancelled ? "×" : "-"))
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
