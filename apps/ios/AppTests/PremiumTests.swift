@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// F8's logic: which source pays for Premium (the site's `premiumSource`, decision G) and when
/// the app may link to web checkout (decision P, fail closed).
@MainActor
@Suite("Premium")
struct PremiumTests {
    // MARK: The source, as the site resolves it

    @Test func freeAccountIsNone() throws {
        #expect(PremiumSource.of(try AccountSnapshotTests.me("me.json"), now: Self.now) == .none)
    }

    @Test func ownStripePlanIsOwn() throws {
        #expect(PremiumSource.of(try AccountSnapshotTests.me("me.premium.json"), now: Self.now) == .own)
    }

    @Test func heldSeatIsSeat() throws {
        #expect(PremiumSource.of(try AccountSnapshotTests.me("me.seat.json"), now: Self.now) == .seat)
    }

    @Test func ownPlanEndingIntoAFutureSeatIsOwnThenSeat() {
        let me = Self.me(billing: Self.stripe(), seat: Self.seat(startsAt: Self.now.addingTimeInterval(86_400)))
        #expect(PremiumSource.of(me, now: Self.now) == .ownThenSeat)
    }

    @Test func seatThatHasStartedBesideOwnPlanStaysOwn() {
        let me = Self.me(billing: Self.stripe(), seat: Self.seat(startsAt: Self.now.addingTimeInterval(-86_400)))
        #expect(PremiumSource.of(me, now: Self.now) == .own)
    }

    @Test func grantIsGrant() {
        #expect(PremiumSource.of(Self.me(billing: Self.grant()), now: Self.now) == .grant)
    }

    @Test(arguments: ["canceled", "incomplete", "unpaid"])
    func lapsedBillingIsNotPremium(status: String) {
        #expect(PremiumSource.of(Self.me(billing: Self.stripe(status: status)), now: Self.now) == .none)
    }

    @Test(arguments: ["active", "trialing", "past_due"])
    func liveStatusesArePremium(status: String) {
        #expect(PremiumSource.of(Self.me(billing: Self.stripe(status: status)), now: Self.now) == .own)
    }

    @Test func lapsedOwnPlanWithASeatIsSeat() {
        let me = Self.me(billing: Self.stripe(status: "canceled"), seat: Self.seat(startsAt: nil))
        #expect(PremiumSource.of(me, now: Self.now) == .seat)
    }

    @Test func resolvedServerSimsAloneIsFallback() {
        #expect(PremiumSource.of(Self.me(serverSims: true), now: Self.now) == .fallback)
    }

    // MARK: The plan's words

    @Test func freePlanReadsFree() throws {
        let plan = PlanSummary(me: try AccountSnapshotTests.me("me.json"), now: Self.now)
        #expect(plan.row == Copy.youPlanFree)
        #expect(!plan.isPremium)
        #expect(!plan.founder)
    }

    @Test func ownPlanReadsTheSiteLineWithItsPrice() throws {
        let plan = PlanSummary(me: try AccountSnapshotTests.me("me.premium.json"), now: Self.now)
        #expect(plan.isPremium)
        #expect(plan.founder)
        #expect(plan.line == Copy.planLineRenewsPrice(date: "27 Oct 2027", price: "$40 a year"))
        #expect(plan.row == Copy.planRenews(date: "27 Oct 2027"))
        #expect(plan.billing == .manage)
    }

    @Test func cancellingPlanReadsItsEnd() {
        let plan = PlanSummary(me: Self.me(billing: Self.stripe(cancelling: true)), now: Self.now)
        #expect(plan.line == Copy.planEnds(date: "27 Oct 2027"))
    }

    @Test func ownThenSeatNamesTheSeat() {
        let me = Self.me(billing: Self.stripe(), seat: Self.seat(startsAt: Self.now.addingTimeInterval(86_400)))
        let plan = PlanSummary(me: me, now: Self.now)
        #expect(plan.line == Copy.planLineEndsThenSeat(date: "27 Oct 2027", guild: "OLYMPUS XXVII"))
        #expect(plan.row == Copy.planEnds(date: "27 Oct 2027"))
        #expect(plan.billing == .manage)
    }

