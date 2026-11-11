@testable import ForeverSixty
import Foundation
import Models

/// The Guild tab's mocked states (design/mocks/ws_app_guild.py, spec 2026-11-10-guild-tab.md) as
/// content values: Olympus XXVII, 38 members, 31 at 60; Thandrel at 60 (member, officer), Felsnap at 36.
enum GuildTabTestData {
    typealias R = RaidFlowTestData

    static let head = GuildTabHead(name: "Olympus XXVII", faction: "alliance", crestURL: nil, line: "38 members · Alliance · US PvP")

    private static func member(_ name: String, _ className: String, _ line: String, _ trail: String) -> GuildMemberLine {
        GuildMemberLine(name: name, className: className, specIcon: nil, line: line, trail: trail)
    }

    static let recent = [
        member("Grunvald", "warrior", "60 · Protection Warrior · ilvl 66", "just now"),
        member("Mirethil", "priest", "60 · Holy Priest · ilvl 63", "12 min ago"),
        member("Kaelra", "mage", "60 · Frost Mage · ilvl 61", "40 min ago"),
        member("Thornwick", "druid", "57 · Feral Druid", "1 h ago"),
        member("Selka", "druid", "60 · Restoration Druid · ilvl 62", "2 h ago"),
    ]

    static let close = [
        member("Vexa", "warlock", "Affliction Warlock", "56"),
        member("Pell", "mage", "Frost Mage", "55"),
        member("Orrin", "paladin", "Holy Paladin", "55"),
    ]

    static let gearGaps = [
        member("Selka", "druid", "Restoration Druid · ilvl 62", "41 HPS"),
        member("Thalrin", "hunter", "Marksmanship Hunter · ilvl 58", "37 DPS"),
        member("Kaelra", "mage", "Frost Mage · ilvl 61", "22 DPS"),
    ]

    static let yourAttendance = member("Thandrel", "warrior", "Fury Warrior", "7 of 8 nights")

    static let missing = [
        member("Brakka", "druid", "Feral Druid · ilvl 60", "2 of 8 nights"),
        member("Thalrin", "hunter", "Marksmanship Hunter · ilvl 58", "3 of 7 nights"),
    ]

    static let logs = [
        GuildLogLine(id: "r3", title: "UBRS, Sunday", line: "Blackrock Spire · 7/8 kills · 1 h 32 min", day: "Sun 3 Jan", clean: false),
        GuildLogLine(id: "r2", title: "Scholomance, Saturday", line: "Scholomance · 10/10 kills · 1 h 08 min", day: "Sat 2 Jan", clean: true),
        GuildLogLine(id: "r1", title: "Onyxia, Wednesday", line: "Onyxia's Lair · 1/1 kills · 41 min", day: "Wed 30 Dec", clean: true),
    ]

    static let progression = GuildProgressLine(title: "Onyxia's Lair", line: "1 of 1 down · 6 pulls")

    static let lootNight = GuildLootNight(zone: "Onyxia's Lair", day: "Wed 30 Dec", drops: [
        .init(itemID: 18205, name: "Eskhandar's Collar", icon: "inv_jewelry_necklace_07", quality: 4, winner: "Kraggor", winnerClass: "warrior"),
        .init(itemID: 17067, name: "Ancient Cornerstone Grimoire", icon: "inv_misc_book_07", quality: 4, winner: "Thandrel", winnerClass: "warrior", yours: true),
        .init(itemID: 18422, name: "Head of Onyxia", icon: "inv_misc_head_dragon_01", quality: 4, winner: nil, winnerClass: nil),
    ])

    static let officerRows = [
        OfficerShortcut(kind: .plan, title: Copy.guildTabPlanRaid, line: "Onyxia Wed, Barrow Deeps Sat", count: nil),
        OfficerShortcut(kind: .applications, title: Copy.guildTabApplications, line: "Mirelle, Holy Priest 60", count: "1 waiting"),
        OfficerShortcut(
            kind: .setup, title: Copy.guildTabSetup, line: "Discord bot, raid schedule and addon sync to go", count: "4 of 8 done"
        ),
    ]

