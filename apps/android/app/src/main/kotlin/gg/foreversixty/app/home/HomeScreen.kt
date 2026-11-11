package gg.foreversixty.app.home

import gg.foreversixty.app.designsystem.screenGround
import android.content.Context
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.compose.LifecycleEventEffect
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.activity.ActivityFeature
import gg.foreversixty.app.home.flows.HomeExits
import gg.foreversixty.app.home.flows.HomeFlowContent
import gg.foreversixty.app.home.flows.LiveHomeFlowService
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.rankings.RankingsFeature
import gg.foreversixty.app.rankings.RankingsReports
import gg.foreversixty.app.rankings.RankingsRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.search.ui.RootSearchBar
import gg.foreversixty.app.reports.ReportsFeature
import gg.foreversixty.app.reports.ReportsRoute
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.sims.SimsFeature
import gg.foreversixty.app.sims.SimsRoute
import gg.foreversixty.app.talents.TalentsFeature
import gg.foreversixty.app.talents.TalentCharacter
import gg.foreversixty.app.talents.TalentsModel
import java.io.File
import kotlinx.coroutines.launch

/** The site pages Home opens in the browser. */
object HomeLinks {
    const val PLANNER = "https://foreversixty.gg/planner"
}

/** What Home needs from the app: the API, the static data, the cache and the feature switch. */
class HomeDependencies(
    val accountApi: AccountApi,
    val client: ApiClient,
    val catalog: ItemCatalog,
    val cache: HomeCache,
    /** Lane F6: the Talents screen's files; null while its feature switch is off. */
    val talents: TalentsFeature? = null,
    private val isOn: () -> Boolean,
    /** Lane F5: Sims, pushed from Home; null while its feature switch is off. */
    private val simsFeature: SimsFeature? = null,
    /** Lane F10: Your reports, pushed over Home; null while its feature switch is off. */
    private val reportsFeature: ReportsFeature? = null,
    /** Lane F11: Rankings and character pages, pushed over Home; null while its switch is off. */
    private val rankingsFeature: RankingsFeature? = null,
    /** Lane F15: whether a failed refresh is offline or stale. */
    val connectivity: Connectivity = Connectivity.Unknown,
) {
    val sims: SimsFeature? get() = simsFeature?.takeIf { it.enabled }

    val reports: ReportsFeature? get() = reportsFeature?.takeIf { it.enabled }

    val rankings: RankingsFeature? get() = rankingsFeature?.takeIf { it.enabled }

    val isEnabled: Boolean get() = isOn()

    companion object {
        fun create(
            context: Context,
            config: ApiConfig,
            accountApi: AccountApi,
            client: ApiClient,
            talents: TalentsFeature?,
            statics: StaticDataStore,
            isSimsOn: () -> Boolean = { false },
            isReportsOn: () -> Boolean = { false },
            isRankingsOn: () -> Boolean = { false },
            mayLinkToPlans: suspend () -> Boolean = { false },
            connectivity: Connectivity = Connectivity.of(context),
            isOn: () -> Boolean,
        ): HomeDependencies {
            val engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets))
            val catalog = ItemCatalog(statics)
            val sims = SimsFeature(context, client, StaticDataClient(engine), catalog, isSimsOn, mayLinkToPlans, connectivity)
            return HomeDependencies(
                accountApi = accountApi,
                client = client,
                catalog = catalog,
                cache = FileHomeCache(File(context.noBackupFilesDir, "home-answers")),
                talents = talents,
                isOn = isOn,
                simsFeature = sims,
                reportsFeature = ReportsFeature.create(context, client, StaticDataClient(engine), sims.specs, isReportsOn, connectivity),
                rankingsFeature = RankingsFeature.create(context, client, StaticDataClient(engine), sims.specs, catalog, isRankingsOn, connectivity),
                connectivity = connectivity,
            )
        }
    }
}

/**
 * The Home tab: finds the character, then shows its answer. Signed in ([session], lane F1), it
 * follows the session's current character under the selector; without a session it shows the
 * main. [openBis] switches to the BiS tab (lane F3 owns it).
 */
@Composable
fun HomeRoute(
    dependencies: HomeDependencies,
    session: SessionManager? = null,
    openBis: () -> Unit = {},
    /** Lane F12: the bar's search button; null while the `search` switch is off. */
    onSearch: (() -> Unit)? = null,
    /** Lane F10: Your reports' signed-out state signs in here. */
    signIn: () -> Unit = {},
    /** The Raid tab's raids, the Activity preview and Home's ways out (spec 2026-11-08 4.1). */
    flows: HomeFlows = HomeFlows(),
    placeholder: @Composable () -> Unit,
) {
    if (!dependencies.isEnabled) {
        placeholder()
        return
    }
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val sessionMe = phase?.me
    if (session != null && sessionMe != null) {
        val character = session.currentCharacter(phase, picked)
        if (character != null) {
            val reports = HomeReports(ReportsViewer.of(sessionMe, character.key), signIn, RankingsViewer.of(sessionMe, character.key))
            HomeCharacterScreen(HomeCharacter.from(character), dependencies, openBis, sessionMe.entitlements.serverSims, reports, flows) {
                RootSearchBar(onSearch) { ProfileBar(session) }
            }
            return
        }
    }
    HomeFromMe(dependencies, openBis, signIn, flows) { RootSearchBar(onSearch) { ProfileBar(session) } }
}

