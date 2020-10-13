import fileinput

# dictionary of card ranks
rank = {
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    'T': 10,
    'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14
}

hands = []

for line in fileinput.input():
    hand = line.split()
    if len(hand) == 4:
        hands.append(hand)

# print(hands)

# Loop through the list of hands and append two values, one - the type, ie
# straight flush, pair, etc., the other value - the highest card in the hand
# int high
for hand in hands:
    # create list of card ranks
    cards = []

    # append integer values of card ranks
    cards.append(rank[hand[1][0]])
    cards.append(rank[hand[2][0]])
    cards.append(rank[hand[3][0]])

    # sort the hands
    cards.sort()
    hand.extend(cards)

    # test if same suit
    if hand[1][1] == hand[2][1] and hand[2][1] == hand[3][1]:
        same_suit = True
    else:
        same_suit = False

    # determine hand types
    # first append is the hand type, second is the value of hands's highest card
    # (except for pairs)

    # --Straight Flush (6)
    if cards[0] + 1 == cards[1] and cards[1] + 1 == cards[2] and same_suit:
        hand.append(6)
    # --Three of a Kind (5)
    elif cards[0] == cards[1] and cards[1] == cards[2]:
        hand.append(5)
    # --Straight (4)
    elif cards[0] + 1 == cards[1] and cards[1] + 1 == cards[2]:
        hand.append(4)
    # --Flush (3)
    elif same_suit:
        hand.append(3)
    # --Pair (2)
    elif cards[0] == cards[1] or cards[1] == cards[2]:
        hand.append(2)
    # --High Card (1)
    else:
        hand.append(1)

# print(hands)
# loop through hands and determine winner
highest_type = hands[0][7]
winners = [hands[0]]
for hand in hands[1:]:
    if hand[7] > highest_type:
        winners = []
        winners.append(hand)
    elif hand[7] == highest_type:
        winners.append(hand)

if len(winners) == 1:
    print(winners[0][0])
    exit(0)

# print(winners)
# pair case
if winners[0][7] == 2:
    pair = []
    single = []
    for winner in winners:
        pair.append(winner[5])
        if winner[4] != winner[5]:
            single.append(winner[4])
        else:
            single.append(winner[6])

    max_pair = max(pair)
    max_single = max(single)

    winners1 = [i for i in winners if i[5] == max_pair]
    if len(winners1) == 1:
        print(winners1[0][0])
        exit(0)
    winners2 = [i for i in winners1 if i[4] == max_single or i[6] == max_single]
    for w in winners2:
        print(w[0], end=' ')
else:
    one = []
    two = []
    three = []
    for winner in winners:
        three.append(winner[6])
        two.append(winner[5])
        one.append(winner[4])

    max_one = max(one)
    max_two = max(two)
    max_three = max(three)

    winners1 = [i for i in winners if i[6] == max_three]
    if len(winners1) == 1:
        print(winners1[0][0])
        exit(0)
    winners2 = [i for i in winners1 if i[5] == max_two]
    if len(winners2) == 1:
        print(winners2[0][0])
        exit(0)
    winners3 = [i for i in winners2 if i[4] == max_one]
    for w in winners3:
        print(w[0], end=' ')
