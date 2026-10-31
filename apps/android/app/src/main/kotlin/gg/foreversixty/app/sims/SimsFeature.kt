package gg.foreversixty.app.sims

import android.content.Context
import android.content.Intent
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.sims.logic.SimsCopy
import gg.foreversixty.app.sims.ui.SimRowFace
import androidx.core.net.toUri
import androidx.lifecycle.compose.LifecycleResumeEffect
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.sims.logic.LiveSimsService
import gg.foreversixty.app.sims.logic.ServerRuns
import gg.foreversixty.app.sims.logic.SimResultPhase
import gg.foreversixty.app.sims.logic.SimResultStore
import gg.foreversixty.app.sims.logic.SimsService
import gg.foreversixty.app.sims.logic.SimsStore
import gg.foreversixty.app.sims.ui.SimResultContent
import gg.foreversixty.app.sims.ui.SimsActions
import gg.foreversixty.app.sims.ui.SimsList
import gg.foreversixty.app.sims.ui.SimsListState
import java.io.File
import java.time.Instant
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/** Where a sim link lands on Home's stack (spec §6.4): Sims, then the result over it. */
class SimsNavigator {
    var showsList by mutableStateOf(false)

    /** A result over Sims: its id and, from the list, its row. */
    var result by mutableStateOf<SimResultRoute?>(null)

    /** Set when a link asks for Home; the scaffold selects it and clears this. */
    var wantsHome by mutableStateOf(false)

    fun open(link: DeepLink): Boolean {
        if (link !is DeepLink.Sim) return false
        showsList = true
        result = SimResultRoute(link.id, null)
        wantsHome = true
        return true
    }
}

data class SimResultRoute(val id: String, val row: gg.foreversixty.app.model.SimRow?)

/** F5's dependencies, built once in the app graph. */
class SimsFeature(
    context: Context,
    apiClient: ApiClient,
    private val statics: StaticDataClient,
    val catalog: ItemCatalog,
    private val isAvailable: () -> Boolean,
    /** Lane F8's rule: See plans only where the storefront may link to web checkout (fail closed). */
    val mayLinkToPlans: suspend () -> Boolean = { false },
) {
    val service: SimsService = LiveSimsService(apiClient)

    /** Beside the raids cache; cleared with it at sign-out. */
    val cache: RaidCache = FileRaidCache(File(context.filesDir, "sims"))
    val specs: SpecCatalog = SpecCatalog.decode(readAsset(context, SpecCatalog.ASSET).orEmpty())
    val navigator = SimsNavigator()
    val now: () -> Instant = Instant::now

    /** Lane F14: a server run's poll answer and its `{Spec} · {kind}` title, for the ongoing notification. */
    var liveProgress: (suspend (SimRow, SimProgress, String) -> Unit)? = null

    /** Sims ships dark behind `sims` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun iconBuild(): String? = statics.activeBuild().dataOrNull()?.build

    companion object {
        private fun readAsset(context: Context, path: String): String? =
            runCatching { context.assets.open(path).bufferedReader().use { it.readText() } }.getOrNull()
    }
}

/** F5.1 pushed on Home, and F5.3 over it; [entitled] is `entitlements.server_sims`. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun SimsRoute(feature: SimsFeature, entitled: Boolean, onBack: () -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val res = LocalResources.current
    val store = remember(feature) {
        SimsStore(feature.service, feature.cache, feature.now) { row, progress ->
            val face = SimRowFace(row, feature.specs, feature.now(), SimsCopy(res))
            feature.liveProgress?.invoke(row, progress, "${face.specName} · ${face.kind}")
        }
    }
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    val plans by produceState(false, feature) { value = feature.mayLinkToPlans() }
    var refreshing by remember { mutableStateOf(false) }
    var active by remember { mutableStateOf(true) }
    LaunchedEffect(store) { store.load() }
    LifecycleResumeEffect(store) {
        active = true
        onPauseOrDispose { active = false }
    }
    // Every 5 s while the screen shows and the app is foreground, for runs still in flight.
    LaunchedEffect(active) {
        while (active) {
            delay(ServerRuns.POLL_INTERVAL)
            if (store.isPolling) store.pollOnce()
        }
    }
    val route = feature.navigator.result
    if (route != null) {
        BackHandler { feature.navigator.result = null }
        SimResultRoute(feature, route) { id -> store.markSeen(id) }
        return
    }
    BackHandler(onBack = onBack)
    val actions = SimsActions(
        open = { row -> row.simId?.let { feature.navigator.result = SimResultRoute(it, row) } },
        runAgain = { id -> scope.launch { store.runAgain(id) } },
        retry = { scope.launch { store.load() } },
        loadMore = { scope.launch { store.loadMore() } },
        openPlans = { context.openInBrowser(gg.foreversixty.app.account.AccountLinks.PREMIUM) },
    )
    PullToRefreshBox(
        isRefreshing = refreshing,
        onRefresh = { scope.launch { refreshing = true; store.load(); refreshing = false } },
        modifier = Modifier.fillMaxSize().background(FsColor.bg),
    ) {
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            SimsList(
                store.phase,
                SimsListState(entitled, store.progress, store.rerun, store.pageFailed, plans),
                feature.specs, build, feature.now(), actions,
            )
        }
    }
}

@Composable
private fun SimResultRoute(feature: SimsFeature, route: SimResultRoute, markSeen: suspend (String) -> Unit) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val store = remember(route.id) {
        SimResultStore(route.id, route.row, feature.service, feature.cache, feature.specs, res, feature.now, markSeen)
    }
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(store) { store.load() }
    val model = (store.phase as? SimResultPhase.Loaded)?.model
    val looks by produceState(emptyMap(), model, build) {
        val b = build
        value = if (model != null && b != null) SimGearLooks.load(model, b, feature.catalog) else emptyMap()
    }
    Box(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
        SimResultContent(
            store.phase, looks, build, feature.now(),
            openItem = { cell ->
                val b = build
                if (model != null && b != null) tooltip = SimGearLooks.request(cell, looks[cell.itemId], model, b)
            },
            openBreakdown = { context.openInBrowser(it) },
            retry = { scope.launch { store.load() } },
        )
    }
    tooltip?.let { ItemTooltipSheet(it, feature.catalog, onDismiss = { tooltip = null }) }
}

private fun Context.openInBrowser(url: String) {
    startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
}
