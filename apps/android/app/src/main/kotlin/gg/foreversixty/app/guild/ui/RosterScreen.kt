package gg.foreversixty.app.guild.ui

import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.common.SavedAnswerStrip
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListScope
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.OutlinedTextFieldDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.minimumInteractiveComponentSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalWindowInfo
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.guild.logic.Loadable
import gg.foreversixty.app.guild.logic.RosterCopy
import gg.foreversixty.app.guild.logic.RosterEntries
import gg.foreversixty.app.guild.logic.RosterEntry
import gg.foreversixty.app.guild.logic.RosterFigure
import gg.foreversixty.app.guild.logic.RosterList
import gg.foreversixty.app.guild.logic.RosterQuery
import gg.foreversixty.app.guild.logic.RosterRole
import gg.foreversixty.app.guild.logic.RosterSort
import gg.foreversixty.app.guild.logic.RosterStanding
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildMemberCounts
import gg.foreversixty.app.raids.ui.ErrorRegion
import gg.foreversixty.app.raids.ui.GameIcon
import gg.foreversixty.app.raids.ui.NameText
import gg.foreversixty.app.raids.ui.isLargeText
import gg.foreversixty.app.raids.ui.raidQuietCard
import java.time.Instant

/** Below this width the attendance bar hides and `{n} of 8` stays (spec 9). */
private const val NARROW_WIDTH_DP = 340

/** The counts card's middle figure's share of the row: its caption is the longest. */
private const val WIDE_FIGURE = 1.6f

/** What the roster screen holds besides its data: the chips, the search and the sort menu. */
class RosterControls(
    query: RosterQuery = RosterQuery(),
    searching: Boolean = false,
    sortMenuOpen: Boolean = false,
) {
    var query by mutableStateOf(query)
    var searching by mutableStateOf(searching)
    var sortMenuOpen by mutableStateOf(sortMenuOpen)
}

/** F16.1 and F16.2: pure over its state, so the screenshot tests draw every state. */
@Composable
fun RosterContent(
    state: Loadable<GuildHome>,
    controls: RosterControls,
    stripTime: (Instant) -> String,
    open: (RosterEntry) -> Unit,
    retry: () -> Unit,
    back: () -> Unit,
) {
    val copy = RosterCopy(LocalResources.current)
    Column(Modifier.fillMaxSize().screenGround()) {
        GuildTopBar(stringResource(R.string.roster_title), back) {
            IconButton(onClick = { controls.searching = !controls.searching }) {
                Icon(Icons.Filled.Search, contentDescription = stringResource(R.string.roster_search), tint = FsColor.strong)
            }
        }
        if (controls.searching) SearchField(controls)
        LazyColumn(Modifier.fillMaxSize()) {
            when (state) {
                Loadable.Loading -> skeleton()
                Loadable.Failed -> item { Box(Modifier.padding(16.dp)) { ErrorRegion(stringResource(R.string.roster_error), retry) } }
                is Loadable.Loaded -> loaded(state, controls, copy, stripTime, open, retry)
            }
        }
    }
}

@Composable
private fun SearchField(controls: RosterControls) {
    OutlinedTextField(
        value = controls.query.search,
        onValueChange = { controls.query = controls.query.copy(search = it) },
        singleLine = true,
        placeholder = { Text(stringResource(R.string.roster_search), style = ScaledType.body, color = FsColor.muted) },
        leadingIcon = { Icon(Icons.Filled.Search, contentDescription = null, tint = FsColor.muted) },
        textStyle = ScaledType.body.copy(color = FsColor.strong),
        colors = OutlinedTextFieldDefaults.colors(
            focusedBorderColor = FsColor.gold,
            unfocusedBorderColor = FsColor.line,
            cursorColor = FsColor.gold,
            focusedContainerColor = FsColor.raised,
            unfocusedContainerColor = FsColor.raised,
        ),
        modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 4.dp),
    )
}