/**
 * Home on its own `GET /v1/me`: the main character, before or without the session's copy. [topBar]
 * is the selector row, which reserves its height while the session's characters arrive.
 */
@Composable
private fun HomeFromMe(dependencies: HomeDependencies, openBis: () -> Unit, signIn: () -> Unit, flows: HomeFlows, topBar: @Composable () -> Unit) {
    var attempt by remember { mutableStateOf(0) }
    val me by produceState<ApiResult<Me>?>(initialValue = null, attempt) { value = dependencies.accountApi.me() }
    val answer = me
    val character = (answer as? ApiResult.Success)?.data?.mainCharacter
    if (answer is ApiResult.Success && character != null) {
        val reports = HomeReports(ReportsViewer.of(answer.data, character.key), signIn, RankingsViewer.of(answer.data, character.key))
        HomeCharacterScreen(HomeCharacter.from(character), dependencies, openBis, answer.data.entitlements.serverSims, reports, flows, topBar)
        return
    }
    Column(Modifier.fillMaxSize().screenGround()) {
        topBar()
        when (answer) {
            null -> Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator(color = FsColor.gold) }
            is ApiResult.Failure -> Column(Modifier.padding(16.dp)) {
                Text(stringResource(R.string.home_answer_failed), style = ScaledType.lead, color = FsColor.wipe)
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, { attempt++ }, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
            }
            is ApiResult.Success -> Column(Modifier.padding(16.dp)) {
                Text(stringResource(R.string.signin_noChars_title), style = ScaledType.sec, color = FsColor.strong)
                Text(stringResource(R.string.signin_noChars_body), style = ScaledType.body, color = FsColor.muted)
            }
        }
    }
}

/** Home for one character: the flows, the screens pushed over them and the tooltip sheet. */
@Composable
fun HomeCharacterScreen(
    character: HomeCharacter,
    dependencies: HomeDependencies,
    openBis: () -> Unit,
    serverSims: Boolean = false,
    reports: HomeReports = HomeReports(),
    flows: HomeFlows = HomeFlows(),
    topBar: @Composable () -> Unit = {},
) {
    val scope = rememberCoroutineScope()
    val store = remember(character.key) { HomeStore.live(dependencies.client, dependencies.cache, dependencies.connectivity) }
    val phase by store.phase.collectAsStateWithLifecycle()
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(character.key) { store.load(character) }
    LifecycleEventEffect(Lifecycle.Event.ON_RESUME) { scope.launch { store.refreshIfStale(character) } }

    val home = phase.homeOrNull()
    val res = LocalResources.current
    val reportsFeature = dependencies.reports
    val rankingsFeature = dependencies.rankings
    val openItem = { look: ItemLook, sourceLine: String?, wornId: Int? ->
        home?.let { tooltip = HomePresenter.tooltipRequest(look, it, character, res, sourceLine, wornId) }
    }

    val sims = dependencies.sims
    if (sims != null && sims.navigator.showsList) {
        SimsRoute(sims, serverSims, onBack = { sims.navigator.showsList = false })
    } else if (reportsFeature != null && reportsFeature.navigator.showsList) {
        ReportsRoute(reportsFeature, reports.viewer, reports.signIn, onBack = { reportsFeature.navigator.showsList = false })
    } else if (rankingsFeature != null && (rankingsFeature.navigator.showsBoard || rankingsFeature.navigator.character != null)) {
        val reportsLink = reportsFeature?.let { RankingsReports(it, reports.viewer) }
        RankingsRoute(rankingsFeature, reports.rankings, reportsLink, onBack = { rankingsFeature.navigator.showsBoard = false })
    } else {
        Column(Modifier.fillMaxSize().screenGround()) {
            val service = remember(dependencies) { LiveHomeFlowService(dependencies.client, dependencies.catalog) }
            val exits = flows.exits.copy(
                openGear = openBis,
                openItem = { line -> openItem(line.look, line.sourceLine, line.wornId) },
                openSims = { sims?.navigator?.showsList = true },
            )
            HomeFlowContent(character, store, phase, service, flows.raids, flows.activity, exits, topBar)
        }
    }
    tooltip?.let { request ->
        ItemTooltipSheet(request, dependencies.catalog, onDismiss = { tooltip = null })
    }
}

/** What the flows Home reads beside its own answer: the Raid tab's feature, Activity, and the shell's exits. */
data class HomeFlows(val raids: RaidsFeature? = null, val activity: ActivityFeature? = null, val exits: HomeExits = HomeExits())

/**
 * Who reads Home's pushed screens: Your reports (lane F10) and how its signed-out state signs in,
 * and Rankings (lane F11).
 */
data class HomeReports(
    val viewer: ReportsViewer = ReportsViewer.SignedOut,
    val signIn: () -> Unit = {},
    val rankings: RankingsViewer = RankingsViewer.SignedOut,
)

private fun HomePhase.homeOrNull(): CharacterHome? = when (this) {
    is HomePhase.Loaded -> home
    is HomePhase.Failed -> cached
    HomePhase.Loading -> null
}
