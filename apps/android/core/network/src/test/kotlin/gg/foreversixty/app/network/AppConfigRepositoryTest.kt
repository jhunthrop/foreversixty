package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpStatusCode
import io.ktor.http.HttpHeaders
import io.ktor.http.headersOf
import kotlinx.coroutines.CompletableDeferred
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import kotlin.time.Duration.Companion.seconds
import kotlin.time.measureTimedValue

class AppConfigRepositoryTest {
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)

    private fun fixtures() = ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source)))

    private fun offline() = ApiClient(
        ForeverHttpClient.create(config, NoToken, MockEngine { respond("", HttpStatusCode.ServiceUnavailable) }),
    )

    @Test
    fun `debug and fixture builds turn every feature on, a live release follows the config`() {
        val repo = AppConfigRepository(offline(), InMemoryAppConfigCache())

        assertTrue(repo.isAvailable(Feature.SignIn, ApiMode.Fixture, isDebugBuild = false))
        assertTrue(repo.isAvailable(Feature.SignIn, ApiMode.Live, isDebugBuild = true))
        assertFalse(repo.isAvailable(Feature.SignIn, ApiMode.Live, isDebugBuild = false))
    }

    @Test
    fun `before any answer every feature takes its bundled default`() {
        val repo = AppConfigRepository(offline(), InMemoryAppConfigCache())

        Feature.entries.forEach { assertEquals(it.key, it.bundledDefault, repo.isEnabled(it)) }
        assertFalse(repo.requiresUpdate("0.0.1"))
        assertFalse(repo.mayLinkToWebCheckout("US"))
    }

    @Test
    fun `the API's answer switches features, and a key it leaves out keeps the default`() = runTest {
        val repo = AppConfigRepository(fixtures(), InMemoryAppConfigCache())

        assertTrue(repo.refresh() is ApiResult.Success)
        assertTrue(repo.isEnabled(Feature.Raids))
        assertFalse(repo.isEnabled(Feature.SignIn))
        assertEquals(Feature.Sims.bundledDefault, repo.isEnabled(Feature.Sims))
        assertTrue(repo.mayLinkToWebCheckout("us"))
        assertFalse(repo.mayLinkToWebCheckout("DE"))
    }

    @Test
    fun `the answer is cached, so the next launch reads it offline`() = runTest {
        val cache = InMemoryAppConfigCache()
        AppConfigRepository(fixtures(), cache).refresh()

        val nextLaunch = AppConfigRepository(offline(), cache)

        assertTrue(nextLaunch.refresh() is ApiResult.Failure)
        assertTrue(nextLaunch.isEnabled(Feature.Raids))
        assertTrue(nextLaunch.current?.activeBuild?.isNotEmpty() == true)
    }

    @Test
    fun `the config is a flow, so the shell follows each answer`() = runTest {
        val repo = AppConfigRepository(fixtures(), InMemoryAppConfigCache())
        assertNull(repo.config.value)

        repo.refresh()

        assertEquals(repo.current, repo.config.value)
        assertTrue(repo.config.value?.features?.get(Feature.Raids.key) == true)
    }

    @Test
    fun `with a cached config the shell never holds`() = runTest {
        val cache = InMemoryAppConfigCache()
        AppConfigRepository(fixtures(), cache).refresh()

        assertTrue(AppConfigRepository(offline(), cache).awaitFirstAnswer(HOLD))
    }

    // The hold is real time (the engine answers on its own thread, so runTest's virtual clock
    // would run past it); each case is a few milliseconds against a one-second hold.

    @Test
    fun `a fresh install holds for its first answer and builds from it`() = runBlocking {
        val repo = AppConfigRepository(fixtures(), InMemoryAppConfigCache())
        val fetch = launch { repo.refresh() }

        val (answered, took) = measureTimedValue { repo.awaitFirstAnswer(HOLD) }

        assertTrue(answered)
        assertTrue("$took", took < HOLD)
        assertTrue(repo.isEnabled(Feature.Raids))
        fetch.join()
    }

    @Test
    fun `a failed first answer ends the hold at once, on the bundled defaults`() = runBlocking {
        val repo = AppConfigRepository(offline(), InMemoryAppConfigCache())
        val fetch = launch { repo.refresh() }

        val (answered, took) = measureTimedValue { repo.awaitFirstAnswer(HOLD) }

        assertFalse(answered)
        assertTrue("$took", took < HOLD)
        assertEquals(Feature.Raids.bundledDefault, repo.isEnabled(Feature.Raids))
        fetch.join()
    }

    @Test
    fun `a slow first answer ends the hold when it passes, and still lands`() = runBlocking {
        val gate = CompletableDeferred<Unit>()
        val slow = ApiClient(
            ForeverHttpClient.create(config, NoToken, MockEngine {
                gate.await()
                respond(requireNotNull(source.read("app.config.json")), HttpStatusCode.OK, headersOf(HttpHeaders.ContentType, "application/json"))
            }),
        )
        val repo = AppConfigRepository(slow, InMemoryAppConfigCache())
        val fetch = launch { repo.refresh() }

        val (answered, took) = measureTimedValue { repo.awaitFirstAnswer(HOLD) }

        assertFalse(answered)
        assertTrue("$took", took >= HOLD)
        assertNull(repo.current)

        gate.complete(Unit)
        fetch.join()
        assertNotNull(repo.config.value)
        assertTrue(repo.isEnabled(Feature.Raids))
    }

    @Test
    fun `an unreadable cache is ignored`() {
        val repo = AppConfigRepository(offline(), InMemoryAppConfigCache("nope"))

        assertNull(repo.current)
        assertEquals(Feature.Raids.bundledDefault, repo.isEnabled(Feature.Raids))
    }

    @Test
    fun `versions compare numerically, part by part`() {
        assertTrue(AppVersion.isBelow("1.9.2", "1.10.0"))
        assertFalse(AppVersion.isBelow("1.10.0", "1.9.2"))
        assertFalse(AppVersion.isBelow("1.0", "1.0.0"))
        assertTrue(AppVersion.isBelow("1.0.0", "1.0.1"))
        assertFalse(AppVersion.isBelow("2", "1.99"))
    }
}

/** The tests' first-launch hold; MainActivity's is its own. */
private val HOLD = 1.seconds
