package gg.foreversixty.app.navigation

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Badge
import androidx.compose.material3.BadgedBox
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.NavigationBarItemDefaults
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.saveable.rememberSaveableStateHolder
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Density
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.account.HomeAccountGate
import gg.foreversixty.app.activity.ActivityFeature
import gg.foreversixty.app.gear.flows.GearDeps
import gg.foreversixty.app.gear.flows.GearOwner
import gg.foreversixty.app.gear.flows.GearRoute
import gg.foreversixty.app.activity.ActivityLanding
import gg.foreversixty.app.activity.ActivityTab
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.guides.GuidesFeature
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.home.HomeFlows
import gg.foreversixty.app.home.flows.HomeExits
import androidx.compose.ui.platform.LocalUriHandler
import gg.foreversixty.app.home.HomeRoute
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.LocalItemPageOpener
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.MenuDestination
import gg.foreversixty.app.menu.ProfileMenuDrawer
import gg.foreversixty.app.menu.MenuScreen
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.planner.PlannerFeature
import gg.foreversixty.app.planner.PlannerScreenHost
import gg.foreversixty.app.planner.logic.PlannerScreen
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.RaidsReports
import gg.foreversixty.app.raids.RaidsRoute
import gg.foreversixty.app.rankings.RankingsRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.reports.ReportsFeature
import gg.foreversixty.app.reports.logic.ReportsViewer
import gg.foreversixty.app.search.SearchExits
import gg.foreversixty.app.search.SearchFeature
import gg.foreversixty.app.search.SearchOverlay
import gg.foreversixty.app.search.SearchOwner
import gg.foreversixty.app.search.ui.RootSearchBar
import gg.foreversixty.app.system.SystemNavigator
import kotlinx.coroutines.launch

/** What the shell draws; built once in MainActivity from the app graph. */
data class ShellFeatures(
    val home: HomeDependencies,
    val bis: BisModel? = null,
    val account: AccountDeps? = null,
    val raids: RaidsFeature? = null,
    val search: SearchFeature? = null,
    val guides: GuidesFeature? = null,
    /** Lane F15: where a launcher shortcut or the Assistant asked to go. */
    val system: SystemNavigator? = null,
    val planner: PlannerFeature? = null,
    val activity: ActivityFeature? = null,
    /** Opens an app link the way MainActivity routes one (a notification, an Activity row). */
    val openLink: (String) -> Unit = {},
)

/**
 * The app shell (spec 2026-11-08 section 3): four tabs, Home, Raid, Gear and Activity, under the
 * profile menu (a Material 3 modal drawer over the tabs). A menu row pushes its screen under the tab
 * the menu was opened on; Back returns to that tab's root.
 */
@Composable
fun ForeverAppScaffold(features: ShellFeatures) {
    var selected by rememberSaveable { mutableStateOf(ForeverTab.Home) }
    var pushed by rememberSaveable { mutableStateOf<MenuDestination?>(null) }
    val drawer = rememberDrawerState(DrawerValue.Closed)
    val scope = rememberCoroutineScope()
    val select = { tab: ForeverTab ->
        selected = tab
        pushed = null
        // Your builds and a shared build sit over the tab like a menu screen; a tab tap leaves them.
        features.planner?.navigator?.close()
        Unit
    }
    FollowRequests(features, select) { pushed = it }
    // Lane F13: the planner's editor covers the shell; the tabs' saved state outlives it.
    val plannerTop = features.planner?.navigator?.top
    if (features.planner != null && plannerTop is PlannerScreen.Editor) {
        // The editor is a full-screen task with no tab bar (spec F13.1).
        PlannerScreenHost(features.planner, plannerTop)
        return
    }
    BackHandler(enabled = pushed != null) { pushed = null }
    val openMenu = { scope.launch { drawer.open() }; Unit }
    val onOpen = { destination: MenuDestination ->
        if (destination == MenuDestination.Planner) features.planner?.open(DeepLink.Builds) else pushed = destination
        Unit
    }
    ProfileMenuDrawer(features, drawer, gesturesEnabled = pushed == null && plannerTop == null, onOpen = onOpen) {
        ShellBody(features, selected, select, pushed, openMenu) { pushed = null }
        val search = features.search
        if (search != null) {
            SearchOverlay(search, searchOwner(features.account), searchExits(features, select), features.home.rankings?.searchLink(), features.guides?.let { g -> { link -> g.navigator.open(link) } })
        }
    }
}

