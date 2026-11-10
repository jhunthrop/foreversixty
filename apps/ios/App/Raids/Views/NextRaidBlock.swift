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
            Text(content.answer).type(FSType.small).foregroundStyle(FSColor.text).padding(.top, FSSpace.space3)
            RaidAnswerControl(content: content, actions: actions, question: true, compact: true).padding(.top, FSSpace.space2)
            if !content.tally.isEmpty { FSTally(content.tally).padding(.top, Self.tallyTop) }
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
    @ScaledMetric(relativeTo: .largeTitle) private var valueLine: CGFloat = 38

    var body: some View {
        if !parts.isEmpty {
            HStack(alignment: .top, spacing: Self.gap) {
                ForEach(Array(parts.enumerated()), id: \.offset) { _, part in
                    VStack(alignment: .leading, spacing: 0) {
                        Text(part.value)
                            .type(TypeRole(face: FSFont.bodySemibold, size: 36, lineHeight: 38, scalesWith: .largeTitle, tabular: true))
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
            // The role tally, the Raid tab's on one line (`2/2 · 5/8 · 11/30`, a short role in wipe red).
            if !content.tally.isEmpty { FSTally(content.tally) }
            RaidAnswerControl(content: content, actions: actions, question: true, compact: true)
                .padding(.bottom, FSSpace.space1)
        }
        .parityRegion(.countdowns)
    }
}

/// The answer as Home's card and the Raid tab's head draw it (F1, brand B1): In, Late, Bench, Out as
/// one segmented group; after Late its minutes, the chosen one selected; the status line (`Saved.`,
/// the refusal with Try again, offline's sentence). Offline the answers disable.
struct RaidAnswerControl: View {
    let content: NextRaidContent
    let actions: NextRaidActions
    /// Home's card asks `How late?` over the minutes; the Raid tab's head sets them bare.
    var question = false
    /// Home's raid line: the 36 pt answers.
    var compact = false

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            FSSegmented(
                Copy.homeRaidAnswers, options: SignupState.allCases.map { FSSegmented.Option($0, $0.label) },
                selection: content.selected, compact: compact, select: actions.answer
            )
            .disabled(!content.answersEnabled)
            .accessibilityValue(content.busy ? Copy.raidsA11ySaving : "")
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

/// In, Late, Bench, Out: four 44 pt buttons 8 apart (iOS, m-ios-home), the current answer in teal (a
/// change of plan is one tap; flows spec appendix C, "the answered state without answers"). `small`:
/// the 36 pt chips of the raid night page, not yet restyled.
struct RaidAnswerButtons: View {
    let selected: SignupState?
    var small = false
    let answer: (SignupState) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        // At the accessibility sizes the answers go two by two, so no word breaks.
        let columns = typeSize.isAccessibilitySize ? 2 : SignupState.allCases.count
        LazyVGrid(columns: Array(repeating: GridItem(.flexible(), spacing: FSSpace.space2), count: columns), spacing: FSSpace.space2) {
            ForEach(SignupState.allCases) { state in
                Button(state.label) { answer(state) }
                    .buttonStyle(FSButtonStyle(state == selected ? .primary : .secondary, small: small, padX: FSButtonStyle.padXTight))
                    .lineLimit(1)
                    .accessibilityAddTraits(state == selected ? .isSelected : [])
            }
        }
        .accessibilityElement(children: .contain)
        .accessibilityLabel(Copy.homeRaidAnswers)
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
struct MinuteChipStyle: ButtonStyle {
    let on: Bool
    let height: CGFloat

    static let radius: CGFloat = 8

    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(.custom(FSFont.bodySemibold, size: 14, relativeTo: .subheadline))
            .foregroundStyle(FSColor.strong)
            .padding(.horizontal, FSSpace.space4)
            .frame(maxWidth: .infinity, minHeight: height)
            .background(on || configuration.isPressed ? FSColor.tealEdge : .clear, in: .rect(cornerRadius: Self.radius))
            .overlay(RoundedRectangle(cornerRadius: Self.radius).strokeBorder(FSColor.rule, lineWidth: FSStroke.hairline))
            .contentShape(.rect)
            .parityItem(.button)
    }
}
