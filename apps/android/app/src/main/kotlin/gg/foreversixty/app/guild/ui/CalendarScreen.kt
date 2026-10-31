package gg.foreversixty.app.guild.ui

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
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.KeyboardArrowDown
import androidx.compose.material3.BasicAlertDialog
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Surface
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Switch
import androidx.compose.material3.SwitchDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.stateDescription
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.guild.logic.CalendarAlert
import gg.foreversixty.app.guild.logic.CalendarLinkNotice
import gg.foreversixty.app.guild.logic.CalendarLinkState
import gg.foreversixty.app.guild.logic.CalendarSettings
import gg.foreversixty.app.guild.logic.CalendarStatus
import gg.foreversixty.app.guild.logic.CalendarWhich
import gg.foreversixty.app.raids.ui.RaidSecondaryButton
import gg.foreversixty.app.raids.ui.SectionLabel
import gg.foreversixty.app.raids.ui.isLargeText
import gg.foreversixty.app.raids.ui.raidQuietCard

/** Everything F16.4 draws, already resolved. */
data class CalendarScreenState(
    val status: CalendarStatus,
    val settings: CalendarSettings,
    val link: CalendarLinkState,
    val linkNotice: CalendarLinkNotice? = null,
    val offline: Boolean = false,
    val resetting: Boolean = false,
)

/** What the screen's controls do. */
class CalendarActions(
    val onSwitch: (Boolean) -> Unit,
    val onWhich: (CalendarWhich) -> Unit,
    val onAlert: (CalendarAlert) -> Unit,
    val openSettings: () -> Unit,
    val copy: () -> Unit,
    val reset: () -> Unit,
)

/** F16.4 and F16.5: You › Calendar. */
@Composable
fun CalendarContent(state: CalendarScreenState, actions: CalendarActions, back: () -> Unit, confirmingReset: Boolean = false) {
    var confirming by remember { mutableStateOf(confirmingReset) }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        GuildTopBar(stringResource(R.string.calendar_title), back)
        Column(
            verticalArrangement = Arrangement.spacedBy(8.dp),
            modifier = Modifier.fillMaxSize().verticalScroll(rememberScrollState()).padding(16.dp),
        ) {
            if (state.status == CalendarStatus.AccessOff) {
                AccessOffCard(actions.openSettings)
                Box(Modifier.heightIn(min = 14.dp))
            }
            SectionLabel(stringResource(R.string.calendar_section))
            SyncCard(state, actions)
            Box(Modifier.heightIn(min = 14.dp))
            SectionLabel(stringResource(R.string.calendar_share))
            LinkCard(state, actions.copy) { confirming = true }
            Text(
                stringResource(R.string.calendar_footer),
                style = FsType.meta,
                color = FsColor.muted,
                modifier = Modifier.padding(top = 6.dp, bottom = 16.dp),
            )
        }
    }
    if (confirming) {
        ResetDialog(
            onReset = {
                confirming = false
                actions.reset()
            },
            onDismiss = { confirming = false },
        )
    }
}

