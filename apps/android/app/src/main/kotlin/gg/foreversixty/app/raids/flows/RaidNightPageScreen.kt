package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.flows.FlowCountdown
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowHeadIcon
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.home.flows.LateChips
import gg.foreversixty.app.home.flows.RaidAnswerButtons
import gg.foreversixty.app.raids.logic.SignupState

/** What the raid page on raid night hands back. */
data class RaidNightPageActions(
    val answer: (SignupState) -> Unit = {},
    val late: (Int) -> Unit = {},
    val roster: () -> Unit = {},
    val where: (ReadyItem) -> Unit = {},
)

/**
 * The raid page on raid night (spec F2, mocks raid-night-live, -signup): the countdown with your
 * answer, Your sign-up with the current answer selected, who is coming by role, and Before you pull,
 * the private checklist from the addon's last export.
 */
@Composable
fun RaidNightPageScreen(content: RaidNightPageContent, build: String?, actions: RaidNightPageActions, error: String? = null) {
    FlowPage {
        // The raid page's head sits a further 16 in (the mock's `.ph` on the raid page).
        Column(Modifier.padding(horizontal = FsSpace.space4)) {
            FlowHead(content.title, content.facts, lead = { FlowHeadIcon(content.icon, build) })
        }
        FlowCountdown(content.eyebrow, content.clock, content.clockLabel, live = content.live) {
            Column(verticalArrangement = Arrangement.spacedBy(TAG_GAP)) {
                FlowNote(content.aside)
                Tag(content.tag, if (content.answered) FsColor.kill else FsColor.muted)
            }
        }
        FlowSection(stringResource(R.string.raid_yourSignup)) {
            RaidAnswerButtons(content.selected, Modifier.padding(top = FsSpace.space3), actions.answer)
            if (content.lateOpen) LateChips(actions.late)
            error?.let { FlowNote(it, Modifier.padding(top = FsSpace.space2), color = FsColor.wipe) }
        }
        FlowSection(stringResource(R.string.raids_whoIsComing), action = stringResource(R.string.raid_night_roster), onAction = actions.roster) {
            content.roles.forEachIndexed { index, (title, names) -> FlowRow(title, names, slim = true, last = index == content.roles.lastIndex) }
        }
        if (content.ready.isNotEmpty()) {
            FlowSection(stringResource(R.string.raid_night_beforePull)) { ReadyList(content.ready, content.readyStamp, actions.where) }
        }
    }
}

/** A small tag (`.tag`): the word in its tone over a faint fill of it. */
@Composable
private fun Tag(text: String, tone: androidx.compose.ui.graphics.Color) {
    Box(Modifier.background(tone.copy(alpha = TAG_FILL), RoundedCornerShape(TAG_RADIUS)).padding(horizontal = FsSpace.space2, vertical = TAG_PAD)) {
        Text(text, style = TAG, color = tone)
    }
}

private val TAG_GAP = 6.dp
private val TAG_RADIUS = 4.dp
private val TAG_PAD = 2.dp
private const val TAG_FILL = 0.14f
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 18.sp, fontWeight = FontWeight.SemiBold)
