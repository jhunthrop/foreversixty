package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.MockRequestHandleScope
import io.ktor.client.engine.mock.respond
import io.ktor.client.request.HttpRequestData
import io.ktor.client.request.HttpResponseData
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpStatusCode
import io.ktor.http.content.TextContent
import io.ktor.http.headersOf
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.Serializable
import kotlinx.serialization.builtins.serializer
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import java.io.IOException

class ApiClientTest {
    @Serializable
    data class Echo(val value: String)

    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Live)
    private val seen = mutableListOf<HttpRequestData>()

    private fun client(
        token: String? = null,
        answer: suspend MockRequestHandleScope.(HttpRequestData) -> HttpResponseData,
    ): ApiClient {
        val engine = MockEngine { request ->
            seen += request
            answer(request)
        }
        return ApiClient(ForeverHttpClient.create(config, { token }, engine))
    }

    private fun MockRequestHandleScope.json(body: String, status: HttpStatusCode = HttpStatusCode.OK) =
        respond(body, status, headersOf(HttpHeaders.ContentType, "application/json"))

    @Test
    fun `ok envelope yields the data and request id`() = runTest {
        val api = client { json("""{"ok":true,"data":{"value":"hi"},"error":null,"request_id":"r1"}""") }

        val result = api.get("/v1/echo", Echo.serializer())

        assertEquals(ApiResult.Success(Echo("hi"), "r1"), result)
    }

    @Test
    fun `error envelope yields a server error with the API's message`() = runTest {
        val body = """{"ok":false,"data":null,"error":{"code":"not_found","message":"no such raid","fields":{"id":"bad"}},"request_id":"r2"}"""
        val api = client { json(body, HttpStatusCode.NotFound) }

        val error = (api.get("/v1/echo", Echo.serializer()) as ApiResult.Failure).error as ApiError.Server

        assertEquals(404, error.status)
        assertEquals("not_found", error.code)
        assertEquals("no such raid", error.message)
        assertEquals(mapOf("id" to "bad"), error.fields)
        assertEquals("r2", error.requestId)
    }

    @Test
    fun `401 is flagged as unauthorized`() = runTest {
        val body = """{"ok":false,"data":null,"error":{"code":"unauthorized","message":"sign in"},"request_id":"r3"}"""
        val api = client { json(body, HttpStatusCode.Unauthorized) }

        val error = (api.get("/v1/me", Echo.serializer()) as ApiResult.Failure).error as ApiError.Server

        assertTrue(error.isUnauthorized)
    }

    @Test
    fun `a body that is not an envelope is a decoding error`() = runTest {
        val api = client { json("<html>bad gateway</html>", HttpStatusCode.BadGateway) }

        val error = (api.get("/v1/echo", Echo.serializer()) as ApiResult.Failure).error

        assertTrue(error is ApiError.Decoding)
        assertEquals(502, (error as ApiError.Decoding).status)
    }

    @Test
    fun `data of the wrong shape is a decoding error carrying the request id`() = runTest {
        val api = client { json("""{"ok":true,"data":{"other":1},"error":null,"request_id":"r4"}""") }

        val error = (api.get("/v1/echo", Echo.serializer()) as ApiResult.Failure).error as ApiError.Decoding

        assertEquals("r4", error.requestId)
    }

    @Test
    fun `a transport failure is a network error`() = runTest {
        val api = client { throw IOException("offline") }

        val error = (api.get("/v1/echo", Echo.serializer()) as ApiResult.Failure).error

        assertTrue(error is ApiError.Network)
    }

    @Test
    fun `the device token is sent as a bearer header`() = runTest {
        val api = client(token = "fsd_abc") { json("""{"ok":true,"data":"x","request_id":"r"}""") }

        api.get("/v1/me", String.serializer())

        assertEquals("Bearer fsd_abc", seen.single().headers[HttpHeaders.Authorization])
    }

    @Test
    fun `no token means no authorization header`() = runTest {
        val api = client(token = null) { json("""{"ok":true,"data":"x","request_id":"r"}""") }

        api.get("/v1/me", String.serializer())

        assertNull(seen.single().headers[HttpHeaders.Authorization])
    }

    @Test
    fun `requests go to the configured base url`() = runTest {
        val api = client { json("""{"ok":true,"data":"x","request_id":"r"}""") }

        api.get("/v1/me", String.serializer())

        assertEquals("https://api.example.test/v1/me", seen.single().url.toString())
    }

    @Test
    fun `post sends the body as json`() = runTest {
        val api = client { json("""{"ok":true,"data":{"value":"back"},"request_id":"r"}""") }

        api.post("/v1/echo", Echo("out"), Echo.serializer(), Echo.serializer())

        val body = seen.single().body as TextContent
        assertEquals("""{"value":"out"}""", body.text)
        assertEquals("application/json", body.contentType.toString())
    }
}
