package gg.foreversixty.app.auth

import java.security.MessageDigest
import java.security.SecureRandom
import java.util.Base64

/** RFC 7636 PKCE with S256, plus the OAuth state nonce. */
object Pkce {
    const val METHOD = "S256"

    /** 32 random bytes give a 43-character verifier, the RFC's minimum length. */
    private const val VERIFIER_BYTES = 32
    private const val STATE_BYTES = 16

    private val base64Url: Base64.Encoder = Base64.getUrlEncoder().withoutPadding()

    fun newVerifier(random: SecureRandom = SecureRandom()): String = randomToken(random, VERIFIER_BYTES)

    fun newState(random: SecureRandom = SecureRandom()): String = randomToken(random, STATE_BYTES)

    /** BASE64URL(SHA256(ASCII(verifier))), unpadded. */
    fun challenge(verifier: String): String {
        require(verifier.length in VERIFIER_LENGTH) { "PKCE verifier must be 43 to 128 characters" }
        val digest = MessageDigest.getInstance("SHA-256").digest(verifier.toByteArray(Charsets.US_ASCII))
        return base64Url.encodeToString(digest)
    }

    private fun randomToken(random: SecureRandom, bytes: Int): String =
        base64Url.encodeToString(ByteArray(bytes).also(random::nextBytes))

    private val VERIFIER_LENGTH = 43..128
}
