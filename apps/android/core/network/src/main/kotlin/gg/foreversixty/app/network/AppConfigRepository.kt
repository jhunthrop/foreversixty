package gg.foreversixty.app.network

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.AppConfig
import kotlinx.serialization.SerializationException
import java.io.File
import java.io.IOException

/** Where the last good `GET /v1/app/config` answer is kept between launches (not a secret). */
interface AppConfigCache {
    fun load(): String?
    fun save(json: String)
}

/** The cache as one file in the app's private storage. */
class FileAppConfigCache(private val file: File) : AppConfigCache {
    override fun load(): String? = try {
        file.takeIf { it.exists() }?.readText()
    } catch (e: IOException) {
        null
    }

    override fun save(json: String) {
        try {
            file.writeText(json)
        } catch (e: IOException) {
            // A full disk loses only the cache; the next launch fetches again.
        }
    }
}

/** For tests and previews. */
class InMemoryAppConfigCache(private var stored: String? = null) : AppConfigCache {
    override fun load(): String? = stored
    override fun save(json: String) {
        stored = json
    }
}

/**
 * The app config, read once at launch ([refresh]), cached, and answered from the cache or the
 * bundled defaults when offline. A feature is on when the API says so, else when its bundled
 * default (apps/shared/features.json) is; a key the API sends that the build does not know is
 * ignored.
 */
class AppConfigRepository(private val client: ApiClient, private val cache: AppConfigCache) {
    @Volatile
    var current: AppConfig? = cache.load()?.let(::decode)
        private set

    /**
     * Fetches the config and caches it. On failure the cached or bundled answer stands and the
     * failure is returned for the caller to log, never shown: the app works offline.
     */
    suspend fun refresh(): ApiResult<AppConfig> {
        val result = client.send(ApiEndpoints.getAppConfig())
        if (result is ApiResult.Success) {
            current = result.data
            cache.save(ApiJson.encodeToString(AppConfig.serializer(), result.data))
        }
        return result
    }

    fun isEnabled(feature: Feature): Boolean = current?.features?.get(feature.key) ?: feature.bundledDefault

    /** True when the API no longer supports [version]: block with the update screen. */
    fun requiresUpdate(version: String): Boolean {
        val minimum = current?.minAppVersion?.android ?: return false
        return AppVersion.isBelow(version, minimum)
    }

    /** Whether the storefront country may show See plans (`web_checkout_regions`). */
    fun mayLinkToWebCheckout(storefront: String): Boolean =
        current?.webCheckoutRegions?.contains(storefront.uppercase()) ?: false

    private fun decode(json: String): AppConfig? = try {
        ApiJson.decodeFromString(AppConfig.serializer(), json)
    } catch (e: SerializationException) {
        null
    } catch (e: IllegalArgumentException) {
        null
    }
}

/** Dotted version strings compared part by part as numbers: "1.10.0" is above "1.9.2". */
object AppVersion {
    fun isBelow(version: String, minimum: String): Boolean {
        val a = parts(version)
        val b = parts(minimum)
        for (i in 0 until maxOf(a.size, b.size)) {
            val x = a.getOrElse(i) { 0 }
            val y = b.getOrElse(i) { 0 }
            if (x != y) return x < y
        }
        return false
    }

    private fun parts(version: String): List<Int> =
        version.split('.').map { part -> part.takeWhile { it.isDigit() }.toIntOrNull() ?: 0 }
}
