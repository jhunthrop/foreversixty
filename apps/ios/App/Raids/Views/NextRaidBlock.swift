import DesignSystem
import SwiftUI

/// What the next raid's card and line hand back.
struct NextRaidActions {
    var answer: (SignupState) -> Void = { _ in }
    var late: (Int) -> Void = { _ in }
    /// Try again after a failed write.
    var retry: () -> Void = {}
    var openRaid: () -> Void = {}
    var openReady: () -> Void = {}
    var previewMention: () -> Void = {}
    var openApplicants: () -> Void = {}
}

/// The next raid on raid day (flows spec F1, round 5; mocks ios-home-raid-day, -late, -officer and
/// m-ios-home): the lead panel, ornate on the cave. The zone, the guild and the time, the countdown,
/// who the answer is for (E7), the four answers with Late's minutes under them, who is coming, what is
/// not ready, and an officer's open items.
struct NextRaidCardView: View {
    let content: NextRaidContent
    var actions = NextRaidActions()

    static let tallyTop: CGFloat = 14

    var body: some View {
        LeadPanel {
            Text(content.zone)
                .type(FSType.cardTitle)
                .foregroundStyle(FSColor.cream)
                .shadow(color: .black.opacity(PanelShadow.opacity), radius: PanelShadow.radius, y: PanelShadow.drop)
                .accessibilityAddTraits(.isHeader)
            Text(content.place).type(FSType.small).foregroundStyle(FSColor.text)
            CardClock(parts: content.clock, label: content.clockLabel).padding(.top, FSSpace.space2)
            // AS5: once answered the Segmented is the answer; the line names only the character (E7), if any.
            if !content.answer.isEmpty {
                Text(content.answer).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, FSSpace.space3)
            }
            RaidAnswerControl(content: content, actions: actions, question: true)
                .padding(.top, content.answer.isEmpty ? FSSpace.space3 : FSSpace.space2)
            // Against a target the meters; without one the counts (`Tanks 3 · Healers 6 · Damage 15`).
            let meters = content.tally.compactMap(\.meter)
            if !meters.isEmpty, meters.count == content.tally.count {
                FSTally(meters).padding(.top, Self.tallyTop)
            } else if !content.tally.isEmpty {
                RaidTallyLine(roles: content.tally).padding(.top, Self.tallyTop)
            }
            if let ready = content.ready {
                Hairline(FSColor.line).padding(.top, Self.tallyTop)
                Button(action: actions.openReady) {
                    PanelRow(ready.title, sub: Text(ready.sub), slim: true, last: content.officer == nil)
                }
                .buttonStyle(RowPressStyle())
            }
            if let officer = content.officer {
                OfficerItemsView(items: officer, actions: actions).padding(.top, FSSpace.space1)
            }
        }
        .parityRegion(.countdowns)
    }
}

/// The lead panel (round 5): the ornate frame over the cave, 18 inside (16 at the foot).
struct LeadPanel<Content: View>: View {
    static var pad: CGFloat { 18 }

    @ViewBuilder let content: Content

    var body: some View {
        PanelBlock(
            frame: .ornate, art: .cave,
            padding: EdgeInsets(top: Self.pad, leading: Self.pad, bottom: FSSpace.space4, trailing: Self.pad)
        ) { content }
    }
}

/// `02 hours  48 min` on the card (round 5: 36/38 semibold, the units 12/16 under, 18 apart), read
/// aloud as one timer.
struct CardClock: View {
    let parts: [FSClockPart]
    let label: String

    static let gap: CGFloat = 18
    @ScaledMetric(relativeTo: FSType.countdown.scalesWith) private var valueLine: CGFloat = FSType.countdown.lineHeight

    var body: some View {
        if !parts.isEmpty {
            HStack(alignment: .top, spacing: Self.gap) {
                ForEach(Array(parts.enumerated()), id: \.offset) { _, part in
                    VStack(alignment: .leading, spacing: 0) {
                        Text(part.value)
                            .type(FSType.countdown)
                            .foregroundStyle(FSColor.strong)
                            .frame(height: valueLine)
                        Text(part.unit)
                            .type(FSType.figureUnit)
                            .foregroundStyle(FSColor.muted)
                    }
                }
            }
            .dynamicTypeSize(...DynamicTypeSize.accessibility1)
            .accessibilityElement(children: .ignore)
            .accessibilityLabel(label)
        }
    }
}

/// The role tally as one line (rows 1 and 2): `Tanks 3 · Healers 6 · Damage 15`, or against a target
/// `Healers 5/8 · 3 short` with the short count in wipe red.
struct RaidTallyLine: View {
    let roles: [RoleLineContent]

