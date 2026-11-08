package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.ExperimentalFoundationApi
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.combinedClickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.requiredSize
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.CustomAccessibilityAction
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.customActions
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import coil3.compose.AsyncImage
import coil3.request.ImageRequest
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.spec.wholeDp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.items.ItemIconLoader
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.planner.logic.PlannerCells
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.TalentCellModel
import gg.foreversixty.app.talents.TalentCellState
import gg.foreversixty.app.talents.TalentConnector
import gg.foreversixty.app.talents.TalentConnectorShape
import gg.foreversixty.app.talents.TalentEditing
import gg.foreversixty.app.talents.accessibilityLabel

/** What the plate draws: the build's trees, the shown tree, and the cells that differ from the band's build. */
class PlateModel(
    val index: TalentIndex,
    val order: List<Int>,
    val build: String,
    val tree: Int,
    /** Talent ids marked by `See what differs`. */
    val differing: Set<Int> = emptySet(),
    /** False with no point left (or read only): nothing empty is drawn learnable. */
    val canSpend: Boolean = true,
    /** The keyboard's talent: the gold focus ring. */
    val selectedId: Int? = null,
)

/** The plate's own controls: the tabs, a cell (add and remove, or the rank sheet read only), Reset. */
class PlateActions(
    val onTree: (Int) -> Unit = {},
    val open: (Int) -> Unit = {},
    val editing: TalentEditing? = null,
    /** Null on a shared build: it has no Reset. */
    val reset: (() -> Unit)? = null,
)

/**
 * The page's one Plate (the mock's TalentEditor on a phone): the tree tabs, the shown tree on the
 * client's art with its prerequisite lines, then the status line, the tree source and Reset.
 */
@Composable
fun PlannerPlate(model: PlateModel, status: (@Composable () -> Unit)?, actions: PlateActions) {
    PlateFrame {
        TreeTabs(model, actions.onTree)
        model.index.trees.getOrNull(model.tree)?.let { tree -> TreeArea(tree, model, actions) }
        HorizontalDivider(color = FsColor.lineSoft)
        Column(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4, vertical = FsSpace.space3), verticalArrangement = Arrangement.spacedBy(FOOT_GAP)) {
            Box(Modifier.fillMaxWidth().heightIn(min = STATUS_LINE)) { status?.invoke() }
            Text(stringResource(R.string.plannerPage_treeSource, model.build), style = FsType.small, color = FsColor.muted, modifier = Modifier.wholeDp())
            actions.reset?.let { reset ->
                Box(Modifier.fillMaxWidth(), contentAlignment = Alignment.CenterEnd) {
                    Text(
                        stringResource(R.string.plannerPage_reset), style = LINK_SMALL, color = FsColor.gold,
                        modifier = Modifier.parityItem(ParityItemKind.Button).clickable(role = Role.Button, onClick = reset),
                    )
                }
            }
        }
    }
}

/** The plate while the trees load: a title bar and the art's place at the ready trees' height. */
@Composable
fun PlannerPlateSkeleton() {
    PlateFrame(Modifier.heightIn(min = PLATE_RESERVED)) {
        Column(Modifier.fillMaxWidth().padding(vertical = FsSpace.space4), horizontalAlignment = Alignment.CenterHorizontally, verticalArrangement = Arrangement.spacedBy(TREE_GAP)) {
            FlowSkeleton(SKELETON_TITLE, FsSpace.space6)
            FlowSkeleton(ART_SKELETON_WIDTH, ART_SKELETON_HEIGHT, Modifier.parityItem(ParityItemKind.Image))
        }
    }
}

/** The plate when the talent files did not load: what failed, the build, and Retry. */
@Composable
fun PlannerPlateError(build: String, retry: () -> Unit) {
    PlateFrame(Modifier.heightIn(min = PLATE_RESERVED)) {
        Column(Modifier.fillMaxWidth().padding(FsSpace.space5), verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
            Text(stringResource(R.string.plannerPage_loadFailed), style = FsType.body.copy(lineHeight = 24.sp), color = FsColor.strong)
            Text(stringResource(R.string.plannerPage_loadFailedDetail, build), style = FsType.small, color = FsColor.muted)
            FlowButton(stringResource(R.string.plannerPage_retry), retry, small = false)
        }
    }
}

@Composable
private fun PlateFrame(modifier: Modifier = Modifier, content: @Composable ColumnScope.() -> Unit) {
    val shape = RoundedCornerShape(PLATE_RADIUS)
    Column(
        modifier
            .fillMaxWidth()
            .parityRegion(ParityRegion.PLATE)
            .background(FsColor.raised, shape)
            .border(1.dp, FsColor.line, shape)
            .padding(1.dp),
        content = content,
    )
}

