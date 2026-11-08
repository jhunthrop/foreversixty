import DesignSystem
import SwiftUI

/// What the next raid's card and line hand back.
struct NextRaidActions {
    var answer: (SignupState) -> Void = { _ in }
    var late: (Int) -> Void = { _ in }
    var openRaid: () -> Void = {}
    var openReady: () -> Void = {}
    var previewMention: () -> Void = {}
    var openApplicants: () -> Void = {}
}

/// The next raid on raid day (flows spec F1, mocks ios-home-raid-day, -late, -officer and the Raid
/// tab): the countdown, who you answer as, the four answers with the current one selected, Late's
/// minutes, who is coming, what is not ready, and an officer's open items.
struct NextRaidCardView: View {
    let content: NextRaidContent
    var actions = NextRaidActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            FSCountdown(eyebrow: content.eyebrow, clock: content.clock) { SmallText(content.when) }
            SmallText(content.answer).padding(.top, FSSpace.space2)
            RaidAnswerButtons(selected: content.selected, answer: actions.answer).padding(.top, FSSpace.space3)
            if content.lateOpen { LateMinutesChips(late: actions.late) }
            if !content.tally.isEmpty { FSTally(content.tally).padding(.top, FSSpace.space4) }
            rows.padding(.top, FSSpace.space2)
        }
        .parityRegion(.countdowns)
    }

    @ViewBuilder
    private var rows: some View {
        VStack(spacing: 0) {
            if let ready = content.ready {
                Button(action: actions.openReady) {
                    FSRow(title: ready.title, sub: ready.sub, density: .slim, inset: 0) { EmptyView() }
                }
                .buttonStyle(RowPressStyle())
            }
            if content.officer != nil {
                // The officer's items are a list of their own (`.blk .rows { margin-top: 8px }`).
                Color.clear.frame(height: FSSpace.space2)
            }
            if let unanswered = content.officer?.unanswered {
                FSRow(
                    title: unanswered.title, sub: unanswered.sub.isEmpty ? nil : unanswered.sub, inset: 0, trailPlacement: .below
                ) {
                    EmptyView()
                } trail: {
                    Button(Copy.homeOfficerPreviewMention, action: actions.previewMention)
                        .buttonStyle(FSButtonStyle(fullWidth: false, small: true))
                }
            }
            if let applicants = content.officer?.applicants {
                Button(action: actions.openApplicants) {
                    FSRow(title: applicants.title, sub: applicants.sub.isEmpty ? nil : applicants.sub, inset: 0) { EmptyView() }
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// The next raid on another day (mocks ios-home-60, -answered): one line with its answers.
struct NextRaidLineView: View {
    let content: NextRaidContent
    let build: String
    var actions = NextRaidActions()

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            Button(action: actions.openRaid) {
                FSRow(title: content.line.title, sub: content.line.sub, inset: 0) {
                    ItemIconView(icon: content.icon, quality: nil, build: build, size: FSSize.icon40).parityItem(.image)
                }
            }
            .buttonStyle(RowPressStyle())
            RaidAnswerButtons(selected: content.selected, answer: actions.answer).padding(.top, FSSpace.space3)
            if content.lateOpen { LateMinutesChips(late: actions.late) }
        }
        .parityRegion(.countdowns)
    }
}

/// In, Late, Bench, Out in four equal columns, the current answer selected (a change of plan is one
/// tap; flows spec appendix C, "the answered state without answers").
struct RaidAnswerButtons: View {
    let selected: SignupState?
    let answer: (SignupState) -> Void

    var body: some View {
        HStack(spacing: FSSpace.space2) {
            ForEach(SignupState.allCases) { state in
                Button(state.label) { answer(state) }
                    .buttonStyle(FSButtonStyle(state == selected ? .primary : .secondary, small: true))
                    .accessibilityAddTraits(state == selected ? .isSelected : [])
            }
        }
        .accessibilityElement(children: .contain)
        .accessibilityLabel(Copy.homeRaidAnswers)
    }
}

/// How late? 15, 30 or 60 min, under the answers after Late (F1); none chosen keeps Late.
struct LateMinutesChips: View {
    static let minutes = [15, 30, 60]
    let late: (Int) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            SmallText(Copy.pushLateQuestion).padding(.top, FSSpace.space2)
            HStack(spacing: FSSpace.space2) {
                ForEach(Self.minutes, id: \.self) { minutes in
                    Button(Copy.pushLateMinutes(n: minutes)) { late(minutes) }
                        .buttonStyle(FSButtonStyle(small: true))
                }
            }
        }
    }
}
