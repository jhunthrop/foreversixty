package gg.foreversixty.app.home.flows

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.activity.ActivityFeature
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.gear.WhereToGo
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeSection
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import java.time.Duration
import java.time.Instant

/** The calls Home makes beside the upgrades answer; live through [ApiClient], fakes in tests. */
interface HomeFlowService {
    suspend fun sections(characterKey: String): List<HomeSection>?
    suspend fun bestRun(home: CharacterHome): RunPick?
}

class LiveHomeFlowService(private val client: ApiClient, private val catalog: ItemCatalog) : HomeFlowService {
    override suspend fun sections(characterKey: String): List<HomeSection>? =
        (client.send(ApiEndpoints.getHome(characterKey)) as? ApiResult.Success)?.data?.sections

    /** Best run from the cached list and loot files (spec F5): the list's band for the answer's band and preset. */
    override suspend fun bestRun(home: CharacterHome): RunPick? {
        val spec = home.spec ?: return null
        val bis = catalog.bis(home.dataBuild, spec) ?: return null
        val loot = catalog.loot(home.dataBuild) ?: return null
        val band = bis.bands.firstOrNull { it.band == home.band && it.preset == home.preset } ?: return null
        return WhereToGo.bestRun(WhereToGo.picks(home.upgrades, band.slots, WhereToGo.sources(loot)))
    }
}

/**
 * What Home has loaded for one character beside its upgrades: the sections, the guild's raids (through
 * the Raid tab's own model and cache), the Activity preview and the Best run pick.
 */
class HomeFlowState(
    private val service: HomeFlowService,
    private val raids: RaidsTabModel?,
    private val activity: ActivityFeature?,
    private val now: () -> Instant,
) {
    var sections: List<HomeSection> by mutableStateOf(emptyList())
        private set
    var bestRun: RunPick? by mutableStateOf(null)
        private set
    var rows: List<ActivityRow> by mutableStateOf(emptyList())
        private set

    val raidContent: RaidsTabContent? get() = (raids?.state as? RaidsTabState.Loaded)?.content

    /** The raid to answer: the next signable one. */
    val nextRaid: RaidView? get() = raidContent?.let { PlannedRaids(it.planned, now()).next }

    /** The raid pulled under four hours ago (spec 4.1 raid live). */
    val liveRaid: RaidView? get() = raidContent?.planned?.firstOrNull { isLive(it, now()) }

    suspend fun load(character: HomeCharacter) {
        service.sections(character.key)?.let { sections = it }
        activity?.preview()?.let { rows = it }
        raids?.let {
            it.currentKey = character.key
            it.load()
        }
    }

    suspend fun refreshBestRun(home: CharacterHome?) {
        bestRun = home?.let { service.bestRun(it) }
    }

    companion object {
        private val LIVE_FOR: Duration = Duration.ofHours(4)

        fun isLive(raid: RaidView, now: Instant): Boolean {
            if (raid.status == "cancelled" || raid.status == "done") return false
            val start = RaidFacts(raid, now).startsAt
            return !now.isBefore(start) && now.isBefore(start.plus(LIVE_FOR))
        }
    }
}
