package gg.foreversixty.app.guides

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.guides.logic.GuideLinks
import gg.foreversixty.app.guides.logic.GuideStore
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesScreen
import gg.foreversixty.app.guides.logic.GuidesSegment
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.logic.Tiers
import gg.foreversixty.app.guides.logic.valueOrNull
import gg.foreversixty.app.guides.ui.ChangeReader
import gg.foreversixty.app.guides.ui.ClassPage
import gg.foreversixty.app.guides.ui.GuideReader
import gg.foreversixty.app.guides.ui.GuideReaderActions
import gg.foreversixty.app.guides.ui.GuideReaderState
import gg.foreversixty.app.guides.ui.GuidesHome
import gg.foreversixty.app.guides.ui.GuidesHomeActions
import gg.foreversixty.app.guides.ui.GuidesHomeState
import gg.foreversixty.app.guides.ui.NewDataState
import gg.foreversixty.app.guides.ui.TierListActions
import gg.foreversixty.app.guides.ui.TierListContent
import gg.foreversixty.app.guides.ui.TierListState
import gg.foreversixty.app.guides.ui.WhatChangedActions
import gg.foreversixty.app.guides.ui.WhatChangedContent
import gg.foreversixty.app.guides.ui.WhatChangedState
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.planner.PlannerFeature
import gg.foreversixty.app.system.logic.ShareText
import gg.foreversixty.app.system.logic.TierRole
import kotlinx.coroutines.launch

/**
 * Guides, Tiers or Patch notes (the menu's tools, spec 2026-11-08): the menu row's [page], or the screen
 * on top of its stack. [character] is the session's current character (null signed out); [openBis] opens
 * a spec's BiS list under Gear; [signIn] starts the sign-in the tier list offers signed out; [selector]
 * fills the bar above the page (the menu's Back).
 */
@Composable
fun GuidesRoute(
    feature: GuidesFeature,
    page: GuidesSegment,
    character: AccountCharacter?,
    openBis: (classSlug: String, spec: String) -> Unit,
    planner: PlannerFeature? = null,
    signIn: (() -> Unit)? = null,
    selector: @Composable () -> Unit = {},
) {
    val store = feature.store
    LaunchedEffect(feature) { if (store.build == null) store.start() }
    val navigator = feature.navigator
    BackHandler(enabled = navigator.top != null) { navigator.pop() }
    when (val top = navigator.top) {
        is GuidesScreen.ClassPage -> ClassPage(top.classSlug, store.index, store.saved, feature.now(), navigator::pop) { navigator.push(GuidesScreen.Guide(it)) }
        is GuidesScreen.Guide -> GuideRoute(feature, top.spec, openBis, planner)
        is GuidesScreen.Change -> ChangeReader(store.changelog.valueOrNull?.entries?.firstOrNull { it.id == top.id }, navigator::pop)
        null -> GuidesRootLayout(selector) { GuidesRoot(feature, page.shown(feature.showsTiers), character, GuideTools(openBis, planner, signIn)) }
    }
}

/** The root's ways out to the rest of the app. */
private class GuideTools(val openBis: (String, String) -> Unit, val planner: PlannerFeature?, val signIn: (() -> Unit)?)

@Composable
private fun GuidesRoot(feature: GuidesFeature, page: GuidesSegment, character: AccountCharacter?, tools: GuideTools) {
    val store = feature.store
    LaunchedEffect(page, store.build) {
        when (page) {
            GuidesSegment.Tiers -> if (store.tiers !is Loadable.Loaded) store.loadTiers()
            GuidesSegment.Changes -> if (store.changelog !is Loadable.Loaded) store.loadChanges()
            GuidesSegment.Guides -> Unit
        }
    }
    val scope = rememberCoroutineScope()
    val index = store.index.valueOrNull
    val yours = index?.let { Guides.yourGuide(it, character) }
    val open = { spec: String -> feature.navigator.push(GuidesScreen.Guide(spec)) }
    when (page) {
        GuidesSegment.Guides -> GuidesHome(
            GuidesHomeState(store.index, store.saved, character, yours),
            GuidesHomeActions(openGuide = open, openClass = { feature.navigator.push(GuidesScreen.ClassPage(it)) }, retry = { scope.launch { store.loadIndex() } }),
        )
        GuidesSegment.Tiers -> TiersRoot(feature, character, tools, open)
        GuidesSegment.Changes -> Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            WhatChangedContent(
                WhatChangedState(store.dates, store.changelog, feature.now().atZone(feature.zone).toLocalDate()),
                feature.now(),
                WhatChangedActions(open = { feature.navigator.push(GuidesScreen.Change(it)) }, retry = { scope.launch { store.loadChanges() } }),
            )
        }
    }
}

