package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.focusable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.input.key.Key
import androidx.compose.ui.input.key.KeyEvent
import androidx.compose.ui.input.key.KeyEventType
import androidx.compose.ui.input.key.key
import androidx.compose.ui.input.key.onKeyEvent
import androidx.compose.ui.input.key.type
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.spec.DpColumn
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.PlannerBand
import gg.foreversixty.app.planner.logic.PlannerCells
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.PlannerEditor
import gg.foreversixty.app.planner.logic.PlannerPageCopy
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.ShareOutcome
import gg.foreversixty.app.planner.logic.TreeDirection
import gg.foreversixty.app.talents.TalentEditing

/** What the editor page draws beyond the model's phase: the tree, the panel, the fields and what was copied. */
data class EditorPageState(
    val phase: EditorPhase,
    val build: String,
    val tree: Int,
    val request: EditorRequest,
    /** The talent the hardware keyboard has reached: the gold focus ring. */
    val keyboardTalent: Int? = null,
    val backTo: String = "",
    val panelOpen: Boolean = false,
    val differsShown: Boolean = false,
    val title: String = "",
    val importCode: String = "",
    val importUnreadable: Boolean = false,
    val copiedForGame: Boolean = false,
    val copiedText: Boolean = false,
    /** The signed-in character's name when the build is its class: `Send this build to Felsnap`. */
    val sendTo: String? = null,
    val signedIn: Boolean = false,
    /** The menu's character opened the planner: its level frames the spent line. */
    val forCharacter: Boolean = false,
    /** `Your talents against the guide ›`, when the Talents compare can open. */
    val canOpenTalents: Boolean = false,
)

/** Every action on the editor page; the defaults let the parity tests draw it inert. */
class EditorActions(
    val back: () -> Unit = {},
    val undo: () -> Unit = {},
    val add: (Int) -> Unit = {},
    val remove: (Int) -> Unit = {},
    val onTree: (Int) -> Unit = {},
    val togglePanel: () -> Unit = {},
    val setClass: (String) -> Unit = {},
    val openRace: () -> Unit = {},
    val openLevel: () -> Unit = {},
    val reset: () -> Unit = {},
    val share: () -> Unit = {},
    val retry: () -> Unit = {},
    /** An arrow key: the talent it reaches is the status line's. */
    val move: (TreeDirection) -> Unit = {},
    /** The status line's talent: the rank sheet with the full text. */
    val openTalent: (Int) -> Unit = {},
    val loadBand: () -> Unit = {},
    val toggleDiffers: () -> Unit = {},
    val openBis: () -> Unit = {},
    val sim: () -> Unit = {},
    val send: () -> Unit = {},
    val copyForGame: () -> Unit = {},
    val addon: () -> Unit = {},
    val onTitle: (String) -> Unit = {},
    val copyText: () -> Unit = {},
    val builds: () -> Unit = {},
    val signIn: () -> Unit = {},
    val onImportCode: (String) -> Unit = {},
    val import: () -> Unit = {},
    val talents: () -> Unit = {},
)

/**
 * The talent planner (mocks `planner-*`, the web `/planner` on a phone): the head with Points left,
 * the Change bar, the note, the trees, the band build, In the game, Point order, Share and Import.
 * Stateless, so the parity tests draw every state.
 */
@Composable
fun PlannerEditorPage(state: EditorPageState, actions: EditorActions, modifier: Modifier = Modifier) {
    val view = (state.phase as? EditorPhase.Ready)?.view
    Column(modifier.fillMaxSize().background(FsColor.bg)) {
        FlowBackBar(state.backTo, actions.back, share = view?.let { actions.share }) {
            if (view != null) UndoButton(view.state.canUndo, actions.undo)
        }
        Column(
            Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .parityContent()
                .then(if (view != null) Modifier.onKeyEvent { treeKey(it, view, actions) }.focusable() else Modifier)
                .padding(bottom = FsSpace.space10),
        ) {
            // On a whole-dp grid (the BiS page's DpColumn), so a 2,500 dp page does not drift down by a pixel per block.
            DpColumn(Modifier.fillMaxWidth()) {
                when (val phase = state.phase) {
                    EditorPhase.Loading -> LoadingBody(state, actions)
                    EditorPhase.Failed -> FailedBody(state, actions)
                    is EditorPhase.Ready -> ReadyBody(phase.view, state, actions)
                }
            }
        }
    }
}

@Composable
private fun UndoButton(canUndo: Boolean, undo: () -> Unit) {
    IconButton(onClick = undo, enabled = canUndo) {
        // An empty history is visibly off: the glyph at 40%.
        Icon(
            painterResource(R.drawable.ic_undo), contentDescription = stringResource(R.string.plan_undo),
            tint = FsColor.strong, modifier = Modifier.alpha(if (canUndo) 1f else DISABLED_ALPHA),
        )
    }
}

