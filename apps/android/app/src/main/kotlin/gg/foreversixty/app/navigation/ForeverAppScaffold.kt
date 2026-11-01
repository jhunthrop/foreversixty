package gg.foreversixty.app.navigation

import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.unit.Density
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.account.CharacterSelectorBar
import gg.foreversixty.app.account.HomeAccountGate
import gg.foreversixty.app.account.YouRoute
import gg.foreversixty.app.auth.session.SessionPhase
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.BisOwner
import gg.foreversixty.app.bis.BisRoute
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.guides.GuidesFeature
import gg.foreversixty.app.guides.GuidesRoute
import gg.foreversixty.app.home.HomeRoute
import gg.foreversixty.app.planner.PlannerFeature
import gg.foreversixty.app.planner.PlannerScreenHost
import gg.foreversixty.app.planner.logic.PlannerScreen
import androidx.compose.runtime.saveable.rememberSaveableStateHolder
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.RaidsReports
import gg.foreversixty.app.raids.RaidsRoute
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.semantics.clearAndSetSemantics
import gg.foreversixty.app.items.LocalItemPageOpener
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.search.SearchExits
import gg.foreversixty.app.search.SearchFeature
import gg.foreversixty.app.search.SearchOverlay
import gg.foreversixty.app.search.SearchOwner
import gg.foreversixty.app.search.ui.RootSearchBar
import kotlinx.coroutines.launch
import gg.foreversixty.app.reports.ReportsFeature
import gg.foreversixty.app.reports.logic.ReportsViewer

/**
 * The app shell: a Material 3 navigation bar over one screen per tab. [account] is lane F1's
 * session and sign-in, null while its feature switch is off.
 */
@Composable
fun ForeverAppScaffold(
    home: HomeDependencies,
    bis: BisModel? = null,
    account: AccountDeps? = null,
    raids: RaidsFeature? = null,
    search: SearchFeature? = null,
    guides: GuidesFeature? = null,
    planner: PlannerFeature? = null,
) {
    var selected by rememberSaveable { mutableStateOf(ForeverTab.Home) }
    // Lane F13: the planner's screens sit over the shell; the shell's saved state outlives them.
    val shell = rememberSaveableStateHolder()
    val plannerTop = planner?.navigator?.top
    if (planner != null && plannerTop is PlannerScreen.Editor) {
        // The editor is a full-screen task with no tab bar (spec F13.1).
        PlannerScreenHost(planner, plannerTop)
        return
    }
    val openBis = { selected = ForeverTab.Bis }
    // Lane F9: a guide link selects the Guides tab, where the guide is pushed.
    val wantsGuides = guides?.navigator?.wantsTab == true
    LaunchedEffect(wantsGuides) {
        if (wantsGuides) {
            selected = ForeverTab.Guides
            guides.navigator.wantsTab = false
        }
    }
    val scope = rememberCoroutineScope()
    // Lane F12: the root bars' search button and the tooltips' `More about this item` row.
    val onSearch = search?.let { feature -> { feature.navigator.openSearch() } }
    val openItemPage = search?.let { feature -> { look: ItemLook -> feature.navigator.openItem(look, selected.label) } }
    val exits = SearchExits { asked ->
        if (bis != null && asked != null) {
            val selection = if (asked.faction.isEmpty()) asked.copy(faction = bis.currentFaction) else asked
            if (selected == ForeverTab.Bis) scope.launch { bis.select(selection) } else bis.request(selection)
        }
        selected = ForeverTab.Bis
    }
    // Lane F4: a raid link or notification selects the Raids tab.
    val wantsRaids = raids?.navigator?.wantsTab == true
    LaunchedEffect(wantsRaids) {
        if (wantsRaids) {
            selected = ForeverTab.Raids
            raids.navigator.wantsTab = false
        }
    }
    // Lanes F5, F10 and F11: a sim, report, rankings or character link selects Home, where its screens are pushed.
    val wantsHome = home.sims?.navigator?.wantsHome == true || home.reports?.navigator?.wantsHome == true ||
        home.rankings?.navigator?.wantsHome == true
    LaunchedEffect(wantsHome) {
        if (wantsHome) {
            selected = ForeverTab.Home
            home.sims?.navigator?.wantsHome = false
            home.reports?.navigator?.wantsHome = false
            home.rankings?.navigator?.wantsHome = false
        }
    }
    val context = LocalContext.current
    val searchOpen = search?.navigator?.stack?.isNotEmpty() == true
    CompositionLocalProvider(LocalItemPageOpener provides openItemPage) {
        Box(Modifier.fillMaxSize()) {
            // While Search covers the app, the tabs under it are hidden from TalkBack.
            Box(if (searchOpen) Modifier.fillMaxSize().clearAndSetSemantics {} else Modifier.fillMaxSize()) {
                ForeverTabScaffold(selected = selected, onSelect = { selected = it }, tabs = ForeverTab.visible(guides != null)) {
                    // Lane F13: a shared build and Your builds are pushed over the current tab.
                    if (planner != null && plannerTop != null) {
                        PlannerScreenHost(planner, plannerTop)
                        return@ForeverTabScaffold
                    }
                    // The tab slots. Each lane replaces its placeholder branch: F2 Home, F4 Raids, F3 BiS, F1 You.
                    shell.SaveableStateProvider(SHELL_KEY) {
                        when (selected) {
                            ForeverTab.Home ->
                                if (account == null) {
                                    HomeRoute(home, openBis = openBis, onSearch = onSearch) { ComingSoon(selected) }
                                } else {
                                    HomeAccountGate(account, openBis) {
                                        HomeRoute(home, account.session, openBis, onSearch, signIn = { account.session.beginSignIn(account.browser(context)) }) { ComingSoon(selected) }
                                    }
                                }
                            ForeverTab.Raids -> if (raids == null) ComingSoon(selected) else RaidsTab(raids, account, home.reports, onSearch)
                            ForeverTab.Bis ->
                                if (bis == null) {
                                    ComingSoon(selected)
                                } else {
                                    BisTab(bis, home, account, onSearch)
                                }
                            ForeverTab.Guides -> if (guides == null) ComingSoon(selected) else GuidesTab(guides, account, bis, planner) { selected = ForeverTab.Bis }
                            ForeverTab.You -> if (account == null) ComingSoon(selected) else YouRoute(account, openBis)
                        }
                    }
                }
            }
            if (search != null) SearchOverlay(search, searchOwner(account), exits)
        }
    }
}

