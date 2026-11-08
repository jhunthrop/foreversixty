package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.ExperimentalFoundationApi
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.combinedClickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material.icons.filled.Star
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.layout.onGloballyPositioned
import androidx.compose.ui.layout.positionInRoot
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.CustomAccessibilityAction
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.customActions
import androidx.compose.ui.semantics.role
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.DpOffset
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.common.SavedAnswerStrip
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowControls
import gg.foreversixty.app.flows.FlowField
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNotice
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowPlateMessage
import gg.foreversixty.app.flows.FlowPlateRetry
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSegments
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.flows.flowBleed
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.rankings.logic.Board
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.Measure
import gg.foreversixty.app.rankings.logic.OwnRow
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsFilters
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.reports.ui.parseColor
import java.time.Instant
import java.time.ZoneId

/** What the board draws, apart from its store so screenshots drive every state. */
data class RankingsScreenState(
    val phase: BoardPhase,
    val filters: RankingsFilters?,
    val encounterNames: Map<String, String>,
    val viewer: RankingsViewer,
    val pageFailed: Boolean = false,
    val watched: Set<String> = emptySet(),
    /** The `reports` switch: a row offers `Report`. */
    val reportsOn: Boolean = false,
)

/** What the board hands back to its route. */
data class RankingsActions(
    val retry: () -> Unit = {},
    /** Choose an encounter: the board of the one picked. */
    val choose: (String) -> Unit = {},
    val openFilters: () -> Unit = {},
    /** The encounter field: the filters sheet's encounter list. */
    val openEncounters: () -> Unit = {},
    /** The class and spec fields: the class and spec picker. */
    val chooseSpec: () -> Unit = {},
    val setMeasure: (Measure) -> Unit = {},
    val openTiers: () -> Unit = {},
    /** The Guilds board, on the site. */
    val openGuilds: () -> Unit = {},
    val open: (RankingRow) -> Unit = {},
    val loadMore: () -> Unit = {},
    val jump: () -> Unit = {},
    /** Your row's top, in root coordinates, so `Jump to your row` can scroll to it. */
    val ownRowAt: (Float) -> Unit = {},
    val watch: (RankingRow) -> Unit = {},
    val openReport: (RankingRow) -> Unit = {},
    val share: (RankingRow) -> Unit = {},
)

/**
 * Rankings (design/mocks/web web-rankings, web-rankings-you, web-rankings-empty): Choose an encounter,
 * or one encounter's board with your rank in the head, the controls bar, the board and the legend.
 */
@Composable
fun RankingsContent(state: RankingsScreenState, now: Instant, actions: RankingsActions = RankingsActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val copy = RankingsCopy(LocalResources.current, zone)
    val board = (state.phase as? BoardPhase.Loaded)?.board
    Column(Modifier.fillMaxWidth()) {
        board?.offlineSince?.let { SavedAnswerStrip(OfflineStamp.format(it, now, zone), board.stale, actions.retry) }
        Column(Modifier.fillMaxWidth().parityContent().padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space6)) {
            when (val phase = state.phase) {
                BoardPhase.Loading -> Loading(state, copy)
                BoardPhase.Failed -> Failed(state, copy, actions)
                is BoardPhase.Choose -> Chooser(phase.encounters, state, copy, actions)
                is BoardPhase.Loaded -> state.filters?.let { BoardView(phase.board, it, state, now, copy, actions) }
            }
        }
    }
}

@Composable
private fun PageHead(state: RankingsScreenState, copy: RankingsCopy) {
    FlowHead(stringResource(R.string.rank_title), listOf(copy.line(state.viewer.signedIn)))
}

