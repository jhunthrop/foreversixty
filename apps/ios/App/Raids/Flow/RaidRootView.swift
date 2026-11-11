import DesignSystem
import SwiftUI

/// What the Raid tab's root hands back.
struct RaidRootActions {
    var raid = NextRaidActions()
    var openPage: (GuildPage) -> Void = { _ in }
    /// The roster filtered to these characters (row 5's not-synced raiders).
    var openRoster: ([String]) -> Void = { _ in }
    /// The raid page's cover list, from Covered's `and N more` (round 5, row 4).
    var openCover: () -> Void = {}
    var openLink: (URL) -> Void = { _ in }
    var apply: (RecruitingRow) -> Void = { _ in }
    var claim: () -> Void = {}
    var findGuild: () -> Void = {}
    var signIn: () -> Void = {}
    var retry: () -> Void = {}
    /// The Guild tab's own sections (spec 2026-11-10-guild-tab.md).
    var guild = GuildTabActions()
}

/// The Guild tab (spec 2026-11-10-guild-tab.md, mocks ios-guild-tab*; the raid as flows spec 4.2 and
/// the mocks ios-raid, -signup, -officer, -live and -no-raid): under the band's large title, the raid
/// first (head with your answer, Your job, Needs you, the officer's rows and assignments, Who is
/// coming, Before you pull), then the guild (head, roster, loot, logs, progression, pages); under 60
/// the guild head leads. No guild: the guilds recruiting your spec, then find or claim yours.
struct RaidRootView: View {
    let content: RaidRootContent
    var actions = RaidRootActions()

    static let skeletons: [CGFloat] = [140, 56, 56, 56]

    var body: some View {
        ScrollView {
            BandPage {
                EmptyView()
            } content: {
                switch content {
                case .loading: loading
                case .failed: failed
                case .signedOut:
                    LonePanel(title: Copy.raidTabWithGuild, lines: []) {
                        FSButton(Copy.signinButton, kind: .primary, action: actions.signIn)
                    }
                case let .noGuild(recruiting): noGuild(recruiting)
                case .waiting:
                    PanelBlock { RaidNote(Copy.raidsWaiting) }.parityRegion(RaidRootView.nextRaidRegion)
                case let .guild(guild): RaidGuildView(content: guild, actions: actions)
                }
            }
        }
        .tabRootScroll()
        .screenGround()
    }

    /// The Raid tab's region for its head: the old Next raid section's name, so every state keeps one.
    static let nextRaidRegion = ParityRegionName.section(Copy.raidTabNextRaid)

    private var loading: some View {
        VStack(spacing: FSSpace.space3) {
            ForEach(Array(Self.skeletons.enumerated()), id: \.offset) { _, height in FSSkeleton(height: height) }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.homeAnswerLoading)
    }

    private var failed: some View {
        PanelBlock {
            PanelRow(Copy.raidsFailed, last: true) {
                EmptyView()
            } trail: {
                Button(Copy.commonTryAgain, action: actions.retry).buttonStyle(FSButtonStyle(fullWidth: false))
            }
        }
    }

    @ViewBuilder
    private func noGuild(_ recruiting: RecruitingContent) -> some View {
        if let raidsOpen = recruiting.raidsOpen {
            Text(raidsOpen).type(RowType.sub).foregroundStyle(FSColor.text).frame(maxWidth: .infinity, alignment: .leading)
        }
        if !recruiting.rows.isEmpty {
            PanelSection(recruiting.heading, frame: .plain) {
                ForEach(Array(recruiting.rows.enumerated()), id: \.element.id) { index, row in
                    PanelRow(row.name, sub: Text(row.line), last: index == recruiting.rows.count - 1) {
                        EmptyView()
                    } trail: {
                        if let applied = row.applied {
                            Text(applied).type(RowType.sub).foregroundStyle(FSColor.muted)
                        } else {
                            Button(Copy.raidTabApply) { actions.apply(row) }.buttonStyle(FSButtonStyle(fullWidth: false))
                        }
                    }
                }
            }
        }
        // Find your guild (spec 2.4): the site's guild finder, then Claim.
        PanelSection(Copy.guildTabFind) {
            Button(action: actions.findGuild) {
                PanelRow(Copy.guildTabSearch, chevron: true)
            }
            .buttonStyle(RowPressStyle())
            .accessibilityIdentifier("guild-search")
            Button(action: actions.claim) {
                PanelRow(Copy.raidTabClaim, last: true, chevron: true)
            }
            .buttonStyle(RowPressStyle())
            .accessibilityIdentifier("guild-claim")
        }
    }
}

