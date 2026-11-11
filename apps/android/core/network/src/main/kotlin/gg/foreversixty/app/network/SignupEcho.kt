package gg.foreversixty.app.network

import gg.foreversixty.app.model.ApiJson
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.contentOrNull
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/**
 * Fixture mode only: a sign-up write (`PUT …/signups/{region}/{ruleset}/{name}`) answers its canned raid with the tapped
 * answer in place, so the fixture app shows the answer moving (In to Late and back), as the server does. The canned file
 * stays the contract's; only the answering character's `state` and `late_minutes` follow the body.
 */
internal object SignupEcho {
    private const val STATE = "state"
    private const val LATE = "late_minutes"
    private const val KEY = "character_key"
    private val PATH = Regex("""/signups/([^/]+)/([^/]+)/([^/]+)$""")

    /** The canned answer with the request's state on the character's sign-up and viewer row; null when not a sign-up. */
    fun apply(body: String, path: String, request: String): String? {
        val (region, ruleset, name) = PATH.find(path)?.destructured ?: return null
        val sent = runCatching { ApiJson.parseToJsonElement(request).jsonObject }.getOrNull() ?: return null
        val state = sent[STATE]?.jsonPrimitive?.contentOrNull ?: return null
        val root = runCatching { ApiJson.parseToJsonElement(body).jsonObject }.getOrNull() ?: return null
        val data = root["data"] as? JsonObject ?: return null
        val raid = data["raid"] as? JsonObject ?: return null
        val key = "$region/$ruleset/$name"
        val late = sent[LATE] ?: JsonNull

        fun echoed(rows: JsonElement?): JsonElement? = (rows as? JsonArray)?.let { list ->
            JsonArray(
                list.map { row ->
                    val obj = row as? JsonObject
                    if (obj == null || (obj[KEY] as? JsonPrimitive)?.contentOrNull != key) {
                        row
                    } else {
                        JsonObject(obj + (STATE to JsonPrimitive(state)) + (if (LATE in obj) mapOf(LATE to late) else emptyMap()))
                    }
                },
            )
        }

        val viewer = raid["viewer"] as? JsonObject
        val newViewer = viewer?.let { v -> JsonObject(v + ("characters" to (echoed(v["characters"]) ?: v["characters"] ?: JsonNull))) }
        val newRaid = JsonObject(
            raid + ("signups" to (echoed(raid["signups"]) ?: raid["signups"] ?: JsonNull)) +
                (if (newViewer != null) mapOf("viewer" to newViewer) else emptyMap()),
        )
        return JsonObject(root + ("data" to JsonObject(data + ("raid" to newRaid)))).toString()
    }
}
