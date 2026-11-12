package gg.foreversixty.app.network

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.AppConfig
import kotlinx.coroutines.CompletableDeferred
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.withTimeoutOrNull
import kotlinx.serialization.SerializationException
import java.io.File
import java.io.IOException
import kotlin.time.Duration

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
    private val state = MutableStateFlow(cache.load()?.let(::decode))

    /** Completes when this process's first [refresh] has answered, with a config or a failure. */
    private val firstAnswer = CompletableDeferred<Unit>()

    /**
     * The config as it stands: the cached answer at launch, then each [refresh]. The shell reads
     * its feature switches from here, so a switch the API turns on shows without a restart.
     */
    val config: StateFlow<AppConfig?> = state.asStateFlow()

    val current: AppConfig? get() = state.value

    /**
     * Fetches the config and caches it. On failure the cached or bundled answer stands and the
     * failure is returned for the caller to log, never shown: the app works offline.
     */
    suspend fun refresh(): ApiResult<AppConfig> {
        val result = client.send(ApiEndpoints.getAppConfig())
        if (result is ApiResult.Success) {
            state.value = result.data
            cache.save(ApiJson.encodeToString(AppConfig.serializer(), result.data))
        }
        firstAnswer.complete(Unit)
        return result
    }

    /**
     * A fresh install has no cached config, so until the first [refresh] answers every switch
     * reads its bundled default (all off) and the shell would show every tab as not yet here.
     * Holds for that first answer, a config or a failure that leaves the defaults, for at most
     * [hold]; returns at once when a config is already known. The refresh itself runs on its
     * own and keeps going after the hold, so a slow answer still lands in [config].
     */
    suspend fun awaitFirstAnswer(hold: Duration): Boolean {
        if (current != null) return true
        withTimeoutOrNull(hold) { firstAnswer.await() }
        return current != null
    }

    fun isEnabled(feature: Feature): Boolean = current?.features?.get(feature.key) ?: feature.bundledDefault

    /**
     * A feature is on in debug and fixture builds whatever the config says, so every lane's screens
     * can be built and captured before the API turns them on; a release build on the live API
     * follows [isEnabled] alone (the Wave 1 brief).
     */
    fun isAvailable(feature: Feature, mode: ApiMode, isDebugBuild: Boolean): Boolean =
        isDebugBuild || mode == ApiMode.Fixture || isEnabled(feature)

    /** True when the API no longer supports [version]: block with the update screen. */
    fun requiresUpdate(version: String): Boolean {
        val minimum = current?.minAppVersion?.android ?: return false
        return AppVersion.isBelow(version, minimum)
    }

    /** Whether the storefront country may show See plans (`web_checkout_regions`). */
    fun mayLinkToWebCheckout(storefront: String): Boolean =
        current?.webCheckoutRegions?.any { it.equals(storefront, ignoreCase = true) } ?: false

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