    var body: some View {
        roles.enumerated().reduce(Text("")) { text, entry in
            let role = entry.element
            let tone: Color = switch role.tone {
            case .short: FSColor.wipe
            case .met: FSColor.kill
            case .plain: FSColor.strong
            }
            let part = Text("\(role.role) \(Text(role.compactTally).foregroundColor(tone))").foregroundColor(FSColor.muted)
            return entry.offset == 0 ? part : Text("\(text)\(Text(RaidFlowPresenter.separator).foregroundColor(FSColor.muted))\(part)")
        }
        .type(RowType.sub)
        .frame(maxWidth: .infinity, alignment: .leading)
        .parityItem(.row)
    }
}

/// An officer's open items (F3): who has not answered with Preview mention, the applicants. On
/// Home's raid-day card and under the Raid tab's Needs you.
struct OfficerItemsView: View {
    let items: OfficerContent
    let actions: NextRaidActions

    var body: some View {
        VStack(spacing: 0) {
            if let unanswered = items.unanswered {
                PanelRow(
                    unanswered.title, sub: unanswered.sub.isEmpty ? nil : Text(unanswered.sub), last: items.applicants == nil
                ) {
                    EmptyView()
                } trail: {
                    Button(Copy.homeOfficerPreviewMention, action: actions.previewMention)
                        .buttonStyle(FSButtonStyle(fullWidth: false))
                }
            }
            if let applicants = items.applicants {
                Button(action: actions.openApplicants) {
                    PanelRow(applicants.title, sub: applicants.sub.isEmpty ? nil : Text(applicants.sub), last: true)
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// The next raid on another day (mocks ios-home-60, -answered): one line with its answers, on a panel.
struct NextRaidLineView: View {
    let content: NextRaidContent
    let build: String
    var actions = NextRaidActions()

    var body: some View {
        PanelBlock {
            Button(action: actions.openRaid) {
                PanelRow(content.line.title, sub: Text(content.line.sub), last: true) {
                    ItemIconView(icon: content.icon, quality: nil, build: build, size: FSSize.icon40).parityItem(.image)
                }
            }
            .buttonStyle(RowPressStyle())
            // The role tally on one line (`Healers 5/8 · 3 short`, a short role in wipe red; `Healers 6` without a target).
            // Round 4, row 6: 8 pt between the tally and the Segmented, as Android.
            if !content.tally.isEmpty { RaidTallyLine(roles: content.tally).padding(.bottom, FSSpace.space2) }
            RaidAnswerControl(content: content, actions: actions, question: true)
                .padding(.bottom, content.tally.isEmpty ? FSSpace.space1 : 0)
        }
        .parityRegion(.countdowns)
    }
}

/// The answer as Home's card and the Raid tab's head draw it (F1, brand B1): In, Late, Bench, Out as
/// the one answer control ([RaidAnswerButtons]); after Late its minutes, the chosen one selected; the status line (`Saved.`,
/// the refusal with Try again, offline's sentence). Offline the answers disable.
struct RaidAnswerControl: View {
    let content: NextRaidContent
    let actions: NextRaidActions
    /// Home's card asks `How late?` over the minutes; the Raid tab's head sets them bare.
    var question = false

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            RaidAnswerButtons(selected: content.selected, busy: content.busy, answer: actions.answer)
                .disabled(!content.answersEnabled)
            if content.lateOpen {
                if question { Text(Copy.pushLateQuestion).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, FSSpace.space3) }
                // Under `How late?` the chips sit straight under the question (the mock's `.me + .chips`).
                LateMinutesRow(chosen: content.lateMinutes, height: question ? FSButtonStyle.smallHeight : FSTarget.min, late: actions.late)
                    .padding(.top, question ? 0 : FSSpace.space2)
                    .disabled(!content.answersEnabled)
                    .accessibilityElement(children: .contain)
                    .accessibilityLabel(Copy.raidTabHowLate)
            }
            if let status = content.status {
                HStack(spacing: FSSpace.space2) {
                    StatusLineView(status: status, reserves: false)
                    if case .error = status, content.answersEnabled {
                        Button(Copy.commonTryAgain, action: actions.retry).buttonStyle(FSButtonStyle(.ghost, fullWidth: false, small: true))
                    }
                }
                .padding(.top, FSSpace.space2)
            }
        }
    }
}

/// In, Late, Bench, Out as one segmented control (Android `FlowSegmented`): one pill on the 1 pt rule
/// with hairline dividers, 44 high; the answer given fills its segment in the teal edge, the others sit
/// on the opaque panel in body ink, and nothing is filled before an answer. The one answer control:
/// Home, the Raid tab, the raid page and raid night all draw it. `pending` rings the answer waiting to
/// send; `busy` reads the shown one as saving. Disabled (offline, locked), the whole control is at 50%.
/// At the accessibility sizes the answers go two by two on the field radius, so no word breaks.
struct RaidAnswerButtons: View {
    let selected: SignupState?
    var pending: SignupState?
    var busy = false
    var accessibilityLabel = Copy.homeRaidAnswers
    let answer: (SignupState) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize
    @Environment(\.isEnabled) private var isEnabled
    @Environment(\.answersOffline) private var offline

    /// The pending ring's inset, dash and width inside a segment.
    private static let pendingInset: CGFloat = 4
    private static let pendingDash: [CGFloat] = [4, 3]
    private static let pendingLine: CGFloat = 1.5
    private static let disabledOpacity = 0.5

    private var isGrid: Bool { typeSize.isAccessibilitySize }
    private var shape: AnyShape {
        isGrid ? AnyShape(RoundedRectangle(cornerRadius: FSRadius.field)) : AnyShape(Capsule())
    }

    var body: some View {
        let rows = isGrid ? [Array(SignupState.allCases.prefix(2)), Array(SignupState.allCases.suffix(2))] : [SignupState.allCases]
        VStack(spacing: 0) {
            ForEach(Array(rows.enumerated()), id: \.offset) { index, row in
                if index > 0 { FSColor.rule.frame(height: FSStroke.hairline) }
                HStack(spacing: 0) {
                    ForEach(Array(row.enumerated()), id: \.element) { position, state in
                        if position > 0 { FSColor.rule.frame(width: FSStroke.hairline) }
                        segment(state)
                    }
                }
                .fixedSize(horizontal: false, vertical: true)
            }
        }
        .background(FSColor.raised)
        .clipShape(shape)
        // The 1 pt rule as a ring of fill round the pill: a stroked capsule draws stray ticks at its ends.
        .padding(FSStroke.hairline)
        .background(FSColor.rule, in: shape)
        .disabled(offline)
        .opacity(isEnabled && !offline ? 1 : Self.disabledOpacity)
        .accessibilityElement(children: .contain)
        .accessibilityLabel(accessibilityLabel)
    }

    private func segment(_ state: SignupState) -> some View {
        let on = state == selected
        return Button { answer(state) } label: {
            Text(state.label)
                .font(ScaledType.segment)
                .lineLimit(1)
                .minimumScaleFactor(0.6)
                .foregroundStyle(on ? FSColor.strong : FSColor.text)
                .frame(maxWidth: .infinity, minHeight: FSTarget.min)
                .background(on ? FSColor.tealEdge : .clear)
                .overlay {
                    if state == pending {
                        RoundedRectangle(cornerRadius: FSRadius.field)
                            .strokeBorder(FSColor.gold, style: StrokeStyle(lineWidth: Self.pendingLine, dash: Self.pendingDash))
                            .padding(Self.pendingInset)
                    }
                }
                .contentShape(.rect)
                .parityItem(.button)
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(on ? .isSelected : [])
        .accessibilityValue(on && busy ? Copy.raidsA11ySaving : state == pending ? Copy.raidsA11yPending : "")
    }
}

/// True while the page is offline: an answer needs a connection, so the answers read as disabled.
private struct AnswersOfflineKey: EnvironmentKey {
    static let defaultValue = false
}

extension EnvironmentValues {
    var answersOffline: Bool {
        get { self[AnswersOfflineKey.self] }
        set { self[AnswersOfflineKey.self] = newValue }
    }
}


/// How late? 15, 30 or 60 min, under the answers after Late (F1); none chosen keeps Late. The question
/// sits `questionTop` under the answers, the 36 pt chips straight under it (the card's `.me` and `.chips`).
struct LateMinutesChips: View {
    static let minutes = [15, 30, 60]
    var questionTop = FSSpace.space2
    let late: (Int) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(Copy.pushLateQuestion).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, questionTop)
            LateMinutesRow(height: FSButtonStyle.smallHeight, late: late)
        }
    }
}

