package gg.foreversixty.app.account

import gg.foreversixty.app.designsystem.screenGround
import gg.foreversixty.app.system.offline.OfflineDataRoute
import gg.foreversixty.app.system.offline.OfflineDataModel
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
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.SessionPhase
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

/** F1.1 bound to the session: the sign-in tap, the spinner and the notice. */
@Composable
fun SignedOutPanel(deps: AccountDeps, onPickClass: () -> Unit, variant: SignedOutVariant = SignedOutVariant.Home) {
    val phase by deps.session.phase.collectAsStateWithLifecycle()
    val context = LocalContext.current
    val signedOutNotice = (phase as? SessionPhase.SignedOut)?.notice
    val notice = signedOutNotice?.let { AccountCopy.notice(it) }
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
        noticeIsFailure = signedOutNotice?.let(AccountCopy::isFailure) ?: true,
    )
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
        phase is SessionPhase.SignedIn -> Box(Modifier.fillMaxSize().screenGround(), Alignment.Center) {
            CircularProgressIndicator(color = FsColor.gold)
        }
        else -> SignedOutPanel(deps, onPickClass)
    }
}

@Composable
private fun SignedInAccount(deps: AccountDeps, start: AccountPage, onBack: () -> Unit) {
    val session = deps.session
    val phase by session.phase.collectAsStateWithLifecycle()
    val picked by session.pickedKey.collectAsStateWithLifecycle()
    val me = phase.let { (it as? SessionPhase.SignedIn)?.me } ?: return
    var page by rememberSaveable(start) { mutableStateOf(start) }
    var deviceCount by remember { mutableStateOf<Int?>(null) }
    val scope = rememberCoroutineScope()
    val context = LocalContext.current
    val version = AccountCopy.version(deps.versionName, deps.versionCode, deps.dataBuild())
    val mayLink by deps.plansLink.mayLink.collectAsStateWithLifecycle()
    val plan = PlanSummary.of(me, Instant.now())
    val offlineData = remember(deps) { deps.offlineData() }
    val back = { if (page == start) onBack() else page = start }
    BackHandler(enabled = page != start) { page = start }
    LaunchedEffect(Unit) {
        deviceCount = (session.devices() as? ApiResult.Success)?.data?.size
    }
    LaunchedEffect(Unit) { deps.plansLink.resolve() }
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
                onPlan = { page = AccountPage.Plan },
                onDevices = { page = AccountPage.Devices },
                onPrivacy = { AccountLinks.open(context, AccountLinks.PRIVACY) },
                onTerms = { AccountLinks.open(context, AccountLinks.TERMS) },
                onAbout = { page = AccountPage.About },
                onDeleteAccount = { page = AccountPage.Delete },
                onSignOut = { scope.launch { session.signOut() } },
                onCalendar = deps.guildExtras?.let { { page = AccountPage.Calendar } },
                onOfflineData = offlineData?.let { { page = AccountPage.OfflineData } },
            ),
            mayLink = mayLink,
            liveSurfaces = deps.liveSurfaces?.let { settings -> { LiveSurfacesCard(settings) } },
            back = back,
        )
        AccountPage.Plan -> PlanScreen(plan, mayLink, back) { AccountLinks.open(context, AccountLinks.BILLING) }
        AccountPage.Devices -> DevicesRoute(session, back) { deviceCount = it }
        AccountPage.Delete -> DeleteRoute(session, me, back)
        AccountPage.About -> AboutScreen(version, back)
        AccountPage.OfflineData -> offlineData?.let { model -> OfflineDataRoute(model) { back() } }
        AccountPage.Calendar -> deps.guildExtras?.let { extras ->
            CalendarRoute(extras, me.user.id.toString(), session.currentCharacter(phase, picked)?.key) { back() }
        }
    }
}

/** Loads `GET /v1/devices`, revokes with `DELETE /v1/devices/{id}`. */
@Composable
private fun DevicesRoute(session: SessionManager, back: () -> Unit, onCount: (Int) -> Unit) {
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
                    AccountActionOutcome.NeedsBrowser -> AccountLinks.open(context, AccountLinks.ACCOUNT)
                    AccountActionOutcome.Failed, AccountActionOutcome.ConfirmMismatch, is AccountActionOutcome.Refused ->
                        failedRevoke = row.id
                }
            }
        },
        onRetry = { scope.launch { load() } },
        back = back,
    )
}

/** Runs `DELETE /v1/me` with what the person typed; on success the session shows signed out. */
@Composable
private fun DeleteRoute(session: SessionManager, me: Me, back: () -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    var typed by rememberSaveable { mutableStateOf("") }
    var deleting by remember { mutableStateOf(false) }
    var outcome by remember { mutableStateOf<AccountActionOutcome?>(null) }
    DeleteAccountScreen(
        consequences = DeletionConsequence.of(me),
        battleTag = me.user.battletag,
        typed = typed,
        onTyped = {
            typed = it
            if (outcome == AccountActionOutcome.ConfirmMismatch) outcome = null
        },
        deleting = deleting,
        outcome = outcome,
        onDelete = {
            deleting = true
            scope.launch {
                val result = session.deleteAccount(typed)
                deleting = false
                outcome = result.takeUnless { it == AccountActionOutcome.Done }
                if (result == AccountActionOutcome.NeedsBrowser) AccountLinks.open(context, AccountLinks.ACCOUNT)
            }
        },
        back = back,
    )
}
