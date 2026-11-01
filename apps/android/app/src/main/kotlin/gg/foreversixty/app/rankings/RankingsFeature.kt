package gg.foreversixty.app.rankings

import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.search.SearchRankings
import gg.foreversixty.app.common.Connectivity
import android.content.Context
import android.content.Intent
import android.content.res.Resources
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.LiveRankingsService
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsService
import gg.foreversixty.app.rankings.logic.RankingsStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import java.io.File
import java.time.Instant
import java.time.ZoneId

/** Where a rankings or character link lands on Home's stack (spec §3.4). */
class RankingsNavigator {
    /** Rankings is pushed on Home. */
    var showsBoard by mutableStateOf(false)

    /** The encounter a link named, read once by the board. */
    var encounter by mutableStateOf<String?>(null)

    /** A character page over Home (or over Rankings). */
    var character by mutableStateOf<CharacterRoute?>(null)

    /** Set when a link asks for Home; the scaffold selects it and clears this. */
    var wantsHome by mutableStateOf(false)

    fun open(link: DeepLink): Boolean {
        when (link) {
            is DeepLink.Rankings -> {
                showsBoard = true
                encounter = link.encounter
                character = null
            }
            is DeepLink.Player -> character = CharacterRoute(link.region, link.ruleset, link.name)
            else -> return false
        }
        wantsHome = true
        return true
    }
}

/** F11's dependencies, built once in the app graph ([create]); tests pass a fake service and cache. */
class RankingsFeature(
    val service: RankingsService,
    /** Cleared at sign-out with the other features' caches. */
    val cache: RaidCache,
    val specs: SpecCatalog,
    /** The class items files for a character's gear and the tooltip sheet. */
    val catalog: ItemCatalog,
    private val build: suspend () -> String?,
    private val isAvailable: () -> Boolean,
    val now: () -> Instant = Instant::now,
    val zone: ZoneId = ZoneId.systemDefault(),
    /** Lane F15: whether a failed refresh is offline or stale. */
    val connectivity: Connectivity = Connectivity.Unknown,
) {
    val navigator = RankingsNavigator()

    /** Lane F15: Builds seen's `Open in planner` (F13); null while `planner` is off. */
    var openPlanner: ((DeepLink.Planner) -> Unit)? = null

    /** Lane F15: `Search again` on F11.7 opens Search with the name in the field; null goes back. */
    var searchAgain: ((String) -> Unit)? = null

    /** Search's way in (lane F15): the board, a character page, the Watching group. */
    fun searchLink(): SearchRankings = SearchRankings(
        open = { link -> navigator.open(link) },
        openBoard = {
            navigator.character = null
            navigator.showsBoard = true
            navigator.wantsHome = true
        },
        watching = {
            (service.watching() as? ApiResult.Success)?.data?.rows.orEmpty().map {
                CharacterSearchRow(it.key, it.region, it.ruleset, it.name, it.className, it.level)
            }
        },
    )

    /** Rankings ships dark behind `rankings` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun iconBuild(): String? = build()

    fun store(viewer: RankingsViewer) = RankingsStore(service, cache, viewer, specs, connectivity, now)

    /**
     * Home's Rankings row (spec F11.1): one read of the current character's board with its
     * defaults; null while it cannot be read, so the row keeps an empty sub line.
     */
    suspend fun homeLine(viewer: RankingsViewer, res: Resources): String? {
        val store = store(viewer)
        store.load()
        val board = (store.phase as? BoardPhase.Loaded)?.board ?: return null
        val filters = store.filters ?: return null
        return RankingsCopy(res, zone).homeLine(board.characterRank, board.total, filters, store.encounterNames)
    }

    companion object {
        fun create(
            context: Context,
            apiClient: ApiClient,
            statics: StaticDataClient,
            specs: SpecCatalog,
            catalog: ItemCatalog,
            isAvailable: () -> Boolean,
            connectivity: Connectivity = Connectivity.Unknown,
        ) = RankingsFeature(
            service = LiveRankingsService(apiClient),
            cache = FileRaidCache(File(context.filesDir, "rankings")),
            specs = specs,
            catalog = catalog,
            build = { statics.activeBuild().dataOrNull()?.build },
            isAvailable = isAvailable,
            connectivity = connectivity,
        )

        /** The system share sheet with a web link. */
        fun share(context: Context, url: String) {
            val send = Intent(Intent.ACTION_SEND).setType("text/plain").putExtra(Intent.EXTRA_TEXT, url)
            context.startActivity(Intent.createChooser(send, null))
        }
    }
}
