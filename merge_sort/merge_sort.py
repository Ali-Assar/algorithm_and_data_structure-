def merge(list1, list2):
    combined = []
    i, j = 0, 0

    while i < len(list1) and j < len(list2):
        if list1[i] < list2[j]:
            combined.append(list1[i])
            i += 1
        else:
            combined.append(list2[j])
            j += 1

    combined.extend(list1[i:])
    combined.extend(list2[j:])
    return combined

def merge_sort(lst):
    if len(lst) <= 1:
        return lst

    mid_index = len(lst) // 2
    left = merge_sort(lst[:mid_index])
    right = merge_sort(lst[mid_index:])
    return merge(left, right)

print(merge_sort([1, 2, 7, 8, 3, 10, 11, 4, 9, 5, 6]))
