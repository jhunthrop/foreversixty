package gg.foreversixty.app.home.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowButtonRow
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowItemRow
import gg.foreversixty.app.activity.ActivityRowView
import gg.foreversixty.app.activity.activityLead
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPanel
import gg.foreversixty.app.flows.FlowPageState
import androidx.compose.foundation.rememberScrollState
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowTallyLine
import gg.foreversixty.app.raids.logic.SignupState

/**
 * The signed-in Home (spec 2026-11-08 4.1) in the round-5 look: the hero with [topBar] and the head,
 * then the blocks of the moment on the ground.
 */
@Composable
fun HomeFlowScreen(
    view: HomeView,
    actions: HomeActions,
    page: FlowPageState = FlowPageState(rememberScrollState()),
    topBar: @Composable () -> Unit = {},
) {
    when (view) {
        is HomeView.Loading -> HomePage(topBar, page, { Head(view.head) }) {
            FlowBlock(ParityRegion.section("Loading"), top = 0.dp) {
                // The raid card, then the rows under it, at their final size: nothing shifts when they fill.
                SKELETONS.forEachIndexed { index, height ->
                    if (index > 0) FlowGap(FsSpace.space3)
                    FlowSkeleton(SKELETON_WIDTH, height)
                }
            }
        }
        is HomeView.Failed -> HomePage(topBar, page, { Head(view.head) }) {
            FlowBlock(ParityRegion.section("Error"), top = 0.dp) {
                FlowPanel {
                    FlowRow(stringResource(R.string.home_error_title), stringResource(R.string.home_error_sub), last = true, trail = {
                        FlowButton(stringResource(R.string.common_tryAgain), actions.onRetry, Modifier.heightIn(min = FsTarget.min))
                    })
                }
            }
        }
        HomeView.NoCharacter -> HomePage(topBar, page, {
            HomeHeroHead(stringResource(R.string.home_new_title), listOf(stringResource(R.string.home_new_line)), className = null)
        }) {
            BrandButton(stringResource(R.string.home_new_check), BrandButtonKind.Primary, actions.onCheckAgain, height = FsTarget.min)
        }
        is HomeView.Blocks -> HomePage(topBar, page, { Head(view.head) }) {
            view.blocks.forEach { Block(it, view.iconBuild, actions) }
        }
    }
}

@Composable
private fun Head(head: HomeHead) {
    HomeHeroHead(head.name, head.facts, head.className.orEmpty())
}

@Composable
private fun Block(block: HomeBlock, build: String?, actions: HomeActions) {
    when (block) {
        is HomeBlock.Upgrades -> Upgrades(block, actions)
        is HomeBlock.NextRaidCard -> FlowBlock(ParityRegion.COUNTDOWNS, top = 0.dp) { HomeRaidCard(block.raid, actions) }
        is HomeBlock.NextRaidLine -> FlowBlock(ParityRegion.COUNTDOWNS, top = 0.dp) { RaidLine(block.raid, build, actions) }
        is HomeBlock.RaidLive -> FlowBlock(ParityRegion.COUNTDOWNS, top = 0.dp) { HomeRaidLive(block.live, actions) }
        is HomeBlock.LastNight -> LastNight(block, build, actions)
        is HomeBlock.Activity -> NewForYou(block, build, actions)
        is HomeBlock.FindGuild -> FlowBlock(ParityRegion.section("Raid with a guild"), top = 0.dp) {
            FlowPanel {
                FlowNote(block.line, Modifier.padding(top = FsSpace.space2), color = FsColor.text)
                FlowButton(stringResource(R.string.home_findGuild), actions.onFindGuild, Modifier.padding(vertical = FsSpace.space3).fillMaxWidth(), small = false)
            }
        }
        is HomeBlock.Notice -> FlowBlock(ParityRegion.section("Notice"), top = 0.dp) {
            FlowPanel { FlowRow(block.text, last = true, onClick = block.href?.let { href -> { actions.onNotice(href) } }) }
        }
    }
}

