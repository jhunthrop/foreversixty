package gg.foreversixty.app.network

/** Live talks to the API; Fixture answers every call from the contract fixtures, offline. */
enum class ApiMode {
    Live,
    Fixture,
    ;

    companion object {
        fun parse(value: String): ApiMode = when (value.lowercase()) {
            "live" -> Live
            "fixture" -> Fixture
            else -> throw IllegalArgumentException("API mode must be live or fixture, got '$value'")
        }
    }
}

/** Where and how the client talks. Build it with [ApiConfig.fromBuildConfig] in the app. */
data class ApiConfig(
    val baseUrl: String,
    val mode: ApiMode,
    val requestTimeoutMillis: Long = DEFAULT_REQUEST_TIMEOUT_MILLIS,
    val connectTimeoutMillis: Long = DEFAULT_CONNECT_TIMEOUT_MILLIS,
) {
    init {
        require(baseUrl.startsWith("https://")) { "API base URL must be https, got '$baseUrl'" }
    }

    companion object {
        const val DEFAULT_REQUEST_TIMEOUT_MILLIS = 15_000L
        const val DEFAULT_CONNECT_TIMEOUT_MILLIS = 10_000L

        /** The base URL and mode this build was compiled with (see core/network/build.gradle.kts). */
        fun fromBuildConfig(): ApiConfig =
            ApiConfig(baseUrl = BuildConfig.API_BASE_URL, mode = ApiMode.parse(BuildConfig.API_MODE))
    }
}

/** Supplies the device bearer token (`fsd_…`) for each request; null when signed out. */
interface TokenProvider {
    suspend fun bearerToken(): String?

    /**
     * The API answered 401 (`unauthorized` or `session_expired`) to a request that carried the
     * token: it is dead, so forget it and let the app sign in again.
     */
    suspend fun discardBearerToken()
}

/** No token: every request goes out without one (fixture previews, sign-in tests). */
object NoToken : TokenProvider {
    override suspend fun bearerToken(): String? = null
    override suspend fun discardBearerToken() = Unit
}
