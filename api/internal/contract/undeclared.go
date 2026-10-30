package contract

import (
	"fmt"
	"slices"
	"strings"
)

// undeclared lists the paths of every object key in value that schema does not declare, so a
// fixture field missing from api/openapi.yaml fails the contract test instead of passing as
// an allowed extra. schema is a generic YAML schema; refs resolves a local "#/components/..."
// pointer. The rules follow JSON Schema's for properties:
//   - allOf, oneOf and anyOf members pool their properties;
//   - additionalProperties (a schema or true) admits any key, a schema checking its value;
//   - an object schema with neither properties nor additionalProperties is open.
func undeclared(schema map[string]any, value any, refs func(string) map[string]any) []string {
	w := walker{refs: refs}
	w.walk("$", schema, value)
	return w.found
}

type walker struct {
	refs  func(string) map[string]any
	found []string
}

func (w *walker) walk(path string, schema map[string]any, value any) {
	members := w.flatten(schema)
	switch v := value.(type) {
	case map[string]any:
		w.walkObject(path, members, v)
	case []any:
		items := w.items(members)
		if items == nil {
			return
		}
		for i, item := range v {
			w.walk(fmt.Sprintf("%s[%d]", path, i), items, item)
		}
	}
}

func (w *walker) walkObject(path string, members []map[string]any, value map[string]any) {
	props := map[string]map[string]any{}
	var extra map[string]any
	open, declares := false, false
	for _, m := range members {
		if p, ok := m["properties"].(map[string]any); ok {
			declares = true
			for key, s := range p {
				if sub, ok := s.(map[string]any); ok && (props[key] == nil || informative(sub)) {
					props[key] = sub
				}
			}
		}
		switch ap := m["additionalProperties"].(type) {
		case bool:
			open = open || ap
		case map[string]any:
			extra = ap
		}
	}
	if !declares && extra == nil {
		open = true
	}
	for _, key := range sortedMapKeys(value) {
		child := path + "." + key
		switch {
		case props[key] != nil:
			w.walk(child, props[key], value[key])
		case extra != nil:
			w.walk(child, extra, value[key])
		case !open:
			w.found = append(w.found, child)
		}
	}
}

// informative is a property schema that says more than a description, so it wins over an
// allOf sibling's bare "data: any JSON" for the same key.
func informative(schema map[string]any) bool {
	for key := range schema {
		if key != "description" {
			return true
		}
	}
	return false
}

func (w *walker) items(members []map[string]any) map[string]any {
	for _, m := range members {
		if items, ok := m["items"].(map[string]any); ok {
			return items
		}
	}
	return nil
}

// flatten resolves refs and pools allOf, oneOf and anyOf into one list of member schemas.
func (w *walker) flatten(schema map[string]any) []map[string]any {
	if ref, ok := schema["$ref"].(string); ok {
		target := w.refs(ref)
		if target == nil {
			w.found = append(w.found, "unresolved "+ref)
			return nil
		}
		return w.flatten(target)
	}
	out := []map[string]any{schema}
	for _, key := range []string{"allOf", "oneOf", "anyOf"} {
		list, _ := schema[key].([]any)
		for _, member := range list {
			if m, ok := member.(map[string]any); ok {
				out = append(out, w.flatten(m)...)
			}
		}
	}
	return out
}

// localRefs resolves "#/a/b" pointers inside root.
func localRefs(root map[string]any) func(string) map[string]any {
	return func(ref string) map[string]any {
		var node any = root
		for _, seg := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
			m, ok := node.(map[string]any)
			if !ok {
				return nil
			}
			node = m[strings.ReplaceAll(strings.ReplaceAll(seg, "~1", "/"), "~0", "~")]
		}
		out, _ := node.(map[string]any)
		return out
	}
}

// sortedMapKeys is a map's keys in order, so walks and test names are stable.
func sortedMapKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
