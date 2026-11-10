package gg.foreversixty.app.push

import gg.foreversixty.app.model.ApiJson
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.SerializationException
import kotlinx.serialization.builtins.ListSerializer

/**
 * One button a push carries in its `actions` (app-v1.md "Push payloads"): the id the apps register,
 * the button's title, and the call a press makes (`POST /v1/guilds/2/applications/31/decline`).
 */
@Serializable
data class PushAction(
    @SerialName("id") val id: String,
    @SerialName("title") val title: String,
    @SerialName("method") val method: String,
    @SerialName("path") val path: String,
)

object PushActions {
    const val KEY_ACTIONS = "actions"

    /** The payload's buttons in its order; none when the key is missing or not the list it should be. */
    fun of(data: Map<String, String>): List<PushAction> {
        val json = data[KEY_ACTIONS]?.takeIf { it.isNotBlank() } ?: return emptyList()
        return try {
            ApiJson.decodeFromString(ListSerializer(PushAction.serializer()), json)
        } catch (_: SerializationException) {
            emptyList()
        } catch (_: IllegalArgumentException) {
            emptyList()
        }
    }
}
