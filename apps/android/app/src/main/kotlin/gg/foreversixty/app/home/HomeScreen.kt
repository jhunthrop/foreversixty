package gg.foreversixty.app.home

import android.content.Context
import android.content.Intent
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
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Text
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.core.net.toUri
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.compose.LifecycleEventEffect
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.CharacterSelectorBar
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FinishedSim
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.sims.SimsFeature
import gg.foreversixty.app.sims.SimsRoute
import gg.foreversixty.app.talents.TalentsFeature
import gg.foreversixty.app.talents.TalentCharacter
import gg.foreversixty.app.talents.TalentsModel
import gg.foreversixty.app.talents.TalentsPhase
import gg.foreversixty.app.talents.TalentsRoute
import java.io.File
import java.time.Instant
import kotlinx.coroutines.launch

/** The site pages Home opens in the browser. */
object HomeLinks {
    const val SETUP = "https://foreversixty.gg/setup"
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
    /** Lane F5: Sims, pushed from the Sims row; null while its feature switch is off. */
    private val simsFeature: SimsFeature? = null,
) {
    val sims: SimsFeature? get() = simsFeature?.takeIf { it.enabled }

    val isEnabled: Boolean get() = isOn()

    companion object {
        fun create(
            context: Context,
            config: ApiConfig,
            accountApi: AccountApi,
            client: ApiClient,
            talents: TalentsFeature?,
            isSimsOn: () -> Boolean = { false },
            mayLinkToPlans: suspend () -> Boolean = { false },
            isOn: () -> Boolean,
        ): HomeDependencies {
            val engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets))
            val catalog = ItemCatalog(StaticDataClient(engine), File(context.cacheDir, "static-data"))
            return HomeDependencies(
                accountApi = accountApi,
                client = client,
                catalog = catalog,
                cache = FileHomeCache(File(context.noBackupFilesDir, "home-answers")),
                talents = talents,
                isOn = isOn,
                simsFeature = SimsFeature(context, client, StaticDataClient(engine), catalog, isSimsOn, mayLinkToPlans),
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
            HomeCharacterScreen(HomeCharacter.from(character), sessionMe.finishedSims, dependencies, openBis, sessionMe.entitlements.serverSims) {
                CharacterSelectorBar(session)
            }
            return
        }
    }
    HomeFromMe(dependencies, openBis) { if (session != null) CharacterSelectorBar(session) }
}

/**
 * Home on its own `GET /v1/me`: the main character, before or without the session's copy. [topBar]
 * is the selector row, which reserves its height while the session's characters arrive.
 */
@Composable
private fun HomeFromMe(dependencies: HomeDependencies, openBis: () -> Unit, topBar: @Composable () -> Unit) {
    var attempt by remember { mutableStateOf(0) }
    val me by produceState<ApiResult<Me>?>(initialValue = null, attempt) { value = dependencies.accountApi.me() }
    val answer = me
    val character = (answer as? ApiResult.Success)?.data?.mainCharacter
    if (answer is ApiResult.Success && character != null) {
        HomeCharacterScreen(HomeCharacter.from(character), answer.data.finishedSims, dependencies, openBis, answer.data.entitlements.serverSims, topBar)
        return
    }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        topBar()
        when (answer) {
            null -> HomeSkeleton()
            is ApiResult.Failure -> Column(Modifier.padding(16.dp)) {
                Text(stringResource(R.string.home_answer_failed), style = FsType.lead, color = FsColor.wipe)
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, { attempt++ }, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
            }
            is ApiResult.Success -> Column(Modifier.padding(16.dp)) {
                Text(stringResource(R.string.signin_noChars_title), style = FsType.sec, color = FsColor.strong)
                Text(stringResource(R.string.signin_noChars_body), style = FsType.body, color = FsColor.muted)
            }
        }
    }
}

