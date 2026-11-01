package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.ExperimentalFoundationApi
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListState
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.NewDataStrip
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.guides.logic.GuideSection
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.reports.ui.BarAction
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.reports.ui.SkeletonBar
import gg.foreversixty.app.reports.ui.scaledSize
import java.time.Instant
import java.util.Locale
import kotlinx.coroutines.launch

/** The new-data strip's state over an open guide (spec §5.7). */
data class NewDataState(val failed: Boolean = false)

/** What the reader draws: the guide, the bar's words while it loads, and the strips over it. */
data class GuideReaderState(
    val phase: Loadable<GuideFile>,
    /** The bar's title while the file loads: `Elemental Shaman` from the index. */
    val title: String = "",
    val classSlug: String? = null,
    val newData: NewDataState? = null,
    /** The section the reader opens on (a capture scrolled to Rotation). */
    val initialSection: GuideSection = GuideSection.Overview,
)

class GuideReaderActions(
    val back: () -> Unit = {},
    val share: (String) -> Unit = {},
    val retry: () -> Unit = {},
    val update: () -> Unit = {},
    /** A section chip pressed: the selection haptic. */
    val selected: () -> Unit = {},
    /** The Gear row: the BiS tab on that spec's list. */
    val openBis: (GuideFile) -> Unit = {},
)

/** F9.2: the head band, the sticky section strip, the at-a-glance card and the guide's sections. */
@Composable
fun GuideReader(state: GuideReaderState, now: Instant, actions: GuideReaderActions) {
    val guide = (state.phase as? Loadable.Loaded)?.value
    val title = guide?.let { specTitle(it.specName, it.className) } ?: state.title
    val share = guide?.let { file -> BarAction(stringResource(R.string.guides_share), { rememberVectorPainter(Icons.Filled.Share) }) { actions.share(Guides.shareLink(file)) } }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        ReportsBar(title, stringResource(R.string.guides_title), actions.back, share, titleColor = classColor(guide?.classSlug ?: state.classSlug))
        state.newData?.let { NewDataStrip(stringResource(if (it.failed) R.string.offline_updateFailed else R.string.offline_newData), stringResource(R.string.offline_update), actions.update, it.failed) }
        when (val phase = state.phase) {
            Loadable.Loading -> ReaderSkeleton()
            Loadable.NotCached -> Text(stringResource(R.string.guides_notCached), style = FsType.body, color = FsColor.muted, modifier = Modifier.padding(GuideSpace.margin))
            Loadable.Failed -> ErrorRegion(stringResource(R.string.guides_error), actions.retry)
            is Loadable.Loaded -> {
                phase.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
                GuideBody(phase.value, state.initialSection, actions)
            }
        }
    }
}

@OptIn(ExperimentalFoundationApi::class)
@Composable
private fun GuideBody(guide: GuideFile, initial: GuideSection, actions: GuideReaderActions) {
    val sections = remember(guide) { Guides.sections(guide) }
    val list = rememberLazyListState()
    val stripOffset = -with(LocalDensity.current) { STRIP_HEIGHT.roundToPx() }
    val current by remember(sections) { derivedStateOf { currentSection(list, sections.size) } }
    LaunchedEffect(guide, initial) {
        // Opening on a section: its label lands just under the sticky strip.
        sections.indexOf(initial).takeIf { it > 0 }?.let { list.scrollToItem(it + FIRST_SECTION_ITEM, stripOffset) }
    }
    LazyColumn(Modifier.fillMaxSize(), state = list, contentPadding = PaddingValues(bottom = GuideSpace.section)) {
        item { Head(guide) }
        stickyHeader { SectionStrip(sections, current, list, stripOffset, actions.selected) }
        item { Glance(guide) }
        itemsIndexed(sections) { _, section -> GuideSectionView(section, guide, actions) }
        item { Column(Modifier.screenMargins()) { SourcesDisclosure(guide.sources) } }
    }
}

/**
 * The chip of the section under the strip (the glance card counts as the first); at the end, the
 * last. A section counts once its label reaches the strip's lower edge.
 */
