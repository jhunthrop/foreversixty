package contract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
	liberrors "github.com/pb33f/libopenapi-validator/errors"
	"github.com/pb33f/libopenapi-validator/schema_validation"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"gopkg.in/yaml.v3"
)

// api/openapi.yaml is the apps' source of truth: their models and endpoints are generated from
// it (web/scripts/gen-mobile-api.mjs). Every fixture under docs/contracts/fixtures/app is
// claimed by the spec, either by a response's `x-fixtures` (an API answer or a static file) or
// by a schema's `x-fixtures` (a push payload), and validates against that schema in strict
// mode: a wrong kind, a missing required field, a null where the spec allows none, or a field
// the spec does not declare all fail. The handlers' own contract tests hold their answers to
// the same fixtures, so the spec, the fixtures and the handlers cannot drift apart.

const specMediaType = "application/json"

// specVersion is the OpenAPI version api/openapi.yaml declares.
const specVersion float32 = 3.1

// fixtureClaim is one fixture and where the spec claims it.
type fixtureClaim struct {
	fixture string
	// method, path, status and server name a response; schema names a component schema and
	// response a shared response (its JSON schema).
	method, path, server string
	status               int
	query                url.Values
	schema, response     string
}

func (c fixtureClaim) String() string {
	if c.schema != "" {
		return fmt.Sprintf("%s as schema %s", c.fixture, c.schema)
	}
	if c.response != "" {
		return fmt.Sprintf("%s as response %s", c.fixture, c.response)
	}
	return fmt.Sprintf("%s as %s %s %d", c.fixture, c.method, c.path, c.status)
}

// specDoc is the raw spec as generic YAML, for the x-fixtures claims and the parameters.
type specDoc struct {
	Servers []struct {
		URL string `yaml:"url"`
	} `yaml:"servers"`
	Paths      map[string]map[string]any `yaml:"paths"`
	Components struct {
		Schemas   map[string]map[string]any `yaml:"schemas"`
		Responses map[string]map[string]any `yaml:"responses"`
	} `yaml:"components"`
}

func specPath() string { return filepath.Join(repoRoot(), "api", "openapi.yaml") }

func loadSpec(t *testing.T) ([]byte, specDoc) {
	t.Helper()
	raw, err := os.ReadFile(specPath())
	if err != nil {
		t.Fatal(err)
	}
	var doc specDoc
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("openapi.yaml is not YAML: %v", err)
	}
	return raw, doc
}

// claimSchema is the generic schema a claim validates against, for the undeclared-key walk.
func claimSchema(t *testing.T, root map[string]any, c fixtureClaim) map[string]any {
	t.Helper()
	refs := localRefs(root)
	if c.schema != "" {
		return refs("#/components/schemas/" + c.schema)
	}
	if c.response != "" {
		return refs("#/components/responses/" + c.response + "/content/" + escapePointer(specMediaType) + "/schema")
	}
	pointer := fmt.Sprintf("#/paths/%s/%s/responses/%d/content/%s/schema",
		escapePointer(c.path), strings.ToLower(c.method), c.status, escapePointer(specMediaType))
	schema := refs(pointer)
	if schema == nil {
		t.Fatalf("%s: no %s schema at %s", c, specMediaType, pointer)
	}
	return schema
}

func escapePointer(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "~", "~0"), "/", "~1")
}

var httpMethods = []string{"get", "put", "post", "patch", "delete"}

// claims lists every x-fixtures entry of the spec, in a stable order.
func claims(t *testing.T, doc specDoc) []fixtureClaim {
	t.Helper()
	var out []fixtureClaim
	for _, path := range sortedMapKeys(doc.Paths) {
		item := doc.Paths[path]
		server := firstServer(item["servers"], doc.Servers[0].URL)
		for _, method := range httpMethods {
			op, ok := item[method].(map[string]any)
			if !ok {
				continue
			}
			out = append(out, operationClaims(t, method, path, server, op)...)
		}
	}
	for _, name := range sortedMapKeys(doc.Components.Schemas) {
		for _, fixture := range stringList(doc.Components.Schemas[name]["x-fixtures"]) {
			out = append(out, fixtureClaim{fixture: fixture, schema: name})
		}
	}
	for _, name := range sortedMapKeys(doc.Components.Responses) {
		for _, fixture := range stringList(doc.Components.Responses[name]["x-fixtures"]) {
			out = append(out, fixtureClaim{fixture: fixture, response: name})
		}
	}
	return out
}

