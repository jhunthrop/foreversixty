package gg.foreversixty.app.guild.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListScope
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.guild.logic.HistoryCopy
import gg.foreversixty.app.guild.logic.Loadable
import gg.foreversixty.app.guild.logic.MoreState
import gg.foreversixty.app.guild.logic.NightFacts
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.ui.DateBlock
import gg.foreversixty.app.raids.ui.ErrorRegion
import gg.foreversixty.app.raids.ui.RaidSecondaryButton
import gg.foreversixty.app.raids.ui.SectionLabel
import gg.foreversixty.app.raids.ui.isLargeText
import gg.foreversixty.app.raids.ui.raidQuietCard
import java.time.Instant

/** F16.3 Raid history: pure over its state; [mine] is the account's character keys. */
@Composable
fun HistoryContent(
    state: Loadable<List<RaidNightRow>>,
    more: MoreState,
    mine: Set<String>,
    clock: RaidClock,
    stripTime: (Instant) -> String,
    open: (NightFacts) -> Unit,
    loadMore: () -> Unit,
    retry: () -> Unit,
    back: () -> Unit,
) {
    val copy = HistoryCopy(LocalResources.current, clock)
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        GuildTopBar(stringResource(R.string.history_title), back)
        LazyColumn(Modifier.fillMaxSize()) {
            when (state) {
                Loadable.Loading -> skeleton()
                Loadable.Failed -> item { Box(Modifier.padding(16.dp)) { ErrorRegion(stringResource(R.string.history_error), retry) } }
                is Loadable.Loaded -> {
                    state.offlineSince?.let { since -> item { OfflineStrip(stringResource(R.string.offline_strip, stripTime(since))) } }
                    if (state.value.isEmpty()) item { Empty() } else nights(state.value, more, mine, copy, clock, open, loadMore)
                }
            }
        }
    }
}

private fun LazyListScope.nights(
    rows: List<RaidNightRow>,
    more: MoreState,
    mine: Set<String>,
    copy: HistoryCopy,
    clock: RaidClock,
    open: (NightFacts) -> Unit,
    loadMore: () -> Unit,
) {
    item { Box(Modifier.padding(start = 16.dp, end = 16.dp, top = 12.dp, bottom = 8.dp)) { SectionLabel(stringResource(R.string.history_label)) } }
    itemsIndexed(rows, key = { _, row -> row.id }) { index, row ->
        Box(Modifier.cardSegment(first = index == 0, last = index == rows.lastIndex)) {
            NightRow(NightFacts(row, mine), copy, clock, open)
        }
    }
    item { MoreRow(more, loadMore) }
    item {
        Text(
            stringResource(R.string.history_footer),
            style = FsType.meta,
            color = FsColor.muted,
            modifier = Modifier.padding(start = 16.dp, end = 16.dp, top = 10.dp, bottom = 24.dp),
        )
    }
}

/** A night (72 min): the date block, the zone, two meta lines, then `You were there` under them. */
@Composable
fun NightRow(facts: NightFacts, copy: HistoryCopy, clock: RaidClock, open: (NightFacts) -> Unit) {
    val large = isLargeText()
    val label = copy.a11y(facts)
    val text = @Composable { modifier: Modifier ->
        Column(verticalArrangement = Arrangement.spacedBy(2.dp), modifier = modifier) {
            Text(facts.row.zone, style = FsType.bodyStrong, color = FsColor.strong)
            Text(copy.fights(facts), style = FsType.meta, color = FsColor.muted, maxLines = if (large) Int.MAX_VALUE else 1, softWrap = large)
            SizeLine(facts, copy, large)
            if (facts.wasThere) Box(Modifier.padding(top = 4.dp)) { Tag(stringResource(R.string.history_there), FsColor.kill) }
        }
    }
    val row = Modifier
        .fillMaxWidth()
        .heightIn(min = 72.dp)
        .clickable { open(facts) }
        .padding(horizontal = 16.dp, vertical = 12.dp)
        .semantics(mergeDescendants = true) {
            contentDescription = label
            role = Role.Button
        }
    if (large) {
        // Largest text: the date block above the zone.
        Column(verticalArrangement = Arrangement.spacedBy(8.dp), modifier = row) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Box(Modifier.weight(1f).clearAndSetSemantics {}) { DateBlock(facts.startedAt, clock, 56.dp, 52.dp) }
                Chevron()
            }
            text(Modifier.clearAndSetSemantics {})
        }
    } else {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.Top, modifier = row) {
            DateBlock(facts.startedAt, clock, 48.dp, 52.dp)
            text(Modifier.weight(1f).clearAndSetSemantics {})
            Chevron()
        }
    }
}

