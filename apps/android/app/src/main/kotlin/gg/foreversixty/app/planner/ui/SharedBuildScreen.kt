package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.SharedBuildPhase
import gg.foreversixty.app.planner.logic.SharedBuildView
import gg.foreversixty.app.reports.ui.BarAction
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.talents.TalentsSkeleton
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** What the viewer page draws beyond the phase: the shown tree, the offline stamp, Compare's presence. */
data class SharedPageState(
    val phase: SharedBuildPhase,
    val tree: Int,
    val offlineStamp: String? = null,
    /** Compare with mine shows only when the current character has an export of that class. */
    val canCompare: Boolean = false,
    val zone: ZoneId = ZoneId.systemDefault(),
    /** The back button's label: the screen it returns to (`Your builds`), else `Back`. */
    val backLabel: String? = null,
)

class SharedActions(
    val back: () -> Unit = {},
    val share: (String) -> Unit = {},
    val editCopy: (SharedBuildView) -> Unit = {},
    val compare: (SharedBuildView) -> Unit = {},
    val onTree: (Int) -> Unit = {},
    val openTalent: (Int) -> Unit = {},
    val retry: () -> Unit = {},
)

/** F13.5 A shared build: the head, Edit a copy (and Compare), and its trees drawn read only. */
@Composable
fun SharedBuildPage(state: SharedPageState, actions: SharedActions) {
    val loaded = (state.phase as? SharedBuildPhase.Loaded)?.view
    val share = loaded?.let { view ->
        BarAction(stringResource(R.string.plan_share), { rememberVectorPainter(Icons.Filled.Share) }) { actions.share(view.build.id) }
    }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        ReportsBar(stringResource(R.string.plan_shared_title), state.backLabel ?: stringResource(R.string.items_slot_back), actions.back, share)
        when (val phase = state.phase) {
            SharedBuildPhase.Loading -> TalentsSkeleton()
            SharedBuildPhase.NotFound -> Message(stringResource(R.string.plan_shared_notFound))
            SharedBuildPhase.Offline -> Message(stringResource(R.string.plan_shared_offline))
            SharedBuildPhase.Failed -> ErrorRegion(stringResource(R.string.plan_shared_failed), actions.retry)
            is SharedBuildPhase.Loaded -> Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
                state.offlineStamp?.let { OfflineStrip(stringResource(R.string.offline_strip, it)) }
                SharedHead(phase.view, state)
                SharedButtons(phase.view, state.canCompare, actions)
                PlannerTree(
                    phase.view.index, phase.view.build.pointOrder, phase.view.build.treeVersion, state.tree, actions.onTree,
                    Modifier.padding(top = 8.dp, bottom = 24.dp), open = actions.openTalent,
                )
            }
        }
    }
}

@Composable
private fun Message(text: String) {
    Text(text, style = FsType.lead, color = FsColor.text, modifier = Modifier.padding(16.dp))
}

@Composable
private fun SharedHead(view: SharedBuildView, state: SharedPageState) {
    val copy = PlannerCopy(LocalResources.current)
    val totals = copy.totals(view.totals)
    val date = view.savedAt?.let { DAY_MONTH.withZone(state.zone).format(it) }.orEmpty()
    val views = view.build.views.toInt()
    val meta = view.build.title?.let { stringResource(R.string.plan_shared_meta, it, date, views) }
        ?: stringResource(R.string.plan_shared_metaNoTitle, date, views)
    Row(
        Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ClassCrest(view.classSlug, 44.dp)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(headLine(stringResource(R.string.plan_shared_head, view.raceName, view.className, totals), totals), style = FsType.bodyStrong, color = FsColor.strong)
            Text(meta, style = FsType.meta, color = FsColor.muted)
        }
    }
}

@Composable
private fun SharedButtons(view: SharedBuildView, canCompare: Boolean, actions: SharedActions) {
    val large = isLargeText()
    val edit = @Composable { modifier: Modifier ->
        BrandButton(stringResource(R.string.plan_edit), BrandButtonKind.Primary, { actions.editCopy(view) }, modifier, height = 48.dp)
    }
    val compare = @Composable { modifier: Modifier ->
        BrandButton(stringResource(R.string.plan_compare), BrandButtonKind.Secondary, { actions.compare(view) }, modifier, height = 48.dp)
    }
    if (large || !canCompare) {
        Column(Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            edit(Modifier)
            if (canCompare) compare(Modifier)
        }
    } else {
        Row(Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 8.dp), horizontalArrangement = Arrangement.spacedBy(10.dp)) {
            edit(Modifier.weight(1f))
            compare(Modifier.weight(1f))
        }
    }
}

/** `{d Mon}`: `4 Nov`. */
val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
