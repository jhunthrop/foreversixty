package gg.foreversixty.app.auth

import gg.foreversixty.app.model.AppExchangeRequest
import gg.foreversixty.app.model.AppExchangeResponse
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiResult
import java.security.SecureRandom

/** Opens a URL in the system browser surface (Custom Tabs in the app, a fake in tests). */
fun interface BrowserLauncher {
    fun open(url: String)
}

/** The outcome of finishing sign-in. */
sealed interface SignInOutcome {
    data class SignedIn(val session: DeviceSession) : SignInOutcome
    /** The player backed out or Battle.net refused. */
    data class Cancelled(val reason: String) : SignInOutcome
    /** The redirect was not for the sign-in in progress, or there was none in progress. */
    data object Rejected : SignInOutcome
    /** The exchange call failed; [result] carries the API or network error. */
    data class ExchangeFailed(val result: ApiResult.Failure) : SignInOutcome
}

/**
 * Battle.net sign-in for the app: PKCE in a Custom Tab, then the code exchange for a device
 * token, which lands in [tokens].
 *
 * The verifier and state live in memory only. If the process dies while the browser is open,
 * the redirect is [SignInOutcome.Rejected] and the player taps Sign in again; F1 may persist
 * the pending pair if that proves common.
 */
class AppSignIn(
    private val baseUrl: String,
    private val client: ApiClient,
    private val tokens: TokenStore,
    private val deviceName: String,
    private val random: SecureRandom = SecureRandom(),
) {
    private data class Pending(val verifier: String, val state: String)

    @Volatile
    private var pending: Pending? = null

    /** Starts sign-in: remembers a fresh verifier and state and opens the start URL. */
    fun begin(browser: BrowserLauncher) {
        val next = Pending(Pkce.newVerifier(random), Pkce.newState(random))
        pending = next
        browser.open(SignInContract.startUrl(baseUrl, Pkce.challenge(next.verifier), next.state))
    }

    /** Finishes sign-in from the App Link redirect. Each begin() is good for one finish. */
    suspend fun complete(redirect: String): SignInOutcome {
        val started = pending ?: return SignInOutcome.Rejected
        return when (val parsed = SignInContract.parseRedirect(redirect, started.state)) {
            is RedirectResult.Code -> {
                pending = null
                exchange(parsed.code, started.verifier)
            }
            is RedirectResult.Denied -> {
                pending = null
                SignInOutcome.Cancelled(parsed.reason)
            }
            RedirectResult.StateMismatch, RedirectResult.Malformed -> SignInOutcome.Rejected
        }
    }

    suspend fun signOut() {
        // TODO(lane F1): call DELETE /v1/devices/self before clearing, so /account drops the row.
        tokens.clear()
    }

    private suspend fun exchange(code: String, verifier: String): SignInOutcome {
        val request = AppExchangeRequest(code, verifier, SignInContract.PLATFORM, deviceName)
        val result = client.post(
            SignInContract.EXCHANGE_PATH,
            request,
            AppExchangeRequest.serializer(),
            AppExchangeResponse.serializer(),
        )
        return when (result) {
            is ApiResult.Success -> {
                val session = DeviceSession(result.data.deviceId, result.data.token, result.data.accountLabel)
                tokens.save(session)
                SignInOutcome.SignedIn(session)
            }
            is ApiResult.Failure -> SignInOutcome.ExchangeFailed(result)
        }
    }
}
