package gg.foreversixty.app.home.flows

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowButtonRow
import gg.foreversixty.app.flows.FlowCountdown
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowItemRow
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowTally
import gg.foreversixty.app.flows.FlowWideButton
import gg.foreversixty.app.raids.logic.SignupState
import androidx.compose.ui.unit.dp

/** The signed-in Home (spec 2026-11-08 4.1): one head, then the blocks of the moment. */
@Composable
fun HomeFlowScreen(view: HomeView, actions: HomeActions) {
    FlowPage {
        when (view) {
            is HomeView.Loading -> {
                Head(view.head)
                FlowBlock(ParityRegion.section("Loading")) {
                    repeat(SKELETONS) { index ->
                        if (index > 0) FlowGap(FsSpace.space3)
                        FlowSkeleton(SKELETON_WIDTH, SKELETON_HEIGHT)
                    }
                }
            }
            is HomeView.Failed -> {
                Head(view.head)
                FlowBlock(ParityRegion.section("Error")) {
                    FlowRow(stringResource(R.string.home_error_title), stringResource(R.string.home_error_sub), last = true, trail = {
                        FlowButton(stringResource(R.string.common_tryAgain), actions.onRetry)
                    })
                }
            }
            HomeView.NoCharacter -> {
                FlowHead(stringResource(R.string.home_new_title), listOf(stringResource(R.string.home_new_line)))
                FlowWideButton(stringResource(R.string.home_new_check), primary = true, onClick = actions.onCheckAgain)
            }
            is HomeView.Blocks -> {
                Head(view.head)
                view.blocks.forEach { Block(it, view.iconBuild, actions) }
            }
        }
    }
}

@Composable
private fun Head(head: HomeHead) {
    FlowHead(head.name, head.facts, classColor(head.className), crest = head.className.orEmpty())
}

@Composable
private fun Block(block: HomeBlock, build: String?, actions: HomeActions) {
    when (block) {
        is HomeBlock.Upgrades -> Upgrades(block, actions)
        is HomeBlock.NextRaidCard -> FlowBlock(ParityRegion.COUNTDOWNS) { RaidCard(block.raid, actions) }
        is HomeBlock.NextRaidLine -> FlowBlock(ParityRegion.COUNTDOWNS) { RaidLine(block.raid, build, actions) }
        is HomeBlock.RaidLive -> FlowBlock(ParityRegion.COUNTDOWNS) { RaidLive(block.live, actions) }
        is HomeBlock.LastNight -> LastNight(block, build, actions)
        is HomeBlock.Activity -> NewForYou(block, build, actions)
        is HomeBlock.FindGuild -> FlowBlock(ParityRegion.section("Raid with a guild")) {
            FlowNote(block.line)
            FlowWideButton(stringResource(R.string.home_findGuild), primary = false, onClick = actions.onFindGuild)
        }
        is HomeBlock.Notice -> FlowBlock(ParityRegion.section("Notice")) {
            FlowRow(block.text, last = true, onClick = block.href?.let { href -> { actions.onNotice(href) } })
        }
    }
}

@Composable
private fun Upgrades(block: HomeBlock.Upgrades, actions: HomeActions) {
    FlowSection(stringResource(R.string.home_upgrades_next), action = stringResource(R.string.home_upgrades_gear), onAction = actions.onGear) {
        block.rows.forEach { row ->
            FlowItemRow(row.look, block.build, row.source, row.gain, row.unit) { actions.onItem(row) }
        }
        block.bestRun?.let { FlowRow(it, slim = true, last = true, onClick = actions.onGear) }
    }
}

