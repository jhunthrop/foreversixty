package gg.foreversixty.app.network

/** The outcome of one API call. Callers `when` over it; nothing here throws. */
sealed interface ApiResult<out T> {
    data class Success<T>(val data: T, val requestId: String) : ApiResult<T>
    data class Failure(val error: ApiError) : ApiResult<Nothing>
}

/** Why a call failed, split by what the UI can do about it. */
sealed interface ApiError {
    /** The API answered `ok: false`. [message] is the API's own player-safe sentence. */
    data class Server(
        val status: Int,
        val code: String,
        val message: String,
        val fields: Map<String, String>,
        val requestId: String,
    ) : ApiError {
        /** The device token is missing, revoked or expired: sign in again. */
        val isUnauthorized: Boolean get() = status == HTTP_UNAUTHORIZED
    }

    /** No answer: offline, DNS, timeout, TLS. Retry later. */
    data class Network(val cause: Throwable) : ApiError

    /** An answer that is not the contract's shape. A bug on one side; report [requestId]. */
    data class Decoding(val status: Int, val requestId: String?, val cause: Throwable) : ApiError

    private companion object {
        const val HTTP_UNAUTHORIZED = 401
    }
}

/** The data on success, else null. */
fun <T> ApiResult<T>.dataOrNull(): T? = when (this) {
    is ApiResult.Success -> data
    is ApiResult.Failure -> null
}