/**
 * The Material 3 alert, laid out by hand so the buttons sit 8 dp off the bottom edge (the stock
 * AlertDialog leaves its 24 dp padding under a 40 dp button row): the question in the body face,
 * Cancel and the destructive Reset.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun ResetDialog(onReset: () -> Unit, onDismiss: () -> Unit) {
    BasicAlertDialog(onDismissRequest = onDismiss) {
        Surface(shape = RoundedCornerShape(DIALOG_RADIUS), color = FsColor.raised, tonalElevation = 6.dp) {
            Column(Modifier.padding(start = 24.dp, end = 24.dp, top = 24.dp, bottom = 8.dp)) {
                Text(stringResource(R.string.calendar_resetAsk), style = FsType.lead, color = FsColor.strong)
                Row(
                    horizontalArrangement = Arrangement.spacedBy(8.dp, Alignment.End),
                    modifier = Modifier.fillMaxWidth().padding(top = 16.dp),
                ) {
                    TextButton(onDismiss) { Text(stringResource(R.string.common_cancel), style = FsType.bodyStrong, color = FsColor.nav) }
                    TextButton(onReset) { Text(stringResource(R.string.calendar_resetConfirm), style = FsType.bodyStrong, color = FsColor.wipe) }
                }
            }
        }
    }
}

@Composable
private fun AccessOffCard(openSettings: () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(10.dp), modifier = Modifier.raidQuietCard().padding(16.dp)) {
        Text(stringResource(R.string.calendar_off), style = FsType.bodyStrong, color = FsColor.wipe)
        Text(stringResource(R.string.calendar_offBody), style = FsType.meta, color = FsColor.muted)
        RaidSecondaryButton(stringResource(R.string.calendar_openSettings), openSettings)
    }
}

@Composable
private fun SyncCard(state: CalendarScreenState, actions: CalendarActions) {
    val accessOff = state.status == CalendarStatus.AccessOff
    val on = state.settings.enabled && !accessOff && state.status != CalendarStatus.Removed
    Column(Modifier.raidQuietCard()) {
        Row(
            horizontalArrangement = Arrangement.spacedBy(12.dp),
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth().heightIn(min = 64.dp).padding(horizontal = 16.dp, vertical = 12.dp),
        ) {
            Column(verticalArrangement = Arrangement.spacedBy(2.dp), modifier = Modifier.weight(1f)) {
                Text(stringResource(R.string.calendar_switch), style = FsType.body, color = FsColor.strong)
                Text(stringResource(R.string.calendar_switchNote), style = FsType.meta, color = FsColor.muted)
                when (state.status) {
                    CalendarStatus.Removed -> Text(stringResource(R.string.calendar_removed), style = FsType.meta, color = FsColor.gold)
                    CalendarStatus.Failed -> Text(stringResource(R.string.calendar_syncFailed), style = FsType.meta, color = FsColor.wipe)
                    else -> Unit
                }
            }
            Switch(
                checked = on,
                onCheckedChange = actions.onSwitch,
                enabled = !accessOff,
                colors = SwitchDefaults.colors(
                    checkedTrackColor = FsColor.gold,
                    checkedThumbColor = FsColor.bg,
                    checkedBorderColor = FsColor.gold,
                    uncheckedTrackColor = FsColor.raised,
                    uncheckedThumbColor = FsColor.muted,
                    uncheckedBorderColor = FsColor.muted,
                    disabledUncheckedTrackColor = FsColor.raised,
                    disabledUncheckedThumbColor = FsColor.muted,
                    disabledUncheckedBorderColor = FsColor.muted,
                ),
                modifier = Modifier.alpha(if (accessOff) DISABLED_ALPHA else 1f),
            )
        }
        HorizontalDivider(color = FsColor.lineSoft)
        ValueMenu(
            label = stringResource(R.string.calendar_which),
            current = state.settings.which,
            options = CalendarWhich.entries,
            word = { stringResource(it.label) },
            onPick = actions.onWhich,
        )
        HorizontalDivider(color = FsColor.lineSoft)
        ValueMenu(
            label = stringResource(R.string.calendar_alert),
            current = state.settings.alert,
            options = CalendarAlert.entries,
            word = { stringResource(it.label) },
            onPick = actions.onAlert,
        )
    }
}

/** A settings row whose muted value opens a menu (spec 7, Roster sort menu's pattern). */
@Composable
private fun <T> ValueMenu(label: String, current: T, options: List<T>, word: @Composable (T) -> String, onPick: (T) -> Unit) {
    var open by remember { mutableStateOf(false) }
    val value = word(current)
    Box {
        val large = isLargeText()
        val modifier = Modifier
            .fillMaxWidth()
            .heightIn(min = 56.dp)
            .clickable { open = true }
            .padding(horizontal = 16.dp, vertical = 12.dp)
            .semantics(mergeDescendants = true) {
                role = Role.DropdownList
                stateDescription = value
            }
        if (large) {
            // Largest text: the value under the label.
            Column(verticalArrangement = Arrangement.spacedBy(2.dp), modifier = modifier) {
                Text(label, style = FsType.body, color = FsColor.strong)
                MenuValue(value)
            }
        } else {
            Row(verticalAlignment = Alignment.CenterVertically, modifier = modifier) {
                Text(label, style = FsType.body, color = FsColor.strong, modifier = Modifier.weight(1f))
                MenuValue(value)
            }
        }
        DropdownMenu(expanded = open, onDismissRequest = { open = false }, containerColor = GameColors.sheet) {
            options.forEach { option ->
                DropdownMenuItem(
                    text = { Text(word(option), style = FsType.body, color = FsColor.strong) },
                    trailingIcon = if (option == current) {
                        { Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.gold) }
                    } else {
                        null
                    },
                    onClick = {
                        open = false
                        onPick(option)
                    },
                )
            }
        }
    }
}

@Composable
private fun MenuValue(value: String) {
    Row(horizontalArrangement = Arrangement.spacedBy(4.dp), verticalAlignment = Alignment.CenterVertically) {
        Text(value, style = FsType.body, color = FsColor.muted)
        Icon(Icons.Filled.KeyboardArrowDown, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(16.dp))
    }
}

@Composable
private fun LinkCard(state: CalendarScreenState, copy: () -> Unit, askReset: () -> Unit) {
    val writable = !state.offline && state.link != CalendarLinkState.Loading
    val shown = state.link is CalendarLinkState.Shown
    Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.raidQuietCard().padding(16.dp)) {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.CenterVertically) {
            GoldTile(painterResource(R.drawable.ic_link))
            Column(verticalArrangement = Arrangement.spacedBy(2.dp), modifier = Modifier.weight(1f)) {
                Text(stringResource(R.string.calendar_link), style = FsType.bodyStrong, color = FsColor.strong)
                Text(stringResource(R.string.calendar_linkNote), style = FsType.meta, color = FsColor.muted)
            }
        }
        if (state.link == CalendarLinkState.Elsewhere) {
            Text(stringResource(R.string.calendar_linkElsewhere), style = FsType.meta, color = FsColor.text)
        }
        val buttons = @Composable { modifier: Modifier ->
            if (shown) {
                BrandButton(stringResource(R.string.calendar_copy), BrandButtonKind.Secondary, copy, modifier, enabled = writable, height = 48.dp)
            }
            BrandButton(
                stringResource(R.string.calendar_reset),
                BrandButtonKind.Secondary,
                askReset,
                modifier,
                enabled = writable && !state.resetting,
                busy = state.resetting,
                height = 48.dp,
            )
        }
        if (isLargeText()) {
            Column(verticalArrangement = Arrangement.spacedBy(10.dp)) { buttons(Modifier.fillMaxWidth()) }
        } else {
            Row(horizontalArrangement = Arrangement.spacedBy(10.dp)) { buttons(Modifier.weight(1f)) }
        }
        state.linkNotice?.let { notice ->
            val (text, color) = when (notice) {
                CalendarLinkNotice.Copied -> R.string.calendar_copied to FsColor.kill
                CalendarLinkNotice.LoadFailed -> R.string.calendar_linkError to FsColor.wipe
                CalendarLinkNotice.ResetFailed -> R.string.calendar_resetError to FsColor.wipe
            }
            Text(stringResource(text), style = FsType.meta, color = color)
        }
    }
}

private const val DISABLED_ALPHA = 0.5f

/** Material 3's dialog corner. */
private val DIALOG_RADIUS = 28.dp
