package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material.icons.filled.KeyboardArrowUp
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.SegmentedButton
import androidx.compose.material3.SegmentedButtonDefaults
import androidx.compose.material3.SingleChoiceSegmentedButtonRow
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.pillSurface
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.GuideSourceLink
import gg.foreversixty.app.raids.ui.SectionLabel
import gg.foreversixty.app.reports.ui.scaledSize

/** Spacing of the guide screens (spec W3 §8): 22 before a section label, 14 inside cards, 16 margins. */
internal object GuideSpace {
    val margin = 16.dp
    val section = 22.dp
    val inCard = 14.dp
    val labelGap = 8.dp
    val rowGap = 12.dp
}

/** A chip or segment's minimum target (48 dp, spec §6) and the 36 dp pill the eye sees. */
internal val TARGET = 48.dp
internal val PILL = 36.dp

/** Equal segments (spec F9.1), every one as tall as the tallest when a label wraps at large text. */
@Composable
internal fun EqualSegments(labels: List<String>, selected: Int, onSelect: (Int) -> Unit) {
    SingleChoiceSegmentedButtonRow(Modifier.fillMaxWidth().height(IntrinsicSize.Min)) {
        labels.forEachIndexed { index, label ->
            SegmentedButton(
                selected = index == selected,
                onClick = { if (index != selected) onSelect(index) },
                shape = SegmentedButtonDefaults.itemShape(index, labels.size),
                modifier = Modifier.fillMaxHeight().heightIn(min = TARGET),
                icon = {},
                colors = SegmentedButtonDefaults.colors(
                    activeContainerColor = GameColors.pressedWash,
                    activeContentColor = FsColor.strong,
                    activeBorderColor = FsColor.line,
                    inactiveContainerColor = GameColors.controlFill,
                    inactiveContentColor = FsColor.muted,
                    inactiveBorderColor = FsColor.line,
                ),
            ) {
                Text(label, style = FsType.segment, maxLines = 2, overflow = TextOverflow.Clip)
            }
        }
    }
}

/** A 36 dp pill in a 48 dp target: a guide section chip or a tier role chip; [selected] is gold. */
@Composable
internal fun PillChip(label: String, selected: Boolean, description: String = label, onClick: () -> Unit, leading: @Composable () -> Unit = {}, trailing: @Composable () -> Unit = {}) {
    val line = if (selected) FsColor.gold else FsColor.line
    val fill = if (selected) GameColors.pressedWash else GameColors.controlFill
    Box(
        Modifier
            .heightIn(min = TARGET)
            .clickable(role = Role.Tab, onClick = onClick)
            .clearAndSetSemantics {
                contentDescription = description
                this.selected = selected
            },
        contentAlignment = Alignment.Center,
    ) {
        Row(
            Modifier.heightIn(min = PILL).pillSurface(fill, line).padding(horizontal = 14.dp, vertical = 6.dp),
            horizontalArrangement = Arrangement.spacedBy(6.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            leading()
            Text(label, style = FsType.bodyStrong, color = if (selected) FsColor.gold else FsColor.text)
            trailing()
        }
    }
}

/** The numbered chip (spec §4): a 36 dp pill, a 22 dp number disc in mono 12, the stat name. Wraps, never cut. */
@OptIn(ExperimentalLayoutApi::class)
@Composable
internal fun NumberedChips(items: List<String>, label: (Int, String) -> String) {
    FlowRow(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        items.forEachIndexed { index, item ->
            val description = label(index, item)
            Row(
                Modifier
                    .heightIn(min = PILL)
                    .pillSurface(GameColors.controlFill, FsColor.line)
                    .padding(start = 7.dp, end = 14.dp, top = 6.dp, bottom = 6.dp)
                    .clearAndSetSemantics { contentDescription = description },
                horizontalArrangement = Arrangement.spacedBy(8.dp),
                verticalAlignment = Alignment.CenterVertically,
            ) {
                Box(Modifier.size(scaledSize(DISC, DISC_CAP)).background(FsColor.cardTop, CircleShape), contentAlignment = Alignment.Center) {
                    Text((index + 1).toString(), style = FsType.mono.copy(fontSize = DISC_TEXT), color = FsColor.gold)
                }
                Text(item, style = FsType.body, color = FsColor.text)
            }
        }
    }
}

private val DISC = 22.dp
private val DISC_CAP = 44.dp
private val DISC_TEXT = 12.sp

/** `Saved` with a check (kill), or `Not saved` with the cloud-down glyph (muted): a class page row's trail. */
@Composable
internal fun SavedMark(saved: Boolean, long: Boolean = false) {
    if (saved) {
        Tag(stringResource(if (long) R.string.guides_saved else R.string.guides_savedShort), FsColor.kill, checked = true)
    } else {
        Row(horizontalArrangement = Arrangement.spacedBy(6.dp), verticalAlignment = Alignment.CenterVertically) {
            Icon(painterResource(R.drawable.ic_cloud_download), contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp))
            Text(stringResource(R.string.guides_notSaved), style = FsType.meta, color = FsColor.muted)
        }
    }
}

