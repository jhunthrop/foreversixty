import DesignSystem
import Models
import Routing
import SwiftUI
import UIKit

/// F4.2 the raid page: loads, then draws `RaidPageContent`.
struct RaidPageScreen: View {
    @State private var model: RaidPageModel
    let clock: RaidClock
    let now: () -> Date
    let currentKey: String?
    let guildName: String
    let guildURL: URL?
    let openNight: (Int64) -> Void
    let back: () -> Void
    let scrollToRoster: Bool

    /// The model is made once per pushed page; a redraw of the stack keeps it.
    init(
        model: @autoclosure () -> RaidPageModel, clock: RaidClock, now: @escaping () -> Date, currentKey: String?,
        guildName: String, guildURL: URL?, openNight: @escaping (Int64) -> Void, back: @escaping () -> Void,
        scrollToRoster: Bool
    ) {
        _model = State(wrappedValue: model())
        self.clock = clock
        self.now = now
        self.currentKey = currentKey
        self.guildName = guildName
        self.guildURL = guildURL
        self.openNight = openNight
        self.back = back
        self.scrollToRoster = scrollToRoster
    }

    var body: some View {
        Group {
            switch model.state {
            case .loading:
                ProgressView().frame(maxWidth: .infinity, maxHeight: .infinity)
            case let .loaded(controller):
                RaidPageContent(
                    controller: controller, clock: clock, now: now, currentKey: currentKey,
                    offlineSince: model.offlineSince, guildName: guildName, guildURL: guildURL, openNight: openNight, scrollToRoster: scrollToRoster
                )
                .refreshable { await model.load() }
            case .failed:
                ScrollView { ErrorRegion(message: Copy.raidFailed) { Task { await model.load() } }.padding(16) }
            case .notAvailable:
                ScrollView {
                    VStack(alignment: .leading, spacing: 12) {
                        Text(Copy.raidNotAvailable).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.text)
                        Button(Copy.raidBack, action: back).buttonStyle(SecondaryButtonStyle())
                    }
                    .quietCard(padding: 16)
                    .padding(16)
                }
            }
        }
        .background(Theme.Palette.background)
        .task { await model.load() }
        .task { await model.stayLive() }
    }
}

/// The raid page's content, pure over the controller so snapshot tests draw every state.
struct RaidPageContent: View {
    let controller: SignupController
    let clock: RaidClock
    let now: () -> Date
    let currentKey: String?
    let offlineSince: Date?
    let guildName: String
    var guildURL: URL?
    let openNight: (Int64) -> Void
    var scrollToRoster = false
    @State private var showsCalendar = false
    @Environment(\.openURL) private var openURL

    private var raid: RaidView { controller.raid }

    var body: some View {
        let facts = controller.facts
        ScrollViewReader { proxy in
            ScrollView {
                VStack(alignment: .leading, spacing: 0) {
                    if let offlineSince {
                        OfflineStrip(text: Copy.offlineStrip(time: clock.stripTime(offlineSince, now: now())))
                    }
                    VStack(alignment: .leading, spacing: 22) {
                        RaidHead(raid: raid, facts: facts, clock: clock, now: now, guildName: guildName)
                        CountBlock(facts: facts, muted: facts.chip == .cancelled)
                        if facts.showsOfficerTools, facts.isSignable {
                            OfficerRow(controller: controller, clock: clock)
                        }
                        SignupCard(controller: controller, facts: facts, clock: clock, currentKey: currentKey)
                        if raid.viewer.raidNight == true { nightRow }
                        if let report = raid.match?.reportId, facts.chip == .done { reportRow(report) }
                        ClassStripView(strip: ClassStrip(raid: raid)) { withAnimation { proxy.scrollTo("roster", anchor: .top) } }
                        notes
                        RosterView(controller: controller, facts: facts).id("roster")
                    }
                    .padding(16)
                }
            }
            .onAppear { if scrollToRoster { proxy.scrollTo("roster", anchor: .top) } }
        }
        .background(Theme.Palette.background)
        .overlay(alignment: .bottom) { toast }
        .navigationTitle(raid.title)
        .navigationBarTitleDisplayMode(.inline)
        .toolbar { toolbar }
        .sheet(isPresented: $showsCalendar) {
            CalendarSheet(event: CalendarEvent(raid: raid, guildName: guildName, guild: controller.guild))
        }
        .sheet(item: Binding(get: { controller.specAsk }, set: { if $0 == nil { controller.closeSpecSheet() } })) { ask in
            if let character = controller.character(ask.characterKey) {
                SpecSheet(ask: ask, character: character, raidTitle: raid.title, date: clock.shortDate(raid.startsAt)) { option in
                    Task { await controller.pickSpec(option) }
                }
            }
        }
    }

