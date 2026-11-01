package gg.foreversixty.app.guild.logic

import gg.foreversixty.app.common.Connectivity
import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidLinks
import java.time.Instant

/** Where a history row goes (spec F16.3). */
sealed interface HistoryTarget {
    /** Your raid night (F4.22), on the Raids stack. */
    data class Night(val raidId: Long) : HistoryTarget

    /** The report on the web (F10 is not built yet). */
    data class Report(val url: String) : HistoryTarget
}

/** Read-only facts of one logged night; [mine] is the account's character keys. */
class NightFacts(val row: RaidNightRow, mine: Set<String>) {
    /** The night starts at its log's creation minus its length. */
    val startedAt: Instant = Instant.parse(row.createdAt).minusMillis(row.durationMs)

    val bosses: Int = row.fights.map { it.name }.distinct().size
    val kills: Int get() = row.killCount
    val wipes: Int get() = row.wipeCount

    /** The night's size, never below a kill's player count. */
    val size: Int = maxOf(row.raiders, row.fights.filter { it.kill }.maxOfOrNull { it.players } ?: 0)

    /** A dungeon of ten or fewer reads `players`, not `raiders`. */
    val isDungeon: Boolean get() = size <= DUNGEON_SIZE

    private val totalMinutes: Long = row.durationMs / MILLIS_PER_MINUTE
    val hours: Int = (totalMinutes / MINUTES_PER_HOUR).toInt()
    val minutes: Int = (totalMinutes % MINUTES_PER_HOUR).toInt()

    val wasThere: Boolean = row.present.any { it.characterKey in mine }

    val target: HistoryTarget = row.raidId?.takeIf { wasThere }?.let { HistoryTarget.Night(it) }
        ?: HistoryTarget.Report(RaidLinks.report(row.id))

    private companion object {
        const val DUNGEON_SIZE = 10
        const val MILLIS_PER_MINUTE = 60_000L
        const val MINUTES_PER_HOUR = 60L
    }
}

/** Every history sentence built from data. */
class HistoryCopy(private val res: Resources, private val clock: RaidClock) {
    /** `1 boss · 1 kill · no wipes`. */
    fun fights(facts: NightFacts): String = listOf(
        res.getQuantityString(R.plurals.history_bosses, facts.bosses, facts.bosses),
        res.getQuantityString(R.plurals.history_kills, facts.kills, facts.kills),
        wipes(facts),
    ).joinToString(SEPARATOR)

    private fun wipes(facts: NightFacts): String =
        if (facts.wipes == 0) res.getString(R.string.history_noWipes) else res.getQuantityString(R.plurals.history_wipes, facts.wipes, facts.wipes)

    /** `38 raiders` or `5 players`. */
    fun size(facts: NightFacts): String =
        res.getQuantityString(if (facts.isDungeon) R.plurals.history_players else R.plurals.history_raiders, facts.size, facts.size)

    /** `1 h 47 m`, or `52 m` under an hour. */
    fun length(facts: NightFacts): String =
        if (facts.hours > 0) res.getString(R.string.history_length, facts.hours, facts.minutes) else res.getString(R.string.history_lengthMinutes, facts.minutes)

    private fun lengthWords(facts: NightFacts): String = listOfNotNull(
        facts.hours.takeIf { it > 0 }?.let { res.getQuantityString(R.plurals.history_a11y_hours, it, it) },
        facts.minutes.takeIf { it > 0 || facts.hours == 0 }?.let { res.getQuantityString(R.plurals.history_a11y_minutes, it, it) },
    ).joinToString(" ")

    /** Spec 6: `Onyxia's Lair, Thursday 17 December, 1 boss, 1 kill, no wipes, 38 raiders, 1 hour 47 minutes. You were there.` */
    fun a11y(facts: NightFacts): String {
        val sentence = listOf(
            facts.row.zone,
            clock.longDate(facts.startedAt),
            res.getQuantityString(R.plurals.history_bosses, facts.bosses, facts.bosses),
            res.getQuantityString(R.plurals.history_kills, facts.kills, facts.kills),
            wipes(facts),
            size(facts),
            lengthWords(facts),
        ).joinToString(", ")
        return if (facts.wasThere) "$sentence. ${res.getString(R.string.history_a11y_there)}" else sentence
    }

    companion object {
        const val SEPARATOR = " · "
    }
}

/** The `Load more` row's state. */
enum class MoreState { Idle, Loading, Failed, End }

/** `GET /v1/guilds/{id}/raids` paged by `next_cursor`, 20 a page; the first page is cached. */
class HistoryPager(
    val guild: Long,
    private val service: GuildService,
    cache: RaidCache,
    connectivity: Connectivity = Connectivity.Unknown,
    now: () -> Instant,
) {
    private val first = CacheFirst("guild-$guild-history", GuildRaids.serializer(), cache, now, connectivity) { service.history(guild, null) }
    private var extra by mutableStateOf<List<RaidNightRow>>(emptyList())
    private var cursor: String? = null

    var more: MoreState by mutableStateOf(MoreState.Idle)
        private set

    val state: Loadable<List<RaidNightRow>>
        get() = when (val page = first.state) {
            Loadable.Loading -> Loadable.Loading
            Loadable.Failed -> Loadable.Failed
            is Loadable.Loaded -> Loadable.Loaded(page.value.rows + extra, page.offlineSince, page.stale)
        }

    suspend fun load() {
        first.load()
        val page = (first.state as? Loadable.Loaded)?.value ?: return
        extra = emptyList()
        cursor = page.nextCursor
        more = if (cursor == null) MoreState.End else MoreState.Idle
    }

    suspend fun loadMore() {
        val next = cursor ?: return
        if (more == MoreState.Loading) return
        more = MoreState.Loading
        more = when (val result = service.history(guild, next)) {
            is ApiResult.Success -> {
                extra = extra + result.data.rows
                cursor = result.data.nextCursor
                if (cursor == null) MoreState.End else MoreState.Idle
            }
            is ApiResult.Failure -> MoreState.Failed
        }
    }

    companion object {
        const val PAGE_SIZE = 20
    }
}
