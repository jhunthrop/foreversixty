package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clipToBounds
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.spec.wholeDp
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.raids.logic.GameClass

/** A head figure (`.pl-fig`): the label, the big value, and its line under it. */
data class PlannerFigure(val label: String, val value: String, val line: String, val gold: Boolean = false)

/** The head's answer: Points left, the set score when the app has it, and `Sim this build ›` when there is one. */
data class PlannerAnswer(val figures: List<PlannerFigure>, val sim: (() -> Unit)?)

/** The answer under the title (`.pl-ans`): the figures side by side, 24 apart, then the sim link. */
@Composable
fun ColumnScope.PlannerAnswerBlock(answer: PlannerAnswer) {
    Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space6)) {
        answer.figures.forEach { Figure(it, Modifier.weight(1f)) }
        repeat(FIGURES - answer.figures.size) { Spacer(Modifier.weight(1f)) }
    }
    answer.sim?.let { FlowLink(stringResource(R.string.plannerPage_simThisBuild), Modifier.padding(top = SIM_GAP).height(LINK_LINE), it) }
}

/** The answer's place while the trees load: two figure-sized skeletons. */
@Composable
fun PlannerAnswerSkeleton() {
    Row(Modifier.fillMaxWidth().heightIn(min = ANSWER_LOADING), horizontalArrangement = Arrangement.spacedBy(FsSpace.space6)) {
        FlowSkeleton(SKELETON_FIGURE_WIDTH, SKELETON_FIGURE_HEIGHT)
        FlowSkeleton(SKELETON_FIGURE_WIDTH, SKELETON_FIGURE_HEIGHT)
    }
}

@Composable
private fun Figure(figure: PlannerFigure, modifier: Modifier) {
    Column(modifier, verticalArrangement = Arrangement.spacedBy(2.dp)) {
        Text(figure.label, style = LABEL, color = FsColor.muted, modifier = Modifier.height(LABEL_LINE))
        Text(figure.value, style = FsType.figureBig, color = if (figure.gold) FsColor.gold else FsColor.strong, modifier = Modifier.height(FIGURE_VALUE))
        Text(figure.line, style = FsType.small, color = FsColor.muted, modifier = Modifier.heightIn(min = FIGURE_LINE).wholeDp())
    }
}

/** What the Change panel picks: the class, the race and the level. */
class PlannerChoices(
    val classSlug: String,
    val race: String,
    val level: Int?,
    val onClass: (String) -> Unit = {},
    val onRace: () -> Unit = {},
    val onLevel: () -> Unit = {},
)

/**
 * The bar under the head (`.cbar`): the build's class and race and Change; the panel holds the
 * class picker, the race and the level. Closed, the panel keeps its place at no height and is
 * hidden from accessibility (the BiS spec page's closed Change panel).
 */
@Composable
fun PlannerControls(bar: String, choices: PlannerChoices, open: Boolean, onToggle: () -> Unit) {
    Column(Modifier.fillMaxWidth().background(BAR).parityRegion(ParityRegion.CONTROLS)) {
        HorizontalDivider(color = FsColor.lineSoft)
        Row(
            Modifier.fillMaxWidth().heightIn(min = BAR_ROW).clickable(role = Role.Button, onClick = onToggle).padding(horizontal = FsSpace.space4),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Text(bar, style = BAR_TEXT, color = FsColor.strong, modifier = Modifier.weight(1f))
            Text(stringResource(R.string.plannerPage_change), style = CHANGE, color = FsColor.gold)
        }
        ParityScrollRegion(ParityRegion.CONTROLS) { ChangePanel(choices, open) }
        HorizontalDivider(color = FsColor.lineSoft)
    }
}