/** The tab bar and the current tab's content: the pushed menu screen, else the tab's root. */
@Composable
private fun ShellBody(
    features: ShellFeatures,
    selected: ForeverTab,
    select: (ForeverTab) -> Unit,
    pushed: MenuDestination?,
    openMenu: () -> Unit,
    closePushed: () -> Unit,
) {
    val search = features.search
    val onSearch = search?.let { feature -> { feature.navigator.openSearch() } }
    val openItemPage = search?.let { feature -> { look: ItemLook -> feature.navigator.openItem(look, selected.label) } }
    val searchOpen = search?.navigator?.stack?.isNotEmpty() == true
    val badges = TabBadges(raid = features.raids?.badge?.count ?: 0, activity = features.activity?.unread ?: 0)
    val plannerTop = features.planner?.navigator?.top
    val shell = rememberSaveableStateHolder()
    CompositionLocalProvider(LocalItemPageOpener provides openItemPage) {
        // While Search covers the app, the tabs under it are hidden from TalkBack.
        Box(if (searchOpen) Modifier.fillMaxSize().clearAndSetSemantics {} else Modifier.fillMaxSize()) {
            ForeverTabScaffold(selected = selected, onSelect = select, badges = badges) {
                when {
                    // Lane F13: a shared build and Your builds are pushed over the current tab.
                    features.planner != null && plannerTop != null -> PlannerScreenHost(features.planner, plannerTop)
                    pushed != null -> MenuScreen(features, pushed, onPickClass = { select(ForeverTab.Gear) }, onBack = closePushed)
                    else -> shell.SaveableStateProvider(selected.name) {
                        CompositionLocalProvider(LocalOpenMenu provides openMenu) { TabRoot(features, selected, select, onSearch) }
                    }
                }
            }
        }
    }
}

/** A tab's root screen. Home and Gear keep their lanes' screens until their own rebuilds. */
@Composable
private fun TabRoot(features: ShellFeatures, tab: ForeverTab, select: (ForeverTab) -> Unit, onSearch: (() -> Unit)?) {
    val account = features.account
    val session = account?.session
    val openGear = { select(ForeverTab.Gear) }
    val context = LocalContext.current
    val uri = LocalUriHandler.current
    val flows = HomeFlows(
        raids = features.raids?.takeIf { it.enabled },
        activity = features.activity,
        exits = HomeExits(
            openActivity = { select(ForeverTab.Activity) },
            openRaidTab = { select(ForeverTab.Raid) },
            openLanding = { landing -> openLanding(features, landing, select) },
            openWeb = { path -> uri.openUri(if (path.startsWith("http")) path else SITE + path) },
        ),
    )
    when (tab) {
        ForeverTab.Home -> when {
            account == null -> HomeRoute(features.home, openBis = openGear, onSearch = onSearch) { Unavailable(tab) }
            publicRankings(features.home, account) -> PublicRankings(features.home, account)
            else -> HomeAccountGate(account, openGear) {
                HomeRoute(features.home, session, openGear, onSearch, signIn = { account.session.beginSignIn(account.browser(context)) }, flows = flows) { Unavailable(tab) }
            }
        }
        ForeverTab.Raid -> features.raids?.let { raids ->
            val title = stringResource(R.string.tab_raid)
            RaidTab(raids, account, features.home.reports) { RootSearchBar(onSearch) { ProfileBar(session, title) } }
        } ?: Unavailable(tab)
        ForeverTab.Gear -> features.bis?.let { GearTab(it, features.home, account, onSearch) } ?: Unavailable(tab)
        ForeverTab.Activity -> features.activity?.let { activity ->
            val title = stringResource(R.string.tab_activity)
            ActivityTab(activity, account, onOpen = { landing -> openLanding(features, landing, select) }) { ProfileBar(session, title) }
        } ?: Unavailable(tab)
    }
}

private fun openLanding(features: ShellFeatures, landing: ActivityLanding, select: (ForeverTab) -> Unit) = when (landing) {
    is ActivityLanding.Link -> features.openLink(landing.link.url())
    is ActivityLanding.Tab -> select(landing.tab)
    is ActivityLanding.Menu -> features.system?.wantsMenu = landing.destination
}

