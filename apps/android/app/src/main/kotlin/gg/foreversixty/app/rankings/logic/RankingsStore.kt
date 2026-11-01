package gg.foreversixty.app.rankings.logic

import gg.foreversixty.app.common.StripDecision
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.common.CacheLimits
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.CharacterRank
import gg.foreversixty.app.model.PhaseBoundary
import gg.foreversixty.app.model.Phases
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankedEncounters
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.model.RankingsPage
import gg.foreversixty.app.model.WatchedCharacter
import gg.foreversixty.app.model.WatchingList
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidCache
import java.time.Instant
import kotlinx.serialization.KSerializer

/** The calls F11 makes. Live goes through [ApiClient]; tests pass a fake. */
interface RankingsService {
    suspend fun rankings(query: RankingsQuery): ApiResult<RankingsPage>
    suspend fun encounters(): ApiResult<RankedEncounters>
    suspend fun phases(): ApiResult<Phases>
    suspend fun character(region: String, ruleset: String, name: String): ApiResult<CharacterProfile>
    suspend fun watching(): ApiResult<WatchingList>
    suspend fun watch(region: String, ruleset: String, name: String): ApiResult<WatchedCharacter>
    suspend fun unwatch(region: String, ruleset: String, name: String): ApiResult<Unit>
}

class LiveRankingsService(private val client: ApiClient) : RankingsService {
    override suspend fun rankings(query: RankingsQuery) = client.send(
        ApiEndpoints.getRankings(
            encounter = query.encounter, metric = query.metric, spec = query.spec, `class` = query.classSlug, phase = query.phase,
            region = query.region, ruleset = query.ruleset, faction = query.faction, since = query.since,
            distinct = query.distinct, character = query.character, page = query.page,
        ),
    )
    override suspend fun encounters() = client.send(ApiEndpoints.getRankingEncounters())
    override suspend fun phases() = client.send(ApiEndpoints.getPhases())
    override suspend fun character(region: String, ruleset: String, name: String) = client.send(ApiEndpoints.getCharacter(region, ruleset, name))
    override suspend fun watching() = client.send(ApiEndpoints.getWatching())
    override suspend fun watch(region: String, ruleset: String, name: String) = client.send(ApiEndpoints.watchCharacter(region, ruleset, name))
    override suspend fun unwatch(region: String, ruleset: String, name: String) = client.send(ApiEndpoints.unwatchCharacter(region, ruleset, name))
}

/** What the board draws (spec F11.1). */
data class Board(
    val rows: List<RankingRow>,
    val total: Int,
    val page: Int,
    val deduped: Boolean,
    val characterRank: CharacterRank?,
    /** Set while the cached board shows under a strip: offline, or [stale] (W3 §5.7). */
    val offlineSince: Instant?,
    /** False once a page read adds no row (fixture mode answers the same page). */
    val exhausted: Boolean = false,
    val stale: Boolean = false,
) {
    val hasMore: Boolean get() = !exhausted && offlineSince == null && rows.size < total

    companion object {
        fun of(page: RankingsPage, offlineSince: Instant?) =
            Board(page.rows, page.total, page.page, page.deduped, page.characterRank, offlineSince)
    }
}

sealed interface BoardPhase {
    data object Loading : BoardPhase
    data class Loaded(val board: Board) : BoardPhase
    data object Failed : BoardPhase
}

/**
 * The rankings board: its defaults (the character's cohort, ruleset, role's measure, the current
 * phase, the encounter it last killed), page one, then pages of 100 as the list scrolls. Every
 * read is cached for offline and cleared at sign-out with the feature's cache.
 */
