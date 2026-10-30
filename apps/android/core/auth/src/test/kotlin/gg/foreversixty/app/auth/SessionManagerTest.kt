package gg.foreversixty.app.auth

import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.CONFIRMED_DELETION
import gg.foreversixty.app.auth.session.CharacterSelection
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.FixtureBrowser
import gg.foreversixty.app.auth.session.MemoryValue
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.auth.session.SessionTokens
import gg.foreversixty.app.auth.session.SignInFailure
import gg.foreversixty.app.auth.session.SignedOutNotice
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.model.AccountDeletionRequest
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.FixtureVariant
import gg.foreversixty.app.network.ForeverHttpClient
import io.ktor.client.engine.HttpClientEngine
import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpStatusCode
import io.ktor.http.headersOf
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.withContext
import kotlinx.coroutines.withTimeout
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.test.TestScope
import kotlinx.coroutines.test.advanceUntilIdle
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import java.time.Instant

@OptIn(kotlinx.coroutines.ExperimentalCoroutinesApi::class)
/** F1's test list, the same list as iOS `SessionStoreTests` and `CharacterSelectionTests`. */
class SessionManagerTest {
    private class FakeTokenStore(token: String? = null) : TokenStore {
        val state = MutableStateFlow(token?.let { DeviceSession("dvc_8Qk2nXw4", it) })
        override val session: Flow<DeviceSession?> = state
        override suspend fun bearerToken(): String? = state.value?.token
        override suspend fun save(session: DeviceSession) { state.value = session }
        override suspend fun clear() { state.value = null }
        override suspend fun discardBearerToken() = clear()
    }

    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    private val fixtures = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))

    private fun fixture(name: String): String = requireNotNull(fixtures.read(name)) { "no fixture $name" }

    private fun answering(status: HttpStatusCode, body: String): HttpClientEngine = MockEngine {
        respond(body, status, headersOf(HttpHeaders.ContentType, "application/json"))
    }

    private fun TestScope.manager(
        token: String? = null,
        engine: HttpClientEngine = FixtureEngine.create(fixtures),
    ): Pair<SessionManager, SessionTokens> {
        val tokens = SessionTokens(FakeTokenStore(token))
        val client = ApiClient(ForeverHttpClient.create(config, tokens, engine))
        val signIn = AppSignIn(config.baseUrl, client, tokens, deviceName = "Pixel 9")
        return SessionManager(client, tokens, signIn, MemoryValue(), backgroundScope) to tokens
    }

    /** The exchange runs on Ktor's own threads, so wait in real time for the phase to settle. */
    private suspend fun SessionManager.awaitPhase(settled: (SessionPhase) -> Boolean): SessionPhase =
        withContext(Dispatchers.Default) { withTimeout(5_000) { phase.first(settled) } }

    @Test
    fun `starts signed out with no token and signed in with one`() = runTest {
        val (out, _) = manager()
        out.start()
        assertEquals(SessionPhase.SignedOut(), out.phase.value)
        val (signedIn, _) = manager(token = "fsd_x")
        signedIn.start()
        assertEquals(SessionPhase.SignedIn(null), signedIn.phase.value)
    }

    @Test
    fun `sign-in exchanges the code, keeps the device id and imports the account`() = runTest {
        val (session, tokens) = manager()
        session.beginSignIn(FixtureBrowser { session.handleRedirect(it) })
        session.awaitPhase { it !is SessionPhase.Importing }
        val me = requireNotNull(session.phase.value.me) { "signed in with no account" }
        assertTrue(me.characters.isNotEmpty())
        assertEquals("dvc_8Qk2nXw4", session.thisDeviceId())
        assertNotNull(tokens.bearerToken())
    }

    @Test
    fun `a spent code shows the expired line and keeps no token`() = runTest {
        val engine = FixtureEngine.create(
            fixtures,
            mapOf("app.signin.exchange" to FixtureVariant.Error("invalid-grant", HttpStatusCode.BadRequest)),
        )
        val (session, tokens) = manager(engine = engine)
        session.beginSignIn(FixtureBrowser { session.handleRedirect(it) })
        session.awaitPhase { it !is SessionPhase.Importing }
        assertEquals(SessionPhase.SignedOut(SignedOutNotice.SignInFailed(SignInFailure.Expired)), session.phase.value)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `closing the Custom Tab returns to signed out with no line`() = runTest {
        val (session, _) = manager()
        session.beginSignIn { }
        assertEquals(SessionPhase.SigningIn, session.phase.value)
        session.onReturnedToApp()
        assertEquals(SessionPhase.SignedOut(), session.phase.value)
    }

    @Test
    fun `sign-in failures map to the spec's three lines`() {
        val server = { code: String -> ApiError.Server(400, code, "x", emptyMap(), "r") }
        assertEquals(SignInFailure.BattleNet, SignInFailure.of(SignInOutcome.Rejected))
        assertEquals(SignInFailure.BattleNet, SignInFailure.of(SignInOutcome.Cancelled("access_denied")))
        assertEquals(SignInFailure.Offline, SignInFailure.of(SignInOutcome.ExchangeFailed(ApiResult.Failure(ApiError.Network(java.io.IOException())))))
        assertEquals(SignInFailure.Expired, SignInFailure.of(SignInOutcome.ExchangeFailed(ApiResult.Failure(server("invalid_grant")))))
        assertEquals(SignInFailure.BattleNet, SignInFailure.of(SignInOutcome.ExchangeFailed(ApiResult.Failure(server("invalid")))))
    }

    @Test
    fun `401 unauthorized and 401 session_expired both clear the token and say so`() = runTest {
        for (file in listOf("error-unauthorized.json", "error-session-expired.json")) {
            val (session, tokens) = manager(token = "fsd_dead", engine = answering(HttpStatusCode.Unauthorized, fixture(file)))
            session.start()
            session.refresh()
            assertEquals(file, SessionPhase.SignedOut(SignedOutNotice.SessionEnded), session.phase.value)
            assertNull(tokens.bearerToken())
        }
    }

    @Test
    fun `a 401 on any other request reaches the session through the token source`() = runTest {
        val (session, tokens) = manager(
            token = "fsd_dead",
            engine = answering(HttpStatusCode.Unauthorized, fixture("error-session-expired.json")),
        )
        session.start()
        backgroundScope.launch { session.watchForDeadTokens() }
        advanceUntilIdle()
        session.devices()
        session.awaitPhase { it is SessionPhase.SignedOut }
        assertEquals(SessionPhase.SignedOut(SignedOutNotice.SessionEnded), session.phase.value)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `sign out forgets the token and shows no line`() = runTest {
        val (session, tokens) = manager(token = "fsd_x")
        session.start()
        session.signOut()
        assertEquals(SessionPhase.SignedOut(), session.phase.value)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `this device is listed first and cannot be revoked`() = runTest {
        val (session, _) = manager(token = "fsd_x")
        val rows = (session.devices(Instant.parse("2026-10-30T21:41:00Z")) as ApiResult.Success).data
        assertTrue(rows.first().isThisDevice)
        assertFalse(rows.first().canRevoke)
        assertTrue(rows.drop(1).all { it.canRevoke })
        assertEquals(gg.foreversixty.app.auth.session.DeviceRow.Meta.LastUsed(RelativeAge.Hours(2)), rows[1].meta)
    }

    @Test
    fun `revoke answers done, or failed for a device that is not yours`() = runTest {
        assertEquals(AccountActionOutcome.Done, manager(token = "fsd_x").first.revoke("dvc_3Hm7pLa9"))
        val missing = manager(token = "fsd_x", engine = answering(HttpStatusCode.NotFound, fixture("devices.revoke.error-not-found.json")))
        assertEquals(AccountActionOutcome.Failed, missing.first.revoke("dvc_nope"))
    }

    @Test
    fun `deleting the account signs out with the date the data goes`() = runTest {
        val (session, tokens) = manager(token = "fsd_x")
        session.start()
        assertEquals(AccountActionOutcome.Done, session.deleteAccount())
        val expected = SignedOutNotice.AccountDeleted(Instant.parse("2026-11-29T21:41:00Z"))
        assertEquals(SessionPhase.SignedOut(expected), session.phase.value)
        assertNull(tokens.bearerToken())
    }

    @Test
    fun `the deletion sends the confirmation phrase`() {
        assertEquals(
            """{"confirm":"delete my account"}""",
            ApiJson.encodeToString(AccountDeletionRequest.serializer(), CONFIRMED_DELETION),
        )
    }

    @Test
    fun `a refused deletion keeps the session`() = runTest {
        val (session, _) = manager(token = "fsd_x", engine = answering(HttpStatusCode.BadRequest, fixture("me.delete.error-invalid.json")))
        session.start()
        assertEquals(AccountActionOutcome.Failed, session.deleteAccount())
        assertTrue(session.phase.value is SessionPhase.SignedIn)
    }

    @Test
    fun `403 browser_session_required opens the site instead of failing`() = runTest {
        val body = """{"ok":false,"data":null,"error":{"code":"browser_session_required","message":"this can only be done on foreversixty.gg in a browser"},"request_id":"r"}"""
        val (session, _) = manager(token = "fsd_x", engine = answering(HttpStatusCode.Forbidden, body))
        assertEquals(AccountActionOutcome.NeedsBrowser, session.deleteAccount())
        assertEquals(AccountActionOutcome.NeedsBrowser, session.revoke("dvc_3Hm7pLa9"))
    }

    @Test
    fun `the picked character wins, then the main, then the first`() {
        val root = ApiJson.parseToJsonElement(fixture("me.json")).jsonObject
        val me = ApiJson.decodeFromJsonElement(Me.serializer(), requireNotNull(root["data"]))
        val other = me.characters.first { it.key != me.mainCharacterKey }
        assertEquals(other.key, CharacterSelection.current(me, other.key)?.key)
        assertEquals(me.mainCharacterKey, CharacterSelection.current(me, "us/pvp/gone")?.key)
        assertEquals(me.mainCharacterKey, CharacterSelection.current(me, null)?.key)
    }

    @Test
    fun `ages read in one unit`() {
        val now = Instant.ofEpochSecond(1_000_000)
        val cases = mapOf(30L to RelativeAge.Now, 300L to RelativeAge.Minutes(5), 7_200L to RelativeAge.Hours(2), 259_200L to RelativeAge.Days(3), -10L to RelativeAge.Now)
        cases.forEach { (seconds, age) -> assertEquals(age, RelativeAge.of(now.minusSeconds(seconds), now)) }
    }

    private fun meOf(name: String): Me {
        val root = ApiJson.parseToJsonElement(fixture(name)).jsonObject
        return ApiJson.decodeFromJsonElement(Me.serializer(), requireNotNull(root["data"]))
    }

    @Test
    fun `a free account with no guild plan or seat loses only its data`() {
        assertTrue(DeletionConsequence.of(meOf("me.json")).isEmpty())
    }

    @Test
    fun `a founder Premium plan names the plan and the founder price`() {
        assertEquals(listOf(DeletionConsequence.PremiumPlan, DeletionConsequence.FounderPrice), DeletionConsequence.of(meOf("me.premium.json")))
    }

    @Test
    fun `a guild plan the account pays for ends with its seats`() {
        assertEquals(listOf(DeletionConsequence.GuildPlan("OLYMPUS XXVII")), DeletionConsequence.of(meOf("me.guild-payer.json")))
    }

    @Test
    fun `a seat held from a guild ends`() {
        assertEquals(listOf(DeletionConsequence.Seat("OLYMPUS XXVII")), DeletionConsequence.of(meOf("me.seat.json")))
    }
}