private fun LazyListScope.loaded(
    state: Loadable.Loaded<GuildHome>,
    controls: RosterControls,
    copy: RosterCopy,
    stripTime: (Instant) -> String,
    open: (RosterEntry) -> Unit,
    retry: () -> Unit,
) {
    val home = state.value
    val entries = RosterEntries.of(home)
    val rows = RosterList.apply(entries, controls.query)
    val onlyCount = RosterList.onlyCount(entries, controls.query.only)
    state.offlineSince?.let { since -> item { SavedAnswerStrip(stripTime(since), state.stale, retry) } }
    item { CountsCard(home.counts, onlyCount) }
    item { Chips(controls, onlyCount) }
    if (rows.isEmpty()) {
        item { NoMatch() }
    } else {
        itemsIndexed(rows, key = { _, entry -> entry.characterKey ?: "name:${entry.name}" }) { index, entry ->
            Box(Modifier.cardSegment(first = index == 0, last = index == rows.lastIndex)) {
                RosterRowView(entry, copy, open)
            }
        }
        item {
            Text(
                stringResource(R.string.roster_footer),
                style = ScaledType.meta,
                color = FsColor.muted,
                modifier = Modifier.padding(start = 16.dp, end = 16.dp, top = 12.dp, bottom = 24.dp),
            )
        }
    }
}

/** The three figures; filtered from the Raid tab ([shown]), the roster's reads `4 of 38` (round 5 row 2). */
@Composable
private fun CountsCard(counts: GuildMemberCounts, shown: Int?) {
    val figures = RosterFigure.of(counts, shown).map { figure ->
        val value = figure.of?.let { total -> stringResource(R.string.roster_counts_of, figure.value, total) } ?: "${figure.value}"
        value to stringResource(figure.kind.caption)
    }
    val figure = @Composable { value: String, caption: String, modifier: Modifier ->
        Column(modifier.semantics(mergeDescendants = true) {}) {
            Text(value, style = ScaledType.stat, color = FsColor.strong)
            Text(caption, style = ScaledType.meta, color = FsColor.muted)
        }
    }
    val card = Modifier.padding(start = 16.dp, end = 16.dp, top = 14.dp).raidQuietCard().padding(horizontal = 16.dp, vertical = 14.dp)
    if (isLargeText()) {
        Column(verticalArrangement = Arrangement.spacedBy(10.dp), modifier = card) {
            figures.forEach { (value, caption) -> figure(value, caption, Modifier) }
        }
    } else {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), modifier = card) {
            figures.forEachIndexed { index, (value, caption) ->
                // The middle caption is the longest; it takes the room it needs.
                figure(value, caption, Modifier.weight(if (index == 1) WIDE_FIGURE else 1f))
            }
        }
    }
}

/**
 * All, Tanks, Healers, DPS, then `Sort: {by}`; from the Raid tab's not-synced line ([onlyCount]) the
 * filter's own chip leads, selected, with its clear control, and All is not selected (round 5 row 2).
 */
@Composable
private fun Chips(controls: RosterControls, onlyCount: Int?) {
    // The chips wrap, so the Sort chip stays whole at the largest text.
    FlowRow(
        horizontalArrangement = Arrangement.spacedBy(8.dp),
        verticalArrangement = Arrangement.spacedBy(4.dp),
        itemVerticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp),
    ) {
        val haptics = rememberHaptics()
        val clearOnly = { controls.query = controls.query.copy(only = emptySet()) }
        onlyCount?.let { FilterChip(stringResource(R.string.roster_chip_notSynced, it), clearOnly) }
        RosterRole.entries.forEach { role ->
            Chip(stringResource(role.label), selected = controls.query.role == role && (role != RosterRole.All || onlyCount == null)) {
                haptics.selection()
                // All is everyone: it clears the not-synced filter too.
                controls.query = controls.query.copy(role = role, only = if (role == RosterRole.All) emptySet() else controls.query.only)
            }
        }
        Box {
            Chip(
                stringResource(R.string.roster_sort, stringResource(controls.query.sort.label)),
                selected = false,
                trailing = true,
            ) { controls.sortMenuOpen = true }
            SortMenu(controls)
        }
    }
}

