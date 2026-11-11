package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.flows.FlowButton
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.flows.FLOW_CHIP_SMALL
import gg.foreversixty.app.flows.FlowChip
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowSegmented
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.raids.ui.StatusLineText

/**
 * A raid card's answers (spec F1, F4.3, R16), one component for Home's card and line and the Raid
 * tab's head: In, Late, Bench and Out as one segmented pill, with Late the minutes under it (the
 * chosen one selected), then the controller's status line (`Saving…`, `Saved.`, the error with Try
 * again, the offline or locked reason). Offline or locked the answers dim and read as disabled.
 */
@Composable
fun RaidAnswers(raid: RaidCardModel, actions: HomeActions, modifier: Modifier = Modifier, late: LateChipsStyle = LateChipsStyle.Inline) {
    val control = raid.control
    Column(modifier) {
        val labels = SignupState.entries.map { stringResource(it.label) }
        FlowSegmented(labels, raid.answer?.let { stringResource(it.label) }, enabled = control.enabled) {
            actions.onAnswer(raid.raidId, SignupState.entries[it])
        }
        if (raid.lateOpen) LateMinuteChips(raid.lateMinutes, control.enabled, late) { actions.onLate(raid.raidId, it) }
        control.status?.let { status ->
            Column(Modifier.padding(top = FsSpace.space2)) { StatusLineText(status, reserves = false) }
        }
        if (control.canRetry) {
            FlowButton(stringResource(R.string.common_tryAgain), { actions.onRetryAnswer(raid.raidId) }, Modifier.heightIn(min = FsTarget.min))
        }
    }
}

/**
 * How Late's minutes sit under the answers, both under `How late?` (a2): [Inline] on the Raid tab's
 * head (48 dp chips, mock m-android-signup), [Asked] on Home's card (the small chip, 48 to the touch,
 * as the Home mock draws it).
 */
enum class LateChipsStyle(val asks: Boolean, val chipHeight: Dp, val gap: Dp) {
    Inline(asks = true, chipHeight = FsTarget.min, gap = FsSpace.space2),
    Asked(asks = true, chipHeight = FLOW_CHIP_SMALL, gap = 0.dp),
}

/** 15, 30 and 60 min, the chosen one teal (F1); they wrap at large text rather than clip their minutes. */
@Composable
private fun LateMinuteChips(selected: Int?, enabled: Boolean, style: LateChipsStyle, onLate: (Int) -> Unit) {
    val label = stringResource(R.string.raid_tab_howLate)
    if (style.asks) FlowNote(stringResource(R.string.push_late_question), Modifier.padding(top = FsSpace.space2))
    FlowRow(
        Modifier.padding(top = style.gap).semantics { contentDescription = label },
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
        verticalArrangement = Arrangement.spacedBy(FsSpace.space2),
    ) {
        LATE_MINUTES.forEach { minutes ->
            val text = stringResource(R.string.push_late_minutes, minutes)
            FlowChip(text, minutes == selected, { onLate(minutes) }, minHeight = style.chipHeight, enabled = enabled)
        }
    }
}

/** Late's minutes, as the notification's chips (RaidPushCategory.LATE_MINUTES). */
private val LATE_MINUTES = RaidPushCategory.LATE_MINUTES
