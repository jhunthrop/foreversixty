package gg.foreversixty.app.bis.spec

import android.content.Context
import android.content.Intent
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.Checkbox
import androidx.compose.material3.CheckboxDefaults
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.core.net.toUri
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.BisPhase
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.common.NewBuildStrip
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowWideButton
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.system.logic.ShareMessage
import gg.foreversixty.app.system.logic.ShareText
import java.time.Instant
import kotlinx.coroutines.launch
import androidx.compose.ui.unit.dp

/**
 * A spec's list pushed under Gear for a player not signed in (mocks `bis-spec-*`): its page over [BisModel],
 * the Change panel, the Sources sheet, the item tooltip, pull to refresh and the offline strip.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BisSpecRoute(model: BisModel, catalog: ItemCatalog, onBack: () -> Unit, onSignIn: () -> Unit) {
    val state by model.state.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val res = LocalResources.current
    val selection = state.context.selection
    var view by rememberSaveable { mutableStateOf(BisSpecView.BySlot) }
    var panel by rememberSaveable { mutableStateOf(false) }
    var off by remember(selection.spec) { mutableStateOf(emptySet<BisSource>()) }
    var sourcesOpen by remember { mutableStateOf(false) }
    var refreshing by remember { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(model) { model.start() }
    BackHandler(onBack = onBack)
    val composer = remember(res) { BisSpecComposer(res, model.catalog) }
    val phase = state.phase
    val published = (phase as? BisPhase.Loaded)?.published
    val build = state.build.orEmpty()
    var searching by rememberSaveable { mutableStateOf(false) }
    var query by rememberSaveable { mutableStateOf("") }
    val content = published?.let { composer.content(BisSpecInput(it.file, it.band, it.items, selection, state.context.className(selection.classSlug), off)) }
    val pick: (BisSelection) -> Unit = { scope.launch { model.select(it) } }
    Column(Modifier.fillMaxSize()) {
        val share = content?.let { page -> { ShareText.send(context, ShareMessage(res.getString(R.string.bisSpec_share, page.title), BisSpecLinks.list(selection))) } }
        FlowBackBar(stringResource(R.string.tab_gear), onBack, share) {
            // The list's own search (the old BiS screen's): the bar's button opens the field under it.
            IconButton(onClick = {
                searching = !searching
                if (!searching) query = ""
            }) {
                Icon(if (searching) Icons.Filled.Close else Icons.Filled.Search, contentDescription = stringResource(R.string.bis_search), tint = FsColor.strong)
            }
        }
        if (searching) {
            OutlinedTextField(
                value = query,
                onValueChange = { query = it },
                placeholder = { Text(stringResource(R.string.bis_search)) },
                singleLine = true,
                modifier = Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4, vertical = FsSpace.space2),
            )
        }
        NewBuildStrip(model) { scope.launch { model.update() } }
        offlineSince(phase)?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(Instant.ofEpochMilli(it), Instant.now()))) }
        PullToRefreshBox(
            isRefreshing = refreshing,
            onRefresh = { scope.launch { refreshing = true; model.refresh(); refreshing = false } },
            modifier = Modifier.fillMaxSize(),
        ) {
            if (content != null) {
                val actions = BisSpecActions(
                    onPick = pick,
                    onView = { view = it },
                    onSources = { sourcesOpen = true },
                    onItem = { look -> tooltip = request(look, selection, build) },
                    onSignIn = onSignIn,
                    onOpenSim = { context.open(BisSpecLinks.SIM) },
                    onPlanner = { context.open(BisSpecLinks.planner(selection)) },
                    onAddon = { context.open(AccountLinks.SETUP) },
                    onGuide = { context.open(BisSpecLinks.guide(selection)) },
                )
                BisSpecPage(content.filtered(query), view, build, panel, { panel = !panel }, actions)
            } else {
                Status(phase) { scope.launch { model.refresh() } }
            }
        }
    }
    if (sourcesOpen && published != null) {
        ModalBottomSheet(
            onDismissRequest = { sourcesOpen = false },
            sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
            containerColor = GameColors.sheet,
        ) {
            SourcesSheet(composer.sources(published.band), off) { source -> off = if (source in off) off - source else off + source }
        }
    }
    tooltip?.let { ItemTooltipSheet(it, catalog, onDismiss = { tooltip = null }) }
}

/** Every source the list draws from, ticked when it is on. */
@Composable
private fun SourcesSheet(sources: List<BisSource>, off: Set<BisSource>, toggle: (BisSource) -> Unit) {
    Column(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space4).padding(bottom = FsSpace.space6)) {
        Text(stringResource(R.string.bisSpec_sources), style = FsType.section, color = FsColor.strong, modifier = Modifier.padding(bottom = FsSpace.space2))
        sources.forEach { source ->
            Row(
                Modifier.fillMaxWidth().clickable(role = Role.Checkbox) { toggle(source) }.padding(vertical = 2.dp),
                verticalAlignment = Alignment.CenterVertically,
            ) {
                Checkbox(source !in off, onCheckedChange = { toggle(source) }, colors = CheckboxDefaults.colors(checkedColor = FsColor.gold, checkmarkColor = FsColor.bg))
                Text(stringResource(source.label), style = FsType.body, color = FsColor.text)
            }
        }
    }
}

/** Loading, no list for this range, not saved offline, or failed. */
@Composable
private fun Status(phase: BisPhase, retry: () -> Unit) {
    FlowPage {
        when (phase) {
            BisPhase.Loading, is BisPhase.Loaded -> {
                FlowSkeleton(240.dp, 32.dp)
                FlowSkeleton(160.dp, 18.dp, Modifier.padding(top = FsSpace.space3))
            }
            is BisPhase.NoList -> FlowNote(stringResource(R.string.home_answer_noList))
            BisPhase.NotCached -> {
                FlowNote(stringResource(R.string.bis_notCached))
                FlowWideButton(stringResource(R.string.common_tryAgain), primary = false, onClick = retry)
            }
            BisPhase.Failed -> {
                FlowNote(stringResource(R.string.bis_failed), color = FsColor.wipe)
                FlowWideButton(stringResource(R.string.common_tryAgain), primary = false, onClick = retry)
            }
        }
    }
}

private fun offlineSince(phase: BisPhase): Long? = when (phase) {
    is BisPhase.Loaded -> phase.offlineSinceMillis
    is BisPhase.NoList -> phase.offlineSinceMillis
    else -> null
}

private fun request(look: ItemLook, selection: BisSelection, build: String) = ItemTooltipRequest(
    itemId = look.id, fallback = look, build = build, classSlug = selection.classSlug, spec = selection.spec,
    specName = null, band = selection.range.band, bandLabel = null, preset = selection.range.preset, faction = selection.faction,
)

private fun Context.open(url: String) {
    startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
}

/** The site's pages the spec page links to. */
object BisSpecLinks {
    private const val SITE = "https://foreversixty.gg"
    const val SIM = "$SITE/sim"

    private fun specSlug(selection: BisSelection) = selection.spec.removePrefix(selection.classSlug + "-")

    fun list(selection: BisSelection) = "$SITE/bis/${selection.classSlug}/${specSlug(selection)}"

    fun guide(selection: BisSelection) = "$SITE/guides/${selection.classSlug}/${specSlug(selection)}"

    fun planner(selection: BisSelection) = "$SITE/planner/${selection.classSlug}"
}