@Composable
private fun SortMenu(controls: RosterControls) {
    val haptics = rememberHaptics()
    DropdownMenu(expanded = controls.sortMenuOpen, onDismissRequest = { controls.sortMenuOpen = false }, containerColor = GameColors.sheet) {
        RosterSort.entries.forEach { sort ->
            val current = controls.query.sort == sort
            DropdownMenuItem(
                text = { Text(stringResource(sort.label), style = ScaledType.body, color = FsColor.strong) },
                trailingIcon = if (current) {
                    { Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.gold) }
                } else {
                    null
                },
                onClick = {
                    haptics.selection()
                    controls.query = controls.query.copy(sort = sort)
                    controls.sortMenuOpen = false
                },
                modifier = Modifier.semantics { selected = current },
            )
        }
    }
}

/** A 40 dp pill in a 48 dp target: gold ring and wash when selected (spec 7, Scope chip). */
@Composable
private fun Chip(text: String, selected: Boolean, trailing: Boolean = false, onClick: () -> Unit) {
    val shape = RoundedCornerShape(FsRadius.control)
    Row(
        horizontalArrangement = Arrangement.spacedBy(6.dp),
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .minimumInteractiveComponentSize()
            .heightIn(min = 40.dp)
            .background(if (selected) FsColor.gold.copy(alpha = 0.10f) else FsColor.strong.copy(alpha = 0.04f), shape)
            .border(1.dp, if (selected) FsColor.gold else FsColor.line, shape)
            .clickable(onClick = onClick)
            .padding(horizontal = 14.dp)
            .semantics {
                role = if (trailing) Role.DropdownList else Role.Tab
                this.selected = selected
            },
    ) {
        Text(text, style = ScaledType.meta, color = if (selected) FsColor.strong else FsColor.text, maxLines = 1)
        if (trailing) Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(16.dp))
    }
}

/**
 * The not-synced filter's chip in [Chip]'s selected style, `Not synced this week · 4`, then its clear
 * control (`×`, a 48 dp target) that shows the whole roster again.
 */
@Composable
private fun FilterChip(text: String, clear: () -> Unit) {
    val shape = RoundedCornerShape(FsRadius.control)
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .minimumInteractiveComponentSize()
            .heightIn(min = 40.dp)
            .background(FsColor.gold.copy(alpha = 0.10f), shape)
            .border(1.dp, FsColor.gold, shape)
            .padding(start = 14.dp)
            .semantics { selected = true },
    ) {
        Text(text, style = ScaledType.meta, color = FsColor.strong, maxLines = 1)
        IconButton(onClick = clear, modifier = Modifier.size(FsTarget.min)) {
            Icon(Icons.Filled.Close, contentDescription = stringResource(R.string.roster_chip_clear), tint = FsColor.gold, modifier = Modifier.size(16.dp))
        }
    }
}

/** One roster row (64 min): the class crest, name and meta, attendance, chevron. */
@Composable
fun RosterRowView(entry: RosterEntry, copy: RosterCopy, open: (RosterEntry) -> Unit) {
    val large = isLargeText()
    val narrow = with(LocalDensity.current) { LocalWindowInfo.current.containerSize.width.toDp() } < NARROW_WIDTH_DP.dp
    val label = copy.a11y(entry)
    Row(
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 64.dp)
            .let { if (entry.opensPage) it.clickable { open(entry) } else it }
            .padding(horizontal = 14.dp, vertical = 10.dp)
            .semantics(mergeDescendants = true) { contentDescription = label },
    ) {
        Box(Modifier.clearAndSetSemantics {}) {
            ClassCrest(entry.className, 32.dp)
        }
        Column(verticalArrangement = Arrangement.spacedBy(2.dp), modifier = Modifier.weight(1f).clearAndSetSemantics {}) {
            Row(horizontalArrangement = Arrangement.spacedBy(4.dp), verticalAlignment = Alignment.CenterVertically) {
                NameText(entry.name, ScaledType.bodyStrong, classColor(entry.className), Modifier.weight(1f, fill = false))
                if (entry.supporter) {
                    Icon(painterResource(R.drawable.ic_supporter), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(13.dp))
                }
            }
            Text(copy.meta(entry), style = ScaledType.meta, color = FsColor.muted)
            // Largest text: the number only, under the name (spec 9).
            if (large) Standing(entry.standing, copy, showsBar = false, alignment = Alignment.Start)
        }
        if (!large) {
            Box(Modifier.clearAndSetSemantics {}) { Standing(entry.standing, copy, showsBar = !narrow, alignment = Alignment.End) }
        }
        if (entry.opensPage) {
            Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.clearAndSetSemantics {})
        }
    }
}

