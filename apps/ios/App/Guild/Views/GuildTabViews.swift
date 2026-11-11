import DesignSystem
import Models
import SwiftUI

/// What the Guild tab's own sections hand back (spec 2026-11-10-guild-tab.md section 2).
struct GuildTabActions {
    var openRoster: () -> Void = {}
    var openLoot: () -> Void = {}
    var openProgression: () -> Void = {}
    var openReport: (String) -> Void = { _ in }
    var openPlanner: () -> Void = {}
    var officer: (OfficerShortcut.Kind) -> Void = { _ in }
}

/// The guild head: the crest (the faction mark when the guild has none), the name, members and
/// faction; below 60 the level spread under it as three figures.
struct GuildHeadBlock: View {
    let head: GuildTabHead
    let spread: GuildSpread?

    static let crest: CGFloat = 48

    var body: some View {
        PanelBlock {
            HStack(spacing: FSSpace.space3) {
                GuildCrestMark(url: head.crestURL, faction: head.faction, size: Self.crest)
                VStack(alignment: .leading, spacing: 0) {
                    Text(head.name)
                        .type(TypeRole(face: FSFont.title, size: 22, lineHeight: 26, scalesWith: .title3))
                        .foregroundStyle(FSColor.cream)
                        .accessibilityAddTraits(.isHeader)
                    Text(head.line).type(RowType.sub).foregroundStyle(FSColor.muted)
                }
                Spacer(minLength: 0)
            }
            .padding(.vertical, FSSpace.space2)
            if let spread { SpreadFigures(spread: spread) }
        }
        .parityRegion(.section(Copy.tabGuild))
    }
}

/// `31 at 60 · 4 55 to 59 · 36 you`: the guild's level spread, your own level in gold.
private struct SpreadFigures: View {
    let spread: GuildSpread

    var body: some View {
        HStack(alignment: .top, spacing: 0) {
            figure(String(spread.atCap), Copy.guildTabAtCap(level: GuildTabPresenter.maxLevel), FSColor.strong)
            figure(
                String(spread.close),
                Copy.guildTabCloseRange(from: GuildTabPresenter.closeFrom, to: GuildTabPresenter.maxLevel - 1),
                FSColor.strong
            )
            figure(String(spread.you), Copy.guildTabYou, FSColor.gold)
        }
        .padding(.vertical, FSSpace.space2)
        .overlay(alignment: .top) { Hairline(FSColor.line) }
        .accessibilityElement(children: .combine)
    }

    private func figure(_ value: String, _ label: String, _ color: Color) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(value)
                .type(TypeRole(face: FSFont.bodyBold, size: 20, lineHeight: 24, scalesWith: .title3, tabular: true))
                .foregroundStyle(color)
            Text(label).type(FSType.small).foregroundStyle(FSColor.muted)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }
}

/// The guild's crest from the site, or its faction mark.
struct GuildCrestMark: View {
    let url: URL?
    let faction: String?
    let size: CGFloat

    var body: some View {
        Group {
            if let url {
                AsyncImage(url: url) { image in
                    image.resizable().scaledToFill()
                } placeholder: {
                    FactionMark(faction: faction ?? "", size: size)
                }
                .frame(width: size, height: size)
                .clipShape(.circle)
            } else {
                FactionMark(faction: faction ?? "", size: size)
            }
        }
        .accessibilityHidden(true)
        .parityItem(.image)
    }
}

/// Roster: Recently on, then below 60 Close to 60; `Roster ›` to the roster page.
struct GuildRosterSection: View {
    let content: GuildTabContent
    let open: () -> Void

    var body: some View {
        if !content.recent.isEmpty || !content.close.isEmpty {
            PanelSection(Copy.guildTabRoster, meta: content.rosterMeta, action: (Copy.raidTabRosterLink, open)) {
                if !content.recent.isEmpty {
                    GroupLabel(Copy.guildTabRecentlyOn)
                    MemberRows(lines: content.recent, last: content.close.isEmpty, levelTrail: false)
                }
                if !content.close.isEmpty {
                    GroupLabel(Copy.guildTabCloseTo(level: GuildTabPresenter.maxLevel))
                    MemberRows(lines: content.close, last: true, levelTrail: true)
                }
            }
        }
    }
}

private struct GroupLabel: View {
    let text: String

    init(_ text: String) { self.text = text }

    var body: some View {
        Text(text).type(FSType.small).fontWeight(.semibold).foregroundStyle(FSColor.muted)
            .padding(.top, FSSpace.space2)
            .accessibilityAddTraits(.isHeader)
    }
}

private struct MemberRows: View {
    let lines: [GuildMemberLine]
    let last: Bool
    let levelTrail: Bool

    var body: some View {
        ForEach(Array(lines.enumerated()), id: \.element.id) { index, line in
            GuildLineRow(
                title: line.name, titleColor: ClassColor.of(line.className), sub: line.line, trail: line.trail,
                trailStrong: levelTrail, last: last && index == lines.count - 1
            ) {
                ClassCrest(className: line.className, size: 32, flat: true)
            }
        }
    }
}

