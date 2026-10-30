package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * POST /v1/auth/app/exchange body (lane A1, docs/contracts/app-v1.md "Sign in"): swaps the
 * one-time code the sign-in redirect carried, plus the PKCE verifier, for a device token.
 */
@Serializable
data class AppExchangeRequest(
    val code: String,
    @SerialName("code_verifier") val codeVerifier: String,
    /** Shown on /account, e.g. "Pixel 9". Optional on the server, which defaults to "Android". */
    @SerialName("device_name") val deviceName: String,
)

/**
 * POST /v1/auth/app/exchange answer (201): the token, readable this once only, the device row
 * it created and the GET /v1/me body, so the first screen needs no second call.
 */
@Serializable
data class AppExchangeResponse(
    val token: String,
    val device: AppDevice,
    val me: Me,
)

/** The device row the exchange created. */
@Serializable
data class AppDevice(
    val id: String,
    val name: String,
    val platform: String,
)

/** POST /v1/auth/app/logout answer. */
@Serializable
data class SignedOut(
    @SerialName("signed_out") val signedOut: Boolean,
)
