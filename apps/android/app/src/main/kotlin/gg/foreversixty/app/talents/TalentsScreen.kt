package gg.foreversixty.app.talents

import android.content.Intent
import androidx.compose.foundation.ExperimentalFoundationApi
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxWithConstraints
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.selection.selectableGroup
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.text.rememberTextMeasurer
import androidx.compose.material.icons.filled.Check
import androidx.compose.ui.hapticfeedback.HapticFeedbackType
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalHapticFeedback
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.stateDescription
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.core.net.toUri
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.home.HomeLinks
import java.time.Instant
import java.util.Locale
import kotlinx.coroutines.launch

/** From font scale 1.5 the segments, tabs and status stack and the bar stops pinning (spec F6 §10). */
private const val LARGE_TEXT_SCALE = 1.5f

/**
 * The Talents screen pushed from Home's Talents row (spec F6 §3): the model, the chosen build
 * and tree, and the rank sheet. [model] is Home's, already loading for the Talents row. [back] returns to Home.
 */
@Composable
fun TalentsRoute(character: HomeCharacter, model: TalentsModel, back: () -> Unit) {
    val phase by model.phase.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val res = LocalResources.current
    var shown by rememberSaveable { mutableStateOf(TalentBuildKind.Yours) }
    var tree by rememberSaveable { mutableStateOf(0) }
    var settled by rememberSaveable { mutableStateOf(false) }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    LaunchedEffect(model) {
        if (model.phase.value !is TalentsPhase.Loaded) model.load()
        val loaded = model.phase.value as? TalentsPhase.Loaded
        if (!settled && loaded != null) {
            settled = true
            shown = loaded.presenter.initialBuild
            tree = loaded.presenter.openingTree
        }
    }
    TalentsPage(
        phase = phase,
        build = model.build,
        offlineStamp = (phase as? TalentsPhase.Loaded)?.offlineSinceMillis?.let {
            OfflineStamp.format(Instant.ofEpochMilli(it), Instant.now())
        },
        shown = shown,
        tree = tree,
        onShow = { shown = it },
        onTree = { tree = it },
        selectedId = sheetId,
        openTalent = { sheetId = it },
        retry = { scope.launch { model.load() } },
        openPlanner = {
            val url = TalentLinks.planner(character.className, character.race)
            context.startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
        },
        back = back,
    )
    val sheet = (phase as? TalentsPhase.Loaded)?.presenter?.let { presenter -> sheetId?.let { presenter.sheet(shown, it, res) } }
    sheet?.let { TalentRankSheet(it, model.build) { sheetId = null } }
}

/** The planner on the site, for the character's class and race (spec F6 §5.1 item 8). */
object TalentLinks {
    fun planner(classSlug: String?, race: String?): String {
        val query = listOfNotNull(classSlug?.let { "class=$it" }, race?.let { "race=${raceSlug(it)}" })
        return if (query.isEmpty()) HomeLinks.PLANNER else "${HomeLinks.PLANNER}?${query.joinToString("&")}"
    }

    /** "Night Elf" to "night-elf", the planner's race slug (`account-character.ts` `raceSlugOf`). */
    fun raceSlug(race: String): String = race.trim().lowercase(Locale.ROOT).split(Regex("\\s+")).joinToString("-")
}

