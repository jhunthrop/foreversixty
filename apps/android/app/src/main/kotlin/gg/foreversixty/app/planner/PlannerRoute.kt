package gg.foreversixty.app.planner

import android.content.Context
import android.content.Intent
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.systemBarsPadding
import androidx.compose.material3.SnackbarDuration
import androidx.compose.material3.SnackbarHostState
import androidx.compose.material3.SnackbarResult
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.platform.LocalView
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.planner.logic.EditorPhase
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.EditorView
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.PlannerEditor
import gg.foreversixty.app.planner.logic.PlannerEditorModel
import gg.foreversixty.app.planner.logic.PlannerKeys
import gg.foreversixty.app.planner.logic.PlannerScreen
import gg.foreversixty.app.planner.logic.ShareOutcome
import gg.foreversixty.app.planner.logic.SharedBuildModel
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.ViewPlannerHaptics
import gg.foreversixty.app.planner.logic.YourBuildsModel
import gg.foreversixty.app.planner.logic.YourBuildsPhase
import gg.foreversixty.app.planner.ui.BuildsActions
import gg.foreversixty.app.planner.ui.BuildsPageState
import gg.foreversixty.app.planner.ui.EditorActions
import gg.foreversixty.app.planner.ui.EditorPageState
import gg.foreversixty.app.planner.ui.LevelSheetContent
import gg.foreversixty.app.planner.ui.NewBuildRow
import gg.foreversixty.app.planner.ui.PlannerEditorPage
import gg.foreversixty.app.planner.ui.PlannerSheet
import gg.foreversixty.app.planner.ui.RaceSheetContent
import gg.foreversixty.app.planner.ui.ResetDialog
import gg.foreversixty.app.planner.ui.ShareSheetContent
import gg.foreversixty.app.planner.ui.ShareSheetState
import gg.foreversixty.app.planner.ui.SharedActions
import gg.foreversixty.app.planner.ui.SharedBuildPage
import gg.foreversixty.app.planner.ui.SharedPageState
import gg.foreversixty.app.planner.ui.YourBuildsPage
import gg.foreversixty.app.planner.ui.openingTree
import gg.foreversixty.app.talents.TalentBuildKind
import gg.foreversixty.app.talents.TalentExport
import gg.foreversixty.app.talents.TalentGuideBuild
import gg.foreversixty.app.talents.TalentRankSheet
import gg.foreversixty.app.talents.TalentsPresenter
import gg.foreversixty.app.talents.TalentRules
import gg.foreversixty.app.model.TalentsFile
import java.time.Instant
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/** The planner's top screen, drawn by the scaffold (the editor over everything, the others in the tab). */
@Composable
fun PlannerScreenHost(feature: PlannerFeature, screen: PlannerScreen) {
    val back = feature.navigator::pop
    BackHandler(onBack = back)
    when (screen) {
        is PlannerScreen.Editor -> PlannerEditorRoute(feature, screen.request, back)
        is PlannerScreen.Viewer -> SharedBuildRoute(feature, screen.id, back)
        PlannerScreen.Builds -> YourBuildsRoute(feature, back)
    }
}

/** Which of the editor's sheets is up. */
private enum class EditorSheet { None, Race, Level, Reset, Share }

@Composable
private fun PlannerEditorRoute(feature: PlannerFeature, request: EditorRequest, back: () -> Unit) {
    val view = LocalView.current
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val model = remember(request) { PlannerEditorModel(feature.source, feature.drafts, request, ViewPlannerHaptics(view)) }
    val phase by model.phase.collectAsStateWithLifecycle()
    var tree by rememberSaveable(request) { mutableStateOf(0) }
    var settled by rememberSaveable(request) { mutableStateOf(false) }
    var keyboardTalent by remember { mutableStateOf<Int?>(null) }
    var sheet by rememberSaveable { mutableStateOf(EditorSheet.None) }
    var title by rememberSaveable { mutableStateOf("") }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    LaunchedEffect(model) {
        model.load()
        val ready = (model.phase.value as? EditorPhase.Ready)?.view
        if (!settled && ready != null) {
            settled = true
            tree = openingTree(ready.file, ready.state.order)
        }
    }
    val ready = (phase as? EditorPhase.Ready)?.view
    PlannerEditorPage(
        EditorPageState(phase, model.build, tree, keyboardTalent),
        EditorActions(
            back = back,
            undo = model::undo,
            add = model::add,
            remove = model::remove,
            onTree = { tree = it },
            openRace = { sheet = EditorSheet.Race },
            openLevel = { sheet = EditorSheet.Level },
            reset = { sheet = EditorSheet.Reset },
            share = { sheet = EditorSheet.Share },
            retry = { scope.launch { model.load() } },
            move = { direction ->
                ready?.index?.trees?.getOrNull(tree)?.let { shown -> PlannerKeys.move(shown, ready.state.touched, direction) }?.let {
                    model.touch(it)
                    keyboardTalent = it
                }
            },
            openTalent = { sheetId = it },
        ),
        Modifier.systemBarsPadding(),
    )
    if (ready == null) return
    sheetId?.let { talentId ->
        rankSheet(ready.file, ready.state.ranks).sheet(TalentBuildKind.Yours, talentId, res)?.let {
            TalentRankSheet(it, model.build) { sheetId = null }
        }
    }
    val dismiss = { sheet = EditorSheet.None }
    when (sheet) {
        EditorSheet.None -> Unit
        EditorSheet.Race -> PlannerSheet(dismiss) {
            RaceSheetContent(ready.races, ready.state.raceSlug) {
                model.setRace(it)
                dismiss()
            }
        }
        EditorSheet.Level -> PlannerSheet(dismiss) { LevelSheetContent(ready.state.level, PlannerEditor.minLevel(ready.state), model::setLevel) }
        EditorSheet.Reset -> ResetDialog(ready.state.spent, clear = {
            model.reset()
            dismiss()
        }, cancel = dismiss)
        EditorSheet.Share -> PlannerSheet(dismiss) {
            ShareSheetContent(
                ShareSheetState(title, ready.sharing, feature.signedIn()),
                onTitle = { title = it },
                proceed = {
                    scope.launch {
                        val outcome = model.share(title)
                        dismiss()
                        if (outcome is ShareOutcome.Shared) openShareSheet(context, outcome.saved.url, shareText(ready, res))
                    }
                },
                cancel = dismiss,
            )
        }
    }
}

