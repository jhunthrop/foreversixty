package gg.foreversixty.app.network

import io.ktor.client.plugins.api.createClientPlugin

/**
 * Who is calling, sent on every API request (docs/contracts/app-v1.md "Client identity"):
 * `X-App: android/<version> (<build>)` and `X-App-Install: <install id>`. The install id is a
 * random UUID made on first launch, tied to no user, and gone with the app on uninstall.
 */
data class ClientIdentity(val version: String, val build: Int, val installId: String) {
    /** `android/1.2.0 (42)`. */
    val appHeaderValue: String get() = "$PLATFORM/$version ($build)"

    companion object {
        const val APP_HEADER = "X-App"
        const val INSTALL_HEADER = "X-App-Install"
        const val PLATFORM = "android"
    }
}

/** Adds [identity]'s two headers to every request the client sends. */
internal fun clientIdentityPlugin(identity: ClientIdentity) = createClientPlugin("ForeverClientIdentity") {
    onRequest { request, _ ->
        request.headers[ClientIdentity.APP_HEADER] = identity.appHeaderValue
        request.headers[ClientIdentity.INSTALL_HEADER] = identity.installId
    }
}
