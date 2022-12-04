from input import calories as calories

calories_caried = (sum(cal) for cal in calories)
most_calories_carried = max(calories_caried)     

print(f"First part: {most_calories_carried}")

sorted_calories = sorted(set(calories_caried), reverse=True)

sum_of_3_most_calories_carried = sum(sorted_calories[:3])

print(f"Second part: {sum_of_3_most_calories_carried}")

    