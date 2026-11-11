package gg.foreversixty.app.menu

import android.content.Intent
import android.provider.Settings
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.DrawerState
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.unit.dp
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.guild.logic.GuildRank
import gg.foreversixty.app.push.NotificationsLine
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.account.AccountLinks
import gg.foreversixty.app.account.AccountPage
import gg.foreversixty.app.account.AccountRoute
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.guides.GuidesRoute
import gg.foreversixty.app.guides.logic.GuidesSegment
import gg.foreversixty.app.navigation.ShellFeatures
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.premium.PlanCopy
import gg.foreversixty.app.premium.PlanSummary
import gg.foreversixty.app.rankings.RankingsReports
import gg.foreversixty.app.rankings.RankingsRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.reports.ReportsRoute
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.sims.SimsRoute
import gg.foreversixty.app.system.offline.OfflineData
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.talents.TalentsModel
import gg.foreversixty.app.talents.TalentCharacter
import gg.foreversixty.app.talents.TalentsRoute
import gg.foreversixty.app.common.NewBuildStrip
import kotlinx.coroutines.launch
import java.time.Instant
import java.time.Year

/** The profile menu as a Material 3 modal drawer, 85% of the width, over the tabs and the tab bar. */
@Composable
fun ProfileMenuDrawer(
    features: ShellFeatures,
    drawer: DrawerState,
    gesturesEnabled: Boolean,
    onOpen: (MenuDestination) -> Unit,
    content: @Composable () -> Unit,
) {
    val scope = rememberCoroutineScope()
    val close = { scope.launch { drawer.close() } }
    ModalNavigationDrawer(
        drawerState = drawer,
        gesturesEnabled = gesturesEnabled || drawer.isOpen,
        scrimColor = FsColor.bg.copy(alpha = SCRIM_ALPHA),
        drawerContent = {
            ModalDrawerSheet(
                drawerState = drawer,
                drawerContainerColor = FsColor.raised,
                // Material 3's drawer: rounded at its open edge (round 5, mock m-android-menu).
                drawerShape = androidx.compose.material3.DrawerDefaults.shape,
                modifier = Modifier.fillMaxWidth(WIDTH).fillMaxHeight(),
                windowInsets = androidx.compose.foundation.layout.WindowInsets(0.dp),
            ) {
                MenuContent(features.account, drawer.isOpen) { action ->
                    close()
                    action(onOpen)
                }
            }
        },
        content = content,
    )
}

/** The menu bound to the session; [act] closes the drawer and runs the row's action. */
@Composable
private fun MenuContent(account: AccountDeps?, open: Boolean, act: ((((MenuDestination) -> Unit)) -> Unit) -> Unit) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val me = (phase as? SessionPhase.SignedIn)?.me
    val model = me?.let { MenuModel.of(it, session.currentCharacter(phase, picked), Instant.now()) }
    var devices by remember { mutableStateOf<List<DeviceRow>?>(null) }
    var notificationGroups by remember { mutableStateOf<Map<PushGroup, Boolean>?>(null) }
    var refresh by remember { mutableStateOf(MenuRefresh.Idle) }
    LaunchedEffect(open, me != null) {
        if (open && session != null && me != null) {
            devices = (session.devices() as? ApiResult.Success)?.data
            // Round 7 row 4: the account's per-group choices, read each time the drawer opens.
            (session.notificationSettings() as? ApiResult.Success)?.let { notificationGroups = NotificationsLine.account(it.data) }
        }
    }
    val offline = remember(account) { account?.offlineData?.invoke() }
    val offlineSize by produceState<String?>(null, open, offline) {
        if (open && offline != null) {
            offline.load()
            value = offline.report?.let { OfflineData.size(it.totalBytes) }
        }
    }
    // The account's choices and the channels' real switches, read each time the drawer opens.
    val officer = model?.you?.guild?.rank?.let { GuildRank.of(it) != GuildRank.Member } == true
    val res = LocalResources.current
    val notifications = remember(open, officer, res, notificationGroups) { NotificationsLine.line(res, context, notificationGroups, officer) }
    val lines = MenuLines(
        notifications = notifications,
        plan = me?.let { PlanCopy.line(PlanSummary.of(it, Instant.now()).row) },
        devices = devices?.size,
        newestDevice = devices?.firstOrNull { !it.isThisDevice }?.name,
        offlineData = offlineSize,
        version = account?.let { AccountCopy.version(it.versionName, it.versionCode, it.dataBuild()) }.orEmpty(),
        year = Year.now().value,
    )
    ProfileMenu(
        model = model,
        lines = lines,
        actions = MenuActions(
            onSwitch = { character -> act { session?.pick(character.key) } },
            onRefresh = {
                if (session != null && refresh != MenuRefresh.Refreshing) {
                    scope.launch { MenuRefresh.run({ refresh = it }, session::refresh) }
                }
            },
            onOpen = { destination -> act { open -> open(destination) } },
            onSignIn = { act { account?.let { it.session.beginSignIn(it.browser(context)) } } },
            onHelp = { act { AccountLinks.open(context, AccountLinks.HELP) } },
            onPrivacy = { act { AccountLinks.open(context, AccountLinks.PRIVACY) } },
        ),
        refresh = refresh,
    )
}