/** The request's class, race and level, for the head before the files are in. */
private class RequestFacts(val classSlug: String, val race: String, val level: Int)

private fun EditorRequest.facts(): RequestFacts = when (this) {
    is EditorRequest.Draft -> RequestFacts(classSlug, PlannerPageCopy.raceName(raceSlug), level ?: PlannerEditor.MAX_LEVEL)
    is EditorRequest.Code -> RequestFacts(classSlug, PlannerPageCopy.raceName(raceSlug), PlannerEditor.MAX_LEVEL)
    is EditorRequest.Copy -> RequestFacts(classSlug, "", PlannerEditor.MAX_LEVEL)
}

@Composable
private fun LoadingBody(state: EditorPageState, actions: EditorActions) {
    val copy = PlannerPageCopy(LocalResources.current)
    val facts = state.request.facts()
    Head(copy.className(facts.classSlug), facts.classSlug, copy.meta(facts.race, facts.level, null, "")) { PlannerAnswerSkeleton() }
    Controls(copy.className(facts.classSlug), facts.race, null, state, actions)
    PlateBlock(top = FsSpace.space5) { PlannerPlateSkeleton() }
    Sections {
        PlannerSkeletonSection(copy.heading(PlannerBand.bandFor(0, facts.level)))
        PlannerSkeletonSection(stringResource(R.string.plannerPage_inGame))
    }
}

@Composable
private fun FailedBody(state: EditorPageState, actions: EditorActions) {
    val copy = PlannerPageCopy(LocalResources.current)
    val facts = state.request.facts()
    val left = PlannerRules.pointsAvailable(facts.level)
    val answer = PlannerAnswer(listOf(PlannerFigure(stringResource(R.string.plannerPage_pointsLeft), left.toString(), copy.spent(0, facts.level, false), gold = true)), sim = actions.sim)
    Head(copy.className(facts.classSlug), facts.classSlug, copy.meta(facts.race, facts.level, null, "")) { PlannerAnswerBlock(answer) }
    Controls(copy.className(facts.classSlug), facts.race, null, state, actions)
    PlateBlock(top = FsSpace.space5) { PlannerPlateError(state.build, actions.retry) }
}

@Composable
private fun ReadyBody(view: EditorView, state: EditorPageState, actions: EditorActions) {
    val res = LocalResources.current
    val copy = PlannerPageCopy(res)
    val editor = view.state
    val band = view.band
    val race = view.race?.name.orEmpty()
    val spec = view.spec?.name?.takeIf { editor.spent > 0 }
    Head(view.className, view.file.classSlug, copy.meta(race, editor.level, spec, view.className)) { PlannerAnswerBlock(answer(view, state, actions, copy, res.getString(R.string.plannerPage_pointsLeft))) }
    Controls(view.className, race, view, state, actions)
    val note = view.shownNote?.let(copy::note)
    note?.let { PlannerNotes(it) }
    val editing = TalentEditing(actions.add, actions.remove, stringResource(R.string.plan_addPoint), stringResource(R.string.plan_removePoint))
    PlateBlock(top = FsSpace.space5) {
        PlannerPlate(
            PlateModel(
                view.index, editor.order, state.build, state.tree,
                differing = if (state.differsShown) band?.differing.orEmpty() else emptySet(),
                canSpend = editor.left > 0, selectedId = state.keyboardTalent,
            ),
            status = { StatusLine(view, state.tree, actions) },
            PlateActions(actions.onTree, actions.openTalent, editing, actions.reset),
        )
    }
    Sections {
        band?.let { compare ->
            PlannerBandSection(
                BandSectionModel(
                    copy.heading(compare.build.band), copy.line(compare),
                    load = copy.loadLabel(compare.build.band).takeUnless { compare.settled },
                    canDiffer = compare.differing.isNotEmpty() && !compare.settled && compare.coveredOf == null,
                    differsShown = state.differsShown, gearRange = copy.range(compare.build.band),
                ),
                BandSectionActions(actions.loadBand, actions.toggleDiffers, actions.openBis),
            )
        }
        PlannerInGameSection(InGameModel(state.sendTo, state.copiedForGame, editor.spent > 0), InGameActions(actions.send, actions.copyForGame, actions.addon))
        PlannerPointOrderSection(view.index, editor.order, state.build)
        val problem = view.shareProblem?.let { shareProblemText(it, res) }
        PlannerShareSection(
            ShareModel(state.title, copy.summary(race, view.className, view.index, view.totals, editor.level), state.copiedText, state.signedIn, problem),
            ShareActions(actions.onTitle, actions.share, actions.copyText, actions.builds, actions.signIn),
        )
        PlannerImportSection(ImportModel(state.importCode, state.importUnreadable), ImportActions(actions.onImportCode, actions.import, actions.addon))
        if (state.canOpenTalents) FlowLink(stringResource(R.string.plannerPage_talentsLink), Modifier.padding(top = FsSpace.space8), actions.talents)
    }
}

