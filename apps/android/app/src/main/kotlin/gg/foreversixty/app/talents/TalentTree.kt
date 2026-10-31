package gg.foreversixty.app.talents

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.interaction.collectIsPressedAsState
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxWithConstraints
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.requiredSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.clickable
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.scale
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.model.Talent

/**
 * The tree's fixed geometry (spec F6 §5.2): 48 dp cells, a 70 dp tier pitch, a column pitch of
 * at most 96 that fits a 360 dp screen, and a region always reserved at seven tiers. iOS: TalentGeometry.
 */
object TalentGeometry {
    val CELL = 48.dp
    val BORDER = 2.dp
    val TIER_PITCH = 70.dp
    private val MAX_COLUMN_PITCH = 96.dp
    private const val COLUMNS = 4
    private const val RESERVED_TIERS = 7
    private val MARGIN = 16.dp
    private val BADGE_HANG = 6.dp
    val ARROWHEAD = 7.dp
    val LINE = 2.dp
    val HALO = 5.dp

    val height: Dp get() = TIER_PITCH * (RESERVED_TIERS - 1) + CELL + BADGE_HANG + 2.dp

    fun columnPitch(width: Dp): Dp = minOf(MAX_COLUMN_PITCH, ((width - MARGIN * 2 - CELL) / (COLUMNS - 1)).value.toInt().dp)

    /** The top-left corner of a cell, the grid centred in [width]. */
    fun origin(cell: TalentCell, width: Dp): Pair<Dp, Dp> {
        val pitch = columnPitch(width)
        val left = (width - (CELL + pitch * (COLUMNS - 1))) / 2
        return Pair(left + pitch * cell.column, TIER_PITCH * cell.tier)
    }
}

/** One tree: the prerequisite lines under the cells (spec F6 §5.2). */
@Composable
fun TalentTreeView(
    cells: List<TalentCellModel>,
    connectors: List<TalentConnector>,
    build: String,
    selectedId: Int?,
    open: (Int) -> Unit,
    modifier: Modifier = Modifier,
) {
    BoxWithConstraints(modifier.fillMaxWidth().height(TalentGeometry.height)) {
        val width = maxWidth
        TalentLines(connectors, width)
        cells.forEach { cell ->
            val (x, y) = TalentGeometry.origin(TalentCell(cell.talent.tier, cell.talent.column), width)
            Box(Modifier.offset(x, y)) {
                TalentCellView(cell, build, selectedId == cell.talent.id) { open(cell.talent.id) }
            }
        }
    }
}

/** The prerequisite lines (`connectors.ts` `connectorPath`), edge to edge, each ending in an arrowhead. */
@Composable
private fun TalentLines(connectors: List<TalentConnector>, width: Dp) {
    val density = LocalDensity.current
    Canvas(Modifier.fillMaxSize().clearAndSetSemantics {}) {
        fun point(cell: TalentCell, dx: Float, dy: Float): Offset {
            val (x, y) = TalentGeometry.origin(cell, width)
            return Offset(x.toPx() + dx, y.toPx() + dy)
        }
        val cell = TalentGeometry.CELL.toPx()
        val half = cell / 2
        val head = TalentGeometry.ARROWHEAD.toPx()
        connectors.forEach { line ->
            val side = if (line.to.column > line.from.column) 1f else -1f
            val path = Path()
            val (tip, direction) = when (line.shape) {
                TalentConnectorShape.Down -> {
                    path.moveTo(point(line.from, half, cell))
                    val tip = point(line.to, half, 0f)
                    path.lineTo(tip)
                    tip to Offset(0f, 1f)
                }
                TalentConnectorShape.Across -> {
                    path.moveTo(point(line.from, half + side * half, half))
                    val tip = point(line.to, half - side * half, half)
                    path.lineTo(tip)
                    tip to Offset(side, 0f)
                }
                TalentConnectorShape.Elbow -> {
                    val start = point(line.from, half, cell)
                    val tip = point(line.to, half - side * half, half)
                    path.moveTo(start)
                    path.lineTo(Offset(start.x, tip.y))
                    path.lineTo(tip)
                    tip to Offset(side, 0f)
                }
            }
            val arrow = arrowhead(tip, direction, head)
            val color = if (line.met) FsColor.gold else FsColor.line
            drawPath(path, FsColor.bg, style = Stroke(TalentGeometry.HALO.toPx(), cap = StrokeCap.Round))
            drawPath(arrow, FsColor.bg, style = Stroke(with(density) { 3.dp.toPx() }))
            drawPath(path, color, style = Stroke(TalentGeometry.LINE.toPx(), cap = StrokeCap.Round))
            drawPath(arrow, color)
        }
    }
}

private fun Path.moveTo(point: Offset) = moveTo(point.x, point.y)
private fun Path.lineTo(point: Offset) = lineTo(point.x, point.y)

private fun arrowhead(tip: Offset, direction: Offset, size: Float): Path {
    val back = tip - direction * size
    val normal = Offset(-direction.y, direction.x) * (size / 2)
    return Path().apply {
        moveTo(tip)
        lineTo(back + normal)
        lineTo(back - normal)
        close()
    }
}

/**
 * One talent: the icon on its tile in the state border, the rank badge, and on Your build the
 * dashed ring and the signed chip (spec F6 §5.2).
 */
