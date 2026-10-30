package gg.foreversixty.app

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
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.core.net.toUri
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverAppScaffold
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
                    ForeverAppScaffold(accountApi = graph.accountApi, bisModel = graph.bisModel, itemCatalog = graph.itemCatalog)
                } else {
                    Box(Modifier.fillMaxSize().background(MaterialTheme.colorScheme.background)) {
                        UpdateRequiredScreen(onOpenStore = { openStore(answer.storeUrl ?: PLAY_STORE_URL) })
                    }
                }
            }
        }
        handleAppLink(intent)
        lifecycleScope.launch { graph.appConfig.refresh() }
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
     * apps/shared/links.json) names the screen. Only the sign-in return is wired so far; the
     * feature lanes open their screens from the other cases.
     */
    private fun handleAppLink(intent: Intent?) {
        val link = intent?.dataString ?: return
        if (DeepLink.parse(link) is DeepLink.SignIn) {
            // TODO(lane F1): surface the SignInOutcome (signed in, cancelled, rejected, failed).
            lifecycleScope.launch { graph.signIn.complete(link) }
        }
    }
}
