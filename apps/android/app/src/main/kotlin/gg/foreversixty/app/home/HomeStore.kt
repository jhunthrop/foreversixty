package gg.foreversixty.app.home

import gg.foreversixty.app.common.CacheLimits
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.common.FailedRefresh
import gg.foreversixty.app.common.StripDecision
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import java.io.File
import java.io.IOException
import java.time.Instant
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.withContext
import kotlinx.serialization.SerializationException
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import kotlinx.serialization.json.longOrNull

/** The last good home answer per character. Sign-out deletes it with [removeAll]. */
interface HomeCache {
    suspend fun load(key: String): HomeCacheEntry?
    suspend fun save(key: String, entry: HomeCacheEntry)
    suspend fun removeAll()
}

data class HomeCacheEntry(
    val home: CharacterHome,
    /** Epoch milliseconds of the fetch. */
    val fetchedAtMillis: Long,
) {
    val fetchedAt: Instant get() = Instant.ofEpochMilli(fetchedAtMillis)

    fun encode(): String = ApiJson.encodeToString(
        JsonObject.serializer(),
        JsonObject(mapOf(HOME to ApiJson.encodeToJsonElement(CharacterHome.serializer(), home), FETCHED_AT to JsonPrimitive(fetchedAtMillis))),
    )

    companion object {
        private const val HOME = "home"
        private const val FETCHED_AT = "fetched_at"

        /** Throws SerializationException or IllegalArgumentException for a file that is not an entry. */
        fun decode(text: String): HomeCacheEntry {
            val root = ApiJson.parseToJsonElement(text).jsonObject
            val home = ApiJson.decodeFromJsonElement(CharacterHome.serializer(), requireNotNull(root[HOME]) { "no home" })
            val at = requireNotNull(root[FETCHED_AT]?.jsonPrimitive?.longOrNull) { "no fetched_at" }
            return HomeCacheEntry(home, at)
        }
    }
}

/**
 * One JSON file per character, in a directory the caller places under `noBackupFilesDir`, so it
 * never reaches a backup (spec §5.0 Caching).
 */
class FileHomeCache(private val directory: File) : HomeCache {
    override suspend fun load(key: String): HomeCacheEntry? = withContext(Dispatchers.IO) {
        val file = file(key).takeIf { it.exists() } ?: return@withContext null
        try {
            HomeCacheEntry.decode(file.readText())
        } catch (_: SerializationException) {
            null
        } catch (_: IllegalArgumentException) {
            null
        } catch (_: IOException) {
            null
        }
    }

    override suspend fun save(key: String, entry: HomeCacheEntry) = withContext(Dispatchers.IO) {
        try {
            directory.mkdirs()
            file(key).writeText(entry.encode())
        } catch (_: IOException) {
            // The cache is a convenience: the next launch fetches again.
        }
    }

    override suspend fun removeAll() {
        withContext(Dispatchers.IO) { directory.deleteRecursively() }
    }

    private fun file(key: String) = File(directory, key.replace('/', '_') + ".json")
}

/** An in-memory cache for tests and previews. */
class MemoryHomeCache(initial: Map<String, HomeCacheEntry> = emptyMap()) : HomeCache {
    private val entries = initial.toMutableMap()
    override suspend fun load(key: String) = entries[key]
    override suspend fun save(key: String, entry: HomeCacheEntry) {
        entries[key] = entry
    }
    override suspend fun removeAll() = entries.clear()
}

/** What Home shows (spec F2.1 to F2.5). */
sealed interface HomePhase {
    /** First load, nothing cached: skeletons. */
    data object Loading : HomePhase

    /**
     * An answer; [offlineSince] is the saved answer's time while it shows under a strip: offline, or
     * [stale] when the server did not answer and the answer is past its limit (W3 §5.7).
     */
    data class Loaded(val home: CharacterHome, val offlineSince: Instant? = null, val stale: Boolean = false) : HomePhase

    /** The server failed. A cached answer still feeds the Next rows. */
    data class Failed(val cached: CharacterHome?) : HomePhase
}

/**
 * Loads the home answer for one character: the cache draws at once, then the API refreshes it.
 * A refresh never replaces content with skeletons; no connection keeps the cached answer under
 * the offline strip; a server error is the error state.
 */
class HomeStore(
    private val fetch: suspend (HomeCharacter) -> ApiResult<CharacterHome>,
    private val cache: HomeCache,
    private val now: () -> Instant = Instant::now,
    private val connectivity: Connectivity = Connectivity.Unknown,
) {
    private val state = MutableStateFlow<HomePhase>(HomePhase.Loading)
    val phase: StateFlow<HomePhase> = state.asStateFlow()

    private var fetchedAt: Instant? = null

    /** Draws the cached answer, then fetches. */
    suspend fun load(character: HomeCharacter) {
        val entry = cache.load(character.key)
        if (entry != null) {
            state.value = HomePhase.Loaded(entry.home)
            fetchedAt = entry.fetchedAt
        } else {
            state.value = HomePhase.Loading
        }
        refresh(character)
    }

    /** Refreshes on foreground only when the answer is older than five minutes. */
    suspend fun refreshIfStale(character: HomeCharacter) {
        val at = fetchedAt
        if (at == null || now().toEpochMilli() - at.toEpochMilli() >= STALE_AFTER_MILLIS) refresh(character)
    }

    suspend fun refresh(character: HomeCharacter) {
        when (val answer = fetch(character)) {
            is ApiResult.Success -> {
                val at = now()
                cache.save(character.key, HomeCacheEntry(answer.data, at.toEpochMilli()))
                fetchedAt = at
                state.value = HomePhase.Loaded(answer.data)
            }
            is ApiResult.Failure -> fail(answer.error, cache.load(character.key))
        }
    }

    private fun fail(error: ApiError, cached: HomeCacheEntry?) {
        val decision = StripDecision.afterFailure(error, cached?.fetchedAt, now(), CacheLimits.HOME, connectivity.isConnected())
        state.value = when (decision) {
            is FailedRefresh.Offline -> HomePhase.Loaded(requireNotNull(cached).home, offlineSince = decision.since)
            is FailedRefresh.Stale -> HomePhase.Loaded(requireNotNull(cached).home, offlineSince = decision.since, stale = true)
            FailedRefresh.KeepSaved -> HomePhase.Loaded(requireNotNull(cached).home)
            FailedRefresh.Error -> HomePhase.Failed(cached?.home)
        }
    }

    companion object {
        const val STALE_AFTER_MILLIS = 5L * 60L * 1000L

        /** The live store: `GET /v1/characters/{region}/{ruleset}/{name}/home`. */
        fun live(client: ApiClient, cache: HomeCache, connectivity: Connectivity = Connectivity.Unknown) = HomeStore(
            fetch = { client.send(ApiEndpoints.getCharacterHome(it.region, it.ruleset, it.slug)) },
            cache = cache,
            connectivity = connectivity,
        )
    }
}
