import Foundation
import Models

/// What deleting this account ends besides its data, read from `GET /v1/me`. The delete screen
/// shows exactly the lines that apply, in this order. A guild plan the account pays for, or a
/// guild claim it holds, is not a line: the server refuses that deletion (409 guild_obligation)
/// and the screen shows its message.
public enum DeletionConsequence: Equatable, Sendable {
    /// An active personal Premium plan is cancelled and does not renew.
    case premiumPlan
    /// The founder price on that plan is gone for good.
    case founderPrice
    /// The Premium seat this account holds from a guild ends.
    case seat(guild: String)

    private static let canceled = "canceled"

    public static func of(_ me: Me) -> [DeletionConsequence] {
        var lines: [DeletionConsequence] = []
        if let billing = me.entitlements.billing, billing.status != canceled {
            lines.append(.premiumPlan)
            if billing.founder == true { lines.append(.founderPrice) }
        }
        if let seat = me.seat { lines.append(.seat(guild: seat.guildName)) }
        return lines
    }
}
