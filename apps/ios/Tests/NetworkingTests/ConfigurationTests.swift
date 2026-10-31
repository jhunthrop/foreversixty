import Foundation
@testable import Networking
import Testing

@Suite("APIConfiguration")
struct ConfigurationTests {
    @Test("reads base URL and mode from Info.plist keys")
    func reads() throws {
        let config = try APIConfiguration(infoDictionary: [
            "FSAPIBaseURL": "https://api.example.test",
            "FSAPIMode": "fixture",
        ])
        #expect(config == APIConfiguration(baseURL: URL(string: "https://api.example.test")!, mode: .fixture))
    }

    @Test("mode defaults to live")
    func defaultsLive() throws {
        let config = try APIConfiguration(infoDictionary: ["FSAPIBaseURL": "https://api.example.test"])
        #expect(config.mode == .live)
    }

    @Test("a missing base URL fails fast")
    func missing() {
        #expect(throws: ConfigurationError.missing("FSAPIBaseURL")) {
            try APIConfiguration(infoDictionary: [:])
        }
    }

    @Test("plain http is refused outside localhost")
    func refusesHTTP() {
        #expect(throws: ConfigurationError.invalid("FSAPIBaseURL", "http://api.example.test")) {
            try APIConfiguration(infoDictionary: ["FSAPIBaseURL": "http://api.example.test"])
        }
    }

    @Test("an unknown mode fails fast")
    func badMode() {
        #expect(throws: ConfigurationError.invalid("FSAPIMode", "staging")) {
            try APIConfiguration(infoDictionary: ["FSAPIBaseURL": "https://a.test", "FSAPIMode": "staging"])
        }
    }
}
