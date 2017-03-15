package routing

import "strings"

type tree struct {
	// Each node can have multiple children
	// (ie. /some/path and /some/thing, "some" would have "path" and "thing" child node).
	children []*tree

	// The part of the path (if path was /some/path, "some" and "path" would be a segment).
	segment string

	// Used to check if segment should be added to params when looking
	// up a route.
	isNamedParam bool

	// The route will be set on leaf nodes only.
	route *Route
}

func newTree() *tree {
	return &tree{segment: "/", isNamedParam: false, route: nil}
}

func (t *tree) addRoute(route *Route) {
	segments := strings.Split(route.Path(), "/")[1:]
	numSegments := len(segments)

	for {
		t, segment, _ := t.traverse(segments)

		// Update an existing segment node
		if t.segment == segment && numSegments == 1 {
			t.route = route
			return
		}

		newTreeNode := &tree{
			segment:      segment,
			isNamedParam: false,
			route:        nil,
		}

		// Check for named route
		if len(segment) > 0 && segment[0] == ':' {
			newTreeNode.isNamedParam = true
		}

		// This is the last segment of the path so add the route.
		if numSegments == 1 {
			newTreeNode.route = route
		}

		t.children = append(t.children, newTreeNode)
		numSegments--

		if numSegments == 0 {
			break
		}
	}
}

func (t *tree) traverse(segments []string) (*tree, string, map[string]string) {
	segment := segments[0]
	params := map[string]string{}

	if len(t.children) > 0 {
		for _, child := range t.children {
			if segment == child.segment || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params[child.segment[1:]] = segment
				}

				next := segments[1:]
				if len(next) > 0 {
					return child.traverse(next)
				}

				return child, segment, params
			}
		}
	}

	return t, segment, params
}

func (t *tree) findRoute(path string) (*Route, map[string]string) {
	node, _, params := t.traverse(strings.Split(path, "/")[1:])

	return node.route, params
}
