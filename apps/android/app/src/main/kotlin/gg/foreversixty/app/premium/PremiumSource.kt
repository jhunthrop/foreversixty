package gg.foreversixty.app.premium

import gg.foreversixty.app.model.Billing
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.MeSeat
import java.time.Instant

/**
 * Which one source pays for the account's Premium; the first match wins. A port of the site's
 * `premiumSource` (web/src/lib/account/plan-line.ts), so the app and the web agree (decision G).
 */
enum class PremiumSource {
    Own,
    OwnThenSeat,
    Grant,
    Seat,

    /** Premium features resolved by the API with no source this read can name. */
    Fallback,
    None,
    ;

    val isPremium: Boolean get() = this != None

    /** The account's own Stripe subscription: Founder and Manage billing belong to these alone. */
    val isOwn: Boolean get() = this == Own || this == OwnThenSeat

    companion object {
        /** The Stripe statuses that still pay for Premium (the site's `LIVE_STATUSES`). */
        val LIVE_STATUSES = setOf("active", "trialing", "past_due")
        private const val GRANT_SOURCE = "grant"

        fun of(me: Me, now: Instant): PremiumSource {
            val billing = liveBilling(me)
            val seat = me.seat
            return when {
                billing != null && billing.source != GRANT_SOURCE ->
                    if (seat != null && startsInFuture(seat, now)) OwnThenSeat else Own
                billing != null -> Grant
                seat != null -> Seat
                me.entitlements.serverSims -> Fallback
                else -> None
            }
        }

        fun liveBilling(me: Me): Billing? = me.entitlements.billing?.takeIf { it.status in LIVE_STATUSES }

        private fun startsInFuture(seat: MeSeat, now: Instant): Boolean =
            seat.startsAt?.let { parseInstant(it) }?.isAfter(now) == true
    }
}

/** An API instant, or null when it does not parse. */
internal fun parseInstant(iso: String): Instant? = runCatching { Instant.parse(iso) }.getOrNull()