@Composable
private fun Loading(state: RankingsScreenState, copy: RankingsCopy) {
    PageHead(state, copy)
    Column(Modifier.padding(top = SECTION_TOP).clearAndSetSemantics {}) {
        FlowSkeleton(SKELETON_TITLE, SKELETON_LINE)
        FlowPlate(Modifier.padding(top = PLATE_TOP)) {
            repeat(SKELETON_ROWS) { index ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                Column(Modifier.padding(FsSpace.space4), verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                    FlowSkeleton(SKELETON_NAME, SKELETON_LINE)
                    FlowSkeleton(SKELETON_SUB, SKELETON_SMALL)
                }
            }
        }
    }
}

@Composable
private fun Failed(state: RankingsScreenState, copy: RankingsCopy, actions: RankingsActions) {
    PageHead(state, copy)
    FlowPlate(Modifier.padding(top = SECTION_TOP)) {
        FlowPlateRetry(stringResource(R.string.rank_error), stringResource(R.string.common_tryAgain), actions.retry)
    }
}

/** No encounter to open on: every ranked encounter, or the Guilds board when there is none. */
@Composable
private fun Chooser(encounters: List<RankedEncounter>, state: RankingsScreenState, copy: RankingsCopy, actions: RankingsActions) {
    PageHead(state, copy)
    TierNotice(actions)
    val meta = if (encounters.isEmpty()) null else pluralStringResource(R.plurals.rank_flows_ranked, encounters.size, encounters.size)
    FlowSection(stringResource(R.string.rank_flows_choose), meta = meta, top = SECTION_TOP) {
        FlowPlate(Modifier.padding(top = PLATE_TOP)) {
            if (encounters.isEmpty()) {
                FlowPlateMessage(listOf(stringResource(R.string.rank_flows_noneYet), stringResource(R.string.rank_flows_tryGuilds))) {
                    FlowLink(stringResource(R.string.rank_flows_openGuilds), onClick = actions.openGuilds)
                }
            } else {
                val line = stringResource(R.string.rank_flows_encounterLine)
                val open = stringResource(R.string.rank_flows_openBoard)
                encounters.forEachIndexed { index, encounter ->
                    val pick = { actions.choose(encounter.slug) }
                    FlowRow(encounter.name, line, inset = FsSpace.space4, last = index == encounters.lastIndex, trail = { FlowLink(open, onClick = pick) }, onClick = pick)
                }
            }
        }
    }
    Legend(region = null)
}

/** `Comparing specs, not characters? Tier list ›`, edge to edge (the mock's summary). */
@Composable
private fun TierNotice(actions: RankingsActions) {
    Box(Modifier.fillMaxWidth().flowBleed().parityRegion(ParityRegion.SUMMARY).padding(horizontal = FsSpace.space4)) {
        FlowNotice(stringResource(R.string.rank_flows_tierQuestion), stringResource(R.string.rank_flows_tierLink), actions.openTiers)
    }
}

@Composable
private fun BoardView(board: Board, filters: RankingsFilters, state: RankingsScreenState, now: Instant, copy: RankingsCopy, actions: RankingsActions) {
    val own = Rankings.placement(board.rows, board.characterRank)
    FlowHead(
        copy.encounterName(filters, state.encounterNames),
        copy.facts(filters, board.updatedAt),
        answer = if (state.viewer.signedIn) ({ Answer(board, filters, own, state, copy, actions) }) else null,
    )
    Controls(filters, state, copy, actions)
    TierNotice(actions)
    val crests = Rankings.showsCrests(board.rows)
    if (own is OwnRow.Pinned) {
        FlowSection(stringResource(R.string.rank_pin), top = SECTION_TOP) {
            FlowPlate(Modifier.padding(top = PLATE_TOP)) {
                BoardRow(own.rank.row.copy(rank = own.rank.rank), board, filters, RowLook(crests, own = true, last = true), now, copy, state, actions)
            }
        }
    }
    FlowSection(copy.cohortTitle(filters.cohort), meta = copy.sectionMeta(board.total, board.deduped), top = SECTION_TOP) {
        FlowPlate(Modifier.padding(top = PLATE_TOP)) {
            if (board.rows.isEmpty()) {
                FlowPlateMessage(listOf(stringResource(R.string.rank_empty)))
            } else {
                board.rows.forEachIndexed { index, row ->
                    val isOwn = own is OwnRow.InPlace && own.index == index
                    val mark = if (isOwn) Modifier.onGloballyPositioned { actions.ownRowAt(it.positionInRoot().y) } else Modifier
                    Box(mark) { BoardRow(row, board, filters, RowLook(crests, isOwn, index == board.rows.lastIndex), now, copy, state, actions) }
                }
            }
            Footer(board, state.pageFailed, actions)
        }
    }
    Legend(region = ParityRegion.FOOTER_NOTE)
}

