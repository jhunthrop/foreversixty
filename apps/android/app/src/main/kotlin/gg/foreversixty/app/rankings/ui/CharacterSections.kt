package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clipToBounds
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.text.LinkAnnotation
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextLinkStyles
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withLink
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.rankings.logic.BestRow
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPageView

/** C5: the Forever Score and its tier as figures, the one `Most to gain` line, the parts under a disclosure. */
@Composable
fun CharacterScore(view: CharacterPageView, actions: CharacterActions, top: Dp) {
    val title = stringResource(R.string.char_score)
    FlowSection(title, top = top) {
        FlowGap(FsSpace.space3)
        val score = view.score
        if (score == null) {
            Text(view.noScoreLine, style = LINE, color = FsColor.text, modifier = Modifier.padding(vertical = LINE_PAD))
            return@FlowSection
        }
        Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space6)) {
            CharacterFigure(stringResource(R.string.char_page_scoreFigure), score.score, modifier = Modifier.weight(1f))
            CharacterFigure(stringResource(R.string.char_page_tier), score.tier, color = ItemQuality(score.tierQuality).text, modifier = Modifier.weight(1f))
        }
        CharacterScoreTrend(score)
        score.mostToGain?.let { most ->
            val how = stringResource(R.string.char_page_howScored)
            Text(
                buildAnnotatedString {
                    append(most)
                    append(" ")
                    withLink(LinkAnnotation.Clickable(HOW_TAG, TextLinkStyles(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold))) { actions.open(CharacterLink.HowScored) }) {
                        append(how)
                    }
                },
                style = LINE,
                color = FsColor.text,
                modifier = Modifier.padding(vertical = LINE_PAD),
            )
        }
        CharacterDisclosure(stringResource(R.string.char_page_parts), ParityRegion.section(title)) {
            score.parts.forEachIndexed { i, (part, sub) -> FlowRow(part, sub, last = i == score.parts.lastIndex) }
        }
    }
}

/** C6, C7: the best kill per encounter (each opens its report) and the builds seen (each opens the planner). */
@Composable
fun CharacterFights(view: CharacterPageView, reportsOn: Boolean, actions: CharacterActions) {
    if (view.best.isNotEmpty()) {
        FlowSection(
            stringResource(R.string.char_page_best),
            action = if (view.allFights > 0) stringResource(R.string.char_page_allFights, view.allFights) else null,
            onAction = { actions.open(CharacterLink.AllFights) },
            top = SECTION_GAP,
            actionColor = FsColor.gold,
        ) {
            FlowGap(FsSpace.space3)
            view.best.forEachIndexed { i, row ->
                FlowRow(
                    row.boss, row.sub, last = i == view.best.lastIndex,
                    trail = { BestTrail(row) },
                    onClick = if (reportsOn) ({ actions.openFight(row.fight) }) else null,
                )
            }
        }
    }
    CharacterRecentFights(view, reportsOn) { actions.openFight(it.fight) }
    if (view.builds.isNotEmpty()) {
        FlowSection(stringResource(R.string.char_builds), top = SECTION_GAP) {
            FlowGap(FsSpace.space3)
            view.builds.forEachIndexed { i, build ->
                FlowRow(build.split, build.sub, last = i == view.builds.lastIndex, onClick = { actions.open(CharacterLink.Planner) })
            }
        }
    }
}

@Composable
private fun BestTrail(row: BestRow) {
    Row(verticalAlignment = Alignment.CenterVertically) {
        Box(Modifier.width(PERCENTILE)) {
            row.percentile?.let { Text(it.toString(), style = FsType.item, color = percentileColor(it)) }
        }
        Spacer(Modifier.width(FsSpace.space3))
        Text(
            buildAnnotatedString {
                append(row.value)
                withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" ${row.unit}") }
            },
            style = FsType.figureSm.copy(fontSize = 18.sp, lineHeight = 26.sp),
            color = FsColor.strong,
        )
    }
}

/** C8: History (the owner's score and parse history, or its lock) and Coach history. */
@Composable
fun CharacterOwner(view: CharacterPageView, actions: CharacterActions) {
    val owner = view.owner ?: return
    FlowSection(stringResource(R.string.char_page_history), top = SECTION_GAP) {
        FlowGap(FsSpace.space3)
        if (owner.history) {
            DisclosureRow(stringResource(R.string.char_page_historyLead)) { actions.open(CharacterLink.History) }
        } else {
            FlowRow(
                stringResource(R.string.char_page_historyLockTitle), stringResource(R.string.char_page_historyLockSub), last = true,
                trail = { FlowLink(stringResource(R.string.char_page_seePlans)) { actions.open(CharacterLink.Plans) } },
            )
        }
    }
    if (owner.coach) {
        FlowSection(stringResource(R.string.char_page_coach), top = SECTION_GAP) {
            FlowGap(FsSpace.space3)
            DisclosureRow(stringResource(R.string.char_page_coachLead)) { actions.open(CharacterLink.Coach) }
        }
    }
}

/**
 * A disclosure (`.disc`): its summary row opens [body] in place. Closed, the body is laid out at no height
 * and out of the semantics tree, as the web's closed `details` keeps its rows, and they count in [region].
 */
@Composable
fun CharacterDisclosure(summary: String, region: String, body: @Composable () -> Unit) {
    var open by remember { mutableStateOf(false) }
    DisclosureRow(summary, open) { open = !open }
    if (open) {
        Column(Modifier.fillMaxWidth()) { body() }
    } else {
        ParityScrollRegion(region) {
            Column(Modifier.fillMaxWidth().height(0.dp).clipToBounds().clearAndSetSemantics {}) {
                Column(Modifier.wrapContentHeight(Alignment.Top, unbounded = true)) { body() }
            }
        }
    }
}

/** The disclosure's 44 row: the turn mark and its words (14/20 semibold); [open] turns the mark down. */
@Composable
private fun DisclosureRow(summary: String, open: Boolean = false, onClick: () -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = DISCLOSURE).clickable(role = Role.Button, onClick = onClick),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(if (open) OPEN_MARK else CLOSED_MARK, style = MARK, color = FsColor.muted)
        Spacer(Modifier.width(FsSpace.space2))
        Text(summary, style = SUMMARY, color = FsColor.text)
    }
}

/** A percentile's colour (`percentileToken`): gold at 100, pink, orange, purple, blue, green, grey under 25. */
internal fun percentileColor(percentile: Int): Color = when {
    percentile >= TOP -> FsColor.parseGold
    percentile >= PINK -> FsColor.parsePink
    percentile >= ORANGE -> FsColor.parseOrange
    percentile >= PURPLE -> FsColor.parsePurple
    percentile >= BLUE -> FsColor.parseBlue
    percentile >= GREEN -> FsColor.parseGreen
    else -> FsColor.parseGreyText
}

private const val TOP = 100
private const val PINK = 99
private const val ORANGE = 95
private const val PURPLE = 75
private const val BLUE = 50
private const val GREEN = 25
private const val HOW_TAG = "how"
private const val CLOSED_MARK = "▸"
private const val OPEN_MARK = "▾"
private val PERCENTILE = 32.dp
private val LINE_PAD = 6.dp
private val DISCLOSURE = 44.dp
private val LINE = FsType.body.copy(lineHeight = 24.sp)
private val SUMMARY = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val MARK = FsType.small.copy(fontSize = 10.sp, lineHeight = 20.sp)