/** Shortcut, link and notification requests: the tab they name, or the menu screen (Guides). */
@Composable
private fun FollowRequests(features: ShellFeatures, select: (ForeverTab) -> Unit, push: (MenuDestination) -> Unit) {
    val system = features.system
    val wantsTab = system?.wantsTab
    val wantsMenu = system?.wantsMenu
    LaunchedEffect(wantsTab, wantsMenu) {
        wantsTab?.let(select)
        wantsMenu?.let(push)
        system?.wantsTab = null
        system?.wantsMenu = null
    }
    // Lane F9: a guide link pushes Guides, where the guide is on top.
    val guides = features.guides
    val wantsGuides = guides?.navigator?.wantsTab == true
    LaunchedEffect(wantsGuides) {
        if (wantsGuides) {
            push(MenuDestination.Guides)
            guides.navigator.wantsTab = false
        }
    }
    // Lane F4: a raid link or notification selects the Raid tab.
    val raids = features.raids
    val wantsRaids = raids?.navigator?.wantsTab == true
    LaunchedEffect(wantsRaids) {
        if (wantsRaids) {
            select(ForeverTab.Raid)
            raids.navigator.wantsTab = false
        }
    }
    // Lanes F5, F10 and F11: a sim, report, rankings or character link selects Home, where its screens are pushed.
    val home = features.home
    val wantsHome = home.sims?.navigator?.wantsHome == true || home.reports?.navigator?.wantsHome == true ||
        home.rankings?.navigator?.wantsHome == true
    LaunchedEffect(wantsHome) {
        if (wantsHome) {
            select(ForeverTab.Home)
            home.sims?.navigator?.wantsHome = false
            home.reports?.navigator?.wantsHome = false
            home.rankings?.navigator?.wantsHome = false
        }
    }
}

/** Search's way out to a list: Gear on that list. */
private fun searchExits(features: ShellFeatures, select: (ForeverTab) -> Unit): SearchExits = SearchExits { asked ->
    val bis = features.bis
    if (bis != null && asked != null) {
        bis.request(if (asked.faction.isEmpty()) asked.copy(faction = bis.currentFaction) else asked)
    }
    select(ForeverTab.Gear)
}

/** Gear (spec 2026-11-08 4.3) for the session's current character. */
@Composable
private fun GearTab(bis: BisModel, home: HomeDependencies, account: AccountDeps?, onSearch: (() -> Unit)?) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val character = phase?.let { session.currentCharacter(it, picked) }
    val waiting = phase == SessionPhase.Importing || (phase is SessionPhase.SignedIn && phase.me == null)
    val deps = remember(home) { GearDeps(home.catalog, home.cache, home.client) }
    val title = stringResource(R.string.tab_gear)
    GearRoute(bis, deps, GearOwner(character?.key, character?.guild?.id, waiting, character?.avatarUrl), onSearch) { ProfileBar(session, title) }
}

/** Raid (lane F4) for the session's current character's guild. */
@Composable
private fun RaidTab(raids: RaidsFeature, account: AccountDeps?, reports: ReportsFeature?, bar: @Composable () -> Unit) {
    val session = account?.session
    val context = LocalContext.current
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val characterKey = phase?.let { session.currentCharacter(it, picked)?.key }
    // Lane F10: Your raid night's `Open the full report` pushes the report on this stack.
    val me = (phase as? SessionPhase.SignedIn)?.me
    val reportsLink = reports?.let { RaidsReports(it, me?.let { m -> ReportsViewer.of(m, characterKey) } ?: ReportsViewer.SignedOut) }
    RaidsRoute(raids, characterKey, reportsLink, signIn = { account?.let { it.session.beginSignIn(it.browser(context)) } }, selector = bar)
}

/**
 * Lane F15 (follow-up of F11): a rankings or player link opens for a signed-out player and for one
 * with no character. Rankings and character pages are public; the board simply has no `Your rank`.
 */
@Composable
private fun publicRankings(home: HomeDependencies, account: AccountDeps): Boolean {
    val nav = home.rankings?.navigator ?: return false
    val phase = account.session.phase.collectAsStateWithLifecycle().value
    val wantsRankings = nav.showsBoard || nav.character != null
    val gated = phase !is SessionPhase.SignedIn || phase.me?.characters?.isEmpty() == true
    return wantsRankings && gated
}

@Composable
private fun PublicRankings(home: HomeDependencies, account: AccountDeps) {
    val rankings = home.rankings ?: return
    val me = account.session.phase.collectAsStateWithLifecycle().value.me
    val viewer = me?.let { RankingsViewer.of(it, null) } ?: RankingsViewer.SignedOut
    RankingsRoute(rankings, viewer, null) { rankings.navigator.showsBoard = false }
}