/** The head's answer: `Your rank`, the character, its ordinal of the cohort, where, and its links. */
@Composable
private fun Answer(board: Board, filters: RankingsFilters, own: OwnRow, state: RankingsScreenState, copy: RankingsCopy, actions: RankingsActions) {
    Text(stringResource(R.string.rank_label), style = BOARD_LABEL, color = FsColor.muted)
    val placed = board.characterRank
    if (placed == null) {
        Text(copy.unranked(filters, state.encounterNames), style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.padding(top = FsSpace.space1))
        Text(stringResource(R.string.rank_noneMeta), style = FsType.small, color = FsColor.muted)
        return
    }
    val row = placed.row
    Row(Modifier.padding(top = FsSpace.space1).height(WHO), verticalAlignment = Alignment.CenterVertically) {
        ClassCrest(row.player.className?.lowercase(), CREST, Modifier.parityItem(ParityItemKind.Image))
        Spacer(Modifier.width(FsSpace.space3))
        Text(row.player.name, style = WHO_NAME, color = classColor(row.player.className), maxLines = 1, modifier = Modifier.weight(1f, fill = false))
        Spacer(Modifier.width(FsSpace.space3))
        Text(
            buildAnnotatedString {
                append(Rankings.ordinal(placed.rank))
                withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) {
                    append(" " + copy.answerUnit(board.total, filters.cohort, board.deduped))
                }
            },
            style = FsType.figureBig,
            color = FsColor.gold,
            maxLines = 1,
        )
    }
    Text(copy.answerNote(filters, state.encounterNames), style = FsType.small, color = FsColor.muted, modifier = Modifier.padding(top = FsSpace.space1))
    Row(Modifier.padding(top = FsSpace.space1), horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
        FlowLink(stringResource(R.string.rank_flows_parses)) { actions.open(row) }
        if (own is OwnRow.InPlace) FlowLink(stringResource(R.string.rank_flows_jump), onClick = actions.jump)
    }
}

/** The controls bar: encounter, class, spec, ruleset, the metric and the board, then Filters. */
@Composable
private fun Controls(filters: RankingsFilters, state: RankingsScreenState, copy: RankingsCopy, actions: RankingsActions) {
    val cohort = filters.cohort
    FlowControls {
        FlowField(copy.encounterName(filters, state.encounterNames), actions.openEncounters)
        FlowField(copy.className(cohort?.classSlug).ifEmpty { stringResource(R.string.rank_flows_everyClassField) }, actions.chooseSpec)
        FlowField(cohort?.specName?.ifEmpty { null } ?: stringResource(R.string.rank_flows_everySpec), actions.chooseSpec)
        FlowField(copy.rulesetLabel(filters.ruleset), actions.openFilters)
        val measures = listOf(Measure.Damage, Measure.Healing)
        FlowSegments(measures.map { copy.measureLabel(it) }, measures.indexOf(filters.measure)) { actions.setMeasure(measures[it]) }
        FlowSegments(listOf(stringResource(R.string.rank_flows_characters), stringResource(R.string.rank_flows_guilds)), 0) { if (it == 1) actions.openGuilds() }
        FlowField(stringResource(R.string.rank_filters), actions.openFilters, button = true)
    }
}

/**
 * One board row (the mock's `.rr`): the rank in its parse colour, the class-coloured name with `You`,
 * who and where, the value and its age, executed and length, the meter against the board's first,
 * and Report and Build. A tap opens the character, a long press the row's menu.
 */
