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
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FinishedSim
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.StaticDataClient
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
    private val isOn: () -> Boolean,
) {
    val isEnabled: Boolean get() = isOn()

    companion object {
        fun create(
            context: Context,
            config: ApiConfig,
            accountApi: AccountApi,
            client: ApiClient,
            appConfig: AppConfigRepository,
        ): HomeDependencies {
            val engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets))
            return HomeDependencies(
                accountApi = accountApi,
                client = client,
                catalog = ItemCatalog(StaticDataClient(engine), File(context.cacheDir, "static-data")),
                cache = FileHomeCache(File(context.noBackupFilesDir, "home-answers")),
                // The switch is on in fixture builds (debug) and wherever the API turns it on.
                isOn = { config.mode == ApiMode.Fixture || appConfig.isEnabled(Feature.CharacterHome) },
            )
        }
    }
}

/**
 * The Home tab: finds the character, then shows its answer. Lane F1's selector replaces the
 * `/v1/me` main-character pick with the player's current character. [openBis] switches to the
 * BiS tab (lane F3 owns it).
 */
@Composable
fun HomeRoute(dependencies: HomeDependencies, openBis: () -> Unit = {}, placeholder: @Composable () -> Unit) {
    if (!dependencies.isEnabled) {
        placeholder()
        return
    }
    var attempt by remember { mutableStateOf(0) }
    val me by produceState<ApiResult<Me>?>(initialValue = null, attempt) { value = dependencies.accountApi.me() }
    when (val answer = me) {
        null -> Column(Modifier.fillMaxSize().background(FsColor.bg)) { HomeSkeleton() }
        is ApiResult.Failure -> Column(Modifier.fillMaxSize().background(FsColor.bg).padding(16.dp)) {
            Text(stringResource(R.string.home_answer_failed), style = HomeType.lead, color = FsColor.wipe)
            SecondaryButton(stringResource(R.string.common_tryAgain), onClick = { attempt++ }, withRetryGlyph = true)
        }
        is ApiResult.Success -> {
            val character = answer.data.mainCharacter
            if (character == null) {
                Column(Modifier.fillMaxSize().background(FsColor.bg).padding(16.dp)) {
                    Text(stringResource(R.string.signin_noChars_title), style = HomeType.sectionTitle, color = FsColor.strong)
                    Text(stringResource(R.string.signin_noChars_body), style = HomeType.body, color = FsColor.muted)
                }
            } else {
                HomeCharacterScreen(HomeCharacter.from(character), answer.data.finishedSims, dependencies, openBis)
            }
        }
    }
}

/** Home for one character: the answer, Your gear and the tooltip sheet. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeCharacterScreen(
    character: HomeCharacter,
    finishedSims: List<FinishedSim>,
    dependencies: HomeDependencies,
    openBis: () -> Unit,
) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val store = remember(character.key) { HomeStore.live(dependencies.client, dependencies.cache) }
    val phase by store.phase.collectAsStateWithLifecycle()
    var refreshing by remember { mutableStateOf(false) }
    var showGear by rememberSaveable { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(character.key) { store.load(character) }
    LifecycleEventEffect(Lifecycle.Event.ON_RESUME) { scope.launch { store.refreshIfStale(character) } }

    val home = phase.homeOrNull()
    val presenter = HomePresenter(character, finishedSims, HomeTime(Instant.now()), LocalResources.current)
    val offline = (phase as? HomePhase.Loaded)?.offlineSince?.let {
        stringResource(R.string.offline_strip, HomeTime(Instant.now()).stamp(it))
    }
    val openItem = { look: ItemLook, sourceLine: String?, wornId: Int? ->
        home?.let { tooltip = HomePresenter.tooltipRequest(look, it, character, sourceLine, wornId) }
    }

    if (showGear && home != null) {
        BackHandler { showGear = false }
        GearPage(home, character, dependencies.catalog) { look, wornId -> openItem(look, null, wornId) }
    } else {
        PullToRefreshBox(
            isRefreshing = refreshing,
            onRefresh = {
                scope.launch {
                    refreshing = true
                    store.refresh(character)
                    refreshing = false
                }
            },
            modifier = Modifier.fillMaxSize().background(FsColor.bg),
        ) {
            Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
                HomeAnswer(phase, presenter, home?.dataBuild.orEmpty(), offline) { intent ->
                    when (intent) {
                        is HomeIntent.OpenItem -> openItem(intent.look, intent.sourceLine, intent.wornId)
                        HomeIntent.OpenGear -> showGear = true
                        HomeIntent.Retry -> scope.launch { store.refresh(character) }
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

/** Your gear, with the build's items file loaded for the slots the answer does not name. */
@Composable
private fun GearPage(home: CharacterHome, character: HomeCharacter, catalog: ItemCatalog, open: (ItemLook, Int?) -> Unit) {
    val res = LocalResources.current
    val items by produceState<Map<Int, Item>>(emptyMap(), home) {
        val probe = HomePresenter.tooltipRequest(ItemLook(0), home, character)
        value = catalog.sources(probe)?.items.orEmpty()
    }
    Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
        GearList(
            rows = GearModel.rows(home, character.gear, items, res),
            intro = stringResource(R.string.gear_intro, HomePresenter.bandHead(home, res).orEmpty()),
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
