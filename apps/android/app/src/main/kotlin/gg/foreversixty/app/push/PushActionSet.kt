package gg.foreversixty.app.push

import gg.foreversixty.app.raids.push.ApplicantPushPayload
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.raids.push.RaidPushPayload

/**
 * The buttons a push draws, chosen by its category alone (app-v1.md "Push payloads"): the raid
 * answers for a raid category carrying its raid, Decline for GUILD_APPLICANT carrying its
 * application, nothing for any other category or a payload missing its ids. A push without buttons
 * still opens its link.
 */
sealed interface PushActionSet {
    data class Raid(val category: String) : PushActionSet
    data class Applicant(val payload: ApplicantPushPayload) : PushActionSet
    data object None : PushActionSet

    companion object {
        fun of(message: PushMessage): PushActionSet {
            val category = message.category ?: return None
            return when {
                category == ApplicantPushPayload.CATEGORY -> ApplicantPushPayload.parse(message.data)?.let(::Applicant) ?: None
                category in RaidPushCategory.ALL && RaidPushPayload.parse(message.data) != null -> Raid(category)
                else -> None
            }
        }
    }
}
