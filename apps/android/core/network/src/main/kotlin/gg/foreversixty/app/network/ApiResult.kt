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
        /**
         * 401 `unauthorized` or `session_expired`: the device token is missing, revoked or
         * expired. The client has already discarded it; show sign-in.
         */
        val isUnauthorized: Boolean get() = status == HTTP_UNAUTHORIZED

        /** 403 `browser_session_required`: billing writes happen on foreversixty.gg in a browser. */
        val needsBrowserSession: Boolean get() = code == BROWSER_SESSION_REQUIRED
    }

    /** A non-2xx answer with no envelope: a static file the CDN does not have (404), or an outage. */
    data class Status(val status: Int) : ApiError

    /** No answer: offline, DNS, timeout, TLS. Retry later. */
    data class Network(val cause: Throwable) : ApiError

    /** An answer that is not the contract's shape. A bug on one side; report [requestId]. */
    data class Decoding(val status: Int, val requestId: String?, val cause: Throwable) : ApiError

    private companion object {
        const val HTTP_UNAUTHORIZED = 401
        const val BROWSER_SESSION_REQUIRED = "browser_session_required"
    }
}

/** The data on success, else null. */
fun <T> ApiResult<T>.dataOrNull(): T? = when (this) {
    is ApiResult.Success -> data
    is ApiResult.Failure -> null
}
