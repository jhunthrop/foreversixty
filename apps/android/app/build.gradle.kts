plugins {
    alias(libs.plugins.forever.android.application)
    alias(libs.plugins.forever.android.compose)
    alias(libs.plugins.roborazzi)
    alias(libs.plugins.forever.contract.fixtures)
    alias(libs.plugins.forever.app.icon)
    alias(libs.plugins.forever.app.version)
    alias(libs.plugins.forever.bundled.data)
    alias(libs.plugins.kotlin.serialization)
}

/**
 * Release signing comes from the environment (CI writes the upload keystore from secrets).
 * Without it the release build is unsigned, which is fine for `assembleRelease` checks.
 */
val uploadKeystore: String? = providers.environmentVariable("FOREVER_UPLOAD_KEYSTORE").orNull

/**
 * The Sentry DSN for release builds (apps/shared/privacy.md). Empty, the placeholder, keeps crash
 * reporting off; the release command passes SENTRY_DSN_ANDROID. Debug builds never get one.
 */
val releaseSentryDsn: String = providers.environmentVariable("SENTRY_DSN_ANDROID").orElse("").get()

android {
    // BuildConfig.VERSION_NAME and VERSION_CODE feed the X-App header and the crash reporter.
    buildFeatures { buildConfig = true }
    // versionName and versionCode come from apps/VERSION (the forever.app-version plugin).
    defaultConfig {
        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
    }
    signingConfigs {
        if (uploadKeystore != null) {
            create("upload") {
                storeFile = file(uploadKeystore)
                storePassword = providers.environmentVariable("FOREVER_UPLOAD_KEYSTORE_PASSWORD").get()
                keyAlias = providers.environmentVariable("FOREVER_UPLOAD_KEY_ALIAS").get()
                keyPassword = providers.environmentVariable("FOREVER_UPLOAD_KEY_PASSWORD").get()
            }
        }
    }
    buildTypes {
        debug {
            applicationIdSuffix = ".debug"
            buildConfigField("String", "SENTRY_DSN", "\"\"")
        }
        release {
            buildConfigField("String", "SENTRY_DSN", "\"$releaseSentryDsn\"")
            isMinifyEnabled = true
            isShrinkResources = true
            proguardFiles(getDefaultProguardFile("proguard-android-optimize.txt"), "proguard-rules.pro")
            if (uploadKeystore != null) signingConfig = signingConfigs.getByName("upload")
        }
    }
    testOptions {
        unitTests.all {
            it.systemProperty("robolectric.pixelCopyRenderMode", "hardware")
            // Robolectric's SDK 36 sandbox reaches into FileDescriptor internals; JDK 17+ needs them opened.
            it.jvmArgs(
                "--add-opens=java.base/java.io=ALL-UNNAMED",
                "--add-exports=java.base/jdk.internal.access=ALL-UNNAMED",
            )
        }
    }
}

dependencies {
    implementation(project(":core:auth"))
    implementation(project(":core:designsystem"))
    implementation(project(":core:model"))
    implementation(project(":core:network"))
    implementation(project(":core:telemetry"))
    implementation(libs.sentry.android.core)
    implementation(libs.androidx.activity.compose)
    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.lifecycle.runtime.compose)
    implementation(libs.compose.material.icons)
    implementation(libs.coil.compose)
    implementation(libs.coil.network.okhttp)
    testImplementation(libs.kotlinx.coroutines.test)

    testImplementation(libs.junit)
    testImplementation(libs.robolectric)
    testImplementation(libs.roborazzi)
    testImplementation(libs.roborazzi.compose)
    testImplementation(libs.compose.ui.test.junit4)
    testImplementation(libs.androidx.test.ext.junit)
    debugImplementation(libs.compose.ui.test.manifest)
}
