import Foundation
import Models

/// The plan as the You row and the Plan screen (spec F8.1 to F8.4) word it, in the site's
/// words (`accountPlanLine`): one line, an optional second line, Founder, what it includes and
/// whether the account has a billing portal of its own.
struct PlanSummary: Equatable {
    enum Billing: Equatable {
        /// The account's own subscription: Manage billing, where the store allows the link.
        case manage
    }

    /// What tapping the You tab's Plan row does.
    enum RowAction: Equatable {
        /// Opens foreversixty.gg/premium in the browser (link regions only).
        case seePlans
        /// Pushes the Plan screen.
        case openPlan
        /// Free outside a link region: the row is not a button.
        case none
    }

    /// The Plan screen's billing line.
    enum BillingLine: Equatable {
        case manageButton
        case managedOnWeb
        case nothing
    }

    let source: PremiumSource
    /// The You row's sub line: `Free`, or the site's plan line without its price (spec 5.8.1).
    let row: String
    /// The Plan card's line: the site's full plan line, price included.
    let line: String
    let line2: String?
    let founder: Bool
    let includes: [String]
    let billing: Billing?

    var isPremium: Bool { source.isPremium }

    /// The row's one VoiceOver / TalkBack label (spec 5.8.1).
    func rowLabel(mayLink: Bool) -> String {
        if !isPremium { return mayLink ? Copy.planA11ySeePlans : Copy.planA11yFree }
        return founder ? Copy.planA11yRowFounder(line: row) : Copy.planA11yRow(line: row)
    }

    init(me: Me, now: Date) {
        let source = PremiumSource.of(me, now: now)
        let words = Self.words(source, me: me, withPrice: true)
        self.source = source
        row = Self.words(source, me: me, withPrice: false).line
        line = words.line
        line2 = words.line2
        founder = source.isOwn && me.entitlements.billing?.founder == true
        includes = Self.includes(me.entitlements)
        billing = source.isOwn ? .manage : nil
    }

    func rowAction(mayLink: Bool) -> RowAction {
        if isPremium { return .openPlan }
        return mayLink ? .seePlans : .none
    }

    func billingLine(mayLink: Bool) -> BillingLine {
        guard billing == .manage else { return .nothing }
        return mayLink ? .manageButton : .managedOnWeb
    }

    /// The site's `formatCents`: `$50 a year`, `$4.99 a month`; nil for an interval it does not know.
    static func price(cents: Int64, interval: String) -> String? {
        let amount = cents % 100 == 0 ? "$\(cents / 100)" : String(format: "$%lld.%02lld", cents / 100, cents % 100)
        switch interval {
        case "yearly": return Copy.planPriceYearly(amount: amount)
        case "monthly": return Copy.planPriceMonthly(amount: amount)
        default: return nil
        }
    }

    // MARK: The site's words per source

    private typealias Words = (line: String, line2: String?)

    private static func words(_ source: PremiumSource, me: Me, withPrice: Bool) -> Words {
        let billing = PremiumSource.liveBilling(me)
        switch source {
        case .own: return billing.map { ownWords($0, withPrice: withPrice) } ?? (Copy.planPremium, nil)
        case .ownThenSeat: return billing.map { ownThenSeatWords($0, seat: me.seat, withSeat: withPrice) } ?? (Copy.planPremium, nil)
        case .grant: return grantWords(billing?.currentPeriodEnd.flatMap(date))
        case .seat: return me.seat.map { (Copy.planLineSeat(guild: $0.guildName), Copy.planSeatGivenBy(name: $0.givenBy)) } ?? (Copy.planPremium, nil)
        case .fallback: return (Copy.planPremium, nil)
        case .none: return (Copy.youPlanFree, nil)
        }
    }

    private static func ownWords(_ billing: Models.Billing, withPrice: Bool) -> Words {
        guard let end = billing.currentPeriodEnd.flatMap(date) else { return (Copy.planPremium, nil) }
        let day = AccountCopy.longDate(end)
        if billing.cancelAtPeriodEnd { return (Copy.planEnds(date: day), nil) }
        guard withPrice, let cents = billing.amountCents, let interval = billing.interval, let price = price(cents: cents, interval: interval) else {
            return (Copy.planRenews(date: day), nil)
        }
        return (Copy.planLineRenewsPrice(date: day, price: price), nil)
    }

    /// The row says only when the own plan ends; the card adds the seat that follows.
    private static func ownThenSeatWords(_ billing: Models.Billing, seat: MeSeat?, withSeat: Bool) -> Words {
        guard let seat, let end = billing.currentPeriodEnd.flatMap(date) else { return ownWords(billing, withPrice: withSeat) }
        let day = AccountCopy.longDate(end)
        return withSeat ? (Copy.planLineEndsThenSeat(date: day, guild: seat.guildName), nil) : (Copy.planEnds(date: day), nil)
    }

    private static func grantWords(_ end: Date?) -> Words {
        guard let end else { return (Copy.planLineGrantNoEnd, Copy.planGivenGrantNoEnd) }
        return (Copy.planLineGrantUntil(date: AccountCopy.longDate(end)), Copy.planGivenGrantUntil)
    }

    /// Spec F8.4: a line for each Premium entitlement the API says is on, in /premium's words and order.
    private static func includes(_ entitlements: Entitlements) -> [String] {
        [
            (entitlements.logCoach, Copy.planFeatureCoach),
            (entitlements.serverSims, Copy.planFeatureServerSims),
            (entitlements.history, Copy.planFeatureHistory),
            (entitlements.notifications, Copy.planFeatureDiscord),
            (entitlements.supporterMark, Copy.planFeatureSupporter),
        ].compactMap { $0.0 ? $0.1 : nil }
    }

    private static func date(_ iso: String) -> Date? { try? Date(iso, strategy: .iso8601) }
}
