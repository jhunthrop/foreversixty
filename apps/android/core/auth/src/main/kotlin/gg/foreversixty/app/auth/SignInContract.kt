package gg.foreversixty.app.auth

import java.net.URI
import java.net.URLDecoder
import java.net.URLEncoder

/**
 * The app sign-in contract (lane A1; docs/contracts/app-v1.md "Sign in"), in one place.
 *
 * 1. A Custom Tab opens `GET /v1/auth/app/start?platform=android&code_challenge=…&
 *    code_challenge_method=S256&state=…`. The API answers a confirm page; the player taps
 *    "Continue to Battle.net" and signs in there.
 * 2. The browser lands on exactly [REDIRECT_URI] `?code=…&state=…` (no `/callback` suffix). The
 *    server picks that landing; a client `redirect_uri` is ignored, so the app sends none.
 * 3. `POST /v1/auth/app/exchange {code, code_verifier, device_name}` answers 201
 *    `{token, device, me}`.
 *
 * Afterwards a 401 (`unauthorized` or `session_expired`) means the token is dead: the client
 * discards it and the app signs in again. Billing writes answer 403 `browser_session_required`.
 */
object SignInContract {
    const val START_PATH = "/v1/auth/app/start"
    const val PLATFORM = "android"

    /** The App Link the API hands the code to with `?code=…&state=…` (or `?error=…`). */
    const val REDIRECT_URI = "https://foreversixty.gg/app/auth"

    /** The URL opened in a Custom Tab to begin Battle.net sign-in. */
    fun startUrl(baseUrl: String, challenge: String, state: String): String {
        val query = listOf(
            "code_challenge" to challenge,
            "code_challenge_method" to Pkce.METHOD,
            "state" to state,
            "platform" to PLATFORM,
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
