package gg.foreversixty.app.account

import gg.foreversixty.app.designsystem.Label12
import gg.foreversixty.app.designsystem.SectionTitle
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.outlined.Info
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.KeyboardCapitalization
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.DeletionConfirmation
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.ScaledType

private val DELETED_ITEMS = listOf(
    R.string.shared_account_delete_deleted_subscription,
    R.string.shared_account_delete_deleted_identity,
    R.string.shared_account_delete_deleted_discord,
    R.string.shared_account_delete_deleted_private,
    R.string.shared_account_delete_deleted_guild,
)

private val KEPT_ITEMS = listOf(
    R.string.shared_account_delete_kept_billing,
    R.string.shared_account_delete_kept_reports,
    R.string.shared_account_delete_kept_characters,
    R.string.shared_account_delete_kept_builds,
)

/**
 * You › Delete account (App Store guideline 5.1.1(v), and the same in the Play app): what is
 * deleted and what is kept, in the web dialog's words; what else ends for this account (its plan,
 * founder price, seat); then the typed confirmation (the BattleTag, or delete) and the destructive
 * button, enabled only once the field matches by the server's rule ([DeletionConfirmation]).
 * [outcome] is the last attempt's when it did not delete; null before the first try.
 */
@Composable
fun DeleteAccountScreen(
    consequences: List<DeletionConsequence>,
    battleTag: String?,
    typed: String,
    onTyped: (String) -> Unit,
    deleting: Boolean,
    outcome: AccountActionOutcome?,
    onDelete: () -> Unit,
    /** Back to the screen under this one: the title moves into the band bar (round 5). */
    back: (() -> Unit)? = null,
) {
    val lines = consequences.map { AccountCopy.consequence(it) }
    val canDelete = !deleting && DeletionConfirmation.matches(typed, battleTag)
    AccountPage(stringResource(R.string.shared_account_delete_title), back) {
        Text(stringResource(R.string.shared_account_delete_intro), style = ScaledType.body, color = FsColor.strong)
        BulletGroup(stringResource(R.string.shared_account_delete_deleted_title), DELETED_ITEMS.map { stringResource(it) }, Icons.Filled.Close)
        BulletGroup(stringResource(R.string.shared_account_delete_kept_title), KEPT_ITEMS.map { stringResource(it) }, Icons.Filled.Check, FsColor.muted)
        if (lines.isNotEmpty()) BulletGroup(stringResource(R.string.shared_account_delete_lose_title), lines, Icons.Outlined.Info)
        // The confirmation on one Plate: the field, its failure in wipe and the button read on the panel's dark.
        AccountCard {
            Column(Modifier.padding(FsSpace.space4), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
                ConfirmField(battleTag, typed, onTyped, enabled = !deleting, mismatch = outcome == AccountActionOutcome.ConfirmMismatch)
                AccountCopy.deletionFailure(outcome)?.let { Text(it, style = ScaledType.body, color = FsColor.wipe) }
                BrandButton(stringResource(R.string.shared_account_delete_button), BrandButtonKind.Destructive, onDelete, busy = deleting, enabled = canDelete)
            }
        }
        Text(stringResource(R.string.shared_account_delete_email), style = ScaledType.meta, color = FsColor.mutedGround)
    }
}

@Composable
private fun ConfirmField(battleTag: String?, typed: String, onTyped: (String) -> Unit, enabled: Boolean, mismatch: Boolean) {
    val label = AccountCopy.deletionPrompt(battleTag)
    Column(verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
        Label12(label, color = FsColor.text)
        OutlinedTextField(
            value = typed,
            onValueChange = onTyped,
            singleLine = true,
            enabled = enabled,
            isError = mismatch,
            keyboardOptions = KeyboardOptions(
                capitalization = KeyboardCapitalization.None,
                autoCorrectEnabled = false,
                imeAction = ImeAction.Done,
            ),
            modifier = Modifier.fillMaxWidth().semantics { contentDescription = label },
        )
        if (mismatch) {
            Text(AccountCopy.deletionMismatch(battleTag), style = ScaledType.meta, color = FsColor.wipe)
        }
    }
}

@Composable
private fun BulletGroup(title: String, lines: List<String>, glyph: ImageVector, tint: Color = FsColor.wipe) {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
        SectionTitle(title)
        AccountCard {
            Column(Modifier.padding(FsSpace.space4), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
                lines.forEach { line ->
                    Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
                        Icon(glyph, contentDescription = null, tint = tint, modifier = Modifier.size(16.dp))
                        Text(line, style = ScaledType.body, color = FsColor.strong)
                    }
                }
            }
        }
    }
}
