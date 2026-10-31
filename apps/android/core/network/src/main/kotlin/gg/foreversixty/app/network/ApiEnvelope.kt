package gg.foreversixty.app.network

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.JsonElement

/**
 * Every API answer: `{ok, data, error, request_id}` (api/internal/httpx/envelope.go).
 * [data] stays raw JSON here; [ApiClient] decodes it with the caller's serializer, so one
 * envelope type serves every endpoint, including those whose data is null.
 */
@Serializable
data class ApiEnvelope(
    val ok: Boolean,
    val data: JsonElement? = null,
    val error: ApiErrorBody? = null,
    @SerialName("request_id") val requestId: String = "",
)

/** api/internal/httpx/envelope.go `ErrorBody`. [message] is safe to show the player. */
@Serializable
data class ApiErrorBody(
    val code: String,
    val message: String,
    val fields: Map<String, String> = emptyMap(),
)