/// The Raid tab in a guild: raid night or the next raid's head with Who is coming, the checklist,
/// Needs you, the guild's pages.
struct RaidGuildView: View {
    let content: RaidGuildContent
    let actions: RaidRootActions

    var body: some View {
        if let stamp = content.offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: stamp)) }
        if let stale = content.staleStamp { StaleStrip(time: stale, retry: actions.retry) }
        // Under 60 the guild leads (spec 2.1); at 60 the raid does (2.2).
        if let tab = content.tab, tab.leveler { GuildHeadBlock(head: tab.head, spread: tab.spread) }
        raid
        if let tab = content.tab {
            if !tab.leveler { GuildHeadBlock(head: tab.head, spread: tab.spread) }
            GuildRosterSection(content: tab, open: actions.guild.openRoster)
            GuildLootSection(night: tab.lootNight, loot: tab.loot, open: actions.guild.openLoot)
            GuildLogsSection(logs: tab.logs, actions: actions.guild)
            GuildProgressSection(progression: tab.progression, open: actions.guild.openProgression)
        }
        GuildPagesSection(content: content.pages, open: actions.openPage)
    }

    @ViewBuilder
    private var raid: some View {
        if let night = content.night {
            RaidNightBlock(content: night)
        } else if let next = content.next {
            RaidHeadView(content: next, notes: content.notes, actions: actions.raid).parityRegion(RaidRootView.nextRaidRegion)
                .environment(\.answersOffline, content.offlineStamp != nil)
            raidsOpen
            // G3: an officer's open items sit directly under the head, the officer's rows with them; then your job.
            needsYou
            officer
            YourJobSection(jobs: content.jobs)
            if !content.coming.isEmpty {
                // With the Guild tab's Roster section on the page, its link is the one way to the roster.
                PanelSection(Copy.raidTabComing, action: content.tab == nil ? (Copy.raidTabRosterLink, { actions.openPage(.roster) }) : nil) {
                    RoleRowsView(roles: content.coming)
                    ComingFoot(
                        answers: content.answers, cover: content.cover, notSynced: content.notSynced,
                        openUnsynced: actions.openRoster, openCover: actions.openCover
                    )
                }
            }
        } else {
            PanelBlock { RaidNote(Copy.raidTabNoRaid).padding(.bottom, FSSpace.space2) }.parityRegion(RaidRootView.nextRaidRegion)
            raidsOpen
        }
        if let ready = content.ready {
            // Once pulled the list is the Ready check, not what comes before the pull.
            PanelSection(content.night == nil ? Copy.raidTabBeforePull : Copy.raidTabReadyCheck) {
                ReadyPanelList(content: ready, open: actions.openLink)
            }
        }
        if content.next == nil {
            needsYou
            officer
        }
    }

    /// Under 60: one quiet line, never a block (coordinator ruling 2026-11-10).
    @ViewBuilder
    private var raidsOpen: some View {
        if let line = content.raidsOpen {
            Text(line).type(RowType.sub).foregroundStyle(FSColor.text)
                .frame(maxWidth: .infinity, alignment: .leading)
                .accessibilityIdentifier("guild-raids-open")
        }
    }

    /// An officer's shortcuts and the raid's assignments, read only (spec 2.3, 2.3a).
    @ViewBuilder
    private var officer: some View {
        if let tab = content.tab { OfficerSection(rows: tab.officer, open: actions.guild.officer) }
        if let assignments = content.assignments { AssignmentsSection(encounters: assignments, openPlanner: actions.guild.openPlanner) }
    }

    @ViewBuilder
    private var needsYou: some View {
        if let needs = content.needsYou {
            PanelSection(Copy.raidTabNeedsYou, meta: needs.meta) { OfficerItemsView(items: needs.items, actions: actions.raid) }
        }
    }
}

/// Who is coming's foot: the other answers' counts, what the cover misses (G4), an officer's
/// not-synced raiders by name (G1), which opens the roster filtered to them.
private struct ComingFoot: View {
    let answers: String?
    let cover: RaidCoverContent?
    let notSynced: NotSyncedContent?
    let openUnsynced: ([String]) -> Void
    let openCover: () -> Void

