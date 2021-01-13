package practiceQuestions.twitter;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;

/*
 * 	Design a simplified version of Twitter where users can post tweets, follow/unfollow another user and is able to see the 10 most recent tweets in the user's news feed. Your design should support the following methods:
		
	postTweet(userId, tweetId): Compose a new tweet.	
	getNewsFeed(userId): Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
	follow(followerId, followeeId): Follower follows a followee.	
	unfollow(followerId, followeeId): Follower unfollows a followee.
	
	Example:	
	Twitter twitter = new Twitter();
	// User 1 posts a new tweet (id = 5).
	twitter.postTweet(1, 5);
	
	// User 1's news feed should return a list with 1 tweet id -> [5].
	twitter.getNewsFeed(1);
		
	// User 1 follows user 2.	
	twitter.follow(1, 2);
		
	// User 2 posts a new tweet (id = 6).
	twitter.postTweet(2, 6);
		
	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	twitter.getNewsFeed(1);
		
	// User 1 unfollows user 2.
	twitter.unfollow(1, 2);
		
	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	twitter.getNewsFeed(1);
 */
public class Twitter {
	/** Initialize your data structure here. */
	// key is the followee, value are the followers
	HashMap<Integer, List<Integer>> userMap;
	// key is the user, value is the list of tweets sent by user
	HashMap<Integer, List<Tweet>> users;
	// global index id for the tweet
	int globalIndex = 0;

	public Twitter() {
		userMap = new HashMap<>();
		users = new HashMap<>();
	}

	/** Compose a new tweet. */

	public void postTweet(int userId, int tweetId) {
		if (users.containsKey(userId)) {
			users.get(userId).add(new Tweet(tweetId, globalIndex));
		} else {
			users.put(userId, new ArrayList<>(Arrays.asList(new Tweet(tweetId, globalIndex))));
		}

		globalIndex++;
	}

	/**
	 * Retrieve the 10 most recent tweet ids in the user's news feed. Each item in
	 * the news feed must be posted by users who the user followed or by the user
	 * herself. Tweets must be ordered from most recent to least recent.
	 */

	public List<Integer> getNewsFeed(int userId) {
		List<Tweet> tweets = new ArrayList<>();
		// tweets from userself
		if (users.containsKey(userId)) {
			tweets.addAll(users.get(userId));
		}
		// tweets from user followed
		for (int user : userMap.keySet()) {
			if (userMap.get(user).contains(userId) && users.containsKey(user)) {
				tweets.addAll(users.get(user));
			}
		}

		Collections.sort(tweets, new Comparator<Tweet>() {
			public int compare(Tweet t1, Tweet t2) {
				return t2.index - t1.index;
			}
		});

		List<Integer> result = new ArrayList<>();
		int i = 0;
		while (i < 10 && i < tweets.size()) {
			result.add(tweets.get(i).id);
			i++;
		}

		return result;
	}

	/**
	 * Follower follows a followee. If the operation is invalid, it should be a
	 * no-op.
	 */

	public void follow(int followerId, int followeeId) {
		if (followerId == followeeId) {
			return;
		}
		if (userMap.containsKey(followeeId)) {
			// don't allow user follow twice
			if (!userMap.get(followeeId).contains(followerId)) {
				userMap.get(followeeId).add(followerId);
			}
		} else {
			userMap.put(followeeId, new ArrayList<>(Arrays.asList(followerId)));
		}
	}

	/**
	 * Follower unfollows a followee. If the operation is invalid, it should be a
	 * no-op.
	 */

	public void unfollow(int followerId, int followeeId) {
		if (userMap.containsKey(followeeId)) {
			userMap.get(followeeId).remove(Integer.valueOf(followerId));
		}
	}
}

class Tweet {
	int id;
	int index;

	public Tweet(int id, int index) {
		this.id = id;
		this.index = index;
	}
}