/** The three tree tabs (`.tedit-tabs`): the tree's icon and points over its name; the shown one gold. */
@Composable
private fun TreeTabs(model: PlateModel, onTree: (Int) -> Unit) {
    val points = PlannerRules.pointsPerTree(model.index, model.order)
    val shape = RoundedCornerShape(FsRadius.control)
    Row(
        Modifier
            .padding(start = FsSpace.space3, end = FsSpace.space3, top = FsSpace.space3)
            .fillMaxWidth()
            .height(IntrinsicSize.Min)
            .clip(shape)
            .border(1.dp, FsColor.line, shape),
    ) {
        model.index.trees.forEachIndexed { i, tree ->
            if (i > 0) Box(Modifier.width(1.dp).fillMaxHeight().background(FsColor.lineSoft))
            val on = i == model.tree
            val color = if (on) FsColor.gold else FsColor.text
            Column(
                Modifier
                    .weight(1f)
                    .fillMaxHeight()
                    .parityItem(ParityItemKind.Button)
                    .background(if (on) FsColor.selected else Color.Transparent)
                    .drawBehind { if (on) drawRect(FsColor.gold, Offset(0f, size.height - TAB_LINE.toPx()), size.copy(height = TAB_LINE.toPx())) }
                    .clickable(role = Role.Tab) { onTree(i) }
                    .semantics { selected = on }
                    .padding(vertical = FsSpace.space2, horizontal = FsSpace.space1),
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.spacedBy(2.dp),
            ) {
                Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(6.dp)) {
                    Box(Modifier.size(TAB_ICON).parityItem(ParityItemKind.Image).clip(RoundedCornerShape(FsRadius.control))) { GameIconImage(tree.icon, model.build, TAB_ICON) }
                    Text(points.getOrElse(i) { 0 }.toString(), style = TAB_POINTS, color = if (on) FsColor.gold else FsColor.strong)
                }
                Text(tree.name, style = TAB_NAME, color = color, maxLines = 1)
            }
        }
    }
}

/** The tree on its art: four columns of 44 cells, 8 apart, on seven tiers, centred in the plate. */
@Composable
private fun TreeArea(tree: TalentTree, model: PlateModel, actions: PlateActions) {
    val ranks = PlannerRules.ranks(model.order)
    val cells = PlannerCells.cells(tree, ranks, model.canSpend)
    Box(Modifier.fillMaxWidth().padding(top = TREE_GAP, bottom = TREE_BOTTOM), contentAlignment = Alignment.TopCenter) {
        Box(Modifier.size(GRID_WIDTH, GRID_HEIGHT)) {
            TreeArt(tree.background, model.build)
            Box(Modifier.padding(GRID_PAD)) {
                GridLines(PlannerCells.connectors(tree, ranks))
                cells.forEach { cell ->
                    Box(Modifier.offset(PITCH * cell.talent.column, PITCH * cell.talent.tier)) {
                        GridCell(cell, model, actions)
                    }
                }
            }
        }
    }
}

@Composable
private fun TreeArt(background: String, build: String) {
    val context = LocalContext.current
    val art = Modifier.fillMaxSize().parityItem(ParityItemKind.Image).clip(RoundedCornerShape(FsRadius.control)).alpha(ART_ALPHA)
    if (background.isEmpty()) {
        Box(art)
        return
    }
    AsyncImage(
        model = ImageRequest.Builder(context).data("${StaticDataClient.SITE_URL}/data/$build/trees/$background.webp").build(),
        imageLoader = ItemIconLoader.get(context),
        contentDescription = null,
        contentScale = ContentScale.Crop,
        modifier = art,
    )
}

/** The prerequisite lines (`.tedit-g path`): a halo in the page colour, then gold when met, else the line. */
@Composable
private fun GridLines(connectors: List<TalentConnector>) {
    Canvas(Modifier.size(GRID_WIDTH - GRID_PAD * 2, GRID_HEIGHT - GRID_PAD * 2).clearAndSetSemantics {}) {
        val cell = CELL.toPx()
        val pitch = PITCH.toPx()
        val half = cell / 2
        fun at(column: Int, tier: Int, dx: Float, dy: Float) = Offset(column * pitch + dx, tier * pitch + dy)
        connectors.forEach { line ->
            val side = if (line.to.column > line.from.column) 1f else -1f
            val path = Path()
            when (line.shape) {
                TalentConnectorShape.Down -> {
                    path.moveTo(at(line.from.column, line.from.tier, half, cell))
                    path.lineTo(at(line.to.column, line.to.tier, half, 0f))
                }
                TalentConnectorShape.Across -> {
                    path.moveTo(at(line.from.column, line.from.tier, half + side * half, half))
                    path.lineTo(at(line.to.column, line.to.tier, half - side * half, half))
                }
                TalentConnectorShape.Elbow -> {
                    val start = at(line.from.column, line.from.tier, half, cell)
                    val end = at(line.to.column, line.to.tier, half - side * half, half)
                    path.moveTo(start)
                    path.lineTo(Offset(start.x, end.y))
                    path.lineTo(end)
                }
            }
            drawPath(path, FsColor.bg, style = Stroke(HALO.toPx(), cap = StrokeCap.Round))
            drawPath(path, if (line.met) FsColor.gold else LINE_OFF, style = Stroke(LINE.toPx(), cap = StrokeCap.Round))
        }
    }
}

private fun Path.moveTo(point: Offset) = moveTo(point.x, point.y)
private fun Path.lineTo(point: Offset) = lineTo(point.x, point.y)

