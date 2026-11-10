import DesignSystem
import SwiftUI

/// What the Raid tab's root hands back.
struct RaidRootActions {
    var raid = NextRaidActions()
    var openPage: (GuildPage) -> Void = { _ in }
    var openLink: (URL) -> Void = { _ in }
    var apply: (RecruitingRow) -> Void = { _ in }
    var claim: () -> Void = {}
    var findGuild: () -> Void = {}
    var signIn: () -> Void = {}
    var retry: () -> Void = {}
}

/// The Raid tab (flows spec 4.2) in the round-5 look (mocks ios-raid, -signup, -officer, -live,
/// -no-raid, -none, -leveler, -out and m-ios-raid): under the band's large title, the one-line raid
/// head with your answer, Who is coming, Before you pull, an officer's Needs you, then the guild's
/// pages; raid night leads while a raid is live. No guild: the guilds recruiting your spec; under 60,
/// when raids start.
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
                case let .leveler(toGo):
                    LonePanel(title: Copy.raidTabLeveler, lines: [toGo, Copy.raidTabLevelerLine]) {
                        FSButton(Copy.homeFindGuildButton, action: actions.findGuild)
                    }
                case let .noGuild(recruiting): noGuild(recruiting)
                case .waiting:
                    PanelBlock { RaidNote(Copy.raidsWaiting) }.parityRegion(RaidRootView.nextRaidRegion)
                case let .guild(guild): RaidGuildView(content: guild, actions: actions)
                }
            }
        }
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
        if !recruiting.rows.isEmpty {
            PanelSection(recruiting.heading, frame: .plain) {
                ForEach(Array(recruiting.rows.enumerated()), id: \.element.id) { index, row in
                    PanelRow(row.name, sub: Text(row.line), last: index == recruiting.rows.count - 1) {
                        EmptyView()
                    } trail: {
                        Button(Copy.raidTabApply) { actions.apply(row) }.buttonStyle(FSButtonStyle(fullWidth: false))
                    }
                }
            }
        }
        LonePanel(title: Copy.raidTabNotHere, lines: [Copy.raidTabClaimLine]) {
            FSButton(Copy.raidTabClaim, action: actions.claim)
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
        if let night = content.night {
            RaidNightBlock(content: night)
        } else if let next = content.next {
            RaidHeadView(content: next, actions: actions.raid).parityRegion(RaidRootView.nextRaidRegion)
                .environment(\.answersOffline, content.offlineStamp != nil)
            if !content.coming.isEmpty {
                PanelSection(Copy.raidTabComing, action: (Copy.raidTabRosterLink, { actions.openPage(.roster) })) {
                    RoleRowsView(roles: content.coming)
                }
            }
        } else {
            PanelBlock { RaidNote(Copy.raidTabNoRaid).padding(.bottom, FSSpace.space2) }.parityRegion(RaidRootView.nextRaidRegion)
        }
        if let ready = content.ready {
            // Once pulled the list is the Ready check, not what comes before the pull.
            PanelSection(content.night == nil ? Copy.raidTabBeforePull : Copy.raidTabReadyCheck) {
                ReadyPanelList(content: ready, open: actions.openLink)
            }
        }
        if let needs = content.needsYou {
            PanelSection(Copy.raidTabNeedsYou, meta: needs.meta) { OfficerItemsView(items: needs.items, actions: actions.raid) }
        }
        GuildPagesSection(content: content.pages, open: actions.openPage)
    }
}

/// The Raid tab's head (spec 4.2, owner 2026-11-09): one compact line, the zone's icon, the raid, its
/// time and the pull, who the answer is for (E7) and the answer as a chip; the four answers under it,
/// and with Late the minutes inline, the chosen one selected (F1).
struct RaidHeadView: View {
    let content: NextRaidContent
    let actions: NextRaidActions
    @Environment(\.dynamicTypeSize) private var typeSize

    /// How far the head's line may shrink to stay on one line beside the chip.
    static let lineScale: CGFloat = 0.8

    var body: some View {
        PanelBlock {
            VStack(alignment: .leading, spacing: 0) {
                Button(action: actions.openRaid) { line }.buttonStyle(RowPressStyle())
                RaidAnswerButtons(selected: content.selected, answer: actions.answer).padding(.top, FSSpace.space3)
                if content.lateOpen {
                    LateMinutesRow(chosen: content.lateMinutes, late: actions.late)
                        .padding(.top, FSSpace.space2)
                        .accessibilityElement(children: .contain)
                        .accessibilityLabel(Copy.raidTabHowLate)
                }
            }
            .padding(.vertical, FSSpace.space2)
        }
    }

    private var line: some View {
        HStack(spacing: FSSpace.space3) {
            ZoneIcon(icon: content.icon)
            VStack(alignment: .leading, spacing: 0) {
                Text(content.zone).type(FSType.item).foregroundStyle(FSColor.strong).accessibilityAddTraits(.isHeader)
                // One line beside the chip, as the mock holds it (`.rl .s { white-space: nowrap }`).
                Text(content.headLine).type(RowType.sub).foregroundStyle(FSColor.muted)
                    .lineLimit(1).minimumScaleFactor(Self.lineScale)
                if let who = content.who { Text(who).type(RowType.sub).foregroundStyle(FSColor.muted) }
                // At the accessibility sizes the chip drops under the lines, so the raid's name keeps the row.
                if typeSize.isAccessibilitySize, let chip = content.chip { AnswerChip(text: chip).padding(.top, FSSpace.space1) }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if !typeSize.isAccessibilitySize, let chip = content.chip { AnswerChip(text: chip) }
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

/// Your answer as a chip at the head's end, teal as the selected answer: `In`, `Late · 30 min`.
struct AnswerChip: View {
    let text: String

    static let padX: CGFloat = 10
    static let padY: CGFloat = 3

    var body: some View {
        Text(text)
            .type(TypeRole(face: FSFont.bodyBold, size: 13, lineHeight: 18, scalesWith: .footnote))
            .foregroundStyle(FSColor.strong)
            .lineLimit(1)
            .padding(.horizontal, Self.padX)
            .padding(.vertical, Self.padY)
            .background(FSColor.teal, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).strokeBorder(FSColor.tealEdge, lineWidth: FSStroke.hairline))
            .fixedSize()
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

/// One role: `Healers 5 of 8` over its names.
private struct RoleRow: View {
    let line: RoleLineContent
    let last: Bool

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            title.type(FSType.item)
            Text(line.names).type(RowType.sub).foregroundStyle(FSColor.muted)
        }
        .padding(.vertical, FSSpace.rowPadY)
        .frame(maxWidth: .infinity, alignment: .leading)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .accessibilityElement(children: .combine)
        .parityItem(.row)
    }

    private var title: Text {
        guard let want = line.want else { return Text("\(line.role) \(line.have)").foregroundColor(FSColor.strong) }
        let count = Text(Copy.raidTabOfWant(have: line.have, want: want)).foregroundColor(line.have >= want ? FSColor.kill : FSColor.wipe)
        return Text("\(line.role) \(count)").foregroundColor(FSColor.strong)
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
struct ReadyPanelList: View {
    let content: ReadyListContent
    let open: (URL) -> Void

    var body: some View {
        ForEach(content.rows) { row in
            Button { if let link = row.link { open(link) } } label: {
                PanelRow(row.title, sub: row.fix.map { Text.line(row.line, link: $0) } ?? Text(row.line)) {
                    ReadyMark(mark: row.mark)
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
