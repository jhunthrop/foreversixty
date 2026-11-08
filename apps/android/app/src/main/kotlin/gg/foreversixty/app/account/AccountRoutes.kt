package gg.foreversixty.app.account

import gg.foreversixty.app.system.offline.OfflineDataRoute
import gg.foreversixty.app.system.offline.OfflineDataModel
import android.content.Context
import android.content.Intent
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import gg.foreversixty.app.flows.FlowBackBar
import androidx.compose.foundation.layout.Column
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
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.auth.session.SignedOutNotice
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.guild.CalendarRoute
import gg.foreversixty.app.guild.GuildExtrasFeature
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.premium.PlanScreen
import gg.foreversixty.app.premium.PlanSummary
import gg.foreversixty.app.premium.PlansLink
import gg.foreversixty.app.widgets.LiveSurfaceSettings
import gg.foreversixty.app.widgets.LiveSurfacesCard
import kotlinx.coroutines.launch
import java.time.Instant

/** The site pages F1 opens in the browser (never an in-app web view). */
object AccountLinks {
    private const val SITE = "https://foreversixty.gg"

    /** Where a `403 browser_session_required` act is done. */
    const val ACCOUNT = "$SITE/account"

    /** F1.3's `Get the addon`. */
    const val SETUP = "$SITE/setup"

    /** You's Privacy and Terms rows (App Store 5.1.1(i)). */
    const val PRIVACY = "$SITE/privacy"
    const val TERMS = "$SITE/terms"

    /** The menu's Help: the Discord server (web/src/data/links.json `discordInvite`). */
    const val HELP = "https://discord.gg/NZhBtPwvHt"

    /** F8's `See plans`, link regions only (decision P). */
    const val PREMIUM = "$SITE/premium"

    /** F8.4's `Manage billing on foreversixty.gg`, link regions only. */
    const val BILLING = "$SITE/account#plan"

    fun open(context: Context, url: String) {
        context.startActivity(Intent(Intent.ACTION_VIEW, url.toUri()))
    }
}

/**
 * What F1 needs from the app graph: the session, the browser sign-in opens in, and this build's
 * version (apps/VERSION through BuildConfig, lane P2) with the static data build, when known.
 */
class AccountDeps(
    val session: SessionManager,
    val browser: (Context) -> BrowserLauncher,
    val versionName: String,
    val versionCode: Int,
    val dataBuild: () -> String?,
    /** Lane F8: See plans and Manage billing, where the store allows them. */
    val plansLink: PlansLink,
    /** Lane F14: You › Live surfaces; null while the `widgets` switch is off. */
    val liveSurfaces: LiveSurfaceSettings? = null,
    /** Lane F16: You › Calendar; null while `guild_extras` is off. */
    val guildExtras: GuildExtrasFeature? = null,
    /** Lane F15: You › Offline data's model; null while `system` is off. */
    val offlineData: () -> OfflineDataModel? = { null },
)

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

/**
 * F1.1 bound to the session: the sign-in tap, the spinner and the notice. A session that ended draws
 * Sign in again (mock `expired`), a deleted account draws Account deleted (mock `delete-done`).
 */
@Composable
fun SignedOutPanel(deps: AccountDeps, onPickClass: () -> Unit, variant: SignedOutVariant = SignedOutVariant.Home) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val context = LocalContext.current
    val signedOutNotice = (phase as? SessionPhase.SignedOut)?.notice
    val signIn = {
        deps.session.dismissNotice()
        deps.session.beginSignIn(deps.browser(context))
    }
    when (signedOutNotice) {
        SignedOutNotice.SessionEnded -> SessionExpiredScreen(signIn)
        SignedOutNotice.AccountDeleted -> AccountDeletedScreen { deps.session.dismissNotice() }
        else -> SignedOutScreen(
            notice = signedOutNotice?.let { AccountCopy.notice(it) },
            signingIn = phase == SessionPhase.SigningIn,
            onSignIn = signIn,
            onPickClass = {
                deps.session.dismissNotice()
                onPickClass()
            },
            variant = variant,
            noticeIsFailure = signedOutNotice?.let(AccountCopy::isFailure) ?: true,
        )
    }
}

/** The account screens the menu's Account rows open (spec 2026-11-08 section 3); Settings is the old You. */
enum class AccountPage { Settings, Plan, Devices, Delete, About, Calendar, OfflineData }

/**
 * An account screen pushed from the menu, starting at [start]; Back from a screen it pushed returns to
 * [start], Back from [start] calls [onBack]. Signed out it shows the sign-in panel.
 */
@Composable
fun AccountRoute(deps: AccountDeps, start: AccountPage, onPickClass: () -> Unit, onBack: () -> Unit) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val me = (phase as? SessionPhase.SignedIn)?.me
    BackHandler(onBack = onBack)
    when {
        me != null -> SignedInAccount(deps, start, onBack)
        phase is SessionPhase.SignedIn -> Box(Modifier.fillMaxSize().background(FsColor.bg), Alignment.Center) {
            CircularProgressIndicator(color = FsColor.gold)
        }
        else -> SignedOutPanel(deps, onPickClass)
    }
}

