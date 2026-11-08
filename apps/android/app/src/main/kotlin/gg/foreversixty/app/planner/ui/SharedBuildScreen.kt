package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.spec.DpColumn
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.planner.logic.PlannerPageCopy
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.SharedBuildView

/** What the shared build page draws beyond the phase: the tree, the toggles and what was copied. */
data class SharedPageState(
    val phase: SharedBuildPhase,
    val tree: Int,
    val offlineStamp: String? = null,
    /** Compare with mine shows only when the current character has an export of that class. */
    val canCompare: Boolean = false,
    /** The back button's label: the screen it returns to. */
    val backTo: String = "",
    val differsShown: Boolean = false,
    val copiedForGame: Boolean = false,
    /** The signed-in character's name when the build is its class. */
    val sendTo: String? = null,
)

class SharedActions(
    val back: () -> Unit = {},
    val share: (String) -> Unit = {},
    val editCopy: (SharedBuildView) -> Unit = {},
    val compare: (SharedBuildView) -> Unit = {},
    val onTree: (Int) -> Unit = {},
    val openTalent: (Int) -> Unit = {},
    val retry: () -> Unit = {},
    val sim: (SharedBuildView) -> Unit = {},
    val send: (SharedBuildView) -> Unit = {},
    val copyForGame: (SharedBuildView) -> Unit = {},
    val addon: () -> Unit = {},
    val toggleDiffers: () -> Unit = {},
    val openBis: (SharedBuildView) -> Unit = {},
)

/**
 * A shared build (mock `build-shared`, the web `/b/:id`): the head with its title and views, Edit a
 * copy, the trees read only, In the game, the guide build it is set against and its point order.
 */
@Composable
fun SharedBuildPage(state: SharedPageState, actions: SharedActions) {
    val loaded = (state.phase as? SharedBuildPhase.Loaded)?.view
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        FlowBackBar(state.backTo, actions.back, share = loaded?.let { view -> { actions.share(view.build.id) } })
        Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState()).parityContent().padding(bottom = FsSpace.space10)) {
            DpColumn(Modifier.fillMaxWidth()) { SharedBody(state, actions) }
        }
    }
}

@Composable
private fun SharedBody(state: SharedPageState, actions: SharedActions) {
    when (val phase = state.phase) {
        SharedBuildPhase.Loading -> PlateBlock(top = FsSpace.space4) { PlannerPlateSkeleton() }
        SharedBuildPhase.NotFound -> Message(stringResource(R.string.plan_shared_notFound), null)
        SharedBuildPhase.Offline -> Message(stringResource(R.string.plan_shared_offline), actions.retry)
        SharedBuildPhase.Failed -> Message(stringResource(R.string.plan_shared_failed), actions.retry)
        is SharedBuildPhase.Loaded -> Loaded(phase.view, state, actions)
    }
}

@Composable
private fun Message(text: String, retry: (() -> Unit)?) {
    Column(Modifier.fillMaxWidth().padding(FsSpace.space4), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        Text(text, style = FsType.body, color = FsColor.muted)
        retry?.let { FlowButton(stringResource(R.string.plannerPage_retry), it, small = false) }
    }
}

@Composable
private fun Loaded(view: SharedBuildView, state: SharedPageState, actions: SharedActions) {
    val copy = PlannerPageCopy(LocalResources.current)
    val band = view.band
    state.offlineStamp?.let { OfflineStrip(stringResource(R.string.offline_strip, it)) }
    val spent = view.build.pointOrder.size
    val spec = band?.build?.spec?.name?.takeIf { spent > 0 }
    val facts = listOf(
        stringResource(R.string.plannerPage_sharedMeta),
        copy.meta(view.raceName, view.level, spec, view.className),
        view.totals.joinToString("/"),
        pluralStringResource(R.plurals.plannerPage_views, view.build.views.toInt(), view.build.views.toInt()),
    )
    val score = band?.let { compare -> copy.score(compare)?.let { (label, value) -> PlannerFigure(label, value, copy.scoreBasis(compare.build.band)) } }
    val left = PlannerFigure(stringResource(R.string.plannerPage_pointsLeft), "0", copy.spent(spent, view.level, character = false), gold = true)
    val answer = PlannerAnswer(listOfNotNull(left, score), sim = { actions.sim(view) })
    Box(Modifier.padding(horizontal = FsSpace.space4)) {
        val title = view.build.title?.takeIf { it.isNotBlank() } ?: stringResource(R.string.plannerPage_h1, view.className)
        FlowHead(title, facts, classColor(view.classSlug), crest = view.classSlug) { PlannerAnswerBlock(answer) }
    }
    Column(
        Modifier.padding(horizontal = FsSpace.space4).fillMaxWidth().parityRegion(ParityRegion.NOTES).padding(top = NOTE_TOP, bottom = FsSpace.space1),
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        Text(stringResource(R.string.plannerPage_sharedNote), style = FsType.body.copy(lineHeight = 24.sp), color = FsColor.text)
        FlowButton(stringResource(R.string.plannerPage_editCopy), { actions.editCopy(view) }, small = false)
    }
    PlateBlock(top = FsSpace.space5) {
        PlannerPlate(
            PlateModel(
                view.index, view.build.pointOrder, view.build.treeVersion, state.tree,
                differing = if (state.differsShown) band?.differing.orEmpty() else emptySet(), canSpend = false,
            ),
            status = null,
            PlateActions(actions.onTree, actions.openTalent),
        )
    }
    Sections {
        PlannerInGameSection(InGameModel(state.sendTo, state.copiedForGame, spent > 0), InGameActions({ actions.send(view) }, { actions.copyForGame(view) }, actions.addon))
        band?.let { compare ->
            PlannerBandSection(
                BandSectionModel(
                    stringResource(R.string.plannerPage_guideHeading), copy.line(compare), load = null,
                    canDiffer = compare.differing.isNotEmpty(), differsShown = state.differsShown, gearRange = copy.range(compare.build.band),
                ),
                BandSectionActions(onDiffers = actions.toggleDiffers, onBis = { actions.openBis(view) }),
            )
        }
        PlannerPointOrderSection(view.index, view.build.pointOrder, view.build.treeVersion)
        if (state.canCompare) FlowLink(stringResource(R.string.plannerPage_compareLink), Modifier.padding(top = FsSpace.space8)) { actions.compare(view) }
    }
}

/** The shared note's place (`.notice`): 14 under the head. */
private val NOTE_TOP = 14.dp
