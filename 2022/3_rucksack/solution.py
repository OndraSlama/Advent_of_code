from input import rucksacks_example as rucksacks
a_z = [chr(i) for i in range(ord('a'), ord('z') + 1)]
A_Z = [chr(i) for i in range(ord('A'), ord('Z') + 1)]

items_in_priority_order = [*a_z, *A_Z]
priorities = {item: i + 1 for i, item in enumerate(items_in_priority_order)}

def _get_common_chars(*strings) -> set:
    return set.intersection(*map(set, strings))

def get_priority_sum_of_misplaced_items():
    sum_of_priorities = 0
    for items in rucksacks:
        first_comp = items[:len(items) // 2]
        sencd_comp = items[len(items) // 2:]
        if common_item := _get_common_chars(first_comp, sencd_comp).pop():
            sum_of_priorities += priorities[common_item]
    return sum_of_priorities

print(f"First part: {get_priority_sum_of_misplaced_items()}")

def _yield_chunked_list(items, chunk_size):
    for i in range(0, len(items), chunk_size):
        yield items[i:i+chunk_size]
        
def get_sum_of_group_badges():
    sum_of_badges = 0
    for group in _yield_chunked_list(rucksacks, 3):
        badge = _get_common_chars(*group).pop()
        sum_of_badges += priorities[badge]
    return sum_of_badges

print(f"Second part: {get_sum_of_group_badges()}")

    