import Models

/// The one count model (round 3, row 1): who is coming to a raid (In and Late) by role, Damage as
/// melee plus ranged. Home's raid card, the Raid tab's Who is coming and the short role all read it,
/// so `Damage N` is the same number on every surface and always equals its split.
struct RaidRoleCounts: Equatable, Sendable {
    let tanks: Int
    let healers: Int
    let melee: Int
    let ranged: Int

    var damage: Int { melee + ranged }

    static let tank = "tank"
    static let healer = "healer"
    static let dps = "dps"
    /// The signup's `group` for a raider who fights in melee.
    static let meleeGroup = "melee"

    /// The signups In or Late, counted by role; a damage raider not in melee counts as ranged.
    static func of(_ raid: RaidView) -> RaidRoleCounts {
        let coming = raid.signups.filter { SignupState(rawValue: $0.state)?.isComing == true }
        let damage = coming.filter { group($0.role) == dps }
        let melee = damage.filter { $0.group == meleeGroup }.count
        return RaidRoleCounts(
            tanks: coming.filter { $0.role == tank }.count,
            healers: coming.filter { $0.role == healer }.count,
            melee: melee,
            ranged: damage.count - melee
        )
    }

    /// The count for a role group (`tank`, `healer`, `dps`).
    func count(_ wire: String) -> Int {
        switch wire {
        case Self.tank: tanks
        case Self.healer: healers
        default: damage
        }
    }

    /// The short role, tanks and healers first (a raid short of damage is rarely held up by it).
    func short(_ target: RaidCompTarget) -> (role: String, missing: Int)? {
        let gaps = [(Self.tank, target.tanks - tanks), (Self.healer, target.healers - healers), (Self.dps, target.melee + target.ranged - damage)]
        return gaps.first { $0.1 > 0 }.map { (role: $0.0, missing: $0.1) }
    }

    /// `tank` and `healer` as sent; every other role is damage.
    static func group(_ role: String) -> String {
        switch role {
        case tank, healer: role
        default: dps
        }
    }
}

extension RaidView {
    /// Who has not answered: the raiders list when the API sends it, else the roster's (row 6: Needs
    /// you and Who is coming's foot count this one population).
    var notAnsweredCount: Int { notSignedRaiders?.count ?? notSigned.count }
}
