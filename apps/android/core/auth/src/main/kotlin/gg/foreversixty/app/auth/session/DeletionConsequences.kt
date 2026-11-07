package gg.foreversixty.app.auth.session

import gg.foreversixty.app.model.Me

/**
 * What deleting this account ends besides its data, read from `GET /v1/me` (iOS
 * `DeletionConsequence`). The delete screen shows exactly the lines that apply, in this order. A
 * guild plan the account pays for, or a guild claim it holds, is not a line: the server refuses
 * that deletion (409 guild_obligation) and the screen shows its message.
 */
sealed interface DeletionConsequence {
    /** An active personal Premium plan is cancelled and does not renew. */
    data object PremiumPlan : DeletionConsequence

    /** The founder price on that plan is gone for good. */
    data object FounderPrice : DeletionConsequence

    /** The Premium seat this account holds from a guild ends. */
    data class Seat(val guild: String) : DeletionConsequence

    companion object {
        private const val CANCELED = "canceled"

        fun of(me: Me): List<DeletionConsequence> = buildList {
            me.entitlements.billing?.takeIf { it.status != CANCELED }?.let { billing ->
                add(PremiumPlan)
                if (billing.founder == true) add(FounderPrice)
            }
            me.seat?.let { add(Seat(it.guildName)) }
        }
    }
}
