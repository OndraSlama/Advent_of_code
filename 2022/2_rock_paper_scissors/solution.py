from input import games as games

my_hand_alias= {
    "X": "A",
    "Y": "B",
    "Z": "C"
}

points_per_hand = {
    "A": 1,
    "B": 2,
    "C": 3
}

points_for_game = {
    "AA": 3,
    "BB": 3,
    "CC": 3,
    "AB": 6,
    "AC": 0,
    "BA": 0,
    "BC": 6,
    "CA": 6,
    "CB": 0
}



points_won_in_first = 0
for game in games:
    opponents_hand = game[0]
    my_hand = my_hand_alias[game[1]]
    points_won_in_first += points_per_hand[my_hand]
    points_won_in_first += points_for_game[opponents_hand + my_hand]


print(f"First part: {points_won_in_first}")

# X - loose
# Y - draw
# Z - win

def determine_my_hand(opponents_hand: str, my_strategy: str) -> str:
    if my_strategy == "X":
        if opponents_hand == "A":
            return "C"
        if opponents_hand == "B":
            return "A"
        if opponents_hand == "C":
            return "B"
    if my_strategy == "Y":
        return opponents_hand
    if my_strategy == "Z":
        if opponents_hand == "A":
            return "B"
        if opponents_hand == "B":
            return "C"
        if opponents_hand == "C":
            return "A"


points_won_in_second = 0
for game in games:
    opponents_hand = game[0]
    my_hand = determine_my_hand(opponents_hand, game[1])
    points_won_in_second += points_per_hand[my_hand]
    points_won_in_second += points_for_game[opponents_hand + my_hand]



print(f"Second part: {points_won_in_second}")

    