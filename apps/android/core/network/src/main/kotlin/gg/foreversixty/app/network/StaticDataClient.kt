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
import kotlinx.serialization.KSerializer
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
    suspend fun <T> get(resource: StaticResource<T>): ApiResult<T> = get(resource.path, resource.response)

    private suspend fun <T> get(path: String, response: KSerializer<T>): ApiResult<T> {
        val answer = try {
            http.get(siteUrl.trimEnd('/') + path) {
                header(HttpHeaders.Accept, ContentType.Application.Json.toString())
            }
        } catch (e: CancellationException) {
            throw e
        } catch (e: IOException) {
            return ApiResult.Failure(ApiError.Network(e))
        }
        val status = answer.status.value
        if (!answer.status.isSuccess()) return ApiResult.Failure(ApiError.Status(status))
        return try {
            ApiResult.Success(ApiJson.decodeFromString(response, answer.bodyAsText()), requestId = "")
        } catch (e: SerializationException) {
            ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        } catch (e: IllegalArgumentException) {
            ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        }
    }

    /** The build to read the lists and BiS from. */
    suspend fun activeBuild(): ApiResult<ActiveBuild> = get(StaticEndpoints.getActiveBuild())

    companion object {
        const val SITE_URL = "https://foreversixty.gg"
        const val DATA_PREFIX = "/data/"
    }
}
