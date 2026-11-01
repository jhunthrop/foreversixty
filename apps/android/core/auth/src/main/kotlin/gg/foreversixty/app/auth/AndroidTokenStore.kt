package gg.foreversixty.app.auth

import android.content.Context
import androidx.datastore.core.DataStoreFactory
import com.google.crypto.tink.Aead
import com.google.crypto.tink.KeyTemplates
import com.google.crypto.tink.RegistryConfiguration
import com.google.crypto.tink.aead.AeadConfig
import com.google.crypto.tink.integration.android.AndroidKeysetManager
import java.io.File

/**
 * The production [TokenStore]: an AES-256-GCM keyset wrapped by an Android Keystore master key,
 * sealing a DataStore file in the app's private files. Build one per process.
 */
object AndroidTokenStore {
    private const val SESSION_FILE = "datastore/device_session.bin"

    fun create(context: Context): TokenStore {
        val app = context.applicationContext
        val store = DataStoreFactory.create(
            serializer = EncryptedSessionSerializer(DeviceKeys.aead(app)),
            produceFile = { File(app.filesDir, SESSION_FILE) },
        )
        return EncryptedTokenStore(store)
    }
}

/** The app's one AES-256-GCM keyset, wrapped by an Android Keystore master key. */
internal object DeviceKeys {
    private const val KEYSET_NAME = "forever_device_session_keyset"
    private const val KEYSET_PREFS = "forever_device_session_keys"
    private const val MASTER_KEY_URI = "android-keystore://forever_device_session_master"

    fun aead(context: Context): Aead {
        AeadConfig.register()
        return AndroidKeysetManager.Builder()
            .withSharedPref(context, KEYSET_NAME, KEYSET_PREFS)
            .withKeyTemplate(KeyTemplates.get("AES256_GCM"))
            .withMasterKeyUri(MASTER_KEY_URI)
            .build()
            .keysetHandle
            .getPrimitive(RegistryConfiguration.get(), Aead::class.java)
    }
}
