package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * POST /v1/auth/app/exchange body (lane A1, plan 2026-10-30): swaps the one-time code the
 * sign-in redirect carried, plus the PKCE verifier, for a device token.
 */
@Serializable
data class AppExchangeRequest(
    val code: String,
    @SerialName("code_verifier") val codeVerifier: String,
    /** "android"; the API records it on the device row. */
    val platform: String,
    /** The device name shown on /account, e.g. "Pixel 9". */
    val name: String,
)

/**
 * POST /v1/auth/app/exchange answer. The same shape POST /v1/devices/claim answers
 * (api/internal/auth/handler.go claim): the token is readable this once only.
 */
@Serializable
data class AppExchangeResponse(
    @SerialName("device_id") val deviceId: String,
    val token: String,
    @SerialName("account_label") val accountLabel: String? = null,
)
