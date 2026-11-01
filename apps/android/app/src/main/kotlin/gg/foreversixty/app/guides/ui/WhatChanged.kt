package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.guides.logic.Changes
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.KeyDateRow
import gg.foreversixty.app.guides.logic.KeyDates
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.model.Changelog
import gg.foreversixty.app.model.ChangelogBlock
import gg.foreversixty.app.model.ChangelogEntry
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.reports.ui.SkeletonBar
import java.time.Instant
import java.time.LocalDate

data class WhatChangedState(
    val dates: Loadable<List<KeyDate>>,
    val changelog: Loadable<Changelog>,
    val today: LocalDate,
)

class WhatChangedActions(val open: (String) -> Unit = {}, val retry: () -> Unit = {})

/** F9.5: Key dates (past ones omitted), then the changelog newest first; each region fails alone. */
@Composable
fun WhatChangedContent(state: WhatChangedState, now: Instant, actions: WhatChangedActions) {
    val offline = listOf(state.dates, state.changelog).mapNotNull { (it as? Loadable.Loaded)?.offlineSince }.minOrNull()
    Column(Modifier.fillMaxWidth()) {
        offline?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
        Column(Modifier.screenMargins().padding(bottom = GuideSpace.section)) {
            Region(stringResource(R.string.changes_dates), state.dates, actions.retry) { dates ->
                val rows = KeyDates.upcoming(dates, state.today)
                if (rows.isNotEmpty()) QuietRows(rows) { DateRow(it) }
            }
            Region(stringResource(R.string.changes_list), state.changelog, actions.retry) { changelog ->
                QuietRows(changelog.entries) { entry -> ChangeRow(entry) { actions.open(entry.id) } }
            }
        }
    }
}

@Composable
private fun <T> Region(title: String, phase: Loadable<T>, retry: () -> Unit, content: @Composable (T) -> Unit) {
    GuideSectionBlock(title) {
        when (phase) {
            Loadable.Loading -> Column(Modifier.fillMaxWidth().quietCard().heightIn(min = SKELETON).cardInside(), verticalArrangement = Arrangement.spacedBy(10.dp)) {
                SkeletonBar(200.dp, 14.dp)
                SkeletonBar(260.dp, 14.dp)
            }
            Loadable.NotCached -> Text(stringResource(R.string.guides_listNotCached), style = FsType.body, color = FsColor.muted)
            Loadable.Failed -> ErrorRegion(stringResource(R.string.changes_error), retry, inset = false)
            is Loadable.Loaded -> content(phase.value)
        }
    }
}

@Composable
private fun DateRow(row: KeyDateRow) {
    val date = row.date
    Row(
        Modifier.fillMaxWidth().heightIn(min = ROW_MIN).cardInside(),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Box(
            Modifier.widthIn(min = BLOCK_WIDTH).heightIn(min = BLOCK_HEIGHT)
                .background(FsColor.bg.copy(alpha = BLOCK_GROUND), RoundedCornerShape(FsRadius.control))
                .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control))
                .padding(6.dp),
            contentAlignment = Alignment.Center,
        ) {
            Text(date.key, style = FsType.mono, color = FsColor.strong, textAlign = TextAlign.Center)
        }
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(date.value, style = FsType.bodyStrong, color = FsColor.strong)
            date.note?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
            row.running?.let { Text(it, style = FsType.meta, color = FsColor.gold) }
        }
    }
}

@Composable
private fun ChangeRow(entry: ChangelogEntry, open: () -> Unit) {
    val copy = GuidesCopy(LocalResources.current)
    Row(
        Modifier.fillMaxWidth().heightIn(min = ROW_MIN).clickable(role = Role.Button, onClick = open).cardInside(),
        horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard),
        verticalAlignment = Alignment.Top,
    ) {
        Text(GuideDates.monthDay(entry.date), style = FsType.monoMeta, color = FsColor.muted, modifier = Modifier.padding(top = 2.dp))
        Text(entry.title, style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Tag(copy.tag(Changes.tag(entry)), FsColor.muted)
    }
}

/** A changelog entry in the guide reader's style: its blocks and the Sources disclosure. */
@Composable
fun ChangeReader(entry: ChangelogEntry?, back: () -> Unit) {
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        ReportsBar(stringResource(R.string.changes_list), stringResource(R.string.guides_title), back)
        if (entry == null) return@Column
        Column(Modifier.verticalScroll(rememberScrollState()).screenMargins().padding(bottom = GuideSpace.section), verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap)) {
            Text(GuideDates.long(entry.date), style = FsType.meta, color = FsColor.muted)
            Text(entry.title, style = FsType.h2, color = FsColor.strong, modifier = Modifier.semantics { heading() })
            entry.blocks.forEach { Block(it) }
            SourcesDisclosure(entry.sources)
        }
    }
}

@Composable
private fun Block(block: ChangelogBlock) {
    when (block.kind) {
        BLOCK_HEADING -> Text(block.text, style = FsType.sec, color = FsColor.strong, modifier = Modifier.padding(top = GuideSpace.labelGap).semantics { heading() })
        BLOCK_ITEM -> Lines(listOf(block.text), bulleted = true)
        else -> Text(block.text, style = FsType.body, color = FsColor.text)
    }
}

private const val BLOCK_HEADING = "h"
private const val BLOCK_ITEM = "li"
private val ROW_MIN = 68.dp
private val SKELETON = 120.dp
private val BLOCK_WIDTH = 56.dp
private val BLOCK_HEIGHT = 48.dp
private const val BLOCK_GROUND = 0.5f
