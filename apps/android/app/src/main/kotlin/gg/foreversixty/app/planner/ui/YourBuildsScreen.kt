package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Add
import androidx.compose.material3.Icon
import androidx.compose.material3.SnackbarHost
import androidx.compose.material3.SnackbarHostState
import androidx.compose.material3.SwipeToDismissBox
import androidx.compose.material3.SwipeToDismissBoxValue
import androidx.compose.material3.Text
import androidx.compose.material3.rememberSwipeToDismissBoxState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.foundation.shape.RoundedCornerShape
import gg.foreversixty.app.designsystem.FsRadius
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.CustomAccessibilityAction
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.customActions
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.planner.logic.BuildRow
import gg.foreversixty.app.planner.logic.BuildsProblem
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.YourBuildsPhase
import gg.foreversixty.app.reports.ui.Chevron
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.ReportsBar
import java.time.ZoneId

/** `Plan a new build`'s row: the current character's class and level. */
data class NewBuildRow(val classSlug: String, val className: String, val level: Int)

data class BuildsPageState(
    val phase: YourBuildsPhase,
    val newBuild: NewBuildRow?,
    val zone: ZoneId = ZoneId.systemDefault(),
)

class BuildsActions(
    val back: () -> Unit = {},
    val open: (BuildRow) -> Unit = {},
    val remove: (BuildRow) -> Unit = {},
    val planNew: (NewBuildRow) -> Unit = {},
    val retry: () -> Unit = {},
)

/** F13.6 Your builds: the rows (swipe to remove), the footer, then `Plan a new build`. */
@Composable
fun YourBuildsPage(state: BuildsPageState, actions: BuildsActions, snackbar: SnackbarHostState = remember { SnackbarHostState() }) {
    Box(Modifier.fillMaxSize().background(FsColor.bg)) {
        Column(Modifier.fillMaxSize()) {
            ReportsBar(stringResource(R.string.plan_builds_title), stringResource(R.string.talents_title), actions.back)
            when (val phase = state.phase) {
                YourBuildsPhase.Loading -> Unit
                YourBuildsPhase.Failed -> ErrorRegion(stringResource(R.string.plan_builds_failed), actions.retry)
                is YourBuildsPhase.Loaded -> BuildsList(phase, state, actions)
            }
        }
        SnackbarHost(snackbar, Modifier.align(Alignment.BottomCenter).padding(16.dp))
    }
}

@Composable
private fun BuildsList(phase: YourBuildsPhase.Loaded, state: BuildsPageState, actions: BuildsActions) {
    LazyColumn(Modifier.fillMaxSize(), contentPadding = androidx.compose.foundation.layout.PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        if (phase.rows.isEmpty()) {
            item { Text(stringResource(R.string.plan_builds_empty), style = FsType.body, color = FsColor.muted) }
        }
        items(phase.rows, key = { it.build.id }) { row -> SwipeRow(row, state.zone, actions) }
        phase.problem?.let { problem ->
            item {
                val text = when (problem) {
                    BuildsProblem.RemoveFailed -> R.string.plan_removeFailed
                    BuildsProblem.UndoFailed -> R.string.common_failed
                }
                Text(stringResource(text), style = FsType.body, color = FsColor.wipe)
            }
        }
        if (phase.rows.isNotEmpty()) {
            item { Text(stringResource(R.string.plan_builds_footer), style = FsType.meta, color = FsColor.muted, modifier = Modifier.padding(top = 4.dp)) }
        }
        state.newBuild?.let { new ->
            item { MicroLabel(stringResource(R.string.plan_builds_newHeader), Modifier.padding(top = 18.dp)) }
            item { NewBuild(new, actions) }
        }
    }
}

/** Swipe-to-dismiss (end to start) takes a row off the list; TalkBack gets `Remove` as an action. */
@Composable
private fun SwipeRow(row: BuildRow, zone: ZoneId, actions: BuildsActions) {
    val swipe = rememberSwipeToDismissBoxState()
    val removeLabel = stringResource(R.string.plan_builds_remove)
    SwipeToDismissBox(
        state = swipe,
        enableDismissFromStartToEnd = false,
        onDismiss = { value -> if (value == SwipeToDismissBoxValue.EndToStart) actions.remove(row) },
        backgroundContent = {
            // Drawn only while a swipe is under way, so nothing shows round a resting card's corners.
            if (swipe.dismissDirection == SwipeToDismissBoxValue.EndToStart) {
                Box(Modifier.fillMaxSize().background(FsColor.wipe.copy(alpha = REMOVE_ALPHA)).padding(end = 20.dp), contentAlignment = Alignment.CenterEnd) {
                    Text(removeLabel, style = FsType.bodyStrong, color = FsColor.strong)
                }
            }
        },
        modifier = Modifier.clip(RoundedCornerShape(FsRadius.panel)).semantics { customActions = listOf(CustomAccessibilityAction(removeLabel) { actions.remove(row); true }) },
    ) {
        BuildRowView(row, zone) { actions.open(row) }
    }
}

@Composable
private fun BuildRowView(row: BuildRow, zone: ZoneId, open: () -> Unit) {
    val copy = PlannerCopy(LocalResources.current)
    val totals = row.totals?.let(copy::totals) ?: row.build.pointOrder.size.toString()
    val date = row.savedAt?.let { DAY_MONTH.withZone(zone).format(it) }.orEmpty()
    val heading = stringResource(R.string.plan_builds_rowTitle, row.raceName, row.className, totals)
    Row(
        Modifier
            .fillMaxWidth()
            .quietCard()
            .clickable(role = Role.Button, onClick = open)
            .heightIn(min = 64.dp)
            .padding(horizontal = 14.dp, vertical = 10.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ClassCrest(row.classSlug, 36.dp)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            val title = row.build.title
            if (title != null) Text(title, style = FsType.bodyStrong, color = FsColor.strong)
            else Text(heading, style = FsType.bodyStrong, color = FsColor.strong)
            Text(stringResource(R.string.plan_builds_meta, row.raceName, row.className, totals, date), style = FsType.meta, color = FsColor.muted)
        }
        val views = row.build.views.toInt()
        Text(pluralStringResource(R.plurals.plan_builds_views, views, views), style = FsType.monoMeta, color = FsColor.text)
        Chevron(20.dp)
    }
}

@Composable
private fun NewBuild(new: NewBuildRow, actions: BuildsActions) {
    Row(
        Modifier
            .fillMaxWidth()
            .quietCard()
            .clickable(role = Role.Button) { actions.planNew(new) }
            .heightIn(min = 56.dp)
            .padding(horizontal = 14.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        Icon(Icons.Filled.Add, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(22.dp))
        Text(stringResource(R.string.plan_builds_newRow, new.className, new.level), style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Chevron(20.dp)
    }
}

private const val REMOVE_ALPHA = 0.35f
