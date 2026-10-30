package gg.foreversixty.app.account

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
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
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Delete
import androidx.compose.material.icons.filled.Person
import androidx.compose.material.icons.filled.Phone
import androidx.compose.material.icons.filled.Star
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing

/** What the You tab's account section shows, already worded (spec F8.1's head; F1's rows). */
data class YouSummary(
    val battletag: String,
    val plan: String,
    val founder: Boolean,
    /** `This phone and 2 more`, null until the device list answers. */
    val devices: String?,
)

/**
 * The You tab's account section: the account head, Plan, Devices, Delete account and Sign out.
 * F7 and F8 add their rows. Stateless apart from the sign-out confirmation.
 */
@Composable
fun YouScreen(
    summary: YouSummary,
    onDevices: () -> Unit,
    onDeleteAccount: () -> Unit,
    onSignOut: () -> Unit,
    confirmingSignOut: Boolean = false,
) {
    var confirming by rememberSaveable { mutableStateOf(confirmingSignOut) }
    AccountPage(stringResource(R.string.you_title)) {
        Row(horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3), verticalAlignment = Alignment.CenterVertically) {
            Tile(Icons.Filled.Person, 44.dp)
            Column {
                Text(summary.battletag, style = AccountType.name, color = FsColor.strong)
                Text(stringResource(R.string.you_signedInWith), style = AccountType.meta, color = FsColor.muted)
            }
        }
        AccountCard {
            AccountRow(Icons.Filled.Star, stringResource(R.string.you_plan), trailing = summary.plan, tag = stringResource(R.string.plan_founder).takeIf { summary.founder })
            HorizontalDivider(color = FsColor.lineSoft)
            AccountRow(Icons.Filled.Phone, stringResource(R.string.you_devices), subtitle = summary.devices, chevron = true, onClick = onDevices)
        }
        AccountCard {
            AccountRow(Icons.Filled.Delete, stringResource(R.string.account_delete_row), chevron = true, tint = FsColor.wipe, onClick = onDeleteAccount)
        }
        BrandButton(stringResource(R.string.you_signOut), BrandButtonKind.Quiet, { confirming = true }, height = 48.dp)
    }
    if (confirming) {
        ConfirmDialog(
            title = stringResource(R.string.you_signOut_confirm, stringResource(R.string.account_deviceWord)),
            body = null,
            action = stringResource(R.string.account_signOut_action),
            onConfirm = {
                confirming = false
                onSignOut()
            },
            onDismiss = { confirming = false },
        )
    }
}

/** A pushed or root account page: its title, then the content on the ground, scrolling. */
@Composable
fun AccountPage(title: String, content: @Composable ColumnScope.() -> Unit) {
    Column(
        Modifier
            .fillMaxSize()
            .background(FsColor.bg)
            .verticalScroll(rememberScrollState())
            .padding(FsSpacing.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpacing.space5),
    ) {
        Text(title.uppercase(), style = AccountType.title, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        content()
    }
}

@Composable
fun AccountCard(content: @Composable ColumnScope.() -> Unit) {
    val shape = RoundedCornerShape(FsRadius.panel)
    Column(Modifier.fillMaxWidth().background(FsColor.raised, shape).border(1.dp, FsColor.lineSoft, shape), content = content)
}

@Composable
private fun Tile(icon: ImageVector, size: androidx.compose.ui.unit.Dp, tint: Color = FsColor.gold) {
    Box(
        Modifier.size(size).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill)),
        contentAlignment = Alignment.Center,
    ) { Icon(icon, contentDescription = null, tint = tint) }
}

/** One list row: a gold tile, the title and its sub line, a trailing value, tag or chevron. */
@Composable
fun AccountRow(
    icon: ImageVector,
    title: String,
    subtitle: String? = null,
    trailing: String? = null,
    tag: String? = null,
    chevron: Boolean = false,
    tint: Color = FsColor.strong,
    onClick: (() -> Unit)? = null,
    end: @Composable () -> Unit = {},
) {
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 56.dp)
            .let { if (onClick != null) it.clickable(onClick = onClick) else it }
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Tile(icon, 32.dp, if (tint == FsColor.strong) FsColor.gold else tint)
        Column(Modifier.weight(1f)) {
            Text(title, style = AccountType.bodyStrong, color = tint)
            subtitle?.let { Text(it, style = AccountType.meta, color = FsColor.muted) }
        }
        tag?.let {
            Text(
                it.uppercase(),
                style = AccountType.micro,
                color = FsColor.gold,
                modifier = Modifier.border(1.dp, FsColor.gold.copy(alpha = 0.4f), RoundedCornerShape(50)).padding(horizontal = 8.dp, vertical = 3.dp),
            )
        }
        trailing?.let { Text(it, style = AccountType.meta, color = FsColor.muted) }
        end()
        if (chevron) Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

/** The destructive confirmation (spec §3.2: `AlertDialog`). */
@Composable
fun ConfirmDialog(title: String, body: String?, action: String, onConfirm: () -> Unit, onDismiss: () -> Unit) {
    AlertDialog(
        onDismissRequest = onDismiss,
        containerColor = FsColor.raised,
        title = { Text(title, style = AccountType.sec, color = FsColor.strong) },
        text = body?.let { { Text(it, style = AccountType.body, color = FsColor.text) } },
        confirmButton = { TextButton(onConfirm) { Text(action, style = AccountType.bodyStrong, color = FsColor.wipe) } },
        dismissButton = { TextButton(onDismiss) { Text(stringResource(R.string.common_cancel), style = AccountType.bodyStrong, color = FsColor.nav) } },
    )
}
