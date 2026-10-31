package gg.foreversixty.app.widgets

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.selection.toggleable
import androidx.compose.material3.HorizontalDivider
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
import gg.foreversixty.app.account.AccountCard
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.widgets.data.LiveKind
import kotlinx.coroutines.launch

/** You › Live surfaces, reading and writing the shared store (spec §5.6 Rules). */
@Composable
fun LiveSurfacesCard(settings: LiveSurfaceSettings) {
    val scope = rememberCoroutineScope()
    var raid by remember { mutableStateOf(true) }
    var sim by remember { mutableStateOf(true) }
    LaunchedEffect(settings) {
        raid = settings.isOn(LiveKind.Raid)
        sim = settings.isOn(LiveKind.Sim)
    }
    LiveSurfacesSection(
        raidOn = raid, simOn = sim, showsSim = settings.showsSim,
        onRaid = { on -> raid = on; scope.launch { settings.set(LiveKind.Raid, on) } },
        onSim = { on -> sim = on; scope.launch { settings.set(LiveKind.Sim, on) } },
    )
}

/** The card itself: `Live surfaces`, then one switch row each, both on by default. */
@Composable
fun LiveSurfacesSection(raidOn: Boolean, simOn: Boolean, showsSim: Boolean, onRaid: (Boolean) -> Unit, onSim: (Boolean) -> Unit) {
    Column {
        Text(
            stringResource(R.string.live_settings).uppercase(), style = FsType.micro, color = FsColor.muted,
            modifier = Modifier.padding(start = FsSpacing.space4, bottom = FsSpacing.space2),
        )
        AccountCard {
            SwitchRow(stringResource(R.string.live_raidSetting), raidOn, onRaid)
            if (showsSim) {
                HorizontalDivider(color = FsColor.lineSoft)
                SwitchRow(stringResource(R.string.live_simSetting), simOn, onSim)
            }
        }
    }
}

@Composable
private fun SwitchRow(title: String, checked: Boolean, onChange: (Boolean) -> Unit) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp)
            .toggleable(value = checked, role = Role.Switch, onValueChange = onChange)
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Text(title, style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
        Switch(
            checked = checked, onCheckedChange = null,
            colors = SwitchDefaults.colors(checkedTrackColor = FsColor.gold, checkedThumbColor = FsColor.bg),
        )
    }
}