func operationClaims(t *testing.T, method, path, server string, op map[string]any) []fixtureClaim {
	t.Helper()
	responses, _ := op["responses"].(map[string]any)
	query := requiredQuery(op)
	var out []fixtureClaim
	for _, code := range sortedMapKeys(responses) {
		response, _ := responses[code].(map[string]any)
		for _, fixture := range stringList(response["x-fixtures"]) {
			status, err := strconv.Atoi(code)
			if err != nil {
				t.Fatalf("%s %s: x-fixtures on response %q, which is not a status", method, path, code)
			}
			out = append(out, fixtureClaim{
				fixture: fixture, method: strings.ToUpper(method), path: path,
				server: server, status: status, query: query,
			})
		}
	}
	return out
}

// requiredQuery gives every required query parameter a value the spec accepts: its first enum
// value (mine=1), else "1".
func requiredQuery(op map[string]any) url.Values {
	q := url.Values{}
	params, _ := op["parameters"].([]any)
	for _, p := range params {
		param, _ := p.(map[string]any)
		if param["in"] != "query" || param["required"] != true {
			continue
		}
		value := "1"
		if schema, ok := param["schema"].(map[string]any); ok {
			if enum, ok := schema["enum"].([]any); ok && len(enum) > 0 {
				value = fmt.Sprint(enum[0])
			}
		}
		q.Set(fmt.Sprint(param["name"]), value)
	}
	return q
}

func firstServer(v any, fallback string) string {
	servers, _ := v.([]any)
	if len(servers) == 0 {
		return fallback
	}
	server, _ := servers[0].(map[string]any)
	return fmt.Sprint(server["url"])
}

var pathParam = regexp.MustCompile(`\{[^}]+\}`)

// concretePath fills every {param} of a path template with a value any parameter accepts.
func concretePath(template string) string { return pathParam.ReplaceAllString(template, "1") }

func newValidator(t *testing.T, raw []byte) validator.Validator {
	t.Helper()
	document, err := libopenapi.NewDocument(raw)
	if err != nil {
		t.Fatalf("libopenapi cannot read openapi.yaml: %v", err)
	}
	v, errs := validator.NewValidator(document)
	if len(errs) > 0 {
		t.Fatalf("libopenapi cannot build openapi.yaml: %v", errs)
	}
	return v
}

func TestEveryFixtureIsClaimedByTheSpec(t *testing.T) {
	_, doc := loadSpec(t)
	claimed := map[string]bool{}
	for _, c := range claims(t, doc) {
		claimed[c.fixture] = true
		if _, err := os.Stat(FixturePath(c.fixture)); err != nil {
			t.Errorf("openapi.yaml claims %s, but there is no such fixture", c)
		}
	}
	files, err := filepath.Glob(filepath.Join(fixtureDir(), "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no fixtures found: %v", err)
	}
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file), ".json")
		if !claimed[name] {
			t.Errorf("%s.json is claimed by no response or schema in api/openapi.yaml: add it to an x-fixtures list", name)
		}
	}
}

func TestEveryFixtureMatchesTheOpenAPISpec(t *testing.T) {
	raw, doc := loadSpec(t)
	v := newValidator(t, raw)
	document, err := libopenapi.NewDocument(raw)
	if err != nil {
		t.Fatal(err)
	}
	model, buildErr := document.BuildV3Model()
	if buildErr != nil {
		t.Fatal(buildErr)
	}
	var root map[string]any
	if err := yaml.Unmarshal(raw, &root); err != nil {
		t.Fatal(err)
	}
	for _, c := range claims(t, doc) {
		t.Run(c.String(), func(t *testing.T) {
			body, err := os.ReadFile(FixturePath(c.fixture))
			if err != nil {
				t.Fatal(err)
			}
			assertDeclared(t, claimSchema(t, root, c), body, localRefs(root))
			if c.schema != "" {
				proxy, ok := model.Model.Components.Schemas.Get(c.schema)
				if !ok {
					t.Fatalf("no schema %s", c.schema)
				}
				assertSchema(t, proxy.Schema(), body)
				return
			}
			if c.response != "" {
				response, ok := model.Model.Components.Responses.Get(c.response)
				if !ok {
					t.Fatalf("no response %s", c.response)
				}
				media, ok := response.Content.Get(specMediaType)
				if !ok {
					t.Fatalf("response %s has no %s", c.response, specMediaType)
				}
				assertSchema(t, media.Schema.Schema(), body)
				return
			}
			assertResponse(t, v, c, body)
		})
	}
}

