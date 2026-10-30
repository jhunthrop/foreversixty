package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * Every gated feature pre-resolved by the API, plus the caller's personal billing state.
 * Mirrors api/internal/auth/handler.go `EntitlementsView`. The app never re-derives these
 * (one behaviour, one place): it reads the flag.
 */
@Serializable
data class Entitlements(
    @SerialName("server_sims") val serverSims: Boolean = false,
    val retention: Boolean = false,
    @SerialName("multi_compare") val multiCompare: Boolean = false,
    val history: Boolean = false,
    val notifications: Boolean = false,
    @SerialName("officer_views") val officerViews: Boolean = false,
    @SerialName("roster_check") val rosterCheck: Boolean = false,
    @SerialName("log_coach") val logCoach: Boolean = false,
    @SerialName("supporter_mark") val supporterMark: Boolean = false,
    val billing: Billing? = null,
    @SerialName("personal_portal") val personalPortal: Boolean = false,
)

/** api/internal/auth/me_billing.go `BillingView` with catalog.PriceFacts inlined. */
@Serializable
data class Billing(
    val plan: String,
    /** "stripe" or "grant". */
    val source: String,
    val status: String,
    @SerialName("current_period_end") val currentPeriodEnd: String? = null,
    @SerialName("cancel_at_period_end") val cancelAtPeriodEnd: Boolean = false,
    val founder: Boolean = false,
    val interval: String? = null,
    @SerialName("amount_cents") val amountCents: Long? = null,
)
