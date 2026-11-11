package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.Endpoint
import gg.foreversixty.app.push.PushAction
import gg.foreversixty.app.push.PushActions

/**
 * The officer's applicant push (`guild.applicant`, category GUILD_APPLICANT): the server sends its
 * buttons in the payload, Accept then Decline or neither (GO N2); each is a POST the receiver makes,
 * the tap opens the application.
 */
object ApplicantPushCategory {
    const val CATEGORY = "GUILD_APPLICANT"
    const val ACTION_ACCEPT = "application_accept"
    const val ACTION_DECLINE = "application_decline"
    private val KNOWN = setOf(ACTION_ACCEPT, ACTION_DECLINE)

    /** The buttons the app answers in the background: the payload's known POSTs, in its order. */
    fun answerable(actions: List<PushAction>): List<PushAction> = actions.filter { it.id in KNOWN && it.isPost }
}

/**
 * An applicant push's payload, read from the FCM data map (strings only); [applicantName] (`Mirelle`)
 * and [applicant] (`Mirelle, Holy Priest 60`) are the server's words for the failure lines (round 7).
 */
data class ApplicantPushPayload(
    val guild: Long,
    val application: Long,
    val title: String,
    val link: String?,
    val actions: List<PushAction>,
    val applicantName: String? = null,
    val applicant: String? = null,
) {
    /**
     * The call [action] makes when it names this application's accept or decline route and nothing
     * else; a payload that points elsewhere is not answered in the background.
     */
    fun call(action: PushAction): Endpoint<GuildApplicationDecision>? = when (action.id) {
        ApplicantPushCategory.ACTION_ACCEPT -> ApiEndpoints.acceptGuildApplication(guild, application)
        ApplicantPushCategory.ACTION_DECLINE -> ApiEndpoints.declineGuildApplication(guild, application)
        else -> null
    }?.takeIf { it.path == action.path }

    companion object {
        fun parse(data: Map<String, String>): ApplicantPushPayload? {
            val guild = data["guild_id"]?.toLongOrNull() ?: return null
            val application = data["application_id"]?.toLongOrNull() ?: return null
            return ApplicantPushPayload(
                guild = guild,
                application = application,
                title = data["title"].orEmpty(),
                link = data["link"]?.takeIf { it.isNotBlank() },
                actions = PushActions.of(data),
                applicantName = data["applicant_name"]?.takeIf { it.isNotBlank() },
                applicant = data["applicant"]?.takeIf { it.isNotBlank() },
            )
        }
    }
}

/** Runs one applicant call; the live one sends it through the API client. */
fun interface ApplicantDecider {
    suspend fun decide(payload: ApplicantPushPayload, action: PushAction): ApiResult<GuildApplicationDecision>?
}

class LiveApplicantDecider(private val client: ApiClient) : ApplicantDecider {
    override suspend fun decide(payload: ApplicantPushPayload, action: PushAction) =
        payload.call(action)?.let { client.send(it) }
}

/** Accepts or declines from the shade, then words the answer that replaces the notification. */
class ApplicantPushAnswerer(private val decider: ApplicantDecider, private val res: Resources) {
    suspend fun answer(payload: ApplicantPushPayload, actionId: String): RaidPushPayload.Answer {
        val action = ApplicantPushCategory.answerable(payload.actions).firstOrNull { it.id == actionId } ?: return failed(payload, actionId)
        val done = if (action.id == ApplicantPushCategory.ACTION_ACCEPT) R.string.push_applicant_accepted else R.string.push_applicant_declined
        return when (decider.decide(payload, action)) {
            is ApiResult.Success -> RaidPushPayload.Answer(res.getString(done), payload.title)
            else -> failed(payload, actionId)
        }
    }

    /**
     * Round 7 row 2: `Could not accept Mirelle` (or decline) over `Mirelle, Holy Priest 60. Open Forever
     * Sixty to try again.`, in the server's words for the applicant; without them, the plain failure.
     */
    fun failed(payload: ApplicantPushPayload, actionId: String): RaidPushPayload.Answer {
        val name = payload.applicantName
        val applicant = payload.applicant
        if (name == null || applicant == null) return RaidPushPayload.Answer(res.getString(R.string.common_failed), payload.title)
        val failed = if (actionId == ApplicantPushCategory.ACTION_DECLINE) R.string.push_applicant_declineFailed else R.string.push_applicant_acceptFailed
        return RaidPushPayload.Answer(res.getString(failed, name), res.getString(R.string.push_applicant_failed, applicant))
    }
}