    var body: some View {
        if answers != nil || cover != nil || notSynced != nil {
            VStack(alignment: .leading, spacing: FSSpace.space2) {
                if let answers { Text(answers).type(RowType.sub).foregroundStyle(FSColor.muted) }
                if let cover {
                    if cover.more > 0 {
                        // The names it leaves out are on the raid page's cover list (round 5, row 4).
                        Button(action: openCover) { CoverLine(content: cover).frame(minHeight: FSTarget.min).contentShape(.rect) }
                            .buttonStyle(RowPressStyle())
                            .parityItem(.row)
                    } else {
                        CoverLine(content: cover)
                    }
                }
                if let notSynced {
                    Button { openUnsynced(notSynced.keys) } label: {
                        HStack(spacing: FSSpace.space1) {
                            Text(notSynced.line).type(RowType.sub).foregroundStyle(FSColor.text)
                                .frame(maxWidth: .infinity, alignment: .leading)
                            Image(systemName: "chevron.right").foregroundStyle(FSColor.muted).accessibilityHidden(true)
                        }
                        .frame(minHeight: FSTarget.min)
                        .contentShape(.rect)
                    }
                    .buttonStyle(RowPressStyle())
                    .parityItem(.row)
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.vertical, FSSpace.rowPadY)
            .overlay(alignment: .top) { Hairline(FSColor.line) }
        }
    }
}

/// The Covered line (row 4; round 4, row 1): `Missing` once in the muted label tone, each gap's spell
/// icon beside its name in wipe red (only the names, never a red wall), ` · ` between the names, then
/// `and N more` in the link gold (round 5, row 4: it opens the cover list). A name never breaks
/// inside, a name with no icon draws no frame, and every word sits on one centre line, so the one
/// icon never drops its name below the others.
private struct CoverLine: View {
    let content: RaidCoverContent
    @Environment(\.raidIconBuild) private var build

    var body: some View {
        WrapLayout(spacing: FSSpace.space1, lineSpacing: FSSpace.space1, centresLines: true) {
            if !content.missing.isEmpty {
                Text(Copy.raidTabCoverLabel).foregroundStyle(FSColor.muted)
            }
            gaps(content.missing + content.short)
            if content.more > 0 {
                Text(Copy.raidTabAndMore(names: "", n: content.more).trimmingCharacters(in: .whitespaces)).foregroundStyle(FSColor.gold)
            }
        }
        .type(RowType.sub)
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(content.line)
    }

    /// Each name with the separator before it (so a wrapped line starts on `· Might`, never a lone dot).
    private func gaps(_ gaps: [CoverGap]) -> some View {
        ForEach(Array(gaps.enumerated()), id: \.offset) { index, gap in
            HStack(spacing: 0) {
                if index > 0 {
                    Text(RaidFlowPresenter.separator.trimmingCharacters(in: .whitespaces)).foregroundStyle(FSColor.muted)
                        .padding(.trailing, FSSpace.space1)
                }
                CoverIcon(icon: gap.icon, build: build ?? "")
                Text("\(Text(gap.name).foregroundColor(FSColor.wipe))\(Text(count(gap)).foregroundColor(FSColor.muted))")
            }
            .lineLimit(1)
            .fixedSize()
        }
    }

    /// ` 1 of 2` after a short gap.
    private func count(_ gap: CoverGap) -> String {
        gap.have.map { " " + Copy.raidTabOfWant(have: $0, want: gap.need) } ?? ""
    }
}

/// A gap's spell icon, drawn only once there is art for it: no stem, or a stem with no file, draws
/// nothing and the name closes up (round 4, row 1), never an empty frame.
private struct CoverIcon: View {
    let icon: String?
    let build: String
    @Environment(\.itemIconSource) private var source
    @State private var loaded = false

    var body: some View {
        Group {
            if let stem, loaded || source.cached(icon: stem, build: build) != nil {
                ItemIconView(icon: stem, quality: nil, build: build, size: FSSize.icon20)
                    .parityItem(.image)
                    .padding(.trailing, FSSpace.space1)
            } else {
                // Holds the load's task without taking any room.
                Color.clear.frame(width: 0, height: 0)
            }
        }
        .task(id: stem) {
            guard let stem, !build.isEmpty else { return }
            loaded = await source.load(icon: stem, build: build) != nil
        }
    }

