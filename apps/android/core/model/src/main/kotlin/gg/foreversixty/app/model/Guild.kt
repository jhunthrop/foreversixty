package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * A guild the account belongs to. Mirrors api/internal/auth/store.go `Guild`; the claim,
 * verify, covered and subscribe blocks are left for the guild lane (F4) to model.
 */
@Serializable
data class Guild(
    val id: Long,
    val region: String,
    val ruleset: String,
    val name: String,
    val rank: String? = null,
    val consent: String? = null,
    val verified: Boolean = false,
    @SerialName("verified_count") val verifiedCount: Int = 0,
    val plan: GuildPlan? = null,
    /** What this account may do here ("approve", "remove", "settings", "award", ...). Draw a control exactly when its word is present. */
    val can: List<String> = emptyList(),
    @SerialName("ruleset_label") val rulesetLabel: String? = null,
    val slug: String? = null,
    @SerialName("crest_url") val crestUrl: String? = null,
    val faction: String? = null,
    @SerialName("roster_count") val rosterCount: Int? = null,
)

/** api/internal/auth/me_billing.go `GuildBillingView`, shown only to verified officers. */
@Serializable
data class GuildPlan(
    val scope: String,
    val tier: String,
    @SerialName("you_hold_seat") val youHoldSeat: Boolean = false,
    val status: String,
    @SerialName("current_period_end") val currentPeriodEnd: String? = null,
    @SerialName("cancel_at_period_end") val cancelAtPeriodEnd: Boolean = false,
    @SerialName("billed_by") val billedBy: String,
    @SerialName("you_are_billing_contact") val youAreBillingContact: Boolean = false,
)
