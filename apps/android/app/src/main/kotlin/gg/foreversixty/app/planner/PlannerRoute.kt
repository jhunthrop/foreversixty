package gg.foreversixty.app.planner

import android.content.ClipData
import android.content.Context
import android.content.Intent
import androidx.compose.ui.platform.ClipEntry
import androidx.compose.ui.platform.LocalClipboard
import androidx.core.net.toUri
import gg.foreversixty.app.planner.logic.ImportOutcome
import gg.foreversixty.app.planner.logic.PlannerCodes
import gg.foreversixty.app.planner.logic.PlannerPageCopy
import gg.foreversixty.app.planner.ui.LoadBandDialog
import gg.foreversixty.app.planner.ui.LevelSheetContent
import gg.foreversixty.app.system.logic.ShareMessage
import gg.foreversixty.app.system.logic.ShareText
import gg.foreversixty.app.talents.TalentsPhase
import gg.foreversixty.app.talents.TalentsPage
import gg.foreversixty.app.planner.logic.SharedCompare
import androidx.activity.compose.BackHandler
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
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
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
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.planner.logic.YourBuildsModel
import gg.foreversixty.app.planner.logic.YourBuildsPhase
import gg.foreversixty.app.planner.ui.BuildsActions
import gg.foreversixty.app.planner.ui.BuildsPageState
import gg.foreversixty.app.planner.ui.EditorActions
import gg.foreversixty.app.planner.ui.EditorPageState
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

/**
 * What the shell gives the planner's screens: the label Back returns to, the sign-in, the app's own
 * links (a BiS list opens in the app), and the Talents compare under the menu's Planner row.
 */
class PlannerHost(
    val backTo: String = "",
    val signIn: () -> Unit = {},
    val openLink: (String) -> Unit = {},
    /** Null when the Talents compare cannot open (no current character, `talents` off). */
    val openTalents: (() -> Unit)? = null,
)

/** The planner's top screen, drawn by the scaffold under the current tab. */
@Composable
fun PlannerScreenHost(feature: PlannerFeature, screen: PlannerScreen, host: PlannerHost = PlannerHost()) {
    val back = feature.navigator::pop
    BackHandler(onBack = back)
    when (screen) {
        is PlannerScreen.Editor -> PlannerEditorRoute(feature, screen.request, host, back)
        is PlannerScreen.Viewer -> SharedBuildRoute(feature, screen.id, host, back)
        PlannerScreen.Builds -> YourBuildsRoute(feature, back)
    }
}

/** Which of the editor's sheets is up. */
private enum class EditorSheet { None, Race, Level, Reset, Share, LoadBand }

/** A link out of the planner: the planner closes so the app under it shows the page. */
private fun leave(feature: PlannerFeature, host: PlannerHost, url: String) {
    feature.navigator.close()
    host.openLink(url)
}

