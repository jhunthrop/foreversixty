import Foundation

/// A raid's zone icon: the server's `zone_icon` where it sends one, else the API's own table
/// (api/internal/zoneicons/zones.json, bundled), else the generic book.
enum RaidIcons {
    static let generic = "inv_misc_book_09"

    private struct Table: Decodable {
        struct Zone: Decodable {
            let key: String
            let icon: String
        }

        let zones: [Zone]
    }

    /// The bundled table by zone key, read once; empty only for a broken build.
    private static let zones: [String: String] = {
        guard let url = Bundle.main.url(forResource: "zones", withExtension: "json"),
              let data = try? Data(contentsOf: url),
              let table = try? JSONDecoder().decode(Table.self, from: data)
        else { return [:] }
        return Dictionary(table.zones.map { ($0.key, $0.icon) }, uniquingKeysWith: { first, _ in first })
    }()

    /// `raid:onyxias-lair` to `inv_misc_head_dragon_black`.
    static func of(zone key: String) -> String {
        zones[key] ?? generic
    }
}
