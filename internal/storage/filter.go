package storage

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Anvinalias/az-blob-downloader/internal/request"
)

type edge struct {
	from string
	to   string
	file string // blob file name
}

type graph struct {
	edges map[string][]edge
}

// newGraph builds a graph from blob names for a given app prefix.
func newGraph(blobs []string, prefix string) *graph {
	g := &graph{edges: make(map[string][]edge)}
	for _, blob := range blobs {
		name := blob
		ext := filepath.Ext(name)
		name = strings.TrimSuffix(name, ext)

		if !strings.HasPrefix(name, prefix+"-") {
			continue
		}
		if strings.HasSuffix(name, "-release") {
			continue // skip sidecar files
		}

		parts := strings.Split(strings.TrimPrefix(name, prefix+"-"), "-")
		if len(parts) != 2 {
			continue
		}
		from, to := parts[0], parts[1]
		e := edge{from: from, to: to, file: blob}
		g.edges[from] = append(g.edges[from], e)
	}
	return g
}
// shortestPath finds the shortest upgrade path from 'from' to 'to'.
func (g *graph) shortestPath(from, to string) ([]edge, bool) {
	prev := make(map[string]edge)
	visited := make(map[string]bool)
	queue := []string{from}
	visited[from] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == to {
			break
		}
		neighbors := g.edges[current]

		for _, e := range neighbors {
			if !visited[e.to] {
				visited[e.to] = true
				prev[e.to] = e
				queue = append(queue, e.to)
			}
		}
	}
	// Reconstruct path
	var path []edge
	for cur := to; cur != from; {
		e, ok := prev[cur]
		if !ok {
			return nil, false
		}
		path = append([]edge{e}, path...)
		cur = e.from
	}
	return path, true
}

// BuildShortestUpgradePath returns the minimal sequence of upgrade steps as base names (without extensions)
func BuildShortestUpgradePath(blobs []string, req *request.Request) ([]string, error) {
	if req.From == req.To {
		return nil, fmt.Errorf("no upgrade path needed: from and to versions are the same (%s)", req.From)
	}
	g := newGraph(blobs, req.Prefix)
	path, ok := g.shortestPath(req.From, req.To)
	if !ok {
		return nil, fmt.Errorf("no upgrade path found for %s from %s to %s", req.Prefix, req.From, req.To)
	}
	var baseNames []string
	for _, edge := range path {
		baseNames = append(baseNames, req.Prefix+"-"+edge.from+"-"+edge.to)
	}
	return baseNames, nil
}