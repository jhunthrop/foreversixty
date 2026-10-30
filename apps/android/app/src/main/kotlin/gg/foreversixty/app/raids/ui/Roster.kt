package gg.foreversixty.app.raids.ui

import android.content.ClipData
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.List
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material.icons.filled.Lock
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.ClipEntry
import androidx.compose.ui.platform.LocalClipboard
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.raids.logic.RaidChip
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidFacts
import gg.foreversixty.app.raids.logic.RosterGroups
import gg.foreversixty.app.raids.logic.RosterRow
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.roleLabel
import kotlinx.coroutines.launch

/** The officer row (spec F4.18), drawn only when `viewer.can` says so. */
@Composable
fun OfficerRow(controller: SignupController, clock: RaidClock) {
    val scope = rememberCoroutineScope()
    val clipboard = LocalClipboard.current
    val facts = controller.facts
    val locked = facts.chip == RaidChip.Locked
    val ended = controller.raid.viewer.toolsEnded == true
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        SectionLabel(stringResource(R.string.officer_label))
        FlowRow(horizontalArrangement = Arrangement.spacedBy(10.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            SecondaryButton(
                text = stringResource(if (locked) R.string.officer_unlock else R.string.officer_lock),
                onClick = { scope.launch { controller.toggleLock() } },
                enabled = !ended && facts.canLock && !controller.offline,
                leading = {
                    if (controller.officerBusy) {
                        CircularProgressIndicator(Modifier.size(16.dp), strokeWidth = 2.dp, color = FsColor.gold)
                    } else {
                        Icon(Icons.Filled.Lock, contentDescription = null, modifier = Modifier.size(16.dp))
                    }
                },
            )
            SecondaryButton(
                text = stringResource(if (controller.copied) R.string.officer_copied else R.string.officer_copy),
                onClick = {
                    val text = controller.copyList(clock)
                    scope.launch { clipboard.setClipEntry(ClipEntry(ClipData.newPlainText(controller.raid.title, text))) }
                },
                enabled = !ended,
                leading = {
                    Icon(if (controller.copied) Icons.Filled.Check else Icons.AutoMirrored.Filled.List, contentDescription = null, modifier = Modifier.size(16.dp))
                },
            )
        }
        Text(
            stringResource(
                when {
                    ended -> R.string.officer_gate
                    locked -> R.string.officer_lockedNote
                    else -> R.string.officer_lockNote
                },
            ),
            style = RaidType.meta,
            color = FsColor.muted,
        )
        controller.officerError?.let { Text(it, style = RaidType.meta, color = FsColor.wipe) }
    }
}

/**
 * Sign-ups (spec F4.2 item 8): groups by role, then Late, Bench, Out, then Not signed up. An
 * officer's rows end in their place button (F4.19); Not signed up become rows with Add.
 */
@Composable
fun Roster(controller: SignupController, facts: RaidFacts) {
    val scope = rememberCoroutineScope()
    val roster = RosterGroups(controller.raid, controller.rosterOverrides.toMap())
    val officer = facts.canMove && facts.isSignable
    val move: (RosterRow, SignupState) -> Unit = { row, state -> scope.launch { controller.move(row, state) } }
    Column(verticalArrangement = Arrangement.spacedBy(14.dp)) {
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp), verticalAlignment = Alignment.Bottom) {
            Text(stringResource(R.string.raid_signups).uppercase(), style = RaidType.h2, color = FsColor.strong, modifier = Modifier.semantics { heading() })
            Text(facts.ratio, style = RaidType.mono, color = FsColor.muted, modifier = Modifier.padding(bottom = 4.dp))
        }
        val res = androidx.compose.ui.platform.LocalResources.current
        gg.foreversixty.app.raids.logic.CompositionRead.of(controller.raid.composition, res)?.let { CompositionBlock(it) }
        roster.groups.forEach { group ->
            Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
                SectionLabel(stringResource(group.kind.label), trailing = "${group.rows.size}")
                val empty = group.kind.emptyLine
                if (group.rows.isEmpty() && empty != null) {
                    Text(stringResource(empty), style = RaidType.meta, color = FsColor.muted, modifier = Modifier.quietCard().padding(14.dp))
                } else {
                    Rows(group.rows, officer, controller, move)
                }
            }
        }
        if (roster.notSigned.isNotEmpty()) {
            Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
                SectionLabel(stringResource(R.string.raid_group_notSigned), trailing = "${roster.notSigned.size}")
                if (officer) {
                    Rows(roster.notSigned, true, controller, move)
                } else {
                    Text(
                        buildAnnotatedString {
                            roster.notSigned.forEachIndexed { index, row ->
                                if (index > 0) append(", ")
                                withStyle(SpanStyle(color = RaidColors.classColor(row.className))) { append(row.name) }
                            }
                        },
                        style = RaidType.meta,
                    )
                }
            }
        }
    }
}

