import Foundation

/// A raid zone's game icon, as the server's `api/internal/zoneicons` names it for Activity rows.
/// The raid and raid-night responses do not carry it yet (asked of the API lane: `zone_icon`);
/// until they do, this mirrors that table, and any other zone takes its generic book.
enum RaidIcons {
    static let generic = "inv_misc_book_09"

    private static let table: [(key: String, name: String, icon: String)] = [
        ("raid:barrow-deeps", "Barrow Deeps", "spell_shadow_raisedead"),
        ("raid:hyjal-summit", "Hyjal Summit", "spell_nature_timestop"),
        ("raid:onyxias-lair", "Onyxia's Lair", "inv_misc_head_dragon_black"),
    ]

    /// The icon of the zone with this key (`raid:onyxias-lair`).
    static func zone(_ key: String) -> String {
        table.first { $0.key == key }?.icon ?? generic
    }

    /// The icon of the zone with this name (`Onyxia's Lair`, either apostrophe).
    static func zoneName(_ name: String) -> String {
        let plain = name.replacingOccurrences(of: "\u{2019}", with: "'")
        return table.first { $0.name == plain }?.icon ?? generic
    }
}
