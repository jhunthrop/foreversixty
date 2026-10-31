package gg.foreversixty.app

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
                    ForeverAppScaffold(
                        home = graph.home,
                        bis = graph.bisModel,
                        account = graph.account.takeIf { graph.isAvailable(Feature.SignIn) },
                        raids = graph.raids.takeIf { it.enabled },
                    )
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
    }

    override fun onResume() {
        super.onResume()
        // A Custom Tab closed without a redirect: back to signed out (spec F1.2's cancel).
        graph.session.onReturnedToApp()
    }

    override fun onNewIntent(intent: Intent) {
        super.onNewIntent(intent)
        handleAppLink(intent)
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
        // Lane F4 review captures, debug builds only: post the fixture raid notification.
        if (BuildConfig.DEBUG && intent?.getBooleanExtra(gg.foreversixty.app.raids.push.RaidDebugPush.EXTRA, false) == true) {
            gg.foreversixty.app.raids.push.RaidDebugPush.post(this)
        }
        val link = intent?.dataString ?: return
        when (val parsed = DeepLink.parse(link)) {
            is DeepLink.SignIn -> graph.session.handleRedirect(link)
            null -> Unit
            // Lane F5: a sim link opens Home, Sims and the result.
            is DeepLink.Sim -> graph.home.sims?.navigator?.open(parsed)
            // Lane F4: raid, raid-night and guild links (and the raid notifications) open the Raids tab.
            else -> if (graph.raids.enabled) graph.raids.navigator.open(parsed)
        }
    }
}
