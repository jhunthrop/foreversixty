import DesignSystem
import Models
import SwiftUI
import UIKit

/// A pushed raid screen while it loads or after it failed.
struct RaidFlowLoading: View {
    let failed: Bool
    let message: String
    let retry: () -> Void

    var body: some View {
        RaidFlowScroll {
            if failed {
                ErrorRegion(message: message, retry: retry).padding(.top, RaidFlowSpacing.section)
            } else {
                ProgressView().frame(maxWidth: .infinity).padding(.top, RaidFlowSpacing.section)
            }
        }
    }
}

/// Not answered on the Raid stack (F3), opened from Preview mention or the `raid.short` push.
struct NotAnsweredScreen: View {
    @State private var model: NotAnsweredModel
    let guildName: String
    let site: any RaidSiteActions
    let clock: RaidClock
    @Environment(\.raidIconBuild) private var build
    @Environment(\.openURL) private var openURL

    init(model: @autoclosure () -> NotAnsweredModel, guildName: String, site: any RaidSiteActions, clock: RaidClock) {
        _model = State(wrappedValue: model())
        self.guildName = guildName
        self.site = site
        self.clock = clock
    }

    var body: some View {
        Group {
            if let raid = model.raid {
                let presenter = NotAnsweredPresenter(raid: raid, guild: guildName, marks: model.marks, clock: clock)
                let content = presenter.content(filter: model.filter)
                NotAnsweredView(
                    content: content, build: build ?? "", status: model.mentionStatus, mentioning: model.mentioning,
                    actions: actions(raid, short: presenter.shortRole)
                )
                    .overlay(alignment: .bottom) {
                        if let error = model.error {
                            SmallText(error, color: FSColor.wipe).padding(FSSpace.space3).background(FSColor.raised)
                        }
                    }
            } else {
                RaidFlowLoading(failed: model.phase == .failed, message: Copy.raidFailed) { Task { await model.load() } }
            }
        }
        .pushedBandBar(model.raid?.zoneName ?? "")
        .toolbar {
            ToolbarItem(placement: .topBarTrailing) {
                ShareLink(item: RaidLinks.web(guild: model.guild, raid: model.raidID)) {
                    Label(Copy.raidMenuShare, systemImage: "square.and.arrow.up")
                }
            }
        }
        .task { await model.load() }
        .refreshable { await model.load() }
    }

    /// With no Discord bot, the raid's page on the site connects one.
    private func mention(role: String?, raid: Int64) {
        Task { if await model.mention(role: role) { openURL(site.raid(guild: model.guild, raid: raid)) } }
    }

    private func actions(_ raid: RaidView, short: RaidShortRole?) -> NotAnsweredActions {
        NotAnsweredActions(
            select: { model.filter = $0 },
            mentionRole: { mention(role: short?.role, raid: raid.id) },
            mentionAll: { mention(role: nil, raid: raid.id) },
            mark: { row, state in model.mark(row.characterKey, state) },
            undo: { model.undo($0.characterKey) },
            message: { openURL($0) },
            editOnSite: { openURL(RaidLinks.web(guild: model.guild, raid: raid.id)) }
        )
    }
}

/// The applicant on the Raid stack (F3), opened from Needs you or the applicant push's Review.
struct ApplicantScreen: View {
    @State private var model: ApplicantModel
    let guildName: String
    let now: () -> Date
    let back: () -> Void

    init(model: @autoclosure () -> ApplicantModel, guildName: String, now: @escaping () -> Date, back: @escaping () -> Void) {
        _model = State(wrappedValue: model())
        self.guildName = guildName
        self.now = now
        self.back = back
    }

    var body: some View {
        Group {
            if case let .loaded(application) = model.phase {
                let content = RaidFlowPresenter(now: now()).applicant(application, guild: guildName, accepted: model.accepted)
                ApplicantView(
                    content: content, busy: model.busy, error: model.error,
                    actions: ApplicantActions(
                        accept: { Task { await model.accept() } },
                        decline: { Task { if await model.decline() { back() } } },
                        copy: { UIPasteboard.general.string = $0 }
                    )
                )
            } else {
                RaidFlowLoading(failed: model.phase == .failed, message: Copy.commonFailed) { Task { await model.load() } }
            }
        }
        .pushedBandBar(applicantName)
        .task { await model.load() }
    }

    private var applicantName: String {
        if case let .loaded(application) = model.phase { return application.character.name }
        return ""
    }
}

/// The raid page on raid night on the Raid stack (F2), opened from Open raid night, the ready
/// line or the `raid.starting` push. Once the night is over, Your raid night answers instead.
struct RaidNightPageScreen: View {
    @State private var model: RaidNightPageModel
    let clock: RaidClock
    let now: () -> Date
    let openRoster: () -> Void
    let after: () -> AnyView
    @Environment(\.raidIconBuild) private var build
    @Environment(\.openURL) private var openURL

    static let afterPhase = "after"

    init(
        model: @autoclosure () -> RaidNightPageModel, clock: RaidClock, now: @escaping () -> Date,
        openRoster: @escaping () -> Void, after: @escaping () -> AnyView
    ) {
        _model = State(wrappedValue: model())
        self.clock = clock
        self.now = now
        self.openRoster = openRoster
        self.after = after
    }

    var body: some View {
        Group {
            if case let .loaded(night) = model.phase {
                if night.phase == Self.afterPhase {
                    after()
                } else {
                    page(night)
                }
            } else {
                RaidFlowLoading(failed: model.phase == .failed, message: Copy.nightFailed) { Task { await model.load() } }
            }
        }
        .task { await model.load() }
    }

    private func page(_ night: RaidNightSummary) -> some View {
        let presenter = RaidFlowPresenter(now: now(), clock: clock)
        let content = presenter.nightPage(night, raid: model.raid, lateOpen: model.answers.lateOpen, answer: model.answers.pending)
        return RaidNightPageView(
            content: content, build: build ?? "", error: model.answers.error,
            actions: RaidNightPageActions(
                answer: { state in Task { await model.answer(state) } },
                late: { minutes in Task { await model.late(minutes) } },
                roster: openRoster,
                openLink: { openURL($0) }
            )
        )
        .refreshable { await model.load() }
        .pushedBandBar(night.zoneName)
        .toolbar {
            ToolbarItem(placement: .topBarTrailing) {
                ShareLink(item: RaidLinks.web(guild: night.guildId, raid: model.raidID)) {
                    Label(Copy.raidMenuShare, systemImage: "square.and.arrow.up")
                }
            }
        }
    }
}

/// The guild's planned raids (the Calendar row), each opening its raid page.
struct RaidCalendarScreen: View {
    let planned: [RaidView]
    let clock: RaidClock
    let now: () -> Date
    let currentKey: String?
    let open: (RaidView) -> Void

    var body: some View {
        RaidFlowScroll {
            RaidFlowHead(title: Copy.raidPagesCalendar)
            let upcoming = planned.filter { $0.startsAt > now().addingTimeInterval(-HomeFlowModel.liveSpan) }.sorted { $0.startsAt < $1.startsAt }
            if upcoming.isEmpty { SmallText(Copy.raidNoRaid).padding(.top, RaidFlowSpacing.section) }
            VStack(spacing: 0) {
                ForEach(upcoming, id: \.id) { raid in
                    PlannedRow(raid: raid, clock: clock, now: now(), currentKey: currentKey) { open(raid) }
                    Hairline()
                }
            }
            .padding(.top, RaidFlowSpacing.line)
        }
        .pushedBandBar(Copy.raidPagesCalendar)
    }
}
