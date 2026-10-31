package gg.foreversixty.app.network

import java.io.File
import java.io.IOException
import java.util.concurrent.ConcurrentHashMap
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

/**
 * The per-build static data store (Wave 3 spec §14: lane F12 owns it, F9 and F15 reuse it): every
 * `/data/{build}/…` file the app reads, in memory for the session and on disk under [directory]
 * for offline use. A path under one build never changes, so a saved copy is answered without a
 * request. [directory] null keeps files in memory only (tests and previews). Never sends the
 * device token: it reads through [StaticDataClient].
 */
class StaticDataStore(private val client: StaticDataClient, private val directory: File? = null) {
    private val memory = ConcurrentHashMap<String, Any>()

    /** The file: from memory, else from disk, else fetched (and kept). A failed fetch is its error. */
    suspend fun <T> get(resource: StaticResource<T>): ApiResult<T> {
        @Suppress("UNCHECKED_CAST")
        (memory[resource.path] as? T)?.let { return ApiResult.Success(it, requestId = "") }
        readDisk(resource)?.let {
            memory[resource.path] = it as Any
            return ApiResult.Success(it, requestId = "")
        }
        return when (val answer = client.fetch(resource)) {
            is ApiResult.Success -> {
                memory[resource.path] = answer.data.value as Any
                writeDisk(resource.path, answer.data.body)
                ApiResult.Success(answer.data.value, requestId = "")
            }
            is ApiResult.Failure -> answer
        }
    }

    /** Swap: keeps only [activeBuild]'s files, on disk and in memory, once a newer build is in use. */
    suspend fun swap(activeBuild: String) = withContext(Dispatchers.IO) {
        val kept = "${StaticDataClient.DATA_PREFIX}$activeBuild/"
        memory.keys.filterNot { it.startsWith(kept) }.forEach { memory.remove(it) }
        directory?.listFiles()?.filter { it.name != activeBuild }?.forEach { it.deleteRecursively() }
    }

    /** Deletes every saved file (sign-out, the Offline data screen's delete). */
    suspend fun deleteAll() = withContext(Dispatchers.IO) {
        memory.clear()
        directory?.listFiles()?.forEach { it.deleteRecursively() }
    }

    /** Size accounting: the bytes the saved files take on disk. */
    suspend fun sizeBytes(): Long = withContext(Dispatchers.IO) {
        directory?.walkTopDown()?.filter { it.isFile }?.sumOf { it.length() } ?: 0L
    }

    private fun diskFile(path: String): File? = directory?.let { root ->
        File(root, path.removePrefix(StaticDataClient.DATA_PREFIX).split('/').filter { it.isNotEmpty() && it != ".." }.joinToString(File.separator))
    }

    private suspend fun <T> readDisk(resource: StaticResource<T>): T? = withContext(Dispatchers.IO) {
        val file = diskFile(resource.path)?.takeIf { it.isFile } ?: return@withContext null
        val body = try {
            file.readText()
        } catch (e: IOException) {
            return@withContext null
        }
        (StaticDataClient.decode(resource, body) as? ApiResult.Success)?.data
    }

    /** Best effort: a full disk loses only the offline copy, never the answer on screen. */
    private suspend fun writeDisk(path: String, body: String) = withContext(Dispatchers.IO) {
        val file = diskFile(path) ?: return@withContext
        try {
            file.parentFile?.mkdirs()
            val temp = File(file.parentFile, "${file.name}.tmp")
            temp.writeText(body)
            if (!temp.renameTo(file)) temp.delete()
        } catch (e: IOException) {
            // The next read fetches again.
        }
    }
}