@Composable
private fun Chevron() {
    Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.clearAndSetSemantics {})
}

/** `38 raiders · 1 h 47 m` with mono figures; one line at the default sizes. */
@Composable
private fun SizeLine(facts: NightFacts, copy: HistoryCopy, large: Boolean) {
    val size = copy.size(facts)
    val count = "${facts.size}"
    val line = buildAnnotatedString {
        withStyle(SpanStyle(fontFamily = FsType.monoMeta.fontFamily, color = FsColor.text)) { append(count) }
        append(size.removePrefix(count))
        append(HistoryCopy.SEPARATOR)
        withStyle(SpanStyle(fontFamily = FsType.monoMeta.fontFamily, color = FsColor.text)) { append(copy.length(facts)) }
    }
    Text(line, style = FsType.meta, color = FsColor.muted, maxLines = if (large) Int.MAX_VALUE else 1, softWrap = large)
}

@Composable
private fun MoreRow(more: MoreState, loadMore: () -> Unit) {
    when (more) {
        MoreState.End -> Unit
        MoreState.Idle, MoreState.Loading -> Row(
            horizontalArrangement = Arrangement.spacedBy(10.dp),
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier
                .padding(horizontal = 16.dp)
                .fillMaxWidth()
                .heightIn(min = 52.dp)
                .clickable(enabled = more == MoreState.Idle, onClick = loadMore)
                .padding(horizontal = 16.dp)
                .semantics(mergeDescendants = true) { role = Role.Button },
        ) {
            Text(stringResource(R.string.history_more), style = FsType.bodyStrong, color = FsColor.gold)
            if (more == MoreState.Loading) CircularProgressIndicator(color = FsColor.gold, strokeWidth = 2.dp, modifier = Modifier.size(16.dp))
        }
        MoreState.Failed -> Column(
            verticalArrangement = Arrangement.spacedBy(10.dp),
            modifier = Modifier.padding(horizontal = 16.dp, vertical = 12.dp).padding(horizontal = 16.dp),
        ) {
            Text(stringResource(R.string.history_moreError), style = FsType.meta, color = FsColor.wipe)
            RaidSecondaryButton(stringResource(R.string.common_tryAgain), loadMore)
        }
    }
}

@Composable
private fun Empty() {
    Column(
        verticalArrangement = Arrangement.spacedBy(4.dp),
        modifier = Modifier.padding(16.dp).raidQuietCard().padding(16.dp),
    ) {
        Text(stringResource(R.string.history_empty_title), style = FsType.lead, color = FsColor.strong)
        Text(stringResource(R.string.history_empty_body), style = FsType.meta, color = FsColor.muted)
    }
}

private fun LazyListScope.skeleton() {
    items(SKELETON_ROWS) { index ->
        Row(
            horizontalArrangement = Arrangement.spacedBy(12.dp),
            modifier = Modifier
                .padding(top = if (index == 0) 40.dp else 0.dp)
                .cardSegment(first = index == 0, last = index == SKELETON_ROWS - 1)
                .heightIn(min = 96.dp)
                .padding(horizontal = 16.dp, vertical = 12.dp)
                .clearAndSetSemantics {},
        ) {
            Box(Modifier.size(48.dp, 52.dp).background(FsColor.lineSoft, RoundedCornerShape(FsRadius.control)))
            Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                SkeletonBar(160.dp, 15.dp)
                SkeletonBar(190.dp, 12.dp)
                SkeletonBar(140.dp, 12.dp)
            }
        }
    }
}

private const val SKELETON_ROWS = 6
