package gg.foreversixty.app.premium

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.AppUser
import gg.foreversixty.app.model.Billing
import gg.foreversixty.app.model.Entitlements
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.MeSeat
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.FixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.InMemoryAppConfigCache
import gg.foreversixty.app.network.NoToken
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * F8's logic, the same list as iOS PremiumTests: which source pays for Premium (the site's
 * `premiumSource`, decision G) and when the app may link to web checkout (decision P, fail closed).
 */
class PremiumTest {
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))

    private fun fixture(name: String): Me {
        val text = requireNotNull(source.read(name)) { "no fixture $name" }
        val data = requireNotNull(ApiJson.parseToJsonElement(text).jsonObject["data"]) { "no data in $name" }
        return ApiJson.decodeFromJsonElement(Me.serializer(), data)
    }

    // The source, as the site resolves it

    @Test fun `a free account is none`() = assertEquals(PremiumSource.None, PremiumSource.of(fixture("me.json"), NOW))

    @Test fun `an own Stripe plan is own`() = assertEquals(PremiumSource.Own, PremiumSource.of(fixture("me.premium.json"), NOW))

    @Test fun `a held seat is seat`() = assertEquals(PremiumSource.Seat, PremiumSource.of(fixture("me.seat.json"), NOW))

    @Test fun `an own plan ending into a future seat is own then seat`() =
        assertEquals(PremiumSource.OwnThenSeat, PremiumSource.of(me(billing = stripe(), seat = seat(NOW.plus(1, ChronoUnit.DAYS))), NOW))

    @Test fun `a seat that has started beside an own plan stays own`() =
        assertEquals(PremiumSource.Own, PremiumSource.of(me(billing = stripe(), seat = seat(NOW.minus(1, ChronoUnit.DAYS))), NOW))

    @Test fun `a grant is grant`() = assertEquals(PremiumSource.Grant, PremiumSource.of(me(billing = grant()), NOW))

    @Test fun `lapsed billing is not Premium`() = listOf("canceled", "incomplete", "unpaid").forEach {
        assertEquals(it, PremiumSource.None, PremiumSource.of(me(billing = stripe(status = it)), NOW))
    }

    @Test fun `live statuses are Premium`() = listOf("active", "trialing", "past_due").forEach {
        assertEquals(it, PremiumSource.Own, PremiumSource.of(me(billing = stripe(status = it)), NOW))
    }

    @Test fun `a lapsed own plan with a seat is seat`() =
        assertEquals(PremiumSource.Seat, PremiumSource.of(me(billing = stripe(status = "canceled"), seat = seat(null)), NOW))

    @Test fun `resolved server sims alone is fallback`() =
        assertEquals(PremiumSource.Fallback, PremiumSource.of(me(serverSims = true), NOW))

    // The plan's words

    @Test fun `a free plan reads Free`() {
        val plan = PlanSummary.of(fixture("me.json"), NOW)
        assertEquals(PlanLine.Free, plan.row)
        assertFalse(plan.isPremium)
        assertFalse(plan.founder)
    }

    @Test fun `an own plan reads the site line with its price on the card and without it on the row`() {
        val plan = PlanSummary.of(fixture("me.premium.json"), NOW)
        assertTrue(plan.founder)
        assertEquals(PlanLine.RenewsPrice(END, PlanPrice("$40", yearly = true)), plan.line)
        assertEquals(PlanLine.Renews(END), plan.row)
    }

    @Test fun `a cancelling plan reads its end`() =
        assertEquals(PlanLine.Ends(END), PlanSummary.of(me(billing = stripe(cancelling = true)), NOW).line)

    @Test fun `own then seat names the seat on the card`() {
        val plan = PlanSummary.of(me(billing = stripe(), seat = seat(NOW.plus(1, ChronoUnit.DAYS))), NOW)
        assertEquals(PlanLine.EndsThenSeat(END, GUILD), plan.line)
        assertEquals(PlanLine.Ends(END), plan.row)
    }

    @Test fun `a seat reads who gave it and has no billing`() {
        val plan = PlanSummary.of(fixture("me.seat.json"), NOW)
        assertEquals(PlanLine.Seat(GUILD), plan.line)
        assertEquals(PlanGiven.Seat("Hunthrop#1894"), plan.line2)
        assertFalse(plan.founder)
        assertEquals(PlanBillingLine.Nothing, plan.billingLine(mayLink = true))
    }

    @Test fun `a grant with and without an end`() {
        val open = PlanSummary.of(me(billing = grant(end = null)), NOW)
        assertEquals(PlanLine.GrantNoEnd, open.line)
        assertEquals(PlanGiven.GrantNoEnd, open.line2)
        val until = PlanSummary.of(me(billing = grant(end = "2027-01-06T15:00:00Z")), NOW)
        assertEquals(PlanLine.GrantUntil(Instant.parse("2027-01-06T15:00:00Z")), until.line)
        assertEquals(PlanGiven.GrantUntil, until.line2)
    }

    @Test fun `prices follow formatCents`() {
        assertEquals(PlanPrice("$6", yearly = false), PlanPrice.of(600, "monthly"))
        assertEquals(PlanPrice("$4.99", yearly = false), PlanPrice.of(499, "monthly"))
        assertEquals(PlanPrice("$50", yearly = true), PlanPrice.of(5000, "yearly"))
        assertNull(PlanPrice.of(5000, "weekly"))
    }

    @Test fun `includes only the true entitlements in the site order`() {
        assertEquals(PlanFeature.entries, PlanSummary.of(fixture("me.premium.json"), NOW).includes)
        assertEquals(listOf(PlanFeature.ServerSims), PlanSummary.of(me(serverSims = true), NOW).includes)
    }

    // See plans, fail closed

    @Test fun `no config means no link`() = assertFalse(WebCheckoutPolicy.mayLink(null, "US", featureOn = true))

    @Test fun `no storefront means no link`() = assertFalse(WebCheckoutPolicy.mayLink(listOf("US"), null, featureOn = true))

    @Test fun `the feature off means no link`() = assertFalse(WebCheckoutPolicy.mayLink(listOf("US"), "US", featureOn = false))

    @Test fun `a US storefront in the regions links`() = assertTrue(WebCheckoutPolicy.mayLink(listOf("US"), "US", featureOn = true))

    @Test fun `other storefronts do not link`() {
        assertFalse(WebCheckoutPolicy.mayLink(listOf("US"), "GB", featureOn = true))
        assertFalse(WebCheckoutPolicy.mayLink(emptyList(), "US", featureOn = true))
    }

    @Test fun `storefront codes read as alpha-2`() {
        assertEquals("US", WebCheckoutPolicy.alpha2("us"))
        listOf("", "USA", "ZZ", "12").forEach { assertNull(it, WebCheckoutPolicy.alpha2(it)) }
    }

    @Test fun `the plans link fails closed before the config endpoint exists`() = runTest {
        val link = PlansLink(repository(FixtureSource { null }), FixedStorefront("US")) { true }
        assertFalse(link.mayLink.value)
        link.resolve()
        assertFalse(link.mayLink.value)
    }

    @Test fun `the plans link follows the config regions`() = runTest {
        listOf("US" to true, "GB" to false, null to false).forEach { (storefront, expected) ->
            val link = PlansLink(repository(source), FixedStorefront(storefront)) { true }
            link.resolve()
            assertEquals("$storefront", expected, link.mayLink.value)
        }
    }

    @Test fun `the plans link stays off with the feature off`() = runTest {
        val link = PlansLink(repository(source), FixedStorefront("US")) { false }
        link.resolve()
        assertFalse(link.mayLink.value)
    }

    @Test fun `the free row links only when allowed and Premium opens the plan`() {
        val free = PlanSummary.of(fixture("me.json"), NOW)
        assertEquals(PlanRowAction.SeePlans, free.rowAction(mayLink = true))
        assertEquals(PlanRowAction.None, free.rowAction(mayLink = false))
        val premium = PlanSummary.of(fixture("me.premium.json"), NOW)
        assertEquals(PlanRowAction.OpenPlan, premium.rowAction(mayLink = false))
    }

    @Test fun `the billing button only in a link region`() {
        val plan = PlanSummary.of(fixture("me.premium.json"), NOW)
        assertEquals(PlanBillingLine.ManageButton, plan.billingLine(mayLink = true))
        assertEquals(PlanBillingLine.ManagedOnWeb, plan.billingLine(mayLink = false))
    }

    private fun repository(fixtures: FixtureSource): AppConfigRepository {
        val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
        return AppConfigRepository(ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(fixtures))), InMemoryAppConfigCache())
    }

    private companion object {
        val NOW: Instant = Instant.parse("2025-11-01T00:00:00Z")
        val END: Instant = Instant.parse("2027-10-27T15:00:00Z")
        const val GUILD = "OLYMPUS XXVII"

        fun stripe(status: String = "active", cancelling: Boolean = false) = Billing(
            plan = "premium", source = "stripe", status = status, currentPeriodEnd = END.toString(),
            cancelAtPeriodEnd = cancelling, amountCents = 5000, interval = "yearly", founder = false,
        )

        fun grant(end: String? = null) =
            Billing(plan = "premium", source = "grant", status = "active", currentPeriodEnd = end, cancelAtPeriodEnd = false)

        fun seat(startsAt: Instant?) =
            MeSeat(guildName = GUILD, givenBy = "Hunthrop#1894", since = NOW.toString(), startsAt = startsAt?.toString())

        fun me(billing: Billing? = null, seat: MeSeat? = null, serverSims: Boolean = false) = Me(
            user = AppUser(id = 1, battletag = "Hunthrop#1894", role = "user", anonymize = false),
            entitlements = Entitlements(
                serverSims = serverSims, retention = false, multiCompare = false, history = false, notifications = false,
                officerViews = false, rosterCheck = false, logCoach = false, supporterMark = false, personalPortal = false,
                billing = billing,
            ),
            discordLinkAvailable = false,
            seat = seat,
        )
    }
}