    private var stem: String? { icon.flatMap { $0.isEmpty ? nil : $0 } }
}

/// The Raid tab's head (spec 4.2, owner 2026-11-09): one compact line, the zone's icon, the raid, its
/// time and the pull, who the answer is for (E7); the four answers under it are the answer (AS5: no
/// chip repeats it), with Late the minutes inline, the chosen one selected (F1).
struct RaidHeadView: View {
    let content: NextRaidContent
    /// The officers' notes, two lines under the head.
    var notes: String?
    let actions: NextRaidActions

    static let notesLines = 2
    /// How far the head's line may shrink to stay on one line.
    static let lineScale: CGFloat = 0.8
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        PanelBlock {
            VStack(alignment: .leading, spacing: 0) {
                Button(action: actions.openRaid) { line }.buttonStyle(RowPressStyle())
                RaidAnswerControl(content: content, actions: actions).padding(.top, FSSpace.space3)
                if let notes {
                    Text((try? AttributedString(markdown: notes)) ?? AttributedString(notes))
                        .type(RowType.sub).foregroundStyle(FSColor.text)
                        .lineLimit(Self.notesLines)
                        .padding(.top, FSSpace.space4)
                }
            }
            .padding(.vertical, FSSpace.space2)
        }
    }

    private var headLineText: some View {
        Text(content.headLine).type(RowType.sub).foregroundStyle(FSColor.muted)
    }

    private var line: some View {
        HStack(spacing: FSSpace.space3) {
            ZoneIcon(icon: content.icon)
            VStack(alignment: .leading, spacing: 0) {
                Text(content.zone).type(FSType.item).foregroundStyle(FSColor.strong).accessibilityAddTraits(.isHeader)
                // One line, as the mock holds it (shrinking a little); at the accessibility sizes it wraps,
                // since the time carries meaning and never truncates (spec 3.2).
                if typeSize.isAccessibilitySize {
                    headLineText.fixedSize(horizontal: false, vertical: true)
                } else {
                    headLineText.lineLimit(1).minimumScaleFactor(Self.lineScale)
                }
                if let invites = content.invites { Text(invites).type(RowType.sub).foregroundStyle(FSColor.muted) }
                if let who = content.who { Text(who).type(RowType.sub).foregroundStyle(FSColor.muted) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .contentShape(.rect)
    }
}

/// A raid's zone icon, 40 in the rounded frame of a row's lead.
struct ZoneIcon: View {
    let icon: String?
    @Environment(\.raidIconBuild) private var build

    var body: some View {
        ItemIconView(icon: icon, quality: nil, build: build ?? "", size: FSSize.icon40)
            .clipShape(.rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline))
            .parityItem(.image)
    }
}

/// Who is coming's rows: the role, its count against the comp (green when met, red when short), the names.
struct RoleRowsView: View {
    let roles: [RoleLineContent]

    var body: some View {
        ForEach(Array(roles.enumerated()), id: \.element.id) { index, line in
            RoleRow(line: line, last: index == roles.count - 1)
        }
    }
}

/// One role (S6): `Healers 5 of 8 · 3 short` in tone over its names in class colour; Damage's split
/// on the title's line, `Damage 15 · 7 melee, 8 ranged` (round 4, row 5, the Android form).
private struct RoleRow: View {
    let line: RoleLineContent
    let last: Bool

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            title.type(FSType.item)
            names.type(RowType.sub)
        }
        .padding(.vertical, FSSpace.rowPadY)
        .frame(maxWidth: .infinity, alignment: .leading)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .accessibilityElement(children: .combine)
        .parityItem(.row)
    }

    private var title: Text {
        let tone: Color = switch line.tone {
        case .met: FSColor.kill
        case .short: FSColor.wipe
        case .plain: FSColor.strong
        }
        let tally = Text(line.tally).foregroundColor(tone)
        let head = Text("\(line.role) \(tally)").foregroundColor(FSColor.strong)
        guard let split = line.split else { return head }
        let tail = Text(RaidFlowPresenter.separator + split).font(.custom(FSFont.body, size: FSType.item.size, relativeTo: .body))
        return Text("\(head)\(tail.foregroundColor(FSColor.muted))")
    }

