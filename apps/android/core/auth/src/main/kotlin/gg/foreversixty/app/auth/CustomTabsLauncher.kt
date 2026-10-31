package gg.foreversixty.app.auth

import android.content.Context
import android.net.Uri
import androidx.browser.customtabs.CustomTabsIntent

/** Opens sign-in in a Custom Tab: the system browser, its cookies and password manager. */
class CustomTabsLauncher(private val context: Context) : BrowserLauncher {
    override fun open(url: String) {
        CustomTabsIntent.Builder()
            .setShowTitle(true)
            .build()
            .launchUrl(context, Uri.parse(url))
    }
}
