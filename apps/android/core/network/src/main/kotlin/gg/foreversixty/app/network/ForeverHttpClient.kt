package gg.foreversixty.app.network

import io.ktor.client.HttpClient
import io.ktor.client.engine.HttpClientEngine
import io.ktor.client.engine.okhttp.OkHttp
import io.ktor.client.plugins.HttpTimeout
import io.ktor.client.plugins.api.createClientPlugin
import io.ktor.client.plugins.defaultRequest
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpStatusCode
import io.ktor.http.URLBuilder
import io.ktor.http.takeFrom

/** Builds the one HttpClient the app uses, for either mode. */
object ForeverHttpClient {

    /**
     * The engine for [config]'s mode: OkHttp for live, the fixture engine for fixture.
     * [fixtures] is required in fixture mode.
     */
    fun engineFor(config: ApiConfig, fixtures: FixtureSource?): HttpClientEngine = when (config.mode) {
        ApiMode.Live -> OkHttp.create()
        ApiMode.Fixture -> FixtureEngine.create(
            requireNotNull(fixtures) { "fixture mode needs a FixtureSource" },
        )
    }

    /**
     * A client rooted at [config]'s base URL that sends `Authorization: Bearer <token>` whenever
     * [tokens] has one, plus [identity]'s `X-App` headers when given (the app always gives them).
     * Status codes do not throw: [ApiClient] reads the envelope on every answer.
     */
    fun create(
        config: ApiConfig,
        tokens: TokenProvider,
        engine: HttpClientEngine,
        identity: ClientIdentity? = null,
    ): HttpClient =
        HttpClient(engine) {
            expectSuccess = false
            install(HttpTimeout) {
                requestTimeoutMillis = config.requestTimeoutMillis
                connectTimeoutMillis = config.connectTimeoutMillis
            }
            install(bearerTokenPlugin(tokens))
            if (identity != null) install(clientIdentityPlugin(identity))
            defaultRequest {
                url.takeFrom(URLBuilder(config.baseUrl))
                headers.append(HttpHeaders.Accept, ContentType.Application.Json.toString())
            }
        }

    /** Sends the token, and discards it when the API answers 401 to a request that carried it. */
    private fun bearerTokenPlugin(tokens: TokenProvider) = createClientPlugin("ForeverBearerToken") {
        onRequest { request, _ ->
            val token = tokens.bearerToken()
            if (token != null && !request.headers.contains(HttpHeaders.Authorization)) {
                request.headers.append(HttpHeaders.Authorization, "Bearer $token")
            }
        }
        onResponse { response ->
            val sentToken = response.call.request.headers.contains(HttpHeaders.Authorization)
            if (response.status == HttpStatusCode.Unauthorized && sentToken) {
                tokens.discardBearerToken()
            }
        }
    }
}
