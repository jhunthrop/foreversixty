import Foundation
import Models

/// What deleting this account ends besides its data, read from `GET /v1/me`. The delete screen
/// and its confirmation show exactly the lines that apply, in this order.
public enum DeletionConsequence: Equatable, Sendable {
    /// An active personal Premium plan is cancelled and does not renew.
    case premiumPlan
    /// The founder price on that plan is gone for good.
    case founderPrice
    /// A guild plan this account pays for is cancelled; its seats end for every member.
    case guildPlan(guild: String)
    /// The Premium seat this account holds from a guild ends.
    case seat(guild: String)

    private static let canceled = "canceled"
    private static let wholePlan = "plan"

    public static func of(_ me: Me) -> [DeletionConsequence] {
        var lines: [DeletionConsequence] = []
        if let billing = me.entitlements.billing, billing.status != canceled {
            lines.append(.premiumPlan)
            if billing.founder == true { lines.append(.founderPrice) }
        }
        lines += me.guilds.compactMap { guild in
            guard let plan = guild.plan, plan.youAreBillingContact, plan.scope == wholePlan, plan.status != canceled
            else { return nil }
            return .guildPlan(guild: guild.name)
        }
        if let seat = me.seat { lines.append(.seat(guild: seat.guildName)) }
        return lines
    }
}