@Composable
private fun PlannerEditorRoute(feature: PlannerFeature, request: EditorRequest, host: PlannerHost, back: () -> Unit) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val haptics = rememberHaptics()
    val clipboard = LocalClipboard.current
    val model = remember(request) { PlannerEditorModel(feature.source, feature.drafts, request, haptics, feature.catalog) }
    val phase by model.phase.collectAsStateWithLifecycle()
    var tree by rememberSaveable(request) { mutableStateOf(0) }
    var settled by rememberSaveable(request) { mutableStateOf(false) }
    var keyboardTalent by remember { mutableStateOf<Int?>(null) }
    var sheet by rememberSaveable { mutableStateOf(EditorSheet.None) }
    var panelOpen by rememberSaveable(request) { mutableStateOf(false) }
    var differsShown by rememberSaveable(request) { mutableStateOf(false) }
    var title by rememberSaveable { mutableStateOf("") }
    var importCode by rememberSaveable(request) { mutableStateOf("") }
    var importUnreadable by remember(request) { mutableStateOf(false) }
    var copied by remember { mutableStateOf<Copied?>(null) }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    LaunchedEffect(model) {
        model.load()
        val ready = (model.phase.value as? EditorPhase.Ready)?.view
        if (!settled && ready != null) {
            settled = true
            tree = openingTree(ready.file, ready.state.order)
        }
    }
    LaunchedEffect(copied) {
        if (copied != null) {
            delay(COPIED_MILLIS)
            copied = null
        }
    }
    val ready = (phase as? EditorPhase.Ready)?.view
    val character = feature.character()
    val sendTo = character?.takeIf { feature.signedIn() && it.className.equals(ready?.file?.classSlug ?: request.classSlug, ignoreCase = true) }?.name
    val spent = ready?.state?.spent ?: 0
    val code = ready?.let { PlannerCodes.fs1(it.index, it.state.order, model.build, it.state.raceSlug.orEmpty()) }
    val copy = { text: String, what: Copied ->
        scope.launch { clipboard.setClipEntry(ClipEntry(ClipData.newPlainText(text, text))) }
        copied = what
    }
    PlannerEditorPage(
        EditorPageState(
            phase, model.build, tree, request, keyboardTalent,
            backTo = host.backTo, panelOpen = panelOpen, differsShown = differsShown, title = title,
            importCode = importCode, importUnreadable = importUnreadable,
            copiedForGame = copied == Copied.Game, copiedText = copied == Copied.Text,
            sendTo = sendTo, signedIn = feature.signedIn(),
            forCharacter = (request as? EditorRequest.Draft)?.level != null,
            canOpenTalents = host.openTalents != null,
        ),
        EditorActions(
            back = back,
            undo = model::undo,
            add = model::add,
            remove = model::remove,
            onTree = { tree = it },
            togglePanel = { panelOpen = !panelOpen },
            setClass = { slug ->
                if (slug != (ready?.file?.classSlug ?: request.classSlug)) {
                    feature.navigator.replaceTop(PlannerScreen.Editor(EditorRequest.Draft(slug, ready?.state?.raceSlug)))
                }
            },
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
            loadBand = { if (spent > 0) sheet = EditorSheet.LoadBand else model.loadBand() },
            toggleDiffers = { differsShown = !differsShown },
            openBis = { ready?.band?.let { leave(feature, host, DeepLink.Bis(it.build.spec.spec, it.build.band.toString()).url()) } },
            sim = { openWeb(context, code?.takeIf { spent > 0 }?.let(PlannerCodes::simFor) ?: PlannerCodes.SIM) },
            send = { code?.let { openWeb(context, PlannerCodes.sitePlanner(it)) } },
            copyForGame = { ready?.takeIf { it.state.spent > 0 }?.let { copy(PlannerCodes.fsb1(it.index, it.state.order, model.build), Copied.Game) } },
            addon = { openWeb(context, PlannerCodes.ADDON) },
            onTitle = { title = it },
            copyText = { ready?.let { copy(summary(it, res), Copied.Text) } },
            builds = feature::openBuilds,
            signIn = host.signIn,
            onImportCode = {
                importCode = it
                importUnreadable = false
            },
            import = {
                when (val outcome = model.import(importCode)) {
                    ImportOutcome.Loaded -> importCode = ""
                    ImportOutcome.Unreadable -> importUnreadable = importCode.isNotBlank()
                    is ImportOutcome.OtherClass -> feature.navigator.replaceTop(PlannerScreen.Editor(EditorRequest.Code(outcome.classSlug, outcome.code)))
                }
            },
            talents = { host.openTalents?.invoke() },
        ),
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
        EditorSheet.LoadBand -> LoadBandDialog(load = {
            model.loadBand()
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
                        if (outcome is ShareOutcome.Shared) ShareText.send(context, ShareMessage(shareText(ready, res), outcome.saved.url))
                    }
                },
                cancel = dismiss,
            )
        }
    }
}

/** What was last copied: its button reads `Copied` for a moment. */
private enum class Copied { Game, Text }

private fun openWeb(context: Context, url: String) {
    context.startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
}

/** `Dwarf Hunter · Beast Mastery 21/Marksmanship 0/Survival 0 · Level 34`: Copy as text. */
private fun summary(view: EditorView, res: android.content.res.Resources): String =
    PlannerPageCopy(res).summary(view.race?.name.orEmpty(), view.className, view.index, view.totals, view.state.level)

/** `Orc Shaman · 30 / 3 / 18`: the share sheet's preview line. */
private fun shareText(view: EditorView, res: android.content.res.Resources): String =
    res.getString(R.string.plan_shareText, view.race?.name.orEmpty(), view.className, PlannerCopy(res).totals(view.totals))

