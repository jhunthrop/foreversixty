package gg.foreversixty.app.home.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FrameCorners
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.panel
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowTally
import gg.foreversixty.app.flows.FlowTallyLine
import gg.foreversixty.app.flows.lineBox

/**
 * The raid card on raid day (spec F1, round 5): the lead panel, ornate on the cave. The zone, the guild
 * and the time, the countdown, who the answer is for (E7), the four answers as one segmented pill with
 * Late's minutes under them, who is coming, what is not ready, and an officer's open items.
 */
@Composable
fun HomeRaidCard(raid: RaidCardModel, actions: HomeActions) {
    LeadPanel {
        Text(raid.zone, style = HOME_CARD_TITLE, color = FsColor.cream, modifier = Modifier.semantics { heading() })
        Text(raid.place, style = FsType.small, color = FsColor.text)
        Clock(raid.clock, raid.clockLabel, Modifier.padding(top = FsSpace.space2))
        raid.me?.let { Text(it, style = FsType.small, color = FsColor.text, modifier = Modifier.padding(top = FsSpace.space3)) }
        RaidAnswers(raid, actions, Modifier.padding(top = FsSpace.space2), LateChipsStyle.Asked)
        // Row 2: the meters against a target, else the plain counts (`Healers 6`), never nothing.
        val tallyTop = Modifier.padding(top = TALLY_TOP)
        if (raid.tally.any { it.want != null }) FlowTally(raid.tally, tallyTop) else FlowTallyLine(raid.tally, tallyTop)
        raid.ready?.let { ready ->
            HorizontalDivider(Modifier.padding(top = TALLY_TOP), color = FsColor.lineWarm)
            FlowRow(ready.title, ready.sub, slim = true, last = true, onClick = { actions.onReady(raid.raidId) })
        }
        raid.officer?.let { Column(Modifier.padding(top = FsSpace.space1)) { RaidOfficerItems(raid.raidId, it, actions) } }
    }
}

/** The raid while it runs: the green dot, the guild and zone, the pull and your answer, Open raid night. */
@Composable
fun HomeRaidLive(live: RaidLiveModel, actions: HomeActions) {
    LeadPanel {
        Row(verticalAlignment = Alignment.CenterVertically) {
            Box(Modifier.size(LIVE_DOT).background(FsColor.kill, CircleShape))
            Spacer(Modifier.width(FsSpace.space2))
            Text(live.eyebrow, style = FsType.bodyStrong, color = FsColor.strong)
        }
        FlowNote(live.line, Modifier.padding(top = FsSpace.space2), color = FsColor.text)
        Text(
            stringResource(R.string.home_live_open),
            style = FsType.bodyStrong,
            color = FsColor.gold,
            modifier = Modifier.padding(top = FsSpace.space2).clickable(role = Role.Button) { actions.onRaidNight(live.raidId) },
        )
    }
}

/** The lead panel: the ornate frame over the cave, 18 inside (16 at the foot), the corners last. */
@Composable
private fun LeadPanel(content: @Composable ColumnScope.() -> Unit) {
    Box(Modifier.fillMaxWidth().panel(PanelFrame.Ornate, cave = true)) {
        Column(Modifier.fillMaxWidth().padding(start = CARD_PAD, end = CARD_PAD, top = CARD_PAD, bottom = FsSpace.space4), content = content)
        FrameCorners(Modifier.matchParentSize())
    }
}

/** `02 hours  48 min`, read aloud as one timer. */
@Composable
private fun Clock(parts: List<ClockPart>, label: String, modifier: Modifier) {
    if (parts.isEmpty()) return
    Row(
        modifier.semantics(mergeDescendants = true) {}.clearAndSetSemantics { contentDescription = label },
        horizontalArrangement = Arrangement.spacedBy(CLOCK_GAP),
    ) {
        parts.forEach { part ->
            Column {
                Text(part.value, style = CLOCK, color = FsColor.strong, modifier = Modifier.lineBox(CLOCK.lineHeight))
                Text(part.unit, style = UNIT, color = FsColor.muted)
            }
        }
    }
}

private val CARD_PAD = 18.dp
private val TALLY_TOP = 14.dp
private val CLOCK_GAP = 18.dp
private val LIVE_DOT = 8.dp
private val CLOCK = FsType.countdown
private val UNIT = FsType.figureUnit
