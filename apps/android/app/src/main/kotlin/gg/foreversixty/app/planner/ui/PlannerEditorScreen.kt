package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.focusable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Add
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.vector.rememberVectorPainter
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
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.PlannerCells
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.ShareOutcome
import gg.foreversixty.app.planner.logic.StripModel
import gg.foreversixty.app.planner.logic.TreeDirection
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.talents.TalentEditing
import gg.foreversixty.app.talents.TalentIconTile
import gg.foreversixty.app.talents.TalentsSkeleton
import java.util.Locale

/** What the editor page draws, beyond the model's phase: the shown tree and the keyboard's talent. */
data class EditorPageState(
    val phase: EditorPhase,
    val build: String,
    val tree: Int,
    /** The talent the hardware keyboard has reached: drawn with F6's selection ring. */
    val keyboardTalent: Int? = null,
)

/** Every action on the editor page; the defaults let the screenshot tests draw it inert. */
class EditorActions(
    val back: () -> Unit = {},
    val undo: () -> Unit = {},
    val add: (Int) -> Unit = {},
    val remove: (Int) -> Unit = {},
    val onTree: (Int) -> Unit = {},
    val openRace: () -> Unit = {},
    val openLevel: () -> Unit = {},
    val reset: () -> Unit = {},
    val share: () -> Unit = {},
    val retry: () -> Unit = {},
    /** An arrow key: the talent it reaches is the strip's. */
    val move: (TreeDirection) -> Unit = {},
)

/**
 * F13.1 Plan a build, full screen with no tab bar: the bar with Undo, the chips, the tabs and F6's
 * tree, the stepper strip and the action bar. Stateless, so the screenshot tests draw every state.
 */
@Composable
fun PlannerEditorPage(state: EditorPageState, actions: EditorActions, modifier: Modifier = Modifier) {
    val view = (state.phase as? EditorPhase.Ready)?.view
    Column(modifier.fillMaxSize().background(FsColor.bg)) {
        EditorBar(canUndo = view?.state?.canUndo == true, actions)
        when (val phase = state.phase) {
            EditorPhase.Loading -> TalentsSkeleton()
            EditorPhase.Failed -> ErrorRegion(stringResource(R.string.talents_failed), actions.retry)
            is EditorPhase.Ready -> EditorBody(phase.view, state, actions, Modifier.weight(1f))
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun EditorBar(canUndo: Boolean, actions: EditorActions) {
    TopAppBar(
        title = {
            Text(
                stringResource(R.string.plan_title).uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong,
                modifier = Modifier.semantics { heading() },
            )
        },
        navigationIcon = {
            IconButton(onClick = actions.back) {
                Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.talents_title), tint = FsColor.strong)
            }
        },
        actions = {
            IconButton(onClick = actions.undo, enabled = canUndo) {
                Icon(
                    painterResource(R.drawable.ic_undo), contentDescription = stringResource(R.string.plan_undo),
                    tint = if (canUndo) FsColor.strong else FsColor.muted,
                )
            }
        },
        colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
        windowInsets = WindowInsets(0, 0, 0, 0),
    )
}

@Composable
private fun EditorBody(view: EditorView, page: EditorPageState, actions: EditorActions, modifier: Modifier) {
    val res = LocalResources.current
    val copy = PlannerCopy(res)
    val editing = TalentEditing(actions.add, actions.remove, stringResource(R.string.plan_addPoint), stringResource(R.string.plan_removePoint))
    // From 1.5x the strip grows past what a pinned bar can hold: it scrolls with the tree, and only
    // the action bar stays pinned (spec §9, `the strip grows`).
    val large = isLargeText()
    val under = @Composable {
        view.shareProblem?.let { problem ->
            Text(
                shareProblemText(problem, res),
                style = FsType.body, color = FsColor.wipe,
                modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp).semantics { liveRegion = LiveRegionMode.Polite },
            )
        }
        PlannerCells.strip(view.index, view.state, page.tree)?.let { StepperStrip(it, page.build, copy, actions) }
    }
    Column(modifier.fillMaxWidth()) {
        ChipsRow(view, actions)
        Column(
            Modifier
                .weight(1f)
                .fillMaxWidth()
                .verticalScroll(rememberScrollState())
                .onKeyEvent { event -> treeKey(event, view, actions) }
                .focusable(),
        ) {
            PlannerTree(
                view.index, view.state.order, page.build, page.tree, actions.onTree,
                Modifier.padding(bottom = 12.dp), selectedId = page.keyboardTalent, editing = editing,
            )
            if (large) under()
        }
        if (!large) under()
        ActionBar(view.state.spent, actions)
    }
}

