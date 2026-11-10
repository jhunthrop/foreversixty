import Foundation

/// What the Gear tab draws for a signed-in player (flows spec 4.3, mock ios-gear): the list it
/// compares against, Where to go tonight, and every slot. Plain values, so the parity tests draw
/// the mock's scenario directly.
struct GearContent: Equatable {
    /// `Fury Warrior · 60 raid · Alliance` over `4 of 17 slots on the list`.
    let summary: HomeLine
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

/// One Every slot row: the list's pick with Worn or its gain, or why the slot has none.
struct GearSlotRow: Equatable, Identifiable {
    enum Trail: Equatable {
        case worn
        case gain(HomeGain)
    }

    enum Content: Equatable {
        case pick(look: HomeItemLook, source: String, trail: Trail?, item: GearItemRef)
        case empty(String)
    }

    let slot: String
    let content: Content

    var id: String { slot }
}
