import Foundation

/// What the Gear tab draws for a signed-in player (flows spec 4.3; round 5, mocks ios-gear and
/// m-ios-gear-frame): the paperdoll strip (the spec, the set's figure), the line naming the list,
/// every slot, and Where to go tonight. Plain values, so the parity tests draw the mock's scenario
/// directly.
struct GearContent: Equatable {
    /// The list's spec: `Fury Warrior`.
    let title: String
    /// The set's figure (`1004.2`) in `unit`; nil for a tank list, whose set figure is not its score.
    let figure: String?
    let unit: String
    /// `60 raid · Alliance · 4 of 17 slots on the list`, ended by Change.
    let facts: String
    /// Under the line: `+157.6 DPS from 13 upgrades · 2 gains not modelled yet · synced 2 h ago`; nil
    /// without the character's own answer.
    var summary: GearSummary?
    /// The best place of each kind; empty when the gains or the loot file are unknown.
    let tonight: [GearPlaceRow]
    let slots: [GearSlotRow]
}

/// Gear's line under the list: what the answer adds up to, then how old the gear it read is, the age
/// in the warning tone past a week (round 4, row 2).
struct GearSummary: Equatable {
    /// `+157.6 DPS from 13 upgrades · 2 gains not modelled yet`.
    var figures: String?
    /// `synced 2 h ago`.
    var synced: String?
    var stale = false

    /// The line as one string (VoiceOver, tests).
    var line: String { [figures, synced].compactMap(\.self).joined(separator: GearPresenter.separator) }
}

/// Gear with no list to draw (round 2): the spec, the line naming the list with Change (so the player
/// can pick another), and one row: no published list, or `We could not load this list` with Try again.
struct GearNotice: Equatable {
    /// `Beast Mastery Hunter`.
    let title: String
    /// `30 to 39 · Horde`, ended by Change.
    let facts: String
    let row: String
    var sub: String?
    var retry = false
}

/// An item a Gear row opens in the item sheet.
struct GearItemRef: Equatable, Sendable {
    let itemID: Int
    /// The list's slot key (`finger1`).
    let slot: String
    let gain: HomeGain?
}

/// One Where to go tonight row: the place, what it gives, the best item's icon.
struct GearPlaceRow: Equatable, Identifiable {
    let id: String
    let title: String
    let sub: String
    let icon: String?
    let quality: Int?
    /// `2 upgrades · +20.4 DPS`: what the place gives in all.
    var figure: String?
    /// The place's best upgrade, which a tap opens.
    let item: GearItemRef
}

/// One slot of the paperdoll: the list's pick with Equipped, its gain, `Not modelled` or `No gain`.
struct GearSlotRow: Equatable, Identifiable {
    enum Trail: Equatable {
        case worn
        case gain(HomeGain)
        /// The answer has the pick but the sim could not measure it.
        case notModelled
        /// Not worn and no gain: `No gain` (round 8, row 3).
        case noGain
    }

    struct Pick: Equatable {
        let look: HomeItemLook
        let trail: Trail?
        let item: GearItemRef
    }

    let slot: String
    /// The slot's word: `Head`, `Main hand`.
    let label: String
    /// Nil when the list has no pick for the slot.
    let pick: Pick?

    var id: String { slot }
}
