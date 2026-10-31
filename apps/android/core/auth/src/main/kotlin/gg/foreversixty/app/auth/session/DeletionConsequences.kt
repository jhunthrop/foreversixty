package gg.foreversixty.app.auth.session

import gg.foreversixty.app.model.Me

/**
 * What deleting this account ends besides its data, read from `GET /v1/me` (iOS
 * `DeletionConsequence`). The delete screen and its confirmation show exactly the lines that
 * apply, in this order.
 */
sealed interface DeletionConsequence {
    /** An active personal Premium plan is cancelled and does not renew. */
    data object PremiumPlan : DeletionConsequence

    /** The founder price on that plan is gone for good. */
    data object FounderPrice : DeletionConsequence

    /** A guild plan this account pays for is cancelled; its seats end for every member. */
    data class GuildPlan(val guild: String) : DeletionConsequence

    /** The Premium seat this account holds from a guild ends. */
    data class Seat(val guild: String) : DeletionConsequence

    companion object {
        private const val CANCELED = "canceled"
        private const val WHOLE_PLAN = "plan"

        fun of(me: Me): List<DeletionConsequence> = buildList {
            me.entitlements.billing?.takeIf { it.status != CANCELED }?.let { billing ->
                add(PremiumPlan)
                if (billing.founder == true) add(FounderPrice)
            }
            me.guilds.forEach { guild ->
                val plan = guild.plan ?: return@forEach
                if (plan.youAreBillingContact && plan.scope == WHOLE_PLAN && plan.status != CANCELED) add(GuildPlan(guild.name))
            }
            me.seat?.let { add(Seat(it.guildName)) }
        }
    }
}
