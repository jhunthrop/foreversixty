package gg.foreversixty.app.account

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.outlined.Info
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing

private val DELETED_ITEMS = listOf(
    R.string.account_delete_item_identity,
    R.string.account_delete_item_characters,
    R.string.account_delete_item_logs,
    R.string.account_delete_item_sims,
    R.string.account_delete_item_links,
)

/**
 * You › Delete account (App Store guideline 5.1.1(v), and the same in the Play app): what is
 * deleted in the privacy page's words; what else ends for this account (its plan, founder price,
 * guild plan, seat); what happens to guild reports; the destructive button and an explicit
 * confirmation that repeats the consequences.
 */
@Composable
fun DeleteAccountScreen(
    consequences: List<DeletionConsequence>,
    deleting: Boolean,
    failed: Boolean,
    onDelete: () -> Unit,
    initiallyConfirming: Boolean = false,
) {
    var confirming by rememberSaveable { mutableStateOf(initiallyConfirming) }
    val lines = consequences.map { AccountCopy.consequence(it) }
    AccountPage(stringResource(R.string.account_delete_title)) {
        Text(stringResource(R.string.account_delete_intro), style = AccountType.body, color = FsColor.text)
        BulletCard(DELETED_ITEMS.map { stringResource(it) }, Icons.Filled.Close)
        if (lines.isNotEmpty()) {
            Text(stringResource(R.string.account_delete_lose_title).uppercase(), style = AccountType.micro, color = FsColor.muted)
            BulletCard(lines, Icons.Outlined.Info)
        }
        // Pending the owner's decision on guild reports: one key, account.delete.reports.
        Text(stringResource(R.string.account_delete_reports), style = AccountType.body, color = FsColor.text)
        Text(stringResource(R.string.account_delete_timing), style = AccountType.body, color = FsColor.text)
        if (failed) Text(stringResource(R.string.account_delete_failed), style = AccountType.body, color = FsColor.wipe)
        BrandButton(stringResource(R.string.account_delete_button), BrandButtonKind.Destructive, { confirming = true }, busy = deleting)
        Text(stringResource(R.string.account_delete_email), style = AccountType.meta, color = FsColor.muted)
    }
    if (confirming) {
        ConfirmDialog(
            title = stringResource(R.string.account_delete_confirm_title),
            body = (listOf(stringResource(R.string.account_delete_confirm_body)) + lines).joinToString("\n\n"),
            action = stringResource(R.string.account_delete_button),
            onConfirm = {
                confirming = false
                onDelete()
            },
            onDismiss = { confirming = false },
        )
    }
}

@Composable
private fun BulletCard(lines: List<String>, glyph: ImageVector) {
    AccountCard {
        Column(Modifier.padding(FsSpacing.space4), verticalArrangement = Arrangement.spacedBy(FsSpacing.space3)) {
            lines.forEach { line ->
                Row(horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3)) {
                    Icon(glyph, contentDescription = null, tint = FsColor.wipe, modifier = Modifier.size(16.dp))
                    Text(line, style = AccountType.body, color = FsColor.strong)
                }
            }
        }
    }
}
