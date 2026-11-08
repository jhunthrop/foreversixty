package gg.foreversixty.app.bis.spec

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.LinkAnnotation
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextLinkStyles
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withLink
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowHead

/**
 * One spec's list (mocks `bis-spec-*`): the head with the set's figure, the Change bar and its panel,
 * the summary with By slot and Where to go, the notes, the plate of cells, What to stack and the guide.
 */
@Composable
fun BisSpecPage(content: BisSpecContent, view: BisSpecView, build: String, panelOpen: Boolean, onPanel: () -> Unit, actions: BisSpecActions) {
    DpColumn(Modifier.fillMaxSize().verticalScroll(rememberScrollState()).parityContent().padding(bottom = FsSpace.space10)) {
        Box(Modifier.padding(horizontal = FsSpace.space4)) {
            FlowHead(content.title, listOf(content.meta), classColor(content.classSlug), crest = content.classSlug, answer = {
                Answer(content.answer, actions)
            })
        }
        Controls(content.controls, panelOpen, onPanel, actions)
        DpColumn(Modifier.padding(horizontal = FsSpace.space4).padding(top = FsSpace.space5)) {
            Summary(content.summary, view, actions)
            if (content.notes.isNotEmpty()) {
                Notes(content.notes)
                Spacer(Modifier.height(FsSpace.space2))
            }
            BisSpecPlate(content, view, build, actions)
            content.weights?.let { BisSpecStack(it, actions.onAddon) }
            Footer(content.guide, actions.onGuide)
        }
    }
}

@Composable
private fun Answer(answer: BisSpecAnswer, actions: BisSpecActions) {
    DpColumn(Modifier.fillMaxWidth(), spacing = FsSpace.space1) {
        Box(Modifier.fillMaxWidth(), contentAlignment = Alignment.TopEnd) {
            Column(horizontalAlignment = Alignment.End, verticalArrangement = Arrangement.spacedBy(2.dp)) {
                Text(answer.label, style = LABEL, color = FsColor.muted, modifier = Modifier.wholeDp())
                answer.value?.let { Text(figure(it, answer.unit), style = FsType.figureBig, color = FsColor.strong, modifier = Modifier.heightIn(min = FIGURE_BOX)) }
            }
        }
        if (answer.facts.isNotEmpty()) {
            Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space5, Alignment.End)) {
                answer.facts.forEach { fact ->
                    Column {
                        Text(fact.label.uppercase(), style = FACT_LABEL, color = FsColor.muted)
                        Text(fact.value, style = FACT_VALUE, color = FsColor.text)
                    }
                }
            }
        }
        val rankingSim = answer.rankingSim
        Text(
            if (rankingSim == null) AnnotatedString(answer.line) else linked(answer.line + " ", rankingSim, actions.onRankingSim),
            style = LINE, color = FsColor.muted, modifier = Modifier.wholeDp(),
        )
        Row(Modifier.fillMaxWidth().heightIn(min = LINKS_HEIGHT), horizontalArrangement = Arrangement.End, verticalAlignment = Alignment.CenterVertically) {
            if (answer.openSim) {
                PageLink(stringResource(R.string.bisSpec_openSim), actions.onOpenSim)
                Text(DOT, style = LINKS, color = FsColor.text)
            }
            PageLink(stringResource(R.string.bisSpec_planner), actions.onPlanner)
        }
    }
}

@Composable
private fun PageLink(text: String, onClick: () -> Unit) {
    Text(text, style = LINKS, color = FsColor.gold, modifier = Modifier.clickable(role = Role.Button, onClick = onClick))
}

/** The bar under the head (`.cbar`): the list's words and Change; the panel holds the choices. */
@Composable
private fun Controls(controls: BisSpecControls, open: Boolean, onPanel: () -> Unit, actions: BisSpecActions) {
    Column(Modifier.fillMaxWidth().background(BAR).parityRegion(ParityRegion.CONTROLS)) {
        HorizontalDivider(color = FsColor.lineSoft)
        Row(
            Modifier
                .fillMaxWidth()
                .clickable(role = Role.Button, onClick = onPanel)
                .padding(horizontal = FsSpace.space4, vertical = BAR_PAD)
                .heightIn(min = BAR_ROW),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Text(controls.bar, style = BAR_TEXT, color = FsColor.strong, modifier = Modifier.weight(1f))
            Text(stringResource(R.string.bisSpec_change), style = CHANGE, color = FsColor.gold)
        }
        ParityScrollRegion(ParityRegion.CONTROLS) { BisSpecPanel(controls, open, actions) }
        HorizontalDivider(color = FsColor.lineSoft)
    }
}

