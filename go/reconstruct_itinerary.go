package main

import "sort"

/*
question: https://leetcode.com/problems/reconstruct-itinerary/

related question: Find Itinerary from a given list of tickets
question: https://www.geeksforgeeks.org/find-itinerary-from-a-given-list-of-tickets/
in order to find the starting point, compare the map with the reversed map,
if the key in map is not in the reversed map, this key is the starting point

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge(u, v),
vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.

Kahn's Algorithm
In this algorithm, we first compute the in-degree values for all vertices.
Then, we start with a vertex whose in-degree is 0 and put it into the end of the output list. Once we choose a vertex,
we update the in-degree values of its adjacent vertices because the vertex and its out edges are removed from the graph.
We can repeat the above process until we choose all vertices. The output list is then a topological sort of the graph.

Depth-first search for Topological sort
The algorithm loops through each node of the graph, in an arbitrary order,
initiating a depth-first search that terminates when it hits any node that has already been visited since the beginning of the topological sort or the node has no outgoing edges (i.e. a leaf node)

L ← Empty list that will contain the sorted nodes
while exists unvisited nodes do
    select an unvisited node n
    visit(n)

function visit(node n)
    if n has a permanent mark then
        return

    for each node m with an edge from n to m do
        visit(m)

    mark n visited
    add n to head of L
*/

/* Hierholzer’s algorithm to find a Eulerian path
Eulerian path is a path that traverses every edge of a graph
1. find the starting vertex v, Initialize 2 stack, cpath and epath
2. push v to cpath
3. let u = cpath.top
4. - if all the edges from u are visited, pop u from cpath and push u to epath
   - if there are edges from u are not visited, select one edge, push the vertex to cpath and delete the path
5. repead 3-4
*/
func findItinerary(tickets [][]string) []string {
	path := make([]string, 0)
	stack := make([]string, 0)
	graph := make(map[string][]string)

	for _, ticket := range tickets {
		_, contains := graph[ticket[0]]
		if !contains {
			graph[ticket[0]] = []string{ticket[1]}
		} else {
			graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
		}
	}
	// sort each value in lexical order
	for key, value := range graph {
		sort.Strings(value)
		graph[key] = value
	}

	stack = append(stack, "JFK")

	for len(stack) > 0 {
		val, exists := graph[stack[len(stack)-1]]
		for exists && len(val) > 0 {
			graph[stack[len(stack)-1]] = val[1:]
			stack = append(stack, val[0])
			val, exists = graph[stack[len(stack)-1]]
		}
		path = prepend(path, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return path
}

func prepend(src []string, value string) []string {
	if len(src) == 0 {
		return append(src, value)
	}
	src = append(src, "")
	copy(src[1:], src)
	src[0] = value
	return src
}
