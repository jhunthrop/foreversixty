import DesignSystem
import SwiftUI

/// What the raid page on raid night hands back.
struct RaidNightPageActions {
    var answer: (SignupState) -> Void = { _ in }
    var late: (Int) -> Void = { _ in }
    var roster: () -> Void = {}
    var openLink: (URL) -> Void = { _ in }
}

/// The raid page on raid night (flows spec F2, mocks ios-raid-night-live, -signup): the countdown
/// with your answer, Your sign-up with the current answer selected, who is coming by role, and
/// Before you pull, the private checklist from the addon's last export.
struct RaidNightPageView: View {
    let content: RaidNightPageContent
    let build: String
    var error: String?
    var actions = RaidNightPageActions()

    /// The raid page's head sits a further 16 in (the mocks' `.ph` on the raid page).
    static let headInset: CGFloat = FSSpace.space4
    static let tagTop: CGFloat = 6

    var body: some View {
        RaidFlowScroll {
            RaidFlowHead(title: content.title, facts: content.facts, inset: Self.headInset) {
                ItemIconView(icon: content.icon, quality: nil, build: build, size: FSSize.crest56).parityItem(.image)
            }
            FSCountdown(eyebrow: content.eyebrow, clock: content.clock, live: content.clock.isEmpty) {
                VStack(alignment: .leading, spacing: Self.tagTop) {
                    SmallText(content.aside)
                    FSTag(content.tag, tone: content.answered ? .kill : .muted)
                }
            }
            FSSection(Copy.raidYourSignup) {
                VStack(alignment: .leading, spacing: 0) {
                    RaidAnswerButtons(selected: content.selected, answer: actions.answer)
                    if content.lateOpen { LateMinutesChips(late: actions.late) }
                    if let error { SmallText(error, color: FSColor.wipe).padding(.top, RaidFlowSpacing.line) }
                }
                .padding(.top, RaidFlowSpacing.buttons)
            }
            .padding(.top, RaidFlowSpacing.section)
            FSSection(Copy.raidsWhoIsComing) {
                SectionLink(title: Copy.raidNightRoster, action: actions.roster)
            } content: {
                RaidRoleRows(roles: content.roles)
            }
            .padding(.top, RaidFlowSpacing.section)
            if let ready = content.ready {
                FSSection(Copy.raidNightBeforePull) { ReadyChecklistView(content: ready, open: actions.openLink) }
                    .padding(.top, RaidFlowSpacing.section)
            }
        }
    }
}