/** The bar and the content, stateless so the screenshot tests draw every state. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TalentsPage(
    phase: TalentsPhase,
    build: String,
    offlineStamp: String?,
    shown: TalentBuildKind,
    tree: Int,
    onShow: (TalentBuildKind) -> Unit = {},
    onTree: (Int) -> Unit = {},
    selectedId: Int? = null,
    openTalent: (Int) -> Unit = {},
    retry: () -> Unit = {},
    openPlanner: () -> Unit = {},
    back: () -> Unit = {},
) {
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        TopAppBar(
            title = { Text(stringResource(R.string.talents_title).uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong) },
            navigationIcon = {
                IconButton(onClick = back) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = null, tint = FsColor.strong) }
            },
            colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
            // The app scaffold already pads for the status bar.
            windowInsets = WindowInsets(0, 0, 0, 0),
        )
        TalentsContent(phase, build, offlineStamp, shown, tree, onShow, onTree, selectedId, openTalent, retry, openPlanner)
    }
}

@OptIn(ExperimentalFoundationApi::class)
@Composable
private fun TalentsContent(
    phase: TalentsPhase,
    build: String,
    offlineStamp: String?,
    shown: TalentBuildKind,
    tree: Int,
    onShow: (TalentBuildKind) -> Unit,
    onTree: (Int) -> Unit,
    selectedId: Int?,
    openTalent: (Int) -> Unit,
    retry: () -> Unit,
    openPlanner: () -> Unit,
) {
    val res = LocalResources.current
    val large = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE
    val list = rememberLazyListState()
    LaunchedEffect(selectedId) {
        // The tapped cell stays above the sheet (spec F6 §7, point 4): bring the tree to the top.
        if (selectedId != null && phase is TalentsPhase.Loaded) list.animateScrollToItem(TREE_ITEM_INDEX, scrollOffset = -SHEET_LIFT_PX)
    }
    LazyColumn(Modifier.fillMaxSize(), state = list) {
        item { offlineStamp?.let { OfflineStrip(stringResource(R.string.offline_strip, it)) } }
        when (phase) {
            TalentsPhase.Loading -> item { TalentsSkeleton() }
            TalentsPhase.Failed -> item {
                Column(Modifier.padding(start = 16.dp, end = 16.dp, top = 8.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
                    Text(stringResource(R.string.talents_failed), style = FsType.body, color = FsColor.wipe, modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite })
                    BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, retry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
                }
            }
            is TalentsPhase.Loaded -> {
                val presenter = phase.presenter
                item { TalentStatusBlock(presenter.status(res), presenter.notes(res), retry, large) }
                if (presenter.hasTree) {
                    val bar: @Composable () -> Unit = {
                        TalentCompareBar(presenter.showsSegments, presenter.tabs(shown), shown, tree, onShow, onTree, large)
                    }
                    if (large) item { bar() } else stickyHeader { bar() }
                    item {
                        TalentTreeView(
                            presenter.cells(shown, tree),
                            presenter.connectors(shown, tree),
                            build,
                            selectedId,
                            openTalent,
                            Modifier.padding(top = 16.dp),
                        )
                    }
                    item {
                        Text(presenter.legend(shown, res), style = FsType.meta, color = FsColor.muted, modifier = Modifier.padding(start = 16.dp, end = 16.dp, top = 8.dp))
                    }
                }
            }
        }
        item { TalentsFooter(openPlanner) }
    }
}

/** The tree's item in the list: after the strip, the status and the bar. */
private const val TREE_ITEM_INDEX = 3
private const val SHEET_LIFT_PX = 0

/** The status block, the screen's one primary element, then the secondary notes (spec F6 §5.1). */
@Composable
fun TalentStatusBlock(status: TalentStatusLine, notes: List<String>, retry: () -> Unit, large: Boolean) {
    Column(Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, top = 8.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
        val tagColor = when (status.tagKind) {
            TalentTagKind.Matches -> FsColor.kill
            TalentTagKind.Differs -> FsColor.gold
            TalentTagKind.Muted -> FsColor.muted
        }
        val sentence = @Composable { modifier: Modifier ->
            Text(status.sentence, style = FsType.body, color = if (status.isError) FsColor.wipe else FsColor.strong, modifier = modifier)
        }
        val described = Modifier.heightIn(min = 28.dp).semantics(mergeDescendants = true) {
            contentDescription = status.accessibilityLabel
            liveRegion = LiveRegionMode.Polite
        }
        if (large) {
            Column(described, verticalArrangement = Arrangement.spacedBy(6.dp)) {
                status.tag?.let { Tag(it, tagColor) }
                sentence(Modifier)
            }
        } else {
            Row(described, horizontalArrangement = Arrangement.spacedBy(10.dp), verticalAlignment = Alignment.CenterVertically) {
                status.tag?.let { Tag(it, tagColor) }
                sentence(Modifier.weight(1f))
            }
        }
        if (status.isError) {
            BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, retry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
        }
        notes.forEach { Text(it, style = FsType.meta, color = FsColor.muted) }
    }
}