/** A menu row's screen, pushed under the current tab; [onBack] returns to the tab's root. */
@Composable
fun MenuScreen(features: ShellFeatures, destination: MenuDestination, onPickClass: () -> Unit, onBack: () -> Unit) {
    val account = features.account
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val me = (phase as? SessionPhase.SignedIn)?.me
    val current = phase?.let { session.currentCharacter(it, picked) }
    val context = LocalContext.current
    val signIn = { account?.let { it.session.beginSignIn(it.browser(context)) } ?: Unit }
    val reportsViewer = me?.let { ReportsViewer.of(it, current?.key) } ?: ReportsViewer.SignedOut
    val scope = rememberCoroutineScope()
    when (destination) {
        MenuDestination.Guides, MenuDestination.Tiers, MenuDestination.PatchNotes -> features.guides?.let { guides ->
            // The row's segment is saved before the root reads it on its first frame.
            remember(destination) { segmentOf(destination)?.also { GuidesSegment.save(guides.segment, it) } }
            GuidesRoute(
                guides,
                current,
                openBis = { guide ->
                    features.bis?.request(BisSelection.forGuide(guide.classSlug, guide.gear?.spec ?: guide.spec, current?.faction))
                    onPickClass()
                },
                planner = features.planner,
            ) { BackButton(onBack) }
        }
        MenuDestination.Sims -> features.home.sims?.let { SimsRoute(it, me?.entitlements?.serverSims == true, onBack) }
        MenuDestination.Logs -> features.home.reports?.let { ReportsRoute(it, reportsViewer, signIn, onBack) }
        MenuDestination.Rankings -> features.home.rankings?.let { rankings ->
            val viewer = me?.let { RankingsViewer.of(it, current?.key) } ?: RankingsViewer.SignedOut
            RankingsRoute(rankings, viewer, features.home.reports?.let { RankingsReports(it, reportsViewer) }, onBack)
        }
        MenuDestination.Notifications -> {
            // Each group is its own channel (PushGroup), so Android's own screen sets them until
            // the in-app settings screen lands with the restyle.
            LaunchedEffect(Unit) {
                context.startActivity(Intent(Settings.ACTION_APP_NOTIFICATION_SETTINGS).putExtra(Settings.EXTRA_APP_PACKAGE, context.packageName))
                onBack()
            }
        }
        // Owner ruling 2026-11-08: the Talents compare page lives under the Talent planner row.
        MenuDestination.Planner -> {
            val talents = features.home.talents
            if (talents != null && current != null) {
                val character = remember(current.key) { HomeCharacter.from(current) }
                val model = remember(current.key) { TalentsModel(talents.source, TalentCharacter.from(character), talents.catalog) }
                LaunchedEffect(model) { model.load() }
                Column(Modifier.fillMaxSize()) {
                    NewBuildStrip(model) { scope.launch { model.load() } }
                    TalentsRoute(character, model, talents.planner, onBack)
                }
            } else {
                LaunchedEffect(Unit) { onBack() }
            }
        }
        MenuDestination.Plan -> account?.let { AccountRoute(it, AccountPage.Plan, onPickClass, onBack) }
        MenuDestination.Devices -> account?.let { AccountRoute(it, AccountPage.Devices, onPickClass, onBack) }
        MenuDestination.OfflineData -> account?.let { AccountRoute(it, AccountPage.OfflineData, onPickClass, onBack) }
        MenuDestination.Settings -> account?.let { AccountRoute(it, AccountPage.Settings, onPickClass, onBack) }
    }
}

@Composable
private fun BackButton(onBack: () -> Unit) {
    IconButton(onClick = onBack) {
        Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.common_back), tint = FsColor.strong)
    }
}

private fun segmentOf(destination: MenuDestination): GuidesSegment? = when (destination) {
    MenuDestination.Guides -> GuidesSegment.Guides
    MenuDestination.Tiers -> GuidesSegment.Tiers
    MenuDestination.PatchNotes -> GuidesSegment.Changes
    else -> null
}

private const val WIDTH = 0.85f
private const val SCRIM_ALPHA = 0.6f
