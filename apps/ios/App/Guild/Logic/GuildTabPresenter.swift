import Auth
import Foundation
import Models

/// The Guild tab's own sections (spec 2026-11-10-guild-tab.md section 2), around the raid the tab
/// already draws: the guild head and level spread, the roster's highlights, the recent logs,
/// progression, the officer's shortcuts, the loot line.
struct GuildTabContent: Equatable {
    var head: GuildTabHead
    /// Below 60: the guild's level spread (the head leads the tab); nil at 60.
    var spread: GuildSpread?
    /// `31 of 38 at 60`, on the Roster section's title row.
    var rosterMeta: String?
    var recent: [GuildMemberLine] = []
    /// Below 60: the members at 55 to 59, highest first.
    var close: [GuildMemberLine] = []
    var logs: [GuildLogLine] = []
    var progression: GuildProgressLine?
    var loot: String?
    var officer: [OfficerShortcut] = []
}

struct GuildTabHead: Equatable {
    let name: String
    let faction: String?
    let crestURL: URL?
    /// `38 members · Alliance`.
    let line: String
}

/// `31 at 60 · 4 at 55 to 59 · you 36`, as three figures.
struct GuildSpread: Equatable {
    let atCap: Int
    let close: Int
    let you: Int
}

struct GuildMemberLine: Equatable, Identifiable {
    var id: String { name }
    let name: String
    let className: String?
    let specIcon: String?
    /// `60 · Protection Warrior`, or the spec alone where the trail carries the level.
    let line: String
    /// `Seen 12 min ago`, or the level.
    let trail: String
}

struct GuildLogLine: Equatable, Identifiable {
    let id: String
    let title: String
    let line: String
    let day: String
    /// Every fight a kill.
    let clean: Bool
}

struct GuildProgressLine: Equatable {
    let title: String
    let line: String
}

/// One officer shortcut (spec 2.3), one tap from the tab.
struct OfficerShortcut: Equatable, Identifiable {
    enum Kind: String { case plan, applications, setup }
    var id: Kind { kind }
    let kind: Kind
    let title: String
    let line: String?
    let count: String?
}

/// What the Guild tab reads for its own sections: the guild home, and for an officer the setup
/// checklist and the pending applications. Each is optional; a call that fails leaves its part out.
struct GuildTabInputs: Equatable {
    var home: GuildHome?
    var setup: GuildSetup?
    var applicants: [ApplicantLine] = []
    var badge: RaidBadge?
    /// The selected character's level and its key (the viewer is left out of the roster's lists).
    var level: Int?
    var characterKey: String?
    var planned: [RaidView] = []
    var firstRaids: Date?
}

/// Turns the guild home into the Guild tab's sections. Pure, so each state is a test of values.
struct GuildTabPresenter {
    let now: Date
    var clock = RaidClock()

    static let maxLevel = RaidFlowPresenter.maxLevel
    /// Close to 60 (spec 2.1): 55 to 59.
    static let closeFrom = 55
    static let recentShown = 5
    static let closeShown = 5
    static let logsShown = 3
    /// A setup step still to do (api/internal/guilds/setup.go SetupTodo).
    static let setupTodo = "todo"

    func content(_ inputs: GuildTabInputs, guildName: String) -> GuildTabContent? {
        guard let home = inputs.home else { return nil }
        let leveler = (inputs.level ?? Self.maxLevel) < Self.maxLevel
        let levels = home.members.compactMap(\.level)
        let atCap = levels.filter { $0 >= Self.maxLevel }.count
        var content = GuildTabContent(head: head(home, name: guildName))
        content.rosterMeta = Copy.guildTabRosterMeta(n: atCap, total: home.members.count, level: Self.maxLevel)
        content.recent = recent(home, excluding: inputs.characterKey)
        if leveler {
            content.spread = GuildSpread(
                atCap: atCap, close: levels.filter { (Self.closeFrom..<Self.maxLevel).contains($0) }.count, you: inputs.level ?? 0
            )
            content.close = close(home, excluding: inputs.characterKey)
        }
        content.logs = home.reports.prefix(Self.logsShown).map(log)
        content.progression = progression(home.summary, firstRaids: inputs.firstRaids)
        content.loot = leveler ? nil : inputs.badge?.lastRaidLoot.map { Copy.raidPagesLootLast(n: $0.count) }
        if home.viewer.officerTools { content.officer = officer(inputs) }
        return content
    }

