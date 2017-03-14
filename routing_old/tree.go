package routing_old

import (
	"net/url"
	"strings"
)

type Tree struct {
	children     []*Tree
	segment      string
	isNamedParam bool
	methods      map[RequestMethod]*Route
}

func (tree *Tree) addNode(method RequestMethod, path string, route *Route) {
	segments := strings.Split(path, "/")[1:]
	segmentCount := len(segments)

	for {
		tree, segment := tree.traverse(segments, nil)

		// Update an existing segment node
		if tree.segment == segment && segmentCount == 1 {
			tree.methods[method] = route
			return
		}

		newTreeNode := Tree{
			segment:      segment,
			isNamedParam: false,
			methods:      make(map[RequestMethod]*Route),
		}

		// Check for named route
		if len(segment) > 0 && segment[0] == ':' {
			newTreeNode.isNamedParam = true
		}

		// This is the last segment of the url so add the route
		if segmentCount == 1 {
			newTreeNode.methods[method] = route
		}

		tree.children = append(tree.children, &newTreeNode)
		segmentCount--

		if segmentCount == 0 {
			break
		}
	}
}

func (tree *Tree) traverse(segments []string, params url.Values) (*Tree, string) {
	segment := segments[0]

	if len(tree.children) > 0 {
		for _, child := range tree.children {
			if segment == child.segment || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.segment[1:], segment)
				}

				next := segments[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				}

				return child, segment
			}
		}
	}

	return tree, segment
}
