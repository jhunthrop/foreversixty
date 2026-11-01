package gg.foreversixty.app.rankings

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableFloatStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.onGloballyPositioned
import androidx.compose.ui.layout.positionInRoot
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.platform.LocalView
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.bis.BisContext
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.ClassSpecSections
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.paperdoll.PaperdollSlot
import gg.foreversixty.app.raids.logic.ViewRaidHaptics
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.CharacterStore
import gg.foreversixty.app.rankings.logic.Cohort
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsStore
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.ui.CharacterActions
import gg.foreversixty.app.rankings.ui.CharacterContent
import gg.foreversixty.app.rankings.ui.CharacterScreenState
import gg.foreversixty.app.rankings.ui.FiltersActions
import gg.foreversixty.app.rankings.ui.FiltersSheet
import gg.foreversixty.app.rankings.ui.FiltersSheetState
import gg.foreversixty.app.rankings.ui.RankingsActions
import gg.foreversixty.app.rankings.ui.RankingsContent
import gg.foreversixty.app.rankings.ui.RankingsScreenState
import gg.foreversixty.app.reports.ReportRoute
import gg.foreversixty.app.reports.ReportRouteScreen
import gg.foreversixty.app.reports.ReportsFeature
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.reports.ui.MetersDragHandle
import kotlinx.coroutines.launch

/** F10's report screen, reachable from a row's menu and a recent fight while the `reports` switch is on. */
class RankingsReports(val feature: ReportsFeature, val viewer: ReportsViewer)

