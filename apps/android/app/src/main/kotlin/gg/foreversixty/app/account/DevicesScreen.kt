package gg.foreversixty.app.account

import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowInsetHead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowStack
import gg.foreversixty.app.flows.FlowStackButton
import gg.foreversixty.app.flows.FlowTag

/** What Devices shows: the rows, or the load failure. */
sealed interface DevicesState {
    data object Loading : DevicesState
    data class Loaded(val rows: List<DeviceRow>) : DevicesState
    data object Failed : DevicesState
}

/** What Devices hands back. */
data class DevicesActions(
    val revoke: (DeviceRow) -> Unit = {},
    val retry: () -> Unit = {},
    val signOutOthers: () -> Unit = {},
)

/**
 * Devices (mock `devices`): Signed in, this device first with its tag, the others with Revoke; Sign out
 * everywhere else; the line on signing this one out. Each revoke asks first. [failedRevoke] is the row
 * whose revoke failed; [othersFailed] says Sign out everywhere else left some signed in.
 */
@Composable
fun DevicesScreen(
    state: DevicesState,
    actions: DevicesActions,
    failedRevoke: String? = null,
    othersFailed: Boolean = false,
    initiallyConfirming: DeviceRow? = null,
) {
    var confirming by remember { mutableStateOf(initiallyConfirming) }
    var confirmingOthers by remember { mutableStateOf(false) }
    FlowPage {
        FlowInsetHead(stringResource(R.string.devices_title))
        when (state) {
            DevicesState.Loading -> FlowSection(stringResource(R.string.devices_section)) {
                repeat(LOADING_ROWS) { FlowSkeleton(SKELETON_WIDTH, SKELETON_HEIGHT, Modifier.padding(vertical = FsSpace.space3)) }
            }
            is DevicesState.Loaded -> {
                FlowSection(stringResource(R.string.devices_section)) {
                    state.rows.forEachIndexed { index, row ->
                        DeviceItem(row, failedRevoke == row.id, last = index == state.rows.lastIndex) { confirming = row }
                    }
                }
                if (state.rows.any { it.canRevoke }) {
                    FlowStack { FlowStackButton(stringResource(R.string.devices_signOutOthers), { confirmingOthers = true }, danger = true) }
                    if (othersFailed) FlowNote(stringResource(R.string.devices_signOutOthers_failed), Modifier.padding(top = FsSpace.space2), FsColor.wipe)
                }
            }
            DevicesState.Failed -> {
                FlowNote(stringResource(R.string.account_devices_loadFailed), Modifier.padding(top = FsSpace.space2), FsColor.wipe)
                FlowStack { FlowStackButton(stringResource(R.string.common_tryAgain), actions.retry) }
            }
        }
        FlowNote(stringResource(R.string.devices_footer), Modifier.padding(top = FsSpace.space4))
    }
    confirming?.let { row ->
        ConfirmDialog(
            title = stringResource(R.string.devices_revoke_title, row.name),
            body = stringResource(R.string.devices_revoke_body),
            action = stringResource(R.string.devices_revoke),
            onConfirm = {
                confirming = null
                actions.revoke(row)
            },
            onDismiss = { confirming = null },
        )
    }
    if (confirmingOthers) {
        ConfirmDialog(
            title = stringResource(R.string.devices_signOutOthers_title),
            body = stringResource(R.string.devices_signOutOthers_body),
            action = stringResource(R.string.devices_signOutOthers),
            onConfirm = {
                confirmingOthers = false
                actions.signOutOthers()
            },
            onDismiss = { confirmingOthers = false },
        )
    }
}

@Composable
private fun DeviceItem(row: DeviceRow, revokeFailed: Boolean, last: Boolean, onRevoke: () -> Unit) {
    FlowRow(
        title = AccountCopy.deviceName(row),
        sub = listOfNotNull(AccountCopy.deviceMeta(row), if (revokeFailed) stringResource(R.string.devices_revoke_failed) else null).joinToString("\n"),
        last = last,
        trail = {
            if (row.canRevoke) {
                FlowButton(stringResource(R.string.devices_revoke), onRevoke, danger = true)
            } else {
                FlowTag(stringResource(R.string.devices_thisTag))
            }
        },
    )
}

private const val LOADING_ROWS = 3
private val SKELETON_WIDTH = 240.dp
private val SKELETON_HEIGHT = 38.dp
