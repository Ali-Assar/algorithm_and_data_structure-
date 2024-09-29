def bubble_sort(lst):
    for i in range(len(lst) - 1, 0, -1):
        swapped = False
        for j in range(i):
            if lst[j] > lst[j + 1]:
                lst[j], lst[j + 1] = lst[j + 1], lst[j]
                swapped = True
        if not swapped:
            break
    return lst

def selection_sort(lst):
    for i in range(len(lst) - 1):
        min_index = i
        for j in range(i + 1, len(lst)):
            if lst[j] < lst[min_index]:
                min_index = j
        if i != min_index:
            lst[i], lst[min_index] = lst[min_index], lst[i]
    return lst

def insertion_sort(lst):
    for i in range(1, len(lst)):
        current_value = lst[i]
        j = i - 1
        while j >= 0 and current_value < lst[j]:
            lst[j + 1] = lst[j]
            j -= 1
        lst[j + 1] = current_value
    return lst

print(bubble_sort([5, 8, 1, 6, 7, 4, 3]))
print(selection_sort([5, 8, 1, 6, 7, 4, 3]))
print(insertion_sort([5, 8, 1, 6, 7, 4, 3]))
