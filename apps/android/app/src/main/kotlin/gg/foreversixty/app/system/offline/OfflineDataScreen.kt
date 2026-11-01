package gg.foreversixty.app.system.offline

import androidx.activity.compose.BackHandler
import androidx.annotation.DrawableRes
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.toggleable
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Switch
import androidx.compose.material3.SwitchDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.ConfirmDialog
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.guild.ui.GuildTopBar
import java.time.Instant
import kotlinx.coroutines.launch

/** Everything the screen draws, so the screenshot tests draw every state without a model. */
data class OfflineDataState(
    val report: OfflineReport?,
    val keepClassesReady: Boolean,
    val wifiOnly: Boolean,
    val online: Boolean,
    val updating: Boolean,
    val note: UpdateNote?,
    val checkedAt: Instant?,
    val confirmingRemove: Boolean = false,
)

class OfflineDataActions(
    val back: () -> Unit = {},
    val keep: (Boolean) -> Unit = {},
    val wifi: (Boolean) -> Unit = {},
    val updateNow: () -> Unit = {},
    val askRemove: () -> Unit = {},
    val remove: () -> Unit = {},
    val cancelRemove: () -> Unit = {},
)

/** You › Offline data, wired to its model. */
@Composable
fun OfflineDataRoute(model: OfflineDataModel, back: () -> Unit) {
    val scope = rememberCoroutineScope()
    var confirming by remember { mutableStateOf(false) }
    BackHandler(onBack = back)
    LaunchedEffect(model) { model.load() }
    val state = OfflineDataState(
        model.report, model.keepClassesReady, model.wifiOnly, model.online, model.updating, model.note, model.checkedAt, confirming,
    )
    OfflineDataScreen(
        state,
        Instant.now(),
        OfflineDataActions(
            back = back,
            keep = { model.keepClassesReady(it) },
            wifi = { model.wifiOnly(it) },
            updateNow = { scope.launch { model.updateNow() } },
            askRemove = { confirming = true },
            remove = {
                confirming = false
                scope.launch { model.remove() }
            },
            cancelRemove = { confirming = false },
        ),
    )
}

/** The screen (spec F15.4, board F15.4). */
@Composable
fun OfflineDataScreen(state: OfflineDataState, now: Instant, actions: OfflineDataActions) {
    val copy = OfflineDataCopy(LocalResources.current)
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        GuildTopBar(stringResource(R.string.offlineData_title), actions.back)
        Column(
            Modifier.fillMaxSize().verticalScroll(rememberScrollState()).padding(horizontal = FsSpacing.space4).padding(bottom = FsSpacing.space6),
            verticalArrangement = Arrangement.spacedBy(FsSpacing.space4),
        ) {
            val report = state.report
            when {
                report == null -> Unit
                report.isEmpty -> EmptyCard()
                else -> Saved(report, copy, now)
            }
            Updates(state, actions)
            UpdateButtons(state, report?.isEmpty != false, actions)
            Text(copy.footer(state.checkedAt, now), style = FsType.meta, color = FsColor.muted)
        }
    }
    if (state.confirmingRemove) {
        ConfirmDialog(
            title = stringResource(R.string.offlineData_removeConfirm),
            body = null,
            action = stringResource(R.string.offlineData_removeAction),
            onConfirm = actions.remove,
            onDismiss = actions.cancelRemove,
        )
    }
}

@Composable
private fun Saved(report: OfflineReport, copy: OfflineDataCopy, now: Instant) {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
        MicroLabel(copy.section(report))
        Card {
            report.rows.forEachIndexed { index, row ->
                if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                DataRow(iconOf(row.category), copy.title(row.category), copy.meta(row, now))
            }
        }
    }
}

@Composable
private fun EmptyCard() {
    Card {
        Column(Modifier.padding(FsSpacing.space4), verticalArrangement = Arrangement.spacedBy(FsSpacing.space1)) {
            Text(stringResource(R.string.offlineData_empty), style = FsType.bodyStrong, color = FsColor.strong)
            Text(stringResource(R.string.offlineData_emptySub), style = FsType.meta, color = FsColor.muted)
        }
    }
}

@Composable
private fun Updates(state: OfflineDataState, actions: OfflineDataActions) {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
        MicroLabel(stringResource(R.string.offlineData_updates))
        Card {
            SwitchRow(stringResource(R.string.offlineData_keep), stringResource(R.string.offlineData_keepNote), state.keepClassesReady, actions.keep)
            HorizontalDivider(color = FsColor.lineSoft)
            SwitchRow(stringResource(R.string.offlineData_wifi), stringResource(R.string.offlineData_wifiNote), state.wifiOnly, actions.wifi)
        }
    }
}

@Composable
private fun UpdateButtons(state: OfflineDataState, empty: Boolean, actions: OfflineDataActions) {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
        val label = stringResource(if (state.updating) R.string.offlineData_updating else R.string.offlineData_updateNow)
        BrandButton(
            label, BrandButtonKind.Primary, actions.updateNow,
            enabled = state.online && !state.updating, height = 48.dp, leadingIcon = Icons.Filled.Refresh,
        )
        val reason = when {
            !state.online -> stringResource(R.string.offline_needsConnection, stringResource(R.string.offlineData_updateNow))
            state.note == UpdateNote.UpToDate -> stringResource(R.string.offlineData_upToDate)
            state.note == UpdateNote.Failed -> stringResource(R.string.offline_updateFailed)
            else -> null
        }
        reason?.let {
            Text(
                it, style = FsType.meta, color = if (state.note == UpdateNote.Failed) FsColor.wipe else FsColor.muted,
                modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite },
            )
        }
        if (!empty) BrandButton(stringResource(R.string.offlineData_remove), BrandButtonKind.Destructive, actions.askRemove, height = 48.dp)
    }
}

@Composable
private fun Card(content: @Composable ColumnScope.() -> Unit) {
    Column(Modifier.fillMaxWidth().quietCard(), content = content)
}

@Composable
private fun DataRow(@DrawableRes icon: Int, title: String, meta: String) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp).padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Box(
            Modifier.size(32.dp).background(FsColor.gold.copy(alpha = TILE_ALPHA), RoundedCornerShape(FsRadius.pill)),
            contentAlignment = Alignment.Center,
        ) { Icon(painterResource(icon), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp)) }
        Column(Modifier.weight(1f)) {
            Text(title, style = FsType.bodyStrong, color = FsColor.strong)
            Text(meta, style = FsType.meta, color = FsColor.muted)
        }
    }
}

@Composable
private fun SwitchRow(title: String, note: String, checked: Boolean, onChange: (Boolean) -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp)
            .toggleable(value = checked, role = Role.Switch, onValueChange = onChange)
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Column(Modifier.weight(1f)) {
            Text(title, style = FsType.bodyStrong, color = FsColor.strong)
            Text(note, style = FsType.meta, color = FsColor.muted)
        }
        Switch(checked = checked, onCheckedChange = null, colors = SwitchDefaults.colors(checkedTrackColor = FsColor.gold, checkedThumbColor = FsColor.bg))
    }
}

@DrawableRes
private fun iconOf(category: OfflineCategory): Int = when (category) {
    OfflineCategory.Guides -> R.drawable.ic_guide
    OfflineCategory.Bis -> R.drawable.ic_shield
    OfflineCategory.Talents -> R.drawable.ic_talents
    OfflineCategory.Search -> R.drawable.ic_search
    OfflineCategory.Icons -> R.drawable.ic_grid
}

private const val TILE_ALPHA = 0.12f
