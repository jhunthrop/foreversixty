import Foundation
import Models

/// Which one source pays for the account's Premium; the first match wins. A port of the site's
/// `premiumSource` (web/src/lib/account/plan-line.ts), so the app and the web agree (decision G).
enum PremiumSource: Equatable {
    case own
    case ownThenSeat
    case grant
    case seat
    /// Premium features resolved by the API with no source this read can name.
    case fallback
    case none

    /// The Stripe statuses that still pay for Premium (the site's `LIVE_STATUSES`).
    static let liveStatuses: Set<String> = ["active", "trialing", "past_due"]

    static func of(_ me: Me, now: Date) -> PremiumSource {
        let billing = liveBilling(me)
        if let billing, billing.source != grantSource {
            return me.seat.map { startsInFuture($0, now: now) } == true ? .ownThenSeat : .own
        }
        if billing != nil { return .grant }
        if me.seat != nil { return .seat }
        return me.entitlements.serverSims ? .fallback : .none
    }

    static func liveBilling(_ me: Me) -> Billing? {
        me.entitlements.billing.flatMap { liveStatuses.contains($0.status) ? $0 : nil }
    }

    var isPremium: Bool { self != .none }

    /// The account's own Stripe subscription: Founder and Manage billing belong to these alone.
    var isOwn: Bool { self == .own || self == .ownThenSeat }

    private static let grantSource = "grant"

    private static func startsInFuture(_ seat: MeSeat, now: Date) -> Bool {
        seat.startsAt.map { $0 > now } ?? false
    }
}
