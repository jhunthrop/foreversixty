import Telemetry
import Testing

@Suite("Crash report scrubbing")
struct ScrubberTests {
    @Test("a URL loses its query and fragment", arguments: [
        ("https://foreversixty.gg/app/auth?code=abc&state=xyz", "https://foreversixty.gg/app/auth"),
        ("https://api.foreversixty.gg/v1/me#frag", "https://api.foreversixty.gg/v1/me"),
        ("https://user:pw@api.foreversixty.gg/v1/me", "https://api.foreversixty.gg/v1/me"),
        ("https://api.foreversixty.gg/v1/sims/12", "https://api.foreversixty.gg/v1/sims/12"),
    ])
    func url(input: String, expected: String) {
        #expect(Scrubber.url(input) == expected)
    }

    @Test("free text loses tokens and credential parameters", arguments: [
        ("token fsd_AbC-123_x expired", "token fsd_[redacted] expired"),
        ("Authorization: Bearer abc.def-ghi", "Authorization: Bearer [redacted]"),
        ("callback /app/auth?code=s3cr3t&state=q failed", "callback /app/auth?code=[redacted]&state=[redacted] failed"),
        ("body code_verifier=v1 and access_token=t2", "body code_verifier=[redacted] and access_token=[redacted]"),
        ("nothing secret here", "nothing secret here"),
    ])
    func text(input: String, expected: String) {
        #expect(Scrubber.text(input) == expected)
    }

    @Test("text that is not a URL is scrubbed as text")
    func notAURL() {
        #expect(Scrubber.url("/app/auth?code=abc") == "/app/auth?code=[redacted]")
    }
}

@Suite("Crash reporting policy")
struct CrashReportingPolicyTests {
    let dsn = "https://public@o0.ingest.sentry.io/0"

    @Test("on only for a release build on the live API with a DSN")
    func onlyInRelease() {
        #expect(CrashReportingPolicy(dsn: dsn, isDebugBuild: false, isLiveAPI: true, version: "1.2.0", build: "42") != nil)
        #expect(CrashReportingPolicy(dsn: dsn, isDebugBuild: true, isLiveAPI: true, version: "1.2.0", build: "42") == nil)
        #expect(CrashReportingPolicy(dsn: dsn, isDebugBuild: false, isLiveAPI: false, version: "1.2.0", build: "42") == nil)
        #expect(CrashReportingPolicy(dsn: "", isDebugBuild: false, isLiveAPI: true, version: "1.2.0", build: "42") == nil)
        #expect(CrashReportingPolicy(dsn: nil, isDebugBuild: false, isLiveAPI: true, version: "1.2.0", build: "42") == nil)
    }

    @Test("the release names the app, version and build")
    func release() {
        let policy = CrashReportingPolicy(dsn: dsn, isDebugBuild: false, isLiveAPI: true, version: "1.2.0", build: "42")
        #expect(policy?.release == "gg.foreversixty.app@1.2.0+42")
    }
}
