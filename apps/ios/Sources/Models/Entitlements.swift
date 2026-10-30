import Foundation

/// Every Premium feature pre-resolved for the caller (Go `EntitlementsView`).
/// The app gates a feature on these flags only; it never re-derives a plan.
public struct Entitlements: Codable, Sendable, Equatable {
    public let serverSims: Bool
    public let retention: Bool
    public let multiCompare: Bool
    public let history: Bool
    public let notifications: Bool
    public let officerViews: Bool
    public let rosterCheck: Bool
    public let logCoach: Bool
    public let supporterMark: Bool
    public let billing: Billing?
    public let personalPortal: Bool
}

/// The caller's personal plan (Go `BillingView`). Prices are left out:
/// the app never shows a purchase flow (plan decision P).
public struct Billing: Codable, Sendable, Equatable {
    public let plan: String
    /// "stripe" or "grant".
    public let source: String
    public let status: String
    public let currentPeriodEnd: String?
    public let cancelAtPeriodEnd: Bool
}
