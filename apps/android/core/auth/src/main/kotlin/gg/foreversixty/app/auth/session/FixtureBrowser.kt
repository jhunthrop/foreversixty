package gg.foreversixty.app.auth.session

import gg.foreversixty.app.auth.BrowserLauncher
import gg.foreversixty.app.auth.SignInContract
import java.net.URI

/**
 * Fixture mode's browser: answers at once with the redirect the API would send, so the whole
 * sign-in (state check, exchange, import) runs against the contract fixtures, offline.
 */
class FixtureBrowser(private val onRedirect: (String) -> Unit) : BrowserLauncher {
    override fun open(url: String) {
        val state = URI(url).rawQuery.orEmpty().split('&')
            .firstOrNull { it.startsWith("state=") }
            ?.substringAfter('=')
            .orEmpty()
        onRedirect("${SignInContract.REDIRECT_URI}?code=$CODE&state=$state")
    }

    companion object {
        const val CODE = "fixture-code"
    }
}