@Composable
private fun Standing(standing: RosterStanding, copy: RosterCopy, showsBar: Boolean, alignment: Alignment.Horizontal) {
    when (standing) {
        is RosterStanding.Attended -> Column(horizontalAlignment = alignment, verticalArrangement = Arrangement.spacedBy(4.dp)) {
            if (showsBar) AttendanceBar(standing)
            Text(copy.attendance(standing), style = ScaledType.meta.copy(fontFeatureSettings = TABULAR), color = FsColor.strong)
        }
        RosterStanding.NoNights -> Text(
            stringResource(R.string.roster_noNights),
            style = ScaledType.meta,
            color = FsColor.muted,
            textAlign = if (alignment == Alignment.End) TextAlign.End else TextAlign.Start,
            modifier = Modifier.widthIn(max = 120.dp),
        )
        RosterStanding.Waiting -> Tag(stringResource(R.string.roster_waitingTag), FsColor.gold)
        RosterStanding.NoAccount -> Tag(stringResource(R.string.roster_noAccountTag), FsColor.muted)
    }
}

/** `nights` segments of 7 x 12: kill for a night present, line for one missed. */
@Composable
private fun AttendanceBar(attended: RosterStanding.Attended) {
    Row(horizontalArrangement = Arrangement.spacedBy(2.dp)) {
        repeat(attended.nights) { night ->
            Box(
                Modifier
                    .size(7.dp, 12.dp)
                    .background(if (night < attended.present) FsColor.kill else FsColor.line, RoundedCornerShape(2.dp)),
            )
        }
    }
}

@Composable
private fun NoMatch() {
    Column(
        verticalArrangement = Arrangement.spacedBy(4.dp),
        modifier = Modifier.padding(horizontal = 16.dp, vertical = 6.dp).raidQuietCard().padding(16.dp),
    ) {
        Text(stringResource(R.string.roster_empty_title), style = ScaledType.lead, color = FsColor.strong)
        Text(stringResource(R.string.roster_empty_body), style = ScaledType.meta, color = FsColor.muted)
    }
}

/** First load: eight rows at their real height. */
private fun LazyListScope.skeleton() {
    item {
        Box(Modifier.padding(start = 16.dp, end = 16.dp, top = 14.dp, bottom = 12.dp).raidQuietCard().heightIn(min = 84.dp).clearAndSetSemantics {})
    }
    items(SKELETON_ROWS) { index ->
        Row(
            horizontalArrangement = Arrangement.spacedBy(12.dp),
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier
                .cardSegment(first = index == 0, last = index == SKELETON_ROWS - 1)
                .heightIn(min = 64.dp)
                .padding(horizontal = 14.dp)
                .clearAndSetSemantics {},
        ) {
            Box(Modifier.size(32.dp).background(FsColor.lineSoft, RoundedCornerShape(6.dp)))
            Column(verticalArrangement = Arrangement.spacedBy(6.dp), modifier = Modifier.weight(1f)) {
                SkeletonBar(120.dp, 15.dp)
                SkeletonBar(190.dp, 12.dp)
            }
            SkeletonBar(70.dp, 12.dp)
        }
    }
}

private const val SKELETON_ROWS = 8
