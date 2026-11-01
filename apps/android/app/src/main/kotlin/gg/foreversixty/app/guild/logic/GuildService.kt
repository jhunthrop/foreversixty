package gg.foreversixty.app.guild.logic

import gg.foreversixty.app.common.StripDecision
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.common.CacheLimits
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.CalendarFeed
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant
import kotlinx.serialization.KSerializer

/** The F16 calls (spec 5.8 Data). Live goes through [ApiClient]; tests pass a fake. */
interface GuildService {
    suspend fun home(guild: Long): ApiResult<GuildHome>
    suspend fun history(guild: Long, cursor: String?): ApiResult<GuildRaids>
    suspend fun calendarFeed(): ApiResult<CalendarFeed>
    suspend fun resetCalendarFeed(): ApiResult<CalendarFeed>
}

class LiveGuildService(private val client: ApiClient) : GuildService {
    override suspend fun home(guild: Long) = client.send(ApiEndpoints.getGuildHome(guild))
    override suspend fun history(guild: Long, cursor: String?) =
        client.send(ApiEndpoints.listGuildRaids(guild, limit = HistoryPager.PAGE_SIZE, cursor = cursor))
    override suspend fun calendarFeed() = client.send(ApiEndpoints.getCalendarFeed())
    override suspend fun resetCalendarFeed() = client.send(ApiEndpoints.resetCalendarFeed())
}

/** A screen's one load: the skeleton, the answer (possibly cached under the strip), or the error. */
sealed interface Loadable<out T> {
    data object Loading : Loadable<Nothing>
    /** [offlineSince] is set while the saved answer shows under a strip: offline, or [stale] (W3 §5.7). */
    data class Loaded<T>(val value: T, val offlineSince: Instant? = null, val stale: Boolean = false) : Loadable<T>
    data object Failed : Loadable<Nothing>
}

/**
 * Cache first, then the answer (the Raids tab's pattern): a refresh that cannot connect keeps the
 * cached answer with its time; any other failure with nothing on screen is the error state.
 */
class CacheFirst<T>(
    private val key: String,
    private val serializer: KSerializer<T>,
    private val cache: RaidCache,
    private val now: () -> Instant,
    private val connectivity: Connectivity = Connectivity.Unknown,
    private val fetch: suspend () -> ApiResult<T>,
) {
    var state: Loadable<T> by mutableStateOf(Loadable.Loading)
        private set

    suspend fun load() {
        if (state == Loadable.Loading) cache.read(key, serializer)?.let { state = Loadable.Loaded(it.value) }
        state = when (val result = fetch()) {
            is ApiResult.Success -> {
                cache.write(key, serializer, result.data, now())
                Loadable.Loaded(result.data)
            }
            is ApiResult.Failure -> {
                val cached = cache.read(key, serializer)
                val saved = StripDecision.savedView(result.error, cached?.value, cached?.savedAt, now(), CacheLimits.GUILD, connectivity.isConnected())
                when {
                    saved != null -> Loadable.Loaded(saved.value, saved.since, saved.stale)
                    state is Loadable.Loaded -> state
                    else -> Loadable.Failed
                }
            }
        }
    }
}
