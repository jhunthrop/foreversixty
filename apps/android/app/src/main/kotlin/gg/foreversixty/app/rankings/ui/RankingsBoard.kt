package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.ExperimentalFoundationApi
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.combinedClickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.widthIn
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.HorizontalDivider
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
import androidx.compose.ui.layout.onGloballyPositioned
import androidx.compose.ui.layout.positionInRoot
import androidx.compose.ui.platform.LocalResources
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
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.rankings.logic.Board
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.OwnRow
import gg.foreversixty.app.rankings.logic.Rankings
import gg.foreversixty.app.rankings.logic.RankingsCopy
import gg.foreversixty.app.rankings.logic.RankingsFilters
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.reports.ui.Chevron
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.reports.ui.SkeletonBar
import gg.foreversixty.app.reports.ui.isLargeText
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
    /** The `reports` switch: a row's menu offers `Open the report`. */
    val reportsOn: Boolean = false,
)

/** What the board hands back to its route. */
data class RankingsActions(
    val back: () -> Unit = {},
    val retry: () -> Unit = {},
    val openFilters: () -> Unit = {},
    val chooseSpec: () -> Unit = {},
    val open: (RankingRow) -> Unit = {},
    val loadMore: () -> Unit = {},
    val jump: () -> Unit = {},
    /** Your row's top, in root coordinates, so `Jump to your row` can scroll to it. */
    val ownRowAt: (Float) -> Unit = {},
    val watch: (RankingRow) -> Unit = {},
    val openReport: (RankingRow) -> Unit = {},
    val share: (RankingRow) -> Unit = {},
)

object RankingsLayout {
    val rowHeight = 68.dp
    val crest = 28.dp
    val rankWidth = 36.dp
    val goldEdge = 3.dp
    const val SKELETON_ROWS = 6
}

/** F11.1 and F11.3 as drawn from a state: the bar, the answer card, the filter line and the board. */
@Composable
fun RankingsContent(state: RankingsScreenState, now: Instant, actions: RankingsActions = RankingsActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val copy = RankingsCopy(LocalResources.current, zone)
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        ReportsBar(stringResource(R.string.rank_title), stringResource(R.string.tab_home), actions.back)
        val board = (state.phase as? BoardPhase.Loaded)?.board
        board?.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now, zone))) }
        Column(Modifier.padding(horizontal = 16.dp).padding(top = 8.dp), verticalArrangement = Arrangement.spacedBy(16.dp)) {
            when (val phase = state.phase) {
                BoardPhase.Loading -> Skeleton(state.viewer.signedIn)
                BoardPhase.Failed -> ErrorRegion(stringResource(R.string.rank_error), actions.retry, inset = false)
                is BoardPhase.Loaded -> state.filters?.let { Loaded(phase.board, it, state, now, copy, actions) }
            }
        }
    }
}

@Composable
private fun Loaded(board: Board, filters: RankingsFilters, state: RankingsScreenState, now: Instant, copy: RankingsCopy, actions: RankingsActions) {
    if (state.viewer.signedIn) AnswerCard(board, filters, state, copy)
    FilterLine(filters, state, copy, actions)
    if (board.rows.isEmpty()) {
        Text(
            stringResource(R.string.rank_empty), style = FsType.body, color = FsColor.text,
            modifier = Modifier.fillMaxWidth().quietCard().padding(16.dp),
        )
        return
    }
    val own = Rankings.placement(board.rows, board.characterRank)
    if (own is OwnRow.InPlace) {
        Text(
            copy.jump(own.rank), style = FsType.bodyStrong, color = FsColor.gold,
            modifier = Modifier.heightIn(min = 48.dp).clickable(role = Role.Button, onClick = actions.jump).padding(vertical = 12.dp),
        )
    }
    val crests = Rankings.showsCrests(board.rows)
    if (own is OwnRow.Pinned) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            MicroLabel(stringResource(R.string.rank_pin))
            Column(Modifier.fillMaxWidth().quietCard()) {
                RankRow(own.rank.row.copy(rank = own.rank.rank), board.total, filters, crests, own = true, now, copy, state, actions)
            }
        }
        HorizontalDivider(color = FsColor.line)
    }
    Column(Modifier.fillMaxWidth().quietCard()) {
        board.rows.forEachIndexed { index, row ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            val isOwn = own is OwnRow.InPlace && own.index == index
            val mark = if (isOwn) Modifier.onGloballyPositioned { actions.ownRowAt(it.positionInRoot().y) } else Modifier
            Box(mark) { RankRow(row, board.total, filters, crests, isOwn, now, copy, state, actions) }
        }
    }
    Footer(board, state.pageFailed, actions)
}

