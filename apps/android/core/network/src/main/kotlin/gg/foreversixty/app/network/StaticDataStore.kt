package gg.foreversixty.app.network

import java.util.concurrent.ConcurrentHashMap
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

/**
 * The app's one static data store (Wave 3 §14, lane F12, unified by lane F15): every
 * `/data/{build}/…` file the app reads, saved in one [cache] (one folder on the phone) for
 * offline use, with one build swap, one delete and one size report. Two reads:
 * - [get]: saved first (memory, then disk), else fetched. For files that never change under a path
 *   (the tooltip's files, the search index, the item pages).
 * - [load]: the network first, the saved copy only when the fetch could not connect. For screens
 *   that show `Offline` when they draw a saved copy (Guides, BiS, Talents).
 * Never sends the device token: it reads through [StaticDataClient].
 */
class StaticDataStore(
    private val client: StaticDataClient,
    private val cache: StaticFileCache = InMemoryStaticFileCache(),
    private val clock: () -> Long = System::currentTimeMillis,
) {
    private val memory = ConcurrentHashMap<String, Any>()

    /** The file: from memory, else from the saved copy, else fetched (and kept). A failed fetch is its error. */
    suspend fun <T> get(resource: StaticResource<T>): ApiResult<T> {
        @Suppress("UNCHECKED_CAST")
        (memory[resource.path] as? T)?.let { return ApiResult.Success(it, requestId = "") }
        withContext(Dispatchers.IO) { cached(resource) }?.let {
            memory[resource.path] = it.value as Any
            return ApiResult.Success(it.value, requestId = "")
        }
        return when (val answer = fetchAndSave(resource)) {
            is ApiResult.Success -> {
                memory[resource.path] = answer.data.value as Any
                ApiResult.Success(answer.data.value, requestId = "")
            }
            is ApiResult.Failure -> answer
        }
    }

    /**
     * The file now, saved; a fetch that fails for lack of a connection answers the saved copy
     * instead. Any other failure (a 404, a body that does not decode) stays an error, because that
     * is the error state, not offline.
     */
    suspend fun <T> load(resource: StaticResource<T>): ApiResult<CachedFile<T>> = when (val answer = fetchAndSave(resource)) {
        is ApiResult.Success -> answer
        is ApiResult.Failure -> {
            val saved = if (answer.error is ApiError.Network) cached(resource) else null
            if (saved != null) ApiResult.Success(saved, requestId = "") else answer
        }
    }

    /** The saved copy alone, with no request: what a screen draws at once before it refreshes. */
    fun <T> cached(resource: StaticResource<T>): CachedFile<T>? {
        val entry = cache.read(resource.path) ?: return null
        val decoded = StaticDataClient.decode(resource, entry.body) as? ApiResult.Success ?: return null
        return CachedFile(decoded.data, entry.savedAtMillis, isFromCache = true)
    }

    /** The saved copy's size in bytes, or null when the file is not on the phone. */
    fun savedBytes(resource: StaticResource<*>): Long? = cache.read(resource.path)?.body?.toByteArray()?.size?.toLong()

    /** Swap: keeps only [activeBuild]'s files, on disk and in memory, once a newer build is in use. */
    suspend fun swap(activeBuild: String) = withContext(Dispatchers.IO) {
        memory.keys.filter { buildOf(it)?.let { build -> build != activeBuild } == true }.forEach { memory.remove(it) }
        cache.keepBuild(activeBuild)
    }

    /** Deletes every saved file (sign-out, the Offline data screen's Remove). */
    suspend fun deleteAll() = withContext(Dispatchers.IO) {
        memory.clear()
        cache.deleteAll()
    }

    /** Every saved file with its size and save time, for the Offline data screen. */
    suspend fun files(): List<SavedStaticFile> = withContext(Dispatchers.IO) { cache.files() }

    /** Size accounting: the bytes the saved files take. */
    suspend fun sizeBytes(): Long = files().sumOf { it.bytes }

    private suspend fun <T> fetchAndSave(resource: StaticResource<T>): ApiResult<CachedFile<T>> =
        when (val answer = client.fetch(resource)) {
            is ApiResult.Success -> {
                val now = clock()
                withContext(Dispatchers.IO) { cache.write(answer.data.body, resource.path, now) }
                ApiResult.Success(CachedFile(answer.data.value, now, isFromCache = false), requestId = "")
            }
            is ApiResult.Failure -> answer
        }

    companion object {
        /** The build a `/data/{build}/…` path belongs to; null for build-less files (`/data/active-build.json`). */
        fun buildOf(path: String): String? {
            val parts = path.removePrefix(StaticDataClient.DATA_PREFIX).split('/')
            return parts.first().takeIf { parts.size > 1 && it.isNotEmpty() }
        }
    }
}
