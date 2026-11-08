package gg.foreversixty.app.account.notifications

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.selection.toggleable
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
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
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.push.PushGroup
import kotlinx.coroutines.launch

/** Notifications, pushed from the menu: bound to the feature, loaded on open. */
@Composable
fun NotificationSettingsRoute(feature: NotificationSettingsFeature) {
    val scope = rememberCoroutineScope()
    LaunchedEffect(feature) { feature.load() }
    NotificationSettingsScreen(
        state = feature.state,
        saveFailed = feature.saveFailed,
        actions = NotificationSettingsActions(
            onGroup = { group, on -> scope.launch { feature.setGroup(group, on) } },
            onReminder = { scope.launch { feature.setReminder(it) } },
            onOfficerDms = { scope.launch { feature.setOfficerDms(it) } },
            onRetry = { scope.launch { feature.load() } },
        ),
    )
}

data class NotificationSettingsActions(
    val onGroup: (PushGroup, Boolean) -> Unit = { _, _ -> },
    val onReminder: (String) -> Unit = {},
    val onOfficerDms: (Boolean) -> Unit = {},
    val onRetry: () -> Unit = {},
)

/**
 * Notifications (mock `notification-settings`): Send me, a switch per group on its title row; When,
 * the raid-day reminder with Change; Discord, the officer DM opt-in (off by default).
 */
@Composable
fun NotificationSettingsScreen(state: NotificationSettingsState, saveFailed: Boolean, actions: NotificationSettingsActions) {
    FlowPage {
        FlowHead(stringResource(R.string.menu_notifications), emptyList())
        when (state) {
            NotificationSettingsState.Loading -> Box(Modifier.fillMaxWidth().padding(FsSpace.space8), Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold)
            }
            NotificationSettingsState.Failed -> Column(Modifier.padding(top = FsSpace.space6)) {
                FlowNote(stringResource(R.string.notif_failed), color = FsColor.wipe)
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.onRetry, height = 44.dp, fill = false)
            }
            is NotificationSettingsState.Loaded -> Loaded(NotificationGroupsSummary.of(state.settings), saveFailed, actions)
        }
    }
}

@Composable
private fun Loaded(summary: NotificationGroupsSummary, saveFailed: Boolean, actions: NotificationSettingsActions) {
    FlowSection(stringResource(R.string.notifications_sendMe)) {
        summary.rows.forEachIndexed { index, row ->
            SwitchRow(stringResource(row.group.title), stringResource(row.group.line), row.on, last = index == summary.rows.lastIndex) {
                actions.onGroup(row.group, it)
            }
        }
        if (saveFailed) FlowNote(stringResource(R.string.common_failed), Modifier.padding(top = FsSpace.space2), FsColor.wipe)
    }
    FlowSection(stringResource(R.string.notifications_when)) {
        FlowRow(
            title = stringResource(R.string.notifications_reminder),
            sub = stringResource(R.string.notifications_reminderLine, NotificationGroupsSummary.clock(summary.reminder)),
            last = true,
            trail = { ReminderPicker(summary.reminder, actions.onReminder) },
        )
    }
    FlowSection(stringResource(R.string.notifications_discord)) {
        SwitchRow(stringResource(R.string.notifications_officerDms), stringResource(R.string.notifications_officerDmsLine), summary.officerDms, last = true) {
            actions.onOfficerDms(it)
        }
    }
}

/** A setting (the mock's `.setrow`): title and line, the switch on the title row; the whole row toggles. */
@Composable
private fun SwitchRow(title: String, line: String, on: Boolean, last: Boolean, set: (Boolean) -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .toggleable(value = on, role = Role.Switch, onValueChange = set)
            .padding(vertical = ROW_PAD),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Column(Modifier.weight(1f)) {
            Text(title, style = FsType.item, color = FsColor.strong)
            Text(line, style = SUB, color = FsColor.muted)
        }
        Spacer(Modifier.width(FsSpace.space3))
        Box(Modifier.heightIn(min = TARGET), contentAlignment = Alignment.Center) {
            Switch(
                checked = on,
                onCheckedChange = null,
                colors = SwitchDefaults.colors(
                    checkedTrackColor = FsColor.gold,
                    checkedThumbColor = FsColor.bg,
                    uncheckedTrackColor = FsColor.raised,
                    uncheckedBorderColor = FsColor.line,
                ),
            )
        }
    }
    if (!last) FlowRule()
}

/** The reminder's Change: a menu of the afternoon and evening hours, the current one marked. */
@Composable
private fun ReminderPicker(hhmm: String, save: (String) -> Unit) {
    var open by remember { mutableStateOf(false) }
    Box {
        FlowLink(stringResource(R.string.notifications_change)) { open = true }
        DropdownMenu(expanded = open, onDismissRequest = { open = false }, containerColor = FsColor.raised) {
            NotificationGroupsSummary.HOURS.forEach { value ->
                val clock = NotificationGroupsSummary.clock(value)
                DropdownMenuItem(
                    text = { Text(if (value == hhmm) "$clock ✓" else clock, style = FsType.body, color = FsColor.strong) },
                    onClick = {
                        open = false
                        if (value != hhmm) save(value)
                    },
                )
            }
        }
    }
}

private val ROW_PAD = 10.dp
private val TARGET = 44.dp
