import Foundation
import Models

/// A sign-up answer, as the API spells it.
enum SignupState: String, CaseIterable, Sendable, Identifiable {
    case signedIn = "in", late, bench, out

    var id: String { rawValue }

    var label: String {
        switch self {
        case .signedIn: Copy.raidStateIn
        case .late: Copy.raidStateLate
        case .bench: Copy.raidStateBench
        case .out: Copy.raidStateOut
        }
    }

    /// In, Late and Bench carry a spec; Out never asks (spec F4.3).
    var carriesSpec: Bool { self != .out }

    /// In and Late draw kill green when checked; Bench and Out draw warm.
    var isComing: Bool { self == .signedIn || self == .late }
}

/// The raid's chip (spec F4.2): its status, or the derived Ended once the end has passed.
enum RaidChip: Equatable, Sendable {
    case planned, locked, cancelled, ended, done

    var label: String {
        switch self {
        case .planned: Copy.raidStatusPlanned
        case .locked: Copy.raidStatusLocked
        case .cancelled: Copy.raidStatusCancelled
        case .ended: Copy.raidStatusEnded
        case .done: Copy.raidStatusDone
        }
    }
}

enum RaidRole {
    static func label(_ role: String) -> String {
        switch role {
        case "tank": Copy.raidRoleTank
        case "healer": Copy.raidRoleHealer
        default: Copy.raidRoleDps
        }
    }
}

/// The game's class order (the roster) and each class's name.
enum GameClass: String, CaseIterable, Sendable {
    case warrior, paladin, hunter, rogue, priest, shaman, mage, warlock, druid

    static func rank(_ slug: String) -> Int {
        allCases.firstIndex { $0.rawValue == slug.lowercased() } ?? allCases.count
    }

    var name: String {
        switch self {
        case .warrior: Copy.raidsClassWarrior
        case .paladin: Copy.raidsClassPaladin
        case .hunter: Copy.raidsClassHunter
        case .rogue: Copy.raidsClassRogue
        case .priest: Copy.raidsClassPriest
        case .shaman: Copy.raidsClassShaman
        case .mage: Copy.raidsClassMage
        case .warlock: Copy.raidsClassWarlock
        case .druid: Copy.raidsClassDruid
        }
    }
}

/// Read-only facts of one raid view at an instant. Nothing here re-derives a rule the API owns:
/// it orders, counts what the view lists, and picks the copy.
struct RaidFacts: Sendable {
    let raid: RaidView
    let now: Date

    var endsAt: Date { raid.startsAt.addingTimeInterval(TimeInterval(raid.durationMinutes * 60)) }

    var chip: RaidChip {
        switch raid.status {
        case "cancelled": .cancelled
        case "done": .done
        default: now >= endsAt ? .ended : (raid.status == "locked" ? .locked : .planned)
        }
    }

    /// Planned or locked, and not ended: a raid the Raids tab offers (spec F4.1, home.next_raid).
    var isSignable: Bool { chip == .planned || chip == .locked }

    var can: Set<String> { Set(raid.viewer.can ?? []) }
    var canLock: Bool { can.contains("lock") }
    var canMove: Bool { can.contains("move") }
    var showsOfficerTools: Bool { canLock || canMove || raid.viewer.toolsEnded == true }

    /// The countdown runs to the invite time when there is one, else to the pull.
    var countdownTarget: Date { raid.inviteAt ?? raid.startsAt }

    /// Signed up counts In plus Late; Bench and Out are not counted.
    var signedUp: Int { raid.counts.signedIn + raid.counts.late }

    var countLine: String {
        raid.size.map { Copy.raidCount(n: signedUp, size: $0) } ?? Copy.raidCountNoSize(n: signedUp)
    }

    /// The bar's fill, 0 to 1; nil (no bar) while the size is unknown.
    var fill: Double? {
        raid.size.flatMap { $0 > 0 ? min(1, Double(signedUp) / Double($0)) : nil }
    }

    /// `3 tanks · 6 healers · 14 DPS · 1 late`; the late clause only when someone is late.
    var breakdown: String { breakdownParts.joined(separator: " · ") }

    var breakdownParts: [String] {
        let counts = raid.counts
        let parts = [Copy.raidsPartTanks(n: counts.tanks), Copy.raidsPartHealers(n: counts.healers), Copy.raidsPartDps(n: counts.dps)]
        return counts.late > 0 ? parts + [Copy.raidsPartLate(n: counts.late)] : parts
    }

    /// `24/40`, beside the roster heading.
    var ratio: String { raid.size.map { "\(signedUp)/\($0)" } ?? "\(signedUp)" }

    /// The planned row's one chip: Cancelled, Locked, Ended, else the viewer's own answer.
    func rowChip(currentKey: String?) -> RowChip {
        switch chip {
        case .cancelled: return .status(.cancelled)
        case .locked: return .status(.locked)
        case .ended, .done: return .status(.ended)
        case .planned: break
        }
        let characters = raid.viewer.characters
        let ordered = characters.filter { $0.characterKey == currentKey } + characters.filter { $0.characterKey != currentKey }
        guard let own = ordered.first(where: { $0.state != nil }), let state = own.state.flatMap(SignupState.init) else {
            return .notSigned
        }
        if state == .bench || state == .out { return .signed(state.label) }
        return .signed(Copy.raidsChipSigned(state: state.label, role: RaidRole.label(own.role)))
    }

    /// The viewer's characters, the current one first, then in the view's order.
    func viewerCharacters(currentKey: String?) -> [RaidViewerCharacter] {
        let characters = raid.viewer.characters
        return characters.filter { $0.characterKey == currentKey } + characters.filter { $0.characterKey != currentKey }
    }
}

enum RowChip: Equatable {
    case status(RaidChip)
    case signed(String)
    case notSigned
}

/// The Raids tab's raids: the next signable one, and the rest of the planned list.
struct PlannedRaids: Sendable {
    let next: RaidView?
    let also: [RaidView]

    init(_ planned: [RaidView], now: Date) {
        let sorted = planned.sorted { $0.startsAt < $1.startsAt }
        let next = sorted.first { RaidFacts(raid: $0, now: now).isSignable }
        self.next = next
        self.also = sorted.filter { $0.id != next?.id }
    }
}