    @Test func seatReadsWhoGaveItAndHasNoBilling() throws {
        let plan = PlanSummary(me: try AccountSnapshotTests.me("me.seat.json"), now: Self.now)
        #expect(plan.line == Copy.planLineSeat(guild: "OLYMPUS XXVII"))
        #expect(plan.line2 == Copy.planSeatGivenBy(name: "Hunthrop#1894"))
        #expect(plan.billing == nil)
        #expect(!plan.founder)
    }

    @Test func grantWithAndWithoutAnEnd() {
        let open = PlanSummary(me: Self.me(billing: Self.grant(end: nil)), now: Self.now)
        #expect(open.line == Copy.planLineGrantNoEnd)
        #expect(open.line2 == Copy.planGivenGrantNoEnd)
        let until = PlanSummary(me: Self.me(billing: Self.grant(end: "2027-01-06T15:00:00Z")), now: Self.now)
        #expect(until.line == Copy.planLineGrantUntil(date: "6 Jan 2027"))
        #expect(until.line2 == Copy.planGivenGrantUntil)
        #expect(until.billing == nil)
    }

    @Test func monthlyPriceWithCents() {
        #expect(PlanSummary.price(cents: 600, interval: "monthly") == "$6 a month")
        #expect(PlanSummary.price(cents: 499, interval: "monthly") == "$4.99 a month")
        #expect(PlanSummary.price(cents: 5000, interval: "yearly") == "$50 a year")
        #expect(PlanSummary.price(cents: 5000, interval: "weekly") == nil)
    }

    @Test func includesOnlyTheTrueEntitlements() throws {
        let premium = PlanSummary(me: try AccountSnapshotTests.me("me.premium.json"), now: Self.now)
        #expect(premium.includes == [Copy.planFeatureCoach, Copy.planFeatureServerSims, Copy.planFeatureHistory, Copy.planFeatureDiscord, Copy.planFeatureSupporter])
        let partial = PlanSummary(me: Self.me(serverSims: true), now: Self.now)
        #expect(partial.includes == [Copy.planFeatureServerSims])
    }

    @Test func rowLabelsReadAsOneElement() throws {
        let free = PlanSummary(me: try AccountSnapshotTests.me("me.json"), now: Self.now)
        #expect(free.rowLabel(mayLink: true) == "Plan, Free, See plans, opens in browser")
        #expect(free.rowLabel(mayLink: false) == "Plan, Free")
        let premium = PlanSummary(me: try AccountSnapshotTests.me("me.premium.json"), now: Self.now)
        #expect(premium.rowLabel(mayLink: false) == "Plan, Premium · renews 27 Oct 2027, Founder")
    }

    // MARK: See plans, fail closed

    @Test func noConfigMeansNoLink() {
        #expect(!WebCheckoutPolicy.mayLink(regions: nil, storefront: "USA", featureOn: true))
    }

    @Test func noStorefrontMeansNoLink() {
        #expect(!WebCheckoutPolicy.mayLink(regions: ["US"], storefront: nil, featureOn: true))
    }

    @Test func featureOffMeansNoLink() {
        #expect(!WebCheckoutPolicy.mayLink(regions: ["US"], storefront: "USA", featureOn: false))
    }

    @Test func usStorefrontInTheRegionsLinks() {
        #expect(WebCheckoutPolicy.mayLink(regions: ["US"], storefront: "USA", featureOn: true))
    }

    @Test func otherStorefrontsDoNotLink() {
        #expect(!WebCheckoutPolicy.mayLink(regions: ["US"], storefront: "GBR", featureOn: true))
        #expect(!WebCheckoutPolicy.mayLink(regions: [], storefront: "USA", featureOn: true))
    }

    @Test(arguments: [("USA", "US"), ("GBR", "GB"), ("deu", "DE"), ("US", "US")])
    func storefrontCodesReadAsAlpha2(code: String, alpha2: String) {
        #expect(WebCheckoutPolicy.alpha2(code) == alpha2)
    }

    @Test(arguments: ["", "ZZZZ", "XKX", "12"])
    func unreadableStorefrontCodesAreNil(code: String) {
        #expect(WebCheckoutPolicy.alpha2(code) == nil)
    }

    /// The store before A3: no config ever answered, so the store has no regions.
    @Test func storeWithoutAConfigNeverLinks() {
        let store = Self.store(FailingTransport())
        #expect(WebCheckoutPolicy.mayLink(regions: store.current?.webCheckoutRegions, storefront: "USA", featureOn: true) == false)
    }

