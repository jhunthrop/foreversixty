package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.layout.Arrangement
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.FsType
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.widthIn
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.GameIconLabel
import gg.foreversixty.app.raids.logic.CompositionRead

/**
 * The composition block at the top of the roster: classes with names and counts, roles against the
 * raid's targets, and the raid buffs present or missing. Every figure comes from the API.
 */
@Composable
fun CompositionBlock(read: CompositionRead) {
    val large = isLargeText()
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.raidQuietCard().padding(14.dp)) {
        SectionLabel(stringResource(R.string.raids_comp_title))
        FlowRow(
            horizontalArrangement = Arrangement.spacedBy(8.dp),
            verticalArrangement = Arrangement.spacedBy(10.dp),
            maxItemsInEachRow = if (large) 1 else CLASS_COLUMNS,
            modifier = Modifier.fillMaxWidth().clearAndSetSemantics { contentDescription = read.classesAccessibilityLabel },
        ) {
            read.classes.forEach { line ->
                Row(
                    horizontalArrangement = Arrangement.spacedBy(8.dp),
                    verticalAlignment = Alignment.CenterVertically,
                    modifier = Modifier.weight(1f),
                ) {
                    ClassCrest(line.className, 24.dp)
                    Text(
                        line.name,
                        style = FsType.meta,
                        color = if (line.count == 0) FsColor.muted else classColor(line.className),
                        maxLines = 1,
                        softWrap = false,
                    )
                    Text("${line.count}", style = FsType.monoMeta, color = if (line.count == 0) FsColor.muted else FsColor.strong)
                }
            }
        }
        FlowRow(horizontalArrangement = Arrangement.spacedBy(16.dp), verticalArrangement = Arrangement.spacedBy(6.dp)) {
            read.roles.forEach { Text(it, style = FsType.bodyStrong, color = FsColor.strong) }
        }
        Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            SectionLabel(stringResource(R.string.raids_comp_buffs))
            FlowRow(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
                read.buffs.forEach { BuffTile(it) }
            }
            Text(read.buffSummary, style = FsType.meta, color = if (read.allCovered) FsColor.kill else FsColor.muted)
        }
    }
}

private const val CLASS_COLUMNS = 3

/**
 * One raid buff: its icon with the short name under it, live art or not (`Fort`, `SPro`), so two
 * buffs never differ by their art alone. TalkBack reads one node with the full name and the count.
 */
@Composable
private fun BuffTile(buff: CompositionRead.BuffLine) {
    Column(
        Modifier.widthIn(min = BUFF_TILE).clearAndSetSemantics { contentDescription = buff.accessibilityLabel },
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(3.dp),
    ) {
        GameIcon(buff.icon, BUFF_TILE, border = if (buff.present) FsColor.kill.copy(alpha = 0.6f) else FsColor.line, dimmed = !buff.present)
        GameIconLabel(buff.label, dimmed = !buff.present)
    }
}

private val BUFF_TILE = 36.dp
