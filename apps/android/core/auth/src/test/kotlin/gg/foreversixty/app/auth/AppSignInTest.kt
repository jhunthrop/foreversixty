package gg.foreversixty.app.auth

import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.FixtureVariant
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpStatusCode
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import java.net.URI
import java.net.URLDecoder

/** The sign-in flow end to end against the contract fixtures, with a fake browser. */
class AppSignInTest {
    private class FakeTokenStore : TokenStore {
        val state = MutableStateFlow<DeviceSession?>(null)
        override val session: Flow<DeviceSession?> = state
        override suspend fun bearerToken(): String? = state.value?.token
        override suspend fun save(session: DeviceSession) { state.value = session }
        override suspend fun clear() { state.value = null }
        override suspend fun discardBearerToken() = clear()
    }

    private class FakeBrowser : BrowserLauncher {
        var opened: String? = null
        override fun open(url: String) { opened = url }

        fun param(name: String): String {
            val query = URI(requireNotNull(opened) { "browser never opened" }).rawQuery
            val raw = query.split('&').first { it.startsWith("$name=") }.substringAfter('=')
            return URLDecoder.decode(raw, Charsets.UTF_8)
        }
    }

    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    private val fixtures = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
    private val tokens = FakeTokenStore()
    private val client = ApiClient(ForeverHttpClient.create(config, tokens, ForeverHttpClient.engineFor(config, fixtures)))
    private val signIn = AppSignIn(config.baseUrl, client, tokens, deviceName = "Pixel 9")
    private val browser = FakeBrowser()

    @Test
    fun `begin opens the start url with an S256 challenge`() {
        signIn.begin(browser)

        assertEquals("S256", browser.param("code_challenge_method"))
        assertEquals(43, browser.param("code_challenge").length)
    }

    @Test
    fun `a good redirect exchanges the code and stores the token`() = runTest {
        signIn.begin(browser)
        val state = browser.param("state")

        val outcome = signIn.complete("${SignInContract.REDIRECT_URI}?code=one-time&state=$state")

        assertTrue("expected signed in, got $outcome", outcome is SignInOutcome.SignedIn)
        val session = (outcome as SignInOutcome.SignedIn).session
        assertTrue(session.token.startsWith("fsd_"))
        assertTrue(session.deviceId.isNotEmpty())
        assertEquals(session.token, tokens.bearerToken())
    }

    @Test
    fun `a spent code answers invalid_grant from the contract fixture`() = runTest {
        val refusing = ApiClient(
            ForeverHttpClient.create(
                config,
                tokens,
                FixtureEngine.create(
                    fixtures,
                    mapOf("app.signin.exchange" to FixtureVariant.Error("invalid-grant", HttpStatusCode.BadRequest)),
                ),
            ),
        )
        val flow = AppSignIn(config.baseUrl, refusing, tokens, deviceName = "Pixel 9")
        flow.begin(browser)

        val outcome = flow.complete("${SignInContract.REDIRECT_URI}?code=spent&state=${browser.param("state")}")

        val error = ((outcome as SignInOutcome.ExchangeFailed).result.error as ApiError.Server)
        assertEquals("invalid_grant", error.code)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `a redirect with a forged state is rejected and stores nothing`() = runTest {
        signIn.begin(browser)

        val outcome = signIn.complete("${SignInContract.REDIRECT_URI}?code=stolen&state=forged")

        assertEquals(SignInOutcome.Rejected, outcome)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `a redirect with no sign-in in progress is rejected`() = runTest {
        assertEquals(SignInOutcome.Rejected, signIn.complete("${SignInContract.REDIRECT_URI}?code=a&state=b"))
    }

    @Test
    fun `a redirect is good for one exchange only`() = runTest {
        signIn.begin(browser)
        val redirect = "${SignInContract.REDIRECT_URI}?code=one-time&state=${browser.param("state")}"
        signIn.complete(redirect)

        assertEquals(SignInOutcome.Rejected, signIn.complete(redirect))
    }

    @Test
    fun `a cancelled sign-in reports the reason`() = runTest {
        signIn.begin(browser)

        val outcome = signIn.complete("${SignInContract.REDIRECT_URI}?error=access_denied")

        assertEquals(SignInOutcome.Cancelled("access_denied"), outcome)
    }

    @Test
    fun `sign out revokes the device on the server, then clears the token`() = runTest {
        tokens.save(DeviceSession("d", "fsd_x"))
        var logout: String? = null
        val recording = ApiClient(
            ForeverHttpClient.create(
                config,
                tokens,
                MockEngine { request ->
                    logout = "${request.method.value} ${request.url.encodedPath} ${request.headers[HttpHeaders.Authorization]}"
                    respond(requireNotNull(fixtures.read("app.logout.json")), HttpStatusCode.OK)
                },
            ),
        )

        AppSignIn(config.baseUrl, recording, tokens, deviceName = "Pixel 9").signOut()

        assertEquals("POST /v1/auth/app/logout Bearer fsd_x", logout)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `sign out clears the token even when the server cannot be reached`() = runTest {
        tokens.save(DeviceSession("d", "fsd_x"))
        val offline = ApiClient(
            ForeverHttpClient.create(config, tokens, MockEngine { respond("", HttpStatusCode.InternalServerError) }),
        )

        AppSignIn(config.baseUrl, offline, tokens, deviceName = "Pixel 9").signOut()

        assertNull(tokens.bearerToken())
    }

    @Test
    fun `a 401 on any call discards the stored token`() = runTest {
        for (code in listOf("unauthorized", "session_expired")) {
            tokens.save(DeviceSession("d", "fsd_dead"))
            val body = """{"ok":false,"data":null,"error":{"code":"$code","message":"Sign in again."},"request_id":"r"}"""
            val dead = ApiClient(
                ForeverHttpClient.create(config, tokens, MockEngine { respond(body, HttpStatusCode.Unauthorized) }),
            )

            val result = AccountApi(dead).me()

            assertTrue(((result as ApiResult.Failure).error as ApiError.Server).isUnauthorized)
            assertNull(code, tokens.bearerToken())
        }
    }
}