/** The two segments and the three tree tabs, pinned under the bar below 1.5x (spec F6 §5.1 item 5). */
@Composable
fun TalentCompareBar(
    showsSegments: Boolean,
    tabs: List<TalentTabModel>,
    shown: TalentBuildKind,
    tree: Int,
    onShow: (TalentBuildKind) -> Unit,
    onTree: (Int) -> Unit,
    large: Boolean,
) {
    val haptics = LocalHapticFeedback.current
    Column(
        Modifier.fillMaxWidth().background(FsColor.bg).padding(start = 16.dp, end = 16.dp, top = 12.dp, bottom = 8.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        if (showsSegments) {
            Segments(shown, large) {
                haptics.performHapticFeedback(HapticFeedbackType.SegmentTick)
                onShow(it)
            }
        }
        val pick = { index: Int ->
            haptics.performHapticFeedback(HapticFeedbackType.SegmentTick)
            onTree(index)
        }
        if (large) StackedTabs(tabs, tree, pick, rowHeight = 56.dp) else FittingTabs(tabs, tree, pick)
    }
}

@Composable
private fun Segments(shown: TalentBuildKind, large: Boolean, onShow: (TalentBuildKind) -> Unit) {
    val shape = RoundedCornerShape(if (large) FsRadius.control else SEGMENT_PILL)
    val items = listOf(TalentBuildKind.Guide to R.string.talents_guide, TalentBuildKind.Yours to R.string.talents_yours)
    val content = @Composable { kind: TalentBuildKind, label: Int, modifier: Modifier ->
        val checked = shown == kind
        Row(
            modifier
                .heightIn(min = 48.dp)
                .background(if (checked) FsColor.gold.copy(alpha = TILE_ALPHA) else FsColor.raised)
                .selectable(checked, role = Role.RadioButton) { onShow(kind) },
            horizontalArrangement = Arrangement.Center,
            verticalAlignment = Alignment.CenterVertically,
        ) {
            if (checked) {
                Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.strong, modifier = Modifier.size(16.dp))
                Spacer(Modifier.width(6.dp))
            }
            Text(stringResource(label), style = FsType.segment, color = if (checked) FsColor.strong else FsColor.muted)
        }
    }
    val frame = Modifier.fillMaxWidth().border(1.dp, FsColor.line, shape).clip(shape).selectableGroup()
    if (large) {
        Column(frame) { items.forEach { (kind, label) -> content(kind, label, Modifier.fillMaxWidth()) } }
    } else {
        Row(frame) { items.forEach { (kind, label) -> content(kind, label, Modifier.weight(1f)) } }
    }
}

/**
 * Three tabs in a row while every label fits on one line, else the stacked list (spec F6 §10).
 * Each tab in the row takes width in proportion to its label, as the iOS row does.
 */
@Composable
private fun FittingTabs(tabs: List<TalentTabModel>, tree: Int, onTree: (Int) -> Unit) {
    val measurer = rememberTextMeasurer()
    val density = LocalDensity.current
    BoxWithConstraints(Modifier.fillMaxWidth()) {
        val widths = tabs.map { tab ->
            val name = measurer.measure(tab.name, if (tab.index == tree) FsType.bodyStrong else FsType.body).size.width
            val points = measurer.measure(tab.points.toString(), FsType.mono).size.width
            val dot = if (tab.differing > 0) with(density) { (DOT_GAP + DOT_SIZE).roundToPx() } else 0
            name + points + dot + with(density) { (LABEL_GAP + TAB_PADDING * 2).roundToPx() }
        }
        val gutters = with(density) { (TAB_GUTTER * (tabs.size - 1)).roundToPx() }
        if (widths.sum() + gutters <= constraints.maxWidth) RowTabs(tabs, widths, tree, onTree) else StackedTabs(tabs, tree, onTree, rowHeight = 44.dp)
    }
}

@Composable
private fun RowTabs(tabs: List<TalentTabModel>, widths: List<Int>, tree: Int, onTree: (Int) -> Unit) {
    val res = LocalResources.current
    Row(Modifier.fillMaxWidth().selectableGroup().drawLine(), horizontalArrangement = Arrangement.spacedBy(TAB_GUTTER)) {
        tabs.forEachIndexed { position, tab ->
            val selected = tab.index == tree
            Box(
                Modifier
                    .weight(widths[position].toFloat())
                    .heightIn(min = 48.dp)
                    .selectable(selected, role = Role.Tab) { onTree(tab.index) }
                    .clearAndSetSemantics {
                        contentDescription = tab.name
                        stateDescription = tab.accessibilityValue(res)
                    },
                contentAlignment = Alignment.Center,
            ) {
                TabLabel(tab, selected)
                if (selected) Box(Modifier.align(Alignment.BottomCenter).fillMaxWidth().height(3.dp).background(FsColor.gold))
            }
        }
    }
}

/** The 1 dp rule under the tab row. */
private fun Modifier.drawLine(): Modifier = drawBehind {
    drawRect(FsColor.line, topLeft = Offset(0f, size.height - 1.dp.toPx()), size = size.copy(height = 1.dp.toPx()))
}

@Composable
private fun TabLabel(tab: TalentTabModel, selected: Boolean) {
    val color = if (selected) FsColor.strong else FsColor.nav
    Row(horizontalArrangement = Arrangement.spacedBy(6.dp), verticalAlignment = Alignment.CenterVertically) {
        Text(tab.name, style = if (selected) FsType.bodyStrong else FsType.body, color = color, softWrap = false)
        Text(tab.points.toString(), style = FsType.mono, color = color, softWrap = false)
        if (tab.differing > 0) DiffDot()
    }
}

