package routing

import (
	"strings"
)

// routeTrie is defined to hide away the traverse method in the trie data structure
// defined below.
type routeTrie interface {
	add(route *Route)
	search(path string) ([]*Route, map[string]string)
}

// trie is a simple trie data structure that allows for multiple routes to be added
// to each node. The trie should be used to get all routes that match a path, any
// matching on methods, regex, host, scheme etc should be handled elsewhere.
type trie struct {
	// Each node can have multiple children
	// (ie. /some/path and /some/thing, "some" would have "path" and "thing" child node).
	children []*trie

	// The part of the path (if path was /some/path, "some" and "path" would be a segment).
	segment string

	// Used to check if segment should be added to params when looking up a route.
	isNamedParam bool

	// Routes will be set on leaf nodes only.
	routes []*Route
}

// newRouteTrie will return a new instance of a trie that implements the routeTrie interface.
// It also sets the first segment to / (this can be overridden by adding a route with "/" as the path.
func newRouteTrie() routeTrie {
	return &trie{segment: "/", isNamedParam: false, routes: nil}
}

// add will insert a new route into the trie.
func (t *trie) add(route *Route) {
	segments := strings.Split(route.Path(), "/")[1:]
	numSegments := len(segments)

	for {
		// Traverse the segments to see if we match, if we don't then the original
		// trie and segment is returned.
		node, segment := t.traverse(segments, nil)

		// If we find a node that matches the last segment then update it and return.
		if len(segments) > 0 && node.segment == segments[len(segments)-1] {
			node.routes = append(node.routes, route)
			return
		}

		newTrieNode := &trie{
			segment:      segment,
			isNamedParam: false,
			routes:       nil,
		}

		// Check for named route.
		if len(segment) > 0 && segment[0] == ':' {
			newTrieNode.isNamedParam = true
		}

		// If the last node is the segment before the one we are adding then add to its child.
		if len(segments) > 1 && node.segment == segments[len(segments)-2] {
			newTrieNode.routes = append(newTrieNode.routes, route)
			node.children = append(node.children, newTrieNode)
			return
		}

		// This is the last segment of the path so add the route.
		if numSegments == 1 {
			newTrieNode.routes = append(newTrieNode.routes, route)
		}

		// Add a new child to the current node.
		node.children = append(node.children, newTrieNode)
		numSegments--

		// Once we have gone through all the segments we are not updating
		// and have added a node for every segment in the path.
		if numSegments == 0 {
			break
		}
	}
}

// search will traverse the trie looking for a match to the path.
// If nothing is found it will return (nil, map[string]string{}).
func (t *trie) search(path string) ([]*Route, map[string]string) {
	params := map[string]string{}

	node, _ := t.traverse(strings.Split(path, "/")[1:], params)

	return node.routes, params
}

// traverse recursively searches the trie based on the path segments and extracts
// named params along the way.
func (t *trie) traverse(segments []string, params map[string]string) (*trie, string) {
	segment := segments[0]

	if len(t.children) > 0 {
		for _, child := range t.children {
			if segment == child.segment || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params[child.segment[1:]] = segment
				}

				next := segments[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				}

				return child, segment
			}
		}
	}

	return t, segment
}