@Composable
private fun SignedInAccount(deps: AccountDeps, start: AccountPage, onBack: () -> Unit) {
    val session = deps.session
    val phase by session.phase.collectAsStateWithLifecycle()
    val me = phase.let { (it as? SessionPhase.SignedIn)?.me } ?: return
    var page by rememberSaveable(start) { mutableStateOf(start) }
    var deviceCount by remember { mutableStateOf<Int?>(null) }
    val version = AccountCopy.version(deps.versionName, deps.versionCode, deps.dataBuild())
    val mayLink by deps.plansLink.mayLink.collectAsStateWithLifecycle()
    val plan = PlanSummary.of(me, Instant.now())
    val back = { if (page == start) onBack() else page = start }
    BackHandler(enabled = page != start) { page = start }
    LaunchedEffect(Unit) {
        deviceCount = (session.devices() as? ApiResult.Success)?.data?.size
    }
    LaunchedEffect(Unit) { deps.plansLink.resolve() }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        FlowBackBar(stringResource(if (page == start) R.string.menu_title else R.string.menu_settings), back)
        AccountPageContent(deps, page, me, plan, mayLink, version, deviceCount, AccountPageNav({ page = it }, back), { deviceCount = it })
    }
}

/** How an account page moves: to another page, or back. */
private class AccountPageNav(val go: (AccountPage) -> Unit, val back: () -> Unit)

@Composable
private fun AccountPageContent(
    deps: AccountDeps,
    page: AccountPage,
    me: Me,
    plan: PlanSummary,
    mayLink: Boolean,
    version: String,
    deviceCount: Int?,
    nav: AccountPageNav,
    onCount: (Int) -> Unit,
) {
    val session = deps.session
    val phase by session.phase.collectAsStateWithLifecycle()
    val picked by session.pickedKey.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val offlineData = remember(deps) { deps.offlineData() }
    val back = nav.back
    when (page) {
        AccountPage.Settings -> YouScreen(
            summary = YouSummary(
                battletag = me.user.battletag.orEmpty(),
                plan = plan,
                character = session.currentCharacter(phase, picked)?.let {
                    CharacterLine(it.name, it.className, AccountCopy.place(it))
                },
                devices = deviceCount?.let {
                    stringResource(R.string.you_devices_sub, stringResource(R.string.account_deviceWord), (it - 1).coerceAtLeast(0))
                },
                version = version,
            ),
            actions = YouActions(
                onSeePlans = { AccountLinks.open(context, AccountLinks.PREMIUM) },
                onPlan = { nav.go(AccountPage.Plan) },
                onDevices = { nav.go(AccountPage.Devices) },
                onPrivacy = { AccountLinks.open(context, AccountLinks.PRIVACY) },
                onTerms = { AccountLinks.open(context, AccountLinks.TERMS) },
                onAbout = { nav.go(AccountPage.About) },
                onDeleteAccount = { nav.go(AccountPage.Delete) },
                onSignOut = { scope.launch { session.signOut() } },
                onCalendar = deps.guildExtras?.let { { nav.go(AccountPage.Calendar) } },
                onOfflineData = offlineData?.let { { nav.go(AccountPage.OfflineData) } },
            ),
            mayLink = mayLink,
            liveSurfaces = deps.liveSurfaces?.let { settings -> { LiveSurfacesCard(settings) } },
        )
        AccountPage.Plan -> PlanScreen(plan, mayLink) { AccountLinks.open(context, AccountLinks.BILLING) }
        AccountPage.Devices -> DevicesRoute(session, onCount)
        AccountPage.Delete -> DeleteRoute(session, me) { back() }
        AccountPage.About -> AboutScreen(version)
        AccountPage.OfflineData -> offlineData?.let { model -> OfflineDataRoute(model) { back() } }
        AccountPage.Calendar -> deps.guildExtras?.let { extras ->
            CalendarRoute(extras, me.user.id.toString(), session.currentCharacter(phase, picked)?.key) { back() }
        }
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
    var othersFailed by remember { mutableStateOf(false) }
    val revoke: suspend (DeviceRow) -> AccountActionOutcome = { row ->
        session.revoke(row.id).also { if (it == AccountActionOutcome.NeedsBrowser) AccountLinks.open(context, AccountLinks.ACCOUNT) }
    }
    DevicesScreen(
        state = state,
        failedRevoke = failedRevoke,
        othersFailed = othersFailed,
        actions = DevicesActions(
            revoke = { row ->
                scope.launch {
                    failedRevoke = when (revoke(row)) {
                        AccountActionOutcome.Done -> null
                        AccountActionOutcome.NeedsBrowser -> failedRevoke
                        AccountActionOutcome.Failed, AccountActionOutcome.ConfirmMismatch, is AccountActionOutcome.Refused -> row.id
                    }
                    load()
                }
            },
            retry = { scope.launch { load() } },
            signOutOthers = {
                val others = (state as? DevicesState.Loaded)?.rows.orEmpty().filter { it.canRevoke }
                scope.launch {
                    othersFailed = others.map { revoke(it) }.any { it != AccountActionOutcome.Done }
                    load()
                }
            },
        ),
    )
}

/** Runs `DELETE /v1/me` with what the person typed; on success the session shows Account deleted. */
@Composable
private fun DeleteRoute(session: SessionManager, me: Me, onCancel: () -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    var typed by rememberSaveable { mutableStateOf("") }
    var deleting by remember { mutableStateOf(false) }
    var outcome by remember { mutableStateOf<AccountActionOutcome?>(null) }
    DeleteAccountScreen(
        DeleteState(DeletionConsequence.of(me), me.user.battletag, typed, deleting, outcome),
        DeleteActions(
            typed = {
                typed = it
                if (outcome == AccountActionOutcome.ConfirmMismatch) outcome = null
            },
            delete = {
                deleting = true
                scope.launch {
                    val result = session.deleteAccount(typed)
                    deleting = false
                    outcome = result.takeUnless { it == AccountActionOutcome.Done }
                    if (result == AccountActionOutcome.NeedsBrowser) AccountLinks.open(context, AccountLinks.ACCOUNT)
                }
            },
            cancel = onCancel,
        ),
    )
}
