import ActivityKit
import Foundation
import Models

/// The raid countdown's Live Activity (spec §5.6). The Codable shape is the API's
/// `RaidActivityAttributes` / `RaidActivityState` (fixture surface.live-raid-start.apns): snake_case
/// keys and Unix-second times, because ActivityKit decodes a push's content-state with a plain
/// JSONDecoder. The figures come from the widget snapshot through WidgetFigures.
struct RaidActivityAttributes: ActivityAttributes {
    struct ContentState: Codable, Hashable, Sendable {
        var startsAt: Int64
        var inviteAt: Int64?
        var status: String
        var myState: String?
        var role: String
        var signedUp: Int
        var size: Int?
        /// Who is coming (flows F2: `18 in, 3 late · 2 tanks, 5 heals`). Optional until the
        /// API's content-state sends them; without them the line reads `38 of 40 signed up`.
        var inCount: Int? = nil
        var late: Int? = nil
        var tanks: Int? = nil
        var healers: Int? = nil
        /// The night's summary on the end event (`status` "done"): `1 kill · you got
        /// Deathbringer · report ready`.
        var kills: Int? = nil
        var loot: RaidLiveLoot? = nil
        var reportReady: Bool? = nil

        enum CodingKeys: String, CodingKey {
            case startsAt = "starts_at", inviteAt = "invite_at", status, myState = "my_state", role
            case signedUp = "signed_up", size
            case inCount = "in", late, tanks, healers, kills, loot, reportReady = "report_ready"
        }

        static let done = "done"
    }

    var raidID: String
    var guildID: String
    var guildName: String
    var zoneName: String
    var characterName: String
    /// The instance's icon under /data/{build}/icons; the first raids' when the server sends none.
    var zoneIcon: String? = nil

    enum CodingKeys: String, CodingKey {
        case raidID = "raid_id", guildID = "guild_id", guildName = "guild_name", zoneName = "zone_name"
        case characterName = "character_name", zoneIcon = "zone_icon"
    }

    /// The icon the views draw: the server's, else the first raids' own.
    var icon: String? { zoneIcon ?? RaidZoneIcon.of(zoneName) }
}

/// The item a player got, as the raid's end event names it.
struct RaidLiveLoot: Codable, Hashable, Sendable {
    var name: String
    var quality: Int
    var icon: String
}

/// The first raid tier's icons, as web/src/lib/first-raids.ts `FIRST_RAIDS` draws them (home's
/// countdowns, the guild's next raid). Any other zone has none until the API sends `zone_icon`.
enum RaidZoneIcon {
    static let firstRaids: [String: String] = [
        "Barrow Deeps": "spell_shadow_raisedead",
        "Hyjal Summit": "spell_nature_timestop",
        "Onyxia's Lair": "inv_misc_head_dragon_black",
    ]

    static func of(_ zoneName: String) -> String? { firstRaids[zoneName] }
}

extension RaidActivityAttributes {
    init(_ raid: WidgetNextRaid, characterName: String) {
        self.init(
            raidID: String(raid.raidId), guildID: String(raid.guildId), guildName: raid.guildName,
            zoneName: raid.zoneName, characterName: characterName, zoneIcon: RaidZoneIcon.of(raid.zoneName)
        )
    }
}

extension RaidActivityAttributes.ContentState {
    init(_ raid: WidgetNextRaid, myState: String?) {
        self.init(
            startsAt: Int64(raid.startsAt.timeIntervalSince1970), inviteAt: raid.inviteAt.map { Int64($0.timeIntervalSince1970) },
            status: raid.status, myState: myState, role: raid.my.role, signedUp: raid.signedUp, size: raid.size
        )
    }

    var startDate: Date { Date(timeIntervalSince1970: TimeInterval(startsAt)) }
    var inviteDate: Date? { inviteAt.map { Date(timeIntervalSince1970: TimeInterval($0)) } }
    /// What the countdown counts to: the invites, else the pull.
    var targetDate: Date { inviteDate ?? startDate }
}

/// A server sim's Live Activity (spec §5.6; fixture surface.live-sim-update.apns). Premium only:
/// it exists only for a server run, which needs `server_sims`.
struct SimActivityAttributes: ActivityAttributes {
    struct ContentState: Codable, Hashable, Sendable {
        var state: String
        var done: Int
        var total: Int?
        var etaSec: Int?
        var dps: Double?

        enum CodingKeys: String, CodingKey {
            case state, done, total, etaSec = "eta_sec", dps
        }

        static let running = "running"
        static let finished = "finished"
        static let failed = "failed"
    }

    var simID: String
    var title: String

    enum CodingKeys: String, CodingKey {
        case simID = "sim_id", title
    }
}

/// The sim Live Activity's words, one place for the lock screen and the Island.
struct SimLiveFigures: Equatable, Sendable {
    let title: String
    let status: String
    let detail: String?
    /// `about 2 min left`, beside the progress while running.
    let left: String?
    let fraction: Double?
    /// Finished: `412` drawn large and the rest of `412 DPS · tap for the result` beside it.
    let value: String?
    let valueRest: String?
    let finished: Bool
    let failed: Bool

    init(title: String, state: SimActivityAttributes.ContentState) {
        self.title = title
        switch state.state {
        case SimActivityAttributes.ContentState.finished:
            status = Copy.liveFinished
            let dps = state.dps.map { String(format: "%.0f", $0) }
            detail = dps.map { Copy.liveFinishedSub(dps: $0) }
            value = dps
            valueRest = zip2(dps, detail).map { dps, line in
                line.hasPrefix(dps) ? String(line.dropFirst(dps.count)).trimmingCharacters(in: .whitespaces) : line
            }
            left = nil
            fraction = 1
            finished = true; failed = false
        case SimActivityAttributes.ContentState.failed:
            status = Copy.liveFailed
            detail = Copy.liveOpenRun
            value = nil; valueRest = nil
            left = nil
            fraction = nil
            finished = false; failed = true
        default:
            status = Copy.simsRunning
            value = nil; valueRest = nil
            left = state.etaSec.map { Copy.liveSimLeft(n: max(1, Int((Double($0) / 60).rounded(.up)))) }
            if let total = state.total, total > 0 {
                detail = Copy.liveSim(n: state.done, m: total)
                fraction = min(1, Double(state.done) / Double(total))
            } else {
                detail = nil
                fraction = nil
            }
            finished = false; failed = false
        }
    }
}

private func zip2<A, B>(_ a: A?, _ b: B?) -> (A, B)? {
    guard let a, let b else { return nil }
    return (a, b)
}
