package gg.foreversixty.app.network

import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.HttpStatusCode
import java.io.File
import kotlinx.coroutines.test.runTest
import kotlinx.io.IOException
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.rules.TemporaryFolder

/** The per-build static data store (Wave 3 spec §14, lane F12), the same list as iOS's StaticDataStoreTests. */
class StaticDataStoreTest {
    @get:Rule
    val temp = TemporaryFolder()

    private val build = "1.60.1.70658"
    private val older = "1.60.1.70532"
    private val search = StaticEndpoints.getSearchIndex(build)
    private val body = requireNotNull(javaClass.classLoader?.getResource("$FIXTURE_ASSET_DIR/cdn.search.json")).readText()

    private var offline = false
    private var requests = 0

    private val engine = MockEngine { _ ->
        requests++
        if (offline) throw IOException("offline")
        respond(body, HttpStatusCode.OK)
    }

    private fun store(directory: File? = temp.root) = StaticDataStore(StaticDataClient(engine), directory)

    @Test
    fun `a file is fetched once, then answered from memory`() = runTest {
        val store = store()
        val first = store.get(search) as ApiResult.Success
        store.get(search)
        assertEquals(build, first.data.build)
        assertEquals(1, requests)
    }

    @Test
    fun `a file saved on disk is answered offline by a new session`() = runTest {
        store().get(search)
        offline = true
        val answer = store().get(search) as ApiResult.Success
        assertEquals(build, answer.data.build)
        assertEquals(1, requests)
    }

    @Test
    fun `offline with nothing saved is the network error`() = runTest {
        offline = true
        val answer = store().get(search) as ApiResult.Failure
        assertTrue(answer.error is ApiError.Network)
    }

    @Test
    fun `swap keeps only the active build`() = runTest {
        val store = store()
        store.get(StaticEndpoints.getSearchIndex(older))
        store.get(search)
        store.swap(build)
        assertFalse(File(temp.root, older).exists())
        assertTrue(File(temp.root, build).exists())
        offline = true
        assertTrue(store.get(search) is ApiResult.Success)
        assertTrue(store.get(StaticEndpoints.getSearchIndex(older)) is ApiResult.Failure)
    }

    @Test
    fun `size is the bytes on disk, and delete empties it`() = runTest {
        val store = store()
        assertEquals(0L, store.sizeBytes())
        store.get(search)
        assertEquals(body.toByteArray().size.toLong(), store.sizeBytes())
        store.deleteAll()
        assertEquals(0L, store.sizeBytes())
        offline = true
        assertTrue(store.get(search) is ApiResult.Failure)
    }

    @Test
    fun `without a directory it keeps files in memory only`() = runTest {
        val store = store(directory = null)
        store.get(search)
        assertEquals(0L, store.sizeBytes())
        offline = true
        assertTrue(store.get(search) is ApiResult.Success)
    }
}