/** A section: its micro label, then [content]. */
@Composable
internal fun GuideSectionBlock(title: String, modifier: Modifier = Modifier, content: @Composable () -> Unit) {
    Column(modifier.fillMaxWidth().padding(top = GuideSpace.section), verticalArrangement = Arrangement.spacedBy(GuideSpace.labelGap)) {
        SectionLabel(title)
        content()
    }
}

/** A quiet card whose rows are split by the soft line. */
@Composable
internal fun <T> QuietRows(items: List<T>, row: @Composable (T) -> Unit) {
    Column(Modifier.fillMaxWidth().quietCard()) {
        items.forEachIndexed { index, item ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            row(item)
        }
    }
}

/** `Sources`, one quiet disclosure at a reader's foot (provenance stays quiet); a label opens its page. */
@Composable
internal fun SourcesDisclosure(sources: List<GuideSourceLink>) {
    if (sources.isEmpty()) return
    val context = LocalContext.current
    var open by rememberSaveable { mutableStateOf(false) }
    Column(Modifier.fillMaxWidth().padding(top = GuideSpace.section).quietCard()) {
        Row(
            Modifier.fillMaxWidth().heightIn(min = TARGET).clickable(role = Role.Button) { open = !open }
                .semantics { selected = open }
                .padding(horizontal = GuideSpace.margin),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Text(stringResource(R.string.guides_sources), style = FsType.bodyStrong, color = FsColor.text, modifier = Modifier.weight(1f))
            Icon(if (open) Icons.Filled.KeyboardArrowUp else Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted)
        }
        if (open) {
            sources.forEach { source ->
                HorizontalDivider(color = FsColor.lineSoft)
                Text(
                    source.label, style = FsType.body, color = FsColor.gold,
                    modifier = Modifier.fillMaxWidth().heightIn(min = TARGET).clickable(role = Role.Button) { AccountLinks.open(context, source.url) }
                        .padding(horizontal = GuideSpace.margin, vertical = GuideSpace.rowGap),
                )
            }
        }
    }
}

/** A plain paragraph in a card, or a list line led by a muted bullet. */
@Composable
internal fun Lines(lines: List<String>, bulleted: Boolean = false) {
    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
        lines.forEach { line ->
            if (bulleted) {
                Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                    Text(BULLET, style = FsType.body, color = FsColor.muted)
                    Text(line, style = FsType.body, color = FsColor.text)
                }
            } else {
                Text(line, style = FsType.body, color = FsColor.text)
            }
        }
    }
}

private const val BULLET = "•"

/** Card padding shorthand. */
internal fun Modifier.cardInside(): Modifier = padding(horizontal = GuideSpace.margin, vertical = GuideSpace.inCard)

internal fun Modifier.screenMargins(): Modifier = fillMaxWidth().padding(horizontal = GuideSpace.margin)
