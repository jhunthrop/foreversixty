package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNotice
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPageView
import gg.foreversixty.app.rankings.logic.HeadFact

/** C2, C3: the crest, the class-coloured name, the fact line, and the answer when the band's list has upgrades. */
@Composable
fun CharacterHead(view: CharacterPageView, actions: CharacterActions) {
    val answer = view.answer
    FlowHead(
        title = view.name,
        facts = emptyList(),
        titleColor = classColor(view.className),
        crest = view.className,
        meta = { view.facts.forEach { Fact(it, actions.open) } },
        answer = answer?.let {
            {
                Row(Modifier.align(Alignment.End), horizontalArrangement = Arrangement.spacedBy(FsSpace.space6), verticalAlignment = Alignment.Bottom) {
                    it.score?.let { score -> CharacterFigure(stringResource(R.string.char_score), score, size = FigureSize.Small) }
                    CharacterFigure(stringResource(R.string.char_page_answer), it.total, unit = it.unit, size = FigureSize.Big)
                }
                Row(Modifier.align(Alignment.End).padding(top = FsSpace.space1), horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
                    FlowLink(stringResource(R.string.char_page_openSim)) { actions.open(CharacterLink.Sim) }
                    FlowLink(stringResource(R.string.char_page_openPlanner)) { actions.open(CharacterLink.Planner) }
                }
            }
        },
    )
}

@Composable
private fun Fact(fact: HeadFact, open: (CharacterLink) -> Unit) {
    when (fact) {
        is HeadFact.Words -> Text(fact.text, style = FsType.small, color = FsColor.muted)
        is HeadFact.Link -> Text(
            fact.text, style = FsType.small.copy(fontWeight = FontWeight.SemiBold), color = FsColor.gold,
            modifier = Modifier.clickable(role = Role.Button) { open(fact.link) },
        )
        is HeadFact.Tag -> Box(Modifier.height(TAG).border(1.dp, FsColor.gold, RoundedCornerShape(FsRadius.pill)).padding(horizontal = FsSpace.space2), contentAlignment = Alignment.Center) {
            Text(fact.text, style = TAG_TEXT, color = FsColor.gold)
        }
    }
}

/** The owner's summary (C11, C16 to C21): the one Next line, then the standing links. */
@Composable
fun CharacterSummary(view: CharacterPageView, actions: CharacterActions) {
    val next = view.next
    if (next == null && view.standing.isEmpty()) return
    Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.SUMMARY)) {
        if (next != null) {
            FlowNotice(next.text, next.action, { next.link?.let(actions.open) }, Modifier.padding(top = NOTICE_DROP), color = FsColor.gold)
        }
        if (view.standing.isNotEmpty()) {
            FlowRow(
                Modifier.fillMaxWidth().padding(bottom = FsSpace.space1),
                horizontalArrangement = Arrangement.spacedBy(FsSpace.space5),
                verticalArrangement = Arrangement.spacedBy(FsSpace.space1),
            ) {
                view.standing.forEach { line ->
                    FlowLink(stringResource(R.string.char_page_standingLink, line.text)) { line.link?.let(actions.open) }
                }
            }
        }
    }
}

enum class FigureSize(val value: TextStyle, val box: Dp) {
    Small(FsType.figureSm.copy(fontSize = 18.sp, lineHeight = 24.sp), 24.dp),
    Normal(FsType.figure, 32.dp),
    Big(FsType.figureBig, 49.dp),
}

/** A Figure (`.fig`): the small caps label over the value, its unit small beside it, in [color]. */
@Composable
fun CharacterFigure(label: String, value: String, modifier: Modifier = Modifier, unit: String? = null, size: FigureSize = FigureSize.Normal, color: Color = FsColor.strong) {
    Column(modifier, verticalArrangement = Arrangement.spacedBy(FIGURE_GAP), horizontalAlignment = Alignment.End.takeIf { size != FigureSize.Normal } ?: Alignment.Start) {
        Text(label.uppercase(), style = LABEL, color = FsColor.muted)
        Box(Modifier.height(size.box), contentAlignment = Alignment.BottomStart) {
            Text(
                buildAnnotatedString {
                    append(value)
                    unit?.takeIf { it.isNotEmpty() }?.let { withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $it") } }
                },
                style = size.value,
                color = color,
            )
        }
    }
}

private val TAG = 20.dp
private val NOTICE_DROP = 5.dp
private val FIGURE_GAP = 2.dp
private val TAG_TEXT = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.SemiBold)
internal val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold, letterSpacing = 0.12.em)
