package gg.foreversixty.app.auth

import java.net.URI
import java.net.URLDecoder
import java.net.URLEncoder

/**
 * The app sign-in contract (plan 2026-10-30, lane A1), in one place.
 *
 * TODO(lane A1): confirm the query parameter names and the redirect path against
 * docs/contracts/app-v1.md once lane C and A1 land; only this file changes.
 */
object SignInContract {
    const val START_PATH = "/v1/auth/app/start"
    const val EXCHANGE_PATH = "/v1/auth/app/exchange"
    const val PLATFORM = "android"

    /** The App Link the API redirects to with `?code=…&state=…` (or `?error=…`). */
    const val REDIRECT_URI = "https://foreversixty.gg/app/auth/callback"

    /** The URL opened in a Custom Tab to begin Battle.net sign-in. */
    fun startUrl(baseUrl: String, challenge: String, state: String): String {
        val query = listOf(
            "code_challenge" to challenge,
            "code_challenge_method" to Pkce.METHOD,
            "state" to state,
            "platform" to PLATFORM,
            "redirect_uri" to REDIRECT_URI,
        ).joinToString("&") { (k, v) -> "$k=${encode(v)}" }
        return "${baseUrl.trimEnd('/')}$START_PATH?$query"
    }

    /** Reads the redirect the App Link delivered. Never trusts a state it did not issue. */
    fun parseRedirect(redirect: String, expectedState: String): RedirectResult {
        val uri = runCatching { URI(redirect) }.getOrNull() ?: return RedirectResult.Malformed
        if (!isRedirect(uri)) return RedirectResult.Malformed
        val params = query(uri.rawQuery)
        params["error"]?.let { return RedirectResult.Denied(it) }
        if (params["state"] != expectedState) return RedirectResult.StateMismatch
        val code = params["code"]?.takeIf { it.isNotBlank() } ?: return RedirectResult.Malformed
        return RedirectResult.Code(code)
    }

    private val redirect: URI = URI(REDIRECT_URI)

    private fun isRedirect(uri: URI): Boolean =
        uri.scheme == redirect.scheme && uri.host == redirect.host && uri.path == redirect.path

    private fun query(raw: String?): Map<String, String> = raw.orEmpty()
        .split('&')
        .filter { it.isNotEmpty() }
        .associate { pair ->
            val key = pair.substringBefore('=')
            decode(key) to decode(pair.substringAfter('=', ""))
        }

    // The String-charset overloads: the Charset ones need API 33 and minSdk is 28.
    private const val UTF8 = "UTF-8"
    private fun encode(value: String): String = URLEncoder.encode(value, UTF8)
    private fun decode(value: String): String = URLDecoder.decode(value, UTF8)
}

/** What the sign-in redirect said. */
sealed interface RedirectResult {
    data class Code(val code: String) : RedirectResult
    /** The player cancelled, or Battle.net refused; [reason] is the API's error word. */
    data class Denied(val reason: String) : RedirectResult
    /** The state does not match the sign-in this app began: drop it (possible CSRF). */
    data object StateMismatch : RedirectResult
    data object Malformed : RedirectResult
}
