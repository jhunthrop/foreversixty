package gg.foreversixty.app.reports

import gg.foreversixty.app.common.Connectivity
import android.content.Context
import android.view.View
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
import androidx.compose.ui.platform.LocalView
import androidx.lifecycle.compose.LifecycleResumeEffect
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.model.ReportRow
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidLinks
import gg.foreversixty.app.raids.logic.ViewRaidHaptics
import gg.foreversixty.app.reports.logic.LiveReportsService
import gg.foreversixty.app.reports.logic.MetersStore
import gg.foreversixty.app.reports.logic.ReportStore
import gg.foreversixty.app.reports.logic.Reports
import gg.foreversixty.app.reports.logic.ReportsService
import gg.foreversixty.app.reports.logic.ReportsStore
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.reports.ui.MetersActions
import gg.foreversixty.app.reports.ui.MetersSheet
import gg.foreversixty.app.reports.ui.MetersSheetState
import gg.foreversixty.app.reports.ui.ReportActions
import gg.foreversixty.app.reports.ui.ReportContent
import gg.foreversixty.app.reports.ui.ReportScreenState
import gg.foreversixty.app.reports.ui.ReportsList
import gg.foreversixty.app.reports.ui.ReportsListActions
import java.io.File
import java.time.Instant
import java.time.ZoneId
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/** A report over Your reports: its id and, from the list, the zone the bar shows while it loads. */
data class ReportRoute(val id: String, val title: String = "")

/** Where a report link lands on Home's stack (spec §3.4): Your reports, then the report over it. */
class ReportsNavigator {
    var showsList by mutableStateOf(false)
    var report by mutableStateOf<ReportRoute?>(null)

    /** Set when a link asks for Home; the scaffold selects it and clears this. */
    var wantsHome by mutableStateOf(false)

    fun open(link: DeepLink): Boolean {
        if (link !is DeepLink.Report) return false
        showsList = true
        report = ReportRoute(link.id)
        wantsHome = true
        return true
    }
}

/** F10's dependencies, built once in the app graph ([create]); tests pass a fake service and cache. */
class ReportsFeature(
    val service: ReportsService,
    /** Beside the raids cache; cleared with it at sign-out. */
    val cache: RaidCache,
    val specs: SpecCatalog,
    private val build: suspend () -> String?,
    private val isAvailable: () -> Boolean,
    val now: () -> Instant = Instant::now,
    /** The zone dates and the offline stamp read in. */
    val zone: ZoneId = ZoneId.systemDefault(),
    /** Lane F15: whether a failed refresh is offline or stale. */
    val connectivity: Connectivity = Connectivity.Unknown,
) {
    val navigator = ReportsNavigator()

    /** Reports ships dark behind `reports` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun iconBuild(): String? = build()

    /** Home's Logs row: the last raid with a boss pull, or null with none. */
    suspend fun lastRaid(): ReportRow? = ReportsStore(service, cache, now).lastRaid()

    companion object {
        fun create(
            context: Context,
            apiClient: ApiClient,
            statics: StaticDataClient,
            specs: SpecCatalog,
            isAvailable: () -> Boolean,
            connectivity: Connectivity = Connectivity.Unknown,
        ) = ReportsFeature(
            service = LiveReportsService(apiClient),
            cache = FileRaidCache(File(context.filesDir, "reports")),
            specs = specs,
            build = { statics.activeBuild().dataOrNull()?.build },
            isAvailable = isAvailable,
            connectivity = connectivity,
        )
    }
}

/** F10.1 pushed on Home, and F10.2 over it. [signIn] is the Battle.net sign-in for the signed-out state. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ReportsRoute(feature: ReportsFeature, viewer: ReportsViewer, signIn: () -> Unit, onBack: () -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val store = remember(feature) { ReportsStore(feature.service, feature.cache, feature.now, connectivity = feature.connectivity) }
    var refreshing by remember { mutableStateOf(false) }
    LaunchedEffect(store) { store.load() }
    val route = feature.navigator.report
    if (route != null) {
        ReportRouteScreen(feature, route, viewer) { feature.navigator.report = null }
        return
    }
    BackHandler(onBack = onBack)
    val refresh: () -> Unit = { scope.launch { refreshing = true; store.load(); refreshing = false } }
    val actions = ReportsListActions(
        back = onBack,
        refresh = refresh,
        open = { row -> feature.navigator.report = ReportRoute(row.id, row.zone) },
        loadMore = { scope.launch { store.loadMore() } },
        getCompanion = { AccountLinks.open(context, AccountLinks.SETUP) },
        signIn = signIn,
    )
    PullToRefreshBox(isRefreshing = refreshing, onRefresh = refresh, modifier = Modifier.fillMaxSize().background(FsColor.bg)) {
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            ReportsList(store.phase, store.pageFailed, viewer, feature.now(), actions, feature.zone)
        }
    }
}

/** F10.2: one report, its live poll, and the fight sheet over it. Raids' night pushes it too. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ReportRouteScreen(feature: ReportsFeature, route: ReportRoute, viewer: ReportsViewer, back: () -> Unit) {
    val context = LocalContext.current
    val view = LocalView.current
    val scope = rememberCoroutineScope()
    val store = remember(route.id, viewer) { ReportStore(route.id, feature.service, feature.cache, viewer, feature.connectivity, feature.now) }
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    var refreshing by remember { mutableStateOf(false) }
    var active by remember { mutableStateOf(true) }
    var meters by remember { mutableStateOf<MetersStore?>(null) }
    LaunchedEffect(store) { store.load() }
    LifecycleResumeEffect(store) {
        active = true
        onPauseOrDispose { active = false }
    }
    // A live report re-reads itself every 20 s while this screen is open and foreground.
    LaunchedEffect(active, store.isLive) {
        while (active && store.isLive) {
            delay(Reports.LIVE_POLL)
            store.loadReport()
        }
    }
    BackHandler(onBack = back)
    val actions = ReportActions(
        back = back,
        openWeb = { AccountLinks.open(context, RaidLinks.report(route.id)) },
        retry = { scope.launch { store.load() } },
        retryParses = { scope.launch { store.loadParses() } },
        retryCoach = { scope.launch { store.loadParses() } },
        openMeters = { target -> meters = MetersStore(route.id, target, feature.service).also { scope.launch { it.load() } } },
    )
    PullToRefreshBox(
        isRefreshing = refreshing,
        onRefresh = { scope.launch { refreshing = true; store.load(); refreshing = false } },
        modifier = Modifier.fillMaxSize().background(FsColor.bg),
    ) {
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            ReportContent(ReportScreenState(store.phase, store.parses, store.coach, route.title), feature.now(), actions, feature.zone)
        }
    }
    meters?.let { sheet -> FightSheet(sheet, feature, build, view) { meters = null } }
}

@Composable
private fun FightSheet(store: MetersStore, feature: ReportsFeature, build: String?, view: View, close: () -> Unit) {
    val scope = rememberCoroutineScope()
    val haptics = remember(view) { ViewRaidHaptics(view) }
    MetersSheet(
        MetersSheetState(store.target, store.metric, store.phase, store.showAll),
        feature.specs,
        build,
        MetersActions(
            select = { metric ->
                haptics.selection()
                scope.launch { store.select(metric) }
            },
            showAll = { store.showAll = true },
            retry = { scope.launch { store.load() } },
        ),
        onDismiss = close,
    )
}
