@testable import ForeverSixty
import DesignSystem
import Foundation
import Models

/// The flows mocks' scenario (design/mocks/ws_app_flows.py): Thandrel, Fury Warrior 60 of Olympus
/// XXVII; Onyxia's Lair, Wed 6 Jan, 8:00 pm CST; Felsnap, Beast Mastery Hunter 34. Next upgrades are the
/// round-5 mock's (m-ios-home): Greenhammer unlocks with the first raids, so Timbermaw Tunic is third.
enum HomeFlowTestData {
    static let build = "1.60.1.70790"
    static let onyxia = "inv_misc_head_dragon_black"

    static let thandrel = HomeHeadContent(name: "Thandrel", className: "warrior", facts: ["Level 60 Fury Warrior", "Olympus XXVII"])
    static let felsnap = HomeHeadContent(name: "Felsnap", className: "hunter", facts: ["Level 34 Beast Mastery Hunter", "6 levels to the 40 list"])

    static func item(_ slot: String, _ name: String, _ icon: String, quality: Int = 4, source: String, gain: String) -> HomeUpgradeItem {
        HomeUpgradeItem(
            slot: slot, look: HomeItemLook(id: slot.hashValue & 0xFFFF, name: name, quality: quality, icon: icon, itemLevel: nil),
            source: source, gain: HomeGain(figure: gain, unit: "DPS"), wornID: nil, sourceLine: source
        )
    }

    static let raidRows = [
        item("head", "Lionheart Helm", "inv_helmet_36", source: "Head · Blacksmithing · replaces Helm of Valor", gain: "+14.5"),
        item("finger", "Don Julio’s Band", "inv_jewelry_ring_35", source: "Finger · Stormpike Guard, needs revered", gain: "+6.8"),
        item("chest", "Timbermaw Tunic", "inv_chest_fur", source: "Chest · Leatherworking", gain: "+5.3"),
    ]

    static let levelerRows = [
        item("main", "Felbough Arm", "inv_spear_08", quality: 3, source: "Main hand · Fel Ancient, Ashenvale", gain: "+1.3"),
        item("hands", "Cartilage Shapers", "inv_gauntlets_18", quality: 3, source: "Hands · quest, Lyn the Ignored", gain: "+0.8"),
        item("finger", "Ring of Precision", "inv_jewelry_ring_14", quality: 3, source: "Finger · Blackfathom Deeps", gain: "+0.6"),
    ]

    /// Under Next upgrades' head (round 2): the synced stamp.
    static let stamp = "Synced 2 h ago from the addon"

    static func upgrades(bestRun: Bool) -> HomeBlockContent {
        .upgrades(HomeUpgradesContent(
            rows: raidRows,
            bestRun: bestRun
                ? HomeBestRunRow(place: "Blackrock Spire", line: "Best run · Blackcrow · 8%", icon: "inv_misc_bone_09", gain: HomeGain(figure: "+1.9", unit: "DPS"))
                : nil,
            note: nil, meta: "12 upgrades · +47.5 DPS", stamp: stamp
        ))
    }

    static let tally = [
        RoleLineContent(role: "Tanks", have: 2, want: 2, names: ""),
        RoleLineContent(role: "Healers", have: 5, want: 8, names: ""),
        RoleLineContent(role: "Damage", have: 11, want: 30, names: ""),
    ]

    static func raid(answered: SignupState? = nil, late: Bool = false, officer: Bool = false) -> NextRaidContent {
        let state = late ? .late : answered
        let who = "as Thandrel, Damage"
        let answer = state == nil ? "You have not answered, \(who)" : "As Thandrel, Damage"
        var content = NextRaidContent(
            guildID: 2, raidID: 29, characterKey: "us/normal/thandrel",
            zone: "Onyxia’s Lair",
            place: "Olympus XXVII · Wed 6 Jan, 8:00 pm CST",
            clock: [FSClockPart("02", "hours"), FSClockPart("48", "min")],
            clockLabel: "2 hours, 48 minutes",
            line: HomeLine(
                title: "Onyxia’s Lair · Wed 6 Jan, 8:00 pm CST",
                sub: state == nil ? "18 in, 3 late · you have not answered, \(who)" : "18 in, 3 late · As Thandrel, Damage"
            ),
            icon: onyxia, answer: answer, selected: state, lateOpen: late, tally: tally,
            ready: HomeLine(title: "2 things before you pull", sub: "No flask in your bags; boots not enchanted · from your last log-out, 9 h ago"),
            officer: officer ? OfficerContent(
                unanswered: HomeLine(title: "5 have not answered, 2 of them healers", sub: "Healers short by 3"),
                applicants: HomeLine(title: "1 applicant", sub: "Mirelle, Holy Priest 60, 2 min ago")
            ) : nil
        )
        // The Late card (m-ios-signup) has just saved its 30 minutes.
        content.status = late ? .saved(Copy.commonSaved) : nil
        return content
    }

