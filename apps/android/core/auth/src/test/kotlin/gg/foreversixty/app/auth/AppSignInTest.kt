package gg.foreversixty.app.auth

import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
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
        assertTrue(requireNotNull(tokens.bearerToken()).startsWith("fsd_"))
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
    fun `sign out clears the token`() = runTest {
        tokens.save(DeviceSession("d", "fsd_x"))

        signIn.signOut()

        assertNull(tokens.bearerToken())
    }
}