/** Home for one character: the answer, Your gear and the tooltip sheet. */
@Composable
fun HomeCharacterScreen(
    character: HomeCharacter,
    finishedSims: List<FinishedSim>,
    dependencies: HomeDependencies,
    openBis: () -> Unit,
    serverSims: Boolean = false,
    topBar: @Composable () -> Unit = {},
) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val store = remember(character.key) { HomeStore.live(dependencies.client, dependencies.cache) }
    val phase by store.phase.collectAsStateWithLifecycle()
    var refreshing by remember { mutableStateOf(false) }
    var showGear by rememberSaveable { mutableStateOf(false) }
    var showTalents by rememberSaveable { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(character.key) { store.load(character) }
    LifecycleEventEffect(Lifecycle.Event.ON_RESUME) { scope.launch { store.refreshIfStale(character) } }

    val home = phase.homeOrNull()
    val res = LocalResources.current
    // Lane F6: one Talents model per character, shared by the Talents row and the pushed screen.
    val talentsModel = remember(character.key) {
        dependencies.talents?.let { TalentsModel(it.source, TalentCharacter.from(character), it.catalog) }
    }
    LaunchedEffect(talentsModel) { talentsModel?.load() }
    val talentsPhase = talentsModel?.phase?.collectAsStateWithLifecycle()?.value
    val talentStatus = (talentsPhase as? TalentsPhase.Loaded)?.presenter?.status(res)
    val presenter = HomePresenter(character, finishedSims, HomeTime(Instant.now()), res, talentStatus)
    val offline = (phase as? HomePhase.Loaded)?.offlineSince?.let {
        stringResource(R.string.offline_strip, HomeTime(Instant.now()).stamp(it))
    }
    val openItem = { look: ItemLook, sourceLine: String?, wornId: Int? ->
        home?.let { tooltip = HomePresenter.tooltipRequest(look, it, character, res, sourceLine, wornId) }
    }

    val sims = dependencies.sims
    val talents = dependencies.talents
    if (sims != null && sims.navigator.showsList) {
        SimsRoute(sims, serverSims, onBack = { sims.navigator.showsList = false })
    } else if (showTalents && talents != null) {
        BackHandler { showTalents = false }
        TalentsRoute(character, talentsModel ?: TalentsModel(talents.source, TalentCharacter.from(character), talents.catalog)) {
            showTalents = false
        }
    } else if (showGear && home != null) {
        BackHandler { showGear = false }
        GearPage(home, character, dependencies.catalog) { look, wornId -> openItem(look, null, wornId) }
    } else {
        Column(Modifier.fillMaxSize().background(FsColor.bg)) {
            topBar()
            HomeAnswerPane(refreshing, onRefresh = {
                scope.launch {
                    refreshing = true
                    store.refresh(character)
                    refreshing = false
                }
            }) {
                HomeAnswer(phase, presenter, home?.dataBuild.orEmpty(), offline) { intent ->
                    when (intent) {
                        is HomeIntent.OpenItem -> openItem(intent.look, intent.sourceLine, intent.wornId)
                        HomeIntent.OpenGear -> showGear = true
                        HomeIntent.Retry -> scope.launch { store.refresh(character) }
                        is HomeIntent.OpenNext -> when (intent.kind) {
                            HomeNextRow.Kind.Sims -> sims?.navigator?.showsList = true
                            HomeNextRow.Kind.Talents -> showTalents = talents != null
                        }
                        is HomeIntent.Action -> when (intent.kind) {
                            HomeAction.Kind.GetAddon -> context.openInBrowser(HomeLinks.SETUP)
                            HomeAction.Kind.OpenPlanner -> context.openInBrowser(HomeLinks.PLANNER)
                            HomeAction.Kind.OpenList -> openBis()
                        }
                    }
                }
            }
        }
    }
    tooltip?.let { request ->
        ItemTooltipSheet(request, dependencies.catalog, onDismiss = { tooltip = null })
    }
}

/** The answer under pull to refresh, scrolling on its own below the selector. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun HomeAnswerPane(refreshing: Boolean, onRefresh: () -> Unit, content: @Composable () -> Unit) {
    PullToRefreshBox(
        isRefreshing = refreshing,
        onRefresh = onRefresh,
        modifier = Modifier.fillMaxSize(),
    ) {
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) { content() }
    }
}

/** Your gear, with the build's items file loaded for the slots the answer does not name. */
@Composable
private fun GearPage(home: CharacterHome, character: HomeCharacter, catalog: ItemCatalog, open: (ItemLook, Int?) -> Unit) {
    val res = LocalResources.current
    val probe = HomePresenter.tooltipRequest(ItemLook(0), home, character, res)
    val sources by produceState<ItemTooltipSources?>(null, home) { value = catalog.sources(probe) }
    val items = sources?.items.orEmpty()
    val blocker = GearModel.offHandBlocker(home, character.gear, items, ItemTooltipModel.band(probe, sources?.bis))
    Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
        GearList(
            rows = GearModel.rows(home, character.gear, items, res, blocker),
            intro = stringResource(R.string.gear_intro, HomePresenter.bandHead(home, character, res).orEmpty()),
            build = home.dataBuild,
            open = open,
        )
    }
}

private fun HomePhase.homeOrNull(): CharacterHome? = when (this) {
    is HomePhase.Loaded -> home
    is HomePhase.Failed -> cached
    HomePhase.Loading -> null
}

private fun Context.openInBrowser(url: String) {
    startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
}