@OptIn(ExperimentalFoundationApi::class)
@Composable
private fun BoardRow(
    row: RankingRow,
    board: Board,
    filters: RankingsFilters,
    look: RowLook,
    now: Instant,
    copy: RankingsCopy,
    state: RankingsScreenState,
    actions: RankingsActions,
) {
    var menu by remember { mutableStateOf(false) }
    val yours = Rankings.isYours(row, state.viewer)
    val label = copy.rowLabel(row, board.total, filters.measure, yours)
    val menuActions = rowActions(row, state, actions)
    val ground = if (look.own) {
        Modifier.background(FsColor.selected).drawBehind {
            drawLine(FsColor.gold, Offset(0f, 0f), Offset(0f, size.height), (GOLD_EDGE * 2).toPx())
        }
    } else {
        Modifier
    }
    Box {
        Row(
            Modifier.fillMaxWidth().parityItem(ParityItemKind.Row).then(ground)
                .combinedClickable(onClick = { actions.open(row) }, onLongClick = { menu = true })
                .padding(start = ROW_SIDE, end = ROW_SIDE, top = ROW_TOP, bottom = ROW_BOTTOM),
        ) {
            Box(Modifier.widthIn(min = RANK_WIDTH).height(NAME_BLOCK), contentAlignment = Alignment.CenterEnd) {
                Text(row.rank.toString(), style = RANK, color = parseColor(Rankings.rankTier(row.rank, board.total)), textAlign = TextAlign.End)
            }
            Spacer(Modifier.width(FsSpace.space3))
            if (look.crest) {
                ClassCrest(row.player.className?.lowercase(), ROW_CREST, Modifier.parityItem(ParityItemKind.Image))
                Spacer(Modifier.width(FsSpace.space2))
            }
            Column(
                Modifier.weight(1f).clearAndSetSemantics {
                    contentDescription = label
                    role = Role.Button
                    customActions = menuActions.map { CustomAccessibilityAction(it.name) { it.run(); true } }
                },
            ) {
                RowBody(row, board, filters, yours, now, copy)
                RowLinks(row, state, actions)
            }
        }
        RowMenu(menu, menuActions) { menu = false }
    }
    if (!look.last) HorizontalDivider(color = FsColor.lineSoft)
}

@Composable
private fun RowBody(row: RankingRow, board: Board, filters: RankingsFilters, yours: Boolean, now: Instant, copy: RankingsCopy) {
    Row(verticalAlignment = Alignment.CenterVertically) {
        Row(Modifier.weight(1f), verticalAlignment = Alignment.CenterVertically) {
            Text(row.player.name, style = NAME, color = classColor(row.player.className), modifier = Modifier.weight(1f, fill = false))
            if (yours) PillTag(stringResource(R.string.rank_you), Modifier.padding(start = TAG_GAP))
        }
        Text(
            buildAnnotatedString {
                append(Rankings.valueText(row, filters.measure))
                withStyle(SpanStyle(fontSize = 12.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" " + copy.unit(filters.measure)) }
            },
            style = NAME,
            color = FsColor.strong,
        )
    }
    Row(verticalAlignment = Alignment.CenterVertically) {
        Text(copy.rowSub(row), style = SUB, color = FsColor.muted, modifier = Modifier.weight(1f))
        Text(copy.age(row.foughtAt, now), style = AGE, color = FsColor.text)
    }
    Row(Modifier.padding(top = FsSpace.space1 + 2.dp).height(FIGURE_LINE), horizontalArrangement = Arrangement.spacedBy(FsSpace.space5)) {
        copy.executed(row)?.let { Figure(stringResource(R.string.rank_flows_executed), it, FsColor.gold) }
        Figure(stringResource(R.string.rank_flows_length), copy.length(row), FsColor.strong)
    }
    val top = board.rows.maxOfOrNull { it.value }?.takeIf { it > 0 } ?: row.value
    Meter((row.value / top).toFloat().coerceIn(0f, 1f), classColor(row.player.className))
}