/// The three minute chips alone (the Raid tab sets them under the answers without the question, 44
/// tall); `chosen` marks the minutes already said, in teal.
struct LateMinutesRow: View {
    var chosen: Int?
    var height = FSTarget.min
    let late: (Int) -> Void

    var body: some View {
        HStack(spacing: FSSpace.space2) {
            ForEach(LateMinutesChips.minutes, id: \.self) { minutes in
                Button(Copy.pushLateMinutes(n: minutes)) { late(minutes) }
                    .buttonStyle(MinuteChipStyle(on: minutes == chosen, height: height))
                    .accessibilityAddTraits(minutes == chosen ? .isSelected : [])
            }
        }
    }
}

/// A choice chip (round 5): the rule's outline, 8 radius, the chosen one filled in the teal edge.
/// `fillsWidth`: the minute chips share their row; a filter chip (Not answered) hugs its label.
struct MinuteChipStyle: ButtonStyle {
    let on: Bool
    let height: CGFloat
    var fillsWidth = true

    static let radius: CGFloat = 8

    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .type(FSType.chip)
            .foregroundStyle(FSColor.strong)
            .padding(.horizontal, FSSpace.space4)
            .frame(maxWidth: fillsWidth ? .infinity : nil, minHeight: height)
            .background(on || configuration.isPressed ? FSColor.tealEdge : .clear, in: .rect(cornerRadius: Self.radius))
            .overlay(RoundedRectangle(cornerRadius: Self.radius).strokeBorder(FSColor.rule, lineWidth: FSStroke.hairline))
            .contentShape(.rect)
            .parityItem(.button)
    }
}