@Composable
private fun SharedBuildRoute(feature: PlannerFeature, id: String, host: PlannerHost, back: () -> Unit) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val clipboard = LocalClipboard.current
    val model = remember(id) { SharedBuildModel(feature.source, id, feature.catalog) }
    val phase by model.phase.collectAsStateWithLifecycle()
    var tree by rememberSaveable(id) { mutableStateOf(0) }
    var settled by rememberSaveable(id) { mutableStateOf(false) }
    var differsShown by rememberSaveable(id) { mutableStateOf(false) }
    var copied by remember { mutableStateOf(false) }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    LaunchedEffect(model) {
        model.load()
        val loaded = (model.phase.value as? SharedBuildPhase.Loaded)?.view
        if (!settled && loaded != null) {
            settled = true
            tree = openingTree(loaded.file, loaded.build.pointOrder)
        }
    }
    LaunchedEffect(copied) {
        if (copied) {
            delay(COPIED_MILLIS)
            copied = false
        }
    }
    val loaded = (phase as? SharedBuildPhase.Loaded)?.view
    val fromList = feature.navigator.stack.dropLast(1).lastOrNull() == PlannerScreen.Builds
    // Lane F15: `Compare with mine` draws F6's compare over this build, with the shared build's words.
    val mine = loaded?.let { feature.compareExport(it.classSlug) }
    var comparing by rememberSaveable(id) { mutableStateOf(false) }
    if (comparing && loaded != null && mine != null) {
        BackHandler { comparing = false }
        SharedCompareRoute(SharedCompare.presenter(loaded, mine), loaded.build.treeVersion) { comparing = false }
        return
    }
    val character = feature.character()
    val sendTo = loaded?.let { view -> character?.takeIf { feature.signedIn() && it.className.equals(view.classSlug, ignoreCase = true) }?.name }
    SharedBuildPage(
        SharedPageState(
            phase = phase,
            tree = tree,
            offlineStamp = loaded?.offlineSinceMillis?.let { OfflineStamp.format(Instant.ofEpochMilli(it), Instant.now()) },
            canCompare = mine != null,
            backTo = if (fromList) res.getString(R.string.plan_yourBuilds) else host.backTo,
            differsShown = differsShown,
            copiedForGame = copied,
            sendTo = sendTo,
        ),
        SharedActions(
            back = back,
            share = { buildId ->
                val text = loaded?.let { res.getString(R.string.plan_shareText, it.raceName, it.className, PlannerCopy(res).totals(it.totals)) }
                ShareText.send(context, ShareMessage(text ?: res.getString(R.string.plan_shared_title), DeepLink.SharedBuild(buildId).url()))
            },
            editCopy = { view -> feature.navigator.push(PlannerScreen.Editor(EditorRequest.Copy(view.classSlug, view.build))) },
            compare = { comparing = true },
            onTree = { tree = it },
            openTalent = { sheetId = it },
            retry = { scope.launch { model.load() } },
            sim = { view -> openWeb(context, PlannerCodes.simForBuild(view.build.id)) },
            send = { view -> openWeb(context, PlannerCodes.sitePlanner(PlannerCodes.fs1(view.index, view.build.pointOrder, view.build.treeVersion, view.raceSlug))) },
            copyForGame = { view ->
                val text = PlannerCodes.fsb1(view.index, view.build.pointOrder, view.build.treeVersion)
                scope.launch { clipboard.setClipEntry(ClipEntry(ClipData.newPlainText(text, text))) }
                copied = true
            },
            addon = { openWeb(context, PlannerCodes.ADDON) },
            toggleDiffers = { differsShown = !differsShown },
            openBis = { view -> view.band?.let { leave(feature, host, DeepLink.Bis(it.build.spec.spec, it.build.band.toString()).url()) } },
        ),
    )
    val sheet = loaded?.let { view -> sheetId?.let { talentId -> rankSheet(view.file, view.ranks).sheet(TalentBuildKind.Yours, talentId, res) } }
    sheet?.let { TalentRankSheet(it, loaded.build.treeVersion) { sheetId = null } }
}

/** `Copied` stays on the button 3 s (the web's CopyForGame). */
private const val COPIED_MILLIS = 3_000L

/** F6's compare over a shared build (F13.5): the same screen as Talents, titled `Compare with mine`. */
@Composable
private fun SharedCompareRoute(presenter: TalentsPresenter, build: String, back: () -> Unit) {
    val res = LocalResources.current
    var shown by rememberSaveable { mutableStateOf(presenter.initialBuild) }
    var tree by rememberSaveable { mutableStateOf(presenter.openingTree) }
    var sheetId by remember { mutableStateOf<Int?>(null) }
    TalentsPage(
        phase = TalentsPhase.Loaded(presenter, null),
        build = build,
        offlineStamp = null,
        shown = shown,
        tree = tree,
        onShow = { shown = it },
        onTree = { tree = it },
        selectedId = sheetId,
        openTalent = { sheetId = it },
        back = back,
        title = res.getString(R.string.plan_cmp_title),
    )
    presenter.sheet(shown, sheetId ?: return, res)?.let { TalentRankSheet(it, build) { sheetId = null } }
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

/** The snackbar's `Removed. Undo` stays 5 s (spec F13.6). */
private const val SNACKBAR_MILLIS = 5_000L
