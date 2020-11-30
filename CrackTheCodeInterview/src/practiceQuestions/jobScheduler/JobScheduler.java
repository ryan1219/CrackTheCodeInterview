package practiceQuestions.jobScheduler;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Collections;
import java.util.Comparator;
import java.util.LinkedList;
import java.util.Scanner;

public class JobScheduler {
	public static void main(String[] args) {
		String input = "path of test file";
		schedule(input);
	}
	
	public static void schedule(String path) {
		LinkedList<Value> list = buildMap(path);
		// sort each job basing on starting time
		Collections.sort(list, new Comparator<Value>() {
			@Override
			public int compare(Value o1, Value o2) {
				return o1.startTime - o2.startTime;
			}
		});
		LinkedList<LinkedList<Value>> result = new LinkedList<>();
		for (Value job : list) {
			boolean processed = false;
			for (LinkedList<Value> l : result) {
				if (l.getLast().startTime + l.getLast().duration < job.startTime) {
					l.add(job);
					processed = true;
					break;
				}

			}

			if (!processed) {
				result.add(new LinkedList<>());
				result.getLast().add(job);
			}
		}

		// print the output in console
		int workerID = 0;
		for (LinkedList<Value> l : result) {
			System.out.println("worker" + workerID);
			for (Value v : l) {
				System.out.println(v.name);
			}
			workerID += 1;
		}
	}

	public static LinkedList<Value> buildMap(String path) {
		LinkedList<Value> list = new LinkedList<>();
		int id = 0;
		try {
			File myObj = new File(path);
			Scanner myReader = new Scanner(myObj);
			while (myReader.hasNextLine()) {
				String data = myReader.nextLine();
				String[] items = data.split(" ");
				if (items.length != 2) {
					continue;
				} else {
					list.add(new Value(Integer.parseInt(items[0]), Integer.parseInt(items[1]), "J" + id));
					id++;
				}

			}
			myReader.close();
		} catch (FileNotFoundException e) {
			System.out.println("An error occurred.");
			e.printStackTrace();
		}

		return list;
	}
}

class Value {
	String name;
	int startTime;
	int duration;

	public Value(int startTime, int duration, String name) {
		super();
		this.name = name;
		this.startTime = startTime;
		this.duration = duration;
	}

	@Override
	public String toString() {
		return "Value [name=" + name + ", startTime=" + startTime + ", duration=" + duration + "]";
	}
}
