package gg.foreversixty.app.network

import java.io.File
import java.io.IOException
import java.util.concurrent.ConcurrentHashMap

/** One saved static file: its body and when it was fetched (epoch milliseconds). */
data class StaticFileEntry(val body: String, val savedAtMillis: Long)

/** A saved file as the Offline data screen sizes it: its `/data/...` path, bytes and save time. */
data class SavedStaticFile(val path: String, val bytes: Long, val savedAtMillis: Long)

/**
 * Where static files are kept for offline use, keyed by their `/data/...` path. A build's files
 * never change under one path, so a newer build simply uses new keys and [keepBuild] drops the old.
 */
interface StaticFileCache {
    fun read(path: String): StaticFileEntry?
    fun write(body: String, path: String, savedAtMillis: Long)

    /** Every saved file, for the size report. */
    fun files(): List<SavedStaticFile>

    /** Drops every `/data/{build}/...` file of another build; build-less files stay. */
    fun keepBuild(build: String)

    fun deleteAll()
}

/**
 * The cache on disk: one file per path under [directory], its modification time the save time.
 * The app passes one folder in `noBackupFilesDir` (spec §5.0 Caching).
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
            if (temp.renameTo(file)) file.setLastModified(savedAtMillis) else temp.delete()
        } catch (e: IOException) {
            // The next good fetch saves it again.
        }
    }

    override fun files(): List<SavedStaticFile> = directory.walkTopDown()
        .filter { it.isFile && !it.name.endsWith(".tmp") }
        .map { SavedStaticFile("/" + it.relativeTo(directory).invariantSeparatorsPath, it.length(), it.lastModified()) }
        .toList()

    override fun keepBuild(build: String) {
        File(directory, DATA_DIR).listFiles()?.filter { it.isDirectory && it.name != build }?.forEach { it.deleteRecursively() }
    }

    override fun deleteAll() {
        directory.listFiles()?.forEach { it.deleteRecursively() }
    }

    private fun fileFor(path: String): File =
        File(directory, path.split('/').filter { it.isNotEmpty() && it != ".." }.joinToString(File.separator))

    private companion object {
        const val DATA_DIR = "data"
    }
}

/** For tests and previews. */
class InMemoryStaticFileCache : StaticFileCache {
    private val entries = ConcurrentHashMap<String, StaticFileEntry>()
    override fun read(path: String): StaticFileEntry? = entries[path]
    override fun write(body: String, path: String, savedAtMillis: Long) {
        entries[path] = StaticFileEntry(body, savedAtMillis)
    }

    override fun files(): List<SavedStaticFile> =
        entries.map { (path, entry) -> SavedStaticFile(path, entry.body.toByteArray().size.toLong(), entry.savedAtMillis) }

    override fun keepBuild(build: String) {
        entries.keys.filter { StaticDataStore.buildOf(it)?.let { other -> other != build } == true }.forEach(entries::remove)
    }

    override fun deleteAll() = entries.clear()
}

/** A static file as the app has it: fetched now, or the saved copy when offline. */
data class CachedFile<T>(val value: T, val savedAtMillis: Long, val isFromCache: Boolean)
