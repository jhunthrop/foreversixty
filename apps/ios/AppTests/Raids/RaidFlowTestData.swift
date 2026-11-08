@testable import ForeverSixty
import DesignSystem
import Foundation

/// The flows mocks' Raid scenario (design/mocks/ws_app_flows.py `raid`, `raid_unanswered`,
/// `applicant`; ws_app.py `raid_night`) as content values: Olympus XXVII's Onyxia's Lair, Wed 6 Jan,
/// 8:00 pm CST, healers short by 3; Mirelle the applicant.
enum RaidFlowTestData {
    typealias H = HomeFlowTestData

    static let pages = GuildPagesContent(guild: "Olympus XXVII", rows: [
        GuildPageRow(page: .roster, sub: "38 members, 31 at 60"),
        GuildPageRow(page: .calendar, sub: "Onyxia Wed, Barrow Deeps Sat"),
        GuildPageRow(page: .loot, sub: "Last raid: 4 pieces"),
        GuildPageRow(page: .sims, sub: "You are #7 of 14 DPS"),
    ])

    static let officer = OfficerContent(
        unanswered: HomeLine(title: "5 have not answered, 2 of them healers", sub: "Healers short by 3"),
        applicants: HomeLine(title: "1 applicant", sub: "Mirelle, Holy Priest 60, 2 min ago")
    )

    static func member(officer: Bool = false) -> RaidRootContent {
        .guild(RaidGuildContent(
            next: H.raid(), needsYou: officer ? NeedsYouContent(meta: "6 open", items: Self.officer) : nil, pages: pages
        ))
    }

    static let noRaid = RaidRootContent.guild(RaidGuildContent(pages: pages))

    static func ready(elixir: Bool) -> ReadyListContent {
        let attuned = elixir ? "Attuned to Onyxia’s Lair" : "Attuned"
        var rows = [ReadyRowContent(key: "attune.onyxias-lair", mark: .ok, title: "Drakefire Amulet", line: attuned, fix: nil, link: nil)]
        if elixir {
            rows.append(ReadyRowContent(
                key: "food", mark: .ok, title: "Elixir of the Mongoose", line: "3 in your bags, from the addon’s last export", fix: nil, link: nil
            ))
        }
        rows += [
            ReadyRowContent(key: "flask", mark: .miss, title: "Flask of the Titans", line: "None in your bags", fix: "Auction house", link: nil),
            ReadyRowContent(
                key: "enchant.feet", mark: .warn, title: "Boots not enchanted", line: "Enchant Boots: Greater Agility",
                fix: "Where to get it", link: URL(string: "https://foreversixty.gg/")
            ),
        ]
        return ReadyListContent(rows: rows, stamp: "From your last log-out, 9 h ago")
    }

    static let live = RaidRootContent.guild(RaidGuildContent(
        night: RaidNightRootContent(
            live: RaidLiveContent(
                eyebrow: "Olympus XXVII · Onyxia’s Lair in progress",
                line: "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage"
            ),
            roles: [
                HomeLine(title: "Tanks 2 of 2", sub: "Grunvald, Brakka"),
                HomeLine(title: "Healers 5 of 8", sub: "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late, 15 min)"),
            ]
        ),
        ready: ready(elixir: false),
        pages: pages
    ))

    static let recruiting = RaidRootContent.noGuild(RecruitingContent(heading: "Guilds recruiting a Fury Warrior", rows: [
        RecruitingRow(guildID: 2, name: "Olympus XXVII", line: "Alliance · Wed and Sat, 8:00 pm CST · needs damage"),
        RecruitingRow(guildID: 7, name: "Dawnbreak", line: "Alliance · Tue and Thu, 9:00 pm EST · needs damage"),
    ]))

    // MARK: Not answered

    static func raider(_ key: String, _ name: String, _ className: String, _ sub: String, role: String, marked: SignupState? = nil) -> NotAnsweredRow {
        // Every one of them lets officers DM them (`officer_dms`), so each row offers Discord DM.
        NotAnsweredRow(
            characterKey: "us/normal/\(key)", name: name, className: className, sub: sub, role: role, marked: marked,
            dm: URL(string: "https://discord.com/users/\(key)")
        )
    }