/** Search's owner: the session's current character and the account's character keys. */
@Composable
private fun searchOwner(account: AccountDeps?): SearchOwner {
    val session = account?.session ?: return SearchOwner()
    val phase = session.phase.collectAsStateWithLifecycle().value
    val picked = session.pickedKey.collectAsStateWithLifecycle().value
    return SearchOwner(session.currentCharacter(phase, picked), phase.me?.characters?.map { it.key }?.toSet().orEmpty())
}

/** The scaffold with the four-tab bar and its badges, stateless so screenshot tests can drive it. */
@Composable
fun ForeverTabScaffold(
    selected: ForeverTab,
    onSelect: (ForeverTab) -> Unit,
    badges: TabBadges = TabBadges(),
    content: @Composable () -> Unit,
) {
    Scaffold(
        containerColor = MaterialTheme.colorScheme.background,
        bottomBar = {
            NavigationBar(containerColor = FsColor.navBar) {
                ForeverTab.entries.forEach { tab ->
                    val count = badges.of(tab)
                    val badge = if (count > 0) stringResource(R.string.tab_badge, count) else null
                    NavigationBarItem(
                        selected = tab == selected,
                        onClick = { onSelect(tab) },
                        icon = {
                            BadgedBox(badge = { if (count > 0) TabBadge(count) }) {
                                Icon(painterResource(tab.icon), contentDescription = null, modifier = Modifier.size(TAB_ICON_SIZE))
                            }
                        },
                        label = { TabLabel(stringResource(tab.label)) },
                        // Spec "Tab-bar icons": gold #ffd36b when selected, muted white otherwise.
                        colors = NavigationBarItemDefaults.colors(
                            selectedIconColor = FsColor.tabOn,
                            selectedTextColor = FsColor.tabOn,
                            unselectedIconColor = FsColor.tabOff,
                            unselectedTextColor = FsColor.tabOff,
                            indicatorColor = FsColor.gold.copy(alpha = INDICATOR_ALPHA),
                        ),
                        modifier = if (badge != null) Modifier.semantics { contentDescription = badge } else Modifier,
                    )
                }
            }
        },
    ) { padding ->
        Box(Modifier.fillMaxSize().padding(padding)) { content() }
    }
}

/** The site a notice's path opens on. */
private const val SITE = "https://foreversixty.gg"

/** The largest font scale a tab label grows to; past it the label stays one line (spec §10). */
private const val TAB_LABEL_MAX_FONT_SCALE = 1.3f

/** The active tab's pill: their gold at 18% (round 5's mock, `.android .tab.on .pill`). */
private const val INDICATOR_ALPHA = 0.18f

/** The owner's tab icons at 24 dp (spec "Tab-bar icons"). */
private val TAB_ICON_SIZE = 24.dp

/** A tab label on one line whatever the font scale: the scale is capped, the icon carries the rest. */
@Composable
private fun TabLabel(text: String) {
    CappedTabScale { Text(text, style = FsType.tab, maxLines = 1, softWrap = false) }
}

/** A tab's unread count: capped as the label is, so at large text it never outgrows the icon it sits on. */
@Composable
private fun TabBadge(count: Int) {
    CappedTabScale {
        Badge(containerColor = FsColor.wipe, contentColor = FsColor.bg) { Text(count.toString(), style = FsType.tab, maxLines = 1) }
    }
}

/** The tab bar's text at most [TAB_LABEL_MAX_FONT_SCALE]: the bar's height is fixed, its words fit it. */
@Composable
private fun CappedTabScale(content: @Composable () -> Unit) {
    val density = LocalDensity.current
    CompositionLocalProvider(
        LocalDensity provides Density(density.density, density.fontScale.coerceAtMost(TAB_LABEL_MAX_FONT_SCALE)),
        content = content,
    )
}

/** A tab whose feature switch is off. */
@Composable
private fun Unavailable(tab: ForeverTab) {
    Box(Modifier.fillMaxSize(), contentAlignment = androidx.compose.ui.Alignment.Center) {
        Text(
            text = stringResource(R.string.shell_arrivesNext, stringResource(tab.label)),
            style = MaterialTheme.typography.bodyLarge,
            color = MaterialTheme.colorScheme.onSurfaceVariant,
        )
    }
}
