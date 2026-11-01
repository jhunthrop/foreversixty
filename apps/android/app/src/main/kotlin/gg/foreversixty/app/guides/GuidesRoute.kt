package gg.foreversixty.app.guides

import gg.foreversixty.app.system.logic.TierRole
import gg.foreversixty.app.system.logic.ShareText
import gg.foreversixty.app.common.rememberHaptics
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
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
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
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
import gg.foreversixty.app.guides.ui.EqualSegments
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
import gg.foreversixty.app.guides.ui.screenMargins
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.sims.ui.ScreenTitle
import kotlinx.coroutines.launch

/**
 * The Guides tab (spec F9): the root with its segments, or the screen on top of its stack.
 * [character] is the session's current character (null signed out); [openBis] switches to the BiS
 * tab on a guide's list; [selector] fills the root bar's leading slot, as on Home, Raids and BiS.
 */
@Composable
fun GuidesRoute(feature: GuidesFeature, character: AccountCharacter?, openBis: (GuideFile) -> Unit, selector: @Composable () -> Unit = {}) {
    val store = feature.store
    LaunchedEffect(feature) { if (store.build == null) store.start() }
    val navigator = feature.navigator
    BackHandler(enabled = navigator.top != null) { navigator.pop() }
    when (val top = navigator.top) {
        is GuidesScreen.ClassPage -> ClassPage(top.classSlug, store.index, store.saved, feature.now(), navigator::pop) { navigator.push(GuidesScreen.Guide(it)) }
        is GuidesScreen.Guide -> GuideRoute(feature, top.spec, openBis)
        is GuidesScreen.Change -> ChangeReader(store.changelog.valueOrNull?.entries?.firstOrNull { it.id == top.id }, navigator::pop)
        null -> GuidesRoot(feature, character, selector)
    }
}

@Composable
private fun GuidesRoot(feature: GuidesFeature, character: AccountCharacter?, selector: @Composable () -> Unit) {
    val store = feature.store
    val haptics = rememberHaptics()
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val res = LocalResources.current
    val segments = GuidesSegment.visible(feature.showsTiers)
    var segment by remember(segments) { mutableStateOf(GuidesSegment.load(feature.segment, feature.showsTiers)) }
    LaunchedEffect(segment, store.build) {
        when (segment) {
            GuidesSegment.Tiers -> if (store.tiers !is Loadable.Loaded) store.loadTiers()
            GuidesSegment.Changes -> if (store.changelog !is Loadable.Loaded) store.loadChanges()
            GuidesSegment.Guides -> Unit
        }
    }
    val index = store.index.valueOrNull
    val yours = index?.let { Guides.yourGuide(it, character) }
    var choice by remember(index, character) { mutableStateOf(index?.let { Tiers.defaultChoice(it, character) } ?: TierChoice(TierChoice.DPS, TierChoice.HORDE)) }
    val open = { spec: String -> feature.navigator.push(GuidesScreen.Guide(spec)) }
    GuidesRootLayout(
        segments = segments,
        segment = segment,
        onSegment = {
            haptics.selection()
            segment = it
            GuidesSegment.save(feature.segment, it)
        },
        selector = selector,
    ) {
        when (segment) {
            GuidesSegment.Guides -> GuidesHome(
                GuidesHomeState(store.index, store.saved, yours),
                GuidesHomeActions(openGuide = open, openClass = { feature.navigator.push(GuidesScreen.ClassPage(it)) }, retry = { scope.launch { store.loadIndex() } }),
            )
            GuidesSegment.Tiers -> TierListContent(
                TierListState(store.tiers, choice, yours?.spec),
                feature.now(),
                TierListActions(choose = {
                    haptics.selection()
                    choice = it
                }, open = open, retry = { scope.launch { store.loadTiers() } }, share = { picked ->
                    ShareText.send(context, ShareText.tiers(res, TierRole.of(picked.role)))
                }),
            )
            GuidesSegment.Changes -> WhatChangedContent(
                WhatChangedState(store.dates, store.changelog, feature.now().atZone(feature.zone).toLocalDate()),
                feature.now(),
                WhatChangedActions(open = { feature.navigator.push(GuidesScreen.Change(it)) }, retry = { scope.launch { store.loadChanges() } }),
            )
        }
    }
}

/** The root: the selector row, `GUIDES`, the segments, then the segment's content, scrolling as one. */
@Composable
fun GuidesRootLayout(
    segments: List<GuidesSegment>,
    segment: GuidesSegment,
    onSegment: (GuidesSegment) -> Unit,
    selector: @Composable () -> Unit = {},
    content: @Composable () -> Unit,
) {
    val res = LocalResources.current
    Column(Modifier.fillMaxSize().background(FsColor.bg).verticalScroll(rememberScrollState())) {
        selector()
        ScreenTitle(stringResource(R.string.guides_title))
        Column(Modifier.screenMargins().padding(vertical = SEGMENT_GAP)) {
            EqualSegments(segments.map { res.getString(it.label) }, segments.indexOf(segment)) { onSegment(segments[it]) }
        }
        content()
    }
}

/** F9.2 over the tab: the saved copy at once, then the network; a new build waits behind the strip. */
@Composable
private fun GuideRoute(feature: GuidesFeature, spec: String, openBis: (GuideFile) -> Unit) {
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
            openBis = openBis,
        ),
    )
}

private val GuidesSegment.label: Int
    get() = when (this) {
        GuidesSegment.Guides -> R.string.guides_segment_guides
        GuidesSegment.Tiers -> R.string.guides_segment_tiers
        GuidesSegment.Changes -> R.string.guides_segment_changes
    }

private val SEGMENT_GAP = 8.dp
