package gg.foreversixty.app

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
import gg.foreversixty.app.navigation.ForeverAppScaffold
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.update.PLAY_STORE_URL
import gg.foreversixty.app.update.UpdateRequiredScreen
import kotlinx.coroutines.launch

class MainActivity : ComponentActivity() {
    private val graph: AppGraph get() = (application as ForeverApplication).graph

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        // The app is dark only, so the bar icons are light whatever the system theme.
        val darkBars = SystemBarStyle.dark(Color.TRANSPARENT)
        enableEdgeToEdge(statusBarStyle = darkBars, navigationBarStyle = darkBars)
        setContent {
            ForeverTheme {
                val gate by graph.updateGate.state.collectAsStateWithLifecycle()
                val answer = gate
                if (answer == null) {
                    // Lane F15: the build-data screens' new-data strip reads the fresh active build.
                    CompositionLocalProvider(LocalLatestBuild provides graph::latestBuild) {
                    ForeverAppScaffold(
                        home = graph.home,
                        bis = graph.bisModel,
                        account = graph.account.takeIf { graph.isAvailable(Feature.SignIn) },
                        raids = graph.raids.takeIf { it.enabled },
                        search = graph.search.takeIf { it.enabled },
                        guides = graph.guides,
                        system = graph.system.navigator,
                    )
                    }
                } else {
                    Box(Modifier.fillMaxSize().background(MaterialTheme.colorScheme.background)) {
                        UpdateRequiredScreen(onOpenStore = { openStore(answer.storeUrl ?: PLAY_STORE_URL) })
                    }
                }
            }
        }
        handleAppLink(intent)
        lifecycleScope.launch { graph.appConfig.refresh() }
        lifecycleScope.launch {
            graph.session.start()
            graph.session.refresh()
        }
        lifecycleScope.launch { graph.session.watchForDeadTokens() }
        // Lane F14: the 30-minute snapshot refresh behind the widgets, tile and countdown.
        if (graph.widgets != null) gg.foreversixty.app.widgets.WidgetRefreshWork.periodic(this)
    }

    override fun onResume() {
        super.onResume()
        // A Custom Tab closed without a redirect: back to signed out (spec F1.2's cancel).
        graph.session.onReturnedToApp()
        // Lane F9: the active build and the player's classes' guides, at most hourly.
        graph.guides?.let { guides -> lifecycleScope.launch { guides.sync() } }
        // Lane F14: every foreground rewrites the snapshot the widgets and the tile read.
        graph.widgets?.let { widgets -> lifecycleScope.launch { widgets.refresher.refresh() } }
        // Lane F15: the dynamic shortcuts follow the data (and go when signed out or switched off).
        lifecycleScope.launch { graph.system.refreshShortcuts() }
    }

    override fun onNewIntent(intent: Intent) {
        super.onNewIntent(intent)
        handleAppLink(intent)
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
        // Lane F4 review captures, debug builds only: post the fixture raid notification.
        if (BuildConfig.DEBUG && intent?.getBooleanExtra(gg.foreversixty.app.raids.push.RaidDebugPush.EXTRA, false) == true) {
            gg.foreversixty.app.raids.push.RaidDebugPush.post(this)
        }
        // Lane F15: a launcher shortcut or the Assistant's Open app feature.
        if (openShortcut(intent)) return
        val link = intent?.dataString ?: return
        when (val parsed = DeepLink.parse(link)) {
            is DeepLink.SignIn -> graph.session.handleRedirect(link)
            null -> Unit
            // Lane F5: a sim link opens Home, Sims and the result.
            is DeepLink.Sim -> graph.home.sims?.navigator?.open(parsed)
            // Lane F10: a report link opens Home, Your reports and the report.
            is DeepLink.Report -> graph.home.reports?.navigator?.open(parsed)
            // Lane F11: a rankings link opens Home and Rankings; a player link opens Home and that character.
            is DeepLink.Rankings, is DeepLink.Player -> graph.home.rankings?.navigator?.open(parsed)
            // Lane F9: a guides link opens the Guides tab and the guide.
            is DeepLink.Guides, is DeepLink.Guide -> graph.guides?.navigator?.open(parsed)
            // Lane F4: raid, raid-night and guild links (and the raid notifications) open the Raids tab.
            else -> if (graph.raids.enabled) graph.raids.navigator.open(parsed)
        }
    }
}

/** The Assistant's `OPEN_APP_FEATURE` parameter (shortcuts.xml). */
private const val ASSISTANT_FEATURE = "feature"
