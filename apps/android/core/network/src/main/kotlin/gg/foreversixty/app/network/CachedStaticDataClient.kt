package gg.foreversixty.app.network

import java.io.File
import java.io.IOException
import java.util.concurrent.ConcurrentHashMap

/** One saved static file: its body and when it was fetched (epoch milliseconds). */
data class StaticFileEntry(val body: String, val savedAtMillis: Long)

/**
 * Where static files are kept for offline use, keyed by their `/data/...` path. A build's files
 * never change under one path, so a newer build simply uses new keys.
 */
interface StaticFileCache {
    fun read(path: String): StaticFileEntry?
    fun write(body: String, path: String, savedAtMillis: Long)
}

/**
 * The cache on disk: one file per path under [directory], its modification time the save time.
 * The app passes a folder in `noBackupFilesDir` (spec §5.0 Caching).
 */
class FileStaticFileCache(private val directory: File) : StaticFileCache {
    override fun read(path: String): StaticFileEntry? {
        val file = fileFor(path)
        return try {
            if (file.exists()) StaticFileEntry(file.readText(), file.lastModified()) else null
        } catch (e: IOException) {
            null
        }
    }

    /** Best effort: a full disk loses only the offline copy, never the answer on screen. */
    override fun write(body: String, path: String, savedAtMillis: Long) {
        val file = fileFor(path)
        try {
            file.parentFile?.mkdirs()
            val temp = File(file.parentFile, "${file.name}.tmp")
            temp.writeText(body)
            if (temp.renameTo(file)) file.setLastModified(savedAtMillis)
        } catch (e: IOException) {
            // The next good fetch saves it again.
        }
    }

    private fun fileFor(path: String): File =
        File(directory, path.split('/').filter { it.isNotEmpty() && it != ".." }.joinToString(File.separator))
}

/** For tests and previews. */
class InMemoryStaticFileCache : StaticFileCache {
    private val entries = ConcurrentHashMap<String, StaticFileEntry>()
    override fun read(path: String): StaticFileEntry? = entries[path]
    override fun write(body: String, path: String, savedAtMillis: Long) {
        entries[path] = StaticFileEntry(body, savedAtMillis)
    }
}

/** A static file as the app has it: fetched now, or the saved copy when offline. */
data class CachedFile<T>(val value: T, val savedAtMillis: Long, val isFromCache: Boolean)

/**
 * Static files with an offline copy (spec §5.0): every good fetch is saved; a fetch that fails
 * for lack of a connection answers the saved copy instead. Any other failure (a 404, a body that
 * does not decode) stays an error, because that is the error state, not offline.
 */
class CachedStaticDataClient(
    private val client: StaticDataClient,
    private val cache: StaticFileCache,
    private val clock: () -> Long = System::currentTimeMillis,
) {
    suspend fun <T> load(resource: StaticResource<T>): ApiResult<CachedFile<T>> =
        when (val answer = client.fetch(resource)) {
            is ApiResult.Success -> {
                val now = clock()
                cache.write(answer.data.body, resource.path, now)
                ApiResult.Success(CachedFile(answer.data.value, now, isFromCache = false), requestId = "")
            }
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
}