    @ToolbarContentBuilder
    private var toolbar: some ToolbarContent {
        ToolbarItemGroup(placement: .topBarTrailing) {
            Button { showsCalendar = true } label: { Image(systemName: "calendar.badge.plus") }
                .accessibilityLabel(Copy.raidMenuCalendar)
            Menu {
                Button { showsCalendar = true } label: { Label(Copy.raidMenuCalendar, systemImage: "calendar.badge.plus") }
                ShareLink(item: shareURL) { Label(Copy.raidMenuShare, systemImage: "square.and.arrow.up") }
                Button { openURL(webURL) } label: { Label(Copy.raidMenuWeb, systemImage: "globe") }
            } label: {
                Image(systemName: "ellipsis")
            }
        }
    }

    private var shareURL: URL { RaidLinks.web(guild: controller.guild, raid: raid.id) }

    /// The site's raid page; the app link when the guild's path is unknown.
    private var webURL: URL { guildURL.map { $0.appending(path: "raids/\(raid.id)") } ?? shareURL }

    private var nightRow: some View {
        Button { openNight(raid.id) } label: {
            HStack {
                Text(Copy.nightTitle).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                Spacer()
                Image(systemName: "chevron.right").foregroundStyle(Theme.Palette.muted).accessibilityHidden(true)
            }
            .padding(.horizontal, 16)
            .frame(minHeight: 56)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
        .quietCard()
    }

    private func reportRow(_ report: String) -> some View {
        Button(Copy.raidReport) { openURL(RaidLinks.report(report)) }.buttonStyle(GoldLinkStyle())
    }

    @ViewBuilder
    private var notes: some View {
        VStack(alignment: .leading, spacing: 10) {
            Text(Copy.raidNotes.uppercased()).font(Theme.Typography.h2).foregroundStyle(Theme.Palette.strong)
                .accessibilityAddTraits(.isHeader)
            if raid.notesHtml.trimmingCharacters(in: .whitespacesAndNewlines).isEmpty {
                Text(Copy.raidNotesEmpty).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.muted)
            } else {
                Text(NotesText.attributed(raid.notesHtml, base: Theme.Typography.bodyText, bold: Theme.Typography.bodyStrong, link: Theme.Palette.accent))
                    .foregroundStyle(Theme.Palette.text)
                    .lineSpacing(7)
                    .tint(Theme.Palette.accent)
            }
        }
    }

    @ViewBuilder
    private var toast: some View {
        if let text = controller.toast {
            HStack(spacing: 10) {
                Image(systemName: "lock").foregroundStyle(Theme.Palette.kill).accessibilityHidden(true)
                Text(text).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.strong)
            }
            .padding(16)
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(Theme.Palette.sheet, in: .rect(cornerRadius: Theme.Radius.panel))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.panel).stroke(Theme.Palette.line))
            .padding(16)
            .transition(.move(edge: .bottom).combined(with: .opacity))
            .onAppear { AccessibilityNotification.Announcement(text).post() }
        }
    }
}

/// The page head (spec F4.2 items 1 to 3): eyebrow, chip, guild, title, when lines, countdown.
struct RaidHead: View {
    let raid: RaidView
    let facts: RaidFacts
    let clock: RaidClock
    let now: () -> Date
    var guildName: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            HStack(spacing: 8) {
                Text(Copy.raidEyebrow.uppercased()).font(Theme.Typography.micro).tracking(1.5).foregroundStyle(Theme.Palette.accent)
                StatusTag(chip: facts.chip)
                if let guildName {
                    Text(guildName).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
                }
            }
            HStack(alignment: .firstTextBaseline, spacing: 0) {
                Text(raid.title.uppercased()).font(Theme.Typography.pageTitle).foregroundStyle(facts.chip == .cancelled ? Theme.Palette.muted : Theme.Palette.strong)
                    .accessibilityAddTraits(.isHeader)
                if raid.title != raid.zoneName {
                    Text(Copy.raidsZoneSuffix(zone: raid.zoneName)).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.muted)
                }
            }
            Text(clock.longDate(raid.startsAt)).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.text)
            TimesLine(raid: raid, clock: clock, showsEnd: true)
            CountdownLine(raid: raid, facts: facts, clock: clock, now: now)
            if let zone = TimeZone(identifier: raid.guildTimezone), clock.offsetDiffers(at: raid.startsAt, guildZone: zone) {
                Text(Copy.raidGuildTime(clock: clock.guildClock(raid.startsAt, guildZone: zone)))
                    .font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
            }
        }
    }
}

/// `Invites 7:45 pm · Pull 8:00 pm · Ends 11:00 pm`, labels muted and times in mono.
struct TimesLine: View {
    let raid: RaidView
    let clock: RaidClock
    let showsEnd: Bool