@Composable
private fun Upgrades(block: HomeBlock.Upgrades, actions: HomeActions) {
    FlowSection(
        stringResource(R.string.home_upgrades_next),
        meta = block.meta,
        action = stringResource(R.string.home_upgrades_gear),
        onAction = actions.onGear,
        top = 0.dp,
        panel = PanelFrame.Plain,
        stamp = block.synced,
        stampWarn = block.syncedStale,
    ) {
        block.rows.forEachIndexed { index, row ->
            if (index > 0) FlowRule()
            FlowItemRow(row.look, block.build, row.source, row.gain, row.unit, inset = 0.dp) { actions.onItem(row) }
        }
        block.bestRun?.let { run ->
            FlowRule()
            FlowItemRow(run.look, block.build, run.source, run.gain, run.unit, inset = 0.dp) { actions.onGear() }
        }
    }
}

@Composable
internal fun RaidOfficerItems(raidId: Long, officer: OfficerItems, actions: HomeActions) {
    officer.unanswered?.let { line ->
        FlowRow(line, officer.short, last = officer.applicants == null, trail = {
            FlowButton(stringResource(R.string.home_officer_preview), { actions.onPreviewMention(raidId) }, Modifier.heightIn(min = FsTarget.min))
        })
    }
    officer.applicants?.let { FlowRow(it, officer.applicantLine, last = true, onClick = actions.onApplicants) }
}

/** In, Late, Bench, Out sharing a row, the current answer selected (F1). */
@Composable
fun RaidAnswerButtons(selected: SignupState?, modifier: Modifier = Modifier, onAnswer: (SignupState) -> Unit) {
    val labels = SignupState.entries.map { stringResource(it.label) }
    FlowSegmented(labels, selected?.let { stringResource(it.label) }, modifier) { onAnswer(SignupState.entries[it]) }
}

/** How late? with the 15, 30 and 60 min chips, shown after Late (F1). */
@Composable
fun LateChips(onLate: (Int) -> Unit) {
    FlowNote(stringResource(R.string.push_late_question), Modifier.padding(top = FsSpace.space2))
    val chips = LATE_MINUTES.map { stringResource(R.string.push_late_minutes, it) }
    FlowButtonRow(chips, null) { onLate(LATE_MINUTES[it]) }
}

/** Any other day: one line with the raid and its answers (Late with its minutes), as the raid-day card's (F1). */
@Composable
private fun RaidLine(raid: RaidCardModel, build: String?, actions: HomeActions) {
    FlowPanel {
        FlowRow(raid.title, raid.sub, FlowLead.Icon(raid.icon, build), last = true, onClick = { actions.onRaid(raid.raidId) })
        // Who is coming in one line, the short role in wipe red (S6, the Raid tab's tally model).
        FlowTallyLine(raid.tally)
        RaidAnswers(raid, actions, Modifier.padding(top = FsSpace.space1, bottom = FsSpace.space2), LateChipsStyle.Asked)
    }
}

@Composable
private fun LastNight(block: HomeBlock.LastNight, build: String?, actions: HomeActions) {
    FlowSection(stringResource(R.string.home_lastNight), top = 0.dp, panel = PanelFrame.Plain) {
        block.rows.forEachIndexed { index, item ->
            val row = item.row
            FlowRow(row.title, item.line, activityLead(row, build), last = index == block.rows.lastIndex && !block.simAgain, onClick = { actions.onActivityRow(row) })
        }
        if (block.simAgain) {
            FlowRow(stringResource(R.string.home_lastNight_sim), stringResource(R.string.home_lastNight_simSub), slim = true, last = true, onClick = actions.onSimAgain)
        }
    }
}

@Composable
private fun NewForYou(block: HomeBlock.Activity, build: String?, actions: HomeActions) {
    val link = if (block.unread > 0) stringResource(R.string.home_allActivity_count, block.unread) else stringResource(R.string.home_allActivity)
    FlowSection(stringResource(R.string.home_newForYou), action = link, onAction = actions.onAllActivity, top = 0.dp, panel = PanelFrame.None) {
        block.rows.forEachIndexed { index, item ->
            val row = item.row
            ActivityRowView(row, item.line, build, last = index == block.rows.lastIndex, onClick = { actions.onActivityRow(row) }, meta = item.age)
        }
    }
}

private val SKELETONS = listOf(365.dp, 56.dp, 56.dp, 56.dp, 56.dp)
private val SKELETON_WIDTH = 384.dp
private val LATE_MINUTES = listOf(15, 30, 60)
