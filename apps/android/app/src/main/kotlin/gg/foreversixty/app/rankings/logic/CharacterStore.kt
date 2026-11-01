package gg.foreversixty.app.rankings.logic

import gg.foreversixty.app.common.StripDecision
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.common.CacheLimits
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidHaptics
import java.time.Instant

sealed interface CharacterPhase {
    data object Loading : CharacterPhase

    /** [offlineSince] is set while the cached card shows under a strip: offline, or [stale] (W3 §5.7). */
    data class Loaded(val profile: CharacterProfile, val offlineSince: Instant? = null, val stale: Boolean = false) : CharacterPhase

    /** 404, also an anonymous account's character (F11.7). */
    data object NotFound : CharacterPhase
    data object Failed : CharacterPhase
}

/** The Watch star (spec §7): null until the watch list is read (and for a signed-out viewer). */
data class WatchState(val watching: Boolean, val failed: Boolean = false, val busy: Boolean = false)

/** One character page: its card (cached for offline) and the Watch star, optimistic and reverted on failure. */
class CharacterStore(
    val route: CharacterRoute,
    private val service: RankingsService,
    private val cache: RaidCache,
    private val viewer: RankingsViewer,
    private val connectivity: Connectivity = Connectivity.Unknown,
    private val now: () -> Instant,
) {
    var phase: CharacterPhase by mutableStateOf(CharacterPhase.Loading)
        private set

    var watch: WatchState? by mutableStateOf(null)
        private set

    private val cacheKey: String get() = CACHE_KEY + "${route.region}-${route.ruleset}-${route.slug}"

    suspend fun load() {
        when (val answer = service.character(route.region, route.ruleset, route.name)) {
            is ApiResult.Success -> {
                cache.write(cacheKey, CharacterProfile.serializer(), answer.data, now())
                phase = CharacterPhase.Loaded(answer.data)
            }
            is ApiResult.Failure -> fail(answer.error)
        }
        if (viewer.signedIn) loadWatching()
    }

    private fun fail(error: ApiError) {
        val cached = cache.read(cacheKey, CharacterProfile.serializer())
        val saved = StripDecision.savedView(error, cached?.value, cached?.savedAt, now(), CacheLimits.CHARACTER, connectivity.isConnected())
        phase = when {
            error is ApiError.Server && error.status == NOT_FOUND -> CharacterPhase.NotFound
            saved != null -> CharacterPhase.Loaded(saved.value, saved.since, saved.stale)
            phase is CharacterPhase.Loaded -> phase
            else -> CharacterPhase.Failed
        }
    }

    private suspend fun loadWatching() {
        val list = (service.watching() as? ApiResult.Success)?.data ?: return
        val key = (phase as? CharacterPhase.Loaded)?.profile?.character?.key
        watch = WatchState(watching = list.rows.any { it.key == key || (it.region == route.region && it.ruleset == route.ruleset && it.name.equals(route.name, ignoreCase = true)) })
    }

    /** The star: filled at once, the write behind it; a failure reverts and says so under the head. */
    suspend fun toggleWatch(haptics: RaidHaptics) {
        val current = watch ?: return
        if (current.busy) return
        val next = !current.watching
        watch = WatchState(watching = next, busy = true)
        val answer = if (next) service.watch(route.region, route.ruleset, route.name) else service.unwatch(route.region, route.ruleset, route.name)
        watch = when (answer) {
            is ApiResult.Success -> {
                if (next) haptics.success()
                WatchState(watching = next)
            }
            is ApiResult.Failure -> {
                haptics.error()
                WatchState(watching = current.watching, failed = true)
            }
        }
    }

    companion object {
        private const val CACHE_KEY = "character-"
        private const val NOT_FOUND = 404
    }
}