    // MARK: Parts

    private func head(_ home: GuildHome, name: String) -> GuildTabHead {
        let members = Copy.raidPagesMembers(n: home.counts.onRoster)
        let faction = home.guild.faction.map(\.capitalized)
        return GuildTabHead(
            name: home.guild.name.isEmpty ? name : home.guild.name,
            faction: home.guild.faction,
            crestURL: home.guild.crestUrl.flatMap(URL.init(string:)),
            line: faction.map { Copy.guildTabHeadLine(members: members, faction: $0) } ?? members
        )
    }

    /// The five most recently seen members (the addon's `seen_at`, read in the viewer's zone as the
    /// site reads it), the viewer left out.
    func recent(_ home: GuildHome, excluding key: String?) -> [GuildMemberLine] {
        home.members
            .filter { $0.characterKey == nil || $0.characterKey != key }
            .compactMap { member in Self.seenAt(member.seenAt).map { (member, $0) } }
            .sorted { $0.1 > $1.1 }
            .prefix(Self.recentShown)
            .map { member, seen in
                GuildMemberLine(
                    name: member.name, className: member.className, specIcon: member.specIcon,
                    line: [member.level.map(String.init), member.specName].compactMap(\.self).joined(separator: " · "),
                    trail: Copy.guildTabSeen(age: RelativeAge(from: seen, to: now).phrase)
                )
            }
    }

    func close(_ home: GuildHome, excluding key: String?) -> [GuildMemberLine] {
        home.members
            .filter { ($0.characterKey == nil || $0.characterKey != key) && (Self.closeFrom..<Self.maxLevel).contains($0.level ?? 0) }
            .sorted { ($0.level ?? 0, $1.name) > ($1.level ?? 0, $0.name) }
            .prefix(Self.closeShown)
            .map { member in
                GuildMemberLine(
                    name: member.name, className: member.className, specIcon: member.specIcon,
                    line: member.specName ?? "", trail: String(member.level ?? 0)
                )
            }
    }

    private func log(_ report: GuildReportRow) -> GuildLogLine {
        GuildLogLine(
            id: report.id, title: report.title,
            line: Copy.guildTabLogLine(
                zone: report.zone, kills: report.killCount, fights: report.fightCount, duration: Self.duration(report.durationMs)
            ),
            day: clock.shortDate(report.createdAt),
            clean: report.fightCount > 0 && report.killCount == report.fightCount
        )
    }

    /// The tier's line: `1 of 1 down · 6 pulls this tier`; before the first pull, when the first raids open.
    func progression(_ summary: GuildSummary, firstRaids: Date?) -> GuildProgressLine? {
        let named = summary.tier?.namedEncounters ?? 0
        let title = summary.tier?.name ?? summary.tier?.raids?.first
        guard let title, named > 0 else { return nil }
        let pulls = summary.tier?.pullsTotal ?? summary.pullsThisTier
        let down = summary.tier?.down ?? summary.namedEncountersDown
        // Before the first pull: when the first raids open; with no pull and no date, nothing yet.
        if pulls == 0 {
            guard let firstRaids, firstRaids > now else { return nil }
            return GuildProgressLine(title: title, line: Copy.guildTabFirstRaids(date: clock.shortDate(firstRaids)))
        }
        let line = [Copy.guildTabDown(down: down, named: named), Copy.guildTabPulls(n: pulls)].joined(separator: " · ")
        return GuildProgressLine(title: title, line: line)
    }