    static let now = Date(timeIntervalSince1970: 1_793_988_000)

    /// `line` as the mocks write it, `body · when`: the time goes on its own meta line (round 3, row 3).
    static func activity(_ id: Int64, _ icon: String, _ title: String, _ line: String, unread: Bool) -> NewForYouRow {
        var parts = line.components(separatedBy: " · ")
        let when = parts.count > 1 ? parts.removeLast() : ""
        let body = parts.joined(separator: " · ")
        return NewForYouRow(
            row: ActivityRow(
                id: id, kind: "bis.changed", group: "gear", title: title, body: body, link: "https://foreversixty.gg/",
                occurredAt: now, read: !unread, icon: icon
            ),
            line: body, when: when
        )
    }

    static func newForYou(_ kind: String, unread: Int = 3) -> HomeBlockContent {
        let rows: [NewForYouRow] = switch kind {
        case "live": [
                activity(3, onyxia, "Brightwen is Late, 15 min", "Onyxia’s Lair · 4 min ago", unread: true),
                activity(2, onyxia, "Selka is In", "Restoration Druid · 20 min ago", unread: true),
                activity(1, onyxia, "Healers short for Onyxia tonight", "3 of 8 missing · 3 h ago", unread: false),
            ]
        case "leveler": [
                activity(3, "inv_misc_book_09", "Level 34: your 30 to 39 list", "Beast Mastery Hunter · 2 h ago", unread: true),
                activity(2, "inv_misc_pocketwatch_01", "Sim done: 58.4 DPS ±1", "+3.1 vs your last · yesterday", unread: false),
                activity(1, "inv_misc_book_09", "Beast Mastery Hunter moved to A", "Leveling tiers, 30 to 39 · 2 d ago", unread: false),
            ]
        default: [
                activity(3, "ability_warrior_innerrage", "Kaelra passed you in the guild sims", "1,012 DPS, you are #7 of 14 · 1 h ago", unread: true),
                activity(2, "inv_misc_pocketwatch_01", "Sim done: 949 DPS ±6", "+21 vs your last · 3 h ago", unread: true),
                activity(1, "inv_misc_book_09", "Fury Warrior BiS changed", "A new pick for Hands · yesterday", unread: false),
            ]
        }
        return .newForYou(NewForYouContent(rows: rows, unread: unread))
    }

    static let raidLive = HomeBlockContent.raidLive(RaidLiveContent(
        eyebrow: "Olympus XXVII · Onyxia’s Lair in progress",
        line: "Pulled 1 h 12 min ago · 18 in, 3 late · you are In as Thandrel, Damage"
    ))

    static let lastNight = HomeBlockContent.lastNight(LastNightContent(
        loot: [LastNightLoot(
            look: HomeItemLook(id: 19364, name: "Deathbringer", quality: 4, icon: "inv_axe_09", itemLevel: nil),
            line: "You got it from Onyxia · main hand, +12.6 DPS over your worn"
        )],
        report: HomeLine(title: "Onyxia’s Lair report", sub: "87th percentile, your best on Onyxia · the coach has 2 notes"),
        reportIcon: onyxia, reportID: "ewlrx3enxbeb",
        simAgain: HomeLine(title: "Sim again with Deathbringer", sub: "Your sim is from before it dropped")
    ))

    static func content(_ blocks: [HomeBlockContent], head: HomeHeadContent = thandrel) -> HomeFlowContent {
        HomeFlowContent(head: head, phase: .loaded(blocks))
    }
}
