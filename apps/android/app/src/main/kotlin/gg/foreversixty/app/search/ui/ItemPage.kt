package gg.foreversixty.app.search.ui

import androidx.annotation.DrawableRes
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListState
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemTooltipPanel
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.search.logic.ItemPageCopy
import gg.foreversixty.app.search.logic.ItemPageModel
import gg.foreversixty.app.search.logic.ItemPagePhase
import gg.foreversixty.app.search.logic.ItemSourceLine
import gg.foreversixty.app.search.logic.ItemSpecLine
import gg.foreversixty.app.search.logic.ItemUsesPhase
import gg.foreversixty.app.sims.ui.Bar
import java.util.Locale
import kotlinx.coroutines.launch

/** The page's sections, in order, each a strip chip (spec F12.4). */
enum class ItemSection { Item, Source, BisFor }

data class ItemPageActions(
    val onBack: () -> Unit = {},
    val onRetry: () -> Unit = {},
    val onRetryUses: () -> Unit = {},
    val onOpenSpec: (ItemSpecLine) -> Unit = {},
)

/**
 * F12.4 and F12.5: back (no share), the head, your pick, the F2.6 tooltip body, where it comes
 * from and the specs that list it. The strip chips pin under the bar once the head scrolls away.
 * [startAt] opens the page scrolled to a section (the F12.5 capture).
 */
@Composable
fun ItemPage(
    phase: ItemPagePhase,
    uses: ItemUsesPhase,
    build: String,
    backLabel: String,
    actions: ItemPageActions,
    startAt: ItemSection? = null,
) {
    val list = rememberLazyListState()
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        ItemBar(titleOf(phase), backLabel, actions.onBack)
        Box(Modifier.weight(1f)) {
            when (phase) {
                ItemPagePhase.Unknown -> Text(
                    stringResource(R.string.itemPage_unknown), style = FsType.lead, color = FsColor.strong,
                    modifier = Modifier.padding(SearchLayout.margin),
                )
                is ItemPagePhase.Loading -> LoadingBody(phase.look, build)
                is ItemPagePhase.Failed -> FailedBody(phase.look, build, actions.onRetry)
                is ItemPagePhase.Loaded -> LoadedBody(phase.model, uses, build, actions, list, startAt)
            }
        }
    }
}

private fun titleOf(phase: ItemPagePhase): String? = when (phase) {
    is ItemPagePhase.Loading -> phase.look.name
    is ItemPagePhase.Failed -> phase.look.name
    is ItemPagePhase.Loaded -> phase.model.look.name
    ItemPagePhase.Unknown -> null
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun ItemBar(title: String?, backLabel: String, onBack: () -> Unit) {
    TopAppBar(
        title = { title?.let { Text(it.uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong, maxLines = 2) } },
        navigationIcon = {
            IconButton(onClick = onBack) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = backLabel, tint = FsColor.strong) }
        },
        colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
        // The overlay already pads for the system bars.
        windowInsets = WindowInsets(0, 0, 0, 0),
    )
}