@Composable
private fun Figure(label: String, value: String, color: Color) {
    Text(
        buildAnnotatedString {
            withStyle(SpanStyle(color = FsColor.muted)) { append("$label ") }
            withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = color)) { append(value) }
        },
        style = FIGURE,
    )
}

/** The value against the board's first (`.meter.cls`): 8 tall in the class colour, 8 under the figures. */
@Composable
private fun Meter(fraction: Float, color: Color) {
    Box(Modifier.padding(top = FsSpace.space2).fillMaxWidth().height(METER).background(FsColor.lineSoft, RoundedCornerShape(METER_RADIUS))) {
        Box(Modifier.fillMaxWidth(fraction).fillMaxHeight().background(color, RoundedCornerShape(METER_RADIUS)))
    }
}

/** `Report` (with the reports switch on) and `Build` (the character's builds seen). */
@Composable
private fun RowLinks(row: RankingRow, state: RankingsScreenState, actions: RankingsActions) {
    Row(Modifier.height(LINKS), horizontalArrangement = Arrangement.spacedBy(FsSpace.space5), verticalAlignment = Alignment.CenterVertically) {
        if (state.reportsOn) {
            Text(stringResource(R.string.rank_flows_report), style = LINK, color = FsColor.gold, modifier = Modifier.clickable(role = Role.Button) { actions.openReport(row) })
        }
        Text(stringResource(R.string.rank_flows_build), style = LINK, color = FsColor.text, modifier = Modifier.clickable(role = Role.Button) { actions.open(row) })
    }
}

/** How a row draws: a crest before the name (a mixed board), yours (the gold edge), the last (no rule after it). */
private data class RowLook(val crest: Boolean, val own: Boolean, val last: Boolean)

/** A small outlined tag (`.pill.tag.you`): `You`. */
@Composable
private fun PillTag(text: String, modifier: Modifier = Modifier) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(modifier.border(1.dp, FsColor.gold, shape).padding(horizontal = TAG_PAD_X, vertical = TAG_PAD_Y)) {
        Text(text, style = TAG, color = FsColor.gold)
    }
}

/**
 * A row's menu, the long press and TalkBack's actions alike: Watch or Stop watching (signed in),
 * Open the report (with the reports switch on), Share.
 */
@Composable
private fun rowActions(row: RankingRow, state: RankingsScreenState, actions: RankingsActions): List<RowAction> {
    val watching = row.player.key in state.watched
    val watch = stringResource(if (watching) R.string.rank_menu_unwatch else R.string.rank_menu_watch)
    val report = stringResource(R.string.rank_menu_report)
    val share = stringResource(R.string.rank_menu_share)
    return listOfNotNull(
        if (state.viewer.signedIn) RowAction(watch, rememberVectorPainter(if (watching) Icons.Filled.Star else StarOutline)) { actions.watch(row) } else null,
        if (state.reportsOn) RowAction(report, painterResource(R.drawable.ic_logs)) { actions.openReport(row) } else null,
        RowAction(share, rememberVectorPainter(Icons.Filled.Share)) { actions.share(row) },
    )
}

/** One entry of a row's menu and TalkBack actions: its words, its glyph and what it does. */
private class RowAction(val name: String, val icon: Painter, val run: () -> Unit)

@Composable
private fun RowMenu(expanded: Boolean, items: List<RowAction>, close: () -> Unit) {
    DropdownMenu(
        expanded = expanded, onDismissRequest = close, offset = MENU_OFFSET,
        containerColor = FsColor.cardTop, border = BorderStroke(1.dp, FsColor.line),
    ) {
        items.forEach { item ->
            DropdownMenuItem(
                text = { Text(item.name, style = ScaledType.body, color = FsColor.strong) },
                leadingIcon = { Icon(item.icon, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(MENU_ICON)) },
                onClick = { close(); item.run() },
            )
        }
    }
}