/** The summary (`.bis-sum`): the count with Sign in and the next band, then By slot and Where to go. */
@Composable
private fun Summary(summary: BisSpecSummary, view: BisSpecView, actions: BisSpecActions) {
    Column(
        Modifier.fillMaxWidth().wholeDp().parityRegion(ParityRegion.SUMMARY).padding(vertical = SUMMARY_PAD),
        verticalArrangement = Arrangement.spacedBy(SUMMARY_PAD),
    ) {
        val signIn = stringResource(R.string.bisSpec_signIn)
        val tail = stringResource(R.string.bisSpec_signInTail)
        val text = buildAnnotatedString {
            append(summary.count + " ")
            withLink(LinkAnnotation.Clickable(SIGN_IN, TextLinkStyles(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold))) { actions.onSignIn() }) {
                append(signIn)
            }
            append(" $tail")
            summary.nextBand?.let { append(" ") }
        }
        InlineButtonText(text, summary.nextBand, SUMMARY_TEXT) { summary.nextPick?.let(actions.onPick) }
        val views = listOf(BisSpecView.BySlot to R.string.bisSpec_bySlot, BisSpecView.WhereToGo to R.string.bisSpec_whereToGo)
        BisSpecSegment(views.map { (v, label) -> BisSpecOption(stringResource(label), v == view) }, height = VIEW_HEIGHT, fill = false) { actions.onView(views[it].first) }
    }
}

@Composable
private fun Notes(notes: List<String>) {
    DpColumn(Modifier.fillMaxWidth().parityRegion(ParityRegion.NOTES), spacing = FsSpace.space2) {
        notes.forEach { Text(it, style = FsType.small, color = FsColor.muted) }
    }
}

@Composable
private fun Footer(guide: String, onGuide: () -> Unit) {
    Text(
        linked(stringResource(R.string.bisSpec_guideLead) + " ", guide, onGuide),
        style = FsType.small,
        color = FsColor.muted,
        modifier = Modifier.padding(top = FsSpace.space12, bottom = FsSpace.space2).fillMaxWidth().wholeDp().parityRegion(ParityRegion.FOOTER_NOTE),
    )
}

/** [lead] then [link] in gold, tappable. */
fun linked(lead: String, link: String, onClick: () -> Unit): AnnotatedString = buildAnnotatedString {
    append(lead)
    withLink(LinkAnnotation.Clickable(link, TextLinkStyles(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold))) { onClick() }) { append(link) }
}

private fun figure(value: String, unit: String): AnnotatedString = buildAnnotatedString {
    append(value)
    withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $unit") }
}

private const val DOT = " · "
private const val SIGN_IN = "sign-in"
private val FIGURE_BOX = 49.dp
private val LINKS_HEIGHT = 24.dp
private val BAR_PAD = 6.dp
private val BAR_ROW = 48.dp
private val SUMMARY_PAD = 10.dp
private val VIEW_HEIGHT = 44.dp
private val BAR = FsColor.bg.copy(alpha = 0.94f)
private val LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val FACT_LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 18.sp, fontWeight = FontWeight.Bold, letterSpacing = 0.08.em)
private val FACT_VALUE = FsType.body.copy(fontSize = 16.sp, lineHeight = 24.sp)
private val LINE = FsType.small.copy(lineHeight = 18.sp)
private val LINKS = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 24.sp)
private val BAR_TEXT = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val CHANGE = FsType.bodyStrong.copy(fontSize = 13.sp, lineHeight = 20.sp)
private val SUMMARY_TEXT = FsType.body.copy(lineHeight = 24.sp)