    static func tab(leveler: Bool = false, officer: Bool = false) -> GuildTabContent {
        var tab = GuildTabContent(head: head)
        tab.leveler = leveler
        tab.spread = leveler
            ? GuildSpread(atCap: 31, close: 4, third: 36, thirdLabel: Copy.guildTabYou, own: true)
            : GuildSpread(atCap: 31, close: 4, third: 24, thirdLabel: Copy.guildTabRaiders, own: false)
        tab.rosterMeta = "31 of 38 at 60"
        tab.recent = recent
        tab.close = leveler ? close : []
        tab.logs = logs
        tab.progression = progression
        tab.lootNight = leveler ? nil : lootNight
        tab.officer = officer ? officerRows : []
        tab.missing = officer ? missing : []
        tab.gearGaps = officer ? gearGaps : []
        tab.yourAttendance = leveler ? nil : yourAttendance
        return tab
    }

    static let job = RaidJobLine(
        encounter: "Onyxia", slot: "whelps_west", title: "Whelps, west cave with Kaelra and Thalrin",
        line: "Onyxia · Pick up the west cave's whelps in phases 2 and 3 and kill them away from the raid."
    )

    private static func picked(_ name: String, _ className: String, own: Bool = false) -> RaidAssignedName {
        RaidAssignedName(characterKey: "us/pvp/\(name.lowercased())", name: name, className: className, group: "", state: "in", own: own, offRole: false)
    }

    static let assignments = [AssignmentEncounterLine(encounterID: 1084, title: "Onyxia", slots: [
        .init(key: "main_tank", label: "Main tank", names: [picked("Grunvald", "warrior")], filled: "1 of 1"),
        .init(key: "off_tank", label: "Off tank (phase 3, fear)", names: [picked("Brakka", "druid")], filled: "1 of 2"),
        .init(key: "main_tank_heals", label: "Main tank healers", names: [
            picked("Mirethil", "priest"), picked("Selka", "druid"), picked("Hollowmend", "priest"), picked("Aenwyn", "paladin"),
        ], filled: "4 of 4"),
        .init(key: "off_tank_heals", label: "Off tank healers", names: [], filled: "0 of 2"),
        .init(key: "fear_cover", label: "Fear cover", names: [picked("Mirethil", "priest")], filled: "1 of 2"),
        .init(key: "whelps_west", label: "Whelps, west cave", names: [
            picked("Thandrel", "warrior", own: true), picked("Kaelra", "mage"), picked("Thalrin", "hunter"),
        ], filled: "3 of 3"),
        .init(key: "whelps_east", label: "Whelps, east cave", names: [], filled: "0 of 3"),
    ])]

    /// The Guild tab's pages once its sections show: Calendar and Guild sims at 60, Calendar and Loot history under 60.
    static func pages(leveler: Bool) -> GuildPagesContent {
        RaidRootComposer.pages(R.pages, tab: tab(leveler: leveler))
    }

    /// At 60 (spec 2.2), and for an officer (2.3, 2.3a).
    static func atCap(officer: Bool = false) -> RaidRootContent {
        guard case var .guild(guild) = R.member(officer: officer) else { return R.member(officer: officer) }
        guild.tab = tab(officer: officer)
        guild.pages = pages(leveler: false)
        guild.jobs = [job]
        guild.assignments = officer ? assignments : nil
        return .guild(guild)
    }

    /// Below 60 (spec 2.1): the guild leads, the raid head answers as at 60, one quiet line.
    static var leveler: RaidRootContent {
        guard case var .guild(guild) = R.member() else { return R.member() }
        guild.next?.who = "As Felsnap, Damage"
        guild.ready = nil
        guild.tab = tab(leveler: true)
        guild.pages = pages(leveler: true)
        guild.raidsOpen = Copy.guildTabRaidsOpen(level: 60, you: 36)
        guild.answers = nil
        guild.cover = nil
        guild.notSynced = nil
        guild.coming = guild.coming.map(\.countOnly)
        return .guild(guild)
    }

    /// No guild (spec 2.4): Olympus already applied to, Dawnbreak with Apply; under 60 the quiet line.
    static func noGuild(leveler: Bool) -> RaidRootContent {
        var olympus = RecruitingRow(guildID: 2, name: "Olympus XXVII", line: "Alliance · Wed and Sat, 8:00 pm CST · needs damage")
        olympus.applied = Copy.guildTabApplied(day: "Mon")
        let dawnbreak = RecruitingRow(guildID: 7, name: "Dawnbreak", line: "Alliance · Tue and Thu, 9:00 pm EST · needs damage")
        var content = RecruitingContent(
            heading: leveler ? "Guilds recruiting a Beast Mastery Hunter" : "Guilds recruiting a Fury Warrior", rows: [olympus, dawnbreak]
        )
        content.raidsOpen = leveler ? Copy.guildTabRaidsOpen(level: 60, you: 36) : nil
        return .noGuild(content)
    }
}
