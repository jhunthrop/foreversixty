package gg.foreversixty.app

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import gg.foreversixty.app.common.LocalLatestBuild
import gg.foreversixty.app.system.logic.AppShortcut
import androidx.compose.runtime.getValue
import android.content.Intent
import android.graphics.Color
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.SystemBarStyle
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.ui.Modifier
import androidx.core.net.toUri
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverAppShell
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.menu.MenuDestination
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.update.PLAY_STORE_URL
import gg.foreversixty.app.update.UpdateRequiredScreen
import gg.foreversixty.app.auth.session.isSignedIn
import gg.foreversixty.app.push.FcmToken
import gg.foreversixty.app.push.PushGroup
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.distinctUntilChanged
import kotlinx.coroutines.flow.filter
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.launch
import kotlin.time.Duration.Companion.seconds

class MainActivity : ComponentActivity() {
    private val app: ForeverApplication get() = application as ForeverApplication
    private val bootstrap: AppBootstrap get() = app.bootstrap
    private val graph: AppGraph get() = app.graph

    /** True once [startShell] ran: the graph exists and the shell draws from it. */
    private val shellStarted = MutableStateFlow(false)

    /** An intent that arrived before the shell started; handled when it does. */
    private var pendingIntent: Intent? = null

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        // The app is dark only, so the bar icons are light whatever the system theme.
        val darkBars = SystemBarStyle.dark(Color.TRANSPARENT)
        enableEdgeToEdge(statusBarStyle = darkBars, navigationBarStyle = darkBars)
        setContent {
            ForeverTheme {
                val gate by bootstrap.updateGate.state.collectAsStateWithLifecycle()
                val started by shellStarted.collectAsStateWithLifecycle()
                val answer = gate
                when {
                    answer != null -> Box(Modifier.fillMaxSize().background(MaterialTheme.colorScheme.background)) {
                        UpdateRequiredScreen(onOpenStore = { openStore(answer.storeUrl ?: PLAY_STORE_URL) })
                    }
                    // A fresh install holds for its first config: the launch window's background, a moment longer.
                    !started -> Box(Modifier.fillMaxSize().background(MaterialTheme.colorScheme.background))
                    else -> Shell()
                }
            }
        }
        pendingIntent = intent
        lifecycleScope.launch { bootstrap.appConfig.refresh() }
        lifecycleScope.launch {
            bootstrap.appConfig.awaitFirstAnswer(FIRST_CONFIG_HOLD)
            startShell()
        }
    }

    /** The tabs over the graph, read again on every config change so a switch the API turns on shows at once. */
    @Composable
    private fun Shell() {
        // Lane F15: the build-data screens' new-data strip reads the fresh active build.
        CompositionLocalProvider(LocalLatestBuild provides graph::latestBuild) {
            ForeverAppShell(graph.appConfig.config) { graph.shellFeatures(::openLink) }
        }
    }

    /** Opens an app link the way an App Link intent does (a notification, an Activity row). */
    private fun openLink(link: String) = handleAppLink(Intent(Intent.ACTION_VIEW, link.toUri()))

    /**
     * Builds the graph from the config as it now stands and starts the account's background work.
     * Runs once the first config is known, or the hold has passed, so no feature is decided from
     * the bundled defaults while the real switches are a moment away.
     */
    private fun startShell() {
        shellStarted.value = true
        handleAppLink(pendingIntent)
        pendingIntent = null
        lifecycleScope.launch {
            graph.session.start()
            graph.session.refresh()
        }
        lifecycleScope.launch { graph.session.watchForDeadTokens() }
        // Lane F15: the dynamic shortcuts follow the session too (its characters arrive after launch).
        lifecycleScope.launch { graph.session.phase.collect { graph.system.refreshShortcuts() } }
        // Push: the groups' channels exist from the first launch; the token registers on each sign-in.
        PushGroup.createChannels(this)
        lifecycleScope.launch {
            graph.session.phase.map { it.isSignedIn }.distinctUntilChanged().filter { it }.collect {
                FcmToken.request(this@MainActivity)
                graph.push.register()
                graph.activity?.refreshUnread()
            }
        }
        // Lane F14: the 30-minute snapshot refresh behind the widgets, tile and countdown.
        if (graph.widgets != null) gg.foreversixty.app.widgets.WidgetRefreshWork.periodic(this)
        // The activity is already in front: the foreground refreshes run now.
        refreshOnForeground()
    }

    override fun onResume() {
        super.onResume()
        // Before the shell starts there is nothing to come back to; startShell runs these itself.
        if (!shellStarted.value) return
        // A Custom Tab closed without a redirect: back to signed out (spec F1.2's cancel).
        graph.session.onReturnedToApp()
        refreshOnForeground()
    }

    /** What every foreground refreshes: guides, the widgets' snapshot, the Activity badge, the shortcuts. */
    private fun refreshOnForeground() {
        // Lane F9: the active build and the player's classes' guides, at most hourly.
        graph.guides?.let { guides -> lifecycleScope.launch { guides.sync() } }
        // Lane F14: every foreground rewrites the snapshot the widgets and the tile read.
        graph.widgets?.let { widgets -> lifecycleScope.launch { widgets.refresher.refresh() } }
        // The Activity badge: whatever arrived while the app was away.
        graph.activity?.let { activity -> lifecycleScope.launch { activity.refreshUnread() } }
        // Lane F15: the dynamic shortcuts follow the data (and go when signed out or switched off).
        lifecycleScope.launch { graph.system.refreshShortcuts() }
    }

    override fun onNewIntent(intent: Intent) {
        super.onNewIntent(intent)
        if (shellStarted.value) handleAppLink(intent) else pendingIntent = intent
    }

    /** A static shortcut's extra, or the Assistant's `feature`; true when it named a screen. */
    private fun openShortcut(intent: Intent?): Boolean {
        intent ?: return false
        val shortcut = AppShortcut.of(intent.getStringExtra(AppShortcut.EXTRA))
            ?: AppShortcut.ofFeature(intent.getStringExtra(ASSISTANT_FEATURE)).takeIf { graph.system.enabled }
            ?: return false
        if (shortcut == AppShortcut.Search) graph.search.takeIf { it.enabled }?.navigator?.openSearch()
        graph.system.navigator.open(shortcut)
        return true
    }

    /**
     * The web page, in the browser app: the selector keeps this app's own /b/ filter from catching the
     * link again (a deep link is never a dead end, spec §3.4).
     */
    private fun openInBrowser(link: String) {
        val view = Intent(Intent.ACTION_VIEW, link.toUri()).addCategory(Intent.CATEGORY_BROWSABLE)
        view.selector = Intent.makeMainSelectorActivity(Intent.ACTION_MAIN, Intent.CATEGORY_APP_BROWSER)
        runCatching { startActivity(view) }
    }

    private fun openStore(url: String) {
        startActivity(Intent(Intent.ACTION_VIEW, url.toUri()))
    }

    /**
     * App Links arrive here (see the TODO in AndroidManifest.xml); [DeepLink] (generated from
     * apps/shared/links.json) names the screen. The sign-in return goes to the session, which
     * shows F1.3 and then signed in or F1.4; the feature lanes open their screens from the others.
     */
    private fun handleAppLink(intent: Intent?) {
        // Lane F8 review captures, debug builds only: the storefront country to show (`GB`, `none`).
        if (BuildConfig.DEBUG) intent?.getStringExtra(gg.foreversixty.app.premium.DebugStorefront.EXTRA)?.let { gg.foreversixty.app.premium.DebugStorefront.code = it }
        // Review captures, debug builds only (round 4 a2): fixture states (`guild.raid-badge=officer`).
        if (BuildConfig.DEBUG) intent?.getStringExtra(DebugFixtures.EXTRA)?.let { DebugFixtures.variants = DebugFixtures.parse(it) }
        if (BuildConfig.DEBUG && intent?.hasExtra(DebugFixtures.LATENCY_EXTRA) == true) {
            DebugFixtures.latencyMillis = intent.getLongExtra(DebugFixtures.LATENCY_EXTRA, 0L).coerceAtLeast(0L)
        }
        // Review captures, debug builds only (round 4 a2): post a push fixture by name as FCM draws it.
        if (BuildConfig.DEBUG) intent?.getStringExtra(gg.foreversixty.app.push.DebugPush.EXTRA)?.let { gg.foreversixty.app.push.DebugPush.post(this, it) }
        // Lane F15 review captures, debug builds only: the real share chooser with the app's share.
        if (BuildConfig.DEBUG) intent?.getStringExtra(gg.foreversixty.app.system.platform.DebugShare.EXTRA)?.let { gg.foreversixty.app.system.platform.DebugShare.open(this, it) }
        // Lane F15: a launcher shortcut or the Assistant's Open app feature.
        if (openShortcut(intent)) return
        val link = intent?.dataString ?: return
        when (val parsed = DeepLink.parse(link)) {
            is DeepLink.SignIn -> graph.session.handleRedirect(link)
            // A captured site link the app has no screen for (or a site-only page, DeepLink.SITE_ONLY) is the web page.
            null -> if (link.startsWith("https://")) openInBrowser(link)
            // Lane F5: a sim link opens Home, Sims and the result.
            is DeepLink.Sim -> graph.home.sims?.navigator?.open(parsed)
            // Lane F10: a report link opens Home, Your reports and the report.
            is DeepLink.Report -> graph.home.reports?.navigator?.open(parsed)
            // Lane F11: a rankings link opens Home and Rankings; a player link opens Home and that character.
            is DeepLink.Rankings, is DeepLink.Player -> graph.home.rankings?.navigator?.open(parsed)
            // Lane F9: a guides link opens the Guides tab and the guide.
            is DeepLink.Guides, is DeepLink.Guide -> graph.guides?.navigator?.open(parsed)
            // Lane F13: the planner, Your builds and a shared build; with `planner` off a /b/ link is the web page.
            is DeepLink.Planner, DeepLink.Builds -> graph.planner?.open(parsed)
            is DeepLink.SharedBuild -> if (graph.planner?.open(parsed) != true) openInBrowser(link)
            // The site's own pages (spec 2026-11-08 F6), each pushed under the tab it belongs to.
            is DeepLink.SimPage -> graph.home.sims?.navigator?.open(DeepLink.Sim(parsed.id))
            is DeepLink.ReportPage -> graph.home.reports?.navigator?.open(DeepLink.Report(parsed.id))
            is DeepLink.CharacterPage -> graph.home.rankings?.navigator?.open(DeepLink.Player(parsed.region, parsed.ruleset, parsed.name))
            is DeepLink.GuidePage -> graph.guides?.navigator?.open(DeepLink.Guide(parsed.className, parsed.spec))
            is DeepLink.GuidesClassPage -> graph.guides?.navigator?.openClass(parsed.className)
            is DeepLink.BisPage -> {
                graph.bisModel?.request(BisSelection.forGuide(parsed.className, parsed.spec, null))
                graph.system.navigator.wantsTab = ForeverTab.Gear
            }
            DeepLink.TiersPage -> graph.system.navigator.wantsMenu = MenuDestination.Tiers
            // The site names the guild, not its id: the Raid tab answers for the player's own guild.
            is DeepLink.GuildRaidPage -> graph.system.navigator.wantsTab = ForeverTab.Guild
            // Lane F4: raid, raid-night and guild links (and the raid notifications) open the Raids tab.
            else -> if (graph.raids.enabled) graph.raids.navigator.open(parsed)
        }
    }
}

/** The Assistant's `OPEN_APP_FEATURE` parameter (shortcuts.xml). */
private const val ASSISTANT_FEATURE = "feature"

/**
 * How long a fresh install waits for its first config before the shell builds from the bundled
 * defaults (every switch off). The fetch is one small GET right after launch; past this the
 * tabs show, and the config still lands into the shell when it arrives.
 */
private val FIRST_CONFIG_HOLD = 4.seconds