@Composable
private fun TiersRoot(feature: GuidesFeature, character: AccountCharacter?, tools: GuideTools, open: (String) -> Unit) {
    val store = feature.store
    val haptics = rememberHaptics()
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val res = LocalResources.current
    val index = store.index.valueOrNull
    val yours = index?.let { Guides.yourGuide(it, character) }
    var choice by remember(index, character) { mutableStateOf(index?.let { Tiers.defaultChoice(it, character) } ?: TierChoice(TierChoice.DPS, TierChoice.ALLIANCE)) }
    // A row's Planner opens the guide's raid build, read from the guide file (saved, or fetched once).
    val openPlanner = tools.planner?.let { planner ->
        { spec: String ->
            scope.launch {
                val build = store.build ?: return@launch
                val guide = GuideStore(spec, feature.files).also { it.load(build) }.phase.valueOrNull ?: return@launch
                Guides.plannerCode(guide.talents)?.let { planner.openCode(guide.classSlug, it) }
            }
            Unit
        }
    }
    TierListContent(
        TierListState(store.tiers, choice, yours?.spec, index, signedIn = character != null),
        feature.now(),
        TierListActions(
            choose = {
                haptics.selection()
                choice = it
            },
            open = open,
            openBis = tools.openBis,
            openPlanner = openPlanner,
            openSim = { spec, faction -> AccountLinks.open(context, GuideLinks.rankingSim(spec, faction)) },
            openChecks = { AccountLinks.open(context, GuideLinks.SIM_CHECKS) },
            retry = { scope.launch { store.loadTiers() } },
            share = { picked -> ShareText.send(context, ShareText.tiers(res, TierRole.of(picked.role))) },
            signIn = tools.signIn,
        ),
    )
}

/** The root: the bar slot (the menu's Back), then the page, which scrolls itself. */
@Composable
fun GuidesRootLayout(selector: @Composable () -> Unit = {}, content: @Composable () -> Unit) {
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        selector()
        content()
    }
}

/** F9.2 over the tab: the saved copy at once, then the network; a new build waits behind the strip. */
@Composable
private fun GuideRoute(feature: GuidesFeature, spec: String, openBis: (String, String) -> Unit, planner: PlannerFeature?) {
    val store = feature.store
    val context = LocalContext.current
    val haptics = rememberHaptics()
    val scope = rememberCoroutineScope()
    // The saved copy is drawn before the first frame: a guide opens from cache with no spinner (spec §11).
    val guide = remember(spec) { GuideStore(spec, feature.files).also { saved -> store.build?.let(saved::drawSaved) } }
    val build = store.build
    LaunchedEffect(guide, build == null) { build?.let { guide.load(it) } }
    val indexed = store.index.valueOrNull?.let { index -> Guides.spec(index, spec)?.let { it to Guides.classOf(index, spec) } }
    val title = indexed?.let { (s, cls) -> "${s.specName} ${cls?.className.orEmpty()}" }.orEmpty()
    GuideReader(
        GuideReaderState(
            phase = guide.phase,
            title = title,
            classSlug = indexed?.second?.classSlug,
            newData = if (store.hasNewData) NewDataState(store.updateFailed) else null,
        ),
        feature.now(),
        GuideReaderActions(
            back = feature.navigator::pop,
            share = { message -> ShareText.send(context, message) },
            retry = { scope.launch { store.build?.let { guide.load(it) } } },
            update = {
                scope.launch {
                    if (store.applyUpdate()) store.build?.let { guide.reload(it) }
                }
            },
            selected = haptics::selection,
            openBis = { file -> openBis(file.classSlug, file.gear?.spec ?: file.spec) },
            openPlanner = planner?.let { feature -> { guide -> Guides.plannerCode(guide.talents)?.let { feature.openCode(guide.classSlug, it) } } },
        ),
    )
}
