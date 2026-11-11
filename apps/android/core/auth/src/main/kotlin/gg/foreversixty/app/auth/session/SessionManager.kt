package gg.foreversixty.app.auth.session

import gg.foreversixty.app.auth.AppSignIn
import gg.foreversixty.app.auth.BrowserLauncher
import gg.foreversixty.app.auth.SignInOutcome
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.AccountDeletionRequest
import gg.foreversixty.app.model.NotificationSettings
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
 * What an account action came to (iOS `AccountActionOutcome`). [NeedsBrowser] is
 * `403 browser_session_required`: the act is done on foreversixty.gg, so the view opens the site
 * and never retries. [ConfirmMismatch] is `400 confirm_mismatch` (the typed confirmation was
 * refused); [Refused] carries the API's own sentence for a 409, 429 or 503, shown verbatim.
 */
sealed interface AccountActionOutcome {
    data object Done : AccountActionOutcome

    data object Failed : AccountActionOutcome

    data object NeedsBrowser : AccountActionOutcome

    data object ConfirmMismatch : AccountActionOutcome

    data class Refused(val message: String) : AccountActionOutcome

    companion object {
        /** The statuses whose message names the reason and the fix (api/openapi.yaml, deleteAccount). */
        internal val VERBATIM_STATUSES = setOf(409, 429, 503)
        internal const val CONFIRM_MISMATCH = "confirm_mismatch"

        internal fun of(error: ApiError): AccountActionOutcome {
            val server = error as? ApiError.Server ?: return Failed
            return when {
                server.needsBrowserSession -> NeedsBrowser
                server.code == CONFIRM_MISMATCH -> ConfirmMismatch
                server.status in VERBATIM_STATUSES -> Refused(server.message)
                else -> Failed
            }
        }
    }
}

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

    /**
     * Reads `GET /v1/me`. A 401 ends the session; any other failure keeps what is on screen. True when
     * the account was read (the menu's Refresh says whether it worked).
     */
    suspend fun refresh(): Boolean {
        if (!phase.value.isSignedIn) return false
        return when (val result = client.send(ApiEndpoints.getMe())) {
            is ApiResult.Success -> {
                if (phase.value.isSignedIn) mutablePhase.value = SessionPhase.SignedIn(result.data)
                true
            }
            is ApiResult.Failure -> {
                if ((result.error as? ApiError.Server)?.isUnauthorized == true) tokenDied()
                false
            }
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

    /** `GET /v1/me/notifications`: the account's per-event, per-channel choices (the drawer's row reads its groups). */
    suspend fun notificationSettings(): ApiResult<NotificationSettings> = client.send(ApiEndpoints.getNotificationSettings())

    suspend fun revoke(id: String): AccountActionOutcome = when (val result = client.send(ApiEndpoints.revokeDevice(id))) {
        is ApiResult.Success -> AccountActionOutcome.Done
        is ApiResult.Failure -> AccountActionOutcome.of(result.error)
    }

    /**
     * `DELETE /v1/me` with what the person typed ([DeletionConfirmation]). On 200 the account is
     * gone and every device signed out, so the app forgets its token without calling logout.
     */
    suspend fun deleteAccount(confirm: String): AccountActionOutcome =
        when (val result = client.send(ApiEndpoints.deleteAccount(AccountDeletionRequest(confirm = confirm)))) {
            is ApiResult.Success -> {
                endingOnPurpose = true
                try {
                    end(SignedOutNotice.AccountDeleted)
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
        // The push token goes with the device on the server (migration 0062): nothing to unregister.
        mutablePhase.value = SessionPhase.SignedOut(notice)
    }
}
