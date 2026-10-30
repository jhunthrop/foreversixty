package gg.foreversixty.app

import android.content.Intent
import android.graphics.Color
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.SystemBarStyle
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.auth.SignInContract
import gg.foreversixty.app.designsystem.ForeverTheme
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

    /** The sign-in redirect arrives as an App Link (see the TODO in AndroidManifest.xml). */
    private fun handleAppLink(intent: Intent?) {
        val link = intent?.dataString ?: return
        if (link.startsWith(SignInContract.REDIRECT_URI)) {
            // TODO(lane F1): surface the SignInOutcome (signed in, cancelled, rejected, failed).
            lifecycleScope.launch { graph.signIn.complete(link) }
        }
    }
}
