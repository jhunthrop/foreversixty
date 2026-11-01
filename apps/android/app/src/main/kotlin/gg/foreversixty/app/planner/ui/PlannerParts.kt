package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.planner.logic.PlannerCells
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.TalentBuildKind
import gg.foreversixty.app.talents.TalentCompareBar
import gg.foreversixty.app.talents.TalentEditing
import gg.foreversixty.app.talents.TalentTreeView

/** From font scale 1.5 the chips and the strip stack, as F6's bar does (spec §9). */
const val PLANNER_LARGE_TEXT = 1.5f

@Composable
fun isLargeText(): Boolean = LocalDensity.current.fontScale >= PLANNER_LARGE_TEXT

/** The tree tabs and F6's tree, one component for the editor and the viewer (spec F13.1, F13.5). */
@Composable
fun PlannerTree(
    index: TalentIndex,
    order: List<Int>,
    build: String,
    tree: Int,
    onTree: (Int) -> Unit,
    modifier: Modifier = Modifier,
    selectedId: Int? = null,
    open: (Int) -> Unit = {},
    editing: TalentEditing? = null,
) {
    val ranks = gg.foreversixty.app.planner.logic.PlannerRules.ranks(order)
    val shown = index.trees.getOrNull(tree) ?: return
    Column(modifier) {
        TalentCompareBar(false, PlannerCells.tabs(index, order), TalentBuildKind.Yours, tree, {}, onTree, isLargeText())
        TalentTreeView(
            PlannerCells.cells(shown, ranks),
            PlannerCells.connectors(shown, ranks),
            build,
            selectedId,
            open,
            Modifier.padding(top = 16.dp),
            editing,
        )
    }
}

/** A planner chip (`Orc ▾`, `Level 60 ▾`): 36 visible inside a 48 target. */
@Composable
fun PlannerChip(label: String, description: String, onClick: () -> Unit) {
    val shape = RoundedCornerShape(FsRadius.pill)
    Box(
        Modifier
            .heightIn(min = 48.dp)
            .clickable(role = Role.Button, onClick = onClick)
            .semantics { contentDescription = description },
        contentAlignment = Alignment.Center,
    ) {
        Row(
            Modifier
                .heightIn(min = 36.dp)
                .background(FsColor.raised, shape)
                .border(1.dp, FsColor.line, shape)
                .padding(start = 14.dp, end = 10.dp),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(4.dp),
        ) {
            Text(label, style = FsType.bodyStrong, color = FsColor.strong)
            Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp))
        }
    }
}

/** The strip's and the level sheet's round `−` and `+` (48). */
@Composable
fun StepButton(icon: Painter, description: String, enabled: Boolean, onClick: () -> Unit, size: Dp = 48.dp) {
    Box(
        Modifier
            .size(size)
            .border(1.dp, FsColor.line, CircleShape)
            .clickable(enabled = enabled, role = Role.Button, onClick = onClick)
            .semantics { contentDescription = description }
            .alpha(if (enabled) 1f else DISABLED_ALPHA),
        contentAlignment = Alignment.Center,
    ) {
        Icon(icon, contentDescription = null, tint = FsColor.strong, modifier = Modifier.size(22.dp))
    }
}

/** `Orc Shaman 30 / 3 / 18` with the totals in mono (spec F13.5 head, F13.6 rows). */
fun headLine(text: String, totals: String): AnnotatedString = buildAnnotatedString {
    append(text)
    val at = text.lastIndexOf(totals)
    if (totals.isNotEmpty() && at >= 0) addStyle(SpanStyle(fontFamily = FsType.mono.fontFamily), at, at + totals.length)
}

/** The tree a file opens on: the most points (F6's rule). */
fun openingTree(file: TalentsFile, order: List<Int>): Int =
    gg.foreversixty.app.talents.TalentRules.openingTree(gg.foreversixty.app.planner.logic.PlannerRules.pointsPerTree(TalentIndex(file), order))

private const val DISABLED_ALPHA = 0.4f