private fun answer(view: EditorView, state: EditorPageState, actions: EditorActions, copy: PlannerPageCopy, pointsLeft: String): PlannerAnswer {
    val editor = view.state
    val band = view.band
    val left = PlannerFigure(pointsLeft, editor.left.toString(), copy.spent(editor.spent, editor.level, state.forCharacter), gold = true)
    val score = band?.let { compare -> copy.score(compare)?.let { (label, value) -> PlannerFigure(label, value, copy.scoreBasis(compare.build.band)) } }
    return PlannerAnswer(listOfNotNull(left, score), sim = actions.sim)
}

/** The status line (`.st2`): a refusal in red, else the touched talent and its rank, which opens the rank sheet. */
@Composable
private fun StatusLine(view: EditorView, tree: Int, actions: EditorActions) {
    val strip = PlannerCells.strip(view.index, view.state, tree)?.takeIf { view.state.touched != null } ?: return
    val refusal = strip.refusal?.let(PlannerCopy(LocalResources.current)::refusal)
    val text = refusal ?: "${strip.talent.name} ${stringResource(R.string.plan_rank, strip.rank, strip.talent.maxRank)}"
    Text(
        text,
        style = STATUS,
        color = if (refusal != null) FsColor.wipe else FsColor.strong,
        maxLines = 2,
        overflow = TextOverflow.Ellipsis,
        modifier = Modifier
            .fillMaxWidth()
            .clickable(role = Role.Button) { actions.openTalent(strip.talent.id) }
            .semantics { liveRegion = LiveRegionMode.Polite },
    )
}

/** The arrows move the status line's talent, `+` and `−` add and remove (hardware keyboard). */
private fun treeKey(event: KeyEvent, view: EditorView, actions: EditorActions): Boolean {
    if (event.type != KeyEventType.KeyDown) return false
    val touched = view.state.touched
    return when (event.key) {
        Key.DirectionLeft -> actions.move(TreeDirection.Left).let { true }
        Key.DirectionRight -> actions.move(TreeDirection.Right).let { true }
        Key.DirectionUp -> actions.move(TreeDirection.Up).let { true }
        Key.DirectionDown -> actions.move(TreeDirection.Down).let { true }
        Key.Plus, Key.Equals, Key.NumPadAdd -> touched?.let { actions.add(it); true } ?: false
        Key.Minus, Key.NumPadSubtract -> touched?.let { actions.remove(it); true } ?: false
        else -> false
    }
}

@Composable
private fun Head(className: String, classSlug: String, meta: String, answer: @Composable ColumnScope.() -> Unit) {
    Box(Modifier.padding(horizontal = FsSpace.space4)) {
        FlowHead(stringResource(R.string.plannerPage_h1, className), listOf(meta), classColor(classSlug), crest = classSlug, answer = answer)
    }
}

@Composable
private fun Controls(className: String, race: String, view: EditorView?, state: EditorPageState, actions: EditorActions) {
    val slug = view?.file?.classSlug ?: state.request.classSlug
    val choices = PlannerChoices(
        classSlug = slug,
        race = race,
        level = view?.state?.level,
        onClass = actions.setClass,
        onRace = actions.openRace,
        onLevel = actions.openLevel,
    )
    PlannerControls(stringResource(R.string.plannerPage_bar, className, race), choices, state.panelOpen, actions.togglePanel)
}

/** The plate block (`.layout-wb`): 20 under the note or the bar, at the page's sides. */
@Composable
fun PlateBlock(top: androidx.compose.ui.unit.Dp, content: @Composable () -> Unit) {
    Box(Modifier.padding(start = FsSpace.space4, end = FsSpace.space4, top = top)) { content() }
}

/** The sections under the plate, at the page's sides. */
@Composable
fun Sections(content: @Composable () -> Unit) {
    DpColumn(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4), content = content)
}

/** The sentence a failed share leaves under Share (spec F13.3). */
fun shareProblemText(problem: ShareOutcome, res: android.content.res.Resources): String = when (problem) {
    is ShareOutcome.Field -> problem.sentence
    ShareOutcome.RateLimited -> res.getString(R.string.plan_rateLimited)
    ShareOutcome.Offline -> res.getString(R.string.plan_offline)
    ShareOutcome.Failed, is ShareOutcome.Shared -> res.getString(R.string.plan_shareFailed)
}

private const val DISABLED_ALPHA = 0.4f
private val STATUS = FsType.body.copy(fontSize = 14.sp, lineHeight = 20.sp)
