import Foundation

/// The one stale rule for a character's synced data (round 4, row 2): past a week the drawer's
/// `Updated … ago`, Home's `Synced … ago` and Gear's `synced … ago` take the warning tone.
enum SyncAge {
    static let staleAfter: TimeInterval = 7 * 86_400

    static func isStale(_ synced: Date, now: Date) -> Bool {
        now.timeIntervalSince(synced) > staleAfter
    }
}
