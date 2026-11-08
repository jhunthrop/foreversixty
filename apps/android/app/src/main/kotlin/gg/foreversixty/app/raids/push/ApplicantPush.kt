package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.flows.RaidFlowService

/**
 * guild.applicant (spec 2026-11-08 F3): officers hear of an application. A tap opens the applicant
 * sheet; the one action, Decline, posts the decline in the background. Accept is never offered from
 * the shade.
 */
data class ApplicantPushPayload(val guild: Long, val application: Long) {
    /** The applicant sheet on the Raids tab. */
    val deepLink: DeepLink get() = DeepLink.GuildApplicant(guild, application)

    companion object {
        const val KIND = "guild.applicant"
        const val CATEGORY = "GUILD_APPLICANT"
        const val ACTION_DECLINE = "application_decline"

        private const val KEY_GUILD = "guild_id"
        private const val KEY_APPLICATION = "application_id"

        fun parse(data: Map<String, String>): ApplicantPushPayload? {
            val guild = data[KEY_GUILD]?.toLongOrNull() ?: return null
            val application = data[KEY_APPLICATION]?.toLongOrNull() ?: return null
            return ApplicantPushPayload(guild, application)
        }
    }
}

/** What Decline from the shade did. */
sealed interface ApplicantDeclineOutcome {
    /** The notification's new words. */
    data class Answered(val title: String, val body: String) : ApplicantDeclineOutcome

    /** A refused or failed decline: the applicant sheet opens. */
    data class Open(val link: DeepLink) : ApplicantDeclineOutcome
}

/**
 * Declines from the shade: the POST, then the answer that replaces the notification, naming the
 * applicant and saying when the bot told them on Discord. A decline the server refuses, or that
 * fails, opens the applicant sheet, where the officer sees the application as it stands.
 */
class ApplicantPushDecliner(private val service: RaidFlowService, private val res: Resources) {
    suspend fun decline(payload: ApplicantPushPayload): ApplicantDeclineOutcome =
        when (val result = service.decline(payload.guild, payload.application)) {
            is ApiResult.Success -> {
                val decision = result.data
                val body = if (decision.discordDm.done) res.getString(R.string.push_applicant_declinedDm) else ""
                ApplicantDeclineOutcome.Answered(res.getString(R.string.push_applicant_declined, decision.application.character.name), body)
            }
            is ApiResult.Failure -> failed(payload)
        }

    /** A decline that did not go through: refused, failed or out of time. */
    fun failed(payload: ApplicantPushPayload): ApplicantDeclineOutcome = ApplicantDeclineOutcome.Open(payload.deepLink)
}
