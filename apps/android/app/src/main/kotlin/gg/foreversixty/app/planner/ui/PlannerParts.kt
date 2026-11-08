package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Icon
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.TalentRules

/** From font scale 1.5 Your builds' rows stack (spec §9). */
const val PLANNER_LARGE_TEXT = 1.5f

@Composable
fun isLargeText(): Boolean = LocalDensity.current.fontScale >= PLANNER_LARGE_TEXT

/**
 * The level sheet's round `−` and `+` (48). [dimmed] draws it at 40% while it stays tappable; a
 * disabled one is dimmed and inert.
 */
@Composable
fun StepButton(icon: Painter, description: String, enabled: Boolean, onClick: () -> Unit, size: Dp = 48.dp, dimmed: Boolean = !enabled) {
    Box(
        Modifier
            .size(size)
            .border(1.dp, FsColor.line, CircleShape)
            .clickable(enabled = enabled, role = Role.Button, onClick = onClick)
            .semantics { contentDescription = description }
            .alpha(if (dimmed) DISABLED_ALPHA else 1f),
        contentAlignment = Alignment.Center,
    ) {
        Icon(icon, contentDescription = null, tint = FsColor.strong, modifier = Modifier.size(22.dp))
    }
}

/** The tree a file opens on: the most points (F6's rule). */
fun openingTree(file: TalentsFile, order: List<Int>): Int =
    TalentRules.openingTree(PlannerRules.pointsPerTree(TalentIndex(file), order))

private const val DISABLED_ALPHA = 0.4f