class RankingsStore(
    private val service: RankingsService,
    private val cache: RaidCache,
    val viewer: RankingsViewer,
    private val specs: SpecCatalog,
    private val connectivity: Connectivity = Connectivity.Unknown,
    private val now: () -> Instant,
) {
    var phase: BoardPhase by mutableStateOf(BoardPhase.Loading)
        private set

    var filters: RankingsFilters? by mutableStateOf(null)
        private set

    var encounters: List<RankedEncounter> by mutableStateOf(emptyList())
        private set

    var phases: List<PhaseBoundary> by mutableStateOf(emptyList())
        private set

    /** The last next-page read failed: the board ends with the error and Try again. */
    var pageFailed: Boolean by mutableStateOf(false)
        private set

    /** The account's watch list by key, for a row's `Watch` / `Stop watching`; empty signed out. */
    var watched: Set<String> by mutableStateOf(emptySet())
        private set

    private var defaults: RankingsFilters? = null

    val encounterNames: Map<String, String> get() = encounters.associate { it.slug to it.name }

    private val board: Board? get() = (phase as? BoardPhase.Loaded)?.board

    /** Reads the defaults once (aimed at [encounter] when a link names one), then page one. */
    suspend fun load(encounter: String? = null) {
        if (filters == null) {
            val start = start(encounter) ?: run {
                phase = BoardPhase.Failed
                return
            }
            defaults = start
            filters = start
        }
        readBoard(keepOnFailure = true)
        if (viewer.signedIn) {
            (service.watching() as? ApiResult.Success)?.data?.let { list -> watched = list.rows.map { it.key }.toSet() }
        }
    }

    /** A row's long-press `Watch` or `Stop watching`; the list changes only when the write lands. */
    suspend fun toggleWatch(row: RankingRow): Boolean {
        val route = CharacterRoute.ofKey(row.player.key, row.player.name) ?: return false
        val key = row.player.key
        val watching = key in watched
        val answer = if (watching) service.unwatch(route.region, route.ruleset, route.name) else service.watch(route.region, route.ruleset, route.name)
        if (answer is ApiResult.Failure) return false
        watched = if (watching) watched - key else watched + key
        return true
    }

    /** A filter changed (every choice applies at once): page one of the new board. */
    suspend fun apply(next: RankingsFilters) {
        if (next == filters) return
        filters = next
        readBoard(keepOnFailure = false)
    }

    suspend fun reset() {
        defaults?.let { apply(it) }
    }

    /** The next page of 100, when the board has one. */
    suspend fun loadMore() {
        val current = board ?: return
        val f = filters ?: return
        if (!current.hasMore) return
        val next = (service.rankings(Rankings.query(f, characterKey, current.page + 1)) as? ApiResult.Success)?.data
        pageFailed = next == null
        next ?: return
        val rows = Rankings.merge(current.rows, next.rows)
        phase = BoardPhase.Loaded(current.copy(rows = rows, total = next.total, page = next.page, exhausted = rows.size == current.rows.size))
    }

    private val characterKey: String? get() = viewer.current?.key

    private suspend fun start(encounter: String?): RankingsFilters? {
        phases = cached(PHASES_KEY, Phases.serializer()) { service.phases() }?.phases.orEmpty()
        encounters = cached(ENCOUNTERS_KEY, RankedEncounters.serializer()) { service.encounters() }?.encounters.orEmpty()
        val slug = encounter?.let(::encounterSlug) ?: defaultEncounter() ?: return null
        return Rankings.defaults(viewer, specs, phases, now(), slug)
    }

    /** A link's encounter, by slug or id; one the API does not list is used as written. */
    private fun encounterSlug(link: String): String = encounters.firstOrNull { it.slug == link || it.id.toString() == link }?.slug ?: link

    /** The one the current character most recently killed, from its character card; else the first listed. */
    private suspend fun defaultEncounter(): String? {
        val character = viewer.current
        val history = character?.let { c ->
            (service.character(c.region, c.ruleset, c.key.substringAfterLast('/')) as? ApiResult.Success)?.data?.history
        }.orEmpty()
        return Rankings.defaultEncounter(history, encounters)
    }

    private suspend fun readBoard(keepOnFailure: Boolean) {
        val f = filters ?: return
        val key = boardKey(f)
        when (val answer = service.rankings(Rankings.query(f, characterKey, 1))) {
            is ApiResult.Success -> {
                cache.write(key, RankingsPage.serializer(), answer.data, now())
                pageFailed = false
                phase = BoardPhase.Loaded(Board.of(answer.data, null))
            }
            is ApiResult.Failure -> {
                val cached = cache.read(key, RankingsPage.serializer())
                val saved = StripDecision.savedView(answer.error, cached?.value, cached?.savedAt, now(), CacheLimits.RANKINGS, connectivity.isConnected())
                phase = when {
                    saved != null -> BoardPhase.Loaded(Board.of(saved.value, saved.since).copy(stale = saved.stale))
                    keepOnFailure && board != null -> phase
                    else -> BoardPhase.Failed
                }
            }
        }
    }

    /** A small read that keeps its last answer for offline use. */
    private suspend fun <T> cached(key: String, serializer: KSerializer<T>, read: suspend () -> ApiResult<T>): T? =
        when (val answer = read()) {
            is ApiResult.Success -> answer.data.also { cache.write(key, serializer, it, now()) }
            is ApiResult.Failure -> cache.read(key, serializer)?.value
        }

    private fun boardKey(f: RankingsFilters): String = BOARD_KEY + Rankings.query(f, characterKey, 1).hashCode().toUInt().toString(HEX)

    companion object {
        private const val PHASES_KEY = "rankings-phases"
        private const val ENCOUNTERS_KEY = "rankings-encounters"
        private const val BOARD_KEY = "rankings-board-"
        private const val HEX = 16
    }
}
