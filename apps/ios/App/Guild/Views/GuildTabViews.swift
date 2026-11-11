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
            HStack(alignment: .top, spacing: FSSpace.space3) {
                GuildCrestMark(url: head.crestURL, faction: head.faction, size: Self.crest)
                VStack(alignment: .leading, spacing: 0) {
                    PanelTitle(head.name)
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
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        // At the accessibility sizes the three figures stack, so no label wraps.
        let layout = typeSize.isAccessibilitySize
            ? AnyLayout(VStackLayout(alignment: .leading, spacing: FSSpace.space2))
            : AnyLayout(HStackLayout(alignment: .top, spacing: 0))
        layout {
            figure(String(spread.atCap), Copy.guildTabAtCap(level: GuildTabPresenter.maxLevel), FSColor.strong)
            figure(
                String(spread.close),
                Copy.guildTabCloseRange(from: GuildTabPresenter.closeFrom, to: GuildTabPresenter.maxLevel - 1),
                FSColor.strong
            )
            figure(String(spread.third), spread.thirdLabel, spread.own ? FSColor.gold : FSColor.strong)
        }
        .padding(.vertical, FSSpace.space2)
        .overlay(alignment: .top) { Hairline(FSColor.line) }
        .accessibilityElement(children: .combine)
    }

    private func figure(_ value: String, _ label: String, _ color: Color) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(value).type(FSType.itemLg).monospacedDigit().foregroundStyle(color)
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

    /// The faction mark's share of the ringed circle.
    static let markShare: CGFloat = 0.72

    var body: some View {
        ZStack {
            Circle().fill(FSColor.raised)
            if let url {
                AsyncImage(url: url) { image in
                    image.resizable().scaledToFill()
                } placeholder: {
                    FactionMark(faction: faction ?? "", size: size * Self.markShare)
                }
                .frame(width: size, height: size)
            } else {
                FactionMark(faction: faction ?? "", size: size * Self.markShare)
            }
        }
        .frame(width: size, height: size)
        .clipShape(.circle)
        .overlay(Circle().stroke(GameColors.faction(faction ?? "").opacity(0.6), lineWidth: FSStroke.crestRing))
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
                    MemberRows(
                        lines: content.recent,
                        last: content.close.isEmpty && content.missing.isEmpty && content.gearGaps.isEmpty && content.yourAttendance == nil,
                        levelTrail: false
                    )
                }
                if !content.close.isEmpty {
                    GroupLabel(Copy.guildTabCloseTo(level: GuildTabPresenter.maxLevel))
                    MemberRows(
                        lines: content.close, last: content.missing.isEmpty && content.gearGaps.isEmpty && content.yourAttendance == nil,
                        levelTrail: true
                    )
                }
                if !content.missing.isEmpty {
                    GroupLabel(Copy.guildTabMissingRaids)
                    MemberRows(lines: content.missing, last: content.gearGaps.isEmpty && content.yourAttendance == nil, levelTrail: false)
                }
                if !content.gearGaps.isEmpty {
                    GroupLabel(Copy.guildTabGearGaps)
                    MemberRows(lines: content.gearGaps, last: content.yourAttendance == nil, levelTrail: false)
                }
                if let yours = content.yourAttendance {
                    GroupLabel(Copy.guildTabYourAttendance)
                    MemberRows(lines: [yours], last: true, levelTrail: false)
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
    /// The row opens a page: iOS's disclosure chevron at its end.
    var chevron = false
    @ViewBuilder let lead: () -> Lead

    var body: some View {
        HStack(spacing: FSSpace.space3) {
            lead()
            VStack(alignment: .leading, spacing: 0) {
                Text(title).type(FSType.item).foregroundStyle(titleColor)
                // The line wraps rather than truncates (the trail keeps its own width).
                if let styledSub {
                    styledSub.type(RowType.sub).fixedSize(horizontal: false, vertical: true)
                } else if !sub.isEmpty {
                    Text(sub).type(RowType.sub).foregroundStyle(subColor).fixedSize(horizontal: false, vertical: true)
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
                    .fixedSize()
            }
            if chevron {
                Image(systemName: "chevron.right")
                    .font(.system(size: 13, weight: .semibold))
                    .foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
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
                        GuildLineRow(title: log.title, sub: log.line, trail: log.day, last: index == logs.count - 1, chevron: true) { EmptyView() }
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
                    GuildLineRow(title: progression.title, sub: progression.line, trail: nil, last: true, chevron: true) { EmptyView() }
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
        PanelSection(Copy.sharedAssignTitle, action: (Copy.guildTabAssignPlanner, openPlanner)) {
            if encounters.isEmpty {
                Text(Copy.sharedAssignNothingYet).type(RowType.sub).foregroundStyle(FSColor.text).padding(.vertical, FSSpace.space2)
            }
            ForEach(encounters) { encounter in
                GroupLabel(encounter.title)
                ForEach(encounter.slots) { slot in
                    GuildLineRow(
                        title: slot.label, titleColor: slot.names.isEmpty ? FSColor.muted : FSColor.strong, sub: "",
                        styledSub: slot.names.isEmpty ? nil : AssignedNames.text(slot.names), trail: slot.filled, last: false
                    ) { EmptyView() }
                }
            }
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
            piece.foregroundColor = name.own ? FSColor.gold : FSColor.muted
            line += piece
        }
        return Text(line)
    }
}

/// Loot at 60: the last raid's pieces, `Loot history ›` to the guild's loot.
struct GuildLootSection: View {
    let night: GuildLootNight?
    let loot: String?
    let open: () -> Void
    @Environment(\.raidIconBuild) private var build

    var body: some View {
        if let night {
            PanelSection(Copy.guildTabLoot, meta: night.day, action: (Copy.guildTabLootLink, open)) {
                ForEach(Array(night.drops.enumerated()), id: \.element.id) { index, drop in
                    GuildLineRow(
                        title: drop.name, titleColor: ItemQuality(value: drop.quality).text,
                        sub: drop.winner == nil ? Copy.guildTabNotAwarded : "",
                        styledSub: drop.winner.map { Self.winner($0, className: drop.winnerClass, yours: drop.yours) },
                        trail: nil, last: index == night.drops.count - 1
                    ) {
                        ItemIconView(icon: drop.icon, quality: drop.quality, build: build ?? "", size: FSSize.icon40).parityItem(.image)
                    }
                }
            }
        } else if let loot {
            PanelSection(Copy.guildTabLoot, action: (Copy.guildTabLootLink, open)) {
                PanelRow(loot, last: true)
            }
        }
    }
    static func winner(_ name: String, className: String?, yours: Bool) -> Text {
        var line = AttributedString(name)
        line.foregroundColor = yours ? FSColor.gold : ClassColor.of(className)
        return Text(line)
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
