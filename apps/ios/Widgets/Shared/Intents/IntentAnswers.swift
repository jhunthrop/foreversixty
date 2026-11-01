import Foundation
import Models
import Networking

/// Every sentence Siri, Shortcuts and the Control say (spec §5.7 Voice and Shortcuts), composed
/// from the data the app already has: the raid view, the widget snapshot and the Home answer.
/// Pure: the zone comes in through the clock, so a test pins every sentence verbatim.
struct IntentAnswers: Sendable {
    let clock: WidgetClock

    /// `In`, `Late`, `Out` (and `Bench`, which only an officer sets).
    static func stateWord(_ state: String) -> String {
        WidgetFigures.stateWord(state) ?? state
    }

    /// `Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?`
    func confirm(name: String, state: String, zone: String, startsAt: Date) -> String {
        Copy.intentConfirmSignup(
            name: name, state: Self.stateWord(state), zone: zone,
            weekday: clock.weekday(startsAt), time: clock.hourTime(startsAt)
        )
    }

    /// `You are In for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.`
    func signedUp(state: String, raid: RaidView, characterName: String, role: String) -> String {
        Copy.intentSignedUp(
            state: Self.stateWord(state), zone: raid.zoneName, weekday: clock.weekday(raid.startsAt),
            time: clock.hourTime(raid.startsAt), role: WidgetFigures.roleWord(role), name: characterName,
            n: Self.signedUpCount(raid.counts), size: raid.size ?? Self.signedUpCount(raid.counts)
        )
    }

    /// `Onyxia's Lair, Thursday at 8 pm. You are In. 25 of 40 signed up.`
    func nextRaid(_ raid: WidgetNextRaid) -> String {
        if raid.status == "cancelled" {
            return Copy.intentCancelled(zone: raid.zoneName, weekday: clock.weekday(raid.startsAt))
        }
        let weekday = clock.weekday(raid.startsAt)
        let time = clock.hourTime(raid.startsAt)
        let size = raid.size ?? raid.signedUp
        guard let state = raid.my.state else {
            return Copy.intentNextRaidUnanswered(zone: raid.zoneName, weekday: weekday, time: time, n: raid.signedUp, size: size)
        }
        return Copy.intentNextRaid(
            zone: raid.zoneName, weekday: weekday, time: time, state: Self.stateWord(state), n: raid.signedUp, size: size
        )
    }

    /// `24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.` The roles count In; Late is its
    /// own part and is left out when nobody is late.
    func whoIsComing(_ raid: RaidView) -> String {
        let counts = raid.counts
        var roles = [
            Copy.intentRolesTanks(n: counts.tanks),
            Copy.intentRolesHealers(n: counts.healers),
            Copy.intentRolesDps(n: counts.dps),
        ]
        if counts.late > 0 { roles.append(Copy.intentRolesLate(n: counts.late)) }
        let signedUp = Self.signedUpCount(counts)
        return Copy.intentWhoIsComing(n: signedUp, size: raid.size ?? signedUp, roles: roles.joined(separator: ", "))
    }

    /// In plus Late: the raid page's band count.
    static func signedUpCount(_ counts: RaidCounts) -> Int {
        counts.signedIn + counts.late
    }

    /// The Spotlight raid title: `Onyxia's Lair · Thu 8 pm`.
    func raidTitle(zone: String, startsAt: Date) -> String {
        Copy.spotlightRaidTitle(zone: zone, weekday: clock.shortWeekday(startsAt), time: clock.hourTime(startsAt))
    }

    /// The refusal sentence for a failed sign-up: the app's own (`raid.locked`, `common.failed`),
    /// or the no-connection line.
    static func refusal(_ error: APIError) -> String {
        switch error {
        case .transport: Copy.intentOffline
        case let .server(_, body, _) where body.code == "raid_locked": Copy.raidLocked
        default: Copy.commonFailed
        }
    }
}

/// The Next upgrade snippet's words (board F15.2): the item, `{Slot} · item level {n} · {source}`
/// and `+5.3 DPS for {name}. Replaces {worn}.`
struct UpgradeAnswer: Equatable, Sendable {
    let itemName: String
    let quality: Int?
    let icon: String?
    let meta: String
    let gain: String

    init?(home: CharacterHome, characterName: String) {
        guard let row = home.upgrades.first else { return nil }
        itemName = row.pick.name ?? row.slotLabel
        quality = row.pick.quality
        icon = row.pick.icon
        meta = Self.meta(row)
        gain = Self.gain(row, characterName: characterName)
    }

    /// The spoken answer: the item, then its meta and gain lines.
    var sentence: String { "\(itemName). \(meta). \(gain)" }

    static func meta(_ row: HomeUpgrade) -> String {
        let level = row.pick.itemLevel ?? 0
        let source = row.pick.sourceLine.trimmingCharacters(in: .whitespaces)
        return source.isEmpty
            ? Copy.intentUpgradeMetaNoSource(slot: row.slotLabel, n: level)
            : Copy.intentUpgradeMeta(slot: row.slotLabel, n: level, source: source)
    }

    static func gain(_ row: HomeUpgrade, characterName: String) -> String {
        let worn = row.worn?.label
        guard let gain = row.gain, row.modelled else {
            return Copy.intentUpgradeNoGain(name: characterName, worn: worn ?? row.slotLabel)
        }
        let words = gain.formatted(.number.precision(.fractionLength(1)).locale(Locale(identifier: "en_US_POSIX")))
        guard let worn else { return Copy.intentUpgradeGainEmpty(gain: words, name: characterName) }
        return Copy.intentUpgradeGain(gain: words, name: characterName, worn: worn)
    }
}