    private func officer(_ inputs: GuildTabInputs) -> [OfficerShortcut] {
        let upcoming = inputs.planned.filter { $0.startsAt > now }.sorted { $0.startsAt < $1.startsAt }
        let planned = upcoming.isEmpty ? Copy.guildTabNonePlanned : upcoming.map { raid in
            Copy.raidPagesCalendarItem(zone: raid.title, weekday: String(clock.weekday(raid.startsAt).prefix(3)))
        }.joined(separator: ", ")
        var rows = [OfficerShortcut(kind: .plan, title: Copy.guildTabPlanRaid, line: planned, count: nil)]
        let waiting = max(inputs.applicants.count, inputs.badge?.applicants ?? 0)
        let newest = inputs.applicants.first.map { Copy.guildTabApplicant(name: $0.name, kind: $0.kind, level: $0.level) }
        let line = newest.map { waiting > 1 ? Copy.guildTabApplicantMore(first: $0, n: waiting - 1) : $0 }
        rows.append(OfficerShortcut(
            kind: .applications, title: Copy.guildTabApplications, line: line ?? Copy.guildTabNoApplications,
            count: waiting > 0 ? Copy.guildTabWaiting(n: waiting) : nil
        ))
        if let setup = inputs.setup, !setup.finished {
            let todo = setup.steps.filter { $0.state == GuildTabPresenter.setupTodo }.map { Self.stepWord($0.key) }.compactMap(\.self)
            rows.append(OfficerShortcut(
                kind: .setup, title: Copy.guildTabSetup,
                line: todo.isEmpty ? nil : Copy.guildTabSetupToGo(steps: RaidFlowPresenter.list(todo)),
                count: Copy.guildTabSetupDone(done: setup.done, total: setup.total)
            ))
        }
        return rows
    }

    // MARK: Words

    /// `1 h 32 min`, `41 min`.
    static func duration(_ ms: Int64) -> String {
        let minutes = Int(ms / 60000)
        let hours = minutes / 60
        return hours > 0 ? Copy.raidTabUntilHours(h: hours, m: minutes % 60) : Copy.raidTabUntilMinutes(m: minutes)
    }

    /// The setup step's name, as the site's checklist titles it.
    static func stepWord(_ key: String) -> String? {
        switch key {
        case "officer_ranks": Copy.guildTabStepRanks
        case "discord_bot": Copy.guildTabStepDiscord
        case "recruiting": Copy.guildTabStepRecruiting
        case "raid_schedule": Copy.guildTabStepSchedule
        case "invite_link": Copy.guildTabStepInvite
        case "addon_sync": Copy.guildTabStepAddon
        case "report_visibility": Copy.guildTabStepVisibility
        case "linked_role": Copy.guildTabStepRole
        default: nil
        }
    }

    /// The member's `seen_at`, their own local `YYYY-MM-DD HH:MM`, read in the viewer's zone (the
    /// addon announces no zone; web's roster-addon.ts reads it the same way).
    static func seenAt(_ text: String?, zone: TimeZone = .current) -> Date? {
        guard let text else { return nil }
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.timeZone = zone
        formatter.dateFormat = "yyyy-MM-dd HH:mm"
        return formatter.date(from: text)
    }
}

/// The guild's pages on the site the Guild tab opens (the web guild page's tabs and its raid planner).
enum GuildWeb {
    /// `{guild page}#progression`.
    static func tab(_ guild: URL, _ tab: String) -> URL {
        var parts = URLComponents(url: guild, resolvingAgainstBaseURL: true)
        parts?.fragment = tab
        return parts?.url ?? guild
    }

    /// The raid's page on the site, where its planner and assignments live.
    static func raid(_ guild: URL, raid: Int64) -> URL {
        guild.appending(path: "raids/\(raid)")
    }
}