    /// The names in class colour, `, ` between; past five `and N more`.
    private var names: Text {
        guard !line.people.isEmpty else { return Text(line.names).foregroundColor(FSColor.muted) }
        let named = line.people.enumerated().reduce(Text("")) { text, entry in
            let name = Text(entry.element.text).foregroundColor(ClassColor.of(entry.element.className))
            return entry.offset == 0 ? name : Text("\(text)\(Text(RaidFlowPresenter.nameSeparator).foregroundColor(FSColor.muted))\(name)")
        }
        guard line.more > 0 else { return named }
        return Text("\(named)\(Text(Copy.raidTabAndMore(names: "", n: line.more)).foregroundColor(FSColor.muted))")
    }
}

/// While the raid runs (spec 4.2): Raid night leads, the pull and your answer, then who is in by role.
struct RaidNightBlock: View {
    let content: RaidNightRootContent

    var body: some View {
        PanelSection(Copy.raidTabRaidNight, frame: .plain) {
            HStack(spacing: FSSpace.space2) {
                Circle().fill(FSColor.kill).frame(width: RaidLiveCard.liveDot, height: RaidLiveCard.liveDot).accessibilityHidden(true)
                Text(content.live.eyebrow).type(FSType.bodyStrong).foregroundStyle(FSColor.strong).accessibilityAddTraits(.isHeader)
            }
            .padding(.top, FSSpace.space3)
            RaidNote(content.live.line).padding(.top, -FSSpace.space1).padding(.bottom, FSSpace.space2)
            if !content.roles.isEmpty {
                Hairline(FSColor.line)
                RoleRowsView(roles: content.roles)
            }
        }
    }
}

/// A line of text on a panel (13/20, the text colour), 8 above.
struct RaidNote: View {
    let text: String
    var strong = false

    init(_ text: String, strong: Bool = false) {
        self.text = text
        self.strong = strong
    }

    var body: some View {
        Text(text)
            .type(FSType.small)
            .foregroundStyle(strong ? FSColor.strong : FSColor.text)
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.top, FSSpace.space2)
    }
}

/// A state with one thing to do: its lines on a panel under the title, then the button.
struct LonePanel<Action: View>: View {
    let title: String
    let lines: [String]
    @ViewBuilder let button: () -> Action

    var body: some View {
        PanelSection(title) {
            ForEach(Array(lines.enumerated()), id: \.offset) { index, line in
                RaidNote(line, strong: index == 0 && lines.count > 1)
            }
            button().padding(.vertical, FSSpace.space3)
        }
    }
}

/// The guild's pages as rows with their figures, titled with the guild's name.
struct GuildPagesSection: View {
    let content: GuildPagesContent
    let open: (GuildPage) -> Void

    var body: some View {
        PanelSection(content.guild) {
            ForEach(Array(content.rows.enumerated()), id: \.element.id) { index, row in
                Button { open(row.page) } label: {
                    PanelRow(row.page.title, sub: row.sub.map { Text($0) }, last: index == content.rows.count - 1, chevron: true)
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// Before you pull on a panel (round 5): each item's mark on its disc, the fix at the end of its line in
/// gold (the row opens it), then the export's age.
/// A check's lead: the mark, and beside it the item's icon when the check names one.
private struct ReadyLead: View {
    let mark: ReadyRowContent.Mark
    let icon: String?
    @Environment(\.raidIconBuild) private var build

    var body: some View {
        HStack(spacing: FSSpace.space2) {
            ReadyMark(mark: mark)
            if let icon {
                ItemIconView(icon: icon, quality: nil, build: build ?? "", size: FSSize.icon28).parityItem(.image)
            }
        }
    }
}

struct ReadyPanelList: View {
    let content: ReadyListContent
    let open: (URL) -> Void

    var body: some View {
        ForEach(content.rows) { row in
            Button { if let link = row.link { open(link) } } label: {
                // A check with nothing to add (s1: `Boots not enchanted`) draws no empty line.
                PanelRow(row.title, sub: row.fix.map { Text.line(row.line, link: $0) } ?? (row.line.isEmpty ? nil : Text(row.line))) {
                    ReadyLead(mark: row.mark, icon: row.icon)
                }
            }
            .buttonStyle(RowPressStyle())
            .disabled(row.link == nil)
        }
        if let stamp = content.stamp {
            Text(stamp).type(FSType.small).foregroundStyle(FSColor.muted).padding(.vertical, FSSpace.space2)
        }
    }
}