/** The answer (spec F11.1): `Your rank`, the ordinal at stat size, its unit, and the best kill's meta. */
@Composable
private fun AnswerCard(board: Board, filters: RankingsFilters, state: RankingsScreenState, copy: RankingsCopy) {
    Column(Modifier.fillMaxWidth().primaryCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(6.dp)) {
        val placed = board.characterRank
        if (placed == null) {
            Text(copy.unranked(filters, state.encounterNames), style = FsType.lead, color = FsColor.strong)
            Text(stringResource(R.string.rank_noneMeta), style = FsType.meta, color = FsColor.muted)
            return@Column
        }
        MicroLabel(stringResource(R.string.rank_label))
        val unit = copy.answerUnit(board.total, filters.cohort, board.deduped)
        val ordinal = Rankings.ordinal(placed.rank)
        if (isLargeText()) {
            Text(ordinal, style = FsType.stat, color = FsColor.strong)
            Text(unit, style = FsType.body, color = FsColor.text)
        } else {
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                Text(ordinal, style = FsType.stat, color = FsColor.strong, modifier = Modifier.alignByBaseline())
                Text(unit, style = FsType.body, color = FsColor.text, modifier = Modifier.alignByBaseline())
            }
        }
        val meta = copy.answerMeta(placed.row, filters, state.encounterNames, state.viewer.current?.name)
        Text(
            buildAnnotatedString {
                val name = meta.name
                if (name != null && meta.text.startsWith(name)) {
                    withStyle(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold)) { append(name) }
                    append(meta.text.removePrefix(name))
                } else {
                    append(meta.text)
                }
            },
            style = FsType.meta, color = FsColor.muted,
        )
    }
}

/** `{Encounter} · {Ruleset} · {Region} · {Phase} · {Measure}`, wrapping, and the 48 dp Filters button. */
@Composable
private fun FilterLine(filters: RankingsFilters, state: RankingsScreenState, copy: RankingsCopy, actions: RankingsActions) {
    val line = copy.filterLine(filters, state.encounterNames)
    val words = @Composable { modifier: Modifier ->
        Column(modifier, verticalArrangement = Arrangement.spacedBy(4.dp)) {
            if (!state.viewer.signedIn) {
                val pick = filters.cohort?.let(copy::cohortName)?.ifEmpty { null } ?: stringResource(R.string.rank_chooseSpec)
                Text(
                    pick, style = FsType.bodyStrong, color = FsColor.gold,
                    modifier = Modifier.heightIn(min = 48.dp).clickable(role = Role.Button, onClick = actions.chooseSpec).padding(vertical = 12.dp),
                )
            }
            Text(line, style = FsType.body, color = FsColor.text)
        }
    }
    val button = @Composable {
        BrandButton(
            stringResource(R.string.rank_filters), BrandButtonKind.Secondary, actions.openFilters,
            height = 48.dp, fill = false, leadingIcon = FiltersIcon,
        )
    }
    if (isLargeText()) {
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            words(Modifier.fillMaxWidth())
            button()
        }
    } else {
        Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(12.dp)) {
            words(Modifier.weight(1f))
            button()
        }
    }
}

/**
 * One board row (68 dp min): the rank in its parse colour, the class-coloured name with `You` on
 * yours, the meta, the value with its unit under it, a chevron. A long press opens the row's menu.
 */
@OptIn(ExperimentalFoundationApi::class)
@Composable
private fun RankRow(
    row: RankingRow,
    total: Int,
    filters: RankingsFilters,
    crest: Boolean,
    own: Boolean,
    now: Instant,
    copy: RankingsCopy,
    state: RankingsScreenState,
    actions: RankingsActions,
) {
    var menu by remember { mutableStateOf(false) }
    val yours = Rankings.isYours(row, state.viewer)
    val label = copy.rowLabel(row, total, filters.measure, yours)
    val menuActions = rowActions(row, state, actions)
    val ground = if (own) {
        Modifier.background(FsColor.cardTop).drawBehind {
            drawLine(FsColor.gold, Offset(0f, 0f), Offset(0f, size.height), (RankingsLayout.goldEdge * 2).toPx())
        }
    } else {
        Modifier
    }
    Box {
        Row(
            Modifier.fillMaxWidth().then(ground)
                .combinedClickable(onClick = { actions.open(row) }, onLongClick = { menu = true })
                .heightIn(min = RankingsLayout.rowHeight).padding(horizontal = 16.dp, vertical = 12.dp)
                .clearAndSetSemantics {
                    contentDescription = label
                    role = Role.Button
                    customActions = menuActions.map { (name, run) -> CustomAccessibilityAction(name) { run(); true } }
                },
            verticalAlignment = if (isLargeText()) Alignment.Top else Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(12.dp),
        ) {
            Text(
                row.rank.toString(), style = RankNumeral, color = parseColor(Rankings.rankTier(row.rank, total)),
                textAlign = TextAlign.Center, modifier = Modifier.widthIn(min = RankingsLayout.rankWidth),
            )
            if (crest) ClassCrest(row.player.className?.lowercase(), RankingsLayout.crest)
            if (isLargeText()) {
                Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(4.dp)) {
                    Name(row, yours)
                    Text(copy.rowMeta(row, now), style = FsType.meta, color = FsColor.muted)
                    Value(row, filters, copy, Alignment.Start)
                }
            } else {
                Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
                    Name(row, yours)
                    Text(copy.rowMeta(row, now), style = FsType.meta, color = FsColor.muted)
                }
                Value(row, filters, copy, Alignment.End)
            }
            Chevron(24.dp)
        }
        RowMenu(menu, menuActions) { menu = false }
    }
}