@Composable
private fun ChangePanel(choices: PlannerChoices, open: Boolean) {
    val closed = if (open) Modifier else Modifier.height(0.dp).clipToBounds().clearAndSetSemantics { }
    Box(Modifier.fillMaxWidth().then(closed)) {
        Column(
            Modifier.fillMaxWidth().wrapContentHeight(Alignment.Top, unbounded = true).padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space3),
            verticalArrangement = Arrangement.spacedBy(FsSpace.space3),
        ) {
            Labelled(stringResource(R.string.plannerPage_classLabel)) { ClassPicker(choices) }
            Labelled(stringResource(R.string.plannerPage_raceLabel)) {
                PanelField(Modifier.parityItem(ParityItemKind.Field).clickable(role = Role.Button, onClick = choices.onRace)) {
                    Text(choices.race, style = FIELD_TEXT, color = FsColor.strong)
                }
            }
            choices.level?.let { level ->
                Labelled(stringResource(R.string.plannerPage_levelLabel)) {
                    PanelField(Modifier.clickable(role = Role.Button, onClick = choices.onLevel)) {
                        Text(stringResource(R.string.plan_level, level), style = FIELD_TEXT, color = FsColor.strong)
                    }
                }
            }
        }
    }
}

/** The nine classes, three to a row, each with its crest (`.seg.clsp`). */
@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun ClassPicker(choices: PlannerChoices) {
    val shape = RoundedCornerShape(FsRadius.control)
    FlowRow(Modifier.fillMaxWidth().border(1.dp, FsColor.line, shape).padding(1.dp), maxItemsInEachRow = PER_ROW) {
        GameClass.entries.forEach { cls ->
            val on = cls.slug == choices.classSlug
            Row(
                Modifier
                    .weight(1f)
                    .height(OPTION_HEIGHT)
                    .parityItem(ParityItemKind.Segment)
                    .background(if (on) FsColor.selected else Color.Transparent)
                    .clickable(role = Role.RadioButton) { choices.onClass(cls.slug) }
                    .semantics { selected = on }
                    .padding(horizontal = FsSpace.space2),
                horizontalArrangement = Arrangement.spacedBy(FsSpace.space2, Alignment.CenterHorizontally),
                verticalAlignment = Alignment.CenterVertically,
            ) {
                ClassCrest(cls.slug, CREST, Modifier.parityItem(ParityItemKind.Image))
                Text(stringResource(cls.label), style = OPTION_TEXT, color = if (on) FsColor.gold else FsColor.text, maxLines = 1)
            }
        }
    }
}

@Composable
private fun Labelled(label: String, body: @Composable () -> Unit) {
    Column {
        Text(label, style = LABEL, color = FsColor.muted, modifier = Modifier.padding(bottom = FsSpace.space1))
        body()
    }
}

@Composable
private fun PanelField(modifier: Modifier, content: @Composable () -> Unit) {
    val shape = RoundedCornerShape(FsRadius.control)
    Box(
        modifier.fillMaxWidth().heightIn(min = OPTION_HEIGHT).border(1.dp, FsColor.line, shape).padding(horizontal = FsSpace.space3),
        contentAlignment = Alignment.CenterStart,
    ) { content() }
}

/** The note under the bar (`.pl-notes .notice`): where the build on screen came from, one line. */
@Composable
fun PlannerNotes(text: String) {
    Text(
        text,
        style = FsType.body.copy(lineHeight = 24.sp),
        color = FsColor.text,
        modifier = Modifier.padding(horizontal = FsSpace.space4).fillMaxWidth().parityRegion(ParityRegion.NOTES).padding(vertical = FsSpace.space1),
    )
}

private const val FIGURES = 2
private const val PER_ROW = 3
private val SIM_GAP = 6.dp
private val FIGURE_LINE = 40.dp
private val LABEL_LINE = 16.dp
private val FIGURE_VALUE = 40.dp
private val LINK_LINE = 24.dp
private val ANSWER_LOADING = 130.dp
private val SKELETON_FIGURE_WIDTH = 150.dp
private val SKELETON_FIGURE_HEIGHT = 64.dp
private val BAR_ROW = 48.dp
private val OPTION_HEIGHT = 44.dp
private val CREST = 22.dp
private val BAR = FsColor.bg.copy(alpha = 0.94f)
private val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val BAR_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val CHANGE = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 20.sp)
private val OPTION_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
private val FIELD_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 16.sp)