    static let notAnswered = NotAnsweredContent(
        title: "Onyxia’s Lair", icon: H.onyxia,
        facts: ["Olympus XXVII", "Wed 6 Jan, 8:00 pm CST", "healers short by 3"],
        chips: [
            NotAnsweredChip(filter: .notAnswered, title: "Not answered 4"), NotAnsweredChip(filter: .signedIn, title: "In 19"),
            NotAnsweredChip(filter: .late, title: "Late 3"), NotAnsweredChip(filter: .bench, title: "Bench 2"),
            NotAnsweredChip(filter: .out, title: "Out 3"), NotAnsweredChip(filter: .all, title: "All 31"),
        ],
        filter: .notAnswered,
        mention: MentionContent(
            role: "Mention the 2 healers in Discord", all: "Mention all 4",
            preview: "@Mirethil @Selka: Onyxia’s Lair is Wed 6 Jan, 8:00 pm CST and we need healers. Answer In, Late, Bench or Out in the app or on the raid page."
        ),
        listTitle: "Not answered", listMeta: "4 of 31, healers first",
        rows: [
            raider("mirethil", "Mirethil", "priest", "Holy Priest · was In last week", role: "healer"),
            raider("selka", "Selka", "druid", "Restoration Druid · was Late last week", role: "healer"),
            raider("grunvald", "Grunvald", "warrior", "Protection Warrior · was In last week", role: "tank", marked: .signedIn),
            raider("varro", "Varro", "rogue", "Combat Rogue · missed last week", role: "dps"),
            raider("ysolde", "Ysolde", "mage", "Fire Mage · new this week", role: "dps"),
        ],
        actionable: true,
        benchTitle: "Healers on the bench",
        bench: [raider("aurelith", "Aurelith", "paladin", "Holy Paladin · on the bench", role: "healer")],
        benchNote: "If Mirethil, Selka and Aurelith are In, healers are full."
    )

    // MARK: The applicant

    static func applicant(accepted: Bool) -> ApplicantContent {
        ApplicantContent(
            name: "Mirelle", className: "priest",
            facts: ["Holy Priest · 60", "Alliance Dwarf", "applied 2 min ago"],
            figures: [
                ApplicantFigure(label: "Forever Score", value: "74", unit: nil),
                ApplicantFigure(label: "Sim", value: "612", unit: "HPS"),
                ApplicantFigure(label: "Item level", value: "63", unit: nil),
            ],
            ready: [
                HomeLine(title: "Attuned to Onyxia’s Lair", sub: "Drakefire Amulet"),
                HomeLine(title: "Enchants", sub: "6 of 8 slots · cloak and boots missing"),
                HomeLine(title: "Best in slot gap", sub: "9 pieces from the Holy Priest 60 raid list"),
            ],
            history: [
                HomeLine(title: "Dawnbreak", sub: "14 raids · last raid 2 November"),
                HomeLine(title: "Ashen Vow", sub: "6 raids · last raid 18 October"),
            ],
            note: "Holy main since launch. Wednesday and Sunday nights work for me.",
            accepted: accepted ? "Accepted into Olympus XXVII. Invite Mirelle in game:" : nil,
            invite: "/ginvite Mirelle",
            after: "Mirelle shows on the roster once the addon or Battle.net sees the character in the guild."
        )
    }

    // MARK: The raid page on raid night

    static func night(answered: Bool) -> RaidNightPageContent {
        RaidNightPageContent(
            title: "Onyxia's Lair", icon: H.onyxia, facts: ["Olympus XXVII", "Wed 6 Jan, 8:00 pm CST"],
            eyebrow: "Olympus XXVII · Onyxia's Lair tonight",
            clock: [FSClockPart("00", "hours"), FSClockPart("42", "min")],
            aside: "Wed 6 Jan, 8:00 pm CST · 18 in, 3 late",
            tag: answered ? "You are In" : "Not signed up", answered: answered,
            selected: answered ? .signedIn : nil, lateOpen: false,
            roles: [
                HomeLine(title: "Tanks 2 of 2", sub: "Grunvald, Brakka"),
                HomeLine(title: "Healers 5 of 8", sub: "Mirethil, Selka, Hollowmend, Aenwyn, Brightwen · 3 short"),
                HomeLine(title: "Damage 11 of 30", sub: "Thandrel, Kaelra, Thalrin and 8 more"),
            ],
            ready: ready(elixir: true)
        )
    }
}