func assertResponse(t *testing.T, v validator.Validator, c fixtureClaim, body []byte) {
	t.Helper()
	target := strings.TrimSuffix(c.server, "/") + concretePath(c.path)
	if len(c.query) > 0 {
		target += "?" + c.query.Encode()
	}
	req, err := http.NewRequest(c.method, target, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := &http.Response{
		StatusCode: c.status,
		Header:     http.Header{"Content-Type": []string{specMediaType}},
		Body:       io.NopCloser(bytes.NewReader(body)),
		Request:    req,
	}
	if ok, errs := v.ValidateHttpResponse(req, resp); !ok {
		t.Errorf("%s does not match its response schema:\n%s", c.fixture, describe(errs))
	}
}

// assertSchema validates a fixture that is no HTTP answer (a push payload) against a schema,
// strictly, as the responses are.
func assertSchema(t *testing.T, s *base.Schema, body []byte) {
	t.Helper()
	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatal(err)
	}
	if ok, errs := schema_validation.NewSchemaValidator().ValidateSchemaObjectWithVersion(s, data, specVersion); !ok {
		t.Errorf("does not match its schema:\n%s", describe(errs))
	}
}

// assertDeclared fails for every fixture key the spec does not declare.
func assertDeclared(t *testing.T, schema map[string]any, body []byte, refs func(string) map[string]any) {
	t.Helper()
	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatal(err)
	}
	for _, path := range undeclared(schema, data, refs) {
		t.Errorf("%s is in the fixture but not declared in api/openapi.yaml", path)
	}
}

func describe(errs []*liberrors.ValidationError) string {
	var b strings.Builder
	for _, e := range errs {
		fmt.Fprintf(&b, "  - %s: %s\n", e.Message, e.Reason)
		for _, f := range e.SchemaValidationErrors {
			fmt.Fprintf(&b, "      %s %s: %s\n", f.FieldPath, f.KeywordLocation, f.Reason)
		}
	}
	return b.String()
}

func stringList(v any) []string {
	items, _ := v.([]any)
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, fmt.Sprint(item))
	}
	return out
}

// Every app call carries the client identity headers and may meet the minimum-version gate, so
// every operation the apps' clients are generated from says so, through the shared components.
func TestEveryAppOperationDeclaresIdentityAndTheUpdateGate(t *testing.T) {
	_, doc := loadSpec(t)
	want := []string{"#/components/parameters/XApp", "#/components/parameters/XAppInstall"}
	for _, path := range sortedMapKeys(doc.Paths) {
		for _, method := range httpMethods {
			op, ok := doc.Paths[path][method].(map[string]any)
			if !ok || !slices.Contains(stringList(op["tags"]), "app") {
				continue
			}
			var refs []string
			params, _ := op["parameters"].([]any)
			for _, p := range params {
				if m, ok := p.(map[string]any); ok {
					refs = append(refs, fmt.Sprint(m["$ref"]))
				}
			}
			for _, ref := range want {
				if !slices.Contains(refs, ref) {
					t.Errorf("%s %s does not list %s", method, path, ref)
				}
			}
			responses, _ := op["responses"].(map[string]any)
			gate, _ := responses["426"].(map[string]any)
			if gate["$ref"] != "#/components/responses/UpdateRequired" {
				t.Errorf("%s %s has no 426 UpdateRequired response", method, path)
			}
		}
	}
}
