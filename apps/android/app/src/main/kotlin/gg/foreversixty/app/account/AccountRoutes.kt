package gg.foreversixty.app.account

import android.content.Context
import android.content.Intent
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.core.net.toUri
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.BrowserLauncher
import gg.foreversixty.app.auth.session.AccountActionOutcome
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.network.ApiResult
import kotlinx.coroutines.launch

/** The site pages F1 opens in the browser (never an in-app web view). */
object AccountLinks {
    private const val SITE = "https://foreversixty.gg"

    /** Where a `403 browser_session_required` act is done. */
    const val ACCOUNT = "$SITE/account"

    /** F1.3's `Get the addon`. */
    const val SETUP = "$SITE/setup"

    fun open(context: Context, url: String) {
        context.startActivity(Intent(Intent.ACTION_VIEW, url.toUri()))
    }
}

/** What F1 needs from the app graph: the session and the browser sign-in opens in. */
class AccountDeps(val session: SessionManager, val browser: (Context) -> BrowserLauncher)

/**
 * Home's account states around the signed-in home (F1.1 to F1.4 and F1.3's no-characters
 * state). Signed in with characters, it draws [signedIn], the home lane's screen.
 */
@Composable
fun HomeAccountGate(deps: AccountDeps, onPickClass: () -> Unit, signedIn: @Composable () -> Unit) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val context = LocalContext.current
    when (val current = phase) {
        is SessionPhase.SignedOut, SessionPhase.SigningIn -> SignedOutPanel(deps, onPickClass)
        SessionPhase.Importing -> ImportScreen()
        is SessionPhase.SignedIn ->
            if (current.me?.characters?.isEmpty() == true) {
                NoCharactersScreen { AccountLinks.open(context, AccountLinks.SETUP) }
            } else {
                signedIn()
            }
    }
}

/** F1.1 bound to the session: the sign-in tap, the spinner and the notice. */
@Composable
fun SignedOutPanel(deps: AccountDeps, onPickClass: () -> Unit, variant: SignedOutVariant = SignedOutVariant.Home) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val context = LocalContext.current
    val notice = (phase as? SessionPhase.SignedOut)?.notice?.let { AccountCopy.notice(it) }
    SignedOutScreen(
        notice = notice,
        signingIn = phase == SessionPhase.SigningIn,
        onSignIn = {
            deps.session.dismissNotice()
            deps.session.beginSignIn(deps.browser(context))
        },
        onPickClass = {
            deps.session.dismissNotice()
            onPickClass()
        },
        variant = variant,
    )
}

private enum class YouPage { Root, Devices, Delete }

/** The You tab: the account section when signed in, the sign-in panel when not. */
@Composable
fun YouRoute(deps: AccountDeps, onPickClass: () -> Unit) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val me = (phase as? SessionPhase.SignedIn)?.me
    when {
        me != null -> SignedInYou(deps)
        phase is SessionPhase.SignedIn -> Box(Modifier.fillMaxSize().background(FsColor.bg), Alignment.Center) {
            CircularProgressIndicator(color = FsColor.gold)
        }
        else -> SignedOutPanel(deps, onPickClass)
    }
}

@Composable
private fun SignedInYou(deps: AccountDeps) {
    val session = deps.session
    val phase by session.phase.collectAsStateWithLifecycle()
    val me = phase.let { (it as? SessionPhase.SignedIn)?.me } ?: return
    var page by rememberSaveable { mutableStateOf(YouPage.Root) }
    var deviceCount by remember { mutableStateOf<Int?>(null) }
    val scope = rememberCoroutineScope()
    BackHandler(enabled = page != YouPage.Root) { page = YouPage.Root }
    LaunchedEffect(Unit) {
        deviceCount = (session.devices() as? ApiResult.Success)?.data?.size
    }
    when (page) {
        YouPage.Root -> YouScreen(
            summary = YouSummary(
                battletag = me.user.battletag.orEmpty(),
                plan = AccountCopy.planLine(me.entitlements.billing),
                founder = me.entitlements.billing?.founder == true,
                devices = deviceCount?.let {
                    stringResource(R.string.you_devices_sub, stringResource(R.string.account_deviceWord), (it - 1).coerceAtLeast(0))
                },
            ),
            onDevices = { page = YouPage.Devices },
            onDeleteAccount = { page = YouPage.Delete },
            onSignOut = { scope.launch { session.signOut() } },
        )
        YouPage.Devices -> DevicesRoute(session) { deviceCount = it }
        YouPage.Delete -> DeleteRoute(session, hasPlan = me.entitlements.billing != null)
    }
}

/** Loads `GET /v1/devices`, revokes with `DELETE /v1/devices/{id}`. */
@Composable
private fun DevicesRoute(session: SessionManager, onCount: (Int) -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    var state by remember { mutableStateOf<DevicesState>(DevicesState.Loading) }
    var failedRevoke by remember { mutableStateOf<String?>(null) }
    val load: suspend () -> Unit = {
        state = when (val result = session.devices()) {
            is ApiResult.Success -> DevicesState.Loaded(result.data).also { onCount(result.data.size) }
            is ApiResult.Failure -> DevicesState.Failed
        }
    }
    LaunchedEffect(Unit) { load() }
    DevicesScreen(
        state = state,
        failedRevoke = failedRevoke,
        onRevoke = { row ->
            scope.launch {
                when (session.revoke(row.id)) {
                    AccountActionOutcome.Done -> {
                        failedRevoke = null
                        load()
                    }
                    AccountActionOutcome.Failed -> failedRevoke = row.id
                    AccountActionOutcome.NeedsBrowser -> AccountLinks.open(context, AccountLinks.ACCOUNT)
                }
            }
        },
        onRetry = { scope.launch { load() } },
    )
}

/** Runs `DELETE /v1/me` after the dialog; on success the session shows signed out with the date. */
@Composable
private fun DeleteRoute(session: SessionManager, hasPlan: Boolean) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    var deleting by remember { mutableStateOf(false) }
    var failed by remember { mutableStateOf(false) }
    DeleteAccountScreen(hasPlan = hasPlan, deleting = deleting, failed = failed, onDelete = {
        deleting = true
        scope.launch {
            val outcome = session.deleteAccount()
            deleting = false
            failed = outcome == AccountActionOutcome.Failed
            if (outcome == AccountActionOutcome.NeedsBrowser) AccountLinks.open(context, AccountLinks.ACCOUNT)
        }
    })
}