/** F11 on Home's stack: a character page over the board (or over Home, from a link), a report over either. */
@Composable
fun RankingsRoute(feature: RankingsFeature, viewer: RankingsViewer, reports: RankingsReports?, onBack: () -> Unit) {
    var report by remember { mutableStateOf<ReportRoute?>(null) }
    val nav = feature.navigator
    val openReport = { id: String -> report = ReportRoute(id) }
    val current = report
    val character = nav.character
    when {
        current != null && reports != null -> ReportRouteScreen(reports.feature, current, reports.viewer) { report = null }
        character != null -> CharacterRouteScreen(feature, character, viewer, reports, openReport) { nav.character = null }
        nav.showsBoard -> BoardRoute(feature, viewer, reports, openReport, onBack)
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun BoardRoute(feature: RankingsFeature, viewer: RankingsViewer, reports: RankingsReports?, openReport: (String) -> Unit, onBack: () -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val store = remember(feature, viewer) { feature.store(viewer) }
    val scroll = rememberScrollState()
    var refreshing by remember { mutableStateOf(false) }
    var filtersOpen by remember { mutableStateOf(false) }
    var choosingEncounter by remember { mutableStateOf(false) }
    var pickerOpen by remember { mutableStateOf(false) }
    var top by remember { mutableFloatStateOf(0f) }
    var ownAt by remember { mutableFloatStateOf(0f) }
    LaunchedEffect(store) {
        store.load(feature.navigator.encounter)
        feature.navigator.encounter = null
    }
    BackHandler(onBack = onBack)
    val refresh: () -> Unit = { scope.launch { refreshing = true; store.load(); refreshing = false } }
    val actions = RankingsActions(
        back = onBack,
        retry = refresh,
        openFilters = { filtersOpen = true },
        chooseSpec = { pickerOpen = true },
        open = { row -> feature.navigator.character = CharacterRoute.ofKey(row.player.key, row.player.name) },
        loadMore = { scope.launch { store.loadMore() } },
        jump = { scope.launch { scroll.animateScrollTo((scroll.value + ownAt - top).toInt().coerceAtLeast(0)) } },
        ownRowAt = { ownAt = it },
        watch = { row -> scope.launch { store.toggleWatch(row) } },
        openReport = { row -> openReport(row.reportId) },
        share = { row -> CharacterCard.shareUrl(row.player.key)?.let { RankingsFeature.share(context, it) } },
    )
    val state = RankingsScreenState(store.phase, store.filters, store.encounterNames, viewer, store.pageFailed, store.watched, reports != null)
    PullToRefreshBox(isRefreshing = refreshing, onRefresh = refresh, modifier = Modifier.fillMaxSize().background(FsColor.bg)) {
        Box(Modifier.fillMaxSize().onGloballyPositioned { top = it.positionInRoot().y }.verticalScroll(scroll)) {
            RankingsContent(state, feature.now(), actions, feature.zone)
        }
    }
    val filters = store.filters
    if (filtersOpen && filters != null) {
        FiltersSheet(
            FiltersSheetState(filters, store.encounters, Rankings.phaseChoices(store.phases, feature.now()), choosingEncounter),
            FiltersActions(
                apply = { next -> scope.launch { store.apply(next) } },
                reset = { scope.launch { store.reset() } },
                chooseEncounter = { choosingEncounter = it },
            ),
        ) {
            filtersOpen = false
            choosingEncounter = false
        }
    }
    if (pickerOpen && filters != null) SpecPicker(feature, store, filters.cohort) { pickerOpen = false }
}

/** Signed out: the F3.2 class and spec choices; the board follows each pick. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun SpecPicker(feature: RankingsFeature, store: RankingsStore, cohort: Cohort?, close: () -> Unit) {
    val scope = rememberCoroutineScope()
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    val specs = feature.specs
    val start = cohort?.let { c -> specs.spec(c.classSlug, c.specName) } ?: specs.specs.firstOrNull() ?: return
    val selection = BisSelection(start.classSlug, start.spec, BisRange.all.last(), BisSelection.factions.first())
    ModalBottomSheet(
        onDismissRequest = close,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = GameColors.sheet,
        dragHandle = { MetersDragHandle() },
    ) {
        Column(
            Modifier.verticalScroll(rememberScrollState()).navigationBarsPadding().padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(22.dp),
        ) {
            ClassSpecSections(BisContext(selection), specs, build.orEmpty()) { picked ->
                val spec = specs.spec(picked.spec) ?: return@ClassSpecSections
                val filters = store.filters ?: return@ClassSpecSections
                scope.launch { store.apply(filters.copy(cohort = Cohort(spec.classSlug, spec.name))) }
            }
        }
    }
}

/** F11.4: one character, its Watch star, its gear's tooltips and its recent fights' reports. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CharacterRouteScreen(
    feature: RankingsFeature,
    route: CharacterRoute,
    viewer: RankingsViewer,
    reports: RankingsReports?,
    openReport: (String) -> Unit,
    back: () -> Unit,
) {
    val context = LocalContext.current
    val view = LocalView.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val store = remember(route, viewer) { CharacterStore(route, feature.service, feature.cache, viewer, feature.now) }
    val haptics = remember(view) { ViewRaidHaptics(view) }
    var refreshing by remember { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(store) { store.load() }
    BackHandler(onBack = back)
    val profile = (store.phase as? CharacterPhase.Loaded)?.profile
    val build by produceState<String?>(null, feature) { value = feature.iconBuild() }
    val classSlug = profile?.let(CharacterCard::classSlug)
    val items by produceState(emptyMap(), build, classSlug) {
        val b = build
        if (b != null && classSlug != null) value = feature.catalog.items(b, classSlug).orEmpty()
    }
    val actions = CharacterActions(
        back = back,
        retry = { scope.launch { store.load() } },
        toggleWatch = { scope.launch { store.toggleWatch(haptics) } },
        share = { RankingsFeature.share(context, CharacterCard.shareUrl(route)) },
        openHow = { AccountLinks.open(context, CharacterCard.HOW_URL) },
        openSocket = { slot -> profile?.let { p -> build?.let { b -> tooltip = tooltipRequest(slot, p, b, feature, res) } } },
        openFight = { fight -> openReport(fight.reportId) },
        openHistory = { AccountLinks.open(context, CharacterCard.historyUrl(route)) },
        searchAgain = back,
    )
    val state = CharacterScreenState(route, store.phase, store.watch, viewer, items, build, reports != null)
    PullToRefreshBox(
        isRefreshing = refreshing,
        onRefresh = { scope.launch { refreshing = true; store.load(); refreshing = false } },
        modifier = Modifier.fillMaxSize().background(FsColor.bg),
    ) {
        Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            CharacterContent(state, feature.specs, feature.now(), actions, feature.zone)
        }
    }
    tooltip?.let { request -> ItemTooltipSheet(request, feature.catalog, onDismiss = { tooltip = null }) }
}

/** The F2.6 sheet for a socket: the worn item in the character's class items, spec and level band. */
private fun tooltipRequest(slot: PaperdollSlot, profile: CharacterProfile, build: String, feature: RankingsFeature, res: android.content.res.Resources): ItemTooltipRequest? {
    val item = slot.item ?: return null
    val range = profile.character.level?.let(BisRange::forLevel)
    return ItemTooltipRequest(
        itemId = item.id, fallback = item, build = build, classSlug = CharacterCard.classSlug(profile).orEmpty(),
        spec = profile.gearSpec, specName = CharacterCard.specName(profile, feature.specs), band = range?.band,
        bandLabel = range?.let { BisLabels.head(res, it) }, preset = range?.preset, faction = null, wornId = item.id,
    )
}
