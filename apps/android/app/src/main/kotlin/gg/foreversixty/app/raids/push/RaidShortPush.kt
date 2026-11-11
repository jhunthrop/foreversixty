package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidPingRequest
import gg.foreversixty.app.model.RaidPingResult
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.Endpoint
import gg.foreversixty.app.push.PushAction
import gg.foreversixty.app.push.PushActions
import gg.foreversixty.app.raids.flows.MentionWords
import gg.foreversixty.app.raids.logic.ShortRole
import kotlinx.serialization.json.contentOrNull
import kotlinx.serialization.json.jsonPrimitive

/**
 * The officer's short-role push (`raid.short`, category RAID_SHORT, GO N1): one button, `Mention
 * healers`, posting the raid's Discord ping for the shortest role; the server leaves it out inside the
 * ping's cooldown, without Discord, or with no role short.
 */
object RaidShortPushCategory {
    const val CATEGORY = "RAID_SHORT"
    const val ACTION_MENTION = "raid_mention_short"

    fun answerable(actions: List<PushAction>): List<PushAction> = actions.filter { it.id == ACTION_MENTION && it.isPost }
}

/** A short-role push's payload, read from the FCM data map (strings only). */
data class RaidShortPushPayload(
    val guild: Long,
    val raid: Long,
    val title: String,
    val link: String?,
    val actions: List<PushAction>,
    /** Who the Mention names, in the server's words (`healers`), for the failure line (round 7). */
    val mentionWho: String? = null,
) {
    /** The ping [action] makes, for its role; null when its path is not this raid's ping. */
    fun call(action: PushAction): Endpoint<RaidPingResult>? =
        ApiEndpoints.pingGuildRaid(guild, raid, RaidPingRequest(role(action))).takeIf { it.path == action.path }

    /** The role the button mentions (`healer`), from its body. */
    fun role(action: PushAction): String? = action.body?.get(ROLE)?.jsonPrimitive?.contentOrNull

    companion object {
        private const val ROLE = "role"

        fun parse(data: Map<String, String>): RaidShortPushPayload? {
            val guild = data["guild_id"]?.toLongOrNull() ?: return null
            val raid = data["raid_id"]?.toLongOrNull() ?: return null
            return RaidShortPushPayload(
                guild, raid, data["title"].orEmpty(), data["link"]?.takeIf { it.isNotBlank() }, PushActions.of(data),
                mentionWho = data["mention_who"]?.takeIf { it.isNotBlank() },
            )
        }
    }
}

/** Runs the ping; the live one sends it through the API client. */
fun interface RaidMentioner {
    suspend fun mention(payload: RaidShortPushPayload, action: PushAction): ApiResult<RaidPingResult>?
}

class LiveRaidMentioner(private val client: ApiClient) : RaidMentioner {
    override suspend fun mention(payload: RaidShortPushPayload, action: PushAction) = payload.call(action)?.let { client.send(it) }
}

/** Mentions from the shade, then words what it did (`Mentioned 3 healers in Discord.`) over the raid. */
class RaidShortPushAnswerer(private val mentioner: RaidMentioner, private val res: Resources) {
    suspend fun answer(payload: RaidShortPushPayload, actionId: String): RaidPushPayload.Answer {
        val action = RaidShortPushCategory.answerable(payload.actions).firstOrNull { it.id == actionId } ?: return failed(payload)
        return when (val result = mentioner.mention(payload, action)) {
            is ApiResult.Success -> {
                val people = ShortRole.people(res, payload.role(action), result.data.mentioned).takeIf { payload.role(action) != null }
                RaidPushPayload.Answer(MentionWords(res).line(result.data, people), payload.title)
            }
            else -> failed(payload)
        }
    }

    /**
     * Round 7 row 2: what failed and how to retry, `Could not mention the healers in Discord` over
     * `Onyxia's Lair. Open the raid to try again.`; who in the server's words (`mention_who`).
     */
    fun failed(payload: RaidShortPushPayload): RaidPushPayload.Answer {
        val title = payload.mentionWho?.let { res.getString(R.string.push_short_failed, it) } ?: res.getString(R.string.common_failed)
        return RaidPushPayload.Answer(title, res.getString(R.string.push_short_retry, payload.title))
    }
}