/** The arrows move the strip's talent, `+` and `−` add and remove (spec F13.1, hardware keyboard). */
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
private fun ChipsRow(view: EditorView, actions: EditorActions) {
    val large = isLargeText()
    val raceName = view.race?.name.orEmpty()
    val chips = @Composable {
        ClassCrest(view.file.classSlug, 30.dp, Modifier.semantics { contentDescription = view.className })
        if (raceName.isNotEmpty()) PlannerChip(raceName, "${stringResource(R.string.plan_raceSheet)}, $raceName", actions.openRace)
        val level = stringResource(R.string.plan_level, view.state.level)
        PlannerChip(level, level, actions.openLevel)
    }
    val left = @Composable { PointsLeft(view.state.left) }
    if (large) {
        Column(Modifier.fillMaxWidth().padding(horizontal = 16.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(10.dp)) { chips() }
            left()
        }
    } else {
        Row(
            Modifier.fillMaxWidth().padding(horizontal = 16.dp),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(10.dp),
        ) {
            chips()
            Spacer(Modifier.weight(1f))
            left()
        }
    }
}

/** `{n} left`, mono 20 gold; kill green at 0 (the points left never collapse, spec §9). */
@Composable
private fun PointsLeft(left: Int) {
    val spoken = stringResource(R.string.plan_a11yLeft, left)
    Text(
        stringResource(R.string.plan_left, left),
        style = FsType.monoDay,
        color = if (left == 0) FsColor.kill else FsColor.gold,
        modifier = Modifier.clearAndSetSemantics {
            contentDescription = spoken
            liveRegion = LiveRegionMode.Polite
        },
    )
}

/** The stepper strip (spec §4): the talent's icon 44, name and rank, its rank text or the refusal, and `−` `+`. */
@Composable
private fun StepperStrip(strip: StripModel, build: String, copy: PlannerCopy, actions: EditorActions) {
    val talent = strip.talent
    val line = strip.refusal?.let(copy::refusal) ?: strip.text
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 96.dp)
            .background(FsColor.raised)
            .drawBehind { drawLine(FsColor.line, Offset(0f, 0f), Offset(size.width, 0f), 1.dp.toPx()) }
            .padding(horizontal = 16.dp, vertical = 10.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        TalentIconTile(talent, strip.state, build, size = 44.dp)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Row(horizontalArrangement = Arrangement.spacedBy(6.dp), verticalAlignment = Alignment.CenterVertically) {
                Text(talent.name, style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f, fill = false))
                Text(stringResource(R.string.plan_rank, strip.rank, talent.maxRank), style = FsType.mono, color = FsColor.text)
            }
            line?.let {
                Text(
                    it, style = FsType.meta, color = if (strip.refusal != null) FsColor.wipe else FsColor.muted,
                    modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite },
                )
            }
        }
        StepButton(painterResource(R.drawable.ic_minus), stringResource(R.string.plan_removePoint), strip.rank > 0, { actions.remove(talent.id) })
        StepButton(rememberVectorPainter(Icons.Filled.Add), stringResource(R.string.plan_addPoint), true, { actions.add(talent.id) })
    }
}

/** Reset (secondary) and Share (the one primary, with the share glyph). */
@Composable
private fun ActionBar(spent: Int, actions: EditorActions) {
    Row(
        Modifier.fillMaxWidth().background(FsColor.raised).padding(start = 16.dp, end = 16.dp, top = 4.dp, bottom = 12.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        BrandButton(stringResource(R.string.plan_reset), BrandButtonKind.Secondary, actions.reset, Modifier.weight(RESET_WEIGHT), enabled = spent > 0, height = 48.dp)
        BrandButton(stringResource(R.string.plan_share), BrandButtonKind.Primary, actions.share, Modifier.weight(1f), height = 48.dp, leadingIcon = Icons.Filled.Share)
    }
}

/** The sentence a failed share leaves under the editor (spec F13.3). */
fun shareProblemText(problem: ShareOutcome, res: android.content.res.Resources): String = when (problem) {
    is ShareOutcome.Field -> problem.sentence
    ShareOutcome.RateLimited -> res.getString(R.string.plan_rateLimited)
    ShareOutcome.Offline -> res.getString(R.string.plan_offline)
    ShareOutcome.Failed, is ShareOutcome.Shared -> res.getString(R.string.plan_shareFailed)
}

private const val RESET_WEIGHT = 0.6f
