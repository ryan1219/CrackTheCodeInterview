package practiceQuestions.twitter;

import static org.junit.Assert.fail;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;

@RunWith(Parameterized.class)
public class TwitterTest {

	final private String testCase;
	final private List<String> commands;
	final private List<List<Integer>> inputs;
	final String expectedOutput;

	public TwitterTest(String testCase, List<String> commands, List<List<Integer>> inputs, String expectedOutput) {
		this.testCase = testCase;
		this.commands = commands;
		this.inputs = inputs;
		this.expectedOutput = expectedOutput;
	}

	@Parameterized.Parameters
	public static Collection<Object[]> testParams() {
		return Arrays.asList(new Object[][] {
				{ "Test Case 0",
						Arrays.asList("Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed",
								"unfollow", "getNewsFeed"),
						Arrays.asList(List.of(), List.of(1, 5), List.of(1), List.of(1, 2), List.of(2, 6), List.of(1),
								List.of(1, 2), List.of(1)),
						"[null,null,[5],null,null,[6,5],null,[5]]" },
				{ "Test Case 1",
						Arrays.asList("Twitter", "postTweet", "getNewsFeed", "follow", "getNewsFeed", "unfollow",
								"getNewsFeed"),
						Arrays.asList(List.of(), List.of(1, 1), List.of(1), List.of(2, 1), List.of(2), List.of(2, 1),
								List.of(2)),
						"[null,null,[1],null,[1],null,[]]" },
				{ "Test Case 2", Arrays.asList("Twitter", "postTweet", "follow", "getNewsFeed"),
						Arrays.asList(List.of(), List.of(1, 5), List.of(1, 1), List.of(1)), "[null,null,null,[5]]" },
				{ "Test Case 3",
						Arrays.asList("Twitter", "postTweet", "postTweet", "unfollow", "follow", "getNewsFeed"),
						Arrays.asList(List.of(), List.of(1, 4), List.of(2, 5), List.of(1, 2), List.of(1, 2),
								List.of(1)),
						"[null,null,null,null,null,[5,4]]" },
				{ "Test Case 4", Arrays.asList("Twitter", "postTweet", "follow", "follow", "getNewsFeed"),
						Arrays.asList(List.of(), List.of(2, 5), List.of(1, 2), List.of(1, 2), List.of(1)),
						"[null,null,null,null,[5]]" },
				{ "Test Case 5",
						Arrays.asList("Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet",
								"postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed",
								"follow", "getNewsFeed"),
						Arrays.asList(List.of(), List.of(2, 5), List.of(1, 3), List.of(1, 101), List.of(2, 13),
								List.of(2, 10), List.of(1, 2), List.of(2, 94), List.of(2, 505), List.of(1, 333),
								List.of(1, 22), List.of(2), List.of(2, 1), List.of(2)),
						"[null,null,null,null,null,null,null,null,null,null,null,[505,94,10,13,5],null,[22,333,505,94,2,10,13,101,3,5]]" } });
	}

	@Test
	public void executeTest() {
		Twitter twitter = new Twitter();
		List<String> outputList = new ArrayList<>();
		System.out.println("Executing " + testCase);
		for (int i = 0; i < commands.size(); i++) {
			List<Integer> input = inputs.get(i);
			switch (commands.get(i)) {
			case "Twitter":
				outputList.add("null");
				break;
			case "postTweet":
				twitter.postTweet(input.get(0), input.get(1));
				outputList.add("null");
				break;
			case "getNewsFeed":
				outputList.add(twitter.getNewsFeed(input.get(0)).toString());
				break;
			case "unfollow":
				twitter.unfollow(input.get(0), input.get(1));
				outputList.add("null");
				break;
			case "follow":
				twitter.follow(input.get(0), input.get(1));
				outputList.add("null");
				break;
			}
		}
		String output = outputList.toString().replaceAll("\\s", "");
		if (!output.equals(expectedOutput)) {
			System.out.println("---------------------");
			System.out.println("Assertion Error");
			System.out.println("---------------------");
			System.out.println("Input");
			System.out.println(commands);
			System.out.println(inputs);
			System.out.println("Output");
			System.out.println(output);
			System.out.println("Expected");
			System.out.println(expectedOutput);
			System.out.println("---------------------");
			fail();
		}
		System.out.println(testCase + " successfully executed");
	}
}
