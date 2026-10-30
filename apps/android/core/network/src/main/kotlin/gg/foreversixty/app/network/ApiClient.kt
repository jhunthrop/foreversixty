package gg.foreversixty.app.network

import gg.foreversixty.app.model.ApiJson
import io.ktor.client.HttpClient
import io.ktor.client.request.HttpRequestBuilder
import io.ktor.client.request.request
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import io.ktor.client.statement.bodyAsText
import io.ktor.http.ContentType
import io.ktor.http.HttpMethod
import io.ktor.http.contentType
import kotlinx.coroutines.CancellationException
import kotlinx.io.IOException
import kotlinx.serialization.KSerializer
import kotlinx.serialization.SerializationException
import kotlinx.serialization.json.JsonNull

/**
 * The typed door to the API. Every call goes through [call], which unwraps the envelope and
 * turns every failure into an [ApiResult.Failure]; the feature lanes never see an exception.
 */
class ApiClient(private val http: HttpClient) {

    suspend fun <T> get(path: String, response: KSerializer<T>): ApiResult<T> =
        call(HttpMethod.Get, path, response) {}

    suspend fun <B, T> post(path: String, body: B, request: KSerializer<B>, response: KSerializer<T>): ApiResult<T> =
        send(HttpMethod.Post, path, body, request, response)

    suspend fun <B, T> put(path: String, body: B, request: KSerializer<B>, response: KSerializer<T>): ApiResult<T> =
        send(HttpMethod.Put, path, body, request, response)

    suspend fun <T> delete(path: String, response: KSerializer<T>): ApiResult<T> =
        call(HttpMethod.Delete, path, response) {}

    private suspend fun <B, T> send(
        method: HttpMethod,
        path: String,
        body: B,
        request: KSerializer<B>,
        response: KSerializer<T>,
    ): ApiResult<T> = call(method, path, response) {
        contentType(ContentType.Application.Json)
        setBody(ApiJson.encodeToString(request, body))
    }

    private suspend fun <T> call(
        method: HttpMethod,
        path: String,
        response: KSerializer<T>,
        configure: HttpRequestBuilder.() -> Unit,
    ): ApiResult<T> {
        val answer = try {
            http.request(path) {
                this.method = method
                configure()
            }
        } catch (e: CancellationException) {
            throw e
        } catch (e: IOException) {
            return ApiResult.Failure(ApiError.Network(e))
        }
        return decode(answer, response)
    }

    private suspend fun <T> decode(answer: HttpResponse, response: KSerializer<T>): ApiResult<T> {
        val status = answer.status.value
        val envelope = try {
            ApiJson.decodeFromString(ApiEnvelope.serializer(), answer.bodyAsText())
        } catch (e: SerializationException) {
            return ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        } catch (e: IllegalArgumentException) {
            return ApiResult.Failure(ApiError.Decoding(status, requestId = null, cause = e))
        }
        return unwrap(status, envelope, response)
    }

    private fun <T> unwrap(status: Int, envelope: ApiEnvelope, response: KSerializer<T>): ApiResult<T> {
        if (!envelope.ok) {
            val error = envelope.error ?: ApiErrorBody(code = "unknown", message = UNKNOWN_ERROR_MESSAGE)
            return ApiResult.Failure(
                ApiError.Server(status, error.code, error.message, error.fields, envelope.requestId),
            )
        }
        return try {
            val data = ApiJson.decodeFromJsonElement(response, envelope.data ?: JsonNull)
            ApiResult.Success(data, envelope.requestId)
        } catch (e: SerializationException) {
            ApiResult.Failure(ApiError.Decoding(status, envelope.requestId, e))
        } catch (e: IllegalArgumentException) {
            ApiResult.Failure(ApiError.Decoding(status, envelope.requestId, e))
        }
    }

    private companion object {
        const val UNKNOWN_ERROR_MESSAGE = "Something went wrong. Try again in a moment."
    }
}
