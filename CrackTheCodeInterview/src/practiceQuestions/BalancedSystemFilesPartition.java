package practiceQuestions;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/*
 * question: The directory structure of a system disk partition is represented as a tree. 
 * Its n directories are numbered from 0 to n-1, where the root directory has the number 0. 
 * The structure of the tree is defined by a parent array, where parent[i] = j means that the directory i is a direct subdirectory of j. 
 * Since the root directory does not have a parent, it will be represented as a parent[0] = -1. 
 * The value in files_size[i] denotes the sum of the sizes in kilobytes of the files located in directory i, but excluding its subdirectories. 
 * The size of the content of a directory is defined as the size of all of its subdirectories. 
 * Partition the tree into two smaller ones by cutting one branch so that the sizes of the resulting subtrees are as close as possible.
 * 
 * Example
 * parent = [-1,0,0,1,1,2]
 * files_size = [1,2,2,1,1,1]
 * Output: 0
 * 
 */
public class BalancedSystemFilesPartition {

	public static void main(String[] args) {
		List<Integer> parent = new ArrayList<>(Arrays.asList(-1, 0, 0, 0, 0, 3, 4, 6, 0, 3));
		List<Integer> fileSize = new ArrayList<>(
				Arrays.asList(298, 2187, 5054, 266, 1989, 6499, 5450, 2205, 5893, 8095));
		System.out.println(mostBalancedPartition(parent, fileSize));
	}

	public static int mostBalancedPartition(List<Integer> parent, List<Integer> files_size) {
		// Write your code here
		HashMap<Integer, List<Integer>> tree = buildTree(parent);
		System.out.println(tree);
		int totalFileSize = files_size.stream().mapToInt(a -> a).sum();
		int min = totalFileSize;
		// cut each edge, calculate the difference of weight, find the min
		for (int node = 1; node < parent.size(); node++) {
			int parentNode = parent.get(node);
			tree.get(parentNode).remove(new Integer(node));
			int subtreeWeight = weight(node, tree, files_size);
			int tmpDiff = Math.abs((totalFileSize - subtreeWeight) - subtreeWeight);
			if (tmpDiff < min) {
				min = tmpDiff;
			}
			tree.get(parentNode).add(node);
		}
		return min;
	}

	// calculate the weight of the subtree
	public static int weight(int root, HashMap<Integer, List<Integer>> tree, List<Integer> files_size) {
		int weight = 0;
		Queue<Integer> queue = new LinkedList<>();
		queue.add(root);
		while (!queue.isEmpty()) {
			int current = queue.remove();
			weight += files_size.get(current);
			if (tree.containsKey(current)) {
				List<Integer> children = tree.get(current);
				for (Integer i : children) {
					queue.add(i);
				}
			}
		}

		return weight;
	}

	// build adjacency list to represent a graph
	public static HashMap<Integer, List<Integer>> buildTree(List<Integer> parent) {
		HashMap<Integer, List<Integer>> graph = new HashMap<>();
		for (int i = 0; i < parent.size(); i++) {
			if (graph.containsKey(parent.get(i))) {
				graph.get(parent.get(i)).add(i);
			} else {
				graph.put(parent.get(i), new ArrayList<>(Arrays.asList(i)));
			}
		}

		return graph;
	}

}
