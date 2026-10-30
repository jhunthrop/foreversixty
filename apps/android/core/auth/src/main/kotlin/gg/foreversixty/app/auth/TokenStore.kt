package gg.foreversixty.app.auth

import androidx.datastore.core.DataStore
import androidx.datastore.core.Serializer
import com.google.crypto.tink.Aead
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.TokenProvider
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.first
import kotlinx.serialization.Serializable
import java.io.InputStream
import java.io.OutputStream
import java.security.GeneralSecurityException

/** The signed-in device: the API's device id and its `fsd_` bearer token. */
@Serializable
data class DeviceSession(
    val deviceId: String,
    val token: String,
    val accountLabel: String? = null,
)

/** Where the device session lives. Also the network's [TokenProvider]. */
interface TokenStore : TokenProvider {
    val session: Flow<DeviceSession?>
    suspend fun save(session: DeviceSession)
    suspend fun clear()
}

/** A [TokenStore] on DataStore whose file holds only Tink AEAD ciphertext. */
class EncryptedTokenStore(private val store: DataStore<DeviceSession?>) : TokenStore {
    override val session: Flow<DeviceSession?> = store.data

    override suspend fun bearerToken(): String? = store.data.first()?.token

    override suspend fun discardBearerToken() = clear()

    override suspend fun save(session: DeviceSession) {
        store.updateData { session }
    }

    override suspend fun clear() {
        store.updateData { null }
    }
}

/**
 * Serialises the session as JSON sealed with [aead]. The associated data binds the ciphertext
 * to this purpose, so a blob lifted from another store will not decrypt here. A blob that fails
 * to decrypt (keystore reset after a restore, say) reads as signed out rather than crashing.
 */
class EncryptedSessionSerializer(private val aead: Aead) : Serializer<DeviceSession?> {
    override val defaultValue: DeviceSession? = null

    override suspend fun readFrom(input: InputStream): DeviceSession? {
        val sealed = input.readBytes()
        if (sealed.isEmpty()) return null
        return try {
            val plain = aead.decrypt(sealed, ASSOCIATED_DATA)
            ApiJson.decodeFromString(DeviceSession.serializer(), plain.decodeToString())
        } catch (_: GeneralSecurityException) {
            null
        }
    }

    override suspend fun writeTo(t: DeviceSession?, output: OutputStream) {
        if (t == null) return
        val plain = ApiJson.encodeToString(DeviceSession.serializer(), t).encodeToByteArray()
        output.write(aead.encrypt(plain, ASSOCIATED_DATA))
    }

    private companion object {
        val ASSOCIATED_DATA = "gg.foreversixty.app/device-session/v1".encodeToByteArray()
    }
}
