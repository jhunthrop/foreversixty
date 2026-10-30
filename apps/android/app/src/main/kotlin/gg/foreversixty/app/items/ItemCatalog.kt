package gg.foreversixty.app.items

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.StaticResource
import java.io.File
import java.io.IOException
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import kotlinx.coroutines.withContext
import kotlinx.serialization.SerializationException

/**
 * The build's static files the tooltip reads (`items/{class}.json`, `sets.json`,
 * `bis/{spec}.json`), fetched once from the CDN and kept per build: in memory for the session, on
 * disk until a newer build is in use (spec §5.0 Caching). Never sends the device token.
 * [directory] null keeps them in memory only (tests and previews).
 */
class ItemCatalog(private val client: StaticDataClient, private val directory: File? = null) {
    private val memory = mutableMapOf<String, Any>()
    private val lock = Mutex()

    /**
     * The files a tooltip needs. Items are required; sets and the BiS file only add lines, so a
     * missing one leaves its lines out rather than failing the sheet. Null when the items file
     * did not load.
     */
    suspend fun sources(request: ItemTooltipRequest): ItemTooltipSources? {
        val items = file(StaticEndpoints.getItems(request.build, request.classSlug)) ?: return null
        val sets = file(StaticEndpoints.getSets(request.build)).orEmpty()
        val bis = request.spec?.let { file(StaticEndpoints.getBis(request.build, it)) }
        return ItemTooltipSources(items.items.associateBy { it.id }, sets, bis)
    }

    /** Forgets every build but [build] on disk, once a newer build is in use. */
    suspend fun prune(build: String) = withContext(Dispatchers.IO) {
        directory?.listFiles()?.filter { it.name != build }?.forEach { it.deleteRecursively() }
    }

    private suspend fun <T> file(resource: StaticResource<T>): T? = lock.withLock {
        @Suppress("UNCHECKED_CAST")
        (memory[resource.path] as? T)?.let { return@withLock it }
        val cached = readDisk(resource)
        if (cached != null) {
            memory[resource.path] = cached as Any
            return@withLock cached
        }
        when (val answer = client.get(resource)) {
            is ApiResult.Success -> {
                memory[resource.path] = answer.data as Any
                writeDisk(resource, answer.data)
                answer.data
            }
            is ApiResult.Failure -> null
        }
    }

    private fun diskFile(path: String): File? =
        directory?.let { File(it, path.removePrefix(StaticDataClient.DATA_PREFIX)) }

    private suspend fun <T> readDisk(resource: StaticResource<T>): T? = withContext(Dispatchers.IO) {
        val file = diskFile(resource.path)?.takeIf { it.exists() } ?: return@withContext null
        try {
            ApiJson.decodeFromString(resource.response, file.readText())
        } catch (_: SerializationException) {
            null
        } catch (_: IOException) {
            null
        } catch (_: IllegalArgumentException) {
            null
        }
    }

    private suspend fun <T> writeDisk(resource: StaticResource<T>, value: T) = withContext(Dispatchers.IO) {
        val file = diskFile(resource.path) ?: return@withContext
        try {
            file.parentFile?.mkdirs()
            file.writeText(ApiJson.encodeToString(resource.response, value))
        } catch (_: IOException) {
            // The cache is a convenience: the next read fetches again.
        }
    }
}
