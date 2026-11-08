package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
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
    /** The notification's new words; [retry] keeps Decline on it. */
    data class Answered(val title: String, val body: String, val retry: Boolean = false) : ApplicantDeclineOutcome

    /** A refusal the shade cannot word (no longer an officer, no such application): the sheet opens. */
    data class Open(val link: DeepLink) : ApplicantDeclineOutcome
}

/**
 * Declines from the shade: the POST, then the answer that replaces the notification. Declined names
 * the applicant and says when the bot told them on Discord; an application another officer already
 * decided says so; a lost connection keeps Decline to try again.
 */
class ApplicantPushDecliner(private val service: RaidFlowService, private val res: Resources) {
    suspend fun decline(payload: ApplicantPushPayload): ApplicantDeclineOutcome =
        when (val result = service.decline(payload.guild, payload.application)) {
            is ApiResult.Success -> {
                val decision = result.data
                val body = if (decision.discordDm.done) res.getString(R.string.push_applicant_declinedDm) else ""
                ApplicantDeclineOutcome.Answered(res.getString(R.string.push_applicant_declined, decision.application.character.name), body)
            }
            is ApiResult.Failure -> refusal(payload, result.error)
        }

    /** The answer when the decline did not finish in time. */
    fun failed(): ApplicantDeclineOutcome.Answered =
        ApplicantDeclineOutcome.Answered(res.getString(R.string.push_applicant_failed), "", retry = true)

    private fun refusal(payload: ApplicantPushPayload, error: ApiError): ApplicantDeclineOutcome = when {
        error !is ApiError.Server -> failed()
        error.code == APPLICATION_CLOSED ->
            ApplicantDeclineOutcome.Answered(res.getString(R.string.push_applicant_decided), res.getString(R.string.push_applicant_decidedLine))
        else -> ApplicantDeclineOutcome.Open(payload.deepLink)
    }

    private companion object {
        /** 409: the application is already accepted, declined or withdrawn. */
        const val APPLICATION_CLOSED = "application_closed"
    }
}
