package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpStatusCode
import kotlinx.coroutines.test.runTest
import kotlinx.io.IOException
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.rules.TemporaryFolder

/** The same list as iOS's CachedStaticDataTests. */
class CachedStaticDataClientTest {
    @get:Rule
    val temp = TemporaryFolder()

    private val resource = StaticEndpoints.getBis("1.60.1.70658", "warrior-fury")
    private val body = requireNotNull(javaClass.classLoader?.getResource("$FIXTURE_ASSET_DIR/cdn.bis.warrior-fury.json")).readText()

    private enum class Mode { Ok, Offline, NotFound }

    private var mode = Mode.Ok

    private val engine = MockEngine { _ ->
        when (mode) {
            Mode.Ok -> respond(body, HttpStatusCode.OK)
            Mode.Offline -> throw IOException("offline")
            Mode.NotFound -> respond("", HttpStatusCode.NotFound)
        }
    }

    private fun client(cache: StaticFileCache, now: Long) =
        CachedStaticDataClient(StaticDataClient(engine), cache, clock = { now })

    @Test
    fun `a fetched file is decoded, saved, and marked fresh`() = runTest {
        val cache = InMemoryStaticFileCache()
        val answer = client(cache, 1_000).load(resource) as ApiResult.Success
        assertEquals("warrior-fury", answer.data.value.spec)
        assertFalse(answer.data.isFromCache)
        assertEquals(1_000L, answer.data.savedAtMillis)
        assertEquals(body, cache.read(resource.path)?.body)
    }

    @Test
    fun `offline answers the saved copy with the time it was saved`() = runTest {
        val cache = InMemoryStaticFileCache()
        client(cache, 1_000).load(resource)
        mode = Mode.Offline
        val answer = client(cache, 9_000).load(resource) as ApiResult.Success
        assertTrue(answer.data.isFromCache)
        assertEquals(1_000L, answer.data.savedAtMillis)
        assertTrue(answer.data.value.bands.isNotEmpty())
    }

    @Test
    fun `offline with nothing saved is still the network error`() = runTest {
        mode = Mode.Offline
        val answer = client(InMemoryStaticFileCache(), 0).load(resource) as ApiResult.Failure
        assertTrue(answer.error is ApiError.Network)
    }

    @Test
    fun `a server error is an error, never the cached copy`() = runTest {
        val cache = InMemoryStaticFileCache()
        client(cache, 0).load(resource)
        mode = Mode.NotFound
        val answer = client(cache, 0).load(resource) as ApiResult.Failure
        assertEquals(ApiError.Status(404), answer.error)
    }

    @Test
    fun `the saved copy reads without a request, for the first draw`() = runTest {
        val cache = InMemoryStaticFileCache()
        val reader = client(cache, 5)
        assertNull(reader.cached(resource))
        reader.load(resource)
        val saved = assertNotNullAndGet(reader.cached(resource))
        assertTrue(saved.isFromCache)
        assertEquals(5L, saved.savedAtMillis)
    }

    @Test
    fun `the disk cache keeps a file per path with its save time`() {
        val cache = FileStaticFileCache(temp.newFolder("static"))
        cache.write("{}", "/data/1.60.1.70658/bis/warrior-fury.json", 1_700_000_000_000)
        val entry = assertNotNullAndGet(cache.read("/data/1.60.1.70658/bis/warrior-fury.json"))
        assertEquals("{}", entry.body)
        assertEquals(1_700_000_000_000L, entry.savedAtMillis)
        assertNull(cache.read("/data/1.60.1.70658/bis/warrior-arms.json"))
    }

    private fun <T> assertNotNullAndGet(value: T?): T {
        assertNotNull(value)
        return requireNotNull(value)
    }
}
