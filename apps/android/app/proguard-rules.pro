# kotlinx.serialization and Ktor ship their own consumer rules; nothing app-specific yet.
# Ktor's OkHttp engine references optional classes it never loads on Android.
-dontwarn org.slf4j.impl.StaticLoggerBinder
-dontwarn java.lang.management.**