/** The board's end: the next 100 loading as it scrolls into view, or the failed page and Try again. */
@Composable
private fun Footer(board: Board, pageFailed: Boolean, actions: RankingsActions) {
    when {
        pageFailed -> {
            HorizontalDivider(color = FsColor.lineSoft)
            FlowPlateRetry(stringResource(R.string.rank_error), stringResource(R.string.common_tryAgain), actions.loadMore)
        }
        board.hasMore && board.rows.isNotEmpty() -> {
            LaunchedEffect(board.page) { actions.loadMore() }
            val loading = stringResource(R.string.rank_loadingMore)
            HorizontalDivider(color = FsColor.lineSoft)
            Box(Modifier.fillMaxWidth().height(PAGER).clearAndSetSemantics { contentDescription = loading }, contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold, modifier = Modifier.size(SPINNER), strokeWidth = 2.dp)
            }
        }
    }
}

/** `Rank colour is the percentile: 100 99+ 95+ 75+ 50+ 25+ under 25`, each in its parse colour. */
@Composable
private fun Legend(region: String?) {
    val tagged = region?.let { Modifier.parityRegion(it) } ?: Modifier
    FlowRow(Modifier.padding(top = LEGEND_TOP).fillMaxWidth().then(tagged), horizontalArrangement = Arrangement.spacedBy(LEGEND_GAP)) {
        Text(stringResource(R.string.rank_flows_legend), style = LEGEND, color = FsColor.muted)
        LEGEND_STEPS.forEach { (label, color) -> Text(label, style = LEGEND_STEP, color = color) }
        Text(stringResource(R.string.rank_flows_under), style = LEGEND_STEP, color = FsColor.parseGreyText)
    }
}

/** The legend's steps: the top rank, then each percentile floor. */
private val LEGEND_STEPS = listOf(
    "100" to FsColor.parseGold,
    "99+" to FsColor.parsePink,
    "95+" to FsColor.parseOrange,
    "75+" to FsColor.parsePurple,
    "50+" to FsColor.parseBlue,
    "25+" to FsColor.parseGreen,
)

private val SECTION_TOP = 20.dp
private val PLATE_TOP = 12.dp
private val CREST = 40.dp
private val WHO = 49.dp
private val ROW_CREST = 28.dp
private val ROW_SIDE = 14.dp
private val ROW_TOP = 12.dp
private val ROW_BOTTOM = 3.dp
private val RANK_WIDTH = 22.dp
private val NAME_BLOCK = 38.dp
private val FIGURE_LINE = 20.dp
private val METER = 8.dp
private val METER_RADIUS = 4.dp
private val LINKS = 44.dp
private val GOLD_EDGE = 3.dp
private val TAG_GAP = 10.dp
private val TAG_PAD_X = 6.dp
private val TAG_PAD_Y = 1.dp
private val PAGER = 44.dp
private val SPINNER = 24.dp
private val MENU_ICON = 20.dp
private val LEGEND_TOP = 48.dp
private val LEGEND_GAP = 14.dp
private val SKELETON_TITLE = 160.dp
private val SKELETON_NAME = 140.dp
private val SKELETON_SUB = 220.dp
private val SKELETON_LINE = 16.dp
private val SKELETON_SMALL = 12.dp
private const val SKELETON_ROWS = 6
private val MENU_OFFSET = DpOffset(56.dp, 0.dp)
private val BOARD_LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val WHO_NAME = FsType.itemLg
private val NAME = FsType.figureSm
private val RANK = FsType.item
private val AGE = FsType.bodyStrong.copy(lineHeight = 18.sp)
private val FIGURE = FsType.small.copy(fontSize = 14.sp)
private val LINK = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
private val LEGEND = FsType.small.copy(lineHeight = 24.sp)
private val LEGEND_STEP = LEGEND.copy(fontWeight = FontWeight.Bold)
