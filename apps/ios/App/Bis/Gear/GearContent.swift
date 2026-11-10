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
    /// The best place of each kind; empty when the gains or the loot file are unknown.
    let tonight: [GearPlaceRow]
    let slots: [GearSlotRow]
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
    /// The place's best upgrade, which a tap opens.
    let item: GearItemRef
}

/// One slot of the paperdoll: the list's pick with Equipped or its gain, or nothing.
struct GearSlotRow: Equatable, Identifiable {
    enum Trail: Equatable {
        case worn
        case gain(HomeGain)
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
