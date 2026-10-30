package gg.foreversixty.app.network

import gg.foreversixty.app.model.AppExchangeRequest
import gg.foreversixty.app.model.AppExchangeResponse
import io.ktor.http.HttpMethod
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class FixtureEngineTest {
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    private val client = ApiClient(
        ForeverHttpClient.create(config, { null }, ForeverHttpClient.engineFor(config, source)),
    )

    @Test
    fun `GET maps to the plain file first`() {
        assertEquals(listOf("me.json", "me.get.json"), FixtureRoutes.candidates(HttpMethod.Get, "/v1/me"))
    }

    @Test
    fun `a write maps to the method file first`() {
        assertEquals(
            listOf("auth-app-exchange.post.json", "auth-app-exchange.json"),
            FixtureRoutes.candidates(HttpMethod.Post, "/v1/auth/app/exchange"),
        )
    }

    @Test
    fun `me is served from the contract fixture`() = runTest {
        val result = AccountApi(client).me()

        assertTrue("expected success, got $result", result is ApiResult.Success)
    }

    @Test
    fun `the exchange is served from the contract fixture`() = runTest {
        val result = client.post(
            "/v1/auth/app/exchange",
            AppExchangeRequest("code", "verifier", "android", "Pixel"),
            AppExchangeRequest.serializer(),
            AppExchangeResponse.serializer(),
        )

        val token = (result as ApiResult.Success).data.token
        assertTrue(token.startsWith("fsd_"))
    }

    @Test
    fun `a path with no fixture answers a not found envelope`() = runTest {
        val result = client.get("/v1/nothing/here", AppExchangeResponse.serializer())

        val error = (result as ApiResult.Failure).error as ApiError.Server
        assertEquals(404, error.status)
        assertEquals("not_found", error.code)
    }

    @Test
    fun `a data-only fixture is wrapped in an envelope`() {
        val wrapped = FixtureEngine.asEnvelope("x.json", """{"value":"v"}""")

        assertTrue(wrapped.contains("\"ok\":true"))
        assertTrue(wrapped.contains("\"request_id\":\"fixture-x.json\""))
    }
}