/**
 * One talent (`.tc2`): its icon in the client's colours (green learnable or partial, gold maxed,
 * grey locked), the rank pill, and the gold dot when it differs from the band's build.
 */
@OptIn(ExperimentalFoundationApi::class)
@Composable
private fun GridCell(cell: TalentCellModel, model: PlateModel, actions: PlateActions) {
    val res = LocalResources.current
    val label = cell.accessibilityLabel(res)
    val hint = stringResource(R.string.talents_a11yHint)
    val id = cell.talent.id
    val editing = actions.editing
    val touch = if (editing == null) {
        Modifier
            .semantics(mergeDescendants = true) {
                contentDescription = label
                role = Role.Button
                onClick(hint) { actions.open(id); true }
            }
            .clickable { actions.open(id) }
    } else {
        Modifier
            .semantics(mergeDescendants = true) {
                contentDescription = label
                role = Role.Button
                onClick(editing.addLabel) { editing.add(id); true }
                customActions = listOf(
                    CustomAccessibilityAction(editing.addLabel) { editing.add(id); true },
                    CustomAccessibilityAction(editing.removeLabel) { editing.remove(id); true },
                )
            }
            // The planner plays its own haptics (a long press only when a point really comes off).
            .combinedClickable(hapticFeedbackEnabled = false, onLongClick = { editing.remove(id) }, onClick = { editing.add(id) })
    }
    val ring = ringColor(cell.state)
    val shape = RoundedCornerShape(CELL_RADIUS)
    Box(Modifier.size(CELL).parityItem(ParityItemKind.Button).then(touch)) {
        if (model.selectedId == id) {
            Box(Modifier.requiredSize(CELL + FOCUS_GROWTH).align(Alignment.Center).border(2.dp, FsColor.gold, RoundedCornerShape(CELL_RADIUS + 4.dp)))
        }
        Box(Modifier.size(CELL).clip(shape).background(CELL_FILL).parityItem(ParityItemKind.Image)) {
            GameIconImage(cell.talent.icon, model.build, CELL, fallback = cell.talent.name.take(2), dimmed = cell.state == TalentCellState.Locked)
        }
        Box(Modifier.size(CELL).border(2.dp, ring, shape))
        Text(
            cell.badge,
            style = RANK,
            color = rankColor(cell.state),
            softWrap = false,
            modifier = Modifier
                .align(Alignment.BottomEnd)
                .offset(5.dp, 5.dp)
                .background(FsColor.bg, CircleShape)
                .border(1.dp, ring, CircleShape)
                .padding(horizontal = FsSpace.space1)
                .clearAndSetSemantics {},
        )
        if (id in model.differing) {
            Box(Modifier.offset((-4).dp, (-4).dp).size(DOT).border(2.dp, FsColor.bg, CircleShape).padding(1.dp).background(FsColor.gold, CircleShape))
        }
    }
}

private fun ringColor(state: TalentCellState): Color = when (state) {
    TalentCellState.Maxed -> FsColor.gold
    TalentCellState.Filled -> FsColor.kill
    TalentCellState.Available -> LEARNABLE
    TalentCellState.Open, TalentCellState.Locked -> FsColor.line
}

private fun rankColor(state: TalentCellState): Color = when (state) {
    TalentCellState.Maxed -> FsColor.gold
    TalentCellState.Filled, TalentCellState.Available -> FsColor.kill
    TalentCellState.Open, TalentCellState.Locked -> FsColor.muted
}

/** The game's TalentFrame green for a talent you can learn now (the mock's `--ring` on `.tc2.avail`). */
private val LEARNABLE = Color(0xFF4FAE5C)
/** An unmet prerequisite line (`.tedit-g path.off`). */
private val LINE_OFF = Color(0xFF3A4256)
/** A cell's tile under its icon (`.tc2`). */
private val CELL_FILL = Color(0xFF10141D)
private val CELL = 44.dp
private val PITCH = 52.dp
private val GRID_PAD = 6.dp
private val GRID_WIDTH = 212.dp
private val GRID_HEIGHT = 368.dp
private val CELL_RADIUS = 5.dp
private val FOCUS_GROWTH = 8.dp
private val DOT = 9.dp
private val HALO = 5.dp
private val LINE = 2.dp
private val TAB_ICON = 28.dp
private val TAB_LINE = 2.dp
private val TREE_GAP = 14.dp
private val TREE_BOTTOM = 18.dp
private val FOOT_GAP = 6.dp
private val STATUS_LINE = 20.dp
private val PLATE_RADIUS = 6.dp
private val PLATE_RESERVED = 577.dp
private val SKELETON_TITLE = 180.dp
private val ART_SKELETON_WIDTH = 216.dp
private val ART_SKELETON_HEIGHT = 384.dp
private const val ART_ALPHA = 0.5f
private val RANK = FsType.small.copy(fontSize = 12.sp, lineHeight = 14.sp, fontWeight = FontWeight.Bold)
private val TAB_POINTS = FsType.bodyStrong.copy(lineHeight = 20.sp, fontWeight = FontWeight.Bold)
private val TAB_NAME = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
private val LINK_SMALL = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
