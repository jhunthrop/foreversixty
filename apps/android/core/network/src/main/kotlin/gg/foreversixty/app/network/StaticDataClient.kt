package gg.foreversixty.app.network

import gg.foreversixty.app.model.ActiveBuild
import gg.foreversixty.app.model.ApiJson
import io.ktor.client.HttpClient
import io.ktor.client.engine.HttpClientEngine
import io.ktor.client.plugins.HttpTimeout
import io.ktor.client.request.get
import io.ktor.client.request.header
import io.ktor.client.statement.bodyAsText
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.isSuccess
import kotlinx.coroutines.CancellationException
import kotlinx.io.IOException
import kotlinx.serialization.SerializationException

/**
 * Reads the site's static data (`https://foreversixty.gg/data/...`, docs/contracts/app-v1.md
 * "Static files"). These files are bare JSON with no envelope, so they never go through
 * [ApiClient], and the device token is never sent to the site. Fixture mode answers them from
 * the `cdn.*` fixtures through the same engine.
 */
class StaticDataClient(
    engine: HttpClientEngine,
    private val siteUrl: String = SITE_URL,
    requestTimeoutMillis: Long = ApiConfig.DEFAULT_REQUEST_TIMEOUT_MILLIS,
    connectTimeoutMillis: Long = ApiConfig.DEFAULT_CONNECT_TIMEOUT_MILLIS,
) {
    private val http = HttpClient(engine) {
        expectSuccess = false
        install(HttpTimeout) {
            this.requestTimeoutMillis = requestTimeoutMillis
            this.connectTimeoutMillis = connectTimeoutMillis
        }
    }

    /** Reads one generated [StaticResource] (`StaticEndpoints`). */
    suspend fun <T> get(resource: StaticResource<T>): ApiResult<T> = when (val answer = fetch(resource)) {
        is ApiResult.Success -> ApiResult.Success(answer.data.value, requestId = "")
        is ApiResult.Failure -> answer
    }

    /** The decoded file and the body it came from, for a caller that keeps the body ([CachedStaticDataClient]). */
    suspend fun <T> fetch(resource: StaticResource<T>): ApiResult<StaticBody<T>> {
        val answer = try {
            http.get(siteUrl.trimEnd('/') + resource.path) {
                header(HttpHeaders.Accept, ContentType.Application.Json.toString())
            }
        } catch (e: CancellationException) {
            throw e
        } catch (e: IOException) {
            return ApiResult.Failure(ApiError.Network(e))
        }
        val status = answer.status.value
        if (!answer.status.isSuccess()) return ApiResult.Failure(ApiError.Status(status))
        val body = answer.bodyAsText()
        return when (val decoded = decode(resource, body, status)) {
            is ApiResult.Success -> ApiResult.Success(StaticBody(decoded.data, body), requestId = "")
            is ApiResult.Failure -> decoded
        }
    }

    /** The build to read the lists and BiS from. */
    suspend fun activeBuild(): ApiResult<ActiveBuild> = get(StaticEndpoints.getActiveBuild())

    companion object {
        /** Decodes a body of [resource], fetched now or saved earlier. */
        fun <T> decode(resource: StaticResource<T>, body: String, status: Int = HTTP_OK): ApiResult<T> = try {
            ApiResult.Success(ApiJson.decodeFromString(resource.response, body), requestId = "")
        } catch (e: SerializationException) {
            ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        } catch (e: IllegalArgumentException) {
            ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        }

        private const val HTTP_OK = 200
        const val SITE_URL = "https://foreversixty.gg"
        const val DATA_PREFIX = "/data/"
    }
}

/** A static file decoded, with the raw body it came from. */
data class StaticBody<T>(val value: T, val body: String)