    var body: some View {
        let end = raid.startsAt.addingTimeInterval(TimeInterval(raid.durationMinutes * 60))
        var parts: [(String, String)] = []
        if let invite = raid.inviteAt { parts.append((Copy.raidInvites(clock: "{}"), clock.clock(invite))) }
        parts.append((Copy.raidPull(clock: "{}"), clock.clock(raid.startsAt)))
        if showsEnd || raid.inviteAt == nil { parts.append((Copy.raidEnds(clock: "{}"), clock.clock(end))) }
        return parts.enumerated().reduce(Text("")) { line, item in
            let (label, time) = item.element
            let pieces = label.components(separatedBy: "{}")
            let separator = item.offset == 0 ? Text("") : Text(" · ").foregroundStyle(Theme.Palette.muted)
            let head = Text(pieces.first ?? "").foregroundStyle(Theme.Palette.muted)
            let figure = Text(time).font(Theme.Typography.mono).foregroundStyle(Theme.Palette.strong)
            let tail = Text(pieces.dropFirst().joined()).foregroundStyle(Theme.Palette.muted)
            return Text("\(line)\(separator)\(head)\(figure)\(tail)")
        }
        .font(Theme.Typography.bodyText)
        .fixedSize(horizontal: false, vertical: true)
    }
}

/// `Invites in 2 days 2 h` (or `Starts …`), the countdown in mono gold, ticking each minute.
struct CountdownLine: View {
    let raid: RaidView
    let facts: RaidFacts
    let clock: RaidClock
    let now: () -> Date

    var body: some View {
        TimelineView(.everyMinute) { _ in
            let current = now()
            let target = facts.countdownTarget
            let countdown = clock.countdown(to: target, end: facts.endsAt, now: current)
            let template = raid.inviteAt != nil ? Copy.raidCountdownInvites(countdown: "{}") : Copy.raidCountdownStarts(countdown: "{}")
            let pieces = template.components(separatedBy: "{}")
            let live = facts.chip == .planned || facts.chip == .locked
            let running = live && target > current
            let head = Text(pieces.first ?? "").foregroundStyle(Theme.Palette.text)
            let figure = Text(countdown).font(Theme.Typography.mono).foregroundStyle(running ? Theme.Palette.accent : Theme.Palette.muted)
            let tail = Text(pieces.dropFirst().joined()).foregroundStyle(Theme.Palette.text)
            Text("\(head)\(figure)\(tail)")
                .font(Theme.Typography.bodyText)
                .accessibilityLabel(Copy.raidA11yCountdown(span: RaidClock.spanWords(target.timeIntervalSince(current))))
        }
    }
}

/// The sign-up card (spec F4.2 item 5, F4.9 cancelled).
struct SignupCard: View {
    let controller: SignupController
    let facts: RaidFacts
    let clock: RaidClock
    let currentKey: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(Copy.raidYourSignup.uppercased()).font(Theme.Typography.sec).foregroundStyle(Theme.Palette.strong)
                .accessibilityAddTraits(.isHeader)
            if facts.chip == .cancelled {
                cancelled
            } else if facts.chip == .done {
                EmptyView()
            } else {
                ForEach(facts.viewerCharacters(currentKey: currentKey), id: \.characterKey) { character in
                    SignupBlockView(controller: controller, character: character)
                }
                if controller.showsReminderAsk {
                    ReminderAskView { turnOn in Task { await controller.answerReminderAsk(turnOn: turnOn) } }
                }
            }
        }
        .primaryCard()
        .animation(.easeOut(duration: 0.2), value: controller.showsReminderAsk)
    }

    @ViewBuilder
    private var cancelled: some View {
        Text(Copy.raidCancelled).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.text)
        if let own = facts.viewerCharacters(currentKey: currentKey).first(where: { $0.state != nil }),
           let state = own.state.flatMap(SignupState.init) {
            Text(Copy.raidCancelledYou(state: state.label, role: RaidRole.label(own.role)))
                .font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
        }
    }
}

/// The classes strip (spec F4.2 item 6): nine crests with In plus Late under each.
struct ClassStripView: View {
    let strip: ClassStrip
    let tap: () -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(alignment: .leading, spacing: 10) {
            SectionLabel(text: Copy.raidClasses)
            let large = typeSize >= .accessibility1
            let columns = Array(repeating: GridItem(.flexible(), spacing: 4), count: large ? 5 : 9)
            Button(action: tap) {
                LazyVGrid(columns: columns, spacing: 10) {
                    ForEach(strip.entries) { entry in
                        VStack(spacing: 4) {
                            CrestView(className: entry.gameClass.rawValue, size: large ? 36 : 30)
                            Text("\(entry.count)").font(Theme.Typography.monoMeta).foregroundStyle(Theme.Palette.strong)
                        }
                        .opacity(entry.count == 0 ? 0.4 : 1)
                    }
                }
            }
            .buttonStyle(.plain)
            .accessibilityElement(children: .ignore)
            .accessibilityLabel(strip.accessibilityLabel)
        }
    }
}

/// The links the raid page shares and opens.
enum RaidLinks {
    /// The raid's https app link (links.json); the site redirects it to its own page without the app.
    static func web(guild: Int64, raid: Int64) -> URL {
        DeepLink.raid(guild: guild, raid: raid).url
    }

    static func report(_ id: String) -> URL {
        StaticSite.url.appending(path: "reports/\(id)")
    }
}