@Composable
private fun Name(row: RankingRow, yours: Boolean) {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(8.dp)) {
        Text(row.player.name, style = FsType.bodyStrong, color = classColor(row.player.className))
        if (yours) Tag(stringResource(R.string.rank_you), FsColor.gold)
    }
}

@Composable
private fun Value(row: RankingRow, filters: RankingsFilters, copy: RankingsCopy, align: Alignment.Horizontal) {
    Column(horizontalAlignment = align) {
        Text(Rankings.valueText(row, filters.measure), style = FsType.mono.copy(fontSize = 17.sp), color = FsColor.strong)
        Text(copy.unit(filters.measure), style = FsType.micro.copy(letterSpacing = 0.sp), color = FsColor.muted)
    }
}

/**
 * A row's menu, the long press and TalkBack's actions alike: Watch or Stop watching (signed in),
 * Open the report (with the reports switch on), Share.
 */
@Composable
private fun rowActions(row: RankingRow, state: RankingsScreenState, actions: RankingsActions): List<Pair<String, () -> Unit>> {
    val watching = row.player.key in state.watched
    val watch = stringResource(if (watching) R.string.rank_menu_unwatch else R.string.rank_menu_watch)
    val report = stringResource(R.string.rank_menu_report)
    val share = stringResource(R.string.rank_menu_share)
    return listOfNotNull(
        if (state.viewer.signedIn) watch to { actions.watch(row) } else null,
        if (state.reportsOn) report to { actions.openReport(row) } else null,
        share to { actions.share(row) },
    )
}

@Composable
private fun RowMenu(expanded: Boolean, items: List<Pair<String, () -> Unit>>, close: () -> Unit) {
    DropdownMenu(expanded = expanded, onDismissRequest = close) {
        items.forEach { (name, run) ->
            DropdownMenuItem(text = { Text(name) }, onClick = { close(); run() })
        }
    }
}

@Composable
private fun Footer(board: Board, pageFailed: Boolean, actions: RankingsActions) {
    when {
        pageFailed -> Row(Modifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.rank_error), style = FsType.meta, color = FsColor.wipe, modifier = Modifier.weight(1f))
            BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.loadMore, height = 48.dp, fill = false)
        }
        board.hasMore -> {
            // The spinner row loads the next 100 as it scrolls into view.
            LaunchedEffect(board.page) { actions.loadMore() }
            val loading = stringResource(R.string.rank_loadingMore)
            Box(Modifier.fillMaxWidth().heightIn(min = 56.dp).clearAndSetSemantics { contentDescription = loading }, contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold, modifier = Modifier.size(24.dp), strokeWidth = 2.dp)
            }
        }
    }
}

/** The loading state: the answer card and six rows at their real heights. */
@Composable
private fun Skeleton(signedIn: Boolean) {
    Column(Modifier.clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(16.dp)) {
        if (signedIn) {
            Column(Modifier.fillMaxWidth().primaryCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
                SkeletonBar(80.dp, 11.dp)
                SkeletonBar(220.dp, 28.dp)
                SkeletonBar(260.dp, 13.dp)
            }
        }
        SkeletonBar(280.dp, 15.dp)
        Column(Modifier.fillMaxWidth().quietCard()) {
            repeat(RankingsLayout.SKELETON_ROWS) { index ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                Row(
                    Modifier.fillMaxWidth().height(RankingsLayout.rowHeight).padding(horizontal = 16.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    horizontalArrangement = Arrangement.spacedBy(12.dp),
                ) {
                    SkeletonBar(20.dp, 16.dp)
                    Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(8.dp)) { SkeletonBar(120.dp, 14.dp); SkeletonBar(200.dp, 11.dp) }
                    SkeletonBar(36.dp, 16.dp)
                }
            }
        }
    }
}

/** The rank numeral: mono 700. */
private val RankNumeral = FsType.mono.copy(fontWeight = FontWeight.Bold, fontSize = 16.sp)