@Composable
private fun StackedTabs(tabs: List<TalentTabModel>, tree: Int, onTree: (Int) -> Unit, rowHeight: Dp) {
    val res = LocalResources.current
    val shape = RoundedCornerShape(FsRadius.control)
    Column(Modifier.fillMaxWidth().border(1.dp, FsColor.line, shape).background(FsColor.raised, shape).selectableGroup()) {
        tabs.forEach { tab ->
            val selected = tab.index == tree
            Row(
                Modifier
                    .fillMaxWidth()
                    .heightIn(min = rowHeight)
                    .background(if (selected) FsColor.gold.copy(alpha = TILE_ALPHA) else FsColor.raised)
                    .selectable(selected, role = Role.Tab) { onTree(tab.index) }
                    .clearAndSetSemantics {
                        contentDescription = tab.name
                        stateDescription = tab.accessibilityValue(res)
                    },
                verticalAlignment = Alignment.CenterVertically,
            ) {
                Box(Modifier.width(3.dp).height(rowHeight).background(if (selected) FsColor.gold else FsColor.raised))
                Text(
                    tab.name,
                    style = FsType.bodyStrong,
                    color = if (selected) FsColor.strong else FsColor.nav,
                    modifier = Modifier.weight(1f).padding(start = 11.dp, end = 8.dp, top = 8.dp, bottom = 8.dp),
                )
                Text(tab.points.toString(), style = FsType.mono, color = if (selected) FsColor.strong else FsColor.nav)
                if (tab.differing > 0) {
                    Spacer(Modifier.width(6.dp))
                    DiffDot()
                }
                Spacer(Modifier.width(14.dp))
            }
        }
    }
}

/** The 8 dp wipe dot on a tab whose tree holds a dashed cell. */
@Composable
private fun DiffDot() {
    Box(Modifier.size(8.dp).background(FsColor.wipe, CircleShape))
}

@Composable
private fun TalentsFooter(openPlanner: () -> Unit) {
    Column(Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, top = 16.dp, bottom = 32.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
        Text(stringResource(R.string.talents_footer), style = FsType.meta, color = FsColor.muted)
        Row(
            Modifier.heightIn(min = 48.dp).clickable(role = Role.Button, onClick = openPlanner),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(6.dp),
        ) {
            Text(stringResource(R.string.talents_planLink), style = FsType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f, fill = false))
            Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(16.dp))
        }
    }
}

/** First load with nothing saved: skeletons at the real sizes, never a spinner (spec F6 §6). */
@Composable
fun TalentsSkeleton() {
    val loading = stringResource(R.string.talents_loading)
    Column(
        Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, top = 8.dp).semantics(mergeDescendants = true) {
            contentDescription = loading
            liveRegion = LiveRegionMode.Polite
        },
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Bar(72, 22)
        Bar(220, 14)
        Box(Modifier.padding(top = 12.dp).fillMaxWidth().height(48.dp).background(FsColor.cardTop, RoundedCornerShape(SEGMENT_PILL)))
        Row(Modifier.fillMaxWidth().height(48.dp), horizontalArrangement = Arrangement.SpaceBetween, verticalAlignment = Alignment.CenterVertically) {
            repeat(3) { Bar(100, 14) }
        }
        BoxWithConstraints(Modifier.padding(top = 16.dp).fillMaxWidth().height(TalentGeometry.height)) {
            SKELETON_CELLS.forEach { cell ->
                val (x, y) = TalentGeometry.origin(cell, maxWidth + 32.dp)
                Box(Modifier.offset(x - 16.dp, y).size(TalentGeometry.CELL).background(FsColor.cardTop, RoundedCornerShape(FsRadius.control)))
            }
        }
        Text(loading, style = FsType.meta, color = FsColor.muted, modifier = Modifier.clearAndSetSemantics {})
    }
}

@Composable
private fun Bar(width: Int, height: Int) {
    Box(Modifier.width(width.dp).height(height.dp).background(FsColor.cardTop, RoundedCornerShape(4.dp)))
}

private val SKELETON_CELLS = listOf(
    TalentCell(0, 1), TalentCell(0, 2), TalentCell(1, 0), TalentCell(1, 1), TalentCell(1, 2), TalentCell(1, 3),
    TalentCell(2, 0), TalentCell(2, 1), TalentCell(3, 0), TalentCell(3, 1), TalentCell(3, 3), TalentCell(4, 1),
    TalentCell(5, 2), TalentCell(6, 1),
)

private val SEGMENT_PILL = 24.dp
private const val TILE_ALPHA = 0.12f
/** The row tab's pieces beside its two labels: the label gap, the dot and its gap, side padding. */
private val LABEL_GAP = 6.dp
private val DOT_GAP = 6.dp
private val DOT_SIZE = 8.dp
private val TAB_PADDING = 2.dp
/** The least space between two row tabs; with less, the tabs stack (F6 ux re-check). */
private val TAB_GUTTER = 12.dp