    @Test func unreachableConfigKeepsNoLink() async {
        let store = Self.store(FailingTransport())
        let error = await store.refresh()
        #expect(error != nil)
        #expect(!WebCheckoutPolicy.mayLink(regions: store.current?.webCheckoutRegions, storefront: "USA", featureOn: true))
    }

    @Test func plansLinkFailsClosedBeforeTheConfigEndpointExists() async {
        let link = PlansLink(appConfig: Self.store(FailingTransport()), storefront: FixedStorefront(code: "USA")) { true }
        #expect(!link.mayLink)
        await link.resolve()
        #expect(!link.mayLink)
    }

    @Test(arguments: [("USA", true), ("GBR", false), (nil, false)] as [(String?, Bool)])
    func plansLinkFollowsTheConfigRegions(storefront: String?, expected: Bool) async {
        let link = PlansLink(appConfig: Self.store(FixtureTransport(directory: AccountSnapshotTests.fixtures)), storefront: FixedStorefront(code: storefront)) { true }
        await link.resolve()
        #expect(link.mayLink == expected)
    }

    @Test func plansLinkStaysOffWithTheFeatureOff() async {
        let link = PlansLink(appConfig: Self.store(FixtureTransport(directory: AccountSnapshotTests.fixtures)), storefront: FixedStorefront(code: "USA")) { false }
        await link.resolve()
        #expect(!link.mayLink)
    }

    @Test func freePlanRowLinksOnlyWhenAllowed() throws {
        let free = PlanSummary(me: try AccountSnapshotTests.me("me.json"), now: Self.now)
        #expect(free.rowAction(mayLink: true) == .seePlans)
        #expect(free.rowAction(mayLink: false) == .none)
        let premium = PlanSummary(me: try AccountSnapshotTests.me("me.premium.json"), now: Self.now)
        #expect(premium.rowAction(mayLink: false) == .openPlan)
        #expect(premium.rowAction(mayLink: true) == .openPlan)
    }

    @Test func billingButtonOnlyInALinkRegion() throws {
        let plan = PlanSummary(me: try AccountSnapshotTests.me("me.premium.json"), now: Self.now)
        #expect(plan.billingLine(mayLink: true) == .manageButton)
        #expect(plan.billingLine(mayLink: false) == .managedOnWeb)
        let seat = PlanSummary(me: try AccountSnapshotTests.me("me.seat.json"), now: Self.now)
        #expect(seat.billingLine(mayLink: true) == .nothing)
        #expect(seat.billingLine(mayLink: false) == .nothing)
    }

    // MARK: Helpers

    static let now = Date(timeIntervalSince1970: 1_761_955_200) // 2025-11-01
    static let unreachable = URL(fileURLWithPath: "/")

    static func store(_ transport: any HTTPTransport) -> AppConfigStore {
        AppConfigStore(client: APIClient(baseURL: unreachable, transport: transport), cache: InMemoryAppConfigCache())
    }

    static func stripe(status: String = "active", cancelling: Bool = false) -> Billing {
        Billing(plan: "premium", source: "stripe", status: status, currentPeriodEnd: "2027-10-27T15:00:00Z",
                cancelAtPeriodEnd: cancelling, amountCents: 5000, interval: "yearly", founder: false)
    }

    static func grant(end: String? = nil) -> Billing {
        Billing(plan: "premium", source: "grant", status: "active", currentPeriodEnd: end, cancelAtPeriodEnd: false)
    }

    static func seat(startsAt: Date?) -> MeSeat {
        MeSeat(guildName: "OLYMPUS XXVII", givenBy: "Hunthrop#1894", since: now, startsAt: startsAt)
    }

    static func me(billing: Billing? = nil, seat: MeSeat? = nil, serverSims: Bool = false) -> Me {
        Me(
            user: AppUser(id: 1, battletag: "Hunthrop#1894", role: "user", anonymize: false),
            characters: [],
            guilds: [],
            entitlements: Entitlements(
                serverSims: serverSims, retention: false, multiCompare: false, history: false, notifications: false,
                officerViews: false, rosterCheck: false, logCoach: false, supporterMark: false, personalPortal: false,
                billing: billing
            ),
            finishedSims: [],
            discordLinkAvailable: false,
            discordLink: nil,
            discordLinkSuggested: nil,
            seat: seat
        )
    }
}

/// The API before lane A3: every request fails at the transport.
private struct FailingTransport: HTTPTransport {
    func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        throw .transport("unreachable")
    }
}