/// A row with a lead, the title and its line, and a short word at the right (`Seen 2 h ago`, `57`, the day).
struct GuildLineRow<Lead: View>: View {
    let title: String
    var titleColor: Color = FSColor.strong
    let sub: String
    /// A line in colours (an assignment slot's names) in place of `sub`.
    var styledSub: Text?
    var subColor: Color = FSColor.muted
    let trail: String?
    var trailStrong = false
    let last: Bool
    @ViewBuilder let lead: () -> Lead

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            lead()
            VStack(alignment: .leading, spacing: 0) {
                Text(title).type(FSType.item).foregroundStyle(titleColor)
                if let styledSub {
                    styledSub.type(RowType.sub)
                } else if !sub.isEmpty {
                    Text(sub).type(RowType.sub).foregroundStyle(subColor)
                }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if let trail {
                Text(trail)
                    .type(trailStrong
                        ? TypeRole(face: FSFont.bodyBold, size: 15, lineHeight: 20, scalesWith: .subheadline, tabular: true)
                        : RowType.sub)
                    .foregroundStyle(trailStrong ? FSColor.strong : FSColor.muted)
                    .multilineTextAlignment(.trailing)
            }
        }
        .padding(.vertical, FSSpace.rowPadY)
        .frame(maxWidth: .infinity, alignment: .leading)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .parityItem(.row)
    }
}

/// Recent logs: the guild's last reports, each opening the report.
struct GuildLogsSection: View {
    let logs: [GuildLogLine]
    let actions: GuildTabActions

    var body: some View {
        if !logs.isEmpty {
            PanelSection(Copy.guildTabLogs) {
                ForEach(Array(logs.enumerated()), id: \.element.id) { index, log in
                    Button { actions.openReport(log.id) } label: {
                        GuildLineRow(title: log.title, sub: log.line, trail: log.day, last: index == logs.count - 1) { EmptyView() }
                    }
                    .buttonStyle(RowPressStyle())
                }
            }
        }
    }
}

/// Progression: the tier's line, and Loot at 60.
/// Progression: the tier's line, opening the guild page's Progression on the site.
struct GuildProgressSection: View {
    let progression: GuildProgressLine?
    let open: () -> Void

    var body: some View {
        if let progression {
            PanelSection(Copy.guildTabProgression) {
                Button(action: open) {
                    GuildLineRow(title: progression.title, sub: progression.line, trail: nil, last: true) { EmptyView() }
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// Your job (spec 2.3a): every slot the viewer holds in the raid, with the partners, the boss and the
/// slot's purpose; left out when they hold none.
struct YourJobSection: View {
    let jobs: [RaidJobLine]

    var body: some View {
        if !jobs.isEmpty {
            PanelSection(Copy.sharedAssignYours) {
                ForEach(Array(jobs.enumerated()), id: \.element.id) { index, job in
                    PanelRow(job.title, sub: Text(job.line), last: index == jobs.count - 1)
                }
            }
        }
    }
}

/// An officer's read-only Assignments (spec 2.3a): each encounter's filled slots, the names (yours in
/// gold), how full; `Open in planner ›` to the raid's page on the site, where assigning stays.
struct AssignmentsSection: View {
    let encounters: [AssignmentEncounterLine]
    let openPlanner: () -> Void

    var body: some View {
        PanelSection(Copy.sharedAssignTitle, action: (Copy.guildTabOpenPlanner, openPlanner)) {
            if encounters.isEmpty {
                Text(Copy.sharedAssignNothingYet).type(RowType.sub).foregroundStyle(FSColor.text).padding(.vertical, FSSpace.space2)
            }
            ForEach(encounters) { encounter in
                GroupLabel(encounter.title)
                ForEach(encounter.slots) { slot in
                    GuildLineRow(
                        title: slot.label, sub: "", styledSub: AssignedNames.text(slot.names), trail: slot.filled, last: false
                    ) { EmptyView() }
                }
            }
            Text(Copy.guildTabAssignOnSite).type(RowType.sub).foregroundStyle(FSColor.muted).padding(.vertical, FSSpace.space2)
        }
        .accessibilityIdentifier("guild-assignments")
    }
}

/// A slot's names in class colour, the viewer's own in gold (the web planner's Tag tones).
enum AssignedNames {
    static func text(_ names: [RaidAssignedName]) -> Text {
        var line = AttributedString()
        for (index, name) in names.enumerated() {
            if index > 0 {
                var comma = AttributedString(", ")
                comma.foregroundColor = FSColor.muted
                line += comma
            }
            var piece = AttributedString(name.name)
            piece.foregroundColor = name.own ? FSColor.gold : ClassColor.of(name.className)
            line += piece
        }
        return Text(line)
    }
}

/// Loot at 60: the last raid's pieces, `Loot history ›` to the guild's loot.
struct GuildLootSection: View {
    let loot: String?
    let open: () -> Void

    var body: some View {
        if let loot {
            PanelSection(Copy.guildTabLoot, action: (Copy.guildTabLootLink, open)) {
                PanelRow(loot, last: true)
            }
        }
    }
}

/// Officer (spec 2.3): Plan a raid, Applications with their count, Setup with its progress; one tap each.
struct OfficerSection: View {
    let rows: [OfficerShortcut]
    let open: (OfficerShortcut.Kind) -> Void

    var body: some View {
        if !rows.isEmpty {
            PanelSection(Copy.guildTabOfficer) {
                ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                    Button { open(row.kind) } label: {
                        PanelRow(
                            row.title, sub: row.line.map { Text($0) }, last: index == rows.count - 1, chevron: true, value: row.count,
                            lead: { EmptyView() }, trail: { EmptyView() }
                        )
                    }
                    .buttonStyle(RowPressStyle())
                    .accessibilityIdentifier("guild-officer-\(row.kind.rawValue)")
                }
            }
        }
    }
}
