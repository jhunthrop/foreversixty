package gg.foreversixty.app

import android.content.Intent
import android.graphics.Color
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.SystemBarStyle
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.navigation.ForeverAppScaffold
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
                ForeverAppScaffold(accountApi = graph.accountApi)
            }
        }
        handleAppLink(intent)
    }

    override fun onNewIntent(intent: Intent) {
        super.onNewIntent(intent)
        handleAppLink(intent)
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