/** The tab content's saved-state key: one, so a tab's screens come back after a planner screen closes. */
private const val SHELL_KEY = "tabs"

/**
 * The BiS tab under the character selector (spec §3.4): the list follows the session's current
 * character, and waits while a signed-in session's characters are still on their way.
 */
@Composable
private fun BisTab(bis: BisModel, home: HomeDependencies, account: AccountDeps?, onSearch: (() -> Unit)?) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val characterKey = phase?.let { session.currentCharacter(it, picked)?.key }
    val waiting = phase == SessionPhase.Importing || (phase is SessionPhase.SignedIn && phase.me == null)
    BisRoute(bis, home.catalog, BisOwner(characterKey, waiting), onSearch) { if (session != null) CharacterSelectorBar(session) }
}

/**
 * Guides (lane F9) for the session's current character, under the same selector as Home and BiS.
 * A guide's Gear row asks BiS for that spec's level-60 raid list and switches to it.
 */
@Composable
private fun GuidesTab(guides: GuidesFeature, account: AccountDeps?, bis: BisModel?, planner: PlannerFeature?, showBis: () -> Unit) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val character = phase?.let { session.currentCharacter(it, picked) }
    GuidesRoute(
        guides,
        character,
        openBis = { guide ->
            if (bis != null) {
                bis.request(BisSelection.forGuide(guide.classSlug, guide.gear?.spec ?: guide.spec, character?.faction))
                showBis()
            }
        },
        planner = planner,
    ) { if (session != null) CharacterSelectorBar(session) }
}

/**
 * Raids (lane F4) for the session's current character's guild, under the same selector as Home
 * and BiS; signed out or with sign-in off, it follows `/v1/me`'s main character.
 */
@Composable
private fun RaidsTab(raids: RaidsFeature, account: AccountDeps?, reports: ReportsFeature?, onSearch: (() -> Unit)?) {
    val session = account?.session
    val phase = session?.phase?.collectAsStateWithLifecycle()?.value
    val picked = session?.pickedKey?.collectAsStateWithLifecycle()?.value
    val characterKey = phase?.let { session.currentCharacter(it, picked)?.key }
    // Lane F10: Your raid night's `Open the full report` pushes the report on this stack.
    val me = (phase as? SessionPhase.SignedIn)?.me
    val reportsLink = reports?.let { RaidsReports(it, me?.let { m -> ReportsViewer.of(m, characterKey) } ?: ReportsViewer.SignedOut) }
    RaidsRoute(raids, characterKey, reportsLink) { RootSearchBar(onSearch) { if (session != null) CharacterSelectorBar(session) } }
}

/** Search's owner: the session's current character and the account's character keys (`You`). */
@Composable
private fun searchOwner(account: AccountDeps?): SearchOwner {
    val session = account?.session ?: return SearchOwner()
    val phase = session.phase.collectAsStateWithLifecycle().value
    val picked = session.pickedKey.collectAsStateWithLifecycle().value
    return SearchOwner(session.currentCharacter(phase, picked), phase.me?.characters?.map { it.key }?.toSet().orEmpty())
}

/** The scaffold with a navigation bar, stateless so screenshot tests can drive it. */
@Composable
fun ForeverTabScaffold(
    selected: ForeverTab,
    onSelect: (ForeverTab) -> Unit,
    tabs: List<ForeverTab> = ForeverTab.visible(guides = false),
    content: @Composable () -> Unit,
) {
    Scaffold(
        containerColor = MaterialTheme.colorScheme.background,
        bottomBar = {
            NavigationBar(containerColor = MaterialTheme.colorScheme.surface) {
                tabs.forEach { tab ->
                    NavigationBarItem(
                        selected = tab == selected,
                        onClick = { onSelect(tab) },
                        icon = { Icon(tab.icon.painter(), contentDescription = null) },
                        label = { TabLabel(stringResource(tab.label)) },
                    )
                }
            }
        },
    ) { padding ->
        Box(Modifier.fillMaxSize().padding(padding)) { content() }
    }
}

/** The largest font scale a tab label grows to; past it the label stays one line (spec §10). */
private const val TAB_LABEL_MAX_FONT_SCALE = 1.3f

/** A tab label on one line whatever the font scale: the scale is capped, the icon carries the rest. */
@Composable
private fun TabLabel(text: String) {
    val density = LocalDensity.current
    CompositionLocalProvider(
        LocalDensity provides Density(density.density, density.fontScale.coerceAtMost(TAB_LABEL_MAX_FONT_SCALE)),
    ) {
        Text(text, maxLines = 1, softWrap = false)
    }
}

@Composable
private fun ComingSoon(tab: ForeverTab) {
    Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
        Text(
            text = stringResource(R.string.shell_arrivesNext, stringResource(tab.label)),
            style = MaterialTheme.typography.bodyLarge,
            color = MaterialTheme.colorScheme.onSurfaceVariant,
        )
    }
}