@Composable
private fun LoadedBody(
    model: ItemPageModel,
    uses: ItemUsesPhase,
    build: String,
    actions: ItemPageActions,
    list: LazyListState,
    startAt: ItemSection?,
) {
    val scope = rememberCoroutineScope()
    var expanded by rememberSaveable { mutableStateOf(false) }
    val first = if (model.pick == null) TOOLTIP_INDEX - 1 else TOOLTIP_INDEX
    fun indexOf(section: ItemSection) = first + section.ordinal
    // A section lands under the pinned strip, not behind it.
    val stripOffset = with(LocalDensity.current) { -STRIP_HEIGHT.roundToPx() }
    LaunchedEffect(startAt) { startAt?.let { list.scrollToItem(indexOf(it), stripOffset) } }
    val pinned by remember { derivedStateOf { list.firstVisibleItemIndex > 0 } }
    // At the end of the page the last section is the one being read, even if it cannot reach the top.
    val current by remember {
        derivedStateOf {
            if (!list.canScrollForward && pinned) {
                ItemSection.BisFor
            } else {
                ItemSection.entries.lastOrNull { indexOf(it) <= list.firstVisibleItemIndex } ?: ItemSection.Item
            }
        }
    }
    LazyColumn(Modifier.fillMaxSize(), state = list) {
        item(key = "head") { Head(model.look, model.headMeta, build) }
        model.pick?.let { pick ->
            item(key = "pick") {
                Row(
                    Modifier.padding(horizontal = SearchLayout.margin).padding(bottom = 16.dp).fillMaxWidth().primaryCard()
                        .semantics(mergeDescendants = true) {}.padding(14.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                ) {
                    ClassCrest(pick.className, 40.dp)
                    Column(Modifier.weight(1f)) {
                        Text(pick.title, style = FsType.bodyStrong, color = FsColor.strong)
                        Text(pick.meta, style = FsType.body, color = FsColor.muted)
                    }
                    Tag(stringResource(R.string.itemPage_yours), FsColor.gold)
                }
            }
        }
        item(key = "tooltip") { Box(Modifier.padding(horizontal = SearchLayout.margin)) { ItemTooltipPanel(model.tooltip) } }
        item(key = "sources") { Sources(uses, actions.onRetryUses) }
        item(key = "specs") { Specs(uses, build, expanded, { expanded = true }, actions.onOpenSpec) }
        item { Spacer(Modifier.height(32.dp)) }
    }
    if (pinned) {
        StripChips(current) { section -> scope.launch { list.animateScrollToItem(indexOf(section), stripOffset) } }
    }
}

/** The tooltip's place in the list when the pick card shows. */
private const val TOOLTIP_INDEX = 2

/** The pinned strip: one row of 48 dp chip targets plus its bottom gap. */
private val STRIP_HEIGHT = 56.dp

@Composable
private fun StripChips(current: ItemSection, onChip: (ItemSection) -> Unit) {
    val labels = listOf(R.string.itemPage_chip_item, R.string.itemPage_chip_source, R.string.itemPage_chip_bisFor)
    Row(
        Modifier.fillMaxWidth().heightIn(min = STRIP_HEIGHT).background(FsColor.bg).horizontalScroll(rememberScrollState())
            .padding(horizontal = SearchLayout.margin),
        horizontalArrangement = Arrangement.spacedBy(8.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        ItemSection.entries.forEachIndexed { i, section ->
            ScopeChip(stringResource(labels[i]), section == current) { onChip(section) }
        }
    }
}

/** Icon 64 in its quality border, the name in quality colour (h2, not uppercased), the meta. */
@Composable
private fun Head(look: ItemLook, meta: String?, build: String) {
    Row(
        Modifier.padding(SearchLayout.margin).fillMaxWidth().semantics(mergeDescendants = true) { heading() },
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(14.dp),
    ) {
        ItemIcon(look.icon, look.quality, build, 64.dp)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(look.name ?: stringResource(R.string.items_notInData), style = FsType.h2, color = ItemQuality(look.quality).text)
            meta?.let { Text(it, style = FsType.body, color = FsColor.muted) }
        }
    }
}

/** The tile at its real size with the index's name; the rest skeleton. */
@Composable
private fun LoadingBody(look: ItemLook, build: String) {
    Column {
        Head(look, null, build)
        Column(Modifier.padding(horizontal = SearchLayout.margin), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            Bar(200.dp, 14.dp)
            Box(Modifier.fillMaxWidth().height(220.dp).quietCard())
            Bar(160.dp, 12.dp)
            Box(Modifier.fillMaxWidth().height(128.dp).quietCard())
        }
    }
}

@Composable
private fun FailedBody(look: ItemLook, build: String, onRetry: () -> Unit) {
    Column {
        Head(look, null, build)
        RegionError(stringResource(R.string.items_failed), onRetry)
    }
}

/** An error in its own region, in wipe red, with Try again (spec §5.0). */
@Composable
private fun RegionError(text: String, onRetry: () -> Unit) {
    Column(Modifier.padding(horizontal = SearchLayout.margin, vertical = 8.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Text(text, style = FsType.body, color = FsColor.wipe)
        BrandButton(stringResource(R.string.search_tryAgain), BrandButtonKind.Secondary, onRetry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
    }
}

@Composable
private fun Foot(text: String) {
    Text(text, style = FsType.body, color = FsColor.muted, modifier = Modifier.padding(horizontal = SearchLayout.margin, vertical = 12.dp))
}

@Composable
private fun Sources(uses: ItemUsesPhase, onRetry: () -> Unit) {
    Column {
        SectionLabel(stringResource(R.string.itemPage_sources))
        when (uses) {
            ItemUsesPhase.Loading -> SkeletonRows()
            ItemUsesPhase.Failed -> RegionError(stringResource(R.string.itemPage_usesFailed), onRetry)
            is ItemUsesPhase.Loaded -> if (uses.sources.isEmpty()) {
                Foot(stringResource(R.string.itemPage_noSources))
            } else {
                RowCard(uses.sources.size) { i -> SourceRow(uses.sources[i]) }
                Foot(stringResource(R.string.itemPage_sourcesFoot))
            }
        }
    }
}

@Composable
private fun SourceRow(line: ItemSourceLine) {
    val description = stringResource(R.string.itemPage_a11y_source, line.title, line.meta)
    Row(
        Modifier.fillMaxWidth().clearAndSetSemantics { contentDescription = description }
            .padding(horizontal = SearchLayout.margin, vertical = 12.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(14.dp),
    ) {
        Box(
            Modifier.size(SearchLayout.itemIcon).background(FsColor.cardTop, RoundedCornerShape(FsRadius.control)),
            contentAlignment = Alignment.Center,
        ) {
            Icon(painterResource(sourceGlyph(line.kind)), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(20.dp))
        }
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(line.title, style = FsType.bodyStrong, color = FsColor.strong)
            if (line.meta.isNotEmpty()) Text(line.meta, style = FsType.body, color = FsColor.muted)
        }
    }
}

/** A source's tile glyph: swords for a boss or PvP, the book for a quest, the shield otherwise. */
@DrawableRes
private fun sourceGlyph(kind: String): Int = when (kind) {
    "raid", "dungeon", "pvp" -> R.drawable.ic_swords
    "quest" -> R.drawable.ic_guide
    else -> R.drawable.ic_shield
}

@Composable
private fun Specs(uses: ItemUsesPhase, build: String, expanded: Boolean, onExpand: () -> Unit, onOpen: (ItemSpecLine) -> Unit) {
    Column {
        SectionLabel(stringResource(R.string.itemPage_bisFor))
        when (uses) {
            ItemUsesPhase.Loading -> SkeletonRows()
            // The sources region carries the error and its Try again; one error per failure.
            ItemUsesPhase.Failed -> Unit
            is ItemUsesPhase.Loaded -> if (uses.specs.isEmpty()) {
                Foot(stringResource(R.string.itemPage_bisForNone))
            } else {
                val (shown, more) = ItemPageCopy.visibleSpecs(uses.specs, expanded)
                RowCard(shown.size) { i -> SpecRow(shown[i], build) { onOpen(shown[i]) } }
                if (more > 0) MoreSpecs(more, onExpand)
                Foot(stringResource(R.string.itemPage_bisForFoot))
            }
        }
    }
}

@Composable
private fun SpecRow(line: ItemSpecLine, build: String, onClick: () -> Unit) {
    val description = stringResource(R.string.itemPage_a11y_spec, line.title, line.meta, line.tag)
    Row(
        Modifier.fillMaxWidth().heightIn(min = SearchLayout.rowMinHeight)
            .clickable(role = Role.Button, onClick = onClick)
            .clearAndSetSemantics { contentDescription = description }
            .padding(horizontal = SearchLayout.margin, vertical = 12.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(14.dp),
    ) {
        line.icon?.let { SpecIcon(it, build, SearchLayout.icon) } ?: ClassCrest(line.classSlug, SearchLayout.icon)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(line.title, style = FsType.bodyStrong, color = classColor(line.classSlug))
            Text(line.meta, style = FsType.body, color = FsColor.muted)
        }
        Tag(line.tag, if (line.isPick) FsColor.kill else FsColor.muted)
        Chevron()
    }
}

@Composable
private fun MoreSpecs(count: Int, onExpand: () -> Unit) {
    Box(Modifier.padding(top = 10.dp)) {
        RowCard(1) {
            TappableRow(onExpand, minHeight = 56.dp) {
                Text(pluralStringResource(R.plurals.itemPage_moreSpecs, count, count), style = FsType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f))
                Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted)
            }
        }
    }
}

@Composable
private fun SkeletonRows() {
    RowCard(2) {
        Row(Modifier.padding(horizontal = SearchLayout.margin, vertical = 12.dp), horizontalArrangement = Arrangement.spacedBy(14.dp)) {
            Bar(SearchLayout.itemIcon, SearchLayout.itemIcon)
            Column(verticalArrangement = Arrangement.spacedBy(8.dp)) { Bar(180.dp, 15.dp); Bar(120.dp, 12.dp) }
        }
    }
}