@Composable
private fun Rows(rows: List<RosterRow>, officer: Boolean, controller: SignupController, move: (RosterRow, SignupState) -> Unit) {
    Column(Modifier.quietCard()) {
        rows.forEachIndexed { index, row ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            RosterRowView(row, officer, controller.rowErrors[row.characterKey]) { move(row, it) }
        }
    }
}

/** One roster row (min 48): spec icon 28, the name in class colour, the spec, the note. */
@Composable
fun RosterRowView(row: RosterRow, officer: Boolean, error: String?, move: (SignupState) -> Unit) {
    val className = GameClass.of(row.className)?.let { stringResource(it.label) }.orEmpty()
    val label = listOf(stringResource(R.string.raids_a11y_roster, row.name, row.specName, className), row.note)
        .filter { it.isNotBlank() }.joinToString(". ")
    Column(
        verticalArrangement = Arrangement.spacedBy(4.dp),
        modifier = Modifier
            .fillMaxWidth()
            .background(if (row.own) RaidColors.rowWash else FsColor.raised.copy(alpha = 0f))
            .heightIn(min = 48.dp)
            .padding(horizontal = 14.dp, vertical = 10.dp),
    ) {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
            GameIcon(row.specIcon, 28.dp)
            Column(Modifier.weight(1f).clearAndSetSemantics { contentDescription = label }) {
                NameText(row.name, RaidType.bodyStrong, RaidColors.classColor(row.className))
                if (row.specName.isNotEmpty()) Text(row.specName, style = RaidType.meta, color = FsColor.muted)
                if (row.note.isNotEmpty()) Text(row.note, style = RaidType.meta, color = FsColor.text)
            }
            if (officer) PlaceMenu(row, move)
        }
        error?.let { Text(it, style = RaidType.meta, color = FsColor.wipe) }
    }
}

/** The menu's items for a place (spec F4.20's table): label, new state, destructive. */
fun placeMenuItems(state: SignupState?): List<Triple<Int, SignupState, Boolean>> = when (state) {
    SignupState.In -> listOf(Triple(R.string.officer_move_bench, SignupState.Bench, false), Triple(R.string.officer_move_remove, SignupState.Out, true))
    SignupState.Bench -> listOf(Triple(R.string.officer_move_in, SignupState.In, false), Triple(R.string.officer_move_remove, SignupState.Out, true))
    SignupState.Late -> listOf(Triple(R.string.officer_move_in, SignupState.In, false), Triple(R.string.officer_move_bench, SignupState.Bench, false))
    SignupState.Out, null -> listOf(Triple(R.string.officer_move_addIn, SignupState.In, false), Triple(R.string.officer_move_addBench, SignupState.Bench, false))
}

/** The officer's place button and its move menu (spec F4.20). */
@Composable
fun PlaceMenu(row: RosterRow, move: (SignupState) -> Unit, initiallyOpen: Boolean = false) {
    var open by remember { mutableStateOf(initiallyOpen) }
    val place = row.state?.let { stringResource(it.label) } ?: stringResource(R.string.officer_add)
    val a11y = stringResource(R.string.raids_a11y_place, row.name, place)
    Box {
        Row(
            horizontalArrangement = Arrangement.spacedBy(4.dp),
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier
                .widthIn(min = 56.dp)
                .heightIn(min = 48.dp)
                .background(FsColor.bg.copy(alpha = 0.4f), RoundedCornerShape(FsRadius.control))
                .border(1.dp, if (open) FsColor.gold else FsColor.line, RoundedCornerShape(FsRadius.control))
                .clickable { open = true }
                .padding(horizontal = 12.dp)
                .clearAndSetSemantics { contentDescription = a11y },
        ) {
            Text(place, style = RaidType.segment, color = FsColor.strong)
            Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.strong, modifier = Modifier.size(16.dp))
        }
        DropdownMenu(expanded = open, onDismissRequest = { open = false }, containerColor = RaidColors.sheet) {
            Text(
                stringResource(
                    R.string.officer_menu_head,
                    row.name,
                    row.state?.let { stringResource(it.label) } ?: stringResource(R.string.raid_chip_notSigned),
                    stringResource(roleLabel(row.role)),
                ),
                style = RaidType.meta,
                color = FsColor.muted,
                modifier = Modifier.padding(horizontal = 16.dp, vertical = 8.dp),
            )
            placeMenuItems(row.state).forEach { (label, state, destructive) ->
                DropdownMenuItem(
                    text = { Text(stringResource(label), style = RaidType.body, color = if (destructive) FsColor.wipe else FsColor.strong) },
                    onClick = {
                        open = false
                        move(state)
                    },
                )
            }
        }
    }
}