/** The system share sheet with the link (spec F13.4). */
private fun openShareSheet(context: Context, url: String, title: String) {
    val send = Intent(Intent.ACTION_SEND).setType(SHARE_TYPE).putExtra(Intent.EXTRA_TEXT, url).putExtra(Intent.EXTRA_TITLE, title)
    context.startActivity(Intent.createChooser(send, title))
}

/** `Orc Shaman · 30 / 3 / 18`: the share sheet's preview line. */
private fun shareText(view: EditorView, res: android.content.res.Resources): String =
    res.getString(R.string.plan_shareText, view.race?.name.orEmpty(), view.className, PlannerCopy(res).totals(view.totals))

@Composable
private fun SharedBuildRoute(feature: PlannerFeature, id: String, back: () -> Unit) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val model = remember(id) { SharedBuildModel(feature.source, id) }
    val phase by model.phase.collectAsStateWithLifecycle()
    var tree by rememberSaveable(id) { mutableStateOf(0) }
    var settled by rememberSaveable(id) { mutableStateOf(false) }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    LaunchedEffect(model) {
        model.load()
        val loaded = (model.phase.value as? SharedBuildPhase.Loaded)?.view
        if (!settled && loaded != null) {
            settled = true
            tree = openingTree(loaded.file, loaded.build.pointOrder)
        }
    }
    val loaded = (phase as? SharedBuildPhase.Loaded)?.view
    val fromList = feature.navigator.stack.dropLast(1).lastOrNull() == PlannerScreen.Builds
    SharedBuildPage(
        SharedPageState(
            phase = phase,
            tree = tree,
            offlineStamp = loaded?.offlineSinceMillis?.let { OfflineStamp.format(Instant.ofEpochMilli(it), Instant.now()) },
            backLabel = if (fromList) res.getString(R.string.plan_yourBuilds) else null,
        ),
        SharedActions(
            back = back,
            share = { buildId -> openShareSheet(context, DeepLink.SharedBuild(buildId).url(), res.getString(R.string.plan_shared_title)) },
            editCopy = { view -> feature.navigator.push(PlannerScreen.Editor(EditorRequest.Copy(view.classSlug, view.build))) },
            onTree = { tree = it },
            openTalent = { sheetId = it },
            retry = { scope.launch { model.load() } },
        ),
    )
    val sheet = loaded?.let { view -> sheetId?.let { talentId -> rankSheet(view.file, view.ranks).sheet(TalentBuildKind.Yours, talentId, res) } }
    sheet?.let { TalentRankSheet(it, loaded.build.treeVersion) { sheetId = null } }
}

/** F6's rank sheet over a planner build's ranks (the editor's strip, a shared build): the same sheet, read only. */
private fun rankSheet(file: TalentsFile, ranks: Map<Int, Int>): TalentsPresenter {
    val trees = TalentRules.trees(file).map { tree -> tree.talents.map { ranks[it.id] ?: 0 } }
    return TalentsPresenter(file, TalentGuideBuild.None, TalentExport.Ranks(trees), level = null, specName = "")
}

@Composable
private fun YourBuildsRoute(feature: PlannerFeature, back: () -> Unit) {
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val model = remember { YourBuildsModel(feature.source) }
    val phase by model.phase.collectAsStateWithLifecycle()
    val snackbar = remember { SnackbarHostState() }
    LaunchedEffect(model) { model.load() }
    val pending = (phase as? YourBuildsPhase.Loaded)?.pending
    LaunchedEffect(pending) {
        if (pending == null) return@LaunchedEffect
        val shown = launch {
            val answer = snackbar.showSnackbar(res.getString(R.string.plan_removed), res.getString(R.string.plan_removedUndo), duration = SnackbarDuration.Indefinite)
            // Outside this effect: the undo clears `pending`, which cancels the effect mid-save.
            if (answer == SnackbarResult.ActionPerformed) scope.launch { model.undo() }
        }
        delay(SNACKBAR_MILLIS)
        if (shown.isActive) {
            snackbar.currentSnackbarData?.dismiss()
            model.expire()
        }
    }
    val character = feature.character()
    val newBuild = character?.className?.let { slug ->
        NewBuildRow(slug, PlannerEditorModel.titleCase(slug), character.level ?: PlannerEditor.MAX_LEVEL)
    }
    YourBuildsPage(
        BuildsPageState(phase, newBuild),
        BuildsActions(
            back = back,
            open = { feature.navigator.push(PlannerScreen.Viewer(it.build.id)) },
            remove = { row -> scope.launch { model.remove(row.build.id) } },
            planNew = { new -> feature.plan(new.classSlug, PlannerEditorModel.raceSlug(character?.race), new.level) },
            retry = { scope.launch { model.load() } },
        ),
        snackbar,
    )
}

private const val SHARE_TYPE = "text/plain"
/** The snackbar's `Removed. Undo` stays 5 s (spec F13.6). */
private const val SNACKBAR_MILLIS = 5_000L
