package gg.foreversixty.app.auth.session

import gg.foreversixty.app.auth.AppSignIn
import gg.foreversixty.app.auth.BrowserLauncher
import gg.foreversixty.app.auth.SignInOutcome
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.AccountDeletionRequest
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.launch
import java.time.Instant

/**
 * What an account action came to. [NeedsBrowser] is `403 browser_session_required`: the act is
 * done on foreversixty.gg, so the view opens the site and never retries.
 */
enum class AccountActionOutcome {
    Done,
    Failed,
    NeedsBrowser,
    ;

    internal companion object {
        fun of(error: ApiError): AccountActionOutcome =
            if ((error as? ApiError.Server)?.needsBrowserSession == true) NeedsBrowser else Failed
    }
}

/** The body `DELETE /v1/me` takes after the person confirmed (api/openapi.yaml). */
val CONFIRMED_DELETION = AccountDeletionRequest(confirm = "delete my account")

/**
 * The account for the whole app (iOS `SessionStore`): the phase, the current character, and the
 * account acts of spec §5.1. Any 401 on any request ends the session through
 * [SessionTokens.discarded]; [scope] runs the redirect's exchange.
 */
class SessionManager(
    private val client: ApiClient,
    private val tokens: SessionTokens,
    private val signInFlow: AppSignIn,
    private val picked: StoredValue,
    private val scope: CoroutineScope,
    private val clearCaches: suspend () -> Unit = {},
) {
    private val mutablePhase = MutableStateFlow<SessionPhase>(SessionPhase.SignedOut())
    private val mutablePicked = MutableStateFlow(picked.load())
    private var endingOnPurpose = false

    val phase: StateFlow<SessionPhase> = mutablePhase.asStateFlow()
    val pickedKey: StateFlow<String?> = mutablePicked.asStateFlow()

    /** Reads the stored session: signed in (me to come) with a token, else signed out. */
    suspend fun start() {
        mutablePhase.value = if (tokens.bearerToken() != null) SessionPhase.SignedIn(null) else SessionPhase.SignedOut()
    }

    /** Ends the session whenever a request is answered 401. Collect once from the app root. */
    suspend fun watchForDeadTokens() {
        tokens.discarded.collect { tokenDied() }
    }

    /** This install's device id, from the sign-in exchange. */
    suspend fun thisDeviceId(): String? = tokens.session.first()?.deviceId

    /** The character Home, Raids and BiS follow (spec §3.4). */
    fun currentCharacter(phase: SessionPhase = this.phase.value, pickedKey: String? = mutablePicked.value): AccountCharacter? =
        phase.me?.let { CharacterSelection.current(it, pickedKey) }

    /** Reads `GET /v1/me`. A 401 ends the session; any other failure keeps what is on screen. */
    suspend fun refresh() {
        if (!phase.value.isSignedIn) return
        when (val result = client.send(ApiEndpoints.getMe())) {
            is ApiResult.Success -> if (phase.value.isSignedIn) mutablePhase.value = SessionPhase.SignedIn(result.data)
            is ApiResult.Failure -> if ((result.error as? ApiError.Server)?.isUnauthorized == true) tokenDied()
        }
    }

    /** F1.2: opens the Custom Tab. The App Link redirect comes back through [handleRedirect]. */
    fun beginSignIn(browser: BrowserLauncher) {
        mutablePhase.value = SessionPhase.SigningIn
        signInFlow.begin(browser)
    }

    /** F1.3: the redirect arrived. Importing at once, then the exchange and `GET /v1/me`. */
    fun handleRedirect(redirect: String) {
        mutablePhase.value = SessionPhase.Importing
        scope.launch { finishSignIn(redirect) }
    }

    /**
     * The app came back to the front. Still signing in means the player closed the Custom Tab
     * without a redirect: back to F1.1 with no line (spec F1.2's cancel).
     */
    fun onReturnedToApp() {
        if (phase.value == SessionPhase.SigningIn) mutablePhase.value = SessionPhase.SignedOut()
    }

    internal suspend fun finishSignIn(redirect: String) {
        when (val outcome = signInFlow.complete(redirect)) {
            is SignInOutcome.SignedIn -> {
                val me = client.send(ApiEndpoints.getMe())
                mutablePhase.value = SessionPhase.SignedIn((me as? ApiResult.Success)?.data)
            }
            else -> mutablePhase.value = SessionPhase.SignedOut(SignInFailure.of(outcome)?.let(SignedOutNotice::SignInFailed))
        }
    }

    /** The tap on F1.1 clears its line (spec F1.4: held until the next tap). */
    fun dismissNotice() {
        val current = phase.value
        if (current is SessionPhase.SignedOut && current.notice != null) mutablePhase.value = SessionPhase.SignedOut()
    }

    fun pick(key: String) {
        picked.save(key)
        mutablePicked.value = key
    }

    /** You › Sign out: revoke this device on the server, then forget everything local. */
    suspend fun signOut() {
        endingOnPurpose = true
        try {
            signInFlow.signOut()
            end(null)
        } finally {
            endingOnPurpose = false
        }
    }

    suspend fun devices(now: Instant = Instant.now()): ApiResult<List<DeviceRow>> =
        when (val result = client.send(ApiEndpoints.listDevices())) {
            is ApiResult.Success -> ApiResult.Success(DeviceList.rows(result.data, thisDeviceId(), now), result.requestId)
            is ApiResult.Failure -> result
        }

    suspend fun revoke(id: String): AccountActionOutcome = when (val result = client.send(ApiEndpoints.revokeDevice(id))) {
        is ApiResult.Success -> AccountActionOutcome.Done
        is ApiResult.Failure -> AccountActionOutcome.of(result.error)
    }

    /**
     * `DELETE /v1/me` after the confirmation. On 202 the server has signed out every device, so
     * the app forgets its token without calling logout and shows the date the data goes.
     */
    suspend fun deleteAccount(): AccountActionOutcome =
        when (val result = client.send(ApiEndpoints.deleteAccount(CONFIRMED_DELETION))) {
            is ApiResult.Success -> {
                val completeBy = runCatching { Instant.parse(result.data.completeBy) }.getOrNull()
                endingOnPurpose = true
                try {
                    end(completeBy?.let { SignedOutNotice.AccountDeleted(it) })
                } finally {
                    endingOnPurpose = false
                }
                AccountActionOutcome.Done
            }
            is ApiResult.Failure -> AccountActionOutcome.of(result.error)
        }

    /** A request learned the token is dead: signed out, with the short line, without asking. */
    suspend fun tokenDied() {
        if (endingOnPurpose || !phase.value.isSignedIn) return
        end(SignedOutNotice.SessionEnded)
    }

    private suspend fun end(notice: SignedOutNotice?) {
        tokens.clear()
        picked.save(null)
        mutablePicked.value = null
        clearCaches()
        // TODO(lane A2, F7): unregister this device's push token before the token goes.
        mutablePhase.value = SessionPhase.SignedOut(notice)
    }
}
