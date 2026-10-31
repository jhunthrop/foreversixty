package gg.foreversixty.app.raids.push

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.RaidViewerCharacter
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.SignupState
import kotlinx.serialization.SerializationException
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/**
 * The raid notification's categories and actions (app-v1.md "Push payloads", spec 6.3). The FCM
 * message names the category in `click_action`; the actions use the same ids as iOS.
 */
object RaidPushCategory {
    const val SIGNUP = "RAID_SIGNUP"
    const val ASK_SPEC = "RAID_ASK_SPEC"
    const val CHOOSE = "RAID_CHOOSE"

    const val ACTION_IN = "signup_in"
    const val ACTION_LATE = "signup_late"
    const val ACTION_OUT = "signup_out"
    const val ACTION_ASK_SPEC = "ask_spec"
    const val ACTION_CHOOSE = "choose"

    /** The sign-up state a background action answers, or null for one that opens the app. */
    fun stateFor(action: String): SignupState? = when (action) {
        ACTION_IN -> SignupState.In
        ACTION_LATE -> SignupState.Late
        ACTION_OUT -> SignupState.Out
        else -> null
    }
}

/** A raid push's payload, read from the FCM data map (strings only). */
data class RaidPushPayload(
    val guild: Long,
    val raid: Long,
    val characterKey: String?,
    val askSpec: Boolean,
    val link: String?,
    val answers: Map<String, Answer>,
) {
    data class Answer(val title: String, val body: String)

    /** The screen a tap opens: the payload's app link, else the raid page. */
    fun deepLink(askingSpec: Boolean = false): DeepLink {
        if (!askingSpec) link?.let(DeepLink::parse)?.let { return it }
        return DeepLink.Raid(guild, raid, spec = if (askingSpec) "ask" else null)
    }

    /** The server's answer for an outcome; without one, the app's own words for saved or failed. */
    fun answer(key: String, res: Resources): Answer = answers[key]
        ?: if (key == FAILED) Answer(res.getString(R.string.common_failed), "") else Answer(res.getString(R.string.common_saved), "")

    companion object {
        const val FAILED = "failed"

        fun parse(data: Map<String, String>): RaidPushPayload? {
            val guild = data["guild_id"]?.toLongOrNull() ?: return null
            val raid = data["raid_id"]?.toLongOrNull() ?: return null
            return RaidPushPayload(
                guild = guild,
                raid = raid,
                characterKey = data["character_key"]?.takeIf { it.isNotEmpty() },
                askSpec = data["ask_spec"] == "true",
                link = data["link"],
                answers = data["answers"]?.let(::parseAnswers).orEmpty(),
            )
        }

        private fun parseAnswers(json: String): Map<String, Answer> = try {
            ApiJson.parseToJsonElement(json).jsonObject.mapValues { (_, value) ->
                val answer = value as JsonObject
                Answer(answer.getValue("title").jsonPrimitive.content, answer.getValue("body").jsonPrimitive.content)
            }
        } catch (_: SerializationException) {
            emptyMap()
        } catch (_: IllegalArgumentException) {
            emptyMap()
        } catch (_: NoSuchElementException) {
            emptyMap()
        }
    }
}

/** What an action did: the answer to post in place of the notification, or a screen to open. */
sealed interface RaidPushOutcome {
    data class Answered(val answer: RaidPushPayload.Answer) : RaidPushOutcome
    data class Open(val link: DeepLink) : RaidPushOutcome
}

/**
 * Answers a background In, Late or Out: the PUT with the chosen character, keeping its note (the
 * PUT writes note too), then the server's answer text for the result.
 */
class RaidPushAnswerer(private val service: RaidService, private val res: Resources) {
    suspend fun answer(payload: RaidPushPayload, state: SignupState): RaidPushOutcome {
        val view = when (val result = service.raid(payload.guild, payload.raid)) {
            is ApiResult.Success -> result.data
            is ApiResult.Failure -> return RaidPushOutcome.Answered(payload.answer(RaidPushPayload.FAILED, res))
        }
        val character = character(payload, view) ?: return RaidPushOutcome.Open(payload.deepLink())
        if (state.carriesSpec && character.raidSpec == null && (character.specOptions?.size ?: 0) >= 2) {
            return RaidPushOutcome.Open(payload.deepLink(askingSpec = true))
        }
        val note = view.signups.firstOrNull { it.characterKey == character.characterKey }?.note.orEmpty()
        val request = SignupRequest(state = state.wire, note = note)
        return when (val result = service.signup(payload.guild, payload.raid, character.characterKey, request)) {
            is ApiResult.Success -> {
                val reply = payload.answer(state.wire, res)
                val moved = result.data.moved
                val suffix = payload.answers["moved"]
                RaidPushOutcome.Answered(
                    if (moved != null && suffix != null) reply.copy(body = reply.body + suffix.body.replace("{name}", moved)) else reply,
                )
            }
            is ApiResult.Failure -> RaidPushOutcome.Answered(payload.answer(refusalKey(result.error), res))
        }
    }

    private fun refusalKey(error: ApiError): String = when ((error as? ApiError.Server)?.code) {
        "raid_locked" -> "locked"
        "raid_cancelled" -> "cancelled"
        else -> RaidPushPayload.FAILED
    }

    companion object {
        /** The server's pick, else the character already signed, else the only verified one. */
        fun character(payload: RaidPushPayload, view: RaidView): RaidViewerCharacter? {
            val verified = view.viewer.characters.filter { it.verified }
            payload.characterKey?.let { key -> return verified.firstOrNull { it.characterKey == key } }
            return verified.firstOrNull { it.state != null } ?: verified.singleOrNull()
        }
    }
}