/** The next raid card (F1): the countdown, who you answer as, the four answers, Late's minutes, who is coming. */
@Composable
fun RaidCard(raid: RaidCardModel, actions: HomeActions) {
    FlowCountdown(raid.eyebrow, raid.clock, raid.clockLabel) { FlowNote(raid.aside) }
    FlowNote(raid.me, Modifier.padding(top = FsSpace.space2))
    Answers(raid, Modifier.padding(top = FsSpace.space3), actions)
    if (raid.lateOpen) {
        FlowNote(stringResource(R.string.push_late_question), Modifier.padding(top = FsSpace.space2))
        val chips = LATE_MINUTES.map { stringResource(R.string.push_late_minutes, it) }
        FlowButtonRow(chips, null) { actions.onLate(raid.raidId, LATE_MINUTES[it]) }
    }
    FlowTally(raid.tally, Modifier.padding(top = FsSpace.space4))
    raid.ready?.let {
        Column(Modifier.padding(top = FsSpace.space2)) { FlowRow(it.title, it.sub, slim = true, last = true, onClick = { actions.onReady(raid.raidId) }) }
    }
    raid.officer?.let { Column(Modifier.padding(top = FsSpace.space2)) { Officer(raid.raidId, it, actions) } }
}

@Composable
private fun Officer(raidId: Long, officer: OfficerItems, actions: HomeActions) {
    officer.unanswered?.let { line ->
        FlowRow(line, officer.short, last = officer.applicants == null, trail = {
            FlowButton(stringResource(R.string.home_officer_preview), { actions.onPreviewMention(raidId) })
        })
    }
    officer.applicants?.let { FlowRow(it, officer.applicantLine, last = true, onClick = actions.onApplicants) }
}

@Composable
private fun Answers(raid: RaidCardModel, modifier: Modifier, actions: HomeActions) {
    val labels = SignupState.entries.map { stringResource(it.label) }
    Column(modifier) {
        FlowButtonRow(labels, raid.answer?.let { stringResource(it.label) }) { actions.onAnswer(raid.raidId, SignupState.entries[it]) }
    }
}

/** Any other day: one line with the raid and its answers, the current one selected (F1). */
@Composable
private fun RaidLine(raid: RaidCardModel, build: String?, actions: HomeActions) {
    FlowRow(raid.title, raid.sub, FlowLead.Icon(raid.icon, build), last = true, onClick = { actions.onRaid(raid.raidId) })
    Answers(raid, Modifier.padding(top = FsSpace.space3), actions)
}

@Composable
private fun RaidLive(live: RaidLiveModel, actions: HomeActions) {
    FlowCountdown(live.eyebrow, emptyList(), "", live = true) {
        Column {
            FlowNote(live.line)
            Text(
                stringResource(R.string.home_live_open),
                style = FsType.bodyStrong,
                color = FsColor.gold,
                modifier = Modifier.padding(top = FsSpace.space2).clickable(role = Role.Button) { actions.onRaidNight(live.raidId) },
            )
        }
    }
}

@Composable
private fun LastNight(block: HomeBlock.LastNight, build: String?, actions: HomeActions) {
    FlowSection(stringResource(R.string.home_lastNight)) {
        block.rows.forEachIndexed { index, item ->
            val row = item.row
            FlowRow(row.title, item.line, FlowLead.Icon(row.icon, build), last = index == block.rows.lastIndex && !block.simAgain, onClick = { actions.onActivityRow(row) })
        }
        if (block.simAgain) {
            FlowRow(stringResource(R.string.home_lastNight_sim), stringResource(R.string.home_lastNight_simSub), slim = true, last = true, onClick = actions.onSimAgain)
        }
    }
}

@Composable
private fun NewForYou(block: HomeBlock.Activity, build: String?, actions: HomeActions) {
    val link = if (block.unread > 0) stringResource(R.string.home_allActivity_count, block.unread) else stringResource(R.string.home_allActivity)
    FlowSection(stringResource(R.string.home_newForYou), action = link, onAction = actions.onAllActivity) {
        block.rows.forEachIndexed { index, item ->
            val row = item.row
            FlowRow(row.title, item.line, FlowLead.Icon(row.icon, build), unread = !row.read, last = index == block.rows.lastIndex, onClick = { actions.onActivityRow(row) })
        }
    }
}

private const val SKELETONS = 5
private val SKELETON_WIDTH = 361.dp
private val SKELETON_HEIGHT = 56.dp
private val LATE_MINUTES = listOf(15, 30, 60)
