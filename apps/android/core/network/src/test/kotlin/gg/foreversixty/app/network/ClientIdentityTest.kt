package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.client.request.HttpRequestData
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpStatusCode
import io.ktor.http.headersOf
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.builtins.serializer
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test

class ClientIdentityTest {
    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Live)
    private val identity = ClientIdentity(version = "1.2.0", build = 42, installId = "0b6f6b8e-0000-4000-8000-000000000001")
    private val seen = mutableListOf<HttpRequestData>()
    private val gate = UpdateGate()

    private fun client(status: HttpStatusCode, body: String): ApiClient {
        val engine = MockEngine { request ->
            seen += request
            respond(body, status, headersOf(HttpHeaders.ContentType, "application/json"))
        }
        return ApiClient(ForeverHttpClient.create(config, NoToken, engine, identity), gate)
    }

    private fun gateBody(storeUrl: String) =
        """{"ok":false,"data":null,"error":{"code":"update_required","message":"Update the app.",""" +
            """"fields":{"store_url":"$storeUrl"}},"request_id":"r"}"""

    @Test
    fun `X-App reads android version and build`() {
        assertEquals("android/1.2.0 (42)", identity.appHeaderValue)
    }

    @Test
    fun `every request carries X-App and X-App-Install`() = runTest {
        client(HttpStatusCode.OK, """{"ok":true,"data":"x","error":null,"request_id":"r"}""")
            .get("/v1/ping", String.serializer())

        val headers = seen.single().headers
        assertEquals("android/1.2.0 (42)", headers["X-App"])
        assertEquals(identity.installId, headers["X-App-Install"])
    }

    @Test
    fun `a 426 update_required fails as UpdateRequired and shuts the gate`() = runTest {
        val store = "https://play.google.com/store/apps/details?id=gg.foreversixty.app"
        val result = client(HttpStatusCode.UpgradeRequired, gateBody(store)).get("/v1/me", String.serializer())

        assertEquals(ApiResult.Failure(ApiError.UpdateRequired(store)), result)
        assertEquals(UpdateAnswer(store), gate.state.value)
    }

    @Test
    fun `a non-https store link keeps the gate but drops the link`() = runTest {
        client(HttpStatusCode.UpgradeRequired, gateBody("market://details?id=x")).get("/v1/me", String.serializer())

        assertEquals(UpdateAnswer(storeUrl = null), gate.state.value)
    }

    @Test
    fun `another 426 code, or update_required on another status, leaves the gate open`() = runTest {
        val other = """{"ok":false,"data":null,"error":{"code":"upgrade","message":"m"},"request_id":"r"}"""
        client(HttpStatusCode.UpgradeRequired, other).get("/v1/me", String.serializer())
        client(HttpStatusCode.BadRequest, gateBody("https://x.test")).get("/v1/me", String.serializer())

        assertNull(gate.state.value)
    }
}
