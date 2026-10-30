package gg.foreversixty.app.account

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Phone
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing

/** What F1.6 shows: the rows, or the load failure. */
sealed interface DevicesState {
    data object Loading : DevicesState
    data class Loaded(val rows: List<DeviceRow>) : DevicesState
    data object Failed : DevicesState
}

/** F1.6 You › Devices, stateless over its state apart from the open confirmation. */
@Composable
fun DevicesScreen(
    state: DevicesState,
    failedRevoke: String?,
    onRevoke: (DeviceRow) -> Unit,
    onRetry: () -> Unit,
    initiallyConfirming: DeviceRow? = null,
) {
    var confirming by remember { mutableStateOf(initiallyConfirming) }
    AccountPage(stringResource(R.string.devices_title)) {
        when (state) {
            DevicesState.Loading -> AccountCard {
                repeat(3) { AccountRow(Icons.Filled.Phone, " ", subtitle = " ") }
            }
            is DevicesState.Loaded -> AccountCard {
                state.rows.forEachIndexed { index, row ->
                    DeviceItem(row, failedRevoke == row.id) { confirming = row }
                    if (index < state.rows.lastIndex) HorizontalDivider(color = FsColor.lineSoft)
                }
            }
            DevicesState.Failed -> Column {
                Text(
                    stringResource(R.string.account_devices_loadFailed),
                    style = AccountType.body,
                    color = FsColor.wipe,
                    modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite },
                )
                BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, onRetry, height = 44.dp)
            }
        }
        Text(stringResource(R.string.devices_footer), style = AccountType.meta, color = FsColor.muted)
    }
    confirming?.let { row ->
        ConfirmDialog(
            title = stringResource(R.string.devices_revoke_title, row.name),
            body = stringResource(R.string.devices_revoke_body),
            action = stringResource(R.string.devices_revoke),
            onConfirm = {
                confirming = null
                onRevoke(row)
            },
            onDismiss = { confirming = null },
        )
    }
}

@Composable
private fun DeviceItem(row: DeviceRow, revokeFailed: Boolean, onRevoke: () -> Unit) {
    Column {
        AccountRow(Icons.Filled.Phone, AccountCopy.deviceName(row), subtitle = AccountCopy.deviceMeta(row)) {
            if (row.canRevoke) {
                BrandButton(stringResource(R.string.devices_revoke), BrandButtonKind.Destructive, onRevoke, height = 36.dp, fill = false)
            }
        }
        if (revokeFailed) {
            Text(
                stringResource(R.string.devices_revoke_failed),
                style = AccountType.meta,
                color = FsColor.wipe,
                modifier = Modifier.padding(start = FsSpacing.space4, end = FsSpacing.space4, bottom = FsSpacing.space3),
            )
        }
    }
}
