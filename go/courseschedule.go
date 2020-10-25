package main

/*
question: https://leetcode.com/problems/course-schedule/

Graph DFS
    DFS-iterative (G, s):                                   //Where G is graph and s is source vertex
      let S be stack
      S.push( s )            //Inserting s in stack
      mark s as visited.
      while ( S is not empty):
        //Pop a vertex from stack to visit next
        v  =  S.top( )
        S.pop( )
        //Push all the neighbours of v in stack that are not visited
        for all neighbours w of v in Graph G:
            if w is not visited :
                S.push( w )
                mark w as visited


    DFS-recursive(G, s):
        mark s as visited
        for all neighbours w of s in Graph G:
            if w is not visited:
                DFS-recursive(G, w)
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	for i := 0; i < len(graph); i++ {
		if !dfs(graph, visited, i) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, visited []bool, current int) bool {
	if visited[current] {
		return false
	}
	visited[current] = true

	for i := 0; i < len(graph[current]); i++ {
		if !dfs(graph, visited, graph[current][i]) {
			return false
		}
	}

	visited[current] = false
	return true
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	return graph
}