private fun currentSection(list: LazyListState, count: Int): Int {
    val last = (count - 1).coerceAtLeast(0)
    if (list.firstVisibleItemIndex > 0 && !list.canScrollForward) return last
    val info = list.layoutInfo.visibleItemsInfo
    val under = info.firstOrNull { it.index != STRIP_ITEM && it.offset + it.size > stripBottom(info) }?.index ?: list.firstVisibleItemIndex
    return (under - FIRST_SECTION_ITEM).coerceIn(0, last)
}

/** Where the sticky strip ends, in list pixels (0 before it sticks). */
private fun stripBottom(info: List<androidx.compose.foundation.lazy.LazyListItemInfo>): Int =
    info.firstOrNull { it.index == STRIP_ITEM }?.let { it.offset + it.size } ?: 0

@Composable
private fun Head(guide: GuideFile) {
    val copy = GuidesCopy(LocalResources.current)
    Row(
        Modifier.fillMaxWidth().surfaceBand().padding(GuideSpace.margin),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        SpecIcon(guide.icon, guide.build, scaledSize(HEAD_ICON, HEAD_ICON_CAP))
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Text(specTitle(guide.specName, guide.className).uppercase(Locale.ROOT), style = FsType.h2, color = classColor(guide.classSlug))
            Text(copy.headMeta(guide), style = FsType.meta, color = FsColor.nav)
        }
    }
}

/** The sticky strip: one chip per section, the current one gold and kept in view. */
@Composable
private fun SectionStrip(sections: List<GuideSection>, current: Int, list: LazyListState, offset: Int, selected: () -> Unit) {
    val copy = GuidesCopy(LocalResources.current)
    val scope = rememberCoroutineScope()
    val strip = rememberLazyListState()
    LaunchedEffect(current) { strip.animateScrollToItem(current) }
    LazyRow(
        Modifier.fillMaxWidth().heightIn(min = STRIP_HEIGHT).background(FsColor.bg),
        state = strip,
        contentPadding = PaddingValues(horizontal = GuideSpace.margin, vertical = 2.dp),
        horizontalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        itemsIndexed(sections) { index, section ->
            PillChip(copy.section(section), selected = index == current, description = copy.sectionLabel(sections, section), onClick = {
                selected()
                scope.launch { list.animateScrollToItem(index + FIRST_SECTION_ITEM, offset) }
            })
        }
    }
}

/** The answer, first (the one warm card): the role and the stat priority as numbered chips. */
@Composable
private fun Glance(guide: GuideFile) {
    val copy = GuidesCopy(LocalResources.current)
    Column(
        Modifier.screenMargins().padding(top = GuideSpace.rowGap).primaryCard().cardInside().fillMaxWidth(),
        verticalArrangement = Arrangement.spacedBy(GuideSpace.labelGap),
    ) {
        guide.roleLine?.let { line ->
            MicroLabel(stringResource(R.string.guides_role))
            Text(line, style = FsType.lead.copy(fontWeight = FsType.body.fontWeight), color = FsColor.strong)
        }
        if (guide.statPriority.isNotEmpty()) {
            MicroLabel(stringResource(R.string.guides_statPriority), Modifier.padding(top = 6.dp))
            NumberedChips(guide.statPriority, copy::statLabel)
        }
    }
}

/** F9.7 over a guide: the head band and the glance card at their heights. */
@Composable
private fun ReaderSkeleton() {
    Column(Modifier.fillMaxWidth()) {
        Row(Modifier.fillMaxWidth().surfaceBand().padding(GuideSpace.margin), horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard)) {
            SkeletonBar(HEAD_ICON, HEAD_ICON)
            Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
                SkeletonBar(200.dp, 24.dp)
                SkeletonBar(240.dp, 12.dp)
            }
        }
        Column(Modifier.screenMargins().padding(top = 60.dp).quietCard().cardInside().fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            SkeletonBar(60.dp, 10.dp)
            SkeletonBar(180.dp, 16.dp)
            SkeletonBar(260.dp, 72.dp)
        }
    }
}

/** List items before the first section: the head, the strip and the glance card. */
private const val FIRST_SECTION_ITEM = 3
private const val STRIP_ITEM = 1

/** The strip: a 48 dp chip target and its padding. */
private val STRIP_HEIGHT = 52.dp
private val HEAD_ICON = 56.dp
private val HEAD_ICON_CAP = 72.dp
