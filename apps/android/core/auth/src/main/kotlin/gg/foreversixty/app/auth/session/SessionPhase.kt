package gg.foreversixty.app.auth.session

import gg.foreversixty.app.auth.SignInOutcome
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.network.ApiError
import java.time.Instant

/** Where the app's account stands; the views draw one screen per phase (spec §5.1). */
sealed interface SessionPhase {
    /** F1.1, with the line the last ending left (F1.4's toast, the 401 line, a deletion). */
    data class SignedOut(val notice: SignedOutNotice? = null) : SessionPhase

    /** F1.2: the Custom Tab is open; the button shows its spinner. */
    data object SigningIn : SessionPhase

    /** F1.3: the redirect arrived; the code is exchanged and the characters are on their way. */
    data object Importing : SessionPhase

    /** Signed in. [me] is null until the first `GET /v1/me` of this launch answers. */
    data class SignedIn(val me: Me?) : SessionPhase
}

val SessionPhase.isSignedIn: Boolean get() = this is SessionPhase.SignedIn

val SessionPhase.me: Me? get() = (this as? SessionPhase.SignedIn)?.me

/** The one line F1.1 shows above its buttons, held until the next tap. */
sealed interface SignedOutNotice {
    data class SignInFailed(val failure: SignInFailure) : SignedOutNotice

    /** A 401 (`unauthorized` or `session_expired`) ended the session without asking. */
    data object SessionEnded : SignedOutNotice

    /** `DELETE /v1/me` was accepted; the data is gone by [completeBy]. */
    data class AccountDeleted(val completeBy: Instant) : SignedOutNotice
}

/** The known sign-in failures (spec F1.4); each has one sentence in the copy catalog. */
enum class SignInFailure {
    /** Battle.net or the redirect refused, the state did not match, or the exchange failed. */
    BattleNet,

    /** The one-time code was spent or expired (`400 invalid_grant`). */
    Expired,

    /** No connection: signing in needs the internet. */
    Offline,
    ;

    companion object {
        private const val INVALID_GRANT = "invalid_grant"

        /** The failure an outcome shows, or null when sign-in succeeded. */
        fun of(outcome: SignInOutcome): SignInFailure? = when (outcome) {
            is SignInOutcome.SignedIn -> null
            is SignInOutcome.Cancelled, SignInOutcome.Rejected -> BattleNet
            is SignInOutcome.ExchangeFailed -> when (val error = outcome.result.error) {
                is ApiError.Network -> Offline
                is ApiError.Server -> if (error.code == INVALID_GRANT) Expired else BattleNet
                else -> BattleNet
            }
        }
    }
}
