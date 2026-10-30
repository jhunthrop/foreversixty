package gg.foreversixty.app.auth

import androidx.datastore.core.DataStoreFactory
import com.google.crypto.tink.Aead
import com.google.crypto.tink.KeysetHandle
import com.google.crypto.tink.RegistryConfiguration
import com.google.crypto.tink.aead.AeadConfig
import com.google.crypto.tink.aead.PredefinedAeadParameters
import kotlinx.coroutines.test.TestScope
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Rule
import org.junit.Test
import org.junit.rules.TemporaryFolder
import java.io.ByteArrayInputStream
import java.io.ByteArrayOutputStream
import java.io.File

/**
 * The store with a JVM keyset in place of the Android Keystore master key: the sealing,
 * round trip and fallbacks are the same code the device runs.
 */
class EncryptedTokenStoreTest {
    @get:Rule
    val folder = TemporaryFolder()

    private fun aead(): Aead {
        AeadConfig.register()
        return KeysetHandle.newBuilder()
            .addEntry(
                KeysetHandle.generateEntryFromParameters(PredefinedAeadParameters.AES256_GCM)
                    .withRandomId()
                    .makePrimary(),
            )
            .build()
            .getPrimitive(RegistryConfiguration.get(), Aead::class.java)
    }

    private fun TestScope.store(file: File, aead: Aead = aead()) = EncryptedTokenStore(
        DataStoreFactory.create(EncryptedSessionSerializer(aead), scope = backgroundScope) { file },
    )

    private val session = DeviceSession(deviceId = "dev_1", token = "fsd_secret-token", accountLabel = "Thrall#1234")

    @Test
    fun `a saved session reads back and supplies the bearer token`() = runTest {
        val tokens = store(File(folder.root, "s.bin"))

        tokens.save(session)

        assertEquals("fsd_secret-token", tokens.bearerToken())
    }

    @Test
    fun `the file on disk never holds the token in clear`() = runTest {
        val file = File(folder.root, "s.bin")
        store(file).save(session)

        assertFalse(file.readBytes().decodeToString().contains("fsd_secret-token"))
    }

    @Test
    fun `clear signs out`() = runTest {
        val tokens = store(File(folder.root, "s.bin"))
        tokens.save(session)

        tokens.clear()

        assertNull(tokens.bearerToken())
    }

    @Test
    fun `a blob sealed by another key reads as signed out`() = runTest {
        val sealed = ByteArrayOutputStream()
        EncryptedSessionSerializer(aead()).writeTo(session, sealed)

        val read = EncryptedSessionSerializer(aead()).readFrom(ByteArrayInputStream(sealed.toByteArray()))

        assertNull(read)
    }
}
