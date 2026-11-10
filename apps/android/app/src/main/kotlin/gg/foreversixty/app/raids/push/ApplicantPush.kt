package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.push.PushAction
import gg.foreversixty.app.push.PushActions

/**
 * The officer's applicant push (`guild.applicant`, category GUILD_APPLICANT): the server sends its
 * buttons in the payload; Decline is a POST the receiver makes, the tap opens the application.
 */
object ApplicantPushCategory {
    const val CATEGORY = "GUILD_APPLICANT"
    const val ACTION_DECLINE = "application_decline"
    private const val POST = "POST"

    /** The buttons the app answers in the background: the payload's known ones, in its order. */
    fun answerable(actions: List<PushAction>): List<PushAction> =
        actions.filter { it.id == ACTION_DECLINE && it.method.equals(POST, ignoreCase = true) }
}

/** An applicant push's payload, read from the FCM data map (strings only). */
data class ApplicantPushPayload(
    val guild: Long,
    val application: Long,
    val title: String,
    val link: String?,
    val actions: List<PushAction>,
) {
    /**
     * The decline call when the payload's [action] names this application's decline route and nothing
     * else; a payload that points elsewhere is not answered in the background.
     */
    fun declineCall(action: PushAction) =
        ApiEndpoints.declineGuildApplication(guild, application).takeIf { it.path == action.path }

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
            )
        }
    }
}

/** Runs one applicant call; the live one sends it through the API client. */
fun interface ApplicantDecliner {
    suspend fun decline(payload: ApplicantPushPayload, action: PushAction): ApiResult<GuildApplicationDecision>?
}

class LiveApplicantDecliner(private val client: ApiClient) : ApplicantDecliner {
    override suspend fun decline(payload: ApplicantPushPayload, action: PushAction) =
        payload.declineCall(action)?.let { client.send(it) }
}

/** Declines from the shade, then words the answer that replaces the notification. */
class ApplicantPushAnswerer(private val decliner: ApplicantDecliner, private val res: Resources) {
    suspend fun answer(payload: ApplicantPushPayload, actionId: String): RaidPushPayload.Answer {
        val action = ApplicantPushCategory.answerable(payload.actions).firstOrNull { it.id == actionId } ?: return failed()
        return when (decliner.decline(payload, action)) {
            is ApiResult.Success -> RaidPushPayload.Answer(res.getString(R.string.push_applicant_declined), payload.title)
            else -> failed()
        }
    }

    fun failed() = RaidPushPayload.Answer(res.getString(R.string.common_failed), res.getString(R.string.push_applicant_failed))
}
