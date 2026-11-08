package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowFigures
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection

/** What the applicant's page hands back. */
data class ApplicantActions(
    val accept: () -> Unit = {},
    val decline: () -> Unit = {},
    val copy: (String) -> Unit = {},
)

/**
 * The applicant (spec F3, mock raid-applicant): the Forever Score, sim and item level, raid readiness,
 * guild history and their note, with Accept into guild and Decline anchored under the page; after
 * Accept, the `/ginvite` line with Copy (mock raid-applicant-accepted).
 */
@Composable
fun ApplicantScreen(content: ApplicantContent, actions: ApplicantActions, busy: Boolean = false, error: String? = null) {
    FlowPage {
        FlowHead(content.name, content.facts, classColor(content.className), crest = content.className)
        val accepted = content.accepted
        if (accepted != null) {
            Accepted(accepted, actions.copy)
        } else {
            Review(content)
            Decision(actions, busy)
        }
        error?.let { FlowNote(it, Modifier.padding(top = FsSpace.space3), color = FsColor.wipe) }
    }
}

@Composable
private fun ColumnScope.Review(content: ApplicantContent) {
    if (content.figures.isNotEmpty()) FlowFigures(content.figures)
    if (content.ready.isNotEmpty()) Lines(stringResource(R.string.raid_applicant_ready), content.ready, slim = true)
    if (content.history.isNotEmpty()) Lines(stringResource(R.string.raid_applicant_history), content.history, slim = false)
    content.note?.let { note ->
        FlowSection(stringResource(R.string.raid_applicant_note)) {
            Text(note, style = FsType.prose, color = FsColor.text, modifier = Modifier.padding(top = FsSpace.space2))
        }
    }
}

@Composable
private fun Lines(title: String, lines: List<ApplicantLine>, slim: Boolean) {
    FlowSection(title) {
        lines.forEachIndexed { index, line -> FlowRow(line.title, line.sub, slim = slim, last = index == lines.lastIndex) }
    }
}

/** Accept into guild and Decline, side by side over a rule after the note (`.sticky`), the role line under them. */
@Composable
private fun Decision(actions: ApplicantActions, busy: Boolean) {
    Column(Modifier.padding(top = FsSpace.space6).fillMaxWidth()) {
        HorizontalDivider(color = FsColor.lineSoft)
        Row(Modifier.padding(top = FsSpace.space3).fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
            FlowButton(stringResource(R.string.raid_applicant_accept), { if (!busy) actions.accept() }, Modifier.weight(1f), primary = true, small = false)
            FlowButton(stringResource(R.string.raid_applicant_decline), { if (!busy) actions.decline() }, Modifier.weight(1f), small = false)
        }
        FlowNote(stringResource(R.string.raid_applicant_role), Modifier.padding(top = FsSpace.space2))
    }
}

/** `Accepted into Olympus XXVII. Invite Mirelle in game:`, the command with Copy, what happens next. */
@Composable
private fun Accepted(accepted: AcceptedContent, copy: (String) -> Unit) {
    Text(accepted.line, style = FsType.lead, color = FsColor.text, modifier = Modifier.padding(top = FsSpace.space3))
    val shape = RoundedCornerShape(FsRadius.control)
    Row(
        Modifier
            .padding(top = FsSpace.space2)
            .fillMaxWidth()
            .border(1.dp, FsColor.line, shape)
            .padding(horizontal = FsSpace.space3, vertical = CODE_PAD),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
    ) {
        Text(accepted.command, style = CODE, color = FsColor.strong, modifier = Modifier.weight(1f))
        FlowButton(stringResource(R.string.raid_applicant_copy), { copy(accepted.command) })
    }
    FlowNote(accepted.after, Modifier.padding(top = FsSpace.space3))
}

private val CODE_PAD = 10.dp
private val CODE = androidx.compose.ui.text.TextStyle(fontFamily = androidx.compose.ui.text.font.FontFamily.Monospace, fontWeight = androidx.compose.ui.text.font.FontWeight.Medium, fontSize = 14.sp, lineHeight = 20.sp)
