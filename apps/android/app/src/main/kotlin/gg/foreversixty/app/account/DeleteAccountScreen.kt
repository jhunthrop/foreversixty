package gg.foreversixty.app.account

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.OutlinedTextFieldDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.KeyboardCapitalization
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.DeletionConfirmation
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.flows.FlowInsetHead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowNotice
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSealPage
import gg.foreversixty.app.flows.FlowStack
import gg.foreversixty.app.flows.FlowStackButton

private val GOES = listOf(R.string.delete_goes_characters, R.string.delete_goes_guild, R.string.delete_goes_devices)

/** The delete screen's state: what was typed, whether the call runs, the last attempt that did not delete. */
data class DeleteState(
    val consequences: List<DeletionConsequence>,
    val battleTag: String?,
    val typed: String,
    val deleting: Boolean = false,
    val outcome: AccountActionOutcome? = null,
)

/** What the delete screen hands back. */
data class DeleteActions(
    val typed: (String) -> Unit = {},
    val delete: () -> Unit = {},
    val cancel: () -> Unit = {},
)

/**
 * Delete your account (mocks `delete-free`, `delete-premium`, `delete-failed`; App Store 5.1.1(v), the
 * same in the Play app): what goes, the plan's warning, when, what stays, then the typed confirmation
 * (the BattleTag, or delete) and Delete my account, enabled once it matches by the server's rule
 * ([DeletionConfirmation]). An attempt that failed shows the failure with Try again and Cancel.
 */
@Composable
fun DeleteAccountScreen(state: DeleteState, actions: DeleteActions) {
    val failure = state.outcome.takeIf { it == AccountActionOutcome.Failed || it is AccountActionOutcome.Refused }
    FlowPage {
        FlowInsetHead(stringResource(R.string.delete_head))
        if (failure != null) {
            FlowNotice((failure as? AccountActionOutcome.Refused)?.message ?: stringResource(R.string.delete_failed))
            FlowStack {
                FlowStackButton(stringResource(R.string.common_tryAgain), actions.delete, primary = true, enabled = !state.deleting)
                FlowStackButton(stringResource(R.string.common_cancel), actions.cancel)
            }
        } else {
            Form(state, actions)
        }
    }
}

@Composable
private fun Form(state: DeleteState, actions: DeleteActions) {
    Text(stringResource(R.string.delete_lead), style = FsType.lead, color = FsColor.text, modifier = Modifier.padding(top = FsSpace.space2))
    FlowSection(stringResource(R.string.delete_whatGoes)) {
        GOES.forEachIndexed { index, line -> FlowRow(stringResource(line), slim = true, last = index == GOES.lastIndex) }
    }
    state.consequences.forEach { FlowNotice(warning(it), gold = true) }
    FlowSection(stringResource(R.string.delete_when)) {
        FlowNote(stringResource(R.string.delete_whenLine), Modifier.padding(top = FsSpace.space3))
    }
    FlowNote(stringResource(R.string.delete_stays), Modifier.padding(top = FsSpace.space4))
    ConfirmField(state, actions.typed)
    val canDelete = !state.deleting && DeletionConfirmation.matches(state.typed, state.battleTag)
    FlowStack {
        val label = if (state.deleting) R.string.shared_account_delete_busy else R.string.delete_action
        FlowStackButton(stringResource(label), actions.delete, danger = true, enabled = canDelete)
        FlowStackButton(stringResource(R.string.common_cancel), actions.cancel)
    }
    FlowNote(stringResource(R.string.shared_account_delete_email), Modifier.padding(top = FsSpace.space4))
}

/** The plan's line: a store subscription bills again unless it is cancelled in the store first. */
@Composable
private fun warning(consequence: DeletionConsequence): String =
    if (consequence == DeletionConsequence.PremiumPlan) stringResource(R.string.shared_account_delete_storeWarning) else AccountCopy.consequence(consequence)

@Composable
private fun ConfirmField(state: DeleteState, onTyped: (String) -> Unit) {
    val label = AccountCopy.deletionPrompt(state.battleTag)
    val mismatch = state.outcome == AccountActionOutcome.ConfirmMismatch
    Column(Modifier.padding(top = FsSpace.space5).fillMaxWidth()) {
        FlowNote(label)
        OutlinedTextField(
            value = state.typed,
            onValueChange = onTyped,
            singleLine = true,
            enabled = !state.deleting,
            isError = mismatch,
            textStyle = FsType.body.copy(color = FsColor.strong),
            colors = OutlinedTextFieldDefaults.colors(
                focusedBorderColor = FsColor.gold,
                unfocusedBorderColor = FsColor.line,
                cursorColor = FsColor.gold,
            ),
            keyboardOptions = KeyboardOptions(
                capitalization = KeyboardCapitalization.None,
                autoCorrectEnabled = false,
                imeAction = ImeAction.Done,
            ),
            modifier = Modifier.padding(top = FsSpace.space2).fillMaxWidth().semantics { contentDescription = label },
        )
        if (mismatch) FlowNote(AccountCopy.deletionMismatch(state.battleTag), Modifier.padding(top = FsSpace.space2), FsColor.wipe)
    }
}

/** Account deleted (mock `delete-done`): the seal, the line, Done. */
@Composable
fun AccountDeletedScreen(onDone: () -> Unit) {
    FlowSealPage(stringResource(R.string.delete_done_title), stringResource(R.string.delete_done_line)) {
        FlowStackButton(stringResource(R.string.delete_done_action), onDone, primary = true)
    }
}

/** Sign in again (mock `expired`): the Battle.net link expired; the seal, the line, Sign in with Battle.net. */
@Composable
fun SessionExpiredScreen(onSignIn: () -> Unit) {
    FlowSealPage(stringResource(R.string.expired_title), stringResource(R.string.expired_line)) {
        FlowStackButton(stringResource(R.string.signin_button), onSignIn, primary = true)
    }
}
