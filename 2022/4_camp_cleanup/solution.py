from input import pairs as pairs


def is_any_set_a_subset(first: set, second: set):
    if first.issubset(second) or second.issubset(first):
        return True
    return False

def are_overlapping(first: set, second: set):
    return bool(first.intersection(second))    

def get_range_from_tuple(range_as_tuple):
    return range(range_as_tuple[0], range_as_tuple[1] + 1)

number_of_subset_pairs = 0
number_of_overlapping_pairs = 0
for pair in pairs:
    first_sections = get_range_from_tuple(pair[0])
    second_sections = get_range_from_tuple(pair[1])

    if is_any_set_a_subset(set(first_sections), set(second_sections)):
        number_of_subset_pairs += 1
    if are_overlapping(set(first_sections), set(second_sections)):
        number_of_overlapping_pairs += 1
       

print(f"First part: {number_of_subset_pairs}")

print(f"Second part: {number_of_overlapping_pairs}")

    