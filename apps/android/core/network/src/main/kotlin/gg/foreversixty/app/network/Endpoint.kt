package gg.foreversixty.app.network

import io.ktor.http.HttpMethod
import kotlinx.serialization.KSerializer

/**
 * One API call: its route, query, JSON body and the type its envelope's data decodes to. The
 * members are generated from api/openapi.yaml (`ApiEndpoints`, web/scripts/gen-mobile-api.mjs);
 * [ApiClient.send] runs one.
 */
class Endpoint<T>(
    val method: HttpMethod,
    /** Absolute API path, e.g. "/v1/me", with its segments already encoded. */
    val path: String,
    val response: KSerializer<T>,
    val query: List<Pair<String, String>> = emptyList(),
    /** The request body, already encoded as the API's JSON. */
    val body: String? = null,
) {
    init {
        require(path.startsWith("/")) { "an endpoint path is absolute: $path" }
    }
}

/** One static file of the site and its type (`StaticEndpoints`); [StaticDataClient.get] reads it. */
class StaticResource<T>(val path: String, val response: KSerializer<T>) {
    init {
        require(path.startsWith(StaticDataClient.DATA_PREFIX)) { "static paths live under /data/: $path" }
    }
}