@Composable
fun TalentCellView(cell: TalentCellModel, build: String, isSelected: Boolean, open: () -> Unit) {
    val res = LocalResources.current
    val interaction = remember { MutableInteractionSource() }
    val pressed by interaction.collectIsPressedAsState()
    val deltaColor = if ((cell.delta ?: 0) > 0) FsColor.wipe else FsColor.gold
    val label = cell.accessibilityLabel(res)
    val hint = stringResource(R.string.talents_a11yHint)
    Box(
        Modifier
            .size(TalentGeometry.CELL)
            .semantics(mergeDescendants = true) {
                contentDescription = label
                role = Role.Button
                onClick(hint) { open(); true }
            }
            .clickable(interaction, indication = null, onClick = open)
            .scale(if (pressed) PRESSED_SCALE else 1f),
    ) {
        if (isSelected) {
            Box(Modifier.requiredSize(TalentGeometry.CELL + 18.dp).align(Alignment.Center).border(2.dp, FsColor.gold, RoundedCornerShape(FsRadius.control + 9.dp)))
        }
        if (cell.delta != null) {
            Box(Modifier.requiredSize(TalentGeometry.CELL + 6.dp).align(Alignment.Center).dashedBorder(deltaColor, 12.dp))
        }
        TalentIconTile(cell.talent, cell.state, build)
        if (pressed) Box(Modifier.size(TalentGeometry.CELL).background(GameColors.pressedWash, RoundedCornerShape(FsRadius.control)))
        TalentBadge(
            cell.badge,
            TalentColors.border(cell.state),
            TalentColors.badgeText(cell.state),
            Modifier.align(Alignment.BottomEnd).offset(6.dp, 6.dp),
        )
        cell.chip?.let { TalentBadge(it, deltaColor, deltaColor, Modifier.align(Alignment.TopStart).offset((-8).dp, (-8).dp)) }
    }
}

/** A 2 dp dashed outline with the given corner radius. */
fun Modifier.dashedBorder(color: Color, radius: Dp, width: Dp = 2.dp): Modifier = drawBehind {
    val stroke = width.toPx()
    drawRoundRect(
        color = color,
        topLeft = Offset(stroke / 2, stroke / 2),
        size = size.copy(width = size.width - stroke, height = size.height - stroke),
        cornerRadius = CornerRadius(radius.toPx()),
        style = Stroke(stroke, pathEffect = PathEffect.dashPathEffect(floatArrayOf(4.dp.toPx(), 3.dp.toPx()))),
    )
}

/** The small mono capsule used for the rank badge and the signed chip; its text caps at 1.3x. */
@Composable
fun TalentBadge(text: String, border: Color, textColor: Color, modifier: Modifier = Modifier) {
    val density = LocalDensity.current
    val capped = Density(density.density, density.fontScale.coerceAtMost(BADGE_SCALE_CEILING))
    androidx.compose.runtime.CompositionLocalProvider(LocalDensity provides capped) {
        Text(
            text,
            style = TextStyle(fontFamily = FontFamily.Monospace, fontWeight = FontWeight.SemiBold, fontSize = 11.sp),
            color = textColor,
            softWrap = false,
            modifier = modifier
                .background(FsColor.bg, RoundedCornerShape(6.dp))
                .border(1.dp, border, RoundedCornerShape(6.dp))
                .padding(horizontal = 5.dp, vertical = 1.dp)
                .clearAndSetSemantics {},
        )
    }
}

/**
 * The 48 dp tile: the icon at 44 on `#131824` inside a 2 dp state border; a locked face at 40%;
 * a missing icon keeps the tile with the name's first two letters (spec F6 §5.2).
 */
@Composable
fun TalentIconTile(talent: Talent, state: TalentCellState, build: String, size: Dp = TalentGeometry.CELL) {
    val inner = size - TalentGeometry.BORDER * 2
    Box(
        Modifier
            .size(size)
            .background(FsColor.cardTop, RoundedCornerShape(FsRadius.control))
            .border(TalentGeometry.BORDER, TalentColors.border(state), RoundedCornerShape(FsRadius.control))
            .clearAndSetSemantics {},
        contentAlignment = Alignment.Center,
    ) {
        Box(Modifier.size(inner).clip(RoundedCornerShape(FsRadius.pill)).alpha(if (state == TalentCellState.Locked) LOCKED_ALPHA else 1f)) {
            GameIconImage(talent.icon, build, inner, fallback = talent.name.take(2))
        }
    }
}

/** The state colours (spec F6 §5.2's table). */
object TalentColors {
    fun border(state: TalentCellState): Color = when (state) {
        TalentCellState.Maxed, TalentCellState.Filled -> FsColor.kill
        TalentCellState.Open, TalentCellState.Locked -> FsColor.line
    }

    fun badgeText(state: TalentCellState): Color = when (state) {
        TalentCellState.Maxed -> FsColor.kill
        TalentCellState.Filled -> FsColor.text
        TalentCellState.Open, TalentCellState.Locked -> FsColor.muted
    }
}

private const val PRESSED_SCALE = 0.96f
private const val LOCKED_ALPHA = 0.4f
private const val BADGE_SCALE_CEILING = 1.3f
