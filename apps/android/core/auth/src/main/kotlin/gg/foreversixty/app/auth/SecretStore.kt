package gg.foreversixty.app.auth

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.core.DataStoreFactory
import androidx.datastore.core.Serializer
import com.google.crypto.tink.Aead
import gg.foreversixty.app.model.ApiJson
import kotlinx.coroutines.flow.first
import kotlinx.serialization.builtins.MapSerializer
import kotlinx.serialization.builtins.serializer
import java.io.File
import java.io.InputStream
import java.io.OutputStream
import java.security.GeneralSecurityException

/** Small secrets kept beside the device token (the calendar link, W3-C17), one value per key. */
interface SecretStore {
    suspend fun read(key: String): String?
    suspend fun write(key: String, value: String)

    /** Sign-out: nothing of the last account stays. */
    suspend fun clear()
}

/** A [SecretStore] on DataStore whose file holds only Tink AEAD ciphertext, like the token. */
class EncryptedSecretStore(private val store: DataStore<Map<String, String>>) : SecretStore {
    override suspend fun read(key: String): String? = store.data.first()[key]

    override suspend fun write(key: String, value: String) {
        store.updateData { it + (key to value) }
    }

    override suspend fun clear() {
        store.updateData { emptyMap() }
    }

    companion object {
        /** One store per [file] per process; [purpose] binds its ciphertext to it. */
        fun create(context: Context, file: String, purpose: String): SecretStore {
            val app = context.applicationContext
            val store = DataStoreFactory.create(
                serializer = SealedMapSerializer(DeviceKeys.aead(app), purpose.encodeToByteArray()),
                produceFile = { File(app.filesDir, file) },
            )
            return EncryptedSecretStore(store)
        }
    }
}

/** The map as JSON sealed with [aead]; a blob that fails to decrypt reads as empty. */
internal class SealedMapSerializer(private val aead: Aead, private val associatedData: ByteArray) : Serializer<Map<String, String>> {
    private val json = MapSerializer(String.serializer(), String.serializer())

    override val defaultValue: Map<String, String> = emptyMap()

    override suspend fun readFrom(input: InputStream): Map<String, String> {
        val sealed = input.readBytes()
        if (sealed.isEmpty()) return emptyMap()
        return try {
            ApiJson.decodeFromString(json, aead.decrypt(sealed, associatedData).decodeToString())
        } catch (_: GeneralSecurityException) {
            emptyMap()
        }
    }

    override suspend fun writeTo(t: Map<String, String>, output: OutputStream) {
        if (t.isEmpty()) return
        output.write(aead.encrypt(ApiJson.encodeToString(json, t).encodeToByteArray(), associatedData))
    }
}
