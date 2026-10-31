package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpStatusCode
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

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
