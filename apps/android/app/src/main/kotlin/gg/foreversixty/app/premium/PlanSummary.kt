package gg.foreversixty.app.premium

import gg.foreversixty.app.model.Billing
import gg.foreversixty.app.model.Entitlements
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.MeSeat
import java.time.Instant

/** One plan line, in the site's words (`accountPlanCopy`, `premiumSourceCopy`); worded by [PlanCopy]. */
sealed interface PlanLine {
    data object Free : PlanLine
    data object Premium : PlanLine
    data class Renews(val date: Instant) : PlanLine
    data class RenewsPrice(val date: Instant, val price: PlanPrice) : PlanLine
    data class Ends(val date: Instant) : PlanLine
    data class EndsThenSeat(val date: Instant, val guild: String) : PlanLine
    data class GrantUntil(val date: Instant) : PlanLine
    data object GrantNoEnd : PlanLine
    data class Seat(val guild: String) : PlanLine
}

/** The muted second line under a plan that is not billed to the account. */
sealed interface PlanGiven {
    data object GrantUntil : PlanGiven
    data object GrantNoEnd : PlanGiven
    data class Seat(val givenBy: String) : PlanGiven
}

/** The site's `formatCents`: `$50 a year`, `$4.99 a month`. */
data class PlanPrice(val amount: String, val yearly: Boolean) {
    companion object {
        private const val CENTS = 100L

        fun of(cents: Long, interval: String): PlanPrice? {
            val yearly = when (interval) {
                "yearly" -> true
                "monthly" -> false
                else -> return null
            }
            val amount = if (cents % CENTS == 0L) "$${cents / CENTS}" else "$%d.%02d".format(cents / CENTS, cents % CENTS)
            return PlanPrice(amount, yearly)
        }
    }
}

/** The Premium entitlements the Plan screen lists, in /premium's order (spec 5.8.1). */
enum class PlanFeature { Coach, ServerSims, History, Discord, Supporter }

/** What tapping the You tab's Plan row does. */
enum class PlanRowAction { SeePlans, OpenPlan, None }

/** The Plan screen's billing line. */
enum class PlanBillingLine { ManageButton, ManagedOnWeb, Nothing }

/**
 * The plan as the You row and the Plan screen (spec F8.1 to F8.4, 5.8.1) show it: the row line
 * without a price, the card line with it, the muted second line, Founder and what it includes.
 */
data class PlanSummary(
    val source: PremiumSource,
    val row: PlanLine,
    val line: PlanLine,
    val line2: PlanGiven?,
    val founder: Boolean,
    val includes: List<PlanFeature>,
) {
    val isPremium: Boolean get() = source.isPremium

    fun rowAction(mayLink: Boolean): PlanRowAction = when {
        isPremium -> PlanRowAction.OpenPlan
        mayLink -> PlanRowAction.SeePlans
        else -> PlanRowAction.None
    }

    fun billingLine(mayLink: Boolean): PlanBillingLine = when {
        !source.isOwn -> PlanBillingLine.Nothing
        mayLink -> PlanBillingLine.ManageButton
        else -> PlanBillingLine.ManagedOnWeb
    }

    companion object {
        fun of(me: Me, now: Instant): PlanSummary {
            val source = PremiumSource.of(me, now)
            val (line, line2) = words(source, me, withPrice = true)
            return PlanSummary(
                source = source,
                row = words(source, me, withPrice = false).first,
                line = line,
                line2 = line2,
                founder = source.isOwn && me.entitlements.billing?.founder == true,
                includes = includes(me.entitlements),
            )
        }

        private fun words(source: PremiumSource, me: Me, withPrice: Boolean): Pair<PlanLine, PlanGiven?> {
            val billing = PremiumSource.liveBilling(me)
            return when (source) {
                PremiumSource.Own -> billing?.let { own(it, withPrice) } ?: (PlanLine.Premium to null)
                PremiumSource.OwnThenSeat -> billing?.let { ownThenSeat(it, me.seat, withPrice) } ?: (PlanLine.Premium to null)
                PremiumSource.Grant -> grant(billing?.currentPeriodEnd?.let(::parseInstant))
                PremiumSource.Seat -> me.seat?.let { PlanLine.Seat(it.guildName) to PlanGiven.Seat(it.givenBy) } ?: (PlanLine.Premium to null)
                PremiumSource.Fallback -> PlanLine.Premium to null
                PremiumSource.None -> PlanLine.Free to null
            }
        }

        private fun own(billing: Billing, withPrice: Boolean): Pair<PlanLine, PlanGiven?> {
            val end = billing.currentPeriodEnd?.let(::parseInstant) ?: return PlanLine.Premium to null
            if (billing.cancelAtPeriodEnd) return PlanLine.Ends(end) to null
            val cents = billing.amountCents
            val interval = billing.interval
            val price = if (withPrice && cents != null && interval != null) PlanPrice.of(cents, interval) else null
            return (price?.let { PlanLine.RenewsPrice(end, it) } ?: PlanLine.Renews(end)) to null
        }

        /** The row says only when the own plan ends; the card adds the seat that follows. */
        private fun ownThenSeat(billing: Billing, seat: MeSeat?, withSeat: Boolean): Pair<PlanLine, PlanGiven?> {
            val end = billing.currentPeriodEnd?.let(::parseInstant)
            if (seat == null || end == null) return own(billing, withSeat)
            return (if (withSeat) PlanLine.EndsThenSeat(end, seat.guildName) else PlanLine.Ends(end)) to null
        }

        private fun grant(end: Instant?): Pair<PlanLine, PlanGiven?> =
            if (end == null) PlanLine.GrantNoEnd to PlanGiven.GrantNoEnd else PlanLine.GrantUntil(end) to PlanGiven.GrantUntil

        /** A line for each Premium entitlement the API says is on, in /premium's words and order. */
        private fun includes(entitlements: Entitlements): List<PlanFeature> = listOfNotNull(
            PlanFeature.Coach.takeIf { entitlements.logCoach },
            PlanFeature.ServerSims.takeIf { entitlements.serverSims },
            PlanFeature.History.takeIf { entitlements.history },
            PlanFeature.Discord.takeIf { entitlements.notifications },
            PlanFeature.Supporter.takeIf { entitlements.supporterMark },
        )
    }
}
